package main

type State int

func parse(s []byte) bool {
	tokens := lex(s)
	if len(tokens) < 2 {
		return false
	}

	stack := make([]Token, 0)
	for i := 0; i < len(tokens); i++ {
		token := tokens[i]
		switch token {
		case LEFT_CURLY_BRACE:
			stack = append(stack, LEFT_CURLY_BRACE)
		case RIGHT_CURLY_BRACE:
			if stack[len(stack)-1] == LEFT_CURLY_BRACE {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		default:
			return false
		}
	}

	return true
}
