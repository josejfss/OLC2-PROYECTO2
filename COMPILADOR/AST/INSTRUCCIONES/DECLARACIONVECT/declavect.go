package declaracionvect

import (
	entorno "OLC2-PROYECTO2/COMPILADOR/ENTORNO"
	generador "OLC2-PROYECTO2/COMPILADOR/GENERADOR"
	interfaces "OLC2-PROYECTO2/COMPILADOR/INTERFACES"
	reportes "OLC2-PROYECTO2/COMPILADOR/REPORTES"
	simbolos "OLC2-PROYECTO2/COMPILADOR/SIMBOLOS"
	"fmt"
	"strconv"
	"time"

	"github.com/colegno/arraylist"
)

type DeclaVector struct {
	Mutable       bool
	Identificador string
	Dimension     *arraylist.List
	ValExp        interfaces.Expresion
	Linea         int
	Columna       int
}

func Ndeclavector(m bool, id string, di *arraylist.List, ve interfaces.Expresion, l int, c int) DeclaVector {
	declavect := DeclaVector{Mutable: m, Identificador: id, Dimension: di, ValExp: ve, Linea: l, Columna: c}
	return declavect
}

func (dvect DeclaVector) Ejecutar_Instruccion(ent *entorno.Entorno, ent2 *entorno.Entorno) interface{} {
	tipvect := dvect.Dimension.GetValue(0).(interfaces.Expresion)
	tipovect := tipvect.Ejecutar_Expresion(ent)
	resultado := dvect.ValExp.Ejecutar_Expresion(ent)
	if !ent.Existe_Variable(dvect.Identificador) {
		if !ent.Existe_ArreVect(dvect.Identificador) {
			if resultado.Tipo == simbolos.VECTOR {
				val_resultado := resultado.Valor.(simbolos.Valor)
				obj_vect := val_resultado.Valor.(simbolos.Simbolo_ArreVect)
				if dvect.Dimension.Len() == 1 {
					if obj_vect.TipoVect == tipovect.Tipo {
						simbarre := simbolos.Simbolo_ArreVect{Mutable: dvect.Mutable,
							EsArreVect:       0,
							TipoVect:         tipovect.Tipo,
							Nombre:           dvect.Identificador,
							Dimensiones:      dvect.Dimension.Len(),
							DimensionesLista: dvect.Dimension.Clone(),
							ValorArreVect:    obj_vect.ValorArreVect,
							ValorLista:       simbolos.ValArreglo.Clone(),
							Lintdim:          obj_vect.Lintdim.Clone(),
							Lintcap:          obj_vect.Lintcap.Clone(),
							PosicionTabla:    ent.Posicion,
							Linea:            dvect.Linea,
							Columna:          dvect.Columna}
						if ent.Nombre_Entorno != ent2.Nombre_Entorno {
							ent2.Posicion = ent2.Posicion - 1
							ent.Guardar_ArregloVector(dvect.Identificador, simbarre)
							ent2.Guardar_ArregloVector(dvect.Identificador, simbarre)
						} else {
							ent.Guardar_ArregloVector(dvect.Identificador, simbarre)
							ent2.Guardar_ArregloVector(dvect.Identificador, simbarre)
						}
						reportes.ReporteSimbolos(dvect.Identificador, "arreglo--"+reportes.ReportObteniendoSimbolos(obj_vect.TipoVect), ent.Nombre_Entorno, "--", strconv.Itoa(dvect.Linea), strconv.Itoa(dvect.Columna))

					} else if tipovect.Tipo == simbolos.STRUCT {
						simbarre := simbolos.Simbolo_ArreVect{Mutable: dvect.Mutable,
							EsArreVect:       0,
							TipoVect:         tipovect.Tipo,
							Nombre:           dvect.Identificador,
							Dimensiones:      dvect.Dimension.Len(),
							DimensionesLista: dvect.Dimension.Clone(),
							ValorArreVect:    obj_vect.ValorArreVect,
							ValorLista:       simbolos.ValArreglo.Clone(),
							Lintdim:          obj_vect.Lintdim.Clone(),
							Lintcap:          obj_vect.Lintcap.Clone(),
							PosicionTabla:    ent.Posicion,
							Linea:            dvect.Linea,
							Columna:          dvect.Columna}
						if ent.Nombre_Entorno != ent2.Nombre_Entorno {
							ent2.Posicion = ent2.Posicion - 1
							ent.Guardar_ArregloVector(dvect.Identificador, simbarre)
							ent2.Guardar_ArregloVector(dvect.Identificador, simbarre)
						} else {
							ent.Guardar_ArregloVector(dvect.Identificador, simbarre)
							ent2.Guardar_ArregloVector(dvect.Identificador, simbarre)
						}
						reportes.ReporteSimbolos(dvect.Identificador, "arreglo--"+reportes.ReportObteniendoSimbolos(tipovect.Tipo), ent.Nombre_Entorno, "--", strconv.Itoa(dvect.Linea), strconv.Itoa(dvect.Columna))
					} else {
						t := time.Now()
						reportes.Nerror("Los tipo no coinciden en los arreglos.", ent.Nombre_Entorno, strconv.Itoa(dvect.Linea), strconv.Itoa(dvect.Columna), t.Format("2006-01-02 15:04:05"))
					}
				} else if obj_vect.Lintdim.Len() != 1 {

					if tipovect.Tipo == obj_vect.TipoVect {
						simbarre := simbolos.Simbolo_ArreVect{Mutable: dvect.Mutable,
							EsArreVect:       0,
							TipoVect:         tipovect.Tipo,
							Nombre:           dvect.Identificador,
							Dimensiones:      simbolos.ValArreglo.Len(),
							DimensionesLista: dvect.Dimension.Clone(),
							ValorArreVect:    obj_vect.ValorArreVect,
							ValorLista:       simbolos.ValArreglo.Clone(),
							Lintdim:          obj_vect.Lintdim.Clone(),
							Lintcap:          obj_vect.Lintcap.Clone(),
							PosicionTabla:    ent.Posicion,
							Linea:            dvect.Linea,
							Columna:          dvect.Columna}
						if ent.Nombre_Entorno != ent2.Nombre_Entorno {
							ent2.Posicion = ent2.Posicion - 1
							ent.Guardar_ArregloVector(dvect.Identificador, simbarre)
							ent2.Guardar_ArregloVector(dvect.Identificador, simbarre)
						} else {
							ent.Guardar_ArregloVector(dvect.Identificador, simbarre)
							ent2.Guardar_ArregloVector(dvect.Identificador, simbarre)
						}
						reportes.ReporteSimbolos(dvect.Identificador, "arreglo--"+reportes.ReportObteniendoSimbolos(obj_vect.TipoVect), ent.Nombre_Entorno, "--", strconv.Itoa(dvect.Linea), strconv.Itoa(dvect.Columna))
					} else if tipovect.Tipo == simbolos.STRUCT {
						simbarre := simbolos.Simbolo_ArreVect{Mutable: dvect.Mutable,
							EsArreVect:       0,
							TipoVect:         tipovect.Tipo,
							Nombre:           dvect.Identificador,
							Dimensiones:      simbolos.ValArreglo.Len(),
							DimensionesLista: dvect.Dimension.Clone(),
							ValorArreVect:    obj_vect.ValorArreVect,
							ValorLista:       simbolos.ValArreglo.Clone(),
							Lintdim:          obj_vect.Lintdim.Clone(),
							Lintcap:          obj_vect.Lintcap.Clone(),
							PosicionTabla:    ent.Posicion,
							Linea:            dvect.Linea,
							Columna:          dvect.Columna}
						if ent.Nombre_Entorno != ent2.Nombre_Entorno {
							ent2.Posicion = ent2.Posicion - 1
							ent.Guardar_ArregloVector(dvect.Identificador, simbarre)
							ent2.Guardar_ArregloVector(dvect.Identificador, simbarre)
						} else {
							ent.Guardar_ArregloVector(dvect.Identificador, simbarre)
							ent2.Guardar_ArregloVector(dvect.Identificador, simbarre)
						}
						reportes.ReporteSimbolos(dvect.Identificador, "arreglo--"+reportes.ReportObteniendoSimbolos(tipovect.Tipo), ent.Nombre_Entorno, "--", strconv.Itoa(dvect.Linea), strconv.Itoa(dvect.Columna))
					} else {
						t := time.Now()
						reportes.Nerror("Los tipo no coinciden en los arreglos.", ent.Nombre_Entorno, strconv.Itoa(dvect.Linea), strconv.Itoa(dvect.Columna), t.Format("2006-01-02 15:04:05"))
					}
				} else {
					t := time.Now()
					reportes.Nerror("Las dimensiones de los arreglos no coinciden.", ent.Nombre_Entorno, strconv.Itoa(dvect.Linea), strconv.Itoa(dvect.Columna), t.Format("2006-01-02 15:04:05"))
				}
			}
		} else {
			t := time.Now()
			reportes.Nerror("No se puede declarar vector.", ent.Nombre_Entorno, strconv.Itoa(dvect.Linea), strconv.Itoa(dvect.Columna), t.Format("2006-01-02 15:04:05"))
		}
	} else {
		t := time.Now()
		reportes.Nerror("Variables con el mismo nombre en el entorno.", ent.Nombre_Entorno, strconv.Itoa(dvect.Linea), strconv.Itoa(dvect.Columna), t.Format("2006-01-02 15:04:05"))
	}
	simbolos.ValArreglo.Clear()
	return 0
}

func (dvect DeclaVector) Compilar_Instruccion(ent *entorno.Entorno, gen *generador.Generador_C3D) interface{} {
	if ent.Existe_ArreVect(dvect.Identificador) {
		gen.Agregar_Comentario("INICIO DECLARACION VECTOR -- " + dvect.Identificador)
		simarreglo := ent.Obtener_ArreVect(dvect.Identificador)
		if simarreglo.TipoVect == simbolos.TEXTO || simarreglo.TipoVect == simbolos.YTEXTO {
			tempo1 := gen.Crear_temporal()
			posi := simarreglo.PosicionTabla
			gen.Agregar_Logica(tempo1 + " = SP + " + strconv.Itoa(posi) + ";\t\t// POSICION: " + dvect.Identificador)
			gen.Agregar_Stack(tempo1, "HP")
			tempo2 := gen.Crear_temporal()
			gen.Agregar_Logica(tempo2 + " = HP;")
			gen.Agregar_Logica("HP = HP +" + strconv.Itoa(simarreglo.ValorLista.Len()) + ";")
			gen.Agregar_Logica("HEAP[int(HP)] = -2;\nHP = HP + 1;")
			tempo3 := gen.Crear_temporal()
			for i := 0; i < simarreglo.ValorLista.Len(); i++ {
				act := simarreglo.ValorLista.GetValue(i).(simbolos.Valor)
				guardar := fmt.Sprintf("%v", act.Valor)
				gen.Agregar_Logica(tempo3 + " = " + tempo2 + " + " + strconv.Itoa(i) + ";")
				gen.Agregar_Logica("HEAP[int(" + tempo3 + ")] = HP;")
				for _, txt := range guardar {
					f := int(txt)
					gen.Agregar_Logica("HEAP[int(HP)] = " + fmt.Sprintf("%v", f) + "; //LETRA-> " + string(txt))
					gen.Agregar_Logica("HP = HP + 1;")
				}
				gen.Agregar_Logica("HEAP[int(HP)] = -1;\nHP = HP + 1;")
			}
		} else {
			tempo1 := gen.Crear_temporal()
			posi := simarreglo.PosicionTabla
			gen.Agregar_Logica(tempo1 + " = SP + " + strconv.Itoa(posi) + ";\t\t// POSICION: " + dvect.Identificador)
			gen.Agregar_Stack(tempo1, "HP")
			tempo2 := gen.Crear_temporal()
			tempo3 := gen.Crear_temporal()
			gen.Agregar_Logica(tempo2 + " = HP;")
			gen.Agregar_Logica("HP = HP +" + strconv.Itoa(simarreglo.ValorLista.Len()) + ";")
			gen.Agregar_Logica("HEAP[int(HP)] = -2;\nHP = HP + 1;")
			for i := 0; i < simarreglo.ValorLista.Len(); i++ {
				act := simarreglo.ValorLista.GetValue(i).(simbolos.Valor)
				gen.Agregar_Logica(tempo3 + " = " + tempo2 + " + " + strconv.Itoa(i) + ";")
				gen.Agregar_Logica("HEAP[int(" + tempo3 + ")] = " + fmt.Sprintf("%v", act.Valor) + ";")
			}
		}
		gen.Agregar_Comentario("FIN DECLARACION VECTOR -- " + dvect.Identificador)
		gen.LiberarTodosTemporales()
	}
	return 0
}
