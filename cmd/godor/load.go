package main

import (
	"errors"
	"fmt"
	"go/types"
	"log"

	"github.com/fatih/structtag"
	"golang.org/x/tools/go/packages"
)

type loadedStruct struct {
	pkg    string
	name   string
	fields []*loadedField
}

func getLoadedStruct(structs map[string]*loadedStruct, pkgPath, structName string) *loadedStruct {
	key := pkgPath + "." + structName
	s, ok := structs[key]
	if !ok {
		s = &loadedStruct{pkgPath, structName, []*loadedField{}}
		structs[key] = s
	}

	return s
}

type loadedField struct {
	name    string
	goType  types.Type
	options []string
	tags    *structtag.Tags
}

func newLoadedField(field *types.Var, options []string, tags *structtag.Tags) *loadedField {
	return &loadedField{field.Name(), field.Type(), options, tags}
}

func load(path string) []*loadedStruct {
	pkgs, err := packages.Load(
		&packages.Config{
			Mode: packages.NeedDeps | packages.NeedTypes,
		},
		path,
	)
	if err != nil || packages.PrintErrors(pkgs) > 0 {
		log.Fatal(err)
	}

	mapStructs := map[string]*loadedStruct{}
	const tagName = "godor"

	for _, pkg := range pkgs {
		for _, scope := range pkg.Types.Scope().Names() {
			object := pkg.Types.Scope().Lookup(scope)
			structObject, ok := object.Type().Underlying().(*types.Struct)
			if !ok {
				continue
			}

			pkgPath := object.Pkg().Path()
			structName := object.Name()
			if !object.Exported() {
				log.Printf("%s.%s struct is unexported, skipping", pkgPath, structName)
				continue
			}

			log.Printf("%s.%s", pkgPath, structName)
			numFields := structObject.NumFields()
			for i := 0; i < numFields; i++ {
				field := structObject.Field(i)
				if !field.Exported() {
					log.Printf("\t%s field is unexported, skipping", field.Name())
					continue
				}

				tags, tagErr := structtag.Parse(structObject.Tag(i))
				if tagErr != nil {
					log.Printf("\t%s field has invalid tag: %s", field.Name(), tags)
					err = errors.Join(err, fmt.Errorf("[%s.%s.%s] has invalid tag: %s", pkgPath, structName, field.Name(), tags))
					continue
				}

				tag, tagErr := tags.Get(tagName)
				if tagErr != nil {
					log.Printf("\t%s field has no %s tag, skipping", field.Name(), tagName)
					continue
				}

				log.Printf("\t%s field has %s tag: %s", field.Name(), tagName, tag)
				tags.Delete(tagName)
				loadedField := newLoadedField(field, append(tag.Options, tag.Name), tags)
				loadedStruct := getLoadedStruct(mapStructs, pkgPath, structName)
				loadedStruct.fields = append(loadedStruct.fields, loadedField)
			}
		}
	}

	if err != nil {
		log.Fatal(err)
	}

	loadedStructs := []*loadedStruct{}
	for _, ls := range mapStructs {
		loadedStructs = append(loadedStructs, ls)
	}

	return loadedStructs
}
