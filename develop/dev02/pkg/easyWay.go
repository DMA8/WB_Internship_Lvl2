package pkg

import (
	"bytes"
	"errors"
	"unicode"
)

//EasyUnpack based on bytes.Buffer.
func EasyUnpack(inpString string) (string, error) {
	var buffer bytes.Buffer
	inpRunes := []rune(inpString)
	if len(inpString) == 0 {
		return "", nil
	} else if unicode.IsDigit(inpRunes[0]) {
		return "", errors.New("string is incorrect")
	}
	for ind := 0; ind < len(inpRunes); ind++ {
		if inpRunes[ind] == '\\' {
			if len(inpRunes[ind: ]) > 2 && unicode.IsDigit(inpRunes[ind + 2]) {
				runeToAdd := inpRunes[ind + 1]
				nTimes := atoiStable(inpRunes[ind + 2:], &ind)
				for i := 0; i < nTimes; i++ {
					buffer.WriteString(string(runeToAdd))
				}
			} else {
				buffer.WriteString(string(inpRunes[ind + 1]))
				if inpRunes[ind + 1] == '\\' {
					ind++
				}
			}
		} else if !unicode.IsDigit(inpRunes[ind]) && // not digit AND
		(ind == len(inpRunes) - 1 || !unicode.IsDigit(inpRunes[ind + 1])) { // last elem OR next rune is not digit
			buffer.WriteString(string(inpRunes[ind]))
		} else if !unicode.IsDigit(inpRunes[ind]) && unicode.IsDigit(inpRunes[ind + 1]){
			runeToAdd := inpRunes[ind]
			nTimes := atoiStable(inpRunes[ind + 1:], &ind)
			for i := 0; i < nTimes; i++ {
				buffer.WriteString(string(runeToAdd))
			}
		}
	}
	return buffer.String(), nil
}
