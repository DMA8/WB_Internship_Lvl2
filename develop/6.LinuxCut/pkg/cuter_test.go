package pkg

import (
	"bufio"
	"io"
	"log"
	"os"
	"testing"
	"strings"
)
//cut -f2 -d "," bigCSV
func TestCutFiles(t *testing.T) {
	correctResult := "../texts/cut_f3_s_d_space_t" // cut -f3 -s -d' ' ../texts/t > cut_f3_s_d_space_t
	myCmdLine := NewCmdLine("-f3`-s`-d `../texts/t")
	myCuter := NewCuter(myCmdLine)
	myCuter.CutFiles()
	if !compareLines(correctResult, myCuter.Output) {
		t.Error("Test 1 failed")
	}

	correctResult2 := "../texts/cut_f2_-d_comma_bigCSV" // cut -f2 -d "," bigCSV > cut_f2_-d_comma_bigCSV
	myCmdLine2 := NewCmdLine("-f2`-d,`../texts/bigCSV")
	myCuter2 := NewCuter(myCmdLine2)
	myCuter2.CutFiles()
	if !compareLines(correctResult2, myCuter2.Output) {
		t.Error("Test 2 failed")
	}

}

func compareLines(file1 string, foundLines []string) bool {
	f1, err := os.Open(file1)
	if err != nil {
		log.Fatal(err)
	}
	defer f1.Close()
	file1Reader := bufio.NewReader(f1)
	for _, v := range foundLines {
		lineFromFile, err := file1Reader.ReadString('\n')
		if lineFromFile != "\n"{
			lineFromFile = strings.TrimSuffix(lineFromFile, "\n")
		}
		if err != nil {
			return false
		}
		if lineFromFile != v {
			return false
		}
	}
	_, err = file1Reader.ReadString('\n')
	return err == io.EOF
}
