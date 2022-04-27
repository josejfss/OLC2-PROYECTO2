package accesoarre

import (
	declaracionarre "OLC2-PROYECTO2/COMPILADOR/AST/INSTRUCCIONES/DECLARACIONARRE"
	entorno "OLC2-PROYECTO2/COMPILADOR/ENTORNO"
	generador "OLC2-PROYECTO2/COMPILADOR/GENERADOR"
	interfaces "OLC2-PROYECTO2/COMPILADOR/INTERFACES"
	reportes "OLC2-PROYECTO2/COMPILADOR/REPORTES"
	simbolos "OLC2-PROYECTO2/COMPILADOR/SIMBOLOS"
	"strconv"
	"time"

	"github.com/colegno/arraylist"
)

type AccessArre struct {
	Nombre   string
	Posicion *arraylist.List
	Linea    int
	Columna  int
}

func Naccessarre(nom string, pos *arraylist.List, lin int, col int) AccessArre {
	accarr := AccessArre{Nombre: nom, Posicion: pos, Linea: lin, Columna: col}
	return accarr
}

func (ac AccessArre) Ejecutar_Expresion(ent *entorno.Entorno) simbolos.Valor {
	if ent.Existe_ArreVect(ac.Nombre) {
		arrvect := ent.Obtener_ArreVect(ac.Nombre)
		conta := 0
		if arrvect.DimensionesLista.Len() == ac.Posicion.Len() {
			for i := 0; i < ac.Posicion.Len(); i++ {
				posact := ac.Posicion.GetValue(i).(interfaces.Expresion).Ejecutar_Expresion(ent)
				if i == 0 {
					tam := arrvect.DimensionesLista.GetValue(i).(declaracionarre.TipoDeclaArre).Dimension.Ejecutar_Expresion(ent)
					if posact.Tipo == simbolos.INTEGER && tam.Tipo == simbolos.INTEGER {
						if posact.Valor.(int) < tam.Valor.(int) {
							conta = conta + 1
						} else {
							t := time.Now()
							reportes.Nerror("La posicion debe ser menor al tamaño.", ent.Nombre_Entorno, strconv.Itoa(ac.Linea), strconv.Itoa(ac.Columna), t.Format("2006-01-02 15:04:05"))
						}
					} else {
						t := time.Now()
						reportes.Nerror("La posicion debe ser de tipo entero.", ent.Nombre_Entorno, strconv.Itoa(ac.Linea), strconv.Itoa(ac.Columna), t.Format("2006-01-02 15:04:05"))
					}
				} else {
					tam := arrvect.DimensionesLista.GetValue(i).(interfaces.Expresion).Ejecutar_Expresion(ent)
					if posact.Tipo == simbolos.INTEGER && tam.Tipo == simbolos.INTEGER {
						if posact.Valor.(int) < tam.Valor.(int) {
							conta = conta + 1
						} else {
							t := time.Now()
							reportes.Nerror("La posicion debe ser menor al tamaño.", ent.Nombre_Entorno, strconv.Itoa(ac.Linea), strconv.Itoa(ac.Columna), t.Format("2006-01-02 15:04:05"))
						}
					} else {
						t := time.Now()
						reportes.Nerror("La posicion debe ser de tipo entero.", ent.Nombre_Entorno, strconv.Itoa(ac.Linea), strconv.Itoa(ac.Columna), t.Format("2006-01-02 15:04:05"))
					}
				}
			}

			if conta == 3 && conta == arrvect.DimensionesLista.Len() {
				posi0 := ac.Posicion.GetValue(0).(interfaces.Expresion).Ejecutar_Expresion(ent)
				posi1 := ac.Posicion.GetValue(1).(interfaces.Expresion).Ejecutar_Expresion(ent)
				posi2 := ac.Posicion.GetValue(1).(interfaces.Expresion).Ejecutar_Expresion(ent)
				tama1 := arrvect.DimensionesLista.GetValue(1).(interfaces.Expresion).Ejecutar_Expresion(ent)
				tama2 := arrvect.DimensionesLista.GetValue(2).(interfaces.Expresion).Ejecutar_Expresion(ent)
				posifinal := (posi2.Valor.(int)*tama2.Valor.(int)+posi1.Valor.(int))*tama1.Valor.(int) + posi0.Valor.(int)
				for i := 0; i < arrvect.ValorLista.Len(); i++ {
					act := arrvect.ValorLista.GetValue(i).(simbolos.Valor)
					if posifinal == i {
						return simbolos.Valor{Tipo: act.Tipo, Valor: act.Valor}
					}
				}
			} else if conta == 2 && conta == arrvect.DimensionesLista.Len() {
				posi0 := ac.Posicion.GetValue(0).(interfaces.Expresion).Ejecutar_Expresion(ent)
				posi1 := ac.Posicion.GetValue(1).(interfaces.Expresion).Ejecutar_Expresion(ent)
				tama0 := arrvect.DimensionesLista.GetValue(0).(declaracionarre.TipoDeclaArre).Dimension.Ejecutar_Expresion(ent)
				posifinal := posi1.Valor.(int)*tama0.Valor.(int) + posi0.Valor.(int)
				for i := 0; i < arrvect.ValorLista.Len(); i++ {
					act := arrvect.ValorLista.GetValue(i).(simbolos.Valor)
					if posifinal == i {
						return simbolos.Valor{Tipo: act.Tipo, Valor: act.Valor}
					}
				}
			} else if conta == 1 && conta == arrvect.DimensionesLista.Len() {
				posi := ac.Posicion.GetValue(0).(interfaces.Expresion).Ejecutar_Expresion(ent)
				for i := 0; i < arrvect.ValorLista.Len(); i++ {
					act := arrvect.ValorLista.GetValue(i).(simbolos.Valor)
					if posi.Valor.(int) == i {
						return simbolos.Valor{Tipo: act.Tipo, Valor: act.Valor}
					}
				}
			}
		} else {
			t := time.Now()
			reportes.Nerror("El tamaño del vector y el acceso deben de ser del mismo tamaño.", ent.Nombre_Entorno, strconv.Itoa(ac.Linea), strconv.Itoa(ac.Columna), t.Format("2006-01-02 15:04:05"))
		}
	} else {
		t := time.Now()
		reportes.Nerror("El vector no existe.", ent.Nombre_Entorno, strconv.Itoa(ac.Linea), strconv.Itoa(ac.Columna), t.Format("2006-01-02 15:04:05"))
	}

	return simbolos.Valor{Tipo: simbolos.NULL, Valor: -1}
}

func (ac AccessArre) Compilar_Expresion(ent *entorno.Entorno, gen *generador.Generador_C3D) simbolos.ValoresC3D {
	gen.Agregar_Comentario("INCIO ACCESO A ARREGLO -- " + ac.Nombre)
	if ent.Existe_ArreVect(ac.Nombre) {
		arrvect := ent.Obtener_ArreVect(ac.Nombre)
		conta := 0
		if arrvect.DimensionesLista.Len() == ac.Posicion.Len() {
			for i := 0; i < ac.Posicion.Len(); i++ {
				posact := ac.Posicion.GetValue(i).(interfaces.Expresion).Ejecutar_Expresion(ent)
				if i == 0 {
					tam := arrvect.DimensionesLista.GetValue(i).(declaracionarre.TipoDeclaArre).Dimension.Ejecutar_Expresion(ent)
					if posact.Tipo == simbolos.INTEGER && tam.Tipo == simbolos.INTEGER {
						if posact.Valor.(int) < tam.Valor.(int) {
							conta = conta + 1
						} else {
							gen.AgregarError("ERROR-POS<TAM", strconv.Itoa(ac.Linea), strconv.Itoa(ac.Columna))
						}
					} else {
						gen.AgregarError("ERROR-TIPO INT", strconv.Itoa(ac.Linea), strconv.Itoa(ac.Columna))
					}
				} else {
					tam := arrvect.DimensionesLista.GetValue(i).(interfaces.Expresion).Ejecutar_Expresion(ent)
					if posact.Tipo == simbolos.INTEGER && tam.Tipo == simbolos.INTEGER {
						if posact.Valor.(int) < tam.Valor.(int) {
							conta = conta + 1
						} else {
							gen.AgregarError("ERROR-POS<TAM", strconv.Itoa(ac.Linea), strconv.Itoa(ac.Columna))
						}
					} else {
						gen.AgregarError("ERROR-TIPO INT", strconv.Itoa(ac.Linea), strconv.Itoa(ac.Columna))
					}
				}
			}

			if conta == 3 && conta == arrvect.DimensionesLista.Len() {
				posi0 := ac.Posicion.GetValue(0).(interfaces.Expresion).Compilar_Expresion(ent, gen)
				posi1 := ac.Posicion.GetValue(1).(interfaces.Expresion).Compilar_Expresion(ent, gen)
				posi2 := ac.Posicion.GetValue(1).(interfaces.Expresion).Compilar_Expresion(ent, gen)
				tama1 := arrvect.DimensionesLista.GetValue(1).(interfaces.Expresion).Compilar_Expresion(ent, gen)
				tama2 := arrvect.DimensionesLista.GetValue(2).(interfaces.Expresion).Compilar_Expresion(ent, gen)

				//TEMPORALES PARA EL CALCULO DE LA POSICION FINAL
				temp1 := gen.Crear_temporal()
				temp2 := gen.Crear_temporal()
				temp3 := gen.Crear_temporal()
				temp4 := gen.Crear_temporal()

				gen.Agregar_Logica(temp1 + " = " + posi2.Valor + " * " + tama2.Valor + ";")
				gen.Agregar_Logica(temp2 + " = " + temp1 + " + " + posi1.Valor + ";")
				gen.Agregar_Logica(temp3 + " = " + temp2 + " * " + tama1.Valor + ";")
				gen.Agregar_Logica(temp4 + " = " + temp3 + " + " + posi0.Valor + ";")

				temp5 := gen.Crear_temporal()
				temp6 := gen.Crear_temporal()
				gen.Agregar_Logica(temp5 + " = SP +" + strconv.Itoa(arrvect.PosicionTabla) + ";")
				gen.Agregar_Logica(temp6 + " = STACK[int(" + temp5 + ")];")

				//MOVIENDOSE A LA POSICION EN EL HEAP
				temp7 := gen.Crear_temporal()
				gen.Agregar_Logica(temp7 + " = " + temp6 + " + " + temp4 + ";")

				temp8 := gen.Crear_temporal()
				gen.Agregar_Logica(temp8 + " = HEAP[int(" + temp7 + ")];")
				return simbolos.ValoresC3D{Valor: temp8, EsTemporal: true, Tipo: arrvect.TipoVect, Label_verdadera: "", Label_false: ""}

			} else if conta == 2 && conta == arrvect.DimensionesLista.Len() {
				posi0 := ac.Posicion.GetValue(0).(interfaces.Expresion).Compilar_Expresion(ent, gen)
				posi1 := ac.Posicion.GetValue(1).(interfaces.Expresion).Compilar_Expresion(ent, gen)
				tama0 := arrvect.DimensionesLista.GetValue(0).(declaracionarre.TipoDeclaArre).Dimension.Compilar_Expresion(ent, gen)

				//TEMPORALES PARA EL CALCULO DE LA POSICION FINAL
				temp1 := gen.Crear_temporal()
				temp2 := gen.Crear_temporal()

				gen.Agregar_Logica(temp1 + " = " + posi1.Valor + " * " + tama0.Valor + ";")
				gen.Agregar_Logica(temp2 + " = " + temp1 + " + " + posi0.Valor + ";")
				//posifinal := posi0.Valor.(int)*tama0.Valor.(int) + posi1.Valor.(int)
				temp3 := gen.Crear_temporal()
				temp4 := gen.Crear_temporal()
				gen.Agregar_Logica(temp3 + " = SP + " + strconv.Itoa(arrvect.PosicionTabla) + ";")
				gen.Agregar_Logica(temp4 + " = STACK[int(" + temp3 + ")];")

				temp5 := gen.Crear_temporal()
				temp6 := gen.Crear_temporal()
				gen.Agregar_Logica(temp5 + " = " + temp4 + " + " + temp2 + ";")
				gen.Agregar_Logica(temp6 + " = HEAP[int(" + temp5 + ")];")
				return simbolos.ValoresC3D{Valor: temp6, EsTemporal: true, Tipo: arrvect.TipoVect, Label_verdadera: "", Label_false: ""}

			} else if conta == 1 && conta == arrvect.DimensionesLista.Len() {
				posi := ac.Posicion.GetValue(0).(interfaces.Expresion).Compilar_Expresion(ent, gen)
				temp1 := gen.Crear_temporal()
				temp2 := gen.Crear_temporal()
				gen.Agregar_Logica(temp1 + " = SP +" + strconv.Itoa(arrvect.PosicionTabla) + ";")
				gen.Agregar_Logica(temp2 + " = STACK[int(" + temp1 + ")];")

				temp3 := gen.Crear_temporal()
				temp4 := gen.Crear_temporal()
				gen.Agregar_Logica(temp3 + " = " + temp2 + " + " + posi.Valor + ";")
				gen.Agregar_Logica(temp4 + " = HEAP[int(" + temp3 + ")];")
				return simbolos.ValoresC3D{Valor: temp4, EsTemporal: true, Tipo: arrvect.TipoVect, Label_verdadera: "", Label_false: ""}
			}
		} else {
			gen.AgregarError("ERROR-TAMAÑO!=POSICION", strconv.Itoa(ac.Linea), strconv.Itoa(ac.Columna))
		}
	} else {
		gen.AgregarError("ERROR-VECTOR NO EXISTE", strconv.Itoa(ac.Linea), strconv.Itoa(ac.Columna))
	}
	gen.Agregar_Comentario("FIN ACCESO A ARREGLO -- " + ac.Nombre)
	gen.LiberarTodosTemporales()
	return simbolos.ValoresC3D{Valor: "0", EsTemporal: false, Tipo: simbolos.INTEGER, Label_verdadera: "", Label_false: ""}
}
