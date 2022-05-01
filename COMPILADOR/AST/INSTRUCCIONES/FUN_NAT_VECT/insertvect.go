package funnatvect

import (
	entorno "OLC2-PROYECTO2/COMPILADOR/ENTORNO"
	generador "OLC2-PROYECTO2/COMPILADOR/GENERADOR"
	interfaces "OLC2-PROYECTO2/COMPILADOR/INTERFACES"
	simbolos "OLC2-PROYECTO2/COMPILADOR/SIMBOLOS"
	"fmt"
	"strconv"

	"github.com/colegno/arraylist"
)

type InsertVect struct {
	Identificador string
	Posicion      interfaces.Expresion
	Insertado     interfaces.Expresion
	Linea         int
	Columna       int
}

func Ninsertvect(nom string, pos interfaces.Expresion, ins interfaces.Expresion, lin int, col int) InsertVect {
	insertado := InsertVect{Identificador: nom, Posicion: pos, Insertado: ins, Linea: lin, Columna: col}
	return insertado
}

func (insertvect InsertVect) Ejecutar_Instruccion(ent *entorno.Entorno, ent2 *entorno.Entorno) interface{} {
	return 0
}

func (insertvect InsertVect) Compilar_Instruccion(ent *entorno.Entorno, gen *generador.Generador_C3D) interface{} {
	gen.Agregar_Comentario("INICIO INSERT VECTOR -- " + insertvect.Identificador)
	if ent.ExisteAcual_ArreVect(insertvect.Identificador) {
		simbvect := ent.Obtener_ArreVect(insertvect.Identificador)
		if simbvect.Mutable {
			if simbvect.ValorLista.Len() < simbvect.EsArreVect {
				insert := insertvect.Insertado.Compilar_Expresion(ent, gen)
				insert1 := insertvect.Insertado.Ejecutar_Expresion(ent)
				possi := insertvect.Posicion.Compilar_Expresion(ent, gen)
				possi1 := insertvect.Posicion.Ejecutar_Expresion(ent)
				if simbvect.TipoVect == insert.Tipo {
					if simbvect.TipoVect == simbolos.TEXTO || simbvect.TipoVect == simbolos.YTEXTO {
						tempo1 := gen.Crear_temporal()
						posi := simbvect.PosicionTabla
						temp2 := gen.Crear_temporal()
						gen.Agregar_Logica(tempo1 + " = SP + " + strconv.Itoa(posi) + ";\t\t// POSICION: " + insertvect.Identificador)
						gen.Agregar_Logica(temp2 + " = STACK[(int)" + tempo1 + "];")
						temp3 := gen.Crear_temporal()
						gen.Agregar_Logica(temp3 + " = " + temp2 + " + " + fmt.Sprintf("%v", possi.Valor) + ";")
						gen.Agregar_Logica("HEAP[(int)" + temp3 + "] = HP;")
						for _, txt := range insert.Valor {
							f := int(txt)
							gen.Agregar_Logica("HEAP[(int)HP] = " + fmt.Sprintf("%v", f) + "; //LETRA-> " + string(txt))
							gen.Agregar_Logica("HP = HP + 1;")
						}
						gen.Agregar_Logica("HEAP[(int)HP] = -1;\nHP = HP + 1;")

						//MODIFICACIONES EN EL VECTOR
						nueva_lista := arraylist.New()
						if simbvect.ValorLista.Len() == 0 {
							simbvect.ValorLista.Add(insert1)
						} else {
							for i := 0; i < simbvect.ValorLista.Len(); i++ {
								act := simbvect.ValorLista.GetValue(i).(simbolos.Valor)
								if possi1.Tipo == simbolos.INTEGER {
									if i == possi1.Valor.(int) {
										nueva_lista.Add(insert1)
									} else {
										nueva_lista.Add(act)
									}
								}
							}
							simbvect.ValorLista = nueva_lista.Clone()
						}

						ent.Guardar_ArregloVector(simbvect.Nombre, simbvect)

					} else {
						tempo1 := gen.Crear_temporal()
						posi := simbvect.PosicionTabla
						temp2 := gen.Crear_temporal()
						gen.Agregar_Logica(tempo1 + " = SP + " + strconv.Itoa(posi) + ";\t\t// POSICION: " + insertvect.Identificador)
						gen.Agregar_Logica(temp2 + " = STACK[(int)" + tempo1 + "];")
						temp3 := gen.Crear_temporal()
						gen.Agregar_Logica(temp3 + " = " + temp2 + " + " + fmt.Sprintf("%v", possi.Valor) + ";")
						gen.Agregar_Logica("HEAP[(int)" + temp3 + "] = " + fmt.Sprintf("%v", insert.Valor) + ";")

						//MODIFICACIONES EN EL VECTOR
						nueva_lista := arraylist.New()
						for i := 0; i < simbvect.ValorLista.Len(); i++ {
							act := simbvect.ValorLista.GetValue(i).(simbolos.Valor)
							if possi1.Tipo == simbolos.INTEGER {
								if i == possi1.Valor.(int) {
									nueva_lista.Add(insert1)
								} else {
									nueva_lista.Add(act)
								}
							}
						}
						simbvect.ValorLista = nueva_lista.Clone()
						ent.Guardar_ArregloVector(simbvect.Nombre, simbvect)
					}

				}
			} else {
				insert := insertvect.Insertado.Compilar_Expresion(ent, gen)
				insert1 := insertvect.Insertado.Ejecutar_Expresion(ent)
				possi := insertvect.Posicion.Compilar_Expresion(ent, gen)
				possi1 := insertvect.Posicion.Ejecutar_Expresion(ent)
				if simbvect.TipoVect == insert.Tipo {
					if simbvect.TipoVect == simbolos.TEXTO || simbvect.TipoVect == simbolos.YTEXTO {
						tempo1 := gen.Crear_temporal()
						posi := simbvect.PosicionTabla
						gen.Agregar_Logica(tempo1 + " = SP + " + strconv.Itoa(posi) + ";\t\t// POSICION: " + insertvect.Identificador)
						gen.Agregar_Stack(tempo1, "HP")
						temp2 := gen.Crear_temporal()
						gen.Agregar_Logica(temp2 + " = HP;")
						simbvect.EsArreVect = simbvect.EsArreVect * 2
						gen.Agregar_Logica("HP = HP +" + strconv.Itoa(simbvect.EsArreVect) + ";")
						gen.Agregar_Logica("HEAP[(int)HP] = -2;\nHP = HP + 1;")
						temp3 := gen.Crear_temporal()

						temp_conca := gen.Crear_temporal()
						etiqueta_entrada := gen.Crear_label()
						gen.Eliminar_label(etiqueta_entrada)
						etiqueta_salida := gen.Crear_label()
						gen.Eliminar_label(etiqueta_salida)
						tempo_valor := temp2
						conta := gen.Crear_temporal()

						gen.Agregar_Logica(etiqueta_entrada + ":")
						gen.Agregar_Logica(temp_conca + " = HEAP[(int)" + tempo_valor + "];")
						gen.Agregar_Logica("if ( " + temp_conca + " == -2) goto " + etiqueta_salida + ";")

						gen.Agregar_Logica(temp3 + " = " + temp2 + " + " + conta + ";")
						gen.Agregar_Logica("HEAP[(int)" + temp3 + "] = " + temp_conca + ";")
						gen.Agregar_Logica(conta + " = " + conta + " + 1;")
						gen.Agregar_Logica(tempo_valor + " = " + tempo_valor + " + 1;\ngoto " + etiqueta_entrada + ";\n")
						gen.Agregar_Logica(etiqueta_salida + ":")
						gen.Agregar_Logica(temp3 + " = " + temp2 + " + " + strconv.Itoa(simbvect.ValorLista.Len()) + ";")
						gen.Agregar_Logica("HEAP[(int)" + temp3 + "] = HP;")
						for _, txt := range insert.Valor {
							f := int(txt)
							gen.Agregar_Logica("HEAP[(int)HP] = " + fmt.Sprintf("%v", f) + "; //LETRA-> " + string(txt))
							gen.Agregar_Logica("HP = HP + 1;")
						}
						gen.Agregar_Logica("HEAP[(int)HP] = -1;\nHP = HP + 1;")

						//MODIFICACIONES EN EL VECTOR
						nueva_lista := arraylist.New()
						for i := 0; i < simbvect.ValorLista.Len(); i++ {
							act := simbvect.ValorLista.GetValue(i).(simbolos.Valor)
							if possi1.Tipo == simbolos.INTEGER {
								if i == possi1.Valor.(int) {
									nueva_lista.Add(insert1)
								} else {
									nueva_lista.Add(act)
								}
							}
						}
						simbvect.ValorLista = nueva_lista.Clone()
						ent.Guardar_ArregloVector(simbvect.Nombre, simbvect)
					} else {
						tempo1 := gen.Crear_temporal()
						posi := simbvect.PosicionTabla
						gen.Agregar_Logica(tempo1 + " = SP + " + strconv.Itoa(posi) + ";\t\t// POSICION: " + insertvect.Identificador)
						gen.Agregar_Stack(tempo1, "HP")
						temp2 := gen.Crear_temporal()
						gen.Agregar_Logica(temp2 + " = HP;")
						simbvect.EsArreVect = simbvect.EsArreVect * 2
						gen.Agregar_Logica("HP = HP +" + strconv.Itoa(simbvect.EsArreVect) + ";")
						gen.Agregar_Logica("HEAP[(int)HP] = -2;\nHP = HP + 1;")
						temp3 := gen.Crear_temporal()
						for i := 0; i < simbvect.ValorLista.Len(); i++ {
							act := simbvect.ValorLista.GetValue(i).(simbolos.Valor)
							if possi1.Tipo == simbolos.INTEGER {
								if i == possi1.Valor.(int) {
									gen.Agregar_Logica(temp3 + " = " + temp2 + " + " + fmt.Sprintf("%v", possi.Valor) + ";")
									gen.Agregar_Logica("HEAP[(int)" + temp3 + "] = " + fmt.Sprintf("%v", insert.Valor) + ";")
								} else {
									gen.Agregar_Logica(temp3 + " = " + temp2 + " + " + strconv.Itoa(i) + ";")
									gen.Agregar_Logica("HEAP[(int)" + temp3 + "] = " + fmt.Sprintf("%v", act.Valor) + ";")
								}
							}
						}

						//MODIFICACIONES EN EL VECTOR
						nueva_lista := arraylist.New()
						for i := 0; i < simbvect.ValorLista.Len(); i++ {
							act := simbvect.ValorLista.GetValue(i).(simbolos.Valor)
							if possi1.Tipo == simbolos.INTEGER {
								if i == possi1.Valor.(int) {
									nueva_lista.Add(insert1)
								} else {
									nueva_lista.Add(act)
								}
							}
						}
						simbvect.ValorLista = nueva_lista.Clone()
						ent.Guardar_ArregloVector(simbvect.Nombre, simbvect)
					}
				}
			}
		}
	}
	gen.Agregar_Comentario("FIN INSERT VECTOR -- " + insertvect.Identificador)
	//gen.LiberarTodosTemporales()
	return 0
}
