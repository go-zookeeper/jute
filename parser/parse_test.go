package parser

import (
	"path/filepath"
	"reflect"
	"testing"
)

func TestParseFile(t *testing.T) {
	files := []string{
		"zookeeper.jute",
	}

	for _, file := range files {
		t.Run(file, func(t *testing.T) {
			_, err := ParseFile(filepath.Join("../testdata", file))
			if err != nil {
				t.Fatalf("failed to parse file: %v", err)
			}
		})
	}
}

func TestParseInclude(t *testing.T) {

	tt := []struct {
		name  string
		input string
		fail  bool
		want  string
	}{
		{"include", `include "myfile.jute"`, false, "myfile.jute"},
		{"unquoted", `include myfile.jute`, true, ""},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			p := &parser{lexer: lex(tc.input)}
			got, err := p.parseInclude()
			if err != nil && !tc.fail {
				t.Fatalf("failed to include module: %v", err)
			}

			if tc.fail {
				return
			}

			if tc.want != got {
				t.Errorf("unexpected include\nwant:\n\t%#+v\ngot:\n\t%#+v", tc.want, got)
			}
		})
	}
}

func TestParseModule(t *testing.T) {
	tt := []struct {
		name  string
		input string
		fail  bool
		want  *Module
	}{
		{"module", "module org.apache.zookeeper.data { class Id { ustring scheme; }}", false,
			&Module{
				Name: "org.apache.zookeeper.data",
				Classes: []*Class{
					&Class{
						Name: "Id",
						Fields: []*Field{
							&Field{Type: &PType{UStringTypeID}, Name: "scheme"},
						},
					},
				},
			},
		},
		{"empty module", "module data {}", false, &Module{Name: "data"}},
		{"bad indentifier", "module 1234 {}", true, nil},
		{"unterminated", "module data {", true, nil},
		{"no open brace", "module data class Data {}", true, nil},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			p := &parser{lexer: lex(tc.input)}
			got, err := p.parseModule()
			if err != nil && !tc.fail {
				t.Fatalf("failed to parse module: %v", err)
			}

			if tc.fail {
				return
			}

			if !reflect.DeepEqual(tc.want, got) {
				t.Errorf("unexpected module\nwant:\n\t%#+v\ngot:\n\t%#+v", tc.want, got)
			}
		})
	}
}

func TestParseClass(t *testing.T) {
	tt := []struct {
		name  string
		input string
		fail  bool
		want  *Class
	}{
		{"class", "class MyClass { int protocolVersion; }", false,
			&Class{Name: "MyClass", Fields: []*Field{
				&Field{Type: &PType{IntTypeID}, Name: "protocolVersion"}}}},
		{"empty class", "class MyClass {}", false,
			&Class{Name: "MyClass", Fields: nil}},
		{"unterminated class", "class MyClass {", true, nil},
		{"bad class name", "class 1234 {}", true, nil},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			p := &parser{lexer: lex(tc.input)}
			got, err := p.parseClass()
			if err != nil && !tc.fail {
				t.Fatalf("failed to parse class: %v", err)
			}

			if tc.fail {
				return
			}

			if !reflect.DeepEqual(tc.want, got) {
				t.Errorf("unexpected class\nwant:\n\t%#+v\ngot:\n\t%#+v", tc.want, got)
			}
		})
	}
}

func TestParseField(t *testing.T) {
	tt := []struct {
		name  string
		input string
		fail  bool
		want  *Field
	}{
		{"long", "long czxid;", false,
			&Field{Type: &PType{LongTypeID}, Name: "czxid"}},
		{"bad type", `"quoted" scheme;`, true, nil},
		{"bad identifier", `long 1234;`, true, nil},
		{"no semicolon", "long czxid\nlong mzxid;", true, nil},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			p := &parser{lexer: lex(tc.input)}
			got, err := p.parseField("")
			if err != nil && !tc.fail {
				t.Fatalf("failed to parse type: %v", err)
			}

			if tc.fail {
				return
			}

			if !reflect.DeepEqual(tc.want, got) {
				t.Errorf("unexpected field\nwant:\n\t%#+v\ngot:\n\t%#+v", tc.want, got)
			}
		})
	}
}
func TestParseType(t *testing.T) {
	tt := []struct {
		name  string
		input string
		fail  bool
		want  Type
	}{
		{"boolean", "boolean", false, &PType{BooleanTypeID}},
		{"int", "int", false, &PType{IntTypeID}},
		{"long", "long", false, &PType{LongTypeID}},
		{"float", "float", false, &PType{FloatTypeID}},
		{"double", "double", false, &PType{DoubleTypeID}},
		{"ustring", "ustring", false, &PType{UStringTypeID}},
		{"buffer", "buffer", false, &PType{BufferTypeID}},
		{"vector", "vector<float>", false, &VectorType{Type: &PType{FloatTypeID}}},
		{"nested vector", "vector<vector<int>>", false,
			&VectorType{Type: &VectorType{Type: &PType{IntTypeID}}}},
		{"vector bad", "vector[int]", true, nil},
		{"vector untermianted", "vector<int", true, nil},
		{"vector bad type", "vector<string>", true, nil},
		{"map", "map<int,ustring>", false, &MapType{
			KeyType: &PType{IntTypeID},
			ValType: &PType{UStringTypeID},
		}},
		{"map bad", "map[int,ustring]", true, nil},
		{"map unterminated", "map<int,ustring", true, nil},
		{"map bad type1", "map<int,string", true, nil},
		{"map bad type2", "map<string,int", true, nil},
		{"map no comma", "map<int>", true, nil},
		{"class", "MyClass", false, &ClassType{ClassName: "MyClass"}},
		{"namespaced class", "org.apache.zookeeper.data.Stat", false,
			&ClassType{Namespace: "org.apache.zookeeper.data", ClassName: "Stat"}},
		{"number", "12345", true, nil},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			p := &parser{lexer: lex(tc.input)}
			got, err := p.parseType("type definition")
			if err != nil && !tc.fail {
				t.Fatalf("failed to parse type: %v", err)
			}

			if tc.fail {
				return
			}

			if !reflect.DeepEqual(tc.want, got) {
				t.Errorf("unexpected type\nwant:\n\t%#+v\ngot:\n\t%#+v", tc.want, got)
			}
		})
	}
}
