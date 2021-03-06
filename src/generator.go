package main

// Generate g
func Generate(root *Root) string {
	return "(function() {" + getBody(root) + "})()"
}

func getBody(root *Root) string {
	value := root.value
	str := ""

	for i := 0; i < len(value); i++ {
		str = str + parseTreeToCode(value[i])
	}

	return str
}

func parseTreeToCode(tree *Tree) string {
	left := ""
	right := ""

	if tree == nil {
		return ""
	}

	if tree.left != nil {
		left = parseTreeToCode(tree.left)
	}
	if tree.right != nil {
		right = parseTreeToCode(tree.right)
	}

	return left + getCode(tree.value) + right
}

func getCode(token interface{}) string {
	switch token.(type) {
	case Token:
		return parseToken(token.(Token))

	case TokenList:
		item := token.(TokenList)
		val := ""
		for i := 0; i < len(item); i++ {
			val += parseToken(item[i])
		}

		return val
	}

	return ""
}

func parseToken(token Token) string {
	switch token.tokenType {
	case Variable:
		return "var "
	case Identifier:
		return token.value
	case Assignment:
		return "="
	case Number:
		return token.value
	case End:
		return ";"
	case LeftParen:
		return "("
	case Function:
		return "function " + token.value + "() {"
	}

	return ""
}
