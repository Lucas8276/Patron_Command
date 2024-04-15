package objetos

import ("fmt")
type Ups struct{
	Bateria bool
}

func (u*Ups) Cargar(){
	if u.Bateria{
	fmt.Printf("La bateria de la ups se encuentra en buen estado")
	}else{
		fmt.Printf("Se producira un cambio de bateria \n")
		u.Bateria=true}
		
}
