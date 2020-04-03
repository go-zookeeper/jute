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
	isPtr     bool
	classType string

	inner1 *goType
	inner2 *goType
}

func (t *goType) String() string {
	sb := strings.Builder{}

	if t.isPtr {
		sb.WriteString("*")
	}

	switch t.typeID {
	case typeBool:
		sb.WriteString("bool")
	case typeByte:
		sb.WriteString("byte")
	case typeInt32:
		sb.WriteString("int32")
	case typeInt64:
		sb.WriteString("int64")
	case typeFloat32:
		sb.WriteString("float32")
	case typeFloat64:
		sb.WriteString("float64")
	case typeString:
		sb.WriteString("string")
	case typeByteSlice:
		sb.WriteString("[]byte")
	case typeSlice:
		sb.WriteString("[]" + t.inner1.String())
	case typeMap:
		sb.WriteString("map[" + t.inner1.String() + "]" + t.inner2.String())
	case typeClass:
		sb.WriteString(t.classType)
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

func (t *goType) isPrimative() bool {
	switch t.typeID {
	case typeBool, typeByte, typeInt32, typeInt64, typeFloat32, typeFloat64, typeString, typeByteSlice:
		return true
	default:
		return false
	}
}

func (t *goType) isNillable() bool {
	switch t.typeID {
	case typeBool, typeByte, typeInt32, typeInt64, typeFloat32, typeFloat64, typeString:
		return t.isPtr
	default:
		return true
	}
}

var primTypeMap = map[parser.PTypeID]struct {
	typeID typeID
	ptr    bool
}{
	parser.BooleanTypeID: {typeBool, false},
	parser.ByteTypeID:    {typeByte, false},
	parser.IntTypeID:     {typeInt32, false},
	parser.LongTypeID:    {typeInt64, false},
	parser.FloatTypeID:   {typeFloat32, false},
	parser.DoubleTypeID:  {typeFloat64, false},
	parser.UStringTypeID: {typeString, true},
	parser.BufferTypeID:  {typeByteSlice, false},
}

func (g *generator) convertType(juteType parser.Type) (*goType, error) {
	switch t := juteType.(type) {
	case *parser.PType:
		if spec, ok := primTypeMap[t.TypeID]; ok {
			return &goType{
				typeID: spec.typeID,
				isPtr:  spec.ptr,
			}, nil
		}
		return nil, fmt.Errorf("unknown primative type %v", t.TypeID)

	case *parser.VectorType:
		innerType, err := g.convertType(t.Type)
		if err != nil {
			return nil, err
		}
		// don't point inner type
		if innerType.typeID != typeClass {
			innerType.isPtr = false
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
		keyType.isPtr = false // never pointer a key type

		valType, err := g.convertType(t.ValType)
		if err != nil {
			return nil, err
		}
		// don't point value types
		if valType.typeID != typeClass {
			valType.isPtr = false
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
			isPtr:     true,
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
	// TODO(bbennett): Since we always use pointers for class references we
	// don't really support external classes used as keys. Do we need this?
	case *parser.MapType:
		return extNamespace(t.ValType)
	}
	return ""
}
