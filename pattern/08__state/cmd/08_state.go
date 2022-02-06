package main

import (
	"../pkg"
	"fmt"
)
	/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern
*/


//конечный автомат
/*
Состояние — это поведенческий паттерн,
позволяющий динамически изменять поведение объекта при смене его состояния.
*/
func main(){
	firstWarior := pkg.NewWarior("first", 100, 100)
	secondWarior := pkg.NewWarior("second", 100, 100)

	firstWarior.DoDamage(99, secondWarior)
	secondWarior.Heal(100)
	firstWarior.DoDamage(100, secondWarior)
	firstWarior.DoDamage(50, secondWarior)
	firstWarior.DoDamage(50, secondWarior)

	fmt.Println(secondWarior)

}