package pkg

import (
	"errors"
	"strconv"
	"unicode"
)



//Unpack  string from a4b3de2 to aaaabbbdee //Also qwe\45 to qwe44444
func Unpack(inpStr string) (string, error){
	// распарсить и подготовить нужную длину под выходную строку
	inpRunes := []rune(inpStr)
	length := getCloseLength(inpRunes)
	outRunes := make([]rune, int(float64(length) * 1.1 ))
	if len(inpRunes) == 0{
		return "", nil
	} else if !unicode.IsLetter(inpRunes[0]) || inpRunes[0] == 92 { // if first letter is not digit -> return
		return "", errors.New("string is not correct")
	}
	unpackEngine(inpRunes, outRunes)
	return string(outRunes), nil
}

// escapeMultiplyRune (for escape case '\') repeats given rune and writes it in given slice. It increments index pointer!
func escapeMultiplyRune(InpRunes, OutRunes []rune, ind, outInd *int ) {
	letter := InpRunes[*ind + 1]
	nLetter := atoiStable(InpRunes[*ind+2: ], ind)
	for i := 0; i < nLetter; i++ {
		OutRunes[*outInd] = letter
		*outInd++
	}
	*ind++
}

// multiplyRune repeats given rune and writes it in given slice. It increments index pointer!
func multiplyRune(InpRunes, OutRunes []rune, ind, outInd *int ) {
	letter := InpRunes[*ind]
	nLetter := atoiStable(InpRunes[*ind + 1: ], ind)
	for i := 0; i < nLetter; i++ {
		OutRunes[*outInd] = letter
		*outInd++
	}
}

// does main job
func unpackEngine(InpRunes, OutRunes []rune) {
	outInd := 0
	for ind := 0; ind < len(InpRunes); ind ++ {
		if ind == len(InpRunes) - 1 || unicode.IsLetter(InpRunes[ind]) && 
		(unicode.IsLetter(InpRunes[ind + 1]) || InpRunes[ind + 1] == '\\') {
				OutRunes[outInd] = InpRunes[ind]
				outInd++
		} else if unicode.IsLetter(InpRunes[ind]) && unicode.IsDigit(InpRunes[ind + 1]) {
			multiplyRune(InpRunes, OutRunes, &ind, &outInd)
		} else if len(InpRunes[ind:]) > 2 && InpRunes[ind] == '\\' && unicode.IsDigit(InpRunes[ind + 2]) {
			escapeMultiplyRune(InpRunes, OutRunes, &ind, &outInd)
			ind++
		} else if len(InpRunes[ind:]) > 1 && InpRunes[ind] == '\\'{
				letter := InpRunes[ind + 1]
				OutRunes[outInd] = letter
				outInd++
				ind++
		}
	}
}

//getCloseLength returns length of unpacked string. It is for getting needed memory at the beginning of the programm
func getCloseLength(runes []rune) int {
	length := 0
	for ind := 0; ind < len(runes); ind++ {
		if unicode.IsLetter(runes[ind]) && ind == len(runes) - 1 || 
		unicode.IsLetter(runes[ind]) && (unicode.IsLetter(runes[ind + 1]) || runes[ind + 1] == '\\'){
			length++
		} else if unicode.IsLetter(runes[ind]) && unicode.IsDigit(runes[ind + 1]) {
			val := atoiStable(runes[ind + 1:], &ind)
			length += val
		} else if runes[ind] == '\\' {
			val := escapeAtoi(runes[ind:], &ind)
			length += val
		}
	}
	return int(float64(length))
}

// atoiStable returns int to strings that ends not only with digit
// "123я" -> 123
func atoiStable(inpRunes []rune, ind *int) int{
	nDigits := 0
	for nDigits < len(inpRunes) && unicode.IsDigit(inpRunes[nDigits]) {
		nDigits++
		*ind++
	}
	outDigit, _ := strconv.Atoi(string(inpRunes[:nDigits]))
	return outDigit
}

//escapeAtoi returns number that starts from the second digit "\45" -> 5; \2100 -> 100; \200???
func escapeAtoi(inpRunes []rune, ind *int) int{
	val := 0
	if inpRunes[0] == '\\' && len(inpRunes) > 2 {
		val = atoiStable(inpRunes[2:], ind)
		if val == 0 {
			val = 1
		}
	} else {
		val = 1
	}
	return val
}
