package objetos

import "fmt"

type Celular struct {
	Marca          string
	Modelo         string
	NivelDeBateria int
}




func (c *Celular) Cargar() {
	if c.NivelDeBateria < 100 {
		c.NivelDeBateria = 100
	}
	fmt.Printf("Ahora tu celu esta cargado\n")
}
