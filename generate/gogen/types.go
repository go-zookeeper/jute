package gogen

import (
	"fmt"
	"strings"

	"github.com/go-zookeeper/jute/parser"
)

type typeID int

const (
	typeBool typeID = iota
	typeByte
	typeInt32
	typeInt64
	typeFloat32
	typeFloat64
	typeString
	typeByteSlice
	typeSlice
	typeMap
	typeClass
)

type goType struct {
	typeID    typeID
	classType string

	inner1 *goType
	inner2 *goType
}

func (t *goType) String() string {
	sb := strings.Builder{}

	switch t.typeID {
	case typeBool:
		sb.WriteString("bool")
	case typeByte:
		sb.WriteString("byte")
	case typeByteSlice:
		sb.WriteString("[]byte")
	case typeClass:
		sb.WriteString(t.classType)
	case typeFloat32:
		sb.WriteString("float32")
	case typeFloat64:
		sb.WriteString("float64")
	case typeInt32:
		sb.WriteString("int32")
	case typeInt64:
		sb.WriteString("int64")
	case typeMap:
		sb.WriteString("map[" + t.inner1.String() + "]" + t.inner2.String())
	case typeSlice:
		sb.WriteString("[]" + t.inner1.String())
	case typeString:
		sb.WriteString("string")
	default:
		panic("unknown type")
	}

	return sb.String()
}

// zeroValue will return the go zero value for the given type
func (t *goType) zeroValue() string {
	switch t.typeID {
	case typeBool:
		return "false"
	case typeByte, typeInt32, typeInt64, typeFloat32, typeFloat64:
		return "0"
	case typeClass:
		return t.classType + "{}"
	case typeString:
		return `""`
	}
	return "nil"
}

func (t *goType) methodSuffix() string {
	switch t.typeID {
	case typeBool:
		return "Boolean"
	case typeByte:
		return "Byte"
	case typeInt32:
		return "Int"
	case typeInt64:
		return "Long"
	case typeFloat32:
		return "Float"
	case typeFloat64:
		return "Double"
	case typeString:
		return "String"
	case typeByteSlice:
		return "Buffer"
	}
	panic("unknown type")
}

func (t *goType) isPrimitive() bool {
	switch t.typeID {
	case typeBool, typeByte, typeByteSlice, typeFloat32, typeFloat64, typeInt32, typeInt64, typeString:
		return true
	default:
		return false
	}
}

func (t *goType) isNillable() bool {
	switch t.typeID {
	case typeBool, typeByte, typeClass, typeFloat32, typeFloat64, typeInt32, typeInt64, typeString:
		return false
	default:
		return true
	}
}

var primTypeMap = map[parser.PTypeID]typeID{
	parser.BooleanTypeID: typeBool,
	parser.ByteTypeID:    typeByte,
	parser.IntTypeID:     typeInt32,
	parser.LongTypeID:    typeInt64,
	parser.FloatTypeID:   typeFloat32,
	parser.DoubleTypeID:  typeFloat64,
	parser.UStringTypeID: typeString,
	parser.BufferTypeID:  typeByteSlice,
}

func (g *generator) convertType(juteType parser.Type) (*goType, error) {
	switch t := juteType.(type) {
	case *parser.PType:
		if typeID, ok := primTypeMap[t.TypeID]; ok {
			return &goType{
				typeID: typeID,
			}, nil
		}
		return nil, fmt.Errorf("unknown primative type %v", t.TypeID)

	case *parser.VectorType:
		innerType, err := g.convertType(t.Type)
		if err != nil {
			return nil, err
		}
		return &goType{
			typeID: typeSlice,
			inner1: innerType,
		}, nil

	case *parser.MapType:
		keyType, err := g.convertType(t.KeyType)
		if err != nil {
			return nil, err
		}

		valType, err := g.convertType(t.ValType)
		if err != nil {
			return nil, err
		}

		return &goType{
			typeID: typeMap,
			inner1: keyType,
			inner2: valType,
		}, nil
	case *parser.ClassType:
		var prefix string
		if t.Namespace != "" {
			prefix += g.moduleMap[t.Namespace].name + "."
		}
		return &goType{
			typeID:    typeClass,
			classType: prefix + t.ClassName,
		}, nil
	}
	return nil, fmt.Errorf("unknown type %T", juteType)
}

// returns the external namespaces for the given type resolving recusively for
// map values and vector inner types.
func extNamespace(typ parser.Type) string {
	switch t := typ.(type) {
	case *parser.ClassType:
		return t.Namespace
	case *parser.VectorType:
		return extNamespace(t.Type)
	case *parser.MapType:
		return extNamespace(t.ValType)
	}
	return ""
}
