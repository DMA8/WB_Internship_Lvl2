package pkg

type Commander interface {
	Execute()
}

type OnStartCommand struct{
	MyEngine	*Engine
}

func (o *OnStartCommand)Execute(){
	o.MyEngine.On()
}

type OnSwitchOffCommand struct{
	MyEngine	*Engine
}

func (o *OnSwitchOffCommand)Execute(){
	o.MyEngine.Off()
}