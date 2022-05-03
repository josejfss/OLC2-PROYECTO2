package imprimir

import (
	interfaz "OLC2-PROYECTO2/OPTIMIZADOR/INTERFAZ"
	objeto "OLC2-PROYECTO2/OPTIMIZADOR/OBJETO"
	"strconv"
)

type Imprimir struct {
	Cadena string
	Casteo string
	Expr   interfaz.Expresion
}

func Nimprimir(cad string, cas string, exp interfaz.Expresion) Imprimir {
	print := Imprimir{Cadena: cad, Casteo: cas, Expr: exp}
	return print
}

func (print Imprimir) Optimizar_Instruccion(block *objeto.Bloque) interface{} {
	exp := print.Expr.Optimizar_Expresion(block)
	prints := objeto.ObjetoBloque{Operacion: false, Opiz: exp.Opiz, Opde: exp.Opde, Valor: "printf(" + print.Cadena + ", " + print.Casteo + exp.Valor + ");"}
	nprint := "print" + strconv.Itoa(objeto.ContadorPrint)
	objeto.ContadorPrint = objeto.ContadorPrint + 1
	block.Guardar_Declaracion(nprint, prints)
	return 0
}
