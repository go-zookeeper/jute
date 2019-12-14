package parser

// Doc is a single jute/jt file and lists the included jute files and
// modules.
type Doc struct {
	Includes []string
	Modules  []*Module
}

// Module defines a namespace of classes.
type Module struct {
	Name    string
	Classes []*Class
}

// Class is a grouping of fields.
type Class struct {
	Name   string
	Fields []*Field
}

// Field is a field inside a class.
type Field struct {
	Type Type
	Name string
}

// Type represents a the type for a field.
type Type interface {
	fieldType()
}

func (*PType) fieldType()      {}
func (*VectorType) fieldType() {}
func (*MapType) fieldType()    {}
func (*ClassType) fieldType()  {}

// PTypeID is an identifier for primative types supported
type PTypeID int

// IDs of primative types
const (
	BooleanTypeID PTypeID = iota + 1
	ByteTypeID
	IntTypeID
	LongTypeID
	FloatTypeID
	DoubleTypeID
	UStringTypeID
	BufferTypeID
)

// PType is a primative type (bool, byte, int, etc)
type PType struct {
	TypeID PTypeID
}

// VectorType is a list of a certain primative type.
type VectorType struct {
	Type Type
}

// MapType is a mapping of a key type to a value type.
type MapType struct {
	KeyType Type
	ValType Type
}

// ClassType represents a reference to another class defined. Namespace is the
// module of where the class is defined.  If Namespace is empty then the class
// is found in the current module.
type ClassType struct {
	Namespace string
	ClassName string
}
