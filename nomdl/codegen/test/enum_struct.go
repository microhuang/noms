// This file was generated by nomdl/codegen.

package test

import (
	"github.com/attic-labs/noms/ref"
	"github.com/attic-labs/noms/types"
)

var __testPackageInFile_enum_struct_CachedRef = __testPackageInFile_enum_struct_Ref()

// This function builds up a Noms value that describes the type
// package implemented by this file and registers it with the global
// type package definition cache.
func __testPackageInFile_enum_struct_Ref() types.Ref {
	p := types.PackageDef{
		Types: types.MapOfStringToTypeRefDef{

			"EnumStruct": __typeRefOfEnumStruct(),
			"Handedness": __typeRefOfHandedness(),
		},
	}.New()
	return types.Ref{R: types.RegisterPackage(&p)}
}

// EnumStruct

type EnumStruct struct {
	m types.Map
}

func NewEnumStruct() EnumStruct {
	return EnumStruct{types.NewMap(
		types.NewString("$name"), types.NewString("EnumStruct"),
		types.NewString("$type"), types.MakeTypeRef("EnumStruct", __testPackageInFile_enum_struct_CachedRef),
		types.NewString("hand"), types.Int32(0),
	)}
}

type EnumStructDef struct {
	Hand Handedness
}

func (def EnumStructDef) New() EnumStruct {
	return EnumStruct{
		types.NewMap(
			types.NewString("$name"), types.NewString("EnumStruct"),
			types.NewString("$type"), types.MakeTypeRef("EnumStruct", __testPackageInFile_enum_struct_CachedRef),
			types.NewString("hand"), types.Int32(def.Hand),
		)}
}

func (s EnumStruct) Def() (d EnumStructDef) {
	d.Hand = Handedness(s.m.Get(types.NewString("hand")).(types.Int32))
	return
}

// Creates and returns a Noms Value that describes EnumStruct.
func __typeRefOfEnumStruct() types.TypeRef {
	return types.MakeStructTypeRef("EnumStruct",
		[]types.Field{
			types.Field{"hand", types.MakeTypeRef("Handedness", types.Ref{}), false},
		},
		types.Choices{},
	)
}

func EnumStructFromVal(val types.Value) EnumStruct {
	// TODO: Validate here
	return EnumStruct{val.(types.Map)}
}

func (s EnumStruct) NomsValue() types.Value {
	return s.m
}

func (s EnumStruct) Equals(other EnumStruct) bool {
	return s.m.Equals(other.m)
}

func (s EnumStruct) Ref() ref.Ref {
	return s.m.Ref()
}

func (s EnumStruct) Type() types.TypeRef {
	return s.m.Get(types.NewString("$type")).(types.TypeRef)
}

func (s EnumStruct) Hand() Handedness {
	return Handedness(s.m.Get(types.NewString("hand")).(types.Int32))
}

func (s EnumStruct) SetHand(val Handedness) EnumStruct {
	return EnumStruct{s.m.Set(types.NewString("hand"), types.Int32(val))}
}

// Handedness

type Handedness uint32

const (
	Right Handedness = iota
	Left
	Switch
)

// Creates and returns a Noms Value that describes Handedness.
func __typeRefOfHandedness() types.TypeRef {
	return types.MakeEnumTypeRef("Handedness", "right",
		"left",
		"switch",
	)
}
