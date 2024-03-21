package main

import (
	"errors"
	"log"
)

func isAlpha(c byte) bool {
	return c == '_' || c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z'

}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func isAlphanumeric(c byte) bool {
	return isAlpha(c) || isDigit(c)
}

type Scanner struct {
	source []byte
	index  int
}

func NewScanner(source []byte) Scanner {
	return Scanner{source, 0}
}

func (s *Scanner) Consume() {
	s.index++
}

func (s *Scanner) Peek() byte {
	if s.index < len(s.source) {
		return s.source[s.index]
	}
	return 0 //eof
}

func (s *Scanner) Match(c string) bool {
	if s.index < len(s.source) {
		return string(s.source[s.index]) == c
	}
	return false
}

func (s *Scanner) MatchNext(c string) bool {
	if s.index+1 < len(s.source) {
		return string(s.source[s.index+1]) == c
	}
	return false
}

func (s *Scanner) Scan() ([]Token, error) {
	tokens := []Token{}
	for {
		if s.Peek() == 0 {
			break
		}
		c := s.Peek()
		switch c {
		case ' ':
			s.Consume()
			continue
		case '(':
			tokens = append(tokens, NewToken(LeftParen, "("))
		case ')':
			tokens = append(tokens, NewToken(RightParen, ")"))
		case '{':
			tokens = append(tokens, NewToken(LeftBrace, "{"))
		case '}':
			tokens = append(tokens, NewToken(RightBrace, "}"))
		case ',':
			tokens = append(tokens, NewToken(Comma, ","))
		case '.':
			tokens = append(tokens, NewToken(Dot, "."))
		case '-':
			tokens = append(tokens, NewToken(Minus, "-"))
		case '+':
			tokens = append(tokens, NewToken(Plus, "+"))
		case ';':
			tokens = append(tokens, NewToken(Semicolon, ";"))
		case '*':
			tokens = append(tokens, NewToken(Star, "*"))
		case '!':
			if s.MatchNext("=") {
				tokens = append(tokens, NewToken(BangEqual, "!="))
				s.Consume()
			} else {
				tokens = append(tokens, NewToken(Bang, "!"))
			}
		case '=':
			if s.MatchNext("=") {
				tokens = append(tokens, NewToken(EqualEqual, "=="))
				s.Consume()
			} else {
				tokens = append(tokens, NewToken(Equal, "="))
			}
		case '<':
			if s.MatchNext("=") {
				tokens = append(tokens, NewToken(LessEqual, "<="))
				s.Consume()
			} else {
				tokens = append(tokens, NewToken(Less, "<"))
			}
		case '>':
			if s.MatchNext("=") {
				tokens = append(tokens, NewToken(GreaterEqual, ">="))
				s.Consume()
			} else {
				tokens = append(tokens, NewToken(Greater, ">"))
			}
		case '/':
			if s.MatchNext("/") {
				for !s.MatchNext("\n") {
					s.Consume()
				}
			} else {
				tokens = append(tokens, NewToken(Slash, "/"))
			}
		case '"':
			j := s.index + 1
			s.Consume()
			for !s.Match("\"") {
				s.Consume()
			}
			if s.Peek() == 0 {
				log.Fatal("scan: case \": string does not end")
			}
			t := NewToken(String, string(s.source[j:s.index]))
			tokens = append(tokens, t)
			s.Consume()
		default:
			if isDigit(c) {
				t, err := s.scanNumber()
				if err != nil {
					log.Fatal(err)
				}
				tokens = append(tokens, t)
				continue
			}
			if isAlpha(c) {
				t, err := s.scanIdentifier()
				if err != nil {
					log.Fatal(err)
				}
				tokens = append(tokens, t)
				continue
			}
			s.Consume()
			continue
		}
		s.Consume()
	}
	tokens = append(tokens, NewToken(Eof, "EOF"))
	return tokens, nil
}

func (s *Scanner) scanNumber() (Token, error) {
	if !isDigit(s.Peek()) {
		return NewToken(Number, ""), errors.New("scanNumber: character is not a digit")
	}
	x := s.index
	for isDigit(s.Peek()) {
		s.Consume()
	}
	if s.Match(".") {
		s.Consume()
		for isDigit(s.Peek()) {
			s.Consume()
		}
		if s.index >= len(s.source) {
			t := NewToken(Number, string(s.source[x:s.index]))
			return t, nil
		}
		if !s.Match(" ") && !s.Match("\n") {
			return NewToken(Number, ""), errors.New("scanNumber: incorrect number format")
		}
	}
	t := NewToken(Number, string(s.source[x:s.index]))
	return t, nil
}

func (s *Scanner) scanIdentifier() (Token, error) {
	var x, y int
	x = s.index
	for isAlphanumeric(s.Peek()) {
		s.Consume()
	}
	y = s.index
	t, ok := keywords[string(s.source[x:y])]
	if !ok {
		return NewToken(Identifier, string(s.source[x:y])), nil
	}
	return t, nil
}
