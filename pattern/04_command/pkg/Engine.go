package pkg

type Engine struct{
	State	bool
}

func (e *Engine)On() {
	e.State = true
}

func (e *Engine)Off() {
	e.State = false
}