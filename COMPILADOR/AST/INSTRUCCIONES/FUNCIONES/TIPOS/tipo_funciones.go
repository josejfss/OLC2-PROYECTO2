package tipos

import (
	interfaces "OLC2-PROYECTO2/COMPILADOR/INTERFACES"
	simbolos "OLC2-PROYECTO2/COMPILADOR/SIMBOLOS"
)

type TiposFunca struct {
	Tips     string
	TiVa     simbolos.TipoExpresion
	TipoReal interfaces.Expresion
	ValArre  interfaces.Expresion
}

func Ntiposfunca(tips string, tvar simbolos.TipoExpresion, tiporeal interfaces.Expresion, valarre interfaces.Expresion) TiposFunca {
	tfunca := TiposFunca{Tips: tips, TiVa: tvar, TipoReal: tiporeal, ValArre: valarre}
	return tfunca
}
