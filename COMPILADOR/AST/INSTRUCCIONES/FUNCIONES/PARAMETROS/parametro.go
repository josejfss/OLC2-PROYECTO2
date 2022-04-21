package parametros

import (
	interfaces "OLC2-PROYECTO2/COMPILADOR/INTERFACES"
	simbolos "OLC2-PROYECTO2/COMPILADOR/SIMBOLOS"
)

type Param struct {
	Identificador string
	Valref        int
	ParametroT    string
	TipoVar       simbolos.TipoExpresion
	Tipo          interfaces.Expresion
	TamArre       interfaces.Expresion
	Linea         int
	Columna       int
}

func Nuevo_parametro(id string, vr int, parat string, tipva simbolos.TipoExpresion, tip interfaces.Expresion, tarr interfaces.Expresion, l int, c int) Param {
	arr := Param{Identificador: id, Valref: vr, ParametroT: parat, TipoVar: tipva, Tipo: tip, TamArre: tarr, Linea: l, Columna: c}
	return arr
}
