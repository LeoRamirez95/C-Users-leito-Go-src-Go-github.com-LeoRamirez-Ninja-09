package main

import "fmt"

type persona struct {
	nombre string
}

func (p *persona) hablar() {
	fmt.Println("Hola desde persona", p.nombre)
}

type humano interface {
	hablar()
}

func diAlgo(h humano) {
	h.hablar()
}

func main() {
	a := persona{
		nombre: "Leo ramirez",
	}
	//diAlgo(&a)
	a.hablar()

}
