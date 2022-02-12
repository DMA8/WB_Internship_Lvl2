package main

import (
	"testing"
	"time"
)

func TestOr(t *testing.T) {
	start1 := time.Now()
	<-or(
		sig(50*time.Millisecond),
		sig(30*time.Millisecond),
		sig(20*time.Millisecond),
		sig(10*time.Millisecond),
	)
	if time.Since(start1) > time.Millisecond * 11 {
		t.Error("Test1 failed")
	}
}
