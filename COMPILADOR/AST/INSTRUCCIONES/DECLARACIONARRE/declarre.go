package declaracionarre

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

type TipoDeclaArre struct {
	Tipito    interfaces.Expresion
	Dimension interfaces.Expresion
}

func Ntipodeclarre(tip interfaces.Expresion, dim interfaces.Expresion) TipoDeclaArre {
	tiparrdecl := TipoDeclaArre{Tipito: tip, Dimension: dim}
	return tiparrdecl
}

type DeclaArre struct {
	Mutable       bool
	Identificador string
	Dimension     *arraylist.List
	ValExp        interfaces.Expresion
	Linea         int
	Columna       int
}

func Ndeclaarre(m bool, id string, dim *arraylist.List, va interfaces.Expresion, l int, c int) DeclaArre {
	arr := DeclaArre{Mutable: m, Identificador: id, Dimension: dim, ValExp: va, Linea: l, Columna: c}
	return arr
}

func (darre DeclaArre) Ejecutar_Instruccion(ent *entorno.Entorno, ent2 *entorno.Entorno) interface{} {
	dimtip := darre.Dimension.GetValue(0).(TipoDeclaArre)
	tip := dimtip.Tipito.Ejecutar_Expresion(ent)
	dim := dimtip.Dimension.Ejecutar_Expresion(ent)
	resultado := darre.ValExp.Ejecutar_Expresion(ent)
	if !ent.Existe_Variable(darre.Identificador) {
		if !ent.Existe_ArreVect(darre.Identificador) {
			if resultado.Tipo == simbolos.ARRAY {
				val_resultado := resultado.Valor.(simbolos.Valor)
				obj_vect := val_resultado.Valor.(simbolos.Simbolo_ArreVect)
				if darre.Dimension.Len() == 1 {
					if obj_vect.TipoVect == tip.Tipo {
						simbarre := simbolos.Simbolo_ArreVect{Mutable: darre.Mutable,
							EsArreVect:       0,
							TipoVect:         tip.Tipo,
							Nombre:           darre.Identificador,
							Dimensiones:      dim.Valor.(int),
							DimensionesLista: darre.Dimension.Clone(),
							ValorArreVect:    obj_vect.ValorArreVect,
							ValorLista:       simbolos.ValArreglo.Clone(),
							Lintdim:          obj_vect.Lintdim.Clone(),
							Lintcap:          obj_vect.Lintcap.Clone(),
							PosicionTabla:    ent.Posicion,
							Linea:            darre.Linea,
							Columna:          darre.Columna}
						if ent.Nombre_Entorno != ent2.Nombre_Entorno {
							ent2.Posicion = ent2.Posicion - 1
							ent.Guardar_ArregloVector(darre.Identificador, simbarre)
							ent2.Guardar_ArregloVector(darre.Identificador, simbarre)
						} else {
							ent.Guardar_ArregloVector(darre.Identificador, simbarre)
							ent2.Guardar_ArregloVector(darre.Identificador, simbarre)
						}
						reportes.ReporteSimbolos(darre.Identificador, "arreglo--"+reportes.ReportObteniendoSimbolos(obj_vect.TipoVect), ent.Nombre_Entorno, "--", strconv.Itoa(darre.Linea), strconv.Itoa(darre.Columna))

					} else if tip.Tipo == simbolos.STRUCT {
						simbarre := simbolos.Simbolo_ArreVect{Mutable: darre.Mutable,
							EsArreVect:       0,
							TipoVect:         tip.Tipo,
							Nombre:           darre.Identificador,
							Dimensiones:      dim.Valor.(int),
							DimensionesLista: darre.Dimension.Clone(),
							ValorArreVect:    obj_vect.ValorArreVect,
							ValorLista:       simbolos.ValArreglo.Clone(),
							Lintdim:          obj_vect.Lintdim.Clone(),
							Lintcap:          obj_vect.Lintcap.Clone(),
							PosicionTabla:    ent.Posicion,
							Linea:            darre.Linea,
							Columna:          darre.Columna}
						if ent.Nombre_Entorno != ent2.Nombre_Entorno {
							ent2.Posicion = ent2.Posicion - 1
							ent.Guardar_ArregloVector(darre.Identificador, simbarre)
							ent2.Guardar_ArregloVector(darre.Identificador, simbarre)
						} else {
							ent.Guardar_ArregloVector(darre.Identificador, simbarre)
							ent2.Guardar_ArregloVector(darre.Identificador, simbarre)
						}
						reportes.ReporteSimbolos(darre.Identificador, "arreglo--"+reportes.ReportObteniendoSimbolos(tip.Tipo), ent.Nombre_Entorno, "--", strconv.Itoa(darre.Linea), strconv.Itoa(darre.Columna))
					} else {
						t := time.Now()
						reportes.Nerror("Los tipo no coinciden en los arreglos.", ent.Nombre_Entorno, strconv.Itoa(darre.Linea), strconv.Itoa(darre.Columna), t.Format("2006-01-02 15:04:05"))
					}
				} else if obj_vect.Lintdim.Len() != 1 {
					tam := 0
					for i := 0; i < darre.Dimension.Len(); i++ {
						if i == 0 {
							act := darre.Dimension.GetValue(i).(TipoDeclaArre)
							act1 := act.Dimension.Ejecutar_Expresion(ent)
							tam = act1.Valor.(int)
						} else {
							act2 := darre.Dimension.GetValue(i).(interfaces.Expresion).Ejecutar_Expresion(ent)
							tam = tam * act2.Valor.(int)
						}
					}
					if tip.Tipo == obj_vect.TipoVect {
						simbarre := simbolos.Simbolo_ArreVect{Mutable: darre.Mutable,
							EsArreVect:       0,
							TipoVect:         tip.Tipo,
							Nombre:           darre.Identificador,
							Dimensiones:      tam,
							DimensionesLista: darre.Dimension.Clone(),
							ValorArreVect:    obj_vect.ValorArreVect,
							ValorLista:       simbolos.ValArreglo.Clone(),
							Lintdim:          obj_vect.Lintdim.Clone(),
							Lintcap:          obj_vect.Lintcap.Clone(),
							PosicionTabla:    ent.Posicion,
							Linea:            darre.Linea,
							Columna:          darre.Columna}
						if ent.Nombre_Entorno != ent2.Nombre_Entorno {
							ent2.Posicion = ent2.Posicion - 1
							ent.Guardar_ArregloVector(darre.Identificador, simbarre)
							ent2.Guardar_ArregloVector(darre.Identificador, simbarre)
						} else {
							ent.Guardar_ArregloVector(darre.Identificador, simbarre)
							ent2.Guardar_ArregloVector(darre.Identificador, simbarre)
						}
						reportes.ReporteSimbolos(darre.Identificador, "arreglo--"+reportes.ReportObteniendoSimbolos(obj_vect.TipoVect), ent.Nombre_Entorno, "--", strconv.Itoa(darre.Linea), strconv.Itoa(darre.Columna))
					} else if tip.Tipo == simbolos.STRUCT {
						simbarre := simbolos.Simbolo_ArreVect{Mutable: darre.Mutable,
							EsArreVect:       0,
							TipoVect:         tip.Tipo,
							Nombre:           darre.Identificador,
							Dimensiones:      tam,
							DimensionesLista: darre.Dimension.Clone(),
							ValorArreVect:    obj_vect.ValorArreVect,
							ValorLista:       simbolos.ValArreglo.Clone(),
							Lintdim:          obj_vect.Lintdim.Clone(),
							Lintcap:          obj_vect.Lintcap.Clone(),
							PosicionTabla:    ent.Posicion,
							Linea:            darre.Linea,
							Columna:          darre.Columna}
						if ent.Nombre_Entorno != ent2.Nombre_Entorno {
							ent2.Posicion = ent2.Posicion - 1
							ent.Guardar_ArregloVector(darre.Identificador, simbarre)
							ent2.Guardar_ArregloVector(darre.Identificador, simbarre)
						} else {
							ent.Guardar_ArregloVector(darre.Identificador, simbarre)
							ent2.Guardar_ArregloVector(darre.Identificador, simbarre)
						}
						reportes.ReporteSimbolos(darre.Identificador, "arreglo--"+reportes.ReportObteniendoSimbolos(tip.Tipo), ent.Nombre_Entorno, "--", strconv.Itoa(darre.Linea), strconv.Itoa(darre.Columna))
					} else {
						t := time.Now()
						reportes.Nerror("Los tipo no coinciden en los arreglos.", ent.Nombre_Entorno, strconv.Itoa(darre.Linea), strconv.Itoa(darre.Columna), t.Format("2006-01-02 15:04:05"))
					}
				} else {
					t := time.Now()
					reportes.Nerror("Las dimensiones de los arreglos no coinciden.", ent.Nombre_Entorno, strconv.Itoa(darre.Linea), strconv.Itoa(darre.Columna), t.Format("2006-01-02 15:04:05"))
				}
			}
		} else {
			t := time.Now()
			reportes.Nerror("No se puede declarar vector.", ent.Nombre_Entorno, strconv.Itoa(darre.Linea), strconv.Itoa(darre.Columna), t.Format("2006-01-02 15:04:05"))
		}
	} else {
		t := time.Now()
		reportes.Nerror("Variables con el mismo nombre en el entorno.", ent.Nombre_Entorno, strconv.Itoa(darre.Linea), strconv.Itoa(darre.Columna), t.Format("2006-01-02 15:04:05"))
	}
	simbolos.ValArreglo.Clear()
	return 0
}

func (darre DeclaArre) Compilar_Instruccion(ent *entorno.Entorno, gen *generador.Generador_C3D) interface{} {

	if ent.Existe_ArreVect(darre.Identificador) {
		gen.Agregar_Comentario("INICIO DECLARACION ARREGLO -- " + darre.Identificador)
		simarreglo := ent.Obtener_ArreVect(darre.Identificador)
		if simarreglo.TipoVect == simbolos.TEXTO || simarreglo.TipoVect == simbolos.YTEXTO {
			tempo1 := gen.Crear_temporal()
			posi := simarreglo.PosicionTabla
			gen.Agregar_Logica(tempo1 + " = SP + " + strconv.Itoa(posi) + ";\t\t// POSICION: " + darre.Identificador)
			gen.Agregar_Stack(tempo1, "HP")
			tempo2 := gen.Crear_temporal()
			gen.Agregar_Logica(tempo2 + " = HP;")
			gen.Agregar_Logica("HP = HP +" + strconv.Itoa(simarreglo.ValorLista.Len()) + ";")
			gen.Agregar_Logica("HEAP[(int)HP] = -2;\nHP = HP + 1;")
			tempo3 := gen.Crear_temporal()
			for i := 0; i < simarreglo.ValorLista.Len(); i++ {
				act := simarreglo.ValorLista.GetValue(i).(simbolos.Valor)
				guardar := fmt.Sprintf("%v", act.Valor)
				gen.Agregar_Logica(tempo3 + " = " + tempo2 + " + " + strconv.Itoa(i) + ";")
				gen.Agregar_Logica("HEAP[(int)" + tempo3 + "] = HP;")
				for _, txt := range guardar {
					f := int(txt)
					gen.Agregar_Logica("HEAP[(int)HP] = " + fmt.Sprintf("%v", f) + "; //LETRA-> " + string(txt))
					gen.Agregar_Logica("HP = HP + 1;")
				}
				gen.Agregar_Logica("HEAP[(int)HP] = -1;\nHP = HP + 1;")
			}
		} else {
			tempo1 := gen.Crear_temporal()
			posi := simarreglo.PosicionTabla
			gen.Agregar_Logica(tempo1 + " = SP + " + strconv.Itoa(posi) + ";\t\t// POSICION: " + darre.Identificador)
			gen.Agregar_Stack(tempo1, "HP")
			tempo2 := gen.Crear_temporal()
			tempo3 := gen.Crear_temporal()
			gen.Agregar_Logica(tempo2 + " = HP;")
			gen.Agregar_Logica("HP = HP +" + strconv.Itoa(simarreglo.Dimensiones) + ";")
			gen.Agregar_Logica("HEAP[(int)HP] = -2;\nHP = HP + 1;")
			if simarreglo.TipoVect == simbolos.INTEGER {
				for i := 0; i < simarreglo.ValorLista.Len(); i++ {
					act := simarreglo.ValorLista.GetValue(i).(simbolos.Valor)
					gen.Agregar_Logica(tempo3 + " = " + tempo2 + " + " + strconv.Itoa(i) + ";")
					gen.Agregar_Logica("HEAP[(int)" + tempo3 + "] = " + fmt.Sprintf("%v", act.Valor) + ";")
				}
			}

		}
		gen.Agregar_Comentario("FIN DECLARACION ARREGLO -- " + darre.Identificador)
		gen.LiberarTodosTemporales()
	}
	return 0
}
