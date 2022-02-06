package main

import "fmt"

/*
Посетитель — это поведенческий паттерн,
который позволяет добавить новую операцию для целой иерархии классов,
не изменяя код этих классов.
*/

func main() {
	var interfaceVisitor VoiceVisitor
	barsik := Cat{"Барсик"}
	voice := Visitor{}
	interfaceVisitor = voice
	barsik.Accept(interfaceVisitor)

}

type IAnimal interface{
	Move() // Не было ни одного метода, но надо было добавить мув и войс
	Voice() // Проблема - чтобы структурки продолжали имплементировать интерфейс,
	//надо добавить метод Voice() во все структурки, НО есть принципы SOLID, Open-Close
	// который не советует изменять уже работающий код, а лишь расширять его
}

type VoiceVisitor interface {
	VoiceC(cat Cat)
	VoiceD(dog Dog)
	VoiceB(bird Bird)
}

type MoveVisitor interface {
	MoveC(cat Cat)
	MoveD(dog Dog)
	MoveB(bird Bird)

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

type Visitor struct {

}

func (v Visitor)VoiceC(cat Cat){
	fmt.Printf("My name is %s And Im doind Meoow \n", cat.Name)
}

func (v Visitor)VoiceD(dog Dog){
	fmt.Println("Гав")
}

func (v Visitor)VoiceB(bird Bird){
	fmt.Println("Чырык-чырык")
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

func (c Cat)Move(){
	fmt.Println("Крадется")
}

func (c Dog)Move(){
	fmt.Println("Бежит")
}
func (c Bird)Move(){
	fmt.Println("Летит")
}

