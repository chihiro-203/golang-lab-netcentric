package main

type rect struct {
	width, height int
}

func (r rect) changeByVal() {
	r.width = 0
	r.height = 0
}

func (r rect) changeByRef() {
	r.width = 0
	r.height = 0
}

func main() {
}
