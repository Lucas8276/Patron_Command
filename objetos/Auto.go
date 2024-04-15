package objetos

import "fmt"

type Auto struct {
	Marca   string
	Modelo  string
	Bateria float64
}



func (a*Auto) Cargar() {
	if a.Bateria < 100.00 {
		fmt.Printf("Cargar bateria del modelo %s, de la marca %s \n",a.Modelo,a.Marca)
		a.Bateria=100
	}
}
