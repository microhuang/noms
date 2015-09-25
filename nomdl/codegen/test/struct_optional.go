// This file was generated by nomdl/codegen.

package test

import (
	"github.com/attic-labs/noms/ref"
	"github.com/attic-labs/noms/types"
)

var __testPackageInFile_struct_optional_CachedRef = __testPackageInFile_struct_optional_Ref()

// This function builds up a Noms value that describes the type
// package implemented by this file and registers it with the global
// type package definition cache.
func __testPackageInFile_struct_optional_Ref() types.Ref {
	p := types.PackageDef{
		Types: types.MapOfStringToTypeRefDef{

			"OptionalStruct": __typeRefOfOptionalStruct(),
		},
	}.New()
	return types.Ref{R: types.RegisterPackage(&p)}
}

// OptionalStruct

type OptionalStruct struct {
	m types.Map
}

func NewOptionalStruct() OptionalStruct {
	return OptionalStruct{types.NewMap(
		types.NewString("$name"), types.NewString("OptionalStruct"),
		types.NewString("$type"), types.MakeTypeRef("OptionalStruct", __testPackageInFile_struct_optional_CachedRef),
	)}
}

type OptionalStructDef struct {
	S string
	B bool
}

func (def OptionalStructDef) New() OptionalStruct {
	return OptionalStruct{
		types.NewMap(
			types.NewString("$name"), types.NewString("OptionalStruct"),
			types.NewString("$type"), types.MakeTypeRef("OptionalStruct", __testPackageInFile_struct_optional_CachedRef),
			types.NewString("s"), types.NewString(def.S),
			types.NewString("b"), types.Bool(def.B),
		)}
}

func (s OptionalStruct) Def() (d OptionalStructDef) {
	if v, ok := s.m.MaybeGet(types.NewString("s")); ok {
		d.S = v.(types.String).String()
	}
	if v, ok := s.m.MaybeGet(types.NewString("b")); ok {
		d.B = bool(v.(types.Bool))
	}
	return
}

// Creates and returns a Noms Value that describes OptionalStruct.
func __typeRefOfOptionalStruct() types.TypeRef {
	return types.MakeStructTypeRef("OptionalStruct",
		[]types.Field{
			types.Field{"s", types.MakePrimitiveTypeRef(types.StringKind), true},
			types.Field{"b", types.MakePrimitiveTypeRef(types.BoolKind), true},
		},
		types.Choices{},
	)
}

func OptionalStructFromVal(val types.Value) OptionalStruct {
	// TODO: Validate here
	return OptionalStruct{val.(types.Map)}
}

func (s OptionalStruct) NomsValue() types.Value {
	return s.m
}

func (s OptionalStruct) Equals(other OptionalStruct) bool {
	return s.m.Equals(other.m)
}

func (s OptionalStruct) Ref() ref.Ref {
	return s.m.Ref()
}

func (s OptionalStruct) Type() types.TypeRef {
	return s.m.Get(types.NewString("$type")).(types.TypeRef)
}

func (s OptionalStruct) S() (v string, ok bool) {
	var vv types.Value
	if vv, ok = s.m.MaybeGet(types.NewString("s")); ok {
		v = vv.(types.String).String()
	}
	return
}

func (s OptionalStruct) SetS(val string) OptionalStruct {
	return OptionalStruct{s.m.Set(types.NewString("s"), types.NewString(val))}
}

func (s OptionalStruct) B() (v bool, ok bool) {
	var vv types.Value
	if vv, ok = s.m.MaybeGet(types.NewString("b")); ok {
		v = bool(vv.(types.Bool))
	}
	return
}

func (s OptionalStruct) SetB(val bool) OptionalStruct {
	return OptionalStruct{s.m.Set(types.NewString("b"), types.Bool(val))}
}
