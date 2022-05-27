package calculator

func Calculate(expression string) (float64, error) {
	tokens, err := ParseTokens(&expression)

	if err != nil {
		return 0, err
	}

	result := process(&tokens)

	return result, nil
}
