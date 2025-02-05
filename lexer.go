package main

type Token int

const (
	IllegalToken Token = iota

	LEFT_CURLY_BRACE  // {
	RIGHT_CURLY_BRACE // }
)

func lex(s []byte) []Token {
	tokens := make([]Token, len(s))
	for i := 0; i < len(s); i++ {
		tokens[i] = getToken(s[i])
	}
	return tokens
}

func getToken(chr byte) Token {
	switch chr {
	case '{':
		return LEFT_CURLY_BRACE
	case '}':
		return RIGHT_CURLY_BRACE
	default:
		return IllegalToken
	}
}
