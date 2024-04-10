package main

import (
	"context"
	"log"
	"testing"
)

func TestInspectMonadTypes(t *testing.T) {
	pkgs, err := LoadPackages(context.Background(), "../..")
	if err != nil {
		t.Error("LoadPackages() failed:", err)
		return
	}
	for _, p := range pkgs {
		for _, inst := range NewMonadTypeInspector(p).InspectMonadTypes() {
			t.Log(inst)
		}
	}
}

func TestInspectDoSyntax(t *testing.T) {
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
		for _, f := range NewMonadDoSyntaxInspector(p, instances).InspectDoSyntax() {
			log.Println("path:", f.Path)
			log.Println("imports:", f.Imports)
			for _, s := range f.Syntax {
				log.Println("syntax:", s.Start, s.End)
				for _, stmt := range s.Block {
					if stmt.CallExpr != nil {
						log.Println("stmt.CallExpr:", stmt.CallExpr.Start, stmt.CallExpr.End)
					}
					log.Println("stmt.InstanceType:", stmt.InstanceType)
					if stmt.ReturnVar != nil {
						log.Println("stmt.ReturnVar:", stmt.ReturnVar.Name, stmt.ReturnVar.Type)
					}
				}
			}
		}
	}
}
