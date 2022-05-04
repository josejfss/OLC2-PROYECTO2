package sif

import (
	reportes "OLC2-PROYECTO2/COMPILADOR/REPORTES"
	bloque "OLC2-PROYECTO2/OPT_COMPILADOR/BLOQUE"
	generando "OLC2-PROYECTO2/OPT_COMPILADOR/GENERANDO"
	inter "OLC2-PROYECTO2/OPT_COMPILADOR/INTER"
	"strconv"
)

type SenteIf struct {
	Condicion inter.Expresion
	Etiqueta  string
	Linea     int
}

func Nsenteif(con inter.Expresion, eti string, lin int) SenteIf {
	sif := SenteIf{Condicion: con, Etiqueta: eti, Linea: lin}
	return sif
}

func (sif SenteIf) Opti_Instruccion(block *bloque.BloquesOpt, genopt *generando.GenerandoOpti) interface{} {
	conds := sif.Condicion.Opti_Expresion(block)
	op_iz := block.Obtener_Temporal(conds.OpeIz)
	op_de := block.Obtener_Temporal(conds.OpeDe)
	if op_iz.Tipo == bloque.TEMPORAL {
		reportes.ReporteOpti("Bloques", "Regla 2", "if("+conds.Valor+") goto "+sif.Etiqueta+";", "if("+op_iz.Valor+conds.Ope+conds.OpeDe+") goto "+sif.Etiqueta+";", strconv.Itoa(sif.Linea))
		genopt.EliminarTempo(op_iz.Asignacion, op_iz.Valor, strconv.Itoa(sif.Linea))
		block.EliminarTemporalLista(op_iz.Asignacion)

		guardar := bloque.Nobjetolista("if", "if("+op_iz.Valor+conds.Ope+conds.OpeDe+") goto "+sif.Etiqueta+";")
		genopt.AgegarCodigoOpt(guardar)
	} else if op_de.Tipo == bloque.TEMPORAL {
		reportes.ReporteOpti("Bloques", "Regla 2", "if("+conds.Valor+") goto "+sif.Etiqueta+";", "if("+conds.OpeIz+conds.Ope+op_de.Valor+") goto "+sif.Etiqueta+";", strconv.Itoa(sif.Linea))
		genopt.EliminarTempo(op_iz.Asignacion, op_iz.Valor, strconv.Itoa(sif.Linea))
		block.EliminarTemporalLista(op_iz.Asignacion)

		guardar := bloque.Nobjetolista("if", "if("+conds.OpeIz+conds.Ope+op_de.Valor+") goto "+sif.Etiqueta+";")
		genopt.AgegarCodigoOpt(guardar)
	} else if op_iz.Tipo == bloque.CONSTANTE {
		reportes.ReporteOpti("Bloques", "Regla 4", "if("+conds.Valor+") goto "+sif.Etiqueta+";", "if("+op_iz.Valor+conds.Ope+conds.OpeDe+") goto "+sif.Etiqueta+";", strconv.Itoa(sif.Linea))
		genopt.EliminarTempo(op_iz.Asignacion, op_iz.Valor, strconv.Itoa(sif.Linea))
		block.EliminarTemporalLista(op_iz.Asignacion)

		guardar := bloque.Nobjetolista("if", "if("+op_iz.Valor+conds.Ope+conds.OpeDe+") goto "+sif.Etiqueta+";")
		genopt.AgegarCodigoOpt(guardar)
	} else if op_de.Tipo == bloque.CONSTANTE {
		reportes.ReporteOpti("Bloques", "Regla 4", "if("+conds.Valor+") goto "+sif.Etiqueta+";", "if("+conds.OpeIz+conds.Ope+op_de.Valor+") goto "+sif.Etiqueta+";", strconv.Itoa(sif.Linea))
		genopt.EliminarTempo(op_iz.Asignacion, op_iz.Valor, strconv.Itoa(sif.Linea))
		block.EliminarTemporalLista(op_iz.Asignacion)

		guardar := bloque.Nobjetolista("if", "if("+conds.OpeIz+conds.Ope+op_de.Valor+") goto "+sif.Etiqueta+";")
		genopt.AgegarCodigoOpt(guardar)
	} else {
		guardar := bloque.Nobjetolista("if", "if("+conds.Valor+") goto "+sif.Etiqueta+";")
		genopt.AgegarCodigoOpt(guardar)
	}
	return 0
}
