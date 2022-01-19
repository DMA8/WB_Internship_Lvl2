package main

import "testing"
func TestGetTime(t *testing.T){
	if _, err := GetTime(); err != nil {
		t.Error("there's an error here ", err)
	}
}