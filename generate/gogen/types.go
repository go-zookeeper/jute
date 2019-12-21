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
