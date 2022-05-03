package modulo

import (
	entorno "OLC2-PROYECTO2/COMPILADOR/ENTORNO"
	generador "OLC2-PROYECTO2/COMPILADOR/GENERADOR"
	reportes "OLC2-PROYECTO2/COMPILADOR/REPORTES"
	simbolos "OLC2-PROYECTO2/COMPILADOR/SIMBOLOS"
	"strconv"

	"github.com/colegno/arraylist"
)

type Modulos struct {
	Publico    bool
	Nombre     string
	ListaInstr *arraylist.List
	Linea      int
	Columna    int
}

func Nmods(pub bool, nom string, list *arraylist.List, lin int, col int) Modulos {
	mods := Modulos{Publico: pub, Nombre: nom, ListaInstr: list, Linea: lin, Columna: col}
	return mods
}

func (mods Modulos) Ejecutar_Instruccion(ent *entorno.Entorno, ent2 *entorno.Entorno) interface{} {
	if !ent.Existe_Modulos(mods.Nombre) {
		simbmodulo := simbolos.Simbolos_Modulos{Publico: mods.Publico,
			Identificador:   mods.Nombre,
			L_Declaraciones: mods.ListaInstr.Clone(),
			Linea:           mods.Linea,
			Columna:         mods.Columna}
		ent.Guardar_Modulos(mods.Nombre, simbmodulo)
		reportes.ReporteBasesDatos(mods.Nombre, strconv.Itoa(mods.Linea))
	}
	return 0
}

func (mods Modulos) Compilar_Instruccion(ent *entorno.Entorno, gen *generador.Generador_C3D) interface{} {

	return 0
}
