package parser

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

type token struct {
	typ tokenType
	val string
}

func (t token) String() string {
	switch {
	case t.typ == tokenEOF:
		return "EOF"
	case t.typ == tokenError:
		return t.val
	case t.typ > tokenKeyword:
		return fmt.Sprintf("<%s>", t.val)
		//	case len(t.val) > 10:
		//return fmt.Sprintf("%.10q...", t.val)
	}
	return fmt.Sprintf("%q", t.val)
}

type tokenType int

const (
	tokenError tokenType = iota // error occured during lexing.
	tokenEOF
	tokenIdentifier
	tokenBlockComment
	tokenLineComment
	tokenSemicolon // ;
	tokenLBrace    // {
	tokenRBrace    // }
	tokenString    // quoted string
	tokenLAngle    // <
	tokenRAngle    // >
	tokenComma     // ,
	// keywords
	tokenKeyword
	tokenInclude
	tokenModule
	tokenClass
	tokenByte
	tokenBoolean
	tokenInt
	tokenLong
	tokenFloat
	tokenDouble
	tokenUString
	tokenBuffer
	tokenVector
	tokenMap
)

var punc = map[rune]tokenType{
	';': tokenSemicolon,
	'{': tokenLBrace,
	'}': tokenRBrace,
	'<': tokenLAngle,
	'>': tokenRAngle,
	',': tokenComma,
}

var key = map[string]tokenType{
	"include": tokenInclude,
	"module":  tokenModule,
	"class":   tokenClass,
	"byte":    tokenByte,
	"boolean": tokenBoolean,
	"int":     tokenInt,
	"long":    tokenLong,
	"float":   tokenFloat,
	"double":  tokenDouble,
	"ustring": tokenUString,
	"buffer":  tokenBuffer,
	"vector":  tokenVector,
	"map":     tokenMap,
}

const eof = -1

// stateFn represents the state of the scanner as a function that returns the next state.
type stateFn func(*lexer) stateFn

type lexer struct {
	input             string
	start, pos, width int
	tokens            chan token
}

// next reutnrs the next run in the input.
func (l *lexer) next() rune {
	if l.pos >= len(l.input) {
		l.width = 0
		return eof
	}

	r, w := utf8.DecodeRuneInString(l.input[l.pos:])
	l.width = w
	l.pos += l.width
	return r
}

func (l *lexer) backup() {
	l.pos -= l.width
}

func (l *lexer) errorf(f string, v ...interface{}) stateFn {
	l.tokens <- token{tokenError, fmt.Sprintf(f, v...)}
	return nil
}

func (l *lexer) emit(t tokenType) {
	l.tokens <- token{t, l.input[l.start:l.pos]}
	l.start = l.pos
}

func (l *lexer) ignore() {
	l.start = l.pos
}

func (l *lexer) nextToken() token {
	return <-l.tokens
}

func (l *lexer) skipSpace() {
	for isWhitespace(l.next()) {
	}
	l.backup()
	l.ignore()
}

func lex(input string) *lexer {
	l := &lexer{
		input:  input,
		tokens: make(chan token),
	}
	go l.run()
	return l
}

func (l *lexer) run() {
	for state := lexStart; state != nil; {
		state = state(l)
	}
	close(l.tokens)
}

func lexStart(l *lexer) stateFn {
	l.skipSpace()

	r := l.next()
	switch {
	case r == '/':
		switch r := l.next(); {
		case r == '/':
			return lexLineComment
		case r == '*':
			return lexBlockComment
		default:
			return l.errorf("bad character %#U", r)
		}
	case punc[r] > 0:
		l.emit(punc[r])
		return lexStart
	case r == '"':
		l.ignore()
		return lexQuoted
	case unicode.IsLetter(r):
		l.backup()
		return lexIdentifier
	case r == eof:
		break
	default:
		return l.errorf("unknown character %#U", r)
	}
	l.emit(tokenEOF)
	return nil
}

func lexLineComment(l *lexer) stateFn {
Loop:
	for {
		switch l.next() {
		case eof, '\n':
			l.backup()
			break Loop
		}
	}
	l.emit(tokenLineComment)
	return lexStart
}

func lexBlockComment(l *lexer) stateFn {
Loop:
	for {
		switch l.next() {
		case eof:
			return l.errorf("unterminated block comment")
		case '*':
			if l.next() == '/' {
				break Loop
			}
			l.backup()
		}
	}
	l.emit(tokenBlockComment)
	return lexStart
}

func lexQuoted(l *lexer) stateFn {
Loop:
	for {
		switch l.next() {
		case eof, '\n':
			return l.errorf("unterminated quoted string")
		case '"':
			l.backup()
			break Loop
		}
	}
	l.emit(tokenString)
	l.next() // eat the last ""
	return lexStart
}

func lexIdentifier(l *lexer) stateFn {
Loop:
	for {
		switch r := l.next(); {
		case isAlphaNumeric(r):
			// absorb
		default:
			l.backup()
			word := l.input[l.start:l.pos]
			switch {
			case key[word] > tokenKeyword:
				l.emit(key[word])
			default:
				l.emit(tokenIdentifier)
			}
			break Loop
		}
	}
	return lexStart
}

// isWhitespace reports if r is space or a new line
func isWhitespace(r rune) bool {
	return r == ' ' || r == '\t' || r == '\r' || r == '\n'
}

// isAlphaNumeric reports whether r is an alphabetic, digit, or underscore.
func isAlphaNumeric(r rune) bool {
	return r == '_' || unicode.IsLetter(r) || unicode.IsDigit(r) || r == '.'
}
