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
	monadDoFFun        = "DoF"
	monadDoExtractFun  = "X"
	lazyPkgPath        = "github.com/arcane-craft/monadic/lazy"
	implDelayableFun   = "ImplDelayable"
	delayFun           = "Delay"
)

type DelayableTypeInspector struct {
	pkg *packages.Package
}

func NewDelayableTypeInspector(pkg *packages.Package) *DelayableTypeInspector {
	return &DelayableTypeInspector{pkg}
}

type DelayableInstanceType struct {
	Name   string
	ArgLen int
}

func (i *DelayableTypeInspector) inspectDelayableInstanceType(x ast.Expr) *DelayableInstanceType {
	ident, ok := x.(*ast.Ident)
	if ok {
		if instance, ok := i.pkg.TypesInfo.Instances[ident]; ok {
			return &DelayableInstanceType{
				Name:   GetPkgPathFromType(instance.Type) + "." + ident.Name,
				ArgLen: instance.TypeArgs.Len(),
			}
		}
	}
	return nil
}

func (i *DelayableTypeInspector) inspectDelayableImpl(implFun *ast.Ident, targetType ast.Expr) *DelayableInstanceType {
	if selObj := i.pkg.TypesInfo.ObjectOf(implFun); selObj != nil {
		if selPkg := selObj.Pkg(); selPkg != nil &&
			selPkg.Path() == lazyPkgPath &&
			implFun.Name == implDelayableFun {
			switch index := targetType.(type) {
			case *ast.IndexExpr:
				return i.inspectDelayableInstanceType(index.X)
			case *ast.IndexListExpr:
				return i.inspectDelayableInstanceType(index.X)
			}
		}
	}
	return nil
}

func (i *DelayableTypeInspector) InspectDelayableTypes() []*DelayableInstanceType {
	var ret []*DelayableInstanceType
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
									inst := i.inspectDelayableImpl(fun, index.Index)
									if inst != nil {
										ret = append(ret, inst)
									}
								case *ast.SelectorExpr:
									inst := i.inspectDelayableImpl(fun.Sel, index.Index)
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

type DelayableSyntaxInspector struct {
	pkg           *packages.Package
	instanceTypes map[string]*DelayableInstanceType
}

func NewDelayableSyntaxInspector(pkg *packages.Package, instances []*DelayableInstanceType) *DelayableSyntaxInspector {
	instanceTypes := make(map[string]*DelayableInstanceType)
	for _, inst := range instances {
		instanceTypes[inst.Name] = inst
	}
	return &DelayableSyntaxInspector{
		pkg:           pkg,
		instanceTypes: instanceTypes,
	}
}

type DelaySyntax struct {
	Extent
	Expr     Extent
	ExprType string
}

func (i *DelayableSyntaxInspector) isDelayFun(funIdent *ast.Ident) bool {
	if delayFunObj := i.pkg.TypesInfo.ObjectOf(funIdent); delayFunObj != nil {
		if delayFunPkg := delayFunObj.Pkg(); delayFunPkg != nil {
			if delayFunPkg.Path() == lazyPkgPath && funIdent.Name == delayFun {
				return true
			}
		}
	}
	return false
}

func (i *DelayableSyntaxInspector) isDelayMethod(sel *ast.SelectorExpr) bool {
	if sel.Sel.Name == delayFun {
		instanceType := i.pkg.TypesInfo.TypeOf(sel.X).String()
		if _, ok := i.instanceTypes[GetNameFromTypeStr(instanceType)]; ok {
			return true
		}
	}
	return false
}

func (i *DelayableSyntaxInspector) inspectDelayFnCall(funExpr ast.Expr, args []ast.Expr) (*Extent, string) {
	var isDelayFunCall bool
	var expr *Extent
	var exprType string
	switch fun := funExpr.(type) {
	case *ast.Ident:
		if len(args) == 1 {
			isDelayFunCall = i.isDelayFun(fun)
		}
	case *ast.SelectorExpr:
		if len(args) == 1 {
			isDelayFunCall = i.isDelayFun(fun.Sel)
		} else {
			isDelayFunCall = i.isDelayMethod(fun)
		}
	case *ast.IndexExpr:
		expr, exprType = i.inspectDelayFnCall(fun.X, args)
	}
	if isDelayFunCall {
		if len(args) == 1 {
			expr = &Extent{
				Start: i.pkg.Fset.Position(args[0].Pos()),
				End:   i.pkg.Fset.Position(args[0].End()),
			}
			exprType = i.pkg.TypesInfo.TypeOf(args[0]).String()
		} else if sel, ok := funExpr.(*ast.SelectorExpr); ok {
			expr = &Extent{
				Start: i.pkg.Fset.Position(sel.X.Pos()),
				End:   i.pkg.Fset.Position(sel.X.End()),
			}
			exprType = i.pkg.TypesInfo.TypeOf(sel.X).String()
		}
	}
	return expr, exprType
}

func (i *DelayableSyntaxInspector) InspectSyntax(callExpr *ast.CallExpr) *DelaySyntax {
	expr, exprType := i.inspectDelayFnCall(callExpr.Fun, callExpr.Args)
	if expr != nil {
		return &DelaySyntax{
			Extent: Extent{
				Start: i.pkg.Fset.Position(callExpr.Pos()),
				End:   i.pkg.Fset.Position(callExpr.End()),
			},
			Expr:     *expr,
			ExprType: exprType,
		}
	}
	return nil
}

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
	Block    []*MonadStmt
	FuncType *Extent
	RetType  Extent
}

func (i *MonadDoSyntaxInspector) isMonadDoFun(funIdent *ast.Ident) bool {
	if doFunObj := i.pkg.TypesInfo.ObjectOf(funIdent); doFunObj != nil {
		if doFunPkg := doFunObj.Pkg(); doFunPkg != nil {
			if doFunPkg.Path() == monadPkgPath &&
				(funIdent.Name == monadDoFun || strings.HasPrefix(funIdent.Name, monadDoFFun)) {
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
										Name: GetRandVarName(i.pkg.Fset.Position(node.Pos()).String()),
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

func (i *MonadDoSyntaxInspector) inspectDoFunCall(funExpr ast.Expr, args []ast.Expr) ([]*MonadStmt, *Extent, Extent) {
	var (
		stmts    []*MonadStmt
		funcType *Extent
		retType  Extent
	)
	var isDoFunCall bool
	var funcName string
	switch fun := funExpr.(type) {
	case *ast.Ident:
		isDoFunCall = i.isMonadDoFun(fun)
		funcName = fun.Name
	case *ast.SelectorExpr:
		isDoFunCall = i.isMonadDoFun(fun.Sel)
		funcName = fun.Sel.Name
	case *ast.IndexExpr:
		s, f, r := i.inspectDoFunCall(fun.X, args)
		stmts = append(stmts, s...)
		funcType = f
		retType = r
	}
	if isDoFunCall {
		if len(args) == 1 {
			funLit, ok := args[0].(*ast.FuncLit)
			if ok {
				if strings.HasPrefix(funcName, monadDoFFun) {
					funcType = &Extent{
						Start: i.pkg.Fset.Position(funLit.Type.Pos()),
						End:   i.pkg.Fset.Position(funLit.Type.End()),
					}
				}
				retType = Extent{
					Start: i.pkg.Fset.Position(funLit.Type.Results.Pos()),
					End:   i.pkg.Fset.Position(funLit.Type.Results.End()),
				}
				stmts = append(stmts, i.inspectDoBlock(funLit.Body)...)
				return stmts, funcType, retType
			}
		}
	}
	return stmts, funcType, retType
}

func (i *MonadDoSyntaxInspector) InspectSyntax(callExpr *ast.CallExpr) *MonadDoSyntax {
	block, funcType, retType := i.inspectDoFunCall(callExpr.Fun, callExpr.Args)
	if len(block) > 0 {
		return &MonadDoSyntax{
			Extent: Extent{
				Start: i.pkg.Fset.Position(callExpr.Pos()),
				End:   i.pkg.Fset.Position(callExpr.End()),
			},
			Block:    block,
			FuncType: funcType,
			RetType:  retType,
		}
	}
	return nil
}

type SyntaxInspector[Syntax any] interface {
	InspectSyntax(callExpr *ast.CallExpr) *Syntax
}

type PackageInspector[Syntax any] struct {
	pkg           *packages.Package
	imports       map[string]map[string]string
	importExtents map[string]Extent
	buildFlags    map[string]Extent

	inspector SyntaxInspector[Syntax]
}

func NewPackageInspector[Syntax any](pkg *packages.Package, inpector SyntaxInspector[Syntax]) *PackageInspector[Syntax] {
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
	return &PackageInspector[Syntax]{
		pkg:           pkg,
		imports:       imports,
		importExtents: importExtents,
		buildFlags:    buildFlags,
		inspector:     inpector,
	}
}

type FileInfo[Syntax any] struct {
	Path         string
	BuildFlag    *Extent
	PkgPath      string
	Imports      map[string]string
	ImportExtent Extent
	Syntax       []*Syntax
}

func (i *PackageInspector[Syntax]) Inspect() []*FileInfo[Syntax] {
	fileMap := make(map[string]*FileInfo[Syntax])
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
			syntax := i.inspector.InspectSyntax(node)
			if syntax != nil {
				fileName := i.pkg.Fset.Position(n.Pos()).Filename
				file := fileMap[fileName]
				if file == nil {
					file = &FileInfo[Syntax]{
						Path:         fileName,
						PkgPath:      i.pkg.PkgPath,
						Imports:      i.imports[fileName],
						ImportExtent: i.importExtents[fileName],
					}
					fileMap[fileName] = file
				}
				file.Syntax = append(file.Syntax, syntax)
				return false
			}
		}
		return true
	})
	var ret []*FileInfo[Syntax]
	for _, f := range fileMap {
		if extent, ok := i.buildFlags[f.Path]; ok {
			f.BuildFlag = &extent
		}
		ret = append(ret, f)
	}
	return ret
}
