package declarar

import (
	reportes "OLC2-PROYECTO2/COMPILADOR/REPORTES"
	bloque "OLC2-PROYECTO2/OPT_COMPILADOR/BLOQUE"
	generando "OLC2-PROYECTO2/OPT_COMPILADOR/GENERANDO"
	inter "OLC2-PROYECTO2/OPT_COMPILADOR/INTER"
	"strconv"
)

type DeclaPunteros struct {
	Nombre   string
	Posicion inter.Expresion
	Valor    inter.Expresion
	Linea    int
}

func Ndeclapunteros(nom string, pos inter.Expresion, val inter.Expresion, lin int) DeclaPunteros {
	pstack := DeclaPunteros{Nombre: nom, Posicion: pos, Valor: val, Linea: lin}
	return pstack
}

func (got DeclaPunteros) Opti_Instruccion(block *bloque.BloquesOpt, genopt *generando.GenerandoOpti) interface{} {
	vali := got.Valor.Opti_Expresion(block)
	pos := got.Posicion.Opti_Expresion(block)
	op_iz := block.Obtener_Temporal(vali.OpeIz)

	if op_iz.Tipo == bloque.TEMPORAL {
		reportes.ReporteOpti("Bloques", "Regla 2", got.Nombre+"[(int)"+pos.Valor+"]="+vali.Valor+";", got.Nombre+"[(int)"+pos.Valor+"]="+op_iz.Asignacion+";", strconv.Itoa(got.Linea))
		genopt.EliminarTempo(op_iz.Asignacion, op_iz.Valor, strconv.Itoa(got.Linea))
		block.EliminarTemporalLista(op_iz.Asignacion)

		guardar := bloque.Nobjetolista("asignacion", got.Nombre+"[(int)"+pos.Valor+"]="+op_iz.Asignacion+";")
		genopt.AgegarCodigoOpt(guardar)
	} else {
		guardar := bloque.Nobjetolista("asignacion", got.Nombre+"[(int)"+pos.Valor+"]="+vali.Valor+";")
		genopt.AgegarCodigoOpt(guardar)
	}
	return 0
}
