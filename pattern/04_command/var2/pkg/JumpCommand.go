package pkg

type JumpCommand struct {} // мы материализовали команду Jump в объект и теперь можем работать с ней как с объектом
// Можем отложить выполнение, отменить его или повторить

func (j JumpCommand)Execute(actor Actor) {
	actor.Jump()
}