package smatch

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

type SenteMatch struct {
	Condicion interfaces.Expresion
	L_instr   *arraylist.List
	Linea     int
	Columna   int
}

func Nmatch(c interfaces.Expresion, l *arraylist.List) SenteMatch {
	mat := SenteMatch{Condicion: c, L_instr: l}
	return mat
}

func (mat SenteMatch) Ejecutar_Instruccion(ent *entorno.Entorno, ent2 *entorno.Entorno) interface{} {

	entorno_match := entorno.Nuevo_Entorno("match_"+ent.Nombre_Entorno, ent)
	//c_pivote := mat.Condicion.Ejecutar_Expresion(entorno_match)
	todito := simbolos.SimboloTodo{Tipo: "match", Nombre: "", Pos: ent.Tamaño}
	ent2.GuardaLTodo(todito)
	//ENTORNO IF
	for i := 0; i < mat.L_instr.Len(); i++ {
		actual := mat.L_instr.GetValue(i).(Matches)
		//c_compa := actual.Condicion_match.Ejecutar_Expresion(entorno_match)
		for i := 0; i < actual.L_instr_match.Len(); i++ {
			instr_actual := actual.L_instr_match.GetValue(i).(interfaces.Instruccion)
			valor_instr := instr_actual.Ejecutar_Instruccion(entorno_match, ent2)

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

	return 0
}

func (mat SenteMatch) Compilar_Instruccion(ent *entorno.Entorno, gen *generador.Generador_C3D) interface{} {
	valtodi := ent.ListaTodo.GetValue(0).(simbolos.SimboloTodo)
	ent.EliminarLTodo()
	if valtodi.Tipo == "match" {
		gen.Agregar_Comentario("ENTRANDO SENTENCIA MATCH")
		gen.Agregar_Logica("SP = SP + " + strconv.Itoa(valtodi.Pos) + ";")
		temp_pivote := gen.Crear_temporal()
		igual := mat.Condicion.Compilar_Expresion(ent, gen)
		gen.Agregar_Logica(temp_pivote + " = " + igual.Valor + ";")
		etiqueta_salida := gen.Crear_label()
		gen.Eliminar_label(etiqueta_salida)
		etiqueta_default := gen.Crear_label()
		gen.Eliminar_label(etiqueta_default)
		salir := "goto " + etiqueta_salida + ";"
		//var entorno_funcion *entorno.Entorno = entorno.Nuevo_Entorno("match_"+ent.Nombre_Entorno, ent)
		for i := 0; i < mat.L_instr.Len(); i++ {
			actual := mat.L_instr.GetValue(i).(Matches)
			comparar := actual.Condicion_match.Compilar_Expresion(ent, gen)
			if !actual.Monton {
				if comparar.Valor != "_" {
					eti_verdad := gen.Crear_label()
					gen.Eliminar_label(eti_verdad)
					eti_falso := gen.Crear_label()
					gen.Eliminar_label(eti_falso)
					gen.Agregar_Logica("if ( " + temp_pivote + " == " + comparar.Valor + " ) { goto " + eti_verdad + "; }\ngoto " + eti_falso + ";\n")
					gen.Agregar_label(eti_verdad)
					if actual.L_instr_match.Len() != 0 {
						for i := 0; i < actual.L_instr_match.Len(); i++ {
							instr_actual := actual.L_instr_match.GetValue(i).(interfaces.Instruccion)
							val := instr_actual.Compilar_Instruccion(ent, gen)
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
							//valor_instr := instr_actual.Compilar_Instruccion(ent, gen)
							// if valor_instr != 0 {
							// 	if valor_instr == "break" {
							// 		gen.Agregar_Logica(salidas)
							// 	} else if valor_instr == "continue" {
							// 		gen.Agregar_Logica(volver_entrada)
							// 	}
							// }
						}
					}
					gen.Agregar_Logica(salir)
					gen.Agregar_label(eti_falso)
				} else {
					if actual.L_instr_match.Len() != 0 {
						for i := 0; i < actual.L_instr_match.Len(); i++ {
							instr_actual := actual.L_instr_match.GetValue(i).(interfaces.Instruccion)
							val := instr_actual.Compilar_Instruccion(ent, gen)
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
							//valor_instr := instr_actual.Compilar_Instruccion(ent, gen)
							// if valor_instr != 0 {
							// 	if valor_instr == "break" {
							// 		gen.Agregar_Logica(salidas)
							// 	} else if valor_instr == "continue" {
							// 		gen.Agregar_Logica(volver_entrada)
							// 	}
							// }
						}
					}
				}
			} else {
				//LISTA DE CONDICIONES JUNTAS
				if actual.Lista_condiciones.Len() != 0 {
					var etiverdad string = ""
					var etifalse string = ""
					for i := 0; i < actual.Lista_condiciones.Len(); i++ {
						nuevo_actual := actual.Lista_condiciones.GetValue(i).(interfaces.Expresion)
						nuevo := nuevo_actual.Compilar_Expresion(ent, gen)
						eti_verdad := gen.Crear_label()
						gen.Eliminar_label(eti_verdad)
						eti_falso := gen.Crear_label()
						gen.Eliminar_label(eti_falso)
						gen.Agregar_Logica("if ( " + temp_pivote + " == " + nuevo.Valor + " ) { goto " + eti_verdad + "; }\ngoto " + eti_falso + ";\n")
						gen.Agregar_label(eti_falso)
						etifalse = eti_falso
						etiverdad += eti_verdad + ":\n"
					}
					pos := gen.Obtener_TamCodigo()
					gen.Obtener_codigo().RemoveAtIndex(pos - 1)
					gen.Agregar_Logica(etiverdad)
					if actual.L_instr_match.Len() != 0 {
						for i := 0; i < actual.L_instr_match.Len(); i++ {
							instr_actual := actual.L_instr_match.GetValue(i).(interfaces.Instruccion)
							val := instr_actual.Compilar_Instruccion(ent, gen)
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
							//valor_instr := instr_actual.Compilar_Instruccion(ent, gen)
							// if valor_instr != 0 {
							// 	if valor_instr == "break" {
							// 		gen.Agregar_Logica(salidas)
							// 	} else if valor_instr == "continue" {
							// 		gen.Agregar_Logica(volver_entrada)
							// 	}
							// }
						}
					}
					gen.Agregar_Logica(salir)
					gen.Agregar_Logica(etifalse + ":")
				}
			}

		}
		gen.Agregar_label(etiqueta_salida)
		gen.Agregar_Logica("SP = SP - " + strconv.Itoa(valtodi.Pos) + ";")
		gen.Agregar_Comentario("SALIENDO SENTENCIA MATCH")
		gen.LiberarTodosTemporales()
	}

	return 0
}

func (mat SenteMatch) Ejecutar_Expresion(ent *entorno.Entorno) simbolos.Valor {
	c_pivote := mat.Condicion.Ejecutar_Expresion(ent)
	for i := 0; i < mat.L_instr.Len(); i++ {
		actual := mat.L_instr.GetValue(i).(Matches)
		c_compa := actual.Condicion_match.Ejecutar_Expresion(ent)

		if !actual.Monton {
			if c_compa.Valor != "_" {
				if c_pivote.Tipo == c_compa.Tipo {
					if c_pivote.Valor == c_compa.Valor {
						if actual.L_instr_match.Len() != 0 {
							for i := 0; i < actual.L_instr_match.Len(); i++ {
								instr_actual := actual.L_instr_match.GetValue(i).(interfaces.Instruccion)
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
						break
					}
				} else {
					t := time.Now()
					reportes.Nerror("Error de tipos en las comparaciones de la Sentencia Match", ent.Nombre_Entorno, strconv.Itoa(mat.Linea), strconv.Itoa(mat.Columna), t.Format("2006-01-02 15:04:05"))
				}
			} else {
				if actual.L_instr_match.Len() != 0 {
					for i := 0; i < actual.L_instr_match.Len(); i++ {
						instr_actual := actual.L_instr_match.GetValue(i).(interfaces.Instruccion)
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
				break
			}
		} else {
			if actual.Lista_condiciones.Len() != 0 {
				for i := 0; i < actual.Lista_condiciones.Len(); i++ {
					nuevo_actual := actual.Lista_condiciones.GetValue(i).(interfaces.Expresion)
					nuevo := nuevo_actual.Ejecutar_Expresion(ent)

					if c_pivote.Tipo == nuevo.Tipo {
						if c_pivote.Valor == nuevo.Valor {
							if actual.L_instr_match.Len() != 0 {
								for i := 0; i < actual.L_instr_match.Len(); i++ {
									instr_actual := actual.L_instr_match.GetValue(i).(interfaces.Instruccion)
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
							break
						}
					} else {
						t := time.Now()
						reportes.Nerror("Error de tipos en las comparaciones de la Sentencia Match", ent.Nombre_Entorno, strconv.Itoa(mat.Linea), strconv.Itoa(mat.Columna), t.Format("2006-01-02 15:04:05"))
					}
				}
			}
		}
	}

	return simbolos.Valor{Tipo: simbolos.NULL, Valor: -1}
}

func (mat SenteMatch) Compilar_Expresion(ent *entorno.Entorno, gen *generador.Generador_C3D) simbolos.ValoresC3D {
	return simbolos.ValoresC3D{Valor: "0", EsTemporal: false, Tipo: simbolos.INTEGER, Label_verdadera: "", Label_false: ""}

}
