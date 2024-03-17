package main

import (
	"errors"
	"log"
	"regexp"
)

func isAlpha(c string) bool {
	return regexp.MustCompile(`^[_a-zA-Z]*$`).MatchString(c)
}

func isAlphanumeric(c string) bool {
	return regexp.MustCompile(`^[_a-zA-Z0-9]*$`).MatchString(c)
}

func isDigit(c string) bool {
	return regexp.MustCompile(`^[0-9]*$`).MatchString(c)
}

func next(source []byte, i int, c string) bool {
	if i >= len(source) {
		return false
	}
	return c == string(source[i])
}

func scanNumber(source []byte, i int) (Token, error) {
	x := i
	if isDigit(string(source[i])) {
		for i < len(source) && isDigit(string(source[i])) {
			i++
		}
		if i < len(source) {
			if next(source, i, ".") {
				i++
				for i < len(source) && isDigit(string(source[i])) {
					i++
				}
				if i >= len(source) {
					t := NewToken(Number, string(source[x:i]))
					return t, nil
				}
				if !next(source, i, " ") && !next(source, i, "\n") {
					return NewToken(Number, ""), errors.New("scanNumber: incorrect number format")
				}
			}
		}
		t := NewToken(Number, string(source[x:i]))
		return t, nil
	}
	return NewToken(Number, ""), errors.New("scanNumber: character is not a digit")
}

func scanIdentifier(source []byte, i int) (Token, error) {
	x := i
	for i < len(source) && isAlphanumeric(string(source[i])) {
		i++
	}
	return NewToken(Identifier, string(source[x:i])), nil
}

func Scan(source []byte) ([]Token, error) {
	tokens := []Token{}
	i := 0
	for {
		if i >= len(source) {
			break
		}

		c := string(source[i])
		switch string(c) {
		case " ":
			i++
			continue
		case "(":
			tokens = append(tokens, NewToken(LeftParen, "("))
		case ")":
			tokens = append(tokens, NewToken(RightParen, ")"))
		case "{":
			tokens = append(tokens, NewToken(LeftBrace, "{"))
		case "}":
			tokens = append(tokens, NewToken(RightBrace, "}"))
		case ",":
			tokens = append(tokens, NewToken(Comma, ","))
		case ".":
			tokens = append(tokens, NewToken(Dot, "."))
		case "-":
			tokens = append(tokens, NewToken(Minus, "-"))
		case "+":
			tokens = append(tokens, NewToken(Plus, "+"))
		case ";":
			tokens = append(tokens, NewToken(Semicolon, ";"))
		case "*":
			tokens = append(tokens, NewToken(Star, "*"))
		case "!":
			if next(source, i+1, "=") {
				tokens = append(tokens, NewToken(BangEqual, "!="))
				i++
			} else {
				tokens = append(tokens, NewToken(Bang, "!"))
			}
		case "=":
			if next(source, i+1, "=") {
				tokens = append(tokens, NewToken(EqualEqual, "=="))
				i++
			} else {
				tokens = append(tokens, NewToken(Equal, "="))
			}
		case "<":
			if next(source, i+1, "=") {
				tokens = append(tokens, NewToken(LessEqual, "<="))
				i++
			} else {
				tokens = append(tokens, NewToken(Less, "<"))
			}
		case ">":
			if next(source, i+1, "=") {
				tokens = append(tokens, NewToken(GreaterEqual, ">="))
				i++
			} else {
				tokens = append(tokens, NewToken(Greater, ">"))
			}
		case "/":
			if next(source, i+1, "/") {
				for i < len(source) && !next(source, i+1, "\n") {
					i++
				}
			} else {
				tokens = append(tokens, NewToken(Slash, "/"))
			}
		case "\"":
			j := i + 1
			for i < len(source) && !next(source, i+1, "\"") {
				i++
			}
			t := NewToken(String, "")
			t.Value = string(source[j : i+1]) // fix this off by one
			tokens = append(tokens, t)
			i++
		default:
			if isDigit(c) {
				t, err := scanNumber(source, i)
				if err != nil {
					log.Fatal(err)
				}
				tokens = append(tokens, t)
				// fix this, it's awful
				i = i + len(t.Lexeme)
				continue
			}
			if isAlpha(c) {
				t, err := scanIdentifier(source, i)
				if err != nil {
					log.Fatal(err)
				}
				// fix this, it's awful
				tokens = append(tokens, t)
				i = i + len(t.Lexeme)
				continue
			}
			i++
			continue
		}
		i++
	}
	tokens = append(tokens, NewToken(Eof, "EOF"))
	return tokens, nil
}
