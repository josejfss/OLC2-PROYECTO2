package simbolos

import "github.com/colegno/arraylist"

type ValoresC3D struct {
	Valor           string
	EsTemporal      bool
	Tipo            TipoExpresion
	Label_verdadera string
	Label_false     string
}

type Etiquetas struct {
	EtiquetaVerdadera string
	EtiquetaFalsa     string
}

var ListaTransferencia = arraylist.New()

type EtiquetasTransferencia struct {
	Etiqueta_Entrada string
	Etiqueta_Salida  string
}
