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
	simbif := objeto.ObjetoBloque{Operacion: true,
		Temporal: sif.Etiqueta,
		Opiz:     conds.Opiz,
		Opde:     conds.Opde,
		Ope:      conds.Ope,
		Valor:    "if (" + conds.Valor + ") goto " + sif.Etiqueta + ";",
		Tipo:     2}
	nomif := "if" + strconv.Itoa(objeto.ContadorIf)
	objeto.ContadorIf = objeto.ContadorIf + 1
	block.Guardar_Declaracion(nomif, simbif)
	block.Guardar_Declaracion1(nomif, simbif)
	block.ListaTemporales.Add(nomif)
	return nomif
}
