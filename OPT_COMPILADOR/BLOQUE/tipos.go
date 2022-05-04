package bloque

type TipoObjeto int

const (
	INTEGER TipoObjeto = iota
	FLOAT
	TEMPORAL
	ETIQUETA
	PUNSTACK
	PUNHEAP
	OPERACION
	CONSTANTE
	NULL
)
