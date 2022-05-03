package declarar

import (
	interfaz "OLC2-PROYECTO2/OPTIMIZADOR/INTERFAZ"
	objeto "OLC2-PROYECTO2/OPTIMIZADOR/OBJETO"
	"strconv"
)

type DeclaHeap struct {
	Pos       interfaz.Expresion
	Expresion interfaz.Expresion
}

func Ndeclaheap(p interfaz.Expresion, exp interfaz.Expresion) DeclaHeap {
	dheap := DeclaHeap{Pos: p, Expresion: exp}
	return dheap
}

func (dheap DeclaHeap) Optimizar_Instruccion(block *objeto.Bloque) interface{} {
	posicion := dheap.Pos.Optimizar_Expresion(block)
	val := dheap.Expresion.Optimizar_Expresion(block)
	simbstack := objeto.ObjetoBloque{
		Operacion: val.Operacion,
		Temporal:  "HEAP",
		Opiz:      val.Opiz,
		Opde:      val.Opde,
		Valor:     "[HEAP[(int)" + posicion.Valor + "]=" + val.Valor + ";"}
	nomstack := "stack" + strconv.Itoa(objeto.ContadorStack)
	objeto.ContadorStack = objeto.ContadorStack + 1
	block.Guardar_Declaracion(nomstack, simbstack)
	return 0
}
