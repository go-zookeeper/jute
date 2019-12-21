package gogen

import (
	"github.com/go-zookeeper/jute/parser"
)

// jute primative types to go
var primaryTypeMap = map[parser.PTypeID]string{
	parser.BooleanTypeID: "bool",
	parser.ByteTypeID:    "byte",
	parser.IntTypeID:     "int32",
	parser.LongTypeID:    "int64",
	parser.FloatTypeID:   "float32",
	parser.DoubleTypeID:  "float64",
	parser.UStringTypeID: "string",
	parser.BufferTypeID:  "[]byte",
}

// jute primative types to the Read/Write* methods.
var primTypeName = map[parser.PTypeID]string{
	parser.BooleanTypeID: "Boolean",
	parser.ByteTypeID:    "Byte",
	parser.IntTypeID:     "Int",
	parser.LongTypeID:    "Long",
	parser.FloatTypeID:   "Float",
	parser.DoubleTypeID:  "Double",
	parser.UStringTypeID: "Ustring",
	parser.BufferTypeID:  "Buffer",
}

// returns the external namespaces for the given type resolving recusively for
// map values and vector inner types.
func extNamespace(typ parser.Type) string {
	switch t := typ.(type) {
	case *parser.ClassType:
		return t.Namespace
	case *parser.VectorType:
		return extNamespace(t.Type)
	// TODO(bbennett): Since we always use pointers for class references we
	// don't really support external classes used as keys. Do we need this?
	case *parser.MapType:
		return extNamespace(t.ValType)
	}
	return ""
}
