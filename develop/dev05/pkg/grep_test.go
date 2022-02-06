package pkg

import (
	"bufio"
	"io"
	"log"
	"os"
	"testing"
	"strings"
)

func TestGreperMain(t *testing.T) {
	correctResult := "../texts/correctGrep_C3_context_t" // grep -C3 context ../texts/t > correctGrep_C3_context_t
	myCmdLine := NewCmdLine("-C3 context ../texts/t")
	myGreper := NewGreper(myCmdLine)
	myGreper.Main()
	if !compareLines(correctResult, myGreper.Output) {
		t.Error("Test 1 failed")
	}

	correctResult2 := "../texts/correctGrep_v_F_point_USA" // grep -v -F .  ../texts/USA_constitution > correctGrep_v_F_point_USA
	myCmdLine2 := NewCmdLine("-v -F . ../texts/USA_constitution")
	myGreper2 := NewGreper(myCmdLine2)
	myGreper2.Main()
	if !compareLines(correctResult2, myGreper2.Output) {
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
		lineFromFile = strings.TrimSuffix(lineFromFile, "\n")
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
