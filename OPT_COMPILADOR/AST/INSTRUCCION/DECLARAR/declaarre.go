package declarar

import (
	reportes "OLC2-PROYECTO2/COMPILADOR/REPORTES"
	bloque "OLC2-PROYECTO2/OPT_COMPILADOR/BLOQUE"
	generando "OLC2-PROYECTO2/OPT_COMPILADOR/GENERANDO"
	inter "OLC2-PROYECTO2/OPT_COMPILADOR/INTER"
	"strconv"
)

type DeclaPStack struct {
	Nombre string
	Valor  inter.Expresion
	Linea  int
}

func Ndeclapstack(nom string, val inter.Expresion, lin int) DeclaPStack {
	pstack := DeclaPStack{Nombre: nom, Valor: val, Linea: lin}
	return pstack
}

func (got DeclaPStack) Opti_Instruccion(block *bloque.BloquesOpt, genopt *generando.GenerandoOpti) interface{} {
	vali := got.Valor.Opti_Expresion(block)
	op_iz := block.Obtener_Temporal(vali.OpeIz)

	if op_iz.Tipo == bloque.TEMPORAL {
		reportes.ReporteOpti("Bloques", "Regla 2", got.Nombre+"="+vali.Valor+";", got.Nombre+"="+op_iz.Asignacion+";", strconv.Itoa(got.Linea))
		genopt.EliminarTempo(op_iz.Asignacion, op_iz.Valor, strconv.Itoa(got.Linea))
		block.EliminarTemporalLista(op_iz.Asignacion)

		guardar := bloque.Nobjetolista("asignacion", got.Nombre+"="+op_iz.Asignacion+";")
		genopt.AgegarCodigoOpt(guardar)
	} else if op_iz.Tipo == bloque.CONSTANTE {
		reportes.ReporteOpti("Bloques", "Regla 4", got.Nombre+"="+vali.Valor+";", got.Nombre+"="+op_iz.Asignacion+";", strconv.Itoa(got.Linea))
		genopt.EliminarTempo(op_iz.Asignacion, op_iz.Valor, strconv.Itoa(got.Linea))
		block.EliminarTemporalLista(op_iz.Asignacion)

		guardar := bloque.Nobjetolista("asignacion", got.Nombre+"="+op_iz.Asignacion+";")
		genopt.AgegarCodigoOpt(guardar)
	} else {
		guardar := bloque.Nobjetolista("asignacion", got.Nombre+"="+vali.Valor+";")
		genopt.AgegarCodigoOpt(guardar)
	}
	return 0
}
