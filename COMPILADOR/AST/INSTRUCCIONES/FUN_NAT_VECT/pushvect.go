package funnatvect

import (
	entorno "OLC2-PROYECTO2/COMPILADOR/ENTORNO"
	generador "OLC2-PROYECTO2/COMPILADOR/GENERADOR"
	interfaces "OLC2-PROYECTO2/COMPILADOR/INTERFACES"
	simbolos "OLC2-PROYECTO2/COMPILADOR/SIMBOLOS"
	"fmt"
	"strconv"
)

type PushVect struct {
	Identificador string
	Pushear       interfaces.Expresion
	Linea         int
	Columna       int
}

func Npushvect(nom string, push interfaces.Expresion, lin int, col int) PushVect {
	pusheo := PushVect{Identificador: nom, Pushear: push, Linea: lin, Columna: col}
	return pusheo
}

func (pushvect PushVect) Ejecutar_Instruccion(ent *entorno.Entorno, ent2 *entorno.Entorno) interface{} {
	// if ent.Existe_ArreVect(pushvect.Identificador) {
	// 	simbvect := ent.Obtener_ArreVect(pushvect.Identificador)
	// 	if simbvect.ValorLista.Len() < simbvect.EsArreVect {
	// 		pusheo := pushvect.Pushear.Ejecutar_Expresion(ent)
	// 		if simbvect.TipoVect == pusheo.Tipo {
	// 			simbvect.ValorLista.Add(pusheo.Valor)
	// 			simbvect.EsArreVect=simbvect.EsArreVect+1
	// 		} else {
	// 			t := time.Now()
	// 			reportes.Nerror("Los tipos no coinciden.", ent.Nombre_Entorno, strconv.Itoa(pushvect.Linea), strconv.Itoa(pushvect.Columna), t.Format("2006-01-02 15:04:05"))
	// 		}
	// 	}
	// } else {
	// 	t := time.Now()
	// 	reportes.Nerror("No existe el vector", ent.Nombre_Entorno, strconv.Itoa(pushvect.Linea), strconv.Itoa(pushvect.Columna), t.Format("2006-01-02 15:04:05"))
	// }
	return 0
}

func (pushvect PushVect) Compilar_Instruccion(ent *entorno.Entorno, gen *generador.Generador_C3D) interface{} {
	gen.Agregar_Comentario("INICIO PUSH VECTOR -- " + pushvect.Identificador)
	if ent.ExisteAcual_ArreVect(pushvect.Identificador) {
		simbvect := ent.Obtener_ArreVect(pushvect.Identificador)
		if simbvect.Mutable {
			if simbvect.ValorLista.Len() < simbvect.EsArreVect {
				pusheo := pushvect.Pushear.Compilar_Expresion(ent, gen)
				pusheo1 := pushvect.Pushear.Ejecutar_Expresion(ent)
				if simbvect.TipoVect == pusheo.Tipo {
					tempo1 := gen.Crear_temporal()
					posi := simbvect.PosicionTabla
					temp2 := gen.Crear_temporal()
					gen.Agregar_Logica(tempo1 + " = SP + " + strconv.Itoa(posi) + ";\t\t// POSICION: " + pushvect.Identificador)
					gen.Agregar_Logica(temp2 + " = STACK[int(" + tempo1 + ")];")
					temp3 := gen.Crear_temporal()
					gen.Agregar_Logica(temp3 + " = " + temp2 + " + " + strconv.Itoa(simbvect.ValorLista.Len()) + ";")
					gen.Agregar_Logica("HEAP[int(" + temp3 + ")] = " + fmt.Sprintf("%v", pusheo.Valor) + ";")

					//MODIFICACIONES EN EL VECTOR
					simbvect.ValorLista.Add(pusheo1)
					ent.Guardar_ArregloVector(simbvect.Nombre, simbvect)

				}
			} else {
				pusheo := pushvect.Pushear.Compilar_Expresion(ent, gen)
				pusheo1 := pushvect.Pushear.Ejecutar_Expresion(ent)
				if simbvect.TipoVect == pusheo.Tipo {
					tempo1 := gen.Crear_temporal()
					posi := simbvect.PosicionTabla
					gen.Agregar_Logica(tempo1 + " = SP + " + strconv.Itoa(posi) + ";\t\t// POSICION: " + pushvect.Identificador)
					gen.Agregar_Stack(tempo1, "HP")
					temp2 := gen.Crear_temporal()
					gen.Agregar_Logica(temp2 + " = HP;")
					simbvect.EsArreVect = simbvect.EsArreVect * 2
					gen.Agregar_Logica("HP = HP +" + strconv.Itoa(simbvect.EsArreVect) + ";")
					gen.Agregar_Logica("HEAP[int(HP)] = -2;\nHP = HP + 1;")
					temp3 := gen.Crear_temporal()
					for i := 0; i < simbvect.ValorLista.Len(); i++ {
						act := simbvect.ValorLista.GetValue(i).(simbolos.Valor)
						gen.Agregar_Logica(temp3 + " = " + temp2 + " + " + strconv.Itoa(i) + ";")
						gen.Agregar_Logica("HEAP[int(" + temp3 + ")] = " + fmt.Sprintf("%v", act.Valor) + ";")
					}
					gen.Agregar_Logica(temp3 + " = " + temp2 + " + " + strconv.Itoa(simbvect.ValorLista.Len()) + ";")
					gen.Agregar_Logica("HEAP[int(" + temp3 + ")] = " + fmt.Sprintf("%v", pusheo.Valor) + ";")

					//MODIFICACIONES EN EL VECTOR
					simbvect.ValorLista.Add(pusheo1)
					ent.Guardar_ArregloVector(simbvect.Nombre, simbvect)
				}
			}
		}
	}
	gen.Agregar_Comentario("FIN PUSH VECTOR -- " + pushvect.Identificador)
	return 0
}
