package sif

import (
	interfaz "OLC2-PROYECTO2/OPTIMIZADOR/INTERFAZ"
	objeto "OLC2-PROYECTO2/OPTIMIZADOR/OBJETO"
	"strconv"
)

type SenteIf struct {
	Condicion interfaz.Expresion
	Etiqueta  string
}

func Nsenteif(con interfaz.Expresion, eti string) SenteIf {
	sif := SenteIf{Condicion: con, Etiqueta: eti}
	return sif
}

func (sif SenteIf) Optimizar_Instruccion(block *objeto.Bloque) interface{} {
	conds := sif.Condicion.Optimizar_Expresion(block)
	simbif := objeto.ObjetoBloque{Operacion: false, Opiz: conds.Opiz, Opde: conds.Opde, Valor: "if (" + conds.Valor + ") goto " + sif.Etiqueta + ";"}
	nomif := "if" + strconv.Itoa(objeto.ContadorIf)
	objeto.ContadorIf = objeto.ContadorIf + 1
	block.Guardar_Declaracion(nomif, simbif)
	return 0
}
