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

type WCapVect struct {
	Mutable       bool
	Identificador string
	Tipo_Vect     interfaces.Expresion
	Capacidad     interfaces.Expresion
	Linea         int
	Columna       int
}

func Nwcvect(mut bool, nom string, tip interfaces.Expresion, cap interfaces.Expresion, lin int, col int) WCapVect {
	nuevo_vector := WCapVect{Mutable: mut,
		Identificador: nom,
		Tipo_Vect:     tip,
		Capacidad:     cap,
		Linea:         lin,
		Columna:       col}
	return nuevo_vector
}

func (wcvect WCapVect) Ejecutar_Instruccion(ent *entorno.Entorno, ent2 *entorno.Entorno) interface{} {
	tipovect := wcvect.Tipo_Vect.Ejecutar_Expresion(ent)
	cap := wcvect.Capacidad.Ejecutar_Expresion(ent)
	if !ent.Existe_Variable(wcvect.Identificador) {
		if !ent.Existe_ArreVect(wcvect.Identificador) {
			if cap.Tipo == simbolos.INTEGER {
				simbarre := simbolos.Simbolo_ArreVect{Mutable: wcvect.Mutable,
					EsArreVect:       cap.Valor.(int),
					TipoVect:         tipovect.Tipo,
					Nombre:           wcvect.Identificador,
					Dimensiones:      simbolos.ValArreglo.Len(),
					DimensionesLista: arraylist.New(),
					ValorArreVect:    make([]interface{}, 0),
					ValorLista:       arraylist.New(),
					Lintdim:          arraylist.New(),
					Lintcap:          arraylist.New(),
					PosicionTabla:    ent.Posicion,
					Linea:            wcvect.Linea,
					Columna:          wcvect.Columna}
				if ent.Nombre_Entorno != ent2.Nombre_Entorno {
					ent2.Posicion = ent2.Posicion - 1
					ent.Guardar_ArregloVector(wcvect.Identificador, simbarre)
					ent2.Guardar_ArregloVector(wcvect.Identificador, simbarre)
				} else {
					ent.Guardar_ArregloVector(wcvect.Identificador, simbarre)
					ent2.Guardar_ArregloVector(wcvect.Identificador, simbarre)
				}
				reportes.ReporteSimbolos(wcvect.Identificador, "arreglo--"+reportes.ReportObteniendoSimbolos(tipovect.Tipo), ent.Nombre_Entorno, "--", strconv.Itoa(wcvect.Linea), strconv.Itoa(wcvect.Columna))
			} else {
				t := time.Now()
				reportes.Nerror("La capacidad debe de ser de tipo entero.", ent.Nombre_Entorno, strconv.Itoa(wcvect.Linea), strconv.Itoa(wcvect.Columna), t.Format("2006-01-02 15:04:05"))
			}

		} else {
			t := time.Now()
			reportes.Nerror("No se puede declarar vector.", ent.Nombre_Entorno, strconv.Itoa(wcvect.Linea), strconv.Itoa(wcvect.Columna), t.Format("2006-01-02 15:04:05"))
		}
	} else {
		t := time.Now()
		reportes.Nerror("Variables con el mismo nombre en el entorno.", ent.Nombre_Entorno, strconv.Itoa(wcvect.Linea), strconv.Itoa(wcvect.Columna), t.Format("2006-01-02 15:04:05"))
	}
	simbolos.ValArreglo.Clear()
	return 0
}

func (wcvect WCapVect) Compilar_Instruccion(ent *entorno.Entorno, gen *generador.Generador_C3D) interface{} {
	if ent.Existe_ArreVect(wcvect.Identificador) {
		simbvect := ent.Obtener_ArreVect(wcvect.Identificador)
		gen.Agregar_Comentario("INICIO DECLARACION VECTOR -- " + wcvect.Identificador)
		tempo1 := gen.Crear_temporal()
		posi := simbvect.PosicionTabla
		gen.Agregar_Logica(tempo1 + " = SP + " + strconv.Itoa(posi) + ";\t\t// POSICION: " + wcvect.Identificador)
		gen.Agregar_Stack(tempo1, "HP")
		gen.Agregar_Logica("HP = HP +" + strconv.Itoa(simbvect.EsArreVect) + ";")
		gen.Agregar_Logica("HEAP[(int)HP] = -2;\nHP = HP + 1;")
		gen.Agregar_Comentario("FIN DECLARACION VECTOR -- " + wcvect.Identificador)
		gen.LiberarTodosTemporales()
	}
	return 0
}
