package cloop

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

type Loop struct {
	Linstr  *arraylist.List
	Linea   int
	Columna int
}

func Nloop(l *arraylist.List, li int, c int) Loop {
	loops := Loop{Linstr: l, Linea: li, Columna: c}
	return loops
}

func (l Loop) Ejecutar_Instruccion(ent *entorno.Entorno, ent2 *entorno.Entorno) interface{} {
	//ENTORNO IF
	entorno_funcion := entorno.Nuevo_Entorno("loop_"+ent.Nombre_Entorno, ent)
	for i := 0; i < l.Linstr.Len(); i++ {
		instr_actual := l.Linstr.GetValue(i).(interfaces.Instruccion)
		valor_instr := instr_actual.Ejecutar_Instruccion(entorno_funcion, ent2)

		if valor_instr != 0 {
			valores := valor_instr.(simbolos.Valor)
			if valores.Tipo == simbolos.BREAK {
				if valores.Valor == "break" {
					return 0
				} else {
					t := time.Now()
					reportes.Nerror("No se puede retornar un valor con la sentencia break", ent.Nombre_Entorno, strconv.Itoa(l.Linea), strconv.Itoa(l.Columna), t.Format("2006-01-02 15:04:05"))
					return 0
				}
			} else if valores.Tipo == simbolos.CONTINUE {
				break
			} else {
				valor_retorno := valor_instr.(simbolos.Valor)
				return valor_retorno
			}

		}
	}
	// contador := 0
	// for contador != 1 {

	// }

	return 0
}

func (l Loop) Ejecutar_Expresion(ent *entorno.Entorno) simbolos.Valor {
	// contador := 0
	// for contador != 1 {

	// }
	for i := 0; i < l.Linstr.Len(); i++ {
		instr_actual := l.Linstr.GetValue(i).(interfaces.Instruccion)
		valor_instr := instr_actual.Ejecutar_Instruccion(ent, ent)

		if valor_instr != 0 {
			valores := valor_instr.(simbolos.Valor)
			if valores.Tipo == simbolos.BREAK {
				if valores.Valor == "break" {
					return simbolos.Valor{Tipo: simbolos.NULL, Valor: 0}
				} else {
					sal := valores.Valor.(simbolos.Valor)
					return sal
				}
			} else if valores.Tipo == simbolos.CONTINUE {
				break
			} else {
				valor_retorno := valor_instr.(simbolos.Valor)
				return valor_retorno
			}

		}
	}
	return simbolos.Valor{Tipo: simbolos.NULL, Valor: -1}
}

func (l Loop) Compilar_Instruccion(ent *entorno.Entorno, gen *generador.Generador_C3D) interface{} {
	gen.Agregar_Comentario("ENTRANDO CICLO LOOP")
	etiqueta_entrada := gen.Crear_label()
	gen.Eliminar_label(etiqueta_entrada)
	etiqueta_salida := gen.Crear_label()
	gen.Eliminar_label(etiqueta_salida)
	volver_entrada := "goto " + etiqueta_entrada + ";"
	salidas := "goto " + etiqueta_salida + ";"
	simbtrans := simbolos.EtiquetasTransferencia{Etiqueta_Entrada: volver_entrada, Etiqueta_Salida: salidas}
	simbolos.ListaTransferencia.Add(simbtrans)
	gen.Agregar_label(etiqueta_entrada)
	for i := 0; i < l.Linstr.Len(); i++ {
		instr := l.Linstr.GetValue(i).(interfaces.Instruccion)
		val := instr.Compilar_Instruccion(ent, gen)
		if val != 0 {
			if val == "break" {
				pos := simbolos.ListaTransferencia.Len()
				valtrans := simbolos.ListaTransferencia.GetValue(pos - 1).(simbolos.EtiquetasTransferencia)
				gen.Agregar_Logica(valtrans.Etiqueta_Salida)
			} else if val == "continue" {
				pos := simbolos.ListaTransferencia.Len()
				valtrans := simbolos.ListaTransferencia.GetValue(pos - 1).(simbolos.EtiquetasTransferencia)
				gen.Agregar_Logica(valtrans.Etiqueta_Entrada)
			}
		}
	}
	gen.Agregar_Logica(volver_entrada)
	if gen.Lista_etiqueta.Len() != 0 {
		for i := 0; i < gen.Lista_etiqueta.Len(); i++ {
			la := gen.Lista_etiqueta.GetValue(i).(string)
			gen.Agregar_label(la)
			gen.Agregar_Logica(salidas)
			gen.Eliminar_label(la)
		}
	}
	gen.Agregar_label(etiqueta_salida)
	gen.Agregar_Comentario("SALIENDO CICLO LOOP")
	pos := simbolos.ListaTransferencia.Len()
	simbolos.ListaTransferencia.RemoveAtIndex(pos - 1)
	return 0
}

func (l Loop) Compilar_Expresion(ent *entorno.Entorno, gen *generador.Generador_C3D) simbolos.ValoresC3D {

	return simbolos.ValoresC3D{Valor: "0", EsTemporal: false, Tipo: simbolos.INTEGER, Label_verdadera: "", Label_false: ""}
}
