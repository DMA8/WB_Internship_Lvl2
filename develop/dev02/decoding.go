package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"unicode"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	var inputString string = `a4bc2d5e`
	//fmt.Scan(&inputString)
	result, err := Unpack(inputString)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Result is", result)
	result, err = EasyUnpack(inputString)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Result is", result)
}

//Unpacks string from a4b3de2 to aaaabbbdee
//Also qwe\45 to qwe44444
func Unpack(inpStr string) (string, error){
	// распарсить и подготовить нужную длину под выходную строку
	inpRunes := []rune(inpStr)
	length := GetCloseLength(inpRunes)
	outRunes := make([]rune, int(float64(length) * 1.1 ))
	if len(inpRunes) == 0{
		return "", nil
	} else if !unicode.IsLetter(inpRunes[0]) || inpRunes[0] == 92 { // if first letter is not digit -> return
		return "", errors.New("string is not correct")
	}
	unpackEngine(inpRunes, outRunes)
	return string(outRunes), nil
}

func EscapeMultiplyRune(InpRunes, OutRunes []rune, ind, outInd *int ) {
	letter := InpRunes[*ind + 1]
	nLetter := AtoiStable(InpRunes[*ind+2: ], ind)
	for i := 0; i < nLetter; i++ {
		OutRunes[*outInd] = letter
		*outInd += 1
	}
	*ind += 1
}
func MultiplyRune(InpRunes, OutRunes []rune, ind, outInd *int ) {
	letter := InpRunes[*ind]
	nLetter := AtoiStable(InpRunes[*ind + 1: ], ind)
	for i := 0; i < nLetter; i++ {
		OutRunes[*outInd] = letter
		*outInd += 1
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
			MultiplyRune(InpRunes, OutRunes, &ind, &outInd)
		} else if len(InpRunes[ind:]) > 2 && InpRunes[ind] == '\\' && unicode.IsDigit(InpRunes[ind + 2]) {
			EscapeMultiplyRune(InpRunes, OutRunes, &ind, &outInd)
			ind++
		} else if len(InpRunes[ind:]) > 1 && InpRunes[ind] == '\\'{
				letter := InpRunes[ind + 1]
				OutRunes[outInd] = letter
				outInd++
				ind++
		}
	}
}

// returns length of unpacked string
func GetCloseLength(runes []rune) int {
	length := 0
	for ind := 0; ind < len(runes); ind++ {
		if unicode.IsLetter(runes[ind]) && ind == len(runes) - 1 || 
		unicode.IsLetter(runes[ind]) && (unicode.IsLetter(runes[ind + 1]) || runes[ind + 1] == '\\'){
			length++
		} else if unicode.IsLetter(runes[ind]) && unicode.IsDigit(runes[ind + 1]) {
			val := AtoiStable(runes[ind + 1:], &ind)
			length += val
		} else if runes[ind] == '\\' {
			val := EscapeAtoi(runes[ind:], &ind)
			length += val
		}
	}
	return int(float64(length))
}

// It return int to strings that ends not only with digit
// "123я" -> 123
func AtoiStable(inpRunes []rune, ind *int) int{
	nDigits := 0
	for nDigits < len(inpRunes) && unicode.IsDigit(inpRunes[nDigits]) {
		nDigits += 1
		*ind += 1
	}
	outDigit, _ := strconv.Atoi(string(inpRunes[:nDigits]))
	return outDigit
}

// "\45" -> 5; \2100 -> 100; \200???
func EscapeAtoi(inpRunes []rune, ind *int) int{
	val := 0
	if inpRunes[0] == '\\' && len(inpRunes) > 2 {
		val = AtoiStable(inpRunes[2:], ind)
		if val == 0 {
			val = 1
		}
	} else {
		val = 1
	}
	return val
}
