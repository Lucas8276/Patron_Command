package interface_

type Cargador_invoker	struct	{}

func (i * Cargador_invoker) EmitoComando(comando CargadorUniversalCommand){

	comando.Cargar()

}