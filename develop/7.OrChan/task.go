package main

import (
	"fmt"
	"time"
	//"sync"
)
// struct{} - используется для подчеркивания того, что канал не передает данные,
// а лишь сигнализирует о событии.

func main(){
	start := time.Now()
	<-or (
	sig(21*time.Second),
	sig(51*time.Second),
	sig(12*time.Second),
	sig(33*time.Second),
	sig(41*time.Second),
	sig(41*time.Hour),
	)
	fmt.Printf("done after %v\n", time.Since(start))

}

func or (channels ... <- chan interface{}) <- chan interface{}{
	done := make(chan interface{})

	for _, v := range channels{
		go func(a <- chan interface{}){
			if _, ok := <- a; !ok {
				done <- 1
			}
		}(v)
	}
	return done
}

func sig(after time.Duration) <- chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(after)
	}()
	return c
}
