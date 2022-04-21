package impresion

import (
	entorno "OLC2-PROYECTO2/COMPILADOR/ENTORNO"
	generador "OLC2-PROYECTO2/COMPILADOR/GENERADOR"
	interfaces "OLC2-PROYECTO2/COMPILADOR/INTERFACES"
	simbolos "OLC2-PROYECTO2/COMPILADOR/SIMBOLOS"
	"fmt"
	"strconv"
)

type Imprimir struct {
	Exp     interfaces.Expresion
	Linea   int
	Columna int
}

func Nimprimir(e interfaces.Expresion, l int, c int) Imprimir {
	print := Imprimir{Exp: e, Linea: l, Columna: c}
	return print
}

func (print Imprimir) Ejecutar_Instruccion(ent *entorno.Entorno, ent2 *entorno.Entorno) interface{} {
	var primera_exp simbolos.Valor = print.Exp.Ejecutar_Expresion(ent)
	prexp := fmt.Sprintf("%v", primera_exp.Valor)
	fmt.Println(prexp)
	todito := simbolos.SimboloTodo{Tipo: "print", Nombre: prexp, Pos: ent2.Posicion}
	ent2.GuardaLTodo(todito)
	return 0
}

func (print Imprimir) Compilar_Instruccion(ent *entorno.Entorno, gen *generador.Generador_C3D) interface{} {
	var ex simbolos.ValoresC3D = print.Exp.Compilar_Expresion(ent, gen)
	if ex.Tipo == simbolos.YTEXTO {
		valtodi := ent.ListaTodo.GetValue(0).(simbolos.SimboloTodo)
		ent.EliminarLTodo()
		if valtodi.Tipo == "print" {
			impr := "/*-IMPRIMIR TEXTO-*/\n"
			temp1 := gen.Crear_temporal()
			temp2 := gen.Crear_temporal()
			impr += temp1 + " = HP;\n"
			for _, txt := range ex.Valor {
				f := int(txt)
				impr += "HEAP[int(HP)] = " + fmt.Sprintf("%v", f) + "; //LETRA-> " + string(txt) + "\n"
				impr += "HP = HP + 1;\n"
			}
			impr += "HEAP[int(HP)] = 10;\nHP = HP + 1;\n"
			impr += "HEAP[int(HP)] = -1;\nHP = HP + 1;\n"
			if valtodi.Pos != 0 {
				impr += temp2 + " = SP +" + strconv.Itoa(valtodi.Pos) + "; //POSICION CADENA EN STACK\n"
				impr += "STACK[int(" + temp2 + ")] = " + temp1 + ";\n"
				impr += "SP = SP + " + strconv.Itoa(valtodi.Pos) + ";\nprint();\nSP = SP - " + strconv.Itoa(valtodi.Pos) + ";\n"
			} else {
				impr += temp2 + " = SP +" + strconv.Itoa(valtodi.Pos) + "; //POSICION CADENA EN STACK\n"
				impr += "STACK[int(" + temp2 + ")] = " + temp1 + ";\n"
				impr += "SP = SP + " + strconv.Itoa(valtodi.Pos) + ";\nprint();\nSP = SP - " + strconv.Itoa(valtodi.Pos) + ";\n"
			}
			impr += "printf(\"%c\", (char)10);\n"
			gen.Agregar_Logica(impr)
		}
	} else if ex.EsTemporal {
		imprs := "/*-IMPRIMIR TEMPORAL-*/\n"

		if ex.Tipo == simbolos.INTEGER {
			imprs += "printf(\"%d\",(int)" + ex.Valor + ");\n"
		} else if ex.Tipo == simbolos.FLOAT {
			imprs += "printf(\"%f\"," + ex.Valor + ");\n"
		} else {
			imprs += "printf(\"%f\"," + ex.Valor + ");\n"
		}
		imprs += "printf(\"%c\",10);"
		gen.Agregar_Logica(imprs)
	} else if ex.Tipo == simbolos.FLOAT {
		imprs := "/*-IMPRIMIR DOUBLE-*/\n"
		imprs += "printf(\"%f\"," + ex.Valor + ");\n"
		imprs += "printf(\"%c\",10);"
		gen.Agregar_Logica(imprs)
	} else if ex.Tipo == simbolos.INTEGER {
		imprs := "/*-IMPRIMIR ENTERO-*/\n"
		imprs += "printf(\"%d\",(int)" + ex.Valor + ");\n"
		imprs += "printf(\"%c\",10);"
		gen.Agregar_Logica(imprs)
	}
	return 0
}
