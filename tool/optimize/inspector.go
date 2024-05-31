package main

import (
	"go/ast"
	"go/token"
	"path"
	"strings"

	"golang.org/x/tools/go/ast/inspector"
	"golang.org/x/tools/go/packages"
)

const (
	monadicPkgPath     = "github.com/arcane-craft/monadic"
	monadPkgPath       = "github.com/arcane-craft/monadic/monad"
	applicativePkgPath = "github.com/arcane-craft/monadic/applicative"
	implMonadFun       = "ImplMonadDo"
	monadDoFun         = "Do"
	monadDoExtractFun  = "X"
)

type MonadTypeInspector struct {
	pkg *packages.Package
}

func NewMonadTypeInspector(pkg *packages.Package) *MonadTypeInspector {
	return &MonadTypeInspector{pkg}
}

type MonadInstanceType struct {
	Name   string
	ArgLen int
}

func (i *MonadTypeInspector) inspectMonadInstanceType(x ast.Expr) *MonadInstanceType {
	ident, ok := x.(*ast.Ident)
	if ok {
		if instance, ok := i.pkg.TypesInfo.Instances[ident]; ok {
			return &MonadInstanceType{
				Name:   GetPkgPathFromType(instance.Type) + "." + ident.Name,
				ArgLen: instance.TypeArgs.Len(),
			}
		}
	}
	return nil
}

func (i *MonadTypeInspector) inspectMonadImpl(implFun *ast.Ident, targetType ast.Expr) *MonadInstanceType {
	if selObj := i.pkg.TypesInfo.ObjectOf(implFun); selObj != nil {
		if selPkg := selObj.Pkg(); selPkg != nil &&
			selPkg.Path() == monadPkgPath &&
			implFun.Name == implMonadFun {
			switch index := targetType.(type) {
			case *ast.IndexExpr:
				return i.inspectMonadInstanceType(index.X)
			case *ast.IndexListExpr:
				return i.inspectMonadInstanceType(index.X)
			}
		}
	}
	return nil
}

func (i *MonadTypeInspector) InspectMonadTypes() []*MonadInstanceType {
	var ret []*MonadInstanceType
	ins := inspector.New(i.pkg.Syntax)
	ins.Preorder([]ast.Node{
		&ast.GenDecl{},
	}, func(n ast.Node) {
		decl := n.(*ast.GenDecl)
		if decl.Tok.String() == ast.Var.String() {
			for _, spec := range decl.Specs {
				valueSpec, ok := spec.(*ast.ValueSpec)
				if ok {
					for _, value := range valueSpec.Values {
						call, ok := value.(*ast.CallExpr)
						if ok {
							index, ok := call.Fun.(*ast.IndexExpr)
							if ok {
								switch fun := index.X.(type) {
								case *ast.Ident:
									inst := i.inspectMonadImpl(fun, index.Index)
									if inst != nil {
										ret = append(ret, inst)
									}
								case *ast.SelectorExpr:
									inst := i.inspectMonadImpl(fun.Sel, index.Index)
									if inst != nil {
										ret = append(ret, inst)
									}
								}
							}
						}
					}
				}
			}
		}
	})
	return ret
}

type MonadDoSyntaxInspector struct {
	pkg           *packages.Package
	instanceTypes map[string]*MonadInstanceType
	imports       map[string]map[string]string
	importExtents map[string]Extent
	buildFlags    map[string]Extent
}

func NewMonadDoSyntaxInspector(pkg *packages.Package, instances []*MonadInstanceType) *MonadDoSyntaxInspector {
	imports := make(map[string]map[string]string)
	importExtents := make(map[string]Extent)
	buildFlags := make(map[string]Extent)
	for _, file := range pkg.Syntax {
		for _, cg := range file.Comments {
			for _, c := range cg.List {
				if strings.Contains(c.Text, "//go:build !monadic_production") || strings.Contains(c.Text, "//go:build monadic_production") {
					fileName := pkg.Fset.Position(c.Pos()).Filename
					buildFlags[fileName] = Extent{
						Start: pkg.Fset.Position(c.Pos()),
						End:   pkg.Fset.Position(c.End()),
					}
				}
			}
		}
		specs := make(map[string]string)
		pkgEndPos := pkg.Fset.Position(file.Name.End())
		importExtents[pkgEndPos.Filename] = Extent{
			Start: pkgEndPos,
			End:   pkgEndPos,
		}
		for _, spec := range file.Imports {
			importPath := strings.Trim(spec.Path.Value, "\"")
			importName := path.Base(importPath)
			if spec.Name != nil {
				importName = spec.Name.Name
				if importName == "." {
					importName = ""
				}
			}
			specs[importPath] = importName
		}
		fileName := pkg.Fset.Position(file.Pos()).Filename
		imports[fileName] = specs
	}
	instanceTypes := make(map[string]*MonadInstanceType)
	for _, inst := range instances {
		instanceTypes[inst.Name] = inst
	}
	return &MonadDoSyntaxInspector{
		pkg:           pkg,
		instanceTypes: instanceTypes,
		imports:       imports,
		importExtents: importExtents,
		buildFlags:    buildFlags,
	}
}

type Extent struct {
	Start token.Position
	End   token.Position
}

type Variable struct {
	Name string
	Type string
}

type MonadStmt struct {
	PreStmts          []Extent
	AnonymousCallExpr *Extent
	CallExpr          *Extent
	InstanceType      string
	ReturnVar         *Variable
}

type MonadDoSyntax struct {
	Extent
	Block   []*MonadStmt
	RetType Extent
}

type FileInfo struct {
	Path         string
	BuildFlag    *Extent
	PkgPath      string
	Imports      map[string]string
	ImportExtent Extent
	Syntax       []*MonadDoSyntax
}

func (i *MonadDoSyntaxInspector) isMonadDoFun(funIdent *ast.Ident) bool {
	if doFunObj := i.pkg.TypesInfo.ObjectOf(funIdent); doFunObj != nil {
		if doFunPkg := doFunObj.Pkg(); doFunPkg != nil {
			if doFunPkg.Path() == monadPkgPath && funIdent.Name == monadDoFun {
				return true
			}
		}
	}
	return false
}

func (i *MonadDoSyntaxInspector) inspectDoBlock(block *ast.BlockStmt) []*MonadStmt {
	var monadStmts []*MonadStmt
	ms := new(MonadStmt)
	for _, stmt := range block.List {
		var isExtraction bool
		switch s := stmt.(type) {
		case *ast.AssignStmt:
			if s.Tok == token.DEFINE && len(s.Rhs) == 1 && len(s.Lhs) == 1 {
				call, ok := s.Rhs[0].(*ast.CallExpr)
				if ok {
					sel, ok := call.Fun.(*ast.SelectorExpr)
					if ok && sel.Sel.Name == monadDoExtractFun {
						instanceType := i.pkg.TypesInfo.TypeOf(sel.X).String()
						if _, ok := i.instanceTypes[GetNameFromTypeStr(instanceType)]; ok {
							ms.InstanceType = instanceType
							ms.CallExpr = &Extent{
								Start: i.pkg.Fset.Position(sel.X.Pos()),
								End:   i.pkg.Fset.Position(sel.X.End()),
							}
							ms.ReturnVar = &Variable{
								Name: s.Lhs[0].(*ast.Ident).Name,
								Type: i.pkg.TypesInfo.TypeOf(s.Lhs[0]).String(),
							}
							monadStmts = append(monadStmts, ms)
							ms = new(MonadStmt)
							isExtraction = true
						}
					}
				}
			}
		case *ast.ExprStmt:
			call, ok := s.X.(*ast.CallExpr)
			if ok {
				sel, ok := call.Fun.(*ast.SelectorExpr)
				if ok && sel.Sel.Name == monadDoExtractFun {
					instanceType := i.pkg.TypesInfo.TypeOf(sel.X).String()
					if _, ok := i.instanceTypes[GetNameFromTypeStr(instanceType)]; ok {
						ms.InstanceType = instanceType
						ms.CallExpr = &Extent{
							Start: i.pkg.Fset.Position(sel.X.Pos()),
							End:   i.pkg.Fset.Position(sel.X.End()),
						}
						monadStmts = append(monadStmts, ms)
						ms = new(MonadStmt)
						isExtraction = true
					}
				}
			}
		}
		if !isExtraction {
			ast.Inspect(stmt, func(n ast.Node) bool {
				if n != nil {
					switch node := n.(type) {
					case *ast.CallExpr:
						switch fun := node.Fun.(type) {
						case *ast.Ident:
							if i.isMonadDoFun(fun) {
								return false
							}
						case *ast.SelectorExpr:
							if i.isMonadDoFun(fun.Sel) {
								return false
							}
							sel := fun
							if sel.Sel.Name == monadDoExtractFun {
								instanceType := i.pkg.TypesInfo.TypeOf(sel.X).String()
								if _, ok := i.instanceTypes[GetNameFromTypeStr(instanceType)]; ok {
									ms.InstanceType = instanceType
									ms.CallExpr = &Extent{
										Start: i.pkg.Fset.Position(sel.X.Pos()),
										End:   i.pkg.Fset.Position(sel.X.End()),
									}
									ms.AnonymousCallExpr = &Extent{
										Start: i.pkg.Fset.Position(node.Pos()),
										End:   i.pkg.Fset.Position(node.End()),
									}
									ms.ReturnVar = &Variable{
										Name: GetRandVarName(),
										Type: i.pkg.TypesInfo.TypeOf(node).String(),
									}
									monadStmts = append(monadStmts, ms)
									ms = new(MonadStmt)
									return false
								}
							}
						}
					}
				}
				return true
			})

			ms.PreStmts = append(ms.PreStmts, Extent{
				Start: i.pkg.Fset.Position(stmt.Pos()),
				End:   i.pkg.Fset.Position(stmt.End()),
			})
			retStmt, ok := stmt.(*ast.ReturnStmt)
			if ok && len(retStmt.Results) == 1 {
				ms.InstanceType = i.pkg.TypesInfo.TypeOf(retStmt.Results[0]).String()
			}
		}
	}
	return append(monadStmts, ms)
}

func (i *MonadDoSyntaxInspector) inspectDoFunCall(funExpr ast.Expr, args []ast.Expr) ([]*MonadStmt, Extent) {
	var (
		stmts   []*MonadStmt
		retType Extent
	)
	var isDoFunCall bool
	switch fun := funExpr.(type) {
	case *ast.Ident:
		isDoFunCall = i.isMonadDoFun(fun)
	case *ast.SelectorExpr:
		isDoFunCall = i.isMonadDoFun(fun.Sel)
	case *ast.IndexExpr:
		s, r := i.inspectDoFunCall(fun.X, args)
		stmts = append(stmts, s...)
		retType = r
	}
	if isDoFunCall {
		if len(args) == 1 {
			funLit, ok := args[0].(*ast.FuncLit)
			if ok {
				retType = Extent{
					Start: i.pkg.Fset.Position(funLit.Type.Results.Pos()),
					End:   i.pkg.Fset.Position(funLit.Type.Results.End()),
				}
				stmts = append(stmts, i.inspectDoBlock(funLit.Body)...)
				return stmts, retType
			}
		}
	}
	return stmts, retType
}

func (i *MonadDoSyntaxInspector) InspectDoSyntax() []*FileInfo {
	fileMap := make(map[string]*FileInfo)
	ins := inspector.New(i.pkg.Syntax)
	ins.Nodes([]ast.Node{
		&ast.GenDecl{},
		&ast.CallExpr{},
	}, func(n ast.Node, _ bool) bool {
		switch node := n.(type) {
		case *ast.GenDecl:
			if node.Tok == token.IMPORT {
				end := i.pkg.Fset.Position(node.End())
				extent := i.importExtents[end.Filename]
				if extent.End.Offset < end.Offset {
					extent.End = end
					i.importExtents[end.Filename] = extent
				}
			}
			return false
		case *ast.CallExpr:
			block, retType := i.inspectDoFunCall(node.Fun, node.Args)
			if len(block) > 0 {
				fileName := i.pkg.Fset.Position(n.Pos()).Filename
				file := fileMap[fileName]
				if file == nil {
					file = &FileInfo{
						Path:         fileName,
						PkgPath:      i.pkg.PkgPath,
						Imports:      i.imports[fileName],
						ImportExtent: i.importExtents[fileName],
					}
					fileMap[fileName] = file
				}
				file.Syntax = append(file.Syntax, &MonadDoSyntax{
					Extent: Extent{
						Start: i.pkg.Fset.Position(node.Pos()),
						End:   i.pkg.Fset.Position(node.End()),
					},
					Block:   block,
					RetType: retType,
				})
				return false
			}
		}
		return true
	})
	var ret []*FileInfo
	for _, f := range fileMap {
		if extent, ok := i.buildFlags[f.Path]; ok {
			f.BuildFlag = &extent
		}
		ret = append(ret, f)
	}
	return ret
}
