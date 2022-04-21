package casteos

import (
	entorno "OLC2-PROYECTO2/COMPILADOR/ENTORNO"
	generador "OLC2-PROYECTO2/COMPILADOR/GENERADOR"
	interfaces "OLC2-PROYECTO2/COMPILADOR/INTERFACES"
	simbolos "OLC2-PROYECTO2/COMPILADOR/SIMBOLOS"
	"fmt"
	"strconv"
)

type CasteEntero struct {
	Valor  interfaces.Expresion
	Cambio simbolos.TipoExpresion
}

func Ncasteentero(v interfaces.Expresion, c simbolos.TipoExpresion) CasteEntero {
	casentero := CasteEntero{Valor: v, Cambio: c}
	return casentero
}

func (ente CasteEntero) Ejecutar_Expresion(ent *entorno.Entorno) simbolos.Valor {
	valores := ente.Valor.Ejecutar_Expresion(ent)
	if ente.Cambio == simbolos.INTEGER {
		return simbolos.Valor{Tipo: simbolos.INTEGER, Valor: valores.Valor}
	} else if ente.Cambio == simbolos.FLOAT {
		//SE HACE CONVERSIONES
		val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", valores.Valor), 64)
		//SE RETORNA SIMBOLO
		return simbolos.Valor{Tipo: simbolos.FLOAT, Valor: val1}
	} else if ente.Cambio == simbolos.CHAR {
		//SE HACE CONVERSIONES
		s := string(rune(valores.Valor.(int)))
		//SE RETORNA SIMBOLO
		return simbolos.Valor{Tipo: simbolos.CHAR, Valor: s}
	} else if ente.Cambio == simbolos.TEXTO {
		//SE HACE CONVERSIONES
		val1, err := strconv.Atoi(fmt.Sprintf("%d", valores.Valor))
		if err != nil {
			fmt.Println(err)
		}
		//SE RETORNA SIMBOLO
		return simbolos.Valor{Tipo: simbolos.TEXTO, Valor: val1}
	} else if ente.Cambio == simbolos.YTEXTO {
		//SE HACE CONVERSIONES
		val1, err := strconv.Atoi(fmt.Sprintf("%d", valores.Valor))
		if err != nil {
			fmt.Println(err)
		}
		//SE RETORNA SIMBOLO
		return simbolos.Valor{Tipo: simbolos.YTEXTO, Valor: val1}
	} else if ente.Cambio == simbolos.BOOLEAN {
		//SE HACE CONVERSIONES
		newbool := valores.Valor.(int) != 0
		//SE RETORNA SIMBOLO
		return simbolos.Valor{Tipo: simbolos.BOOLEAN, Valor: newbool}
	}
	return simbolos.Valor{Tipo: simbolos.NULL, Valor: -1}
}

func (ente CasteEntero) Compilar_Expresion(ent *entorno.Entorno, gen *generador.Generador_C3D) simbolos.ValoresC3D {
	return simbolos.ValoresC3D{Valor: "0", EsTemporal: false, Tipo: simbolos.INTEGER, Label_verdadera: "", Label_false: ""}
}
