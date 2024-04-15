package objetostest_test

import (
	"fmt"
	"interface_cargador/interface_"
	"interface_cargador/objetos"
	"testing"
)

func TestCargar_Auto(t*testing.T){
	a := objetos.Auto{Marca: "Nissan",Modelo: "Gtr",Bateria: 30.7}
	invocadorsupremo := &interface_.Cargador_invoker{}
	invocadorsupremo.EmitoComando(&a)
	if a.Bateria==100{
		fmt.Printf("La bateria del auto se cargó")
	}
	
}

func TestCargar_Celular(t*testing.T){
	c := objetos.Celular{Marca: "Nokia",Modelo: "1100",NivelDeBateria: 57}
	invocadorsupremo := &interface_.Cargador_invoker{}
	invocadorsupremo.EmitoComando(&c)
	if c.NivelDeBateria==100{
	fmt.Printf("Se cargo el celular")
	}
}

func TestCargar_Ups(t*testing.T){
	u := objetos.Ups{Bateria: false}
	invocadorsupremo := &interface_.Cargador_invoker{}
	invocadorsupremo.EmitoComando(&u)
	if u.Bateria==true{
	fmt.Printf("Se hizo el remplaazo correspondiente")
	}
}

func TestCompleto(t*testing.T){
	a := objetos.Auto{Marca: "Nissan",Modelo: "Gtr",Bateria: 30.7}
	c := objetos.Celular{Marca: "Nokia",Modelo: "1100",NivelDeBateria: 57}
	u := objetos.Ups{Bateria: false}
	invocadorsupremo := &interface_.Cargador_invoker{}
	invocadorsupremo.EmitoComando(&a)
	invocadorsupremo.EmitoComando(&c)
	invocadorsupremo.EmitoComando(&u)
	if a.Bateria==100&& c.NivelDeBateria==100&& u.Bateria==true{
		fmt.Printf("Se paso con exito")
	}
}
/*
func TestCargar_Celular(t *testing.T) {
	c := objetos.Celular{Marca: "Samsung", Modelo: "Galaxy S21"}
	salidaCelular  := "Cargando Celular,Marca:Samsung,Modelo:Galaxy S21"

	if c.Cargar() != salidaCelular  {
		t.Errorf("Carga incorrecta, se esperaba '%s'",salidaCelular )
	}
}

func TestCargar_Auto(t *testing.T) {
	auto := objetos.Auto{Marca: "Toyota", Modelo: "Corolla"}
	salidaAuto := "Cargando Auto, Marca:Toyota, Modelo:Corolla" // Sin espacio adicional

	if mensaje_auto := auto.Cargar(); mensaje_auto != salidaAuto {
		t.Errorf("Carga incorrecta para Auto, se esperaba '%s ' pero se obtuvo '%s'", salidaAuto, mensaje_auto)
	}
}

func Test_AutoYCelular(t*testing.T){
	c := objetos.Celular{Marca: "Nokia",Modelo: "1100"}
	a :=objetos.Auto{Marca: "Fiat",	Modelo:"Uno"}
	
	salidaCelular := "Cargando Celular,Marca:Nokia,Modelo:1100" 
	salidaAuto := "Cargando Auto, Marca:Fiat, Modelo:Uno" 
	
	if mensaje_celular := c.Cargar();mensaje_celular!=salidaCelular {
		t.Errorf("Carga incorrecta para Celular, se esperaba '%s ' pero se obtuvo '%s'", salidaCelular, mensaje_celular)
	}

	if mensaje_auto := a.Cargar();mensaje_auto!=salidaAuto {
		t.Errorf("Carga incorrecta para Auto, se esperaba '%s ' pero se obtuvo '%s'", salidaAuto, mensaje_auto)
	}	
}
*/