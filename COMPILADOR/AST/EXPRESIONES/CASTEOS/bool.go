package casteos

import (
	entorno "OLC2-PROYECTO2/COMPILADOR/ENTORNO"
	generador "OLC2-PROYECTO2/COMPILADOR/GENERADOR"
	interfaces "OLC2-PROYECTO2/COMPILADOR/INTERFACES"
	simbolos "OLC2-PROYECTO2/COMPILADOR/SIMBOLOS"
)

type CasteBool struct {
	Valor  interfaces.Expresion
	Cambio simbolos.TipoExpresion
}

func Ncastebool(v interfaces.Expresion, c simbolos.TipoExpresion) CasteBool {
	casfloat := CasteBool{Valor: v, Cambio: c}
	return casfloat
}

func (bools CasteBool) Ejecutar_Expresion(ent *entorno.Entorno) simbolos.Valor {
	valores := bools.Valor.Ejecutar_Expresion(ent)
	if bools.Cambio == simbolos.INTEGER {
		//SE VERIFICA SI ES TRUE O FALSE
		if valores.Valor.(bool) {
			//SE RETORNA SIMBOLO
			return simbolos.Valor{Tipo: simbolos.INTEGER, Valor: 1}
		} else {
			//SE RETORNA SIMBOLO
			return simbolos.Valor{Tipo: simbolos.INTEGER, Valor: 0}
		}
	} else if bools.Cambio == simbolos.FLOAT {
		//SE VERIFICA SI ES TRUE O FALSE
		if valores.Valor.(bool) {
			//SE RETORNA SIMBOLO
			return simbolos.Valor{Tipo: simbolos.FLOAT, Valor: 1.0}
		} else {
			//SE RETORNA SIMBOLO
			return simbolos.Valor{Tipo: simbolos.FLOAT, Valor: 0.0}
		}
	} else if bools.Cambio == simbolos.CHAR {
		//SE VERIFICA SI ES TRUE O FALSE
		if valores.Valor.(bool) {
			//SE RETORNA SIMBOLO
			return simbolos.Valor{Tipo: simbolos.CHAR, Valor: 1.0}
		} else {
			//SE RETORNA SIMBOLO
			return simbolos.Valor{Tipo: simbolos.CHAR, Valor: 0.0}
		}
	} else if bools.Cambio == simbolos.TEXTO {
		//SE VERIFICA SI ES TRUE O FALSE
		if valores.Valor.(bool) {
			//SE RETORNA SIMBOLO
			return simbolos.Valor{Tipo: simbolos.TEXTO, Valor: "true"}
		} else {
			//SE RETORNA SIMBOLO
			return simbolos.Valor{Tipo: simbolos.TEXTO, Valor: "false"}
		}
	} else if bools.Cambio == simbolos.YTEXTO {
		//SE VERIFICA SI ES TRUE O FALSE
		if valores.Valor.(bool) {
			//SE RETORNA SIMBOLO
			return simbolos.Valor{Tipo: simbolos.YTEXTO, Valor: "true"}
		} else {
			//SE RETORNA SIMBOLO
			return simbolos.Valor{Tipo: simbolos.YTEXTO, Valor: "false"}
		}
	} else if bools.Cambio == simbolos.BOOLEAN {
		//SE RETORNA SIMBOLO
		return simbolos.Valor{Tipo: simbolos.BOOLEAN, Valor: valores.Valor}
	}
	return simbolos.Valor{Tipo: simbolos.NULL, Valor: -1}
}

func (bools CasteBool) Compilar_Expresion(ent *entorno.Entorno, gen *generador.Generador_C3D) simbolos.ValoresC3D {
	return simbolos.ValoresC3D{Valor: "0", EsTemporal: false, Tipo: simbolos.INTEGER, Label_verdadera: "", Label_false: ""}
}
