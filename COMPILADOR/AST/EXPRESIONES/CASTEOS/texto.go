package casteos

import (
	entorno "OLC2-PROYECTO2/COMPILADOR/ENTORNO"
	generador "OLC2-PROYECTO2/COMPILADOR/GENERADOR"
	interfaces "OLC2-PROYECTO2/COMPILADOR/INTERFACES"
	simbolos "OLC2-PROYECTO2/COMPILADOR/SIMBOLOS"
	"strconv"
)

type CasteTexto struct {
	Valor  interfaces.Expresion
	Cambio simbolos.TipoExpresion
}

func Ncastetexto(v interfaces.Expresion, c simbolos.TipoExpresion) CasteTexto {
	castxt := CasteTexto{Valor: v, Cambio: c}
	return castxt
}

func (txt CasteTexto) Ejecutar_Expresion(ent *entorno.Entorno) simbolos.Valor {
	valores := txt.Valor.Ejecutar_Expresion(ent)
	if txt.Cambio == simbolos.INTEGER {
		//SE HACE CONVERSIONES
		val1, _ := strconv.Atoi(valores.Valor.(string))
		//SE RETORNA SIMBOLO
		return simbolos.Valor{Tipo: simbolos.INTEGER, Valor: val1}
	} else if txt.Cambio == simbolos.FLOAT {
		//SE HACE CONVERSIONES
		s, _ := strconv.ParseFloat(valores.Valor.(string), 64)
		//SE RETORNA SIMBOLO
		return simbolos.Valor{Tipo: simbolos.FLOAT, Valor: s}
	} else if txt.Cambio == simbolos.CHAR {
		//SE HACE CONVERSIONES
		r := []rune(valores.Valor.(string))
		//SE RETORNA SIMBOLO
		return simbolos.Valor{Tipo: simbolos.CHAR, Valor: r[0]}
	} else if txt.Cambio == simbolos.TEXTO {
		//SE RETORNA SIMBOLO
		return simbolos.Valor{Tipo: simbolos.TEXTO, Valor: valores.Valor}
	} else if txt.Cambio == simbolos.YTEXTO {
		//SE RETORNA SIMBOLO
		return simbolos.Valor{Tipo: simbolos.YTEXTO, Valor: valores.Valor}
	} else if txt.Cambio == simbolos.BOOLEAN {
		//SE HACE CONVERSIONES
		newbool, _ := strconv.ParseBool(valores.Valor.(string))
		//SE RETORNA SIMBOLO
		return simbolos.Valor{Tipo: simbolos.BOOLEAN, Valor: newbool}
	}
	return simbolos.Valor{Tipo: simbolos.NULL, Valor: -1}
}

func (txt CasteTexto) Compilar_Expresion(ent *entorno.Entorno, gen *generador.Generador_C3D) simbolos.ValoresC3D {
	return simbolos.ValoresC3D{Valor: "0", EsTemporal: false, Tipo: simbolos.INTEGER, Label_verdadera: "", Label_false: ""}
}
