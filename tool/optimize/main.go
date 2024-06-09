package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"strings"

	"golang.org/x/tools/go/packages"
)

func FormatCode(ctx context.Context, filePath string) error {
	goplsPath, err := exec.LookPath("gopls")
	if err != nil {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return fmt.Errorf("os.UserHomeDir() failed: %w", err)
		}
		goplsPath = filepath.Join(homeDir, "go", "bin", "gopls")
		if len(goplsPath) <= 0 {
			return fmt.Errorf("executable file \"gopls\" not found")
		}
	}
	cmd := exec.CommandContext(ctx, goplsPath, "format", "-w", filePath)
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("%s failed: %w", cmd.String(), err)
	}
	cmd = exec.CommandContext(ctx, goplsPath, "imports", "-w", filePath)
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("%s failed: %w", cmd.String(), err)
	}
	return nil
}

func TranslateSyntax[Type, Syntax any](
	ctx context.Context, rootDir string, firstRun bool,
	inpectTypes func(p *packages.Package) []*Type,
	inspectSyntax func(p *packages.Package, instTypes []*Type) SyntaxInspector[Syntax],
	generate func(info *FileInfo[Syntax], writer io.Writer) error,
) error {

	var finished bool
	var buildFlags []string

	for !finished {
		finished = true

		pkgs, err := LoadPackages(ctx, rootDir, buildFlags...)
		if err != nil {
			return fmt.Errorf("load source packages failed: %w", err)
		}
		var instTypes []*Type
		for _, p := range pkgs {
			instTypes = append(instTypes, inpectTypes(p)...)
			for path, dep := range p.Imports {
				if p.PkgPath != path {
					instTypes = append(instTypes, inpectTypes(dep)...)
				}
			}
		}

		var totalFiles []*FileInfo[Syntax]
		for _, p := range pkgs {
			files := NewPackageInspector(p, inspectSyntax(p, instTypes)).Inspect()
			for _, info := range files {
				func() {
					ext := filepath.Ext(info.Path)
					newFile := strings.TrimSuffix(info.Path, ext) + "_monadic_prod" + ext
					if _, err := os.Stat(newFile); err == nil {
						if !firstRun {
							return
						}
					}
					file, err := os.OpenFile(newFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
					if err != nil {
						fmt.Println("open file", newFile, "failed:", err)
						return
					}
					defer file.Close()

					err = generate(info, file)
					if err != nil {
						fmt.Println("generate code of", info.Path, "failed:", err)
						return
					}

					if err = FormatCode(ctx, newFile); err != nil {
						fmt.Println("format code of", newFile, "failed:", err)
					}

					if len(buildFlags) > 0 {
						if err := os.Remove(info.Path); err != nil {
							fmt.Println("remove file", info.Path, "failed:", err)
							return
						}
						if err := os.Rename(newFile, info.Path); err != nil {
							fmt.Println("rename file", newFile, "failed:", err)
							return
						}
					} else {
						buf := bytes.NewBuffer(nil)
						if info.BuildFlag == nil {
							buf.Write([]byte(GenBuildFlags(true)))
						}
						bs, err := os.ReadFile(info.Path)
						if err != nil {
							fmt.Println("read file", info.Path, "failed:", err)
							return
						}
						buf.Write(bs)
						tmpOldFileNam := info.Path + ".old"
						if err := os.WriteFile(tmpOldFileNam, buf.Bytes(), 0644); err != nil {
							fmt.Println("write file", tmpOldFileNam, "failed:", err)
							return
						}
						if err := os.Remove(info.Path); err != nil {
							fmt.Println("remove file", info.Path, "failed:", err)
							return
						}
						if err := os.Rename(tmpOldFileNam, info.Path); err != nil {
							fmt.Println("rename file", tmpOldFileNam, "failed:", err)
							return
						}
					}
				}()
			}
			totalFiles = append(totalFiles, files...)
		}
		if len(buildFlags) <= 0 || len(totalFiles) > 0 {
			finished = false
		}
		buildFlags = []string{"-tags=monadic_production"}
	}
	return nil
}

func main() {
	var rootDir string
	if len(os.Args) < 2 {
		pwd, err := os.Getwd()
		if err != nil {
			fmt.Println("get workspace failed:", err)
			os.Exit(2)
			return
		}
		rootDir = pwd
	} else {
		rootDir = os.Args[1]
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	err := TranslateSyntax(ctx, rootDir, true,
		func(p *packages.Package) []*DelayableInstanceType {
			return NewDelayableTypeInspector(p).InspectDelayableTypes()
		},
		func(p *packages.Package, instTypes []*DelayableInstanceType) SyntaxInspector[DelaySyntax] {
			return NewDelayableSyntaxInspector(p, instTypes)
		},
		func(info *FileInfo[DelaySyntax], writer io.Writer) error {
			return GenerateDelaySyntax(info, writer)
		},
	)
	if err != nil {
		fmt.Println("translate delay syntax failed:", err)
		return
	}

	err = TranslateSyntax(ctx, rootDir, false,
		func(p *packages.Package) []*MonadInstanceType {
			return NewMonadTypeInspector(p).InspectMonadTypes()
		},
		func(p *packages.Package, instTypes []*MonadInstanceType) SyntaxInspector[MonadDoSyntax] {
			return NewMonadDoSyntaxInspector(p, instTypes)
		},
		func(info *FileInfo[MonadDoSyntax], writer io.Writer) error {
			return GenerateMonadDoSyntax(info, writer)
		},
	)
	if err != nil {
		fmt.Println("translate monad do syntax failed:", err)
		return
	}
}
