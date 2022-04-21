package casteos

import (
	entorno "OLC2-PROYECTO2/COMPILADOR/ENTORNO"
	generador "OLC2-PROYECTO2/COMPILADOR/GENERADOR"
	interfaces "OLC2-PROYECTO2/COMPILADOR/INTERFACES"
	simbolos "OLC2-PROYECTO2/COMPILADOR/SIMBOLOS"
)

type CasteChar struct {
	Valor  interfaces.Expresion
	Cambio simbolos.TipoExpresion
}

func Ncastechar(v interfaces.Expresion, c simbolos.TipoExpresion) CasteChar {
	caschar := CasteChar{Valor: v, Cambio: c}
	return caschar
}

func (char CasteChar) Ejecutar_Expresion(ent *entorno.Entorno) simbolos.Valor {
	valores := char.Valor.Ejecutar_Expresion(ent)
	if char.Cambio == simbolos.INTEGER {
		//SE HACE CONVERSIONES
		f := valores.Valor.(int32)
		var h int = int(f)
		//SE RETORNA SIMBOLO
		return simbolos.Valor{Tipo: simbolos.INTEGER, Valor: h}
	} else if char.Cambio == simbolos.FLOAT {
		//SE HACE CONVERSIONES
		f := valores.Valor.(int32)
		var h int = int(f)
		j := float64(h)
		//SE RETORNA SIMBOLO
		return simbolos.Valor{Tipo: simbolos.FLOAT, Valor: j}
	} else if char.Cambio == simbolos.CHAR {
		//SE RETORNA SIMBOLO
		return simbolos.Valor{Tipo: simbolos.CHAR, Valor: valores.Valor}
	} else if char.Cambio == simbolos.TEXTO {
		//SE HACE CONVERSIONES
		s := string(rune(valores.Valor.(int32)))
		//SE RETORNA SIMBOLO
		return simbolos.Valor{Tipo: simbolos.TEXTO, Valor: s}
	} else if char.Cambio == simbolos.YTEXTO {
		//SE HACE CONVERSIONES
		s := string(rune(valores.Valor.(int32)))
		//SE RETORNA SIMBOLO
		return simbolos.Valor{Tipo: simbolos.YTEXTO, Valor: s}
	} else if char.Cambio == simbolos.BOOLEAN {
		//SE HACE CONVERSIONES
		newbool := valores.Valor.(int32) != 0
		//SE RETORNA SIMBOLO
		return simbolos.Valor{Tipo: simbolos.BOOLEAN, Valor: newbool}
	}
	return simbolos.Valor{Tipo: simbolos.NULL, Valor: -1}
}

func (float CasteChar) Compilar_Expresion(ent *entorno.Entorno, gen *generador.Generador_C3D) simbolos.ValoresC3D {
	return simbolos.ValoresC3D{Valor: "0", EsTemporal: false, Tipo: simbolos.INTEGER, Label_verdadera: "", Label_false: ""}
}
