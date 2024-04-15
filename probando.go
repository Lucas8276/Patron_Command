package main

import (
	
	"interface_cargador/interface_"
	"interface_cargador/objetos"

)

func main() {
	invocadorsupremo := &interface_.Cargador_invoker{}
	auto := objetos.Auto{Marca: "Toyota", Modelo: "Corolla", Bateria: 56}
	celular := objetos.Celular{Marca: "Samsung", Modelo: "A54", NivelDeBateria: 56}
	ups:= objetos.Ups{Bateria:false }
	invocadorsupremo.EmitoComando(&auto)
	invocadorsupremo.EmitoComando(&celular)
	invocadorsupremo.EmitoComando(&ups)
	
}
