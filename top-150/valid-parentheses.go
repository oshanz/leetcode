

type ParenthesesValidator []rune

func (p *ParenthesesValidator) Push(r rune) bool {
	if r == '(' || r == '{' || r == '[' {
		*p = append(*p, r)
		return true
	}
	if p.Valid(r) {
		*p = (*p)[:len(*p)-1]
		return true
	} else {
		return false
	}
}

func (p *ParenthesesValidator) Valid(r rune) bool {
	if len(*p) == 0 {
		return false
	}
	lastChar := (*p)[len(*p)-1]
	if lastChar == '[' && r == ']' {
		return true
	}
	if lastChar == '{' && r == '}' {
		return true
	}
	if lastChar == '(' && r == ')' {
		return true
	}
	return false
}

func isValid(s string) bool {
	var brackets ParenthesesValidator
	for _, char := range s {
		valid := brackets.Push(char)
		if !valid {
			return false
		}
	}
	if len(brackets) == 0 {
		return true
	}
	return false
}


