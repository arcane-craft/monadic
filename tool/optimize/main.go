package main

import (
	"bytes"
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

	var finished bool
	var buildFlags []string
	for !finished {
		finished = true
		pkgs, err := LoadPackages(context.Background(), rootDir, buildFlags...)
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
					newFile := strings.TrimSuffix(info.Path, ext) + "_monadic_prod" + ext
					file, err := os.OpenFile(newFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
					if err != nil {
						fmt.Println("open file", newFile, "failed:", err)
						return
					}
					defer file.Close()

					err = Generate(info, file)
					if err != nil {
						fmt.Println("generate code of", info.Path, "failed:", err)
						return
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
			if len(files) > 0 {
				finished = false
			}
		}
		buildFlags = []string{"-tags=monadic_production"}
	}
}
