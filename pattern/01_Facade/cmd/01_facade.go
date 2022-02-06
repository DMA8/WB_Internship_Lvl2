package main

/*
Фасад — это структурный паттерн проектирования, который
предоставляет простой интерфейс к сложной системе классов, библиотеке или фреймворку

`Когда вам нужно представить простой или урезанный
интерфейс к сложной подсистеме.`

+ Изолирует клиентов от компонентов системы.
+ Уменьшает зависимость между подсистемой и клиентами.
- Фасад рискует стать божественным объектом,привязанным ко всем классам программы.

*/
import (
	"fmt"
	"../pkg"
)
type FilmWatcherCascade struct {
	MyKettle	pkg.Kettle
	MyFilm		pkg.Film
	MyCellPhone	pkg.CellPhone
}

func (f *FilmWatcherCascade) StartPreperationAndWatch(film string) {
	f.MyKettle.IsHot = true
	f.MyFilm.Name = film
	f.MyFilm.FindRefForFreeFilm()
	f.MyCellPhone.SilentMode = true
	fmt.Printf("Чайник горячий, %t\n", f.MyKettle.IsHot)
	fmt.Printf("Фильм нашелся по ссылке %s\n", f.MyFilm.Ref)
	fmt.Printf("Телефон не побеспокоит тебя, он в тихом режиме %t\n", f.MyCellPhone.SilentMode)
}

func main() {
	var myFascade FilmWatcherCascade
	myFascade.StartPreperationAndWatch("Зелёная миля")
}