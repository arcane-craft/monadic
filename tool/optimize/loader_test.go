package main

import (
	"context"
	"testing"
)

func TestLoadPackages(t *testing.T) {
	pkgs, err := LoadPackages(context.Background(), "../..")
	if err != nil {
		t.Error("LoadPackages() failed:", err)
		return
	}
	for _, p := range pkgs {
		t.Log("Name:", p.Name)
		t.Log("PkgPath:", p.PkgPath)
		t.Log("Types:", p.Types.Name())
		t.Log("Imports:", p.Imports)
		t.Log("GoFiles:", p.GoFiles)
		for k := range p.TypesInfo.Instances {
			t.Log("TypesInfo.Instances:", k.String(), p.Fset.Position(k.Pos()).String())
		}
		t.Log()
	}
}
