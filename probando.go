
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
/*
package main

import "fmt"

// Request representa una solicitud de compra
type Request struct {
    Item     string
    Quantity int
}

// Handler define la interfaz para manejar solicitudes
type Handler interface {
    HandleRequest(req *Request)
    SetNext(handler Handler)
}

// ConcreteHandler implementa la interfaz Handler
type ConcreteHandler struct {
    Next    Handler
    Command Command
}

func (h *ConcreteHandler) SetNext(handler Handler) {
    h.Next = handler
}

func (h *ConcreteHandler) HandleRequest(req *Request) {
    if h.Command.CanHandle(req) {
        h.Command.Execute(req)
    } else if h.Next != nil {
        fmt.Printf("%s no puede manejar la solicitud de %d unidades de %s. Pasando al siguiente nivel.\n", h.Command, req.Quantity, req.Item)
        h.Next.HandleRequest(req)
    } else {
        fmt.Printf("Ningún nivel puede manejar la solicitud de %d unidades de %s.\n", req.Quantity, req.Item)
    }
}

// Command define la interfaz para los comandos
type Command interface {
    CanHandle(req *Request) bool
    Execute(req *Request)
}

// ConcreteCommand implementa la interfaz Command
type ConcreteCommand struct {
    Threshold int
    Stock     map[string]int
    Name      string
}

func (c *ConcreteCommand) CanHandle(req *Request) bool {
    stock, ok := c.Stock[req.Item]
    return ok && req.Quantity <= stock && req.Quantity <= c.Threshold
}

func (c *ConcreteCommand) Execute(req *Request) {
    fmt.Printf("Compra de %d unidades de %s aprobada por %s.\n", req.Quantity, req.Item, c.Name)
    c.Stock[req.Item] -= req.Quantity
}

func main() {
    // Inicializar el stock de productos
    stock := map[string]int{
        "Camiseta": 10,
        "Pantalón": 5,
        "Zapatos":  3,
    }

    // Crear una cadena de comandos para manejar las solicitudes de compra
    command1 := &ConcreteCommand{Threshold: 5, Stock: stock, Name: "Nivel 1"}
    command2 := &ConcreteCommand{Threshold: 10, Stock: stock, Name: "Nivel 2"}
    command3 := &ConcreteCommand{Threshold: 20, Stock: stock, Name: "Nivel 3"}

    // Configurar la cadena de responsabilidad
    handler1 := &ConcreteHandler{Command: command1}
    handler2 := &ConcreteHandler{Command: command2}
    handler3 := &ConcreteHandler{Command: command3}

    handler1.SetNext(handler2)
    handler2.SetNext(handler3)

    // Procesar solicitudes de compra
    requests := []*Request{
        {Item: "Camiseta", Quantity: 3},
        {Item: "Pantalón", Quantity: 8},
        {Item: "Zapatos", Quantity: 4},
        {Item: "Corbata", Quantity: 2}, // Producto no disponible en el stock
    }

    for _, req := range requests {
        fmt.Printf("Procesando solicitud de %d unidades de %s...\n", req.Quantity, req.Item)
        handler1.HandleRequest(req)
        fmt.Println("------------------------------")
    }

    // Mostrar el stock actualizado
    fmt.Println("Stock actualizado:")
    for item, quantity := range stock {
        fmt.Printf("%s: %d\n", item, quantity)
    }
}
*/
