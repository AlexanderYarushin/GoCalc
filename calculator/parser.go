package calculator

import (
	"GoCalc/utils"
	"errors"
	"unicode"
)

/*
	return:
		numberStack signStack error
*/

func ParseTokens(expression *string) ([]string, error) {

	if err := validate(expression); err != nil {
		return nil, err
	}

	var str = *expression

	utils.ClearSpace(&str)

	var result []string
	var tmp string

	for i := 0; i < len(str); i++ {
		byteStr := string(str[i])

		if Signs[byteStr].Index != 0 {
			if len(tmp) != 0 {
				result = append(result, tmp)
			}
			result = append(result, byteStr)
			tmp = ""
			continue
		}

		if utils.IsNumber(&byteStr) || byteStr == "." {
			tmp += byteStr

			if i == len(str)-1 {
				result = append(result, tmp)
			}
		}

	}

	return result, nil
}

func validate(expression *string) error {
	openBrackets := 0
	closeBrackets := 0

	for _, value := range *expression {
		if string(value) == "(" {
			openBrackets++
		}

		if string(value) == ")" {
			closeBrackets++
		}

		if unicode.IsLetter(value) {
			return errors.New("unexpected symbol '" + string(value) + "'")
		}

	}

	if openBrackets != closeBrackets {
		return errors.New("the number of open and closed brackets does not match")
	}

	return nil
}
