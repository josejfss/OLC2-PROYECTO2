package funnatvect

import (
	entorno "OLC2-PROYECTO2/COMPILADOR/ENTORNO"
	generador "OLC2-PROYECTO2/COMPILADOR/GENERADOR"
	interfaces "OLC2-PROYECTO2/COMPILADOR/INTERFACES"
	reportes "OLC2-PROYECTO2/COMPILADOR/REPORTES"
	simbolos "OLC2-PROYECTO2/COMPILADOR/SIMBOLOS"
	"strconv"
	"time"

	"github.com/colegno/arraylist"
)

type NewVect struct {
	Mutable       bool
	Identificador string
	Tipo_Vect     interfaces.Expresion
	Linea         int
	Columna       int
}

func Nnewvect(mut bool, nom string, tip interfaces.Expresion, lin int, col int) NewVect {
	nuevo_vector := NewVect{Mutable: mut,
		Identificador: nom,
		Tipo_Vect:     tip,
		Linea:         lin,
		Columna:       col}
	return nuevo_vector
}

func (nvect NewVect) Ejecutar_Instruccion(ent *entorno.Entorno, ent2 *entorno.Entorno) interface{} {
	tipovect := nvect.Tipo_Vect.Ejecutar_Expresion(ent)
	if !ent.Existe_Variable(nvect.Identificador) {
		if !ent.Existe_ArreVect(nvect.Identificador) {
			simbarre := simbolos.Simbolo_ArreVect{Mutable: nvect.Mutable,
				EsArreVect:       0,
				TipoVect:         tipovect.Tipo,
				Nombre:           nvect.Identificador,
				Dimensiones:      simbolos.ValArreglo.Len(),
				DimensionesLista: arraylist.New(),
				ValorArreVect:    make([]interface{}, 0),
				ValorLista:       arraylist.New(),
				Lintdim:          arraylist.New(),
				Lintcap:          arraylist.New(),
				PosicionTabla:    ent.Posicion,
				Linea:            nvect.Linea,
				Columna:          nvect.Columna}
			if ent.Nombre_Entorno != ent2.Nombre_Entorno {
				ent2.Posicion = ent2.Posicion - 1
				ent.Guardar_ArregloVector(nvect.Identificador, simbarre)
				ent2.Guardar_ArregloVector(nvect.Identificador, simbarre)
			} else {
				ent.Guardar_ArregloVector(nvect.Identificador, simbarre)
				ent2.Guardar_ArregloVector(nvect.Identificador, simbarre)
			}
			reportes.ReporteSimbolos(nvect.Identificador, "arreglo--"+reportes.ReportObteniendoSimbolos(tipovect.Tipo), ent.Nombre_Entorno, "--", strconv.Itoa(nvect.Linea), strconv.Itoa(nvect.Columna))
		} else {
			t := time.Now()
			reportes.Nerror("No se puede declarar vector.", ent.Nombre_Entorno, strconv.Itoa(nvect.Linea), strconv.Itoa(nvect.Columna), t.Format("2006-01-02 15:04:05"))
		}
	} else {
		t := time.Now()
		reportes.Nerror("Variables con el mismo nombre en el entorno.", ent.Nombre_Entorno, strconv.Itoa(nvect.Linea), strconv.Itoa(nvect.Columna), t.Format("2006-01-02 15:04:05"))
	}
	simbolos.ValArreglo.Clear()
	return 0
}

func (nvect NewVect) Compilar_Instruccion(ent *entorno.Entorno, gen *generador.Generador_C3D) interface{} {
	if ent.Existe_ArreVect(nvect.Identificador) {
		simbvect := ent.Obtener_ArreVect(nvect.Identificador)
		gen.Agregar_Comentario("INICIO DECLARACION VECTOR -- " + nvect.Identificador)
		tempo1 := gen.Crear_temporal()
		posi := simbvect.PosicionTabla
		gen.Agregar_Logica(tempo1 + " = SP + " + strconv.Itoa(posi) + ";\t\t// POSICION: " + nvect.Identificador)
		gen.Agregar_Stack(tempo1, "HP")
		gen.Agregar_Logica("HP = HP + 1;")
		gen.Agregar_Logica("HEAP[(int)HP] = -2;\nHP = HP + 1;")
		gen.Agregar_Comentario("FIN DECLARACION VECTOR -- " + nvect.Identificador)
		gen.LiberarTodosTemporales()
	}
	return 0
}
