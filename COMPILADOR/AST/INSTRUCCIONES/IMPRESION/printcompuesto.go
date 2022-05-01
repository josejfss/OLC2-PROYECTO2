package impresion

import (
	entorno "OLC2-PROYECTO2/COMPILADOR/ENTORNO"
	generador "OLC2-PROYECTO2/COMPILADOR/GENERADOR"
	interfaces "OLC2-PROYECTO2/COMPILADOR/INTERFACES"
	simbolos "OLC2-PROYECTO2/COMPILADOR/SIMBOLOS"
	"fmt"
	"strconv"
	"strings"

	"github.com/colegno/arraylist"
)

type Impres struct {
	Exp      interfaces.Expresion
	ListaExp *arraylist.List
	Linea    int
	Columna  int
}

func Nimpres(e interfaces.Expresion, le *arraylist.List, l int, c int) Impres {
	imp := Impres{Exp: e, ListaExp: le, Linea: l, Columna: c}
	return imp
}

func (imp Impres) Ejecutar_Instruccion(ent *entorno.Entorno, ent2 *entorno.Entorno) interface{} {
	var primera_exp simbolos.Valor = imp.Exp.Ejecutar_Expresion(ent)
	prexp := fmt.Sprintf("%v", primera_exp.Valor)
	var concatenacion string = ""

	todito := simbolos.SimboloTodo{Tipo: "print", Nombre: concatenacion, Pos: ent2.Posicion}
	ent2.GuardaLTodo(todito)

	if imp.ListaExp.Len() != 0 {
		if strings.Contains(prexp, "{}") {
			conteo := strings.Count(prexp, "{}")
			if conteo == imp.ListaExp.Len() {
				separador := strings.Split(prexp, "{}")
				for i := 0; i < len(separador); i++ {

					if i != imp.ListaExp.Len() {
						ex := imp.ListaExp.GetValue(i).(interfaces.Expresion)
						var ss simbolos.Valor = ex.Ejecutar_Expresion(ent)
						despues := fmt.Sprintf("%v", ss.Valor)
						antes := separador[i] + despues
						concatenacion += antes
					} else {
						concatenacion += separador[i]
					}
				}
			}
		} else if strings.Contains(prexp, "{:?}") {
			conteo := strings.Count(prexp, "{:?}")
			if conteo == imp.ListaExp.Len() {
				separador := strings.Split(prexp, "{:?}")
				for i := 0; i < len(separador); i++ {
					if i != imp.ListaExp.Len() {
						ex := imp.ListaExp.GetValue(i).(interfaces.Expresion)
						var ss simbolos.Valor = ex.Ejecutar_Expresion(ent)
						simbs := ss.Valor.(simbolos.Simbolo_ArreVect)
						despues := ""
						for i := 0; i < simbs.ValorLista.Len(); i++ {
							act := simbs.ValorLista.GetValue(i).(simbolos.Valor)
							despues += fmt.Sprintf("%v", act.Valor)
							if i != simbs.ValorLista.Len()-1 {
								despues += ","
							}
						}
						concatenacion += separador[i] + despues
					} else {
						concatenacion += separador[i]
					}
				}
			}
		}
	}
	fmt.Println(concatenacion)
	return 0
}

func (imp Impres) Compilar_Instruccion(ent *entorno.Entorno, gen *generador.Generador_C3D) interface{} {
	var primera_exp simbolos.Valor = imp.Exp.Ejecutar_Expresion(ent)
	prexp := fmt.Sprintf("%v", primera_exp.Valor)
	valtodi := ent.ListaTodo.GetValue(0).(simbolos.SimboloTodo)
	ent.EliminarLTodo()

	impr := "/*-IMPRIMIR TEXTO-*/\n"
	temp1 := gen.Crear_temporal()
	temp2 := gen.Crear_temporal()
	impr += temp1 + " = HP;\n"

	if valtodi.Tipo == "print" {
		if imp.ListaExp.Len() != 0 {
			if strings.Contains(prexp, "{}") {
				conteo := strings.Count(prexp, "{}")
				if conteo == imp.ListaExp.Len() {
					separador := strings.Split(prexp, "{}")
					for i := 0; i < len(separador); i++ {

						if i != imp.ListaExp.Len() {
							ex := imp.ListaExp.GetValue(i).(interfaces.Expresion)
							var ss simbolos.ValoresC3D = ex.Compilar_Expresion(ent, gen)

							//LA PRIMERA VEZ QUE ESTA EN EL FOR
							if i == 0 {
								//SE GUARDA EN EL HEAP EL TXT QUE SE VA A IMPRIMIR
								for _, txt := range separador[i] {
									f := int(txt)
									impr += "HEAP[(int)HP] = " + fmt.Sprintf("%v", f) + "; //LETRA-> " + string(txt) + "\n"
									impr += "HP = HP + 1;\n"
								}
								//SE GUARDA EL HEAP EL VALOR PARA INDICAR EL FIN DE LA CADENA
								impr += "HEAP[(int)HP] = -1;\nHP = HP + 1;\n"
								//SE GUARDA LA REFERENCIA DEL HEAP EN EL STACK
								impr += temp2 + " = SP + " + strconv.Itoa(valtodi.Pos) + "; //POSICION CADENA EN STACK\n"
								//SE GUARDA EN EL STACK LA CADENA DEL HEAP
								impr += "STACK[(int)" + temp2 + "] = " + temp1 + ";\n"
								//SE CAMBIA DE ENTORNO - SE LLAMA EL METODO PRINT - SE REGRESE AL OTRO ENTORNO
								impr += "SP = SP + " + strconv.Itoa(valtodi.Pos) + ";\nprint();\nSP = SP - " + strconv.Itoa(valtodi.Pos) + ";\n"
								//SE IMPRIME EL VALOR DE LA EXPRESION
								if ss.Tipo == simbolos.INTEGER {
									impr += "printf(\"%d\",(int)" + ss.Valor + ");\n"
								} else if ss.Tipo == simbolos.FLOAT {
									impr += "printf(\"%f\"," + ss.Valor + ");\n"
								} else if ss.Tipo == simbolos.STRUCT {
									impr += ss.Valor
								} else if ss.EsTemporal {
									temp_conca := gen.Crear_temporal()
									etiqueta_entrada := gen.Crear_label()
									gen.Eliminar_label(etiqueta_entrada)
									etiqueta_salida := gen.Crear_label()
									gen.Eliminar_label(etiqueta_salida)
									tempo_valor := ss.Valor
									impr += etiqueta_entrada + ":\n"
									impr += temp_conca + " = HEAP[(int)" + tempo_valor + "];\n"
									impr += "if ( " + temp_conca + " == -1) goto " + etiqueta_salida + ";\n"
									impr += "printf(\"%c\", (char)" + temp_conca + ");\n"
									impr += tempo_valor + " = " + tempo_valor + " + 1;\ngoto " + etiqueta_entrada + ";\n"
									impr += etiqueta_salida + ":\n"
								} else {
									//SE GUARDA EN EL HEAP EL TXT QUE SE VA A IMPRIMIR
									//SE CAMBIA EL EL VALOR DEL HEAP EN EL TEMPORAL
									impr += temp1 + " = HP;\n"
									for _, txt := range ss.Valor {
										f := int(txt)
										impr += "HEAP[(int)HP] = " + fmt.Sprintf("%v", f) + "; //LETRA-> " + string(txt) + "\n"
										impr += "HP = HP + 1;\n"
									}
									//SE GUARDA EL HEAP EL VALOR PARA INDICAR EL FIN DE LA CADENA
									impr += "HEAP[(int)HP] = -1;\nHP = HP + 1;\n"
									//SE GUARDA LA REFERENCIA DEL HEAP EN EL STACK
									impr += temp2 + " = SP + " + strconv.Itoa(valtodi.Pos) + "; //POSICION CADENA EN STACK\n"
									//SE GUARDA EN EL STACK LA CADENA DEL HEAP
									impr += "STACK[(int)" + temp2 + "] = " + temp1 + ";\n"
									//SE CAMBIA DE ENTORNO - SE LLAMA EL METODO PRINT - SE REGRESE AL OTRO ENTORNO
									impr += "SP = SP + " + strconv.Itoa(valtodi.Pos) + ";\nprint();\nSP = SP - " + strconv.Itoa(valtodi.Pos) + ";\n"
								}
								//impr += "printf(\"%c\", (char)10);\n"
							} else {
								//SE IMPRIME LOS VALORES DESPUES DE LAS LLAVES
								//SE CAMBIA EL EL VALOR DEL HEAP EN EL TEMPORAL
								impr += temp1 + " = HP;\n"
								//SE GUARDA EN EL TXT EN EL HEAP
								for _, txt := range separador[i] {
									f := int(txt)
									impr += "HEAP[(int)HP] = " + fmt.Sprintf("%v", f) + "; //LETRA-> " + string(txt) + "\n"
									impr += "HP = HP + 1;\n"
								}
								//SE GUARDA EL HEAP EL VALOR PARA INDICAR EL FIN DE LA CADENA
								impr += "HEAP[(int)HP] = -1;\nHP = HP + 1;\n"
								//SE GUARDA LA REFERENCIA DEL HEAP EN EL STACK
								impr += temp2 + " = SP + " + strconv.Itoa(valtodi.Pos) + "; //POSICION CADENA EN STACK\n"
								//SE GUARDA EN EL STACK LA CADENA DEL HEAP
								impr += "STACK[(int)" + temp2 + "] = " + temp1 + ";\n"
								//SE CAMBIA DE ENTORNO - SE LLAMA EL METODO PRINT - SE REGRESA AL OTRO ENTORNO
								impr += "SP = SP + " + strconv.Itoa(valtodi.Pos) + ";\nprint();\nSP = SP - " + strconv.Itoa(valtodi.Pos) + ";\n"
								//SE IMPRIME EL VALOR DE LA EXPRESION
								if ss.Tipo == simbolos.INTEGER {
									impr += "printf(\"%d\",(int)" + ss.Valor + ");\n"
								} else if ss.Tipo == simbolos.FLOAT {
									impr += "printf(\"%f\"," + ss.Valor + ");\n"
								} else if ss.EsTemporal {
									temp_conca := gen.Crear_temporal()
									etiqueta_entrada := gen.Crear_label()
									gen.Eliminar_label(etiqueta_entrada)
									etiqueta_salida := gen.Crear_label()
									gen.Eliminar_label(etiqueta_salida)
									tempo_valor := ss.Valor
									impr += etiqueta_entrada + ":\n"
									impr += temp_conca + " = HEAP[(int)" + tempo_valor + "];\n"
									impr += "if ( " + temp_conca + " == -1) goto " + etiqueta_salida + ";\n"
									impr += "printf(\"%c\", (char)" + temp_conca + ");\n"
									impr += tempo_valor + " = " + tempo_valor + " + 1;\ngoto " + etiqueta_entrada + ";\n"
									impr += etiqueta_salida + ":\n"
								} else {
									//SE GUARDA EN EL HEAP EL TXT QUE SE VA A IMPRIMIR
									//SE CAMBIA EL EL VALOR DEL HEAP EN EL TEMPORAL
									impr += temp1 + " = HP;\n"
									for _, txt := range ss.Valor {
										f := int(txt)
										impr += "HEAP[(int)HP] = " + fmt.Sprintf("%v", f) + "; //LETRA-> " + string(txt) + "\n"
										impr += "HP = HP + 1;\n"
									}
									//SE GUARDA EL HEAP EL VALOR PARA INDICAR EL FIN DE LA CADENA
									impr += "HEAP[(int)HP] = -1;\nHP = HP + 1;\n"
									//SE GUARDA LA REFERENCIA DEL HEAP EN EL STACK
									impr += temp2 + " = SP + " + strconv.Itoa(valtodi.Pos) + "; //POSICION CADENA EN STACK\n"
									//SE GUARDA EN EL STACK LA CADENA DEL HEAP
									impr += "STACK[(int)" + temp2 + "] = " + temp1 + ";\n"
									//SE CAMBIA DE ENTORNO - SE LLAMA EL METODO PRINT - SE REGRESE AL OTRO ENTORNO
									impr += "SP = SP + " + strconv.Itoa(valtodi.Pos) + ";\nprint();\nSP = SP - " + strconv.Itoa(valtodi.Pos) + ";\n"
								}
								//impr += "printf(\"%c\", (char)10);\n"
							}

						} else {
							//concatenacion += separador[i]
							if separador[i] != "" {
								//SE GUARDA LA REFERENCIA DEL HEAP
								impr += temp1 + " = HP;\n"
								//SE GUARDA TXT EN EL HEAP
								for _, txt := range separador[i] {
									f := int(txt)
									impr += "HEAP[(int)HP] = " + fmt.Sprintf("%v", f) + "; //LETRA-> " + string(txt) + "\n"
									impr += "HP = HP + 1;\n"
								}
								//SE GUARDA EN EL HEAP EL VALOR PARA INDICAR EL FIN DE LA CADENA
								impr += "HEAP[(int)HP] = -1;\nHP = HP + 1;\n"
								//SE GUARDA LA REFERENCIA DEL HEAP EN EL STACK
								impr += temp2 + " = SP + " + strconv.Itoa(valtodi.Pos) + "; //POSICION CADENA EN STACK\n"
								//SE GUARDA EN EL STACK LA CADENA DEL HEAP
								impr += "STACK[(int)" + temp2 + "] = " + temp1 + ";\n"
								//SE CAMBIA DE ENTORNO - SE LLAMA EL METODO PRINT - SE REGRESA AL OTRO ENTORNO
								impr += "SP = SP + " + strconv.Itoa(valtodi.Pos) + ";\nprint();\nSP = SP - " + strconv.Itoa(valtodi.Pos) + ";\n"
							}
						}
					}
				}
			} else if strings.Contains(prexp, "{:?}") {
				conteo := strings.Count(prexp, "{:?}")
				if conteo == imp.ListaExp.Len() {
					separador := strings.Split(prexp, "{:?}")
					for i := 0; i < len(separador); i++ {

						if i != imp.ListaExp.Len() {
							ex := imp.ListaExp.GetValue(i).(interfaces.Expresion)
							var ss simbolos.ValoresC3D = ex.Compilar_Expresion(ent, gen)

							//LA PRIMERA VEZ QUE ESTA EN EL FOR
							if i == 0 {
								//SE GUARDA EN EL HEAP EL TXT QUE SE VA A IMPRIMIR
								for _, txt := range separador[i] {
									f := int(txt)
									impr += "HEAP[(int)HP] = " + fmt.Sprintf("%v", f) + "; //LETRA-> " + string(txt) + "\n"
									impr += "HP = HP + 1;\n"
								}
								//SE GUARDA EL HEAP EL VALOR PARA INDICAR EL FIN DE LA CADENA
								impr += "HEAP[(int)HP] = -1;\nHP = HP + 1;\n"
								//SE GUARDA LA REFERENCIA DEL HEAP EN EL STACK
								impr += temp2 + " = SP + " + strconv.Itoa(valtodi.Pos) + "; //POSICION CADENA EN STACK\n"
								//SE GUARDA EN EL STACK LA CADENA DEL HEAP
								impr += "STACK[(int)" + temp2 + "] = " + temp1 + ";\n"
								//SE CAMBIA DE ENTORNO - SE LLAMA EL METODO PRINT - SE REGRESE AL OTRO ENTORNO
								impr += "SP = SP + " + strconv.Itoa(valtodi.Pos) + ";\nprint();\nSP = SP - " + strconv.Itoa(valtodi.Pos) + ";\n"
								//SE IMPRIME EL VALOR DE LA EXPRESION
								if ss.EsTemporal {
									if ss.Tipo == simbolos.INTEGER {
										temp_conca := gen.Crear_temporal()
										etiqueta_entrada := gen.Crear_label()
										gen.Eliminar_label(etiqueta_entrada)
										etiqueta_salida := gen.Crear_label()
										gen.Eliminar_label(etiqueta_salida)
										tempo_valor := ss.Valor
										impr += etiqueta_entrada + ":\n"
										impr += temp_conca + " = HEAP[(int)" + tempo_valor + "];\n"
										impr += "if ( " + temp_conca + " == -2) goto " + etiqueta_salida + ";\n"
										impr += "printf(\"%d\",(int)" + temp_conca + ");\n"
										impr += "printf(\"%c\", (char) 32);\n"
										impr += tempo_valor + " = " + tempo_valor + " + 1;\ngoto " + etiqueta_entrada + ";\n"
										impr += etiqueta_salida + ":\n"
									} else if ss.Tipo == simbolos.FLOAT {
										fmt.Println("ytexto")
									} else if ss.Tipo == simbolos.TEXTO {
										fmt.Println("ytexto")
									} else if ss.Tipo == simbolos.YTEXTO {
										temp_conca := gen.Crear_temporal()
										etiqueta_entrada := gen.Crear_label()
										gen.Eliminar_label(etiqueta_entrada)
										etiqueta_salida := gen.Crear_label()
										gen.Eliminar_label(etiqueta_salida)
										tempo_valor := ss.Valor

										tempo_conca1 := gen.Crear_temporal()
										etiqueta_entrada1 := gen.Crear_label()
										gen.Eliminar_label(etiqueta_entrada1)
										etiqueta_salida1 := gen.Crear_label()
										gen.Eliminar_label(etiqueta_salida1)
										tempo_valor1 := gen.Crear_temporal()

										impr += etiqueta_entrada + ":\n"
										impr += temp_conca + " = HEAP[(int)" + tempo_valor + "];\n"
										impr += "if ( " + temp_conca + " == -2) goto " + etiqueta_salida + ";\n"
										impr += tempo_valor1 + " = " + temp_conca + ";\n"
										impr += etiqueta_entrada1 + ":\n"
										impr += tempo_conca1 + " = HEAP[(int)" + tempo_valor1 + "];\n"
										impr += "if ( " + tempo_conca1 + " == -1) goto " + etiqueta_salida1 + ";\n"
										impr += "printf(\"%c\", (char)" + tempo_conca1 + ");\n"
										impr += tempo_valor1 + " = " + tempo_valor1 + " + 1;\ngoto " + etiqueta_entrada1 + ";\n"
										impr += etiqueta_salida1 + ":\n"
										impr += "printf(\"%c\", (char) 32);\n"
										impr += tempo_valor + " = " + tempo_valor + " + 1;\ngoto " + etiqueta_entrada + ";\n"
										impr += etiqueta_salida + ":\n"
									}

								} else {
									//SE GUARDA EN EL HEAP EL TXT QUE SE VA A IMPRIMIR
									//SE CAMBIA EL EL VALOR DEL HEAP EN EL TEMPORAL
									impr += temp1 + " = HP;\n"
									for _, txt := range ss.Valor {
										f := int(txt)
										impr += "HEAP[(int)HP] = " + fmt.Sprintf("%v", f) + "; //LETRA-> " + string(txt) + "\n"
										impr += "HP = HP + 1;\n"
									}
									//SE GUARDA EL HEAP EL VALOR PARA INDICAR EL FIN DE LA CADENA
									impr += "HEAP[(int)HP] = -1;\nHP = HP + 1;\n"
									//SE GUARDA LA REFERENCIA DEL HEAP EN EL STACK
									impr += temp2 + " = SP + " + strconv.Itoa(valtodi.Pos) + "; //POSICION CADENA EN STACK\n"
									//SE GUARDA EN EL STACK LA CADENA DEL HEAP
									impr += "STACK[(int)" + temp2 + "] = " + temp1 + ";\n"
									//SE CAMBIA DE ENTORNO - SE LLAMA EL METODO PRINT - SE REGRESE AL OTRO ENTORNO
									impr += "SP = SP + " + strconv.Itoa(valtodi.Pos) + ";\nprint();\nSP = SP - " + strconv.Itoa(valtodi.Pos) + ";\n"
								}
								//impr += "printf(\"%c\", (char)10);\n"
							} else {
								//SE IMPRIME LOS VALORES DESPUES DE LAS LLAVES
								//SE CAMBIA EL EL VALOR DEL HEAP EN EL TEMPORAL
								impr += temp1 + " = HP;\n"
								//SE GUARDA EN EL TXT EN EL HEAP
								for _, txt := range separador[i] {
									f := int(txt)
									impr += "HEAP[(int)HP] = " + fmt.Sprintf("%v", f) + "; //LETRA-> " + string(txt) + "\n"
									impr += "HP = HP + 1;\n"
								}
								//SE GUARDA EL HEAP EL VALOR PARA INDICAR EL FIN DE LA CADENA
								impr += "HEAP[(int)HP] = -1;\nHP = HP + 1;\n"
								//SE GUARDA LA REFERENCIA DEL HEAP EN EL STACK
								impr += temp2 + " = SP + " + strconv.Itoa(valtodi.Pos) + "; //POSICION CADENA EN STACK\n"
								//SE GUARDA EN EL STACK LA CADENA DEL HEAP
								impr += "STACK[(int)" + temp2 + "] = " + temp1 + ";\n"
								//SE CAMBIA DE ENTORNO - SE LLAMA EL METODO PRINT - SE REGRESA AL OTRO ENTORNO
								impr += "SP = SP + " + strconv.Itoa(valtodi.Pos) + ";\nprint();\nSP = SP - " + strconv.Itoa(valtodi.Pos) + ";\n"
								//SE IMPRIME EL VALOR DE LA EXPRESION
								if ss.EsTemporal {
									if ss.Tipo == simbolos.INTEGER {
										temp_conca := gen.Crear_temporal()
										etiqueta_entrada := gen.Crear_label()
										gen.Eliminar_label(etiqueta_entrada)
										etiqueta_salida := gen.Crear_label()
										gen.Eliminar_label(etiqueta_salida)
										tempo_valor := ss.Valor
										impr += etiqueta_entrada + ":\n"
										impr += temp_conca + " = HEAP[(int)" + tempo_valor + "];\n"
										impr += "if ( " + temp_conca + " == -2) goto " + etiqueta_salida + ";\n"
										impr += "printf(\"%d\",(int)" + temp_conca + ");\n"
										impr += "printf(\"%c\", (char) 32);\n"
										impr += tempo_valor + " = " + tempo_valor + " + 1;\ngoto " + etiqueta_entrada + ";\n"
										impr += etiqueta_salida + ":\n"
									} else if ss.Tipo == simbolos.FLOAT {
										fmt.Println("ytexto")
									} else if ss.Tipo == simbolos.TEXTO {
										fmt.Println("ytexto")
									} else if ss.Tipo == simbolos.YTEXTO {
										fmt.Println("ytexto")
									}

								} else {
									//SE GUARDA EN EL HEAP EL TXT QUE SE VA A IMPRIMIR
									//SE CAMBIA EL EL VALOR DEL HEAP EN EL TEMPORAL
									impr += temp1 + " = HP;\n"
									for _, txt := range ss.Valor {
										f := int(txt)
										impr += "HEAP[(int)HP] = " + fmt.Sprintf("%v", f) + "; //LETRA-> " + string(txt) + "\n"
										impr += "HP = HP + 1;\n"
									}
									//SE GUARDA EL HEAP EL VALOR PARA INDICAR EL FIN DE LA CADENA
									impr += "HEAP[(int)HP] = -1;\nHP = HP + 1;\n"
									//SE GUARDA LA REFERENCIA DEL HEAP EN EL STACK
									impr += temp2 + " = SP + " + strconv.Itoa(valtodi.Pos) + "; //POSICION CADENA EN STACK\n"
									//SE GUARDA EN EL STACK LA CADENA DEL HEAP
									impr += "STACK[(int)" + temp2 + "] = " + temp1 + ";\n"
									//SE CAMBIA DE ENTORNO - SE LLAMA EL METODO PRINT - SE REGRESE AL OTRO ENTORNO
									impr += "SP = SP + " + strconv.Itoa(valtodi.Pos) + ";\nprint();\nSP = SP - " + strconv.Itoa(valtodi.Pos) + ";\n"
								}
								//impr += "printf(\"%c\", (char)10);\n"
							}

						} else {
							//concatenacion += separador[i]
							if separador[i] != "" {
								//SE GUARDA LA REFERENCIA DEL HEAP
								impr += temp1 + " = HP;\n"
								//SE GUARDA TXT EN EL HEAP
								for _, txt := range separador[i] {
									f := int(txt)
									impr += "HEAP[(int)HP] = " + fmt.Sprintf("%v", f) + "; //LETRA-> " + string(txt) + "\n"
									impr += "HP = HP + 1;\n"
								}
								//SE GUARDA EN EL HEAP EL VALOR PARA INDICAR EL FIN DE LA CADENA
								impr += "HEAP[(int)HP] = -1;\nHP = HP + 1;\n"
								//SE GUARDA LA REFERENCIA DEL HEAP EN EL STACK
								impr += temp2 + " = SP + " + strconv.Itoa(valtodi.Pos) + "; //POSICION CADENA EN STACK\n"
								//SE GUARDA EN EL STACK LA CADENA DEL HEAP
								impr += "STACK[(int)" + temp2 + "] = " + temp1 + ";\n"
								//SE CAMBIA DE ENTORNO - SE LLAMA EL METODO PRINT - SE REGRESA AL OTRO ENTORNO
								impr += "SP = SP + " + strconv.Itoa(valtodi.Pos) + ";\nprint();\nSP = SP - " + strconv.Itoa(valtodi.Pos) + ";\n"
							}
						}
					}
				}
			}
		}
	}
	impr += "printf(\"%c\", (char)10);\n"
	gen.Agregar_Logica(impr)
	//ent.EliminarLTodo()
	return 0
}
