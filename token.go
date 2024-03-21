package main

// tokens
const (
	LeftParen = iota
	RightParen
	LeftBrace
	RightBrace
	Comma
	Dot
	Minus
	Plus
	Semicolon
	Slash
	Star
	Bang
	BangEqual
	Equal
	EqualEqual
	Greater
	GreaterEqual
	Less
	LessEqual
	Identifier
	String
	Number
	And
	Class
	Else
	False
	Fun
	For
	If
	Nil
	Or
	Print
	Return
	Super
	This
	True
	Var
	While
	Eof
)

var keywords = map[string]Token{
	"and":    {Type: And, Lexeme: "and"},
	"class":  {Type: Class, Lexeme: "class"},
	"else":   {Type: Else, Lexeme: "else"},
	"false":  {Type: False, Lexeme: "false"},
	"for":    {Type: For, Lexeme: "for"},
	"fun":    {Type: Fun, Lexeme: "fun"},
	"if":     {Type: If, Lexeme: "if"},
	"nil":    {Type: Nil, Lexeme: "nil"},
	"or":     {Type: Or, Lexeme: "or"},
	"print":  {Type: Print, Lexeme: "print"},
	"return": {Type: Return, Lexeme: "return"},
	"super":  {Type: Super, Lexeme: "super"},
	"this":   {Type: This, Lexeme: "this"},
	"true":   {Type: True, Lexeme: "true"},
	"var":    {Type: Var, Lexeme: "var"},
	"while":  {Type: While, Lexeme: "while"},
}

type Token struct {
	Type   int
	Lexeme string
}

func NewToken(tokenType int, lexeme string) Token {
	return Token{tokenType, lexeme}
}
