package retorno

import objeto "OLC2-PROYECTO2/OPTIMIZADOR/OBJETO"

type Returns struct {
	Reto string
}

func Nreturns(r string) Returns {
	ret := Returns{Reto: r}
	return ret
}

func (ret Returns) Optimizar_Instruccion(block *objeto.Bloque) interface{} {
	return ret.Reto
}
