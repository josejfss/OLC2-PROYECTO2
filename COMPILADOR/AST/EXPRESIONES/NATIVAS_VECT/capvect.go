package nativasvect

import (
	entorno "OLC2-PROYECTO2/COMPILADOR/ENTORNO"
	generador "OLC2-PROYECTO2/COMPILADOR/GENERADOR"
	reportes "OLC2-PROYECTO2/COMPILADOR/REPORTES"
	simbolos "OLC2-PROYECTO2/COMPILADOR/SIMBOLOS"
	"strconv"
	"time"
)

type CapVect struct {
	Identificador string
	Linea         int
	Columna       int
}

func Ncapvect(nom string, lin int, col int) CapVect {
	capvector := CapVect{Identificador: nom, Linea: lin, Columna: col}
	return capvector
}

func (capacidad CapVect) Ejecutar_Expresion(ent *entorno.Entorno) simbolos.Valor {
	if ent.Existe_ArreVect(capacidad.Identificador) {
		simbvect := ent.Obtener_ArreVect(capacidad.Identificador)
		return simbolos.Valor{Tipo: simbolos.INTEGER, Valor: simbvect.EsArreVect}
	} else {
		t := time.Now()
		reportes.Nerror("No existe el vector", ent.Nombre_Entorno, strconv.Itoa(capacidad.Linea), strconv.Itoa(capacidad.Columna), t.Format("2006-01-02 15:04:05"))
	}
	//NO SE ACEPTO LOS TIPOS
	return simbolos.Valor{Tipo: simbolos.NULL, Valor: -1}
}

func (capacidad CapVect) Compilar_Expresion(ent *entorno.Entorno, gen *generador.Generador_C3D) simbolos.ValoresC3D {
	gen.Agregar_Comentario("INICIANDO FUNCION NATIVA CAPACIDAD VECTOR")
	if ent.Existe_ArreVect(capacidad.Identificador) {
		simbvect := ent.Obtener_ArreVect(capacidad.Identificador)
		nuevo_temporal := gen.Crear_temporal()
		gen.Agregar_Logica(nuevo_temporal + " = " + strconv.Itoa(simbvect.EsArreVect) + ";")
		return simbolos.ValoresC3D{Valor: nuevo_temporal, EsTemporal: true, Tipo: simbolos.INTEGER, Label_verdadera: "", Label_false: ""}
	}
	gen.Agregar_Comentario("FINALIZANDO FUNCION NATIVA CAPACIDAD VECTOR")
	return simbolos.ValoresC3D{Valor: "0", EsTemporal: false, Tipo: simbolos.INTEGER, Label_verdadera: "", Label_false: ""}
}
