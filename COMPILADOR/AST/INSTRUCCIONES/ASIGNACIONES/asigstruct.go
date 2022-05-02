package asignaciones

import (
	entorno "OLC2-PROYECTO2/COMPILADOR/ENTORNO"
	generador "OLC2-PROYECTO2/COMPILADOR/GENERADOR"
	interfaces "OLC2-PROYECTO2/COMPILADOR/INTERFACES"
	simbolos "OLC2-PROYECTO2/COMPILADOR/SIMBOLOS"
	"fmt"
	"strconv"
)

type AsigStruct struct {
	Identificador string
	Nombre        string
	Valor         interfaces.Expresion
	Linea         int
	Columna       int
}

func Nasigstruct(nom string, nom2 string, val interfaces.Expresion, lin int, col int) AsigStruct {
	astruct := AsigStruct{Identificador: nom, Nombre: nom2, Valor: val, Linea: lin, Columna: col}
	return astruct
}

func (as AsigStruct) Ejecutar_Instruccion(ent *entorno.Entorno, ent2 *entorno.Entorno) interface{} {
	return 0
}

func (as AsigStruct) Compilar_Instruccion(ent *entorno.Entorno, gen *generador.Generador_C3D) interface{} {
	gen.Agregar_Comentario("INICIO ASIGNACION STRUCT MUTABLE")
	nuevo_valor := as.Valor.Compilar_Expresion(ent, gen)
	if ent.Existe_Variable(as.Identificador) {
		simb := ent.Obtener_Variable(as.Identificador)
		if simb.Mutable {
			if ent.Existe_Structs(simb.ValorVariable.(string)) {
				temp1 := gen.Crear_temporal()
				temp2 := gen.Crear_temporal()
				gen.Agregar_Logica(temp1 + " = SP + " + strconv.Itoa(simb.PosicionTabla) + ";")
				gen.Agregar_Logica(temp2 + " = STACK[(int)" + temp1 + "];")
				entstruct := ent.Entorno_Anterior.Retornar_Entorno("struct_" + simb.ValorVariable.(string))
				if entstruct.Existe_Variable(as.Nombre) {
					varstruct := entstruct.Obtener_Variable(as.Nombre)
					temp3 := gen.Crear_temporal()
					gen.Agregar_Logica(temp3 + " = " + temp2 + " + " + strconv.Itoa(varstruct.PosicionTabla) + ";")
					if nuevo_valor.Tipo == simbolos.INTEGER {
						gen.Agregar_Logica("HEAP[(int)" + temp3 + "] = " + nuevo_valor.Valor + ";")
					} else if nuevo_valor.Tipo == simbolos.FLOAT {
						gen.Agregar_Logica("HEAP[(int)" + temp3 + "] = " + nuevo_valor.Valor + ";")
					} else if nuevo_valor.Tipo == simbolos.TEXTO || nuevo_valor.Tipo == simbolos.YTEXTO {
						gen.Agregar_Logica("HEAP[(int)" + temp3 + "] = HP;")
						for _, txt := range nuevo_valor.Valor {
							f := int(txt)
							gen.Agregar_Logica("HEAP[(int)HP] = " + fmt.Sprintf("%v", f) + "; //LETRA-> " + string(txt) + ";")
							gen.Agregar_Logica("HP = HP + 1;")
						}
						gen.Agregar_Logica("HEAP[(int)HP] = -1;\nHP = HP + 1;")
					}
				}
			}
		}
	}
	gen.Agregar_Comentario("FIN ASIGNACION STRUCT MUTABLE")
	return 0
}
