package objetos

import (
	"fmt"

)

type Auto struct{
	
	Marca string
	Modelo string

}

func (a Auto) Cargar()string{
	return fmt.Sprintf("Cargando Auto, Marca:%s, Modelo:%s",a.Marca,a.Modelo)

}