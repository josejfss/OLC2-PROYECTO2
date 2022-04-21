package casteos

import (
	entorno "OLC2-PROYECTO2/COMPILADOR/ENTORNO"
	generador "OLC2-PROYECTO2/COMPILADOR/GENERADOR"
	interfaces "OLC2-PROYECTO2/COMPILADOR/INTERFACES"
	simbolos "OLC2-PROYECTO2/COMPILADOR/SIMBOLOS"
	"fmt"
	"strconv"
)

type CasteFloat struct {
	Valor  interfaces.Expresion
	Cambio simbolos.TipoExpresion
}

func Ncastefloat(v interfaces.Expresion, c simbolos.TipoExpresion) CasteFloat {
	casfloat := CasteFloat{Valor: v, Cambio: c}
	return casfloat
}

func (float CasteFloat) Ejecutar_Expresion(ent *entorno.Entorno) simbolos.Valor {
	valores := float.Valor.Ejecutar_Expresion(ent)
	if float.Cambio == simbolos.INTEGER {
		//SE HACE CONVERSIONES
		val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", valores.Valor), 64)
		valint := int(val1)
		//SE RETORNA SIMBOLO
		return simbolos.Valor{Tipo: simbolos.INTEGER, Valor: valint}
	} else if float.Cambio == simbolos.FLOAT {
		//SE RETORNA SIMBOLO
		return simbolos.Valor{Tipo: simbolos.FLOAT, Valor: valores.Valor}
	} else if float.Cambio == simbolos.CHAR {
		//SE HACE CONVERSIONES
		s := string(rune(valores.Valor.(float64)))
		//SE RETORNA SIMBOLO
		return simbolos.Valor{Tipo: simbolos.CHAR, Valor: s}
	} else if float.Cambio == simbolos.TEXTO {
		//SE HACE CONVERSIONES
		val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", valores.Valor), 64)
		s := fmt.Sprintf("%f", val1)
		//SE RETORNA SIMBOLO
		return simbolos.Valor{Tipo: simbolos.TEXTO, Valor: s}
	} else if float.Cambio == simbolos.YTEXTO {
		//SE HACE CONVERSIONES
		val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", valores.Valor), 64)
		s := fmt.Sprintf("%f", val1)
		//SE RETORNA SIMBOLO
		return simbolos.Valor{Tipo: simbolos.TEXTO, Valor: s}
	} else if float.Cambio == simbolos.BOOLEAN {
		//SE HACE CONVERSIONES
		newbool := valores.Valor.(float64) != 0
		//SE RETORNA SIMBOLO
		return simbolos.Valor{Tipo: simbolos.BOOLEAN, Valor: newbool}
	}
	return simbolos.Valor{Tipo: simbolos.NULL, Valor: -1}
}

func (float CasteFloat) Compilar_Expresion(ent *entorno.Entorno, gen *generador.Generador_C3D) simbolos.ValoresC3D {
	return simbolos.ValoresC3D{Valor: "0", EsTemporal: false, Tipo: simbolos.INTEGER, Label_verdadera: "", Label_false: ""}
}
