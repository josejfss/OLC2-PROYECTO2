package accesovect

import (
	entorno "OLC2-PROYECTO2/COMPILADOR/ENTORNO"
	generador "OLC2-PROYECTO2/COMPILADOR/GENERADOR"
	interfaces "OLC2-PROYECTO2/COMPILADOR/INTERFACES"
	reportes "OLC2-PROYECTO2/COMPILADOR/REPORTES"
	simbolos "OLC2-PROYECTO2/COMPILADOR/SIMBOLOS"
	"strconv"
	"time"
)

type AccessVect struct {
	Nombre   string
	Posicion interfaces.Expresion
	Linea    int
	Columna  int
}

func Naccessvect(nom string, pos interfaces.Expresion, lin int, col int) AccessVect {
	accvect := AccessVect{Nombre: nom, Posicion: pos, Linea: lin, Columna: col}
	return accvect
}

func (ac AccessVect) Ejecutar_Expresion(ent *entorno.Entorno) simbolos.Valor {
	if ent.Existe_ArreVect(ac.Nombre) {
		arrvect := ent.Obtener_ArreVect(ac.Nombre)
		posi := ac.Posicion.Ejecutar_Expresion(ent)
		for i := 0; i < arrvect.ValorLista.Len(); i++ {
			act := arrvect.ValorLista.GetValue(i).(simbolos.Valor)
			if posi.Valor.(int) == i {
				return simbolos.Valor{Tipo: act.Tipo, Valor: act.Valor}
			}
		}

	} else {
		t := time.Now()
		reportes.Nerror("El vector no existe.", ent.Nombre_Entorno, strconv.Itoa(ac.Linea), strconv.Itoa(ac.Columna), t.Format("2006-01-02 15:04:05"))
	}

	return simbolos.Valor{Tipo: simbolos.NULL, Valor: -1}
}

func (ac AccessVect) Compilar_Expresion(ent *entorno.Entorno, gen *generador.Generador_C3D) simbolos.ValoresC3D {
	gen.Agregar_Comentario("INCIO ACCESO A VECTOR -- " + ac.Nombre)
	if ent.Existe_ArreVect(ac.Nombre) {
		arrvect := ent.Obtener_ArreVect(ac.Nombre)
		if arrvect.DimensionesLista.Len() == 1 {
			posi := ac.Posicion.Compilar_Expresion(ent, gen)
			temp1 := gen.Crear_temporal()
			temp2 := gen.Crear_temporal()
			gen.Agregar_Logica(temp1 + " = SP +" + strconv.Itoa(arrvect.PosicionTabla) + ";")
			gen.Agregar_Logica(temp2 + " = STACK[int(" + temp1 + ")];")

			etiqueta_error := gen.Crear_label()
			gen.Eliminar_label(etiqueta_error)
			gen.Agregar_Logica("if (" + posi.Valor + " > " + strconv.Itoa(arrvect.Dimensiones) + ") goto " + etiqueta_error + ";")

			temp3 := gen.Crear_temporal()
			temp4 := gen.Crear_temporal()
			gen.Agregar_Logica(temp3 + " = " + temp2 + " + " + posi.Valor + ";")
			gen.Agregar_Logica(temp4 + " = HEAP[int(" + temp3 + ")];")
			gen.Agregar_Logica(etiqueta_error + ":")
			gen.AgregarError("INDICE FUERA DE LIMITE", strconv.Itoa(ac.Linea), strconv.Itoa(ac.Columna))
			return simbolos.ValoresC3D{Valor: temp4, EsTemporal: true, Tipo: arrvect.TipoVect, Label_verdadera: "", Label_false: ""}
		}

	} else {
		gen.AgregarError("ERROR-VECTOR NO EXISTE", strconv.Itoa(ac.Linea), strconv.Itoa(ac.Columna))
	}
	gen.Agregar_Comentario("FIN ACCESO A VECTOR -- " + ac.Nombre)
	gen.LiberarTodosTemporales()
	return simbolos.ValoresC3D{Valor: "0", EsTemporal: false, Tipo: simbolos.INTEGER, Label_verdadera: "", Label_false: ""}
}
