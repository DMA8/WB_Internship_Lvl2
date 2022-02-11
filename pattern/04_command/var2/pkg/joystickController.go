package pkg

type Commander interface {
	Execute(Actor)
}

type Actor interface{
	Jump()
}


// XYAB - клавиши джойстика
type Controller struct {
	X	Commander
	Y	Commander
	A	Commander
	B	Commander
}

func NewController(X, Y, A, B Commander) *Controller{
	return &Controller{
		X: X,
		Y: Y,
		A: A,
		B: B,
	}
}

func (c Controller)InpCommandHandler(cmd string, actor Actor){
	switch cmd{
	case "X":
		c.X.Execute(actor)
	case "Y":
		c.Y.Execute(actor)
	case "A":
		c.A.Execute(actor)
	case "B":
		c.B.Execute(actor)
	}
}
