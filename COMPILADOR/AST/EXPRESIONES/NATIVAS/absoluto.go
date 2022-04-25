package nativas

import (
	entorno "OLC2-PROYECTO2/COMPILADOR/ENTORNO"
	generador "OLC2-PROYECTO2/COMPILADOR/GENERADOR"
	interfaces "OLC2-PROYECTO2/COMPILADOR/INTERFACES"
	reportes "OLC2-PROYECTO2/COMPILADOR/REPORTES"
	simbolos "OLC2-PROYECTO2/COMPILADOR/SIMBOLOS"
	"fmt"
	"math"
	"strconv"
	"time"
)

type Absoluto struct {
	Val     interfaces.Expresion
	Linea   int
	Columna int
}

func Nabsoluto(v interfaces.Expresion, l int, c int) Absoluto {
	abs := Absoluto{Val: v, Linea: l, Columna: c}
	return abs
}

func (abs Absoluto) Ejecutar_Expresion(ent *entorno.Entorno) simbolos.Valor {
	resultado := abs.Val.Ejecutar_Expresion(ent)
	if resultado.Tipo == simbolos.INTEGER {
		val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", resultado.Valor), 64)
		abs := math.Abs(val1)
		absi := int(abs)
		return simbolos.Valor{Tipo: resultado.Tipo, Valor: absi}
	} else if resultado.Tipo == simbolos.FLOAT {
		val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", resultado.Valor), 64)
		abs := math.Abs(val1)
		absi := int(abs)
		return simbolos.Valor{Tipo: resultado.Tipo, Valor: absi}
	} else {
		t := time.Now()
		reportes.Nerror("No es posible realizar la funcion: Abs", ent.Nombre_Entorno, strconv.Itoa(abs.Linea), strconv.Itoa(abs.Columna), t.Format("2006-01-02 15:04:05"))
	}
	return simbolos.Valor{Tipo: simbolos.NULL, Valor: -1}
}

func (abs Absoluto) Compilar_Expresion(ent *entorno.Entorno, gen *generador.Generador_C3D) simbolos.ValoresC3D {
	resultado := abs.Val.Compilar_Expresion(ent, gen)
	nuevo_temporal := gen.Crear_temporal()
	if resultado.Tipo == simbolos.INTEGER {
		gen.Agregar_Logica(nuevo_temporal + " = fabs(" + resultado.Valor + "); \t\t//FUN. NATIVA ABSOLUTO")
		return simbolos.ValoresC3D{Valor: nuevo_temporal, EsTemporal: true, Tipo: simbolos.FLOAT, Label_verdadera: "", Label_false: ""}
	} else if resultado.Tipo == simbolos.FLOAT {
		gen.Agregar_Logica(nuevo_temporal + " = fabs(" + resultado.Valor + "); \t\t//FUN. NATIVA ABSOLUTO")
		return simbolos.ValoresC3D{Valor: nuevo_temporal, EsTemporal: true, Tipo: simbolos.FLOAT, Label_verdadera: "", Label_false: ""}
	} else {
		gen.AgregarError("ERROR-TIPOS--ABSOLUTO", strconv.Itoa(abs.Linea), strconv.Itoa(abs.Columna))
	}
	return simbolos.ValoresC3D{Valor: "0", EsTemporal: false, Tipo: simbolos.INTEGER, Label_verdadera: "", Label_false: ""}
}
