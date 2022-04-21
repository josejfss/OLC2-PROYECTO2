package funciones

import (
	primitivo "OLC2-PROYECTO2/COMPILADOR/AST/EXPRESIONES/PRIMITIVO"
	declaracionvar "OLC2-PROYECTO2/COMPILADOR/AST/INSTRUCCIONES/DECLARACIONVAR"
	parametros "OLC2-PROYECTO2/COMPILADOR/AST/INSTRUCCIONES/FUNCIONES/PARAMETROS"
	tipos "OLC2-PROYECTO2/COMPILADOR/AST/INSTRUCCIONES/FUNCIONES/TIPOS"
	entorno "OLC2-PROYECTO2/COMPILADOR/ENTORNO"
	generador "OLC2-PROYECTO2/COMPILADOR/GENERADOR"
	interfaces "OLC2-PROYECTO2/COMPILADOR/INTERFACES"
	reportes "OLC2-PROYECTO2/COMPILADOR/REPORTES"
	simbolos "OLC2-PROYECTO2/COMPILADOR/SIMBOLOS"
	"fmt"
	"reflect"
	"strconv"
	"time"

	"github.com/colegno/arraylist"
)

type Funciones struct {
	Publico     bool
	NombreFunca string
	LParametros *arraylist.List
	TipoFun     tipos.TiposFunca
	Linstrucc   *arraylist.List
	Linea       int
	Columna     int
}

func Nfunciones(p bool, nf string, lp *arraylist.List, tfu tipos.TiposFunca, lins *arraylist.List, lin int, col int) Funciones {
	funcas := Funciones{Publico: p, NombreFunca: nf, LParametros: lp, TipoFun: tfu, Linstrucc: lins, Linea: lin, Columna: col}
	return funcas
}

func (fun Funciones) Ejecutar_Instruccion(ent *entorno.Entorno, ent2 *entorno.Entorno) interface{} {

	if !ent.Existe_Funciones(fun.NombreFunca) {
		ent.Guardar_Funciones(fun.Publico, fun.NombreFunca, fun.TipoFun.TiVa, "", fun.LParametros.Clone(), fun.Linstrucc.Clone(), fun.Linea, fun.Columna)
		ent.ListaTodo.Add(simbolos.SimboloTodo{Tipo: "funcion", Nombre: fun.NombreFunca})
		ent2.Guardar_Funciones(fun.Publico, fun.NombreFunca, fun.TipoFun.TiVa, "", fun.LParametros.Clone(), fun.Linstrucc.Clone(), fun.Linea, fun.Columna)
		ent2.ListaTodo.Add(simbolos.SimboloTodo{Tipo: "funcion", Nombre: fun.NombreFunca})
		conca := ""
		if fun.LParametros.Len() != 0 {
			for i := 0; i < fun.LParametros.Len(); i++ {
				parpivote := fun.LParametros.GetValue(i).(parametros.Param)
				conca += parpivote.Identificador
				if i != fun.LParametros.Len()-1 {
					conca += ","
				}
			}
		}
		if conca == "" {
			conca += "--"
		}
		reportes.ReporteSimbolos(fun.NombreFunca, "funcion--"+reportes.ReportObteniendoSimbolos(fun.TipoFun.TiVa), ent.Nombre_Entorno, conca, strconv.Itoa(fun.Linea), strconv.Itoa(fun.Columna))
	}
	entorno1 := entorno.Nuevo_Entorno("fn_"+fun.NombreFunca, ent)
	entorno2 := entorno.Nuevo_Entorno("fn_"+fun.NombreFunca, ent2)

	if fun.LParametros.Len() != 0 {
		for i := 0; i < fun.LParametros.Len(); i++ {
			parpivote := fun.LParametros.GetValue(i).(parametros.Param)
			if parpivote.ParametroT == "variable" {
				var declapar declaracionvar.Declaracion
				declapar.Identificador = parpivote.Identificador
				declapar.TipoDecla = parpivote.TipoVar
				vals := primitivo.Nuevo_Dato_Primitivo(Valor_Explicito(parpivote.TipoVar), parpivote.TipoVar)
				declapar.Valor_exp = vals
				declapar.Linea = parpivote.Linea
				declapar.Columna = parpivote.Columna
				declapar.Ejecutar_Instruccion(entorno1, entorno2)
				poselimnar := entorno2.ObtTamLTodo()
				entorno2.ListaTodo.RemoveAtIndex(poselimnar - 1)
			}
			// else if parpivote.ParametroT == "vector" {
			// 	var declavecpar declaracionvect.DeclaVector
			// 	declavecpar.Identificador = parpivote.Identificador
			// 	nuevadim := arraylist.New()
			// 	declavecpar.Dimension = nuevadim
			// } else if parpivote.ParametroT == "arreglo" {

			// }
		}
	}

	for i := 0; i < fun.Linstrucc.Len(); i++ {
		instr_actual := fun.Linstrucc.GetValue(i).(interfaces.Instruccion)
		if interfaces.TYPEOF(instr_actual) == "stransferencia.Romper" {
			t := time.Now()
			reportes.Nerror("No puede venir un break en una funcion.", entorno1.Nombre_Entorno, strconv.Itoa(fun.Linea), strconv.Itoa(fun.Columna), t.Format("2006-01-02 15:04:05"))
		} else if interfaces.TYPEOF(instr_actual) == "stransferencia.Continuar" {
			t := time.Now()
			reportes.Nerror("No puede venir un continue en una funcion.", entorno1.Nombre_Entorno, strconv.Itoa(fun.Linea), strconv.Itoa(fun.Columna), t.Format("2006-01-02 15:04:05"))
		} else if interfaces.TYPEOF(instr_actual) == "funciones.Llamada" {

		} else {
			valor_instr_actual := instr_actual.Ejecutar_Instruccion(entorno1, entorno2)
			if valor_instr_actual != 0 {
				if reflect.TypeOf(valor_instr_actual) != reflect.TypeOf(simbolos.Valor{}) {
					fmt.Println("error")
					t := time.Now()
					reportes.Nerror("Retorno no adecuado.", entorno1.Nombre_Entorno, strconv.Itoa(fun.Linea), strconv.Itoa(fun.Columna), t.Format("2006-01-02 15:04:05"))
					return simbolos.Valor{Tipo: simbolos.NULL, Valor: -1}
				}
				if fun.TipoFun.TiVa == valor_instr_actual.(simbolos.Valor).Tipo {
					valor_retorno := valor_instr_actual.(simbolos.Valor)
					simfunca := ent.Obtener_Funciones(fun.NombreFunca)
					simfunca.ValVectArre = valor_retorno
					ent.Guardar_Retorno(fun.NombreFunca, simfunca)
					return valor_retorno
				} else {
					fmt.Println("error")
					t := time.Now()
					reportes.Nerror("El tipo de la funcion no coincide en el entorno.", entorno1.Nombre_Entorno, strconv.Itoa(fun.Linea), strconv.Itoa(fun.Columna), t.Format("2006-01-02 15:04:05"))
					return simbolos.Valor{Tipo: simbolos.NULL, Valor: -1}
				}
			}
		}
	}
	simfunca := ent.Obtener_Funciones(fun.NombreFunca)
	simfunca.ValVectArre = 0
	ent.Guardar_Retorno(fun.NombreFunca, simfunca)
	entorno1.Posicion = entorno1.Posicion + 1
	entorno1.Tamaño = entorno1.Tamaño + 1
	ent.Guardar_Entorno(entorno1)

	ent2.Guardar_Retorno(fun.NombreFunca, simfunca)
	entorno2.Posicion = entorno2.Posicion + 1
	entorno2.Tamaño = entorno2.Tamaño + 1
	ent2.Guardar_Entorno(entorno2)
	return 0
}

func (fun Funciones) Compilar_Instruccion(ent *entorno.Entorno, gen *generador.Generador_C3D) interface{} {
	valtodi := ent.ObtFunca(fun.NombreFunca)
	if fun.NombreFunca != "main" {
		if valtodi.Nombre == fun.NombreFunca {
			if fun.TipoFun.TiVa == simbolos.NULL {
				if ent.Existe_Funciones(fun.NombreFunca) {
					ento := ent.Retornar_Entorno("fn_" + fun.NombreFunca)
					gen.Agregar_Comentario("------------EMPEZANDO METODO " + fun.NombreFunca + "------------")
					gen.Agregar_Logica("void " + fun.NombreFunca + "() {")

					// if fun.LParametros.Len() != 0 {
					// 	// gen.Agregar_Comentario("EMPEZANDO DECLARACION PARAMETROS")
					// 	// for i := 0; i < fun.LParametros.Len(); i++ {
					// 	// 	parpivote := fun.LParametros.GetValue(i).(parametros.Param)
					// 	// 	if parpivote.ParametroT == "variable" {
					// 	// 		var declapar declaracionvar.Declaracion
					// 	// 		declapar.Identificador = parpivote.Identificador
					// 	// 		declapar.TipoDecla = parpivote.TipoVar
					// 	// 		vals := primitivo.Nuevo_Dato_Primitivo(Valor_Explicito(parpivote.TipoVar), parpivote.TipoVar)
					// 	// 		declapar.Valor_exp = vals
					// 	// 		declapar.Linea = parpivote.Linea
					// 	// 		declapar.Columna = parpivote.Columna
					// 	// 		declapar.Compilar_Instruccion(ento, gen)
					// 	// 	}
					// 	// 	// else if parpivote.ParametroT == "vector" {
					// 	// 	// 	var declavecpar declaracionvect.DeclaVector
					// 	// 	// 	declavecpar.Identificador = parpivote.Identificador
					// 	// 	// 	nuevadim := arraylist.New()
					// 	// 	// 	declavecpar.Dimension = nuevadim
					// 	// 	// } else if parpivote.ParametroT == "arreglo" {

					// 	// 	// }
					// 	// }
					// 	// gen.Agregar_Comentario("FINALIZANDO DECLARACION PARAMETROS")
					// }

					for i := 0; i < fun.Linstrucc.Len(); i++ {
						instr := fun.Linstrucc.GetValue(i).(interfaces.Instruccion)
						instr.Compilar_Instruccion(ento, gen)
					}
					gen.Agregar_Logica("\nreturn;\n}")
					gen.Agregar_Comentario("------------FINALIZANDO METODO " + fun.NombreFunca + "------------")
				}
			} else {
				if ent.Existe_Funciones(fun.NombreFunca) {
					ento := ent.Retornar_Entorno("fn_" + fun.NombreFunca)
					gen.Agregar_Comentario("------------EMPEZANDO FUNCION " + fun.NombreFunca + "------------")
					gen.Agregar_Logica("func " + fun.NombreFunca + "() {")

					if fun.LParametros.Len() != 0 {
						gen.Agregar_Comentario("EMPEZANDO DECLARACION PARAMETROS")
						for i := 0; i < fun.LParametros.Len(); i++ {
							parpivote := fun.LParametros.GetValue(i).(parametros.Param)
							if parpivote.ParametroT == "variable" {
								var declapar declaracionvar.Declaracion
								declapar.Identificador = parpivote.Identificador
								declapar.TipoDecla = parpivote.TipoVar
								vals := primitivo.Nuevo_Dato_Primitivo(Valor_Explicito(parpivote.TipoVar), parpivote.TipoVar)
								declapar.Valor_exp = vals
								declapar.Linea = parpivote.Linea
								declapar.Columna = parpivote.Columna
								declapar.Compilar_Instruccion(ento, gen)
							}
							// else if parpivote.ParametroT == "vector" {
							// 	var declavecpar declaracionvect.DeclaVector
							// 	declavecpar.Identificador = parpivote.Identificador
							// 	nuevadim := arraylist.New()
							// 	declavecpar.Dimension = nuevadim
							// } else if parpivote.ParametroT == "arreglo" {

							// }
						}
						gen.Agregar_Comentario("FINALIZANDO DECLARACION PARAMETROS")
					}

					for i := 0; i < fun.Linstrucc.Len(); i++ {
						instr := fun.Linstrucc.GetValue(i).(interfaces.Instruccion)
						instr.Compilar_Instruccion(ento, gen)
					}
					gen.Agregar_Logica("}")
					gen.Agregar_Comentario("------------FINALIZANDO FUNCION " + fun.NombreFunca + "------------")
				}
			}
		}

	} else {
		if ent.Existe_Funciones(fun.NombreFunca) {
			if valtodi.Nombre == fun.NombreFunca {
				ento := ent.Retornar_Entorno("fn_" + fun.NombreFunca)
				// gen.Agregar_Comentario("------------EMPEZANDO FUNCION " + fun.NombreFunca + "------------")
				// gen.Agregar_Logica("int " + fun.NombreFunca + "() {")
				for i := 0; i < fun.Linstrucc.Len(); i++ {
					instr := fun.Linstrucc.GetValue(i).(interfaces.Instruccion)
					instr.Compilar_Instruccion(ento, gen)
				}
				// gen.Agregar_Logica("\n\nreturn 0; \n}")
				// gen.Agregar_Comentario("------------FINALIZANDO FUNCION " + fun.NombreFunca + "------------")
			}
		}
	}
	return 0
}

func Valor_Explicito(tipo simbolos.TipoExpresion) interface{} {
	if simbolos.INTEGER == tipo {
		return 0
	} else if simbolos.FLOAT == tipo {
		return 0.0
	} else if simbolos.YTEXTO == tipo {
		return ""
	} else if simbolos.TEXTO == tipo {
		return ""
	} else if simbolos.BOOLEAN == tipo {
		return false
	}
	return 0
}
