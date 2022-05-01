package funnatvect

import (
	entorno "OLC2-PROYECTO2/COMPILADOR/ENTORNO"
	generador "OLC2-PROYECTO2/COMPILADOR/GENERADOR"
	interfaces "OLC2-PROYECTO2/COMPILADOR/INTERFACES"
	simbolos "OLC2-PROYECTO2/COMPILADOR/SIMBOLOS"
	"strconv"

	"github.com/colegno/arraylist"
)

type RemoveVect struct {
	Identificador string
	Posicion      interfaces.Expresion
	Linea         int
	Columna       int
}

func Nremovevect(ide string, pos interfaces.Expresion, lin int, col int) RemoveVect {
	remove := RemoveVect{Identificador: ide, Posicion: pos, Linea: lin, Columna: col}
	return remove
}

func (removect RemoveVect) Ejecutar_Instruccion(ent *entorno.Entorno, ent2 *entorno.Entorno) interface{} {
	return 0
}

func (removect RemoveVect) Compilar_Instruccion(ent *entorno.Entorno, gen *generador.Generador_C3D) interface{} {
	gen.Agregar_Comentario("INICIO REMOVE VECTOR -- " + removect.Identificador)
	if ent.ExisteAcual_ArreVect(removect.Identificador) {
		simbvect := ent.Obtener_ArreVect(removect.Identificador)
		if simbvect.Mutable {
			possi := removect.Posicion.Compilar_Expresion(ent, gen)
			possi1 := removect.Posicion.Ejecutar_Expresion(ent)

			tempo1 := gen.Crear_temporal()
			posi := simbvect.PosicionTabla
			temp_conca := gen.Crear_temporal()
			etiqueta_entrada := gen.Crear_label()
			gen.Eliminar_label(etiqueta_entrada)
			etiqueta_salida := gen.Crear_label()
			gen.Eliminar_label(etiqueta_salida)
			etiqueta_salida1 := gen.Crear_label()
			gen.Eliminar_label(etiqueta_salida1)

			gen.Agregar_Logica(tempo1 + " = SP + " + strconv.Itoa(posi) + ";\t\t// POSICION: " + removect.Identificador)
			temp2 := gen.Crear_temporal()
			gen.Agregar_Logica(temp2 + " = STACK[(int)" + tempo1 + "];")
			gen.Agregar_Stack(tempo1, "HP")
			temp3 := gen.Crear_temporal()
			gen.Agregar_Logica(temp3 + " = HP;")
			gen.Agregar_Logica("HP = HP + " + strconv.Itoa(simbvect.EsArreVect) + ";")
			gen.Agregar_Logica("HEAP[(int)HP] = -2;\nHP = HP + 1;")
			tempo_valor := temp2
			conta := gen.Crear_temporal()

			gen.Agregar_Logica(conta + " = 0;")
			gen.Agregar_Logica(etiqueta_entrada + ":")
			gen.Agregar_Logica(temp_conca + " = HEAP[(int)" + tempo_valor + "];")
			gen.Agregar_Logica("if ( " + conta + " == " + possi.Valor + ") goto " + etiqueta_salida1 + ";")
			gen.Agregar_Logica("if ( " + temp_conca + " == -2) goto " + etiqueta_salida + ";")

			gen.Agregar_Logica("HEAP[(int)" + temp3 + "] = " + temp_conca + ";")
			gen.Agregar_Logica(temp3 + " = " + temp3 + " + 1;")
			gen.Agregar_Logica(tempo_valor + " = " + tempo_valor + " + 1;")
			gen.Agregar_Logica(conta + " = " + conta + " + 1;\ngoto " + etiqueta_entrada + ";\n")
			gen.Agregar_Logica(etiqueta_salida1 + ":")
			gen.Agregar_Logica(tempo_valor + " = " + tempo_valor + " + 1;")
			gen.Agregar_Logica(conta + " = " + conta + " + 1;\ngoto " + etiqueta_entrada + ";\n")
			gen.Agregar_Logica(etiqueta_salida + ":")

			tconca := gen.Crear_temporal()
			etiqueta_entrada1 := gen.Crear_label()
			gen.Eliminar_label(etiqueta_entrada1)
			etiqueta_salida2 := gen.Crear_label()
			gen.Eliminar_label(etiqueta_salida2)
			otrotemp := gen.Crear_temporal()

			gen.Agregar_Logica(etiqueta_entrada1 + ":")
			gen.Agregar_Logica(tconca + " = HEAP[(int)" + temp3 + "];")
			gen.Agregar_Logica("if ( " + tconca + " != 0) goto " + etiqueta_salida2 + ";")
			gen.Agregar_Logica("HEAP[(int)" + temp3 + "] = HP;")
			gen.Agregar_Logica(otrotemp + " = HP;\nHP = HP + 1;")
			gen.Agregar_Logica("HEAP[(int)" + otrotemp + "] = -1;")
			gen.Agregar_Logica(temp3 + " = " + temp3 + " + 1;\ngoto " + etiqueta_entrada1 + ";")
			gen.Agregar_Logica(etiqueta_salida2 + ":")

			// gen.Agregar_Logica("HEAP[(int)" + temp3 + "] = HP;")
			// gen.Agregar_Logica("HEAP[(int)HP] = -1;\nHP = HP + 1;")

			//MODIFICACIONES EN EL VECTOR
			nueva_lista := *arraylist.New()
			for i := 0; i < simbvect.ValorLista.Len(); i++ {
				act := simbvect.ValorLista.GetValue(i).(simbolos.Valor)
				if possi1.Tipo == simbolos.INTEGER {
					if i != possi1.Valor.(int) {
						nueva_lista.Add(act)
					}
				}
			}
			simbvect.ValorLista = &nueva_lista
			ent.Guardar_ArregloVector(simbvect.Nombre, simbvect)
		}
	}
	gen.Agregar_Comentario("FIN REMOVE VECTOR -- " + removect.Identificador)
	//gen.LiberarTodosTemporales()
	return 0
}
