package receiver

type Light struct {
	IsActive bool
}

func (r *Light) TurnOn() {
	r.IsActive = true
}

func (r *Light) TurnOff() {
	r.IsActive = false
}
