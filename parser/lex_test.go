package parser

import "testing"

type lexTest struct {
	name   string
	input  string
	tokens []token
}

func mkToken(typ tokenType, text string) token {
	return token{
		typ: typ,
		val: text,
	}
}

var (
	tEOF       = mkToken(tokenEOF, "")
	tLBrace    = mkToken(tokenLBrace, "{")
	tRBrace    = mkToken(tokenRBrace, "}")
	tLAngle    = mkToken(tokenLAngle, "<")
	tRAngle    = mkToken(tokenRAngle, ">")
	tComma     = mkToken(tokenComma, ",")
	tSemicolon = mkToken(tokenSemicolon, ";")
	tInclude   = mkToken(tokenInclude, "include")
	tModule    = mkToken(tokenModule, "module")
	tClass     = mkToken(tokenClass, "class")
	tByte      = mkToken(tokenByte, "byte")
	tBoolean   = mkToken(tokenBoolean, "boolean")
	tInt       = mkToken(tokenInt, "int")
	tLong      = mkToken(tokenLong, "long")
	tFloat     = mkToken(tokenFloat, "float")
	tDouble    = mkToken(tokenDouble, "double")
	tUString   = mkToken(tokenUString, "ustring")
	tBuffer    = mkToken(tokenBuffer, "buffer")
	tVector    = mkToken(tokenVector, "vector")
	tMap       = mkToken(tokenMap, "map")
)

var sampleDDL = `
module links {
    class Link {
        ustring URL;
        boolean isRelative;
        ustring anchorText;
    };
}
`

var lexTests = []lexTest{
	{"empty", "", []token{tEOF}},
	{"spaces", " \t\n", []token{tEOF}},
	{"punctuation", "{}<>,;", []token{
		tLBrace,
		tRBrace,
		tLAngle,
		tRAngle,
		tComma,
		tSemicolon,
		tEOF,
	}},
	{"line comment", "// this is a line comment\n\n", []token{
		mkToken(tokenLineComment, "// this is a line comment"),
		tEOF,
	}},
	{"block comment", "/** line 1\n * line 2\n *line 3\n */", []token{
		mkToken(tokenBlockComment, "/** line 1\n * line 2\n *line 3\n */"),
		tEOF,
	}},
	{"bad comment", "/whoops", []token{
		mkToken(tokenError, "bad character U+0077 'w'"),
	}},
	{"unterminated block comment", "/* i am error", []token{
		mkToken(tokenError, "unterminated block comment"),
	}},
	{"quoted string", `"suck it trebek"`, []token{
		mkToken(tokenString, "suck it trebek"),
		tEOF,
	}},
	{"unquoted string", `"abc`, []token{
		mkToken(tokenError, "unterminated quoted string"),
	}},
	{"keywords", "include module class", []token{
		tInclude,
		tModule,
		tClass,
		tEOF,
	}},
	{"ptypes", "byte boolean int long float double ustring buffer", []token{
		tByte,
		tBoolean,
		tInt,
		tLong,
		tFloat,
		tDouble,
		tUString,
		tBuffer,
		tEOF,
	}},
	{"ctypes", "vector<ustring> map<int, float>", []token{
		tVector,
		tLAngle,
		tUString,
		tRAngle,
		tMap,
		tLAngle,
		tInt,
		tComma,
		tFloat,
		tRAngle,
		tEOF,
	}},
	{"identifier", "myThing myThing2 com.github.fake.news", []token{
		mkToken(tokenIdentifier, "myThing"),
		mkToken(tokenIdentifier, "myThing2"),
		mkToken(tokenIdentifier, "com.github.fake.news"),
		tEOF,
	}},
	{"bad character", "!", []token{
		mkToken(tokenError, "unknown character U+0021 '!'"),
	}},
	{"complete ddl", sampleDDL, []token{
		tModule,
		mkToken(tokenIdentifier, "links"),
		tLBrace,
		tClass,
		mkToken(tokenIdentifier, "Link"),
		tLBrace,
		tUString,
		mkToken(tokenIdentifier, "URL"),
		tSemicolon,
		tBoolean,
		mkToken(tokenIdentifier, "isRelative"),
		tSemicolon,
		tUString,
		mkToken(tokenIdentifier, "anchorText"),
		tSemicolon,
		tRBrace,
		tSemicolon,
		tRBrace,
		tEOF,
	}},
}

func collect(input string) []token {
	tokens := []token{}
	l := lex(input)
	for {
		token := l.nextToken()
		tokens = append(tokens, token)
		if token.typ == tokenEOF || token.typ == tokenError {
			break
		}
	}
	return tokens
}

func TestLex(t *testing.T) {
	for _, tc := range lexTests {
		t.Run(tc.name, func(t *testing.T) {
			tokens := collect(tc.input)

			if !equal(tokens, tc.tokens) {
				t.Errorf("got\n\t%+v\nexpected:\n\t%+v", tokens, tc.tokens)
			}
		})
	}

}

func equal(t1, t2 []token) bool {
	if len(t1) != len(t2) {
		return false
	}

	for i := range t1 {
		if t1[i].typ != t2[i].typ {
			return false
		}
		if t1[i].val != t2[i].val {
			return false
		}
	}

	return true
}
