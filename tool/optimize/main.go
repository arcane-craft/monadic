package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

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

	pkgs, err := LoadPackages(context.Background(), rootDir)
	if err != nil {
		fmt.Println("load source packages failed:", err)
		os.Exit(2)
		return
	}
	var instances []*MonadInstanceType
	for _, p := range pkgs {
		instances = append(instances, NewMonadTypeInspector(p).InspectMonadTypes()...)
		for path, dep := range p.Imports {
			if p.PkgPath != path {
				instances = append(instances, NewMonadTypeInspector(dep).InspectMonadTypes()...)
			}
		}
	}
	for _, p := range pkgs {
		files := NewMonadDoSyntaxInspector(p, instances).InspectDoSyntax()
		for _, info := range files {
			func() {
				ext := filepath.Ext(info.Path)
				newFile := strings.TrimSuffix(info.Path, ext) + "_desugar" + ext
				file, err := os.OpenFile(newFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
				if err != nil {
					fmt.Println("open file", newFile, "failed:", err)
					os.Exit(2)
					return
				}
				defer file.Close()

				err = Generate(info, file)
				if err != nil {
					fmt.Println("generate code of", info.Path, "failed:", err)
					os.Exit(2)
					return
				}
				err = os.Rename(info.Path, info.Path+".origin")
				if err != nil {
					fmt.Println("rename origin code file", info.Path, "failed:", err)
					os.Exit(2)
					return
				}
			}()
		}
	}
}
