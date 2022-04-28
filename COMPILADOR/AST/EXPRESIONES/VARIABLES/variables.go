package variables

import (
	entorno "OLC2-PROYECTO2/COMPILADOR/ENTORNO"
	generador "OLC2-PROYECTO2/COMPILADOR/GENERADOR"
	reportes "OLC2-PROYECTO2/COMPILADOR/REPORTES"
	simbolos "OLC2-PROYECTO2/COMPILADOR/SIMBOLOS"
	"strconv"
	"time"
)

type Variable struct {
	Identificador string
	Linea         int
	Columna       int
}

func NVariable(i string, l int, c int) Variable {
	variables := Variable{Identificador: i, Linea: l, Columna: c}
	return variables
}

func (vari Variable) Ejecutar_Expresion(ent *entorno.Entorno) simbolos.Valor {

	if vari.Identificador == "_" {
		return simbolos.Valor{Tipo: simbolos.NULL, Valor: "_"}
	} else if ent.Existe_Variable(vari.Identificador) {
		simb := ent.Obtener_Variable(vari.Identificador)
		return simbolos.Valor{Tipo: simb.TipoVariable, Valor: simb.ValorVariable}
	} else if ent.Existe_ArreVect(vari.Identificador) {
		simb := ent.Obtener_ArreVect(vari.Identificador)
		return simbolos.Valor{Tipo: simb.TipoVect, Valor: simb}
	} else {
		t := time.Now()
		reportes.Nerror("No existe la variable: "+vari.Identificador, ent.Nombre_Entorno, strconv.Itoa(vari.Linea), strconv.Itoa(vari.Columna), t.Format("2006-01-02 15:04:05"))
	}
	return simbolos.Valor{Tipo: simbolos.NULL, Valor: -1}
}

func (vari Variable) Compilar_Expresion(ent *entorno.Entorno, gen *generador.Generador_C3D) simbolos.ValoresC3D {
	if vari.Identificador == "_" {
		return simbolos.ValoresC3D{Valor: "_", EsTemporal: false, Tipo: simbolos.NULL, Label_verdadera: "", Label_false: ""}
	} else if ent.Existe_Variable(vari.Identificador) {
		obtvar := ent.Obtener_Variable(vari.Identificador)
		temp1 := gen.Crear_temporal()
		temp2 := gen.Crear_temporal()
		gen.Agregar_Logica(temp1 + " = SP + " + strconv.Itoa(obtvar.PosicionTabla) + ";")
		gen.Agregar_Logica(temp2 + " = STACK[int(" + temp1 + ")];")
		return simbolos.ValoresC3D{Valor: temp2, EsTemporal: true, Tipo: obtvar.TipoVariable, Label_verdadera: "", Label_false: ""}
	} else if ent.Existe_ArreVect(vari.Identificador) {
		obtvar := ent.Obtener_ArreVect(vari.Identificador)
		temp1 := gen.Crear_temporal()
		temp2 := gen.Crear_temporal()
		gen.Agregar_Logica(temp1 + " = SP + " + strconv.Itoa(obtvar.PosicionTabla) + ";")
		gen.Agregar_Logica(temp2 + " = STACK[int(" + temp1 + ")];")
		return simbolos.ValoresC3D{Valor: temp2, EsTemporal: true, Tipo: obtvar.TipoVect, Label_verdadera: "", Label_false: ""}
	}
	return simbolos.ValoresC3D{Valor: "0", EsTemporal: false, Tipo: simbolos.INTEGER, Label_verdadera: "", Label_false: ""}

}
