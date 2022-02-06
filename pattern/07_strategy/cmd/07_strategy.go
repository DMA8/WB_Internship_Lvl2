package main

import (
	"fmt"

	"../pkg"
)

/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern
*/
func main(){
	salaryTable := pkg.SalaryTable{map[string]int{"Petrov": 12, "Sidorov": 18, "Myhtarov": 5}, nil}
	decreseALot := pkg.DecreaserSalaryBiggest{}
	decreseALittle := pkg.DecreaserSalarySmallest{}
	howMuchToOptimize:="big"
	if howMuchToOptimize == "little" {
		salaryTable.OptimizationAlgo = decreseALittle
	} else if howMuchToOptimize == "big" {
		salaryTable.OptimizationAlgo = decreseALot
	}
	salaryTable.Optimize()
	fmt.Println(salaryTable.Salary)
}