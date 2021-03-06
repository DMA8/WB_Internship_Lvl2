/*
=== Базовая задача ===

Создать программу печатающую точное время с использованием NTP библиотеки.Инициализировать как go module.
Использовать библиотеку https://github.com/beevik/ntp.
Написать программу печатающую текущее время / точное время с использованием этой библиотеки.

Программа должна быть оформлена с использованием как go module.
Программа должна корректно обрабатывать ошибки библиотеки: распечатывать их в STDERR и возвращать ненулевой код выхода в OS.
Программа должна проходить проверки go vet и golint.
*/
package main

import (
	"fmt"
	"log"
	"time"
	"github.com/beevik/ntp"
)

func main() {
	time, err := GetTime()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(time)
}

//GetTime provides curent time from remote serv through ntp
func GetTime()(time.Time, error){
	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	return time, err
}
