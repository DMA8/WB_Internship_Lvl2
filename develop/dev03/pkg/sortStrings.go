package pkg

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
	"io/ioutil"
	"log"
)

// SortStrings is made for sorting lines in given files. 
type SortStrings struct{
	Flags			*CmdLine
	InpLines		[]string
	SortedStrings	[]string
	Sorter			Sorter
}

// NewSortStrings - constuctor for SortStrings
func NewSortStrings(flags *CmdLine, sorter Sorter, inpLines []string) *SortStrings{
	return &SortStrings{
		flags,
		inpLines,
		make([]string, 0),
		sorter,
	}
}

// Sort sorts lines 
func (s *SortStrings)Sort(){
	if s.Flags.MonthSort{
		s.keyColumnSortAsMonth()
	} else if s.Flags.SuffixSort{
		s.keyColumnSortAsSuffix()
	} else if s.Flags.NumberSort{
		s.keyColumnSortAsFloat()
	} else {
		s.keyColumnSortAsString()
	}

	if s.Flags.UniqueLines{
		s.onlyUniqueLines()
	}

	if s.Flags.ReversedOrder{
		s.reversedOrder()
	}
	
	if s.Flags.CheckSort{
		s.checkOrder()
	}
}

//simpleSort sorts given lines as strings. PivotColumn = 1
func(s *SortStrings)simpleSort(){
	copy(s.SortedStrings, s.InpLines)
	s.Sorter.Sort(s.SortedStrings)
}

//keyColumnSortAsString sorts given lines as strings in given column
func (s *SortStrings)keyColumnSortAsString(){
	var columnValue string
	Columns := make([]string, 0, len(s.InpLines))
	HashColumnString := make(map[string][]string)
	for _, v := range s.InpLines{
		splittedLine := strings.Split(v, " ")
		if s.Flags.PivotColumn <= len(splittedLine){
			columnValue = splittedLine[s.Flags.PivotColumn - 1]
		} else {
			columnValue = " "
		}
		Columns = append(Columns, columnValue)
		HashColumnString[columnValue] = append(HashColumnString[columnValue], v)
	}
	s.Sorter.Sort(Columns)
	for _, v := range Columns{
		s.Sorter.Sort(HashColumnString[v])
		s.SortedStrings = append(s.SortedStrings, HashColumnString[v]...)
	}
}

//keyColumnSortAsFloat sorts given lines as numbers in given column
func (s *SortStrings)keyColumnSortAsFloat(){
	var columnValue float64
	Columns := make([]float64, 0, len(s.InpLines))
	HashColumnString := make(map[float64][]string)
	for _, v := range s.InpLines{
		splittedLine := strings.Split(v, " ")
		if s.Flags.PivotColumn <= len(splittedLine){
			columnValue = safeFloatConvert(splittedLine[s.Flags.PivotColumn - 1])
		} else {
			columnValue = 0
		}
		Columns = append(Columns, columnValue)
		HashColumnString[columnValue] = append(HashColumnString[columnValue], v)
	}
	s.Sorter.Sort(Columns)
	vLast := 0.0
	for i, v := range Columns{
		if i != 0{
			if vLast == v {
				continue
			}
		}
		s.Sorter.Sort(HashColumnString[v])
		s.SortedStrings = append(s.SortedStrings, HashColumnString[v]...)
		vLast = v
	}
}

//keyColumnSortAsMonth sorts lines in order the Month given in particular column
func (s *SortStrings)keyColumnSortAsMonth(){
	var month = map[string]int{
		"JAN": 1,
		"FEB": 2,
		"MAR": 3,
		"APR": 4,
		"MAY": 5,
		"JUN": 6,
		"JUL": 7,
		"AUG": 8,
		"SEP": 9,
		"OCT": 10,
		"NOV": 11,
		"DEC": 12,
	}
	var columnValue int
	Columns := make([]int, 0, len(s.InpLines))
	HashColumnString := make(map[int][]string)
	for _, v := range s.InpLines{
		splittedLine := strings.Split(v, " ")
		if s.Flags.PivotColumn <= len(splittedLine) && len(splittedLine[s.Flags.PivotColumn - 1]) >= 3{
			prefix := strings.ToUpper(splittedLine[s.Flags.PivotColumn - 1][:3])
			columnValue = month[prefix]
		} else {
			columnValue = 0
		}
		Columns = append(Columns, columnValue)
		HashColumnString[columnValue] = append(HashColumnString[columnValue], v)
	}
	s.Sorter.Sort(Columns)
	vLast := 0
	for i, v := range Columns{
		if i != 0{
			if vLast == v {
				continue
			}
		}
		s.Sorter.Sort(HashColumnString[v])
		s.SortedStrings = append(s.SortedStrings, HashColumnString[v]...)
		vLast = v
	}
}

//keyColumnSortAsSuffix sorts columns as suffix (Megabytes, Gigabytes...)
func (s *SortStrings)keyColumnSortAsSuffix(){
	var suffixes = map[string]int{
		"K":  1 << 10,
		"M":  1 << 20,
		"G":  1 << 30,
		"T":  1 << 40,
	}
	var columnValue float64
	Columns := make([]float64, 0, len(s.InpLines))
	HashColumnString := make(map[float64][]string)
	for _, v := range s.InpLines{
		splittedLine := strings.Split(v, " ")
		if s.Flags.PivotColumn <= len(splittedLine){
			digit := safeFloatConvert(splittedLine[s.Flags.PivotColumn - 1])
			for _, c := range splittedLine[s.Flags.PivotColumn - 1]{
				if c >= '0' && c <= '9' || c == '.'{
					continue
				} else{
					power := float64(suffixes[strings.ToUpper(string(c))])
					columnValue = digit * power

					break
				}
			}
		} else {
			columnValue = 0.0
		}
		Columns = append(Columns, columnValue)
		HashColumnString[columnValue] = append(HashColumnString[columnValue], v)
	}
	s.Sorter.Sort(Columns)
	vLast := 0.0
	for i, v := range Columns{
		if i != 0{
			if vLast == v {
				continue
			}
		}
		s.Sorter.Sort(HashColumnString[v])
		s.SortedStrings = append(s.SortedStrings, HashColumnString[v]...)
		vLast = v
	}
}

//safeFloatConvert Parses float as standart func but not return 0 when not digit or separator is encounted
func safeFloatConvert(inpString string) (float64){
	outValue := 0.0
	pointSeen := 0
	indx := 0
	for i := range inpString {
		if inpString[i] == ',' || inpString[i] == '.' && pointSeen < 1{
			pointSeen ++
			indx++
			continue
		}
		if pointSeen > 1 || !unicode.IsDigit(rune(inpString[i])){
			break
		}
		indx++
	}
	outValue, _ = strconv.ParseFloat(inpString[:indx], 64)
	return outValue
}

//checkOrder check wether any line is not ordered
func (s *SortStrings)checkOrder(){
	for i := range s.InpLines{
		if s.InpLines[i] != s.SortedStrings[i]{
			fmt.Printf("sort: %d disorder: %s\n", i + 1, s.InpLines[i])
			return
		}
	}
}

//reversedOrder convert sorted in ascending order strings to descending order
func (s *SortStrings)reversedOrder(){
	i := len(s.SortedStrings) - 1
	j := 0
	for i > j {
		s.SortedStrings[i], s.SortedStrings[j] = s.SortedStrings[j], s.SortedStrings[i]
		i--
		j++
	}
}

//onlyUniqueLines removes string duplicates in sorted collection
func (s *SortStrings)onlyUniqueLines(){
	j := 1
	for i := 1; i < len(s.SortedStrings); i++ {
		if s.SortedStrings[i] != s.SortedStrings[i-1] {
			s.SortedStrings[j] = s.SortedStrings[i]
			j++
		}
	}
	s.SortedStrings = s.SortedStrings[:j]
}

//ReadFile reads slice of given fileNames and returns slice of all strings in the files
func (s *SortStrings)ReadFile(names []string) {
	s.InpLines = make([]string, 0)
	for _, v := range names {
		i, err := ioutil.ReadFile(v)
		if err != nil {
			log.Fatal(err)
		}
		s.InpLines = append(s.InpLines,  strings.Split(string(i), "\n")...)
	}
}