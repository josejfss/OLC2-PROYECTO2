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
		Valor:     "[HEAP[(int)" + posicion.Valor + "]=" + val.Valor + ";",
		Tipo:      0}
	nomstack := "heap" + strconv.Itoa(objeto.ContadorHeap)
	objeto.ContadorHeap = objeto.ContadorHeap + 1
	block.Guardar_Declaracion(nomstack, simbstack)
	block.Guardar_Declaracion1(nomstack, simbstack)
	block.ListaTemporales.Add(nomstack)
	return nomstack
}
