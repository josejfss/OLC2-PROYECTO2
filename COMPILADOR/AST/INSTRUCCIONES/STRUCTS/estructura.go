package structs

import (
	primitivo "OLC2-PROYECTO2/COMPILADOR/AST/EXPRESIONES/PRIMITIVO"
	declaracionvar "OLC2-PROYECTO2/COMPILADOR/AST/INSTRUCCIONES/DECLARACIONVAR"
	entorno "OLC2-PROYECTO2/COMPILADOR/ENTORNO"
	generador "OLC2-PROYECTO2/COMPILADOR/GENERADOR"
	reportes "OLC2-PROYECTO2/COMPILADOR/REPORTES"
	simbolos "OLC2-PROYECTO2/COMPILADOR/SIMBOLOS"
	"strconv"
	"time"

	"github.com/colegno/arraylist"
)

type DefStruct struct {
	Nombre string
	Tipo   simbolos.TipoExpresion
}

func Ndefstruct(nom string, tip simbolos.TipoExpresion) DefStruct {
	defstructs := DefStruct{Nombre: nom, Tipo: tip}
	return defstructs
}

type StructRust struct {
	Identificador string
	Definiciones  *arraylist.List
	Linea         int
	Columna       int
}

func Nstructrust(id string, def *arraylist.List, lin int, col int) StructRust {
	structrust := StructRust{Identificador: id, Definiciones: def, Linea: lin, Columna: col}
	return structrust
}

func (stru StructRust) Ejecutar_Instruccion(ent *entorno.Entorno, ent2 *entorno.Entorno) interface{} {
	if !ent.Existe_Structs(stru.Identificador) {
		reportes.ReporteSimbolos(stru.Identificador, "struct--"+reportes.ReportObteniendoSimbolos(simbolos.NULL), ent.Nombre_Entorno, "--", strconv.Itoa(stru.Linea), strconv.Itoa(stru.Columna))

		entorno1 := entorno.Nuevo_Entorno("struct_"+stru.Identificador, ent)
		entorno2 := entorno.Nuevo_Entorno("struct_"+stru.Identificador, ent2)
		for i := 0; i < stru.Definiciones.Len(); i++ {
			act := stru.Definiciones.GetValue(i).(DefStruct)
			var declapar declaracionvar.Declaracion
			declapar.Identificador = act.Nombre
			declapar.TipoDecla = act.Tipo
			vals := primitivo.Nuevo_Dato_Primitivo(ValorInicial(act.Tipo), act.Tipo)
			declapar.Valor_exp = vals
			declapar.Linea = stru.Linea
			declapar.Columna = stru.Columna
			declapar.Ejecutar_Instruccion(entorno1, entorno2)
		}

		simbstruct := simbolos.Simbolo_Struct{Identificador: stru.Identificador, L_Atributos: stru.Definiciones.Clone(), Tamaño: entorno1.Tamaño, Linea: stru.Linea, Columna: stru.Columna}
		ent.Guardar_Struct(stru.Identificador, simbstruct)
		ent2.Guardar_Struct(stru.Identificador, simbstruct)

		ent.Guardar_Entorno(entorno1)
		ent2.Guardar_Entorno(entorno2)
	} else {
		t := time.Now()
		reportes.Nerror("Ya existe un Struct con ese nombre.", ent.Nombre_Entorno, strconv.Itoa(stru.Linea), strconv.Itoa(stru.Columna), t.Format("2006-01-02 15:04:05"))
	}
	return 0
}

func (stru StructRust) Compilar_Instruccion(ent *entorno.Entorno, gen *generador.Generador_C3D) interface{} {
	return 0
}

func ValorInicial(sim simbolos.TipoExpresion) interface{} {
	if sim == simbolos.INTEGER {
		return 0
	} else if sim == simbolos.BOOLEAN {
		return false
	} else if sim == simbolos.FLOAT {
		return 0.0
	} else if sim == simbolos.YTEXTO {
		return ""
	} else if sim == simbolos.TEXTO {
		return ""
	}

	return 0
}
