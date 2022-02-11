package main

import "fmt"

/*
Посетитель — это поведенческий паттерн,
который позволяет добавить новую операцию для целой иерархии классов,
не изменяя код этих классов.
*/

func main() {
	barsik := Cat{"Барсик"}
	sharik := Dog{"Шарик"}
	kesha := Bird{"Кеша"}
	voice := Visitor{}
	barsik.Accept(voice)
	sharik.Accept(voice)
	kesha.Accept(voice)

}

// type IAnimal interface{
// 	Move() // надо было добавить мув и войс
// 	Voice() // Проблема - чтобы структурки продолжали имплементировать интерфейс,
// 	//надо добавить метод Voice() во все структурки, НО есть принципы SOLID, Open-Close
// 	// который не советует изменять уже работающий код, а лишь расширять его
// }

type VoiceVisitor interface {
	VoiceC(cat Cat)
	VoiceD(dog Dog)
	VoiceB(bird Bird)
}

type Cat struct {
	Name	string
}

type Dog struct {
	Name	string
}

type Bird struct {
	Name	string	
}

type Visitor struct {}

func (v Visitor)VoiceC(cat Cat){
	fmt.Printf("My name is %s And Im doind Meoow \n", cat.Name)
}

func (v Visitor)VoiceD(dog Dog){
	fmt.Printf("My name is %s And Im doind Гав\n", dog.Name)
}

func (v Visitor)VoiceB(bird Bird){
	fmt.Printf("My name is %s And Im doind Чырык-чырык\n", bird.Name)
}

func (c Cat)Accept(visitor VoiceVisitor){
	visitor.VoiceC(c)
}

func (d Dog)Accept(visitor VoiceVisitor){
	visitor.VoiceD(d)
}

func (b Bird)Accept(visitor VoiceVisitor){
	visitor.VoiceB(b)
}

