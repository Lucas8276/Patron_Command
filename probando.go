package main

import (
	
	"interface_cargador/interface_"
	"interface_cargador/objetos"
)

func main() {
	invocadorsupremo:= &interface_.Cargador_invoker{}
	auto:= objetos.Auto{Marca: "Toyota",Modelo: "Corolla\n"}
	celular:= objetos.Celular{Marca: "Samsung",Modelo: "A54"}

	invocadorsupremo.EmitoComando(auto)
	invocadorsupremo.EmitoComando(celular)
}

