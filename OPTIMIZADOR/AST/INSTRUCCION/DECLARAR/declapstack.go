package declarar

import (
	interfaz "OLC2-PROYECTO2/OPTIMIZADOR/INTERFAZ"
	objeto "OLC2-PROYECTO2/OPTIMIZADOR/OBJETO"
)

type DeclaPStack struct {
	Nombre string
	Valor  interfaz.Expresion
}

func Ndeclapstack(nom string, val interfaz.Expresion) DeclaPStack {
	pstack := DeclaPStack{Nombre: nom, Valor: val}
	return pstack
}

func (decla DeclaPStack) Optimizar_Instruccion(block *objeto.Bloque) interface{} {
	valopt := decla.Valor.Optimizar_Expresion(block)
	simbdecla := objeto.ObjetoBloque{
		Operacion: valopt.Operacion,
		Temporal:  decla.Nombre,
		Opiz:      valopt.Opiz,
		Opde:      valopt.Opde,
		Ope:       valopt.Ope,
		Valor:     decla.Nombre + "=" + valopt.Valor + ";",
		Tipo:      0}
	block.Guardar_Declaracion(decla.Nombre, simbdecla)
	return decla.Nombre
}
