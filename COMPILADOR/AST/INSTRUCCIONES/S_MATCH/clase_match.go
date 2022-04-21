package smatch

import (
	interfaces "OLC2-PROYECTO2/COMPILADOR/INTERFACES"

	"github.com/colegno/arraylist"
)

type Matches struct {
	Condicion_match   interfaces.Expresion
	L_instr_match     *arraylist.List
	Monton            bool
	Lista_condiciones *arraylist.List
}

func Nmatches(con interfaces.Expresion, lmatch *arraylist.List, mt bool, lcon *arraylist.List) Matches {
	matchado := Matches{Condicion_match: con, L_instr_match: lmatch, Monton: mt, Lista_condiciones: lcon}
	return matchado
}
