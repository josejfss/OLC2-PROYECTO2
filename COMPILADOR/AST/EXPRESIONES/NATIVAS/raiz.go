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

type Raiz struct {
	Val     interfaces.Expresion
	Linea   int
	Columna int
}

func Nraiz(v interfaces.Expresion, l int, c int) Raiz {
	sqrt := Raiz{Val: v, Linea: l, Columna: c}
	return sqrt
}

func (raiz Raiz) Ejecutar_Expresion(ent *entorno.Entorno) simbolos.Valor {
	resultado := raiz.Val.Ejecutar_Expresion(ent)
	if resultado.Tipo == simbolos.INTEGER {
		val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", resultado.Valor), 64)
		sqrt := math.Sqrt(val1)
		sqrt1 := int(sqrt)
		return simbolos.Valor{Tipo: resultado.Tipo, Valor: sqrt1}
	} else if resultado.Tipo == simbolos.FLOAT {
		val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", resultado.Valor), 64)
		sqrt := math.Sqrt(val1)
		sqrt1 := int(sqrt)
		return simbolos.Valor{Tipo: resultado.Tipo, Valor: sqrt1}
	} else {
		t := time.Now()
		reportes.Nerror("No es posible realizar la funcion: Sqrt", ent.Nombre_Entorno, strconv.Itoa(raiz.Linea), strconv.Itoa(raiz.Columna), t.Format("2006-01-02 15:04:05"))
	}
	return simbolos.Valor{Tipo: simbolos.NULL, Valor: -1}
}

func (raiz Raiz) Compilar_Expresion(ent *entorno.Entorno, gen *generador.Generador_C3D) simbolos.ValoresC3D {
	resultado := raiz.Val.Compilar_Expresion(ent, gen)
	nuevo_temporal := gen.Crear_temporal()
	if resultado.Tipo == simbolos.INTEGER {
		gen.Agregar_Logica(nuevo_temporal + " = sqrt(" + resultado.Valor + "); \t\t//FUN. NATIVA RAIZ")
		return simbolos.ValoresC3D{Valor: nuevo_temporal, EsTemporal: true, Tipo: simbolos.FLOAT, Label_verdadera: "", Label_false: ""}
	} else if resultado.Tipo == simbolos.FLOAT {
		gen.Agregar_Logica(nuevo_temporal + " = sqrt(" + resultado.Valor + "); \t\t//FUN. NATIVA RAIZ")
		return simbolos.ValoresC3D{Valor: nuevo_temporal, EsTemporal: true, Tipo: simbolos.FLOAT, Label_verdadera: "", Label_false: ""}
	} else {
		gen.AgregarError("ERROR-TIPOS--RAIZ", strconv.Itoa(raiz.Linea), strconv.Itoa(raiz.Columna))
	}
	return simbolos.ValoresC3D{Valor: "0", EsTemporal: false, Tipo: simbolos.INTEGER, Label_verdadera: "", Label_false: ""}
}
