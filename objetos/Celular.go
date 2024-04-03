package objetos

import (
	"fmt"
	
)

type Celular struct{

	Marca string
	Modelo string

}

func (c Celular) Cargar()string{
	return fmt.Sprintf("Cargando Celular,Marca:%s,Modelo:%s",c.Marca,c.Modelo)
	
}