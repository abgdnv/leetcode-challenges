package main

func parseBoolExpr(expression string) bool {
	return parse(expression)
}

func parse(s string) bool {
	if s == "t" {
		return true
	}
	if s == "f" {
		return false
	}

	if s[0] == '!' {
		return !parse(s[2 : len(s)-1])
	}
	if s[0] == '&' {
		items := split(s[2 : len(s)-1])
		result := true
		for _, v := range items {
			result = result && parse(v)
			if !result {
				return false
			}
		}
		return result
	}
	if s[0] == '|' {
		items := split(s[2 : len(s)-1])
		result := false
		for _, v := range items {
			result = result || parse(v)
			if result {
				return true
			}
		}
		return result
	}

	return false
}

func split(s string) []string {
	begin, end := 0, 0
	ob, cb := 0, 0
	strs := []string{}
	for end < len(s) {
		if s[end] == ',' && ob == cb {
			strs = append(strs, s[begin:end])
			begin, end = end+1, end+2
			continue
		}
		switch s[end] {
		case '(':
			ob++
		case ')':
			cb++
		}
		end++
	}
	if begin < len(s) {
		strs = append(strs, s[begin:])
	}
	return strs
}
