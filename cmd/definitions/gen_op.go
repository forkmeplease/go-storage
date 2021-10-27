//go:build tools
// +build tools

package main

import (
	"github.com/Xuanwo/gg"
	"github.com/Xuanwo/templateutils"
	log "github.com/sirupsen/logrus"
)

func generateOperation(data *Data, path string) {
	f := gg.NewGroup()
	f.AddLineComment("Code generated by go generate cmd/definitions; DO NOT EDIT.")
	f.AddPackage("types")
	f.NewImport().
		AddPath("context").
		AddPath("io").
		AddPath("net/http").
		AddPath("time")

	for _, in := range data.Interfaces() {
		f.AddLineComment("%s %s", in.DisplayName(), in.Description)

		inter := f.NewInterface(in.DisplayName())
		if in.Name == "servicer" || in.Name == "storager" {
			inter.NewFunction("String").AddResult("", "string")
		}

		for _, op := range in.SortedOps() {
			pname := templateutils.ToPascal(op.Name)

			inter.AddLineComment("%s %s", pname, op.Description)
			gop := inter.NewFunction(pname)

			for _, p := range op.ParsedParams() {
				gop.AddParameter(p.Name, p.Type)
			}
			for _, r := range op.ParsedResults() {
				gop.AddResult(r.Name, r.Type)
			}

			// We need to generate XxxWithContext functions if not local.
			if !op.Local {
				inter.AddLineComment("%sWithContext %s", pname, op.Description)
				gop := inter.NewFunction(pname + "WithContext")

				// Insert context param.
				gop.AddParameter("ctx", "context.Context")
				for _, p := range op.ParsedParams() {
					gop.AddParameter(p.Name, p.Type)
				}
				for _, r := range op.ParsedResults() {
					gop.AddResult(r.Name, r.Type)
				}
			}
			// Insert an empty for different functions.
			inter.AddLine()
		}
		inter.NewFunction("mustEmbedUnimplemented" + in.DisplayName())

		stubName := "Unimplemented" + in.DisplayName()
		f.AddLineComment("%s must be embedded to have forward compatible implementations.", stubName).
			AddLineComment("").
			AddLineComment("%s %s", in.DisplayName(), in.Description)
		f.NewStruct(stubName)

		f.NewFunction("mustEmbedUnimplemented"+in.DisplayName()).
			WithReceiver("s", stubName).
			AddBody(gg.Line())
		f.NewFunction("String").
			WithReceiver("s", stubName).
			AddResult("", "string").
			AddBody(
				gg.Return(gg.Lit(stubName)))

		for _, op := range in.SortedOps() {
			pname := templateutils.ToPascal(op.Name)

			gop := f.NewFunction(pname).
				WithReceiver("s", stubName)

			for _, p := range op.ParsedParams() {
				gop.AddParameter(p.Name, p.Type)
			}
			for _, r := range op.ParsedResults() {
				gop.AddResult(r.Name, r.Type)
			}
			// If not local, we need to set error
			if !op.Local {
				gop.AddBody(gg.S(`err = NewOperationNotImplementedError("%s")`, op.Name))
			}
			gop.AddBody(gg.Return())

			// We need to generate XxxWithContext functions if not local.
			if !op.Local {
				gop := f.NewFunction(pname+"WithContext").
					WithReceiver("s", stubName)

				// Insert context param.
				gop.AddParameter("ctx", "context.Context")
				for _, p := range op.ParsedParams() {
					gop.AddParameter(p.Name, p.Type)
				}
				for _, r := range op.ParsedResults() {
					gop.AddResult(r.Name, r.Type)
				}
				gop.AddBody(
					gg.S(`err = NewOperationNotImplementedError("%s")`, op.Name),
					gg.Return(),
				)
			}
		}
	}

	err := f.WriteFile(path)
	if err != nil {
		log.Fatalf("generate to %s: %v", path, err)
	}
}
