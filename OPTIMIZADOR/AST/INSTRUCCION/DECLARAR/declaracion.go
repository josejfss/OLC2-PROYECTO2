package declarar

import (
	interfaz "OLC2-PROYECTO2/OPTIMIZADOR/INTERFAZ"
	objeto "OLC2-PROYECTO2/OPTIMIZADOR/OBJETO"
)

type Declaracion struct {
	Temporal string
	Valor    interfaz.Expresion
	Linea    int
}

func Ndeclaracion(tem string, val interfaz.Expresion, lin int) Declaracion {
	decla := Declaracion{Temporal: tem, Valor: val, Linea: lin}
	return decla
}

func (decla Declaracion) Optimizar_Instruccion(block *objeto.Bloque) interface{} {
	valopt := decla.Valor.Optimizar_Expresion(block)
	simbdecla := objeto.ObjetoBloque{
		Operacion:  valopt.Operacion,
		EsTemporal: valopt.EsTemporal,
		Constante:  valopt.Constante,
		Temporal:   decla.Temporal,
		Opiz:       valopt.Opiz,
		Opde:       valopt.Opde,
		Ope:        valopt.Ope,
		Valor:      valopt.Valor,
		Tipo:       1,
		Linea:      decla.Linea,
	}
	block.Guardar_Declaracion(decla.Temporal, simbdecla)
	block.Guardar_Declaracion1(decla.Temporal, simbdecla)
	block.ListaTemporales.Add(decla.Temporal)
	return decla.Temporal
}