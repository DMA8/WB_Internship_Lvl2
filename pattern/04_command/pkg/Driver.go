package pkg

type Driver struct{
	CommandQueue	[]Commander
}

func (d Driver)Execute(){
	for _, v := range d.CommandQueue {
		v.Execute()
	} 
}