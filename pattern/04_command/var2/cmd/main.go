package main

import (
	"../pkg"
)
//https://martalex.gitbooks.io/gameprogrammingpatterns/content/chapter-2/2.1-command.html
func main() {
	jumpCmd := pkg.JumpCommand{} //создали команду
	hero := pkg.GameActor{Name:"Hero"}//бизнес логика
	controller := pkg.NewController(jumpCmd, nil, nil, nil)//интерфейс
	controller.InpCommandHandler("X", hero)//связалли интерфейс с коммандой
}