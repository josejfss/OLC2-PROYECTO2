package imprimir

import (
	reportes "OLC2-PROYECTO2/COMPILADOR/REPORTES"
	bloque "OLC2-PROYECTO2/OPT_COMPILADOR/BLOQUE"
	generando "OLC2-PROYECTO2/OPT_COMPILADOR/GENERANDO"
	inter "OLC2-PROYECTO2/OPT_COMPILADOR/INTER"
	"strconv"
)

type Imprimir struct {
	Cadena string
	Casteo string
	Expr   inter.Expresion
	Linea  int
}

func Nimprimir(cad string, cas string, exp inter.Expresion, lin int) Imprimir {
	print := Imprimir{Cadena: cad, Casteo: cas, Expr: exp, Linea: lin}
	return print
}

func (print Imprimir) Opti_Instruccion(block *bloque.BloquesOpt, genopt *generando.GenerandoOpti) interface{} {
	exp := print.Expr.Opti_Expresion(block)
	op_iz := block.Obtener_Temporal(exp.OpeIz)
	if op_iz.Tipo == bloque.TEMPORAL {
		reportes.ReporteOpti("Bloques", "Regla 2", "printf("+print.Cadena+", "+print.Casteo+exp.Valor+");", "printf("+print.Cadena+", "+print.Casteo+op_iz.Valor+");", strconv.Itoa(print.Linea))
		genopt.EliminarTempo(op_iz.Asignacion, op_iz.Valor, strconv.Itoa(print.Linea))
		block.EliminarTemporalLista(op_iz.Asignacion)

		guardar := bloque.Nobjetolista("print", "printf("+print.Cadena+", "+print.Casteo+op_iz.Valor+");")
		genopt.AgegarCodigoOpt(guardar)
	} else if op_iz.Tipo == bloque.CONSTANTE {
		reportes.ReporteOpti("Bloques", "Regla 4", "printf("+print.Cadena+", "+print.Casteo+exp.Valor+");", "printf("+print.Cadena+", "+print.Casteo+op_iz.Valor+");", strconv.Itoa(print.Linea))
		genopt.EliminarTempo(op_iz.Asignacion, op_iz.Valor, strconv.Itoa(print.Linea))
		block.EliminarTemporalLista(op_iz.Asignacion)

		guardar := bloque.Nobjetolista("print", "printf("+print.Cadena+", "+print.Casteo+op_iz.Valor+");")
		genopt.AgegarCodigoOpt(guardar)
	} else {
		guardar := bloque.Nobjetolista("print", "printf("+print.Cadena+", "+print.Casteo+exp.Valor+");")
		genopt.AgegarCodigoOpt(guardar)
	}
	return 0
}
