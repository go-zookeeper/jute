package parser

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type parser struct {
	name  string
	lexer *lexer

	tokens    [2]token
	peekCount int
	namespace string
}

func Parse(name, input string) (*Doc, error) {
	p := &parser{
		name:  name,
		lexer: lex(input),
	}
	return p.parse()
}

func ParseFile(filename string) (*Doc, error) {
	p, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}
	return Parse(filename, string(p))
}

func (p *parser) errorf(format string, args ...interface{}) error {
	format = fmt.Sprintf("%s: %s", p.name, format)
	return fmt.Errorf(format, args...)
}

func (p *parser) next() token {
	if p.peekCount > 0 {
		p.peekCount--
	} else {
		p.tokens[0] = p.nextToken()
	}
	return p.tokens[p.peekCount]
}

func (p *parser) peek() token {
	if p.peekCount > 0 {
		return p.tokens[p.peekCount-1]
	}
	p.peekCount = 1
	p.tokens[0] = p.nextToken()
	return p.tokens[0]

}

func (p *parser) nextToken() token {
	t := p.lexer.nextToken()
	for {
		// Skip comments for now
		if t.typ != tokenBlockComment && t.typ != tokenLineComment {
			break
		}
		t = p.lexer.nextToken()
	}
	return t
}

func (p *parser) expect(typ tokenType, context string) (token, error) {
	t := p.next()
	if t.typ != typ {
		return token{}, p.errorf("unexpected %s in %s (want %s)", t, context, typ)
	}
	return t, nil
}

func (p *parser) unexpected(t token, context string) error {
	return p.errorf("unexpected %s in %s", t, context)
}

func (p *parser) parse() (*Doc, error) {
	doc := &Doc{}
Loop:
	for {
		switch t := p.peek(); {
		case t.typ == tokenEOF:
			break Loop
		case t.typ == tokenInclude:
			include, err := p.parseInclude()
			if err != nil {
				return nil, err
			}

			doc.Includes = append(doc.Includes, include)
		case t.typ == tokenModule:
			module, err := p.parseModule()
			if err != nil {
				return nil, err
			}

			doc.Modules = append(doc.Modules, module)
		default:
			return nil, p.errorf("unexpected %s, expected 'module' or 'include'", t)
		}
	}

	return doc, nil
}

func (p *parser) parseInclude() (string, error) {
	context := "include statement"
	_, err := p.expect(tokenInclude, context)
	if err != nil {
		return "", err
	}

	t, err := p.expect(tokenString, context)
	if err != nil {
		return "", err
	}
	return t.val, nil
}

func (p *parser) parseModule() (*Module, error) {
	context := "module definition"

	_, err := p.expect(tokenModule, context)
	if err != nil {
		return nil, err
	}

	t, err := p.expect(tokenIdentifier, context)
	if err != nil {
		return nil, err
	}

	p.namespace = t.val
	mod := &Module{
		Name: t.val,
	}

	_, err = p.expect(tokenLBrace, context)
	if err != nil {
		return nil, err
	}

Loop:
	for {
		switch t := p.peek(); {
		case t.typ == tokenClass:
			class, err := p.parseClass()
			if err != nil {
				return nil, err
			}
			mod.Classes = append(mod.Classes, class)

		case t.typ == tokenRBrace:
			p.next() // read brace
			break Loop

		default:
			return nil, p.unexpected(t, context)
		}
	}
	return mod, nil
}

func (p *parser) parseClass() (*Class, error) {
	context := "class definition"

	_, err := p.expect(tokenClass, context)
	if err != nil {
		return nil, err
	}

	t, err := p.expect(tokenIdentifier, context)
	if err != nil {
		return nil, err
	}
	class := &Class{
		Name: t.val,
	}

	if _, err := p.expect(tokenLBrace, context); err != nil {
		return nil, err
	}

	for {
		t := p.peek()
		if t.typ == tokenRBrace {
			p.next() // read brace
			break
		}

		field, err := p.parseField("")
		if err != nil {
			return nil, err
		}
		class.Fields = append(class.Fields, field)
	}

	return class, nil
}

func (p *parser) parseField(context string) (*Field, error) {
	if context == "" {
		context = "field definition"
	}

	fieldType, err := p.parseType(context)
	if err != nil {
		return nil, err
	}

	t, err := p.expect(tokenIdentifier, context)
	if err != nil {
		return nil, err
	}

	_, err = p.expect(tokenSemicolon, context)
	if err != nil {
		return nil, err
	}

	return &Field{
		Type: fieldType,
		Name: t.val,
	}, nil
}

var typMap = map[tokenType]PTypeID{
	tokenBoolean: BooleanTypeID,
	tokenByte:    ByteTypeID,
	tokenInt:     IntTypeID,
	tokenLong:    LongTypeID,
	tokenFloat:   FloatTypeID,
	tokenDouble:  DoubleTypeID,
	tokenUString: UStringTypeID,
	tokenBuffer:  BufferTypeID,
}

func (p *parser) parseType(context string) (Type, error) {
	t := p.peek()
	switch {
	case typMap[t.typ] > 0:
		p.next() // consume the type
		return &PType{typMap[t.typ]}, nil
	case t.typ == tokenVector:
		return p.parseVector()
	case t.typ == tokenMap:
		return p.parseMap()
	case t.typ == tokenIdentifier:
		return p.parseClassRefType()
	}
	return nil, p.unexpected(t, context)
}

func (p *parser) parseVector() (*VectorType, error) {
	context := "vector type"
	if _, err := p.expect(tokenVector, context); err != nil {
		return nil, err
	}

	if _, err := p.expect(tokenLAngle, context); err != nil {
		return nil, err
	}

	typ, err := p.parseType(context)
	if err != nil {
		return nil, err
	}

	if _, err := p.expect(tokenRAngle, context); err != nil {
		return nil, err
	}

	return &VectorType{
		Type: typ,
	}, nil
}

func (p *parser) parseMap() (*MapType, error) {
	context := "map type"

	if _, err := p.expect(tokenMap, context); err != nil {
		return nil, err
	}

	if _, err := p.expect(tokenLAngle, context); err != nil {
		return nil, err
	}

	keyType, err := p.parseType(context)
	if err != nil {
		return nil, err
	}

	if _, err := p.expect(tokenComma, context); err != nil {
		return nil, err
	}

	valType, err := p.parseType(context)
	if err != nil {
		return nil, err
	}

	if _, err := p.expect(tokenRAngle, context); err != nil {
		return nil, err
	}

	return &MapType{
		KeyType: keyType,
		ValType: valType,
	}, nil

}

func (p *parser) parseClassRefType() (*ClassType, error) {
	context := "type"

	t, err := p.expect(tokenIdentifier, context)
	if err != nil {
		return nil, err
	}

	var namespace, ident string
	i := strings.LastIndex(t.val, ".")
	if i > 0 {
		namespace = t.val[:i]
		ident = t.val[i+1:]
	} else {
		ident = t.val
	}

	if p.namespace == namespace {
		namespace = ""
	}

	return &ClassType{
		Namespace: namespace,
		ClassName: ident,
	}, nil
}
