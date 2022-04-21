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
							//despues := fmt.Sprintf("%v", ss.Valor)
							//antes := separador[i] + despues
							//concatenacion += antes

							//LA PRIMERA VEZ QUE ESTA EN EL FOR
							if i == 0 {
								//SE GUARDA EN EL HEAP EL TXT QUE SE VA A IMPRIMIR
								for _, txt := range separador[i] {
									f := int(txt)
									impr += "HEAP[int(HP)] = " + fmt.Sprintf("%v", f) + "; //LETRA-> " + string(txt) + "\n"
									impr += "HP = HP + 1;\n"
								}
								//SE GUARDA EL HEAP EL VALOR PARA INDICAR EL FIN DE LA CADENA
								impr += "HEAP[int(HP)] = -1;\nHP = HP + 1;\n"
								//SE GUARDA LA REFERENCIA DEL HEAP EN EL STACK
								impr += temp2 + " = SP + " + strconv.Itoa(valtodi.Pos-1) + "; //POSICION CADENA EN STACK\n"
								//SE GUARDA EN EL STACK LA CADENA DEL HEAP
								impr += "STACK[int(" + temp2 + ")] = " + temp1 + ";\n"
								//SE CAMBIA DE ENTORNO - SE LLAMA EL METODO PRINT - SE REGRESE AL OTRO ENTORNO
								impr += "SP = SP + " + strconv.Itoa(valtodi.Pos-1) + ";\nprint();\nSP = SP - " + strconv.Itoa(valtodi.Pos-1) + ";\n"
								//SE IMPRIME EL VALOR DE LA EXPRESION
								if ss.Tipo == simbolos.INTEGER {
									impr += "printf(\"%d\",(int)" + ss.Valor + ");\n"
								} else if ss.Tipo == simbolos.FLOAT {
									impr += "printf(\"%f\"," + ss.Valor + ");\n"
								} else {
									impr += "printf(\"%f\"," + ss.Valor + ");\n"
								}
							} else {
								//SE IMPRIME LOS VALORES DESPUES DE LAS LLAVES
								//SE CAMBIA EL EL VALOR DEL HEAP EN EL TEMPORAL
								impr += temp1 + " = HP;\n"
								//SE GUARDA EN EL TXT EN EL HEAP
								for _, txt := range separador[i] {
									f := int(txt)
									impr += "HEAP[int(HP)] = " + fmt.Sprintf("%v", f) + "; //LETRA-> " + string(txt) + "\n"
									impr += "HP = HP + 1;\n"
								}
								//SE GUARDA EL HEAP EL VALOR PARA INDICAR EL FIN DE LA CADENA
								impr += "HEAP[int(HP)] = -1;\nHP = HP + 1;\n"
								//SE GUARDA LA REFERENCIA DEL HEAP EN EL STACK
								impr += temp2 + " = SP + " + strconv.Itoa(valtodi.Pos-1) + "; //POSICION CADENA EN STACK\n"
								//SE GUARDA EN EL STACK LA CADENA DEL HEAP
								impr += "STACK[int(" + temp2 + ")] = " + temp1 + ";\n"
								//SE CAMBIA DE ENTORNO - SE LLAMA EL METODO PRINT - SE REGRESA AL OTRO ENTORNO
								impr += "SP = SP + " + strconv.Itoa(valtodi.Pos-1) + ";\nprint();\nSP = SP - " + strconv.Itoa(valtodi.Pos-1) + ";\n"
								//SE IMPRIME EL VALOR DE LA EXPRESION
								if ss.Tipo == simbolos.INTEGER {
									impr += "printf(\"%d\",(int)" + ss.Valor + ");\n"
								} else if ss.Tipo == simbolos.FLOAT {
									impr += "printf(\"%f\"," + ss.Valor + ");\n"
								} else {
									impr += "printf(\"%f\"," + ss.Valor + ");\n"
								}
							}

						} else {
							//concatenacion += separador[i]
							if separador[i] != "" {
								//SE GUARDA LA REFERENCIA DEL HEAP
								impr += temp1 + " = HP;\n"
								//SE GUARDA TXT EN EL HEAP
								for _, txt := range separador[i] {
									f := int(txt)
									impr += "HEAP[int(HP)] = " + fmt.Sprintf("%v", f) + "; //LETRA-> " + string(txt) + "\n"
									impr += "HP = HP + 1;\n"
								}
								//SE GUARDA EN EL HEAP EL VALOR PARA INDICAR EL FIN DE LA CADENA
								impr += "HEAP[int(HP)] = -1;\nHP = HP + 1;\n"
								//SE GUARDA LA REFERENCIA DEL HEAP EN EL STACK
								impr += temp2 + " = SP + " + strconv.Itoa(valtodi.Pos) + "; //POSICION CADENA EN STACK\n"
								//SE GUARDA EN EL STACK LA CADENA DEL HEAP
								impr += "STACK[int(" + temp2 + ")] = " + temp1 + ";\n"
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
