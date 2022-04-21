package cwhile

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

type C_While struct {
	Condi   interfaces.Expresion
	L_instr *arraylist.List
	Linea   int
	Columna int
}

func Nwhile(c interfaces.Expresion, linstr *arraylist.List, l int, co int) C_While {
	whi := C_While{Condi: c, L_instr: linstr, Linea: l, Columna: co}
	return whi
}

func (w C_While) Ejecutar_Instruccion(ent *entorno.Entorno, ent2 *entorno.Entorno) interface{} {
	entorno_funcion := entorno.Nuevo_Entorno("while_"+ent.Nombre_Entorno, ent)
	var condi simbolos.Valor = w.Condi.Ejecutar_Expresion(entorno_funcion)
	if condi.Tipo == simbolos.BOOLEAN {
		for i := 0; i < w.L_instr.Len(); i++ {
			instr_actual := w.L_instr.GetValue(i).(interfaces.Instruccion)
			valor_instr := instr_actual.Ejecutar_Instruccion(entorno_funcion, ent2)

			if valor_instr != 0 {
				valores := valor_instr.(simbolos.Valor)
				if valores.Tipo == simbolos.BREAK {
					if valores.Valor == "break" {
						return 0
					} else {
						t := time.Now()
						reportes.Nerror("No se puede retornar un valor con la sentencia break", ent.Nombre_Entorno, strconv.Itoa(w.Linea), strconv.Itoa(w.Columna), t.Format("2006-01-02 15:04:05"))
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
		// for condi.Valor.(bool) {

		// 	condi = w.Condi.Ejecutar_Expresion(entorno_funcion)
		// }
	} else {
		t := time.Now()
		reportes.Nerror("Se esperaba una condicion de tipo bool", ent.Nombre_Entorno, strconv.Itoa(w.Linea), strconv.Itoa(w.Columna), t.Format("2006-01-02 15:04:05"))
	}

	return 0
}

func (w C_While) Compilar_Instruccion(ent *entorno.Entorno, gen *generador.Generador_C3D) interface{} {
	gen.Agregar_Comentario("ENTRANDO CICLO WHILE")
	etiqueta_entrada := gen.Crear_label()
	gen.Eliminar_label(etiqueta_entrada)
	etiqueta_salida := gen.Crear_label()
	gen.Eliminar_label(etiqueta_salida)
	volver_entrada := "goto " + etiqueta_entrada + ";"
	salidas := "goto " + etiqueta_salida + ";"
	simbtrans := simbolos.EtiquetasTransferencia{Etiqueta_Entrada: volver_entrada, Etiqueta_Salida: salidas}
	simbolos.ListaTransferencia.Add(simbtrans)
	gen.Agregar_label(etiqueta_entrada)
	condi := w.Condi.Compilar_Expresion(ent, gen)
	//gen.Agregar_Logica(condi.Valor)
	gen.Agregar_label(condi.Label_verdadera)
	gen.Eliminar_label(condi.Label_verdadera)
	for i := 0; i < w.L_instr.Len(); i++ {
		instr := w.L_instr.GetValue(i).(interfaces.Instruccion)
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
	gen.Agregar_Comentario("SALIENDO CICLO WHILE")
	pos := simbolos.ListaTransferencia.Len()
	simbolos.ListaTransferencia.RemoveAtIndex(pos - 1)
	return 0
}
