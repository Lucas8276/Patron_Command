package objetostest_test

import (
	"interface_cargador/objetos"

	"testing"
)

func TestCargar_Celular(t *testing.T) {
	c := objetos.Celular{Marca: "Samsung", Modelo: "Galaxy S21"}
	expectedOutput := "Cargando Celular,Marca:Samsung,Modelo:Galaxy S21"

	if c.Cargar() != expectedOutput {
		t.Errorf("Carga incorrecta, se esperaba '%s'", expectedOutput)
	}
}

func TestCargar_Auto(t *testing.T) {
	auto := objetos.Auto{Marca: "Toyota", Modelo: "Corolla"}
	expectedOutput := "Cargando Auto, Marca:Toyota, Modelo:Corolla" // Sin espacio adicional

	if output := auto.Cargar(); output != expectedOutput {
		t.Errorf("Carga incorrecta para Auto, se esperaba '%s ' pero se obtuvo '%s'", expectedOutput, output)
	}
}
