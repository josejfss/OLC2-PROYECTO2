package declarar

import (
	interfaz "OLC2-PROYECTO2/OPTIMIZADOR/INTERFAZ"
	objeto "OLC2-PROYECTO2/OPTIMIZADOR/OBJETO"
	"strconv"
)

type DeclaStack struct {
	Pos       interfaz.Expresion
	Expresion interfaz.Expresion
}

func Ndeclastack(p interfaz.Expresion, exp interfaz.Expresion) DeclaStack {
	dstack := DeclaStack{Pos: p, Expresion: exp}
	return dstack
}

func (dstack DeclaStack) Optimizar_Instruccion(block *objeto.Bloque) interface{} {
	posicion := dstack.Pos.Optimizar_Expresion(block)
	val := dstack.Expresion.Optimizar_Expresion(block)
	simbstack := objeto.ObjetoBloque{
		Operacion: val.Operacion,
		Temporal:  "STACK",
		Opiz:      val.Opiz,
		Opde:      val.Opde,
		Valor:     "[STACK[(int)" + posicion.Valor + "]=" + val.Valor + ";",
		Tipo:      0}
	nomstack := "stack" + strconv.Itoa(objeto.ContadorStack)
	objeto.ContadorStack = objeto.ContadorStack + 1
	block.Guardar_Declaracion(nomstack, simbstack)
	return nomstack
}
