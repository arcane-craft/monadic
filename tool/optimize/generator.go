package main

import (
	"fmt"
	"io"
	"os"
	"slices"
	"strings"
)

func GenNewLine() string {
	return "\n"
}

func GenImport(pkgName, pkgPath string) string {
	if len(pkgName) <= 0 {
		pkgName = "."
	}
	return fmt.Sprintf("import %s \"%s\"\n", pkgName, pkgPath)
}

func GenReturn(expr string) string {
	return fmt.Sprintf("return %s", expr)
}

func GenZero(pkgName, ty string) string {
	if len(pkgName) > 0 {
		pkgName += "."
	}
	return fmt.Sprintf("%sZero[%s]()", pkgName, ty)
}

func GenBind(pkgName, monadStmt, block string) string {
	if len(pkgName) > 0 {
		pkgName += "."
	}
	return fmt.Sprintf("%sBind(%s, %s)", pkgName, monadStmt, block)
}

func GenThen(pkgName, monadStmt, block string) string {
	if len(pkgName) > 0 {
		pkgName += "."
	}
	return fmt.Sprintf("%sThen(%s, %s)", pkgName, monadStmt, block)
}

func GenBindBlock(inName, inType, outType string, stmts []string) string {
	return fmt.Sprintf("func (%s %s) %s { %s }", inName, inType, outType, strings.Join(stmts, ";"))
}

func GenThenBlock(outType string, stmts []string) string {
	return fmt.Sprintf("func () %s { %s }", outType, strings.Join(stmts, ";"))
}

type ReplaceBlock struct {
	Old Extent
	New string
}

func readExtent(reader io.ReaderAt, extent Extent) (string, error) {
	bs := make([]byte, extent.End.Offset-extent.Start.Offset)
	n, err := reader.ReadAt(bs, int64(extent.Start.Offset))
	if err != nil {
		return "", fmt.Errorf("reader.ReadAt() failed: %w", err)
	}
	if n < len(bs) {
		return "", fmt.Errorf("the lenght of read bytes is not enough")
	}
	return string(bs), nil
}

func readExtentList(reader io.ReaderAt, extents []Extent) ([]string, error) {
	var ret []string
	for _, e := range extents {
		str, err := readExtent(reader, e)
		if err != nil {
			return nil, fmt.Errorf("readExtent() failed: %w", err)
		}
		ret = append(ret, str)
	}
	return ret, nil
}

func Generate(info *FileInfo, writer io.Writer) error {
	file, err := os.Open(info.Path)
	if err != nil {
		return fmt.Errorf("os.Open() failed: %w", err)
	}
	defer file.Close()

	monadPkgName, ok := info.Imports[monadPkgPath]
	if !ok {
		return fmt.Errorf("expect import '%s' not found", monadPkgPath)
	}
	var blocks []*ReplaceBlock
	addImports := make(map[string]string)
	monadicPkgName, ok := info.Imports[monadicPkgPath]
	if !ok {
		monadicPkgName = GetRandPkgName()
		addImports[monadicPkgPath] = monadicPkgName
	}
	for _, s := range info.Syntax {
		slices.Reverse(s.Block)
		var lastStmts []string
		var finalInstanceType string
		for _, b := range s.Block {
			contBlock, err := readExtentList(file, b.PreStmts)
			if err != nil {
				return fmt.Errorf("readExtentList() failed: %w", err)
			}
			if b.CallExpr != nil {
				str, err := readExtent(file, *b.CallExpr)
				if err != nil {
					return fmt.Errorf("readExtent() failed: %w", err)
				}
				var operation string
				if b.ReturnVar != nil {
					pkgPath := GetPkgPathFromTypeStr(b.ReturnVar.Type)
					varType := b.ReturnVar.Type
					if pkgPath == info.PkgPath {
						varType = ResetTypePkgNameStr(b.ReturnVar.Type, "")
					} else if len(pkgPath) > 0 {
						pkgName, ok := info.Imports[pkgPath]
						if !ok {
							pkgName = GetRandPkgName()
							addImports[pkgPath] = pkgName
						}
						varType = ResetTypePkgNameStr(b.ReturnVar.Type, pkgName)
					}
					operation = GenReturn(GenBind(monadPkgName, str, GenBindBlock(b.ReturnVar.Name, varType, finalInstanceType, lastStmts)))
				} else {
					operation = GenReturn(GenThen(monadPkgName, str, GenThenBlock(finalInstanceType, lastStmts)))
				}
				contBlock = append(contBlock, operation)
			}
			lastStmts = contBlock
			if len(finalInstanceType) == 0 {
				pkgPath := GetPkgPathFromTypeStr(b.InstanceType)
				finalInstanceType = b.InstanceType
				if pkgPath == info.PkgPath {
					finalInstanceType = ResetTypePkgNameStr(b.InstanceType, "")
				} else if len(pkgPath) > 0 {
					pkgName, ok := info.Imports[pkgPath]
					if !ok {
						return fmt.Errorf("expect import '%s' not found", pkgPath)
					}
					finalInstanceType = ResetTypePkgNameStr(b.InstanceType, pkgName)
				}
			}
		}
		blocks = append(blocks, &ReplaceBlock{
			Old: s.Extent,
			New: GenThen(monadPkgName, GenZero(monadicPkgName, finalInstanceType), GenThenBlock(finalInstanceType, lastStmts)),
		})
	}
	slices.SortFunc(blocks, func(a, b *ReplaceBlock) int {
		return a.Old.Start.Offset - b.Old.Start.Offset
	})

	if _, err := io.CopyN(writer, file, int64(info.ImportExtent.End.Offset+1)); err != nil {
		return fmt.Errorf("io.CopyN() failed: %w", err)
	}
	for p, n := range addImports {
		if _, err := writer.Write([]byte(GenNewLine() + GenImport(n, p))); err != nil {
			return fmt.Errorf("writer.Write() failed: %w", err)
		}
	}
	lastOffset := info.ImportExtent.End.Offset + 1
	for _, b := range blocks {
		if _, err := io.CopyN(writer, file, int64(b.Old.Start.Offset-lastOffset)); err != nil {
			return fmt.Errorf("io.CopyN() failed: %w", err)
		}
		if _, err := writer.Write([]byte(b.New)); err != nil {
			return fmt.Errorf("writer.Write() failed: %w", err)
		}
		if _, err := file.Seek(int64(b.Old.End.Offset-b.Old.Start.Offset), 1); err != nil {
			return fmt.Errorf("file.Seek() failed: %w", err)
		}
		lastOffset = b.Old.End.Offset
	}
	if _, err := io.Copy(writer, file); err != nil {
		return fmt.Errorf("io.Copy() failed: %w", err)
	}
	return nil
}
