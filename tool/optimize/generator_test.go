package main

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestGenerate(t *testing.T) {
	pkgs, err := LoadPackages(context.Background(), "../..")
	if err != nil {
		t.Error("LoadPackages() failed:", err)
		return
	}
	var instances []*MonadInstanceType
	for _, p := range pkgs {
		instances = append(instances, NewMonadTypeInspector(p).InspectMonadTypes()...)
	}
	for _, p := range pkgs {
		files := NewMonadDoSyntaxInspector(p, instances).InspectDoSyntax()
		for _, info := range files {
			func() {
				ext := filepath.Ext(info.Path)
				file, err := os.OpenFile(strings.TrimSuffix(info.Path, ext)+"_monadic_prod"+ext, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
				if err != nil {
					t.Errorf("os.Open() failed: %s", err)
					return
				}
				defer file.Close()

				err = Generate(info, file)
				if err != nil {
					t.Errorf("Generate() failed: %s", err)
					return
				}
			}()
		}
	}
}
