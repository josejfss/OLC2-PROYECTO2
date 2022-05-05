package scondicional

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

type Si struct {
	Condicion   interfaces.Expresion
	Lista_instr *arraylist.List
	Lista_else  *arraylist.List
	Linea       int
	Columna     int
}

func Nsi(condi interfaces.Expresion, lifs *arraylist.List, lelse *arraylist.List) Si {
	ifs := Si{Condicion: condi, Lista_instr: lifs, Lista_else: lelse}
	return ifs
}

func (ifs Si) Ejecutar_Instruccion(ent *entorno.Entorno, ent2 *entorno.Entorno) interface{} {

	if_condi := ifs.Condicion.Ejecutar_Expresion(ent)
	todito := simbolos.SimboloTodo{Tipo: "if", Nombre: "", Pos: ent.Tamaño}
	ent2.GuardaLTodo(todito)
	if if_condi.Tipo == simbolos.BOOLEAN {
		//ENTORNO IF
		var entorno_funcion *entorno.Entorno = entorno.Nuevo_Entorno("if_"+ent.Nombre_Entorno, ent)
		if ifs.Lista_instr.Len() != 0 {
			for i := 0; i < ifs.Lista_instr.Len(); i++ {
				instr_actual := ifs.Lista_instr.GetValue(i).(interfaces.Instruccion)
				valor_instr := instr_actual.Ejecutar_Instruccion(entorno_funcion, ent2)

				if valor_instr != 0 {
					valores := valor_instr.(simbolos.Valor)
					if valores.Tipo == simbolos.BREAK {
						if valores.Valor == "break" {
							return 0
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
		}
		if ifs.Lista_else.Len() != 0 {
			for i := 0; i < ifs.Lista_else.Len(); i++ {
				instr_actual := ifs.Lista_else.GetValue(i).(interfaces.Instruccion)
				valor_instr := instr_actual.Ejecutar_Instruccion(entorno_funcion, ent2)

				if valor_instr != 0 {
					valores := valor_instr.(simbolos.Valor)
					if valores.Tipo == simbolos.BREAK {
						if valores.Valor == "break" {
							return 0
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
		}
	} else {
		t := time.Now()
		reportes.Nerror("No se puede realizar el if: condicion no es de tipo bool", ent.Nombre_Entorno, strconv.Itoa(ifs.Linea), strconv.Itoa(ifs.Columna), t.Format("2006-01-02 15:04:05"))
	}
	return 0
}

func (ifs Si) Ejecutar_Expresion(ent *entorno.Entorno) simbolos.Valor {
	if_condi := ifs.Condicion.Ejecutar_Expresion(ent)
	if if_condi.Tipo == simbolos.BOOLEAN {
		if if_condi.Valor.(bool) {
			if ifs.Lista_instr.Len() != 0 {
				for i := 0; i < ifs.Lista_instr.Len(); i++ {
					instr_actual := ifs.Lista_instr.GetValue(i).(interfaces.Instruccion)
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
			}
		} else {
			if ifs.Lista_else.Len() != 0 {
				for i := 0; i < ifs.Lista_else.Len(); i++ {
					instr_actual := ifs.Lista_else.GetValue(i).(interfaces.Instruccion)
					valor_instr := instr_actual.Ejecutar_Instruccion(ent, ent)

					if valor_instr != 0 {
						valores := valor_instr.(simbolos.Valor)
						if valores.Tipo == simbolos.BREAK {
							if valores.Valor == "break" {
								return simbolos.Valor{Tipo: simbolos.NULL, Valor: 0}
							} else {
								if i == ifs.Lista_else.Len()-1 {
									valor_retorno := valor_instr.(simbolos.Valor)
									return valor_retorno
								}
							}
						} else if valores.Tipo == simbolos.CONTINUE {
							break
						} else {
							if i == ifs.Lista_else.Len()-1 {
								valor_retorno := valor_instr.(simbolos.Valor)
								return valor_retorno
							}
						}

					}
				}
			}
		}
	} else {
		t := time.Now()
		reportes.Nerror("No se puede realizar el if: condicion no es de tipo bool", ent.Nombre_Entorno, strconv.Itoa(ifs.Linea), strconv.Itoa(ifs.Columna), t.Format("2006-01-02 15:04:05"))
	}
	return simbolos.Valor{Tipo: simbolos.NULL, Valor: 0}
}

func (ifs Si) Compilar_Instruccion(ent *entorno.Entorno, gen *generador.Generador_C3D) interface{} {
	valtodi := ent.ListaTodo.GetValue(0).(simbolos.SimboloTodo)
	ent.EliminarLTodo()
	if valtodi.Tipo == "if" {
		gen.Agregar_Comentario("ENTRANDO SENTENCIA IF")
		condis := ifs.Condicion.Compilar_Expresion(ent, gen)
		gen.Agregar_Logica("SP = SP + " + strconv.Itoa(valtodi.Pos) + ";")
		etiqueta_salida := gen.Crear_label()
		gen.Eliminar_label(etiqueta_salida)
		salidas := "goto " + etiqueta_salida + ";"
		gen.Agregar_label(condis.Label_verdadera)
		gen.Eliminar_label(condis.Label_verdadera)
		for i := 0; i < ifs.Lista_instr.Len(); i++ {
			instr := ifs.Lista_instr.GetValue(i).(interfaces.Instruccion)
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
		gen.Agregar_Logica(salidas)
		gen.Agregar_label(condis.Label_false)
		gen.Eliminar_label(condis.Label_false)
		for i := 0; i < ifs.Lista_else.Len(); i++ {
			instr := ifs.Lista_else.GetValue(i).(interfaces.Instruccion)
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
		gen.Agregar_Logica(salidas)
		if gen.Lista_etiqueta.Len() != 0 {
			for i := 0; i < gen.Lista_etiqueta.Len(); i++ {
				la := gen.Lista_etiqueta.GetValue(i).(string)
				gen.Agregar_label(la)
				gen.Agregar_Logica(salidas)
				// gen.Eliminar_label(la)
			}
			for i := 0; i < gen.Lista_etiqueta.Len()+1; i++ {
				la := gen.Lista_etiqueta.GetValue(i).(string)
				// gen.Agregar_label(la)
				// gen.Agregar_Logica(salidas)
				gen.Eliminar_label(la)
			}
		}
		gen.Agregar_label(etiqueta_salida)
		gen.Agregar_Logica("SP = SP - " + strconv.Itoa(valtodi.Pos) + ";")
		gen.Agregar_Comentario("SALIENDO SENTENCIA IF")
	}

	return 0
}

func (ifs Si) Compilar_Expresion(ent *entorno.Entorno, gen *generador.Generador_C3D) simbolos.ValoresC3D {

	return simbolos.ValoresC3D{Valor: "0", EsTemporal: false, Tipo: simbolos.INTEGER, Label_verdadera: "", Label_false: ""}
}
