package logica

import (
	entorno "OLC2-PROYECTO2/COMPILADOR/ENTORNO"
	generador "OLC2-PROYECTO2/COMPILADOR/GENERADOR"
	interfaces "OLC2-PROYECTO2/COMPILADOR/INTERFACES"
	reportes "OLC2-PROYECTO2/COMPILADOR/REPORTES"
	simbolos "OLC2-PROYECTO2/COMPILADOR/SIMBOLOS"
	"fmt"
	"strconv"
	"time"
)

type OpNot struct {
	Operacion interfaces.Expresion
	Linea     int
	Columna   int
}

func Nopnot(op interfaces.Expresion, l int, c int) OpNot {
	nots := OpNot{Operacion: op, Linea: l, Columna: c}
	return nots
}

func (not OpNot) Ejecutar_Expresion(ent *entorno.Entorno) simbolos.Valor {
	//SE OBTIENE EL VALOR DEL OPERADOR UNARIO
	var unarias simbolos.Valor = not.Operacion.Ejecutar_Expresion(ent)
	//SI EL TIPO ES ENTERO
	if unarias.Tipo == simbolos.BOOLEAN {
		//SE RETORNA SIMBOLO
		return simbolos.Valor{Tipo: simbolos.BOOLEAN, Valor: !unarias.Valor.(bool)}
		//SI EL TIPO ES DECIMAL
	} else {
		//ERROR
		t := time.Now()
		reportes.Nerror("No es posible realizar la operacion: Not", ent.Nombre_Entorno, strconv.Itoa(not.Linea), strconv.Itoa(not.Columna), t.Format("2006-01-02 15:04:05"))
	}
	//SINO SE CUMPLEN LAS CONDICIONES SE RETORNA UN SIMBOLO DE TIPO NULL CON VALOR NULL
	return simbolos.Valor{Tipo: simbolos.NULL, Valor: -1}
}

func (not OpNot) Compilar_Expresion(ent *entorno.Entorno, gen *generador.Generador_C3D) simbolos.ValoresC3D {
	var opera_unario simbolos.ValoresC3D = not.Operacion.Compilar_Expresion(ent, gen)
	if opera_unario.Tipo == simbolos.BOOLEAN {
		return simbolos.ValoresC3D{Valor: opera_unario.Valor, EsTemporal: opera_unario.EsTemporal, Tipo: opera_unario.Tipo, Label_verdadera: opera_unario.Label_false, Label_false: opera_unario.Label_verdadera}
	} else {
		fmt.Println("ERROR, tipos")
		gen.AgregarError("ERROR-TIPOS--NOT", strconv.Itoa(not.Linea), strconv.Itoa(not.Columna))
	}
	return simbolos.ValoresC3D{Valor: "0", EsTemporal: false, Tipo: simbolos.INTEGER, Label_verdadera: "", Label_false: ""}
}
