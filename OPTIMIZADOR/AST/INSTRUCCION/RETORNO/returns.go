package retorno

import (
	objeto "OLC2-PROYECTO2/OPTIMIZADOR/OBJETO"
	"strconv"
)

type Returns struct {
	Reto string
}

func Nreturns(r string) Returns {
	ret := Returns{Reto: r}
	return ret
}

func (ret Returns) Optimizar_Instruccion(block *objeto.Bloque) interface{} {
	nom := "return" + strconv.Itoa(objeto.ContadorReturn)
	objeto.ContadorReturn = objeto.ContadorReturn + 1
	simbs := objeto.ObjetoBloque{
		Operacion: false,
		Temporal:  ret.Reto,
		Opiz:      ret.Reto,
		Opde:      ret.Reto,
		Valor:     ret.Reto + "\n}",
	}
	block.Guardar_Declaracion(nom, simbs)
	block.Guardar_Declaracion1(nom, simbs)
	block.ListaTemporales.Add(nom)
	return nom
}
