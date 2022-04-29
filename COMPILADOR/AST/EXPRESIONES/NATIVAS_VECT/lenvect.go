package nativasvect

import (
	entorno "OLC2-PROYECTO2/COMPILADOR/ENTORNO"
	generador "OLC2-PROYECTO2/COMPILADOR/GENERADOR"
	reportes "OLC2-PROYECTO2/COMPILADOR/REPORTES"
	simbolos "OLC2-PROYECTO2/COMPILADOR/SIMBOLOS"
	"strconv"
	"time"
)

type LenVect struct {
	Identificador string
	Linea         int
	Columna       int
}

func Nlenvect(nom string, lin int, col int) LenVect {
	lenvector := LenVect{Identificador: nom, Linea: lin, Columna: col}
	return lenvector
}

func (lenght LenVect) Ejecutar_Expresion(ent *entorno.Entorno) simbolos.Valor {
	if ent.Existe_ArreVect(lenght.Identificador) {
		simbvect := ent.Obtener_ArreVect(lenght.Identificador)
		return simbolos.Valor{Tipo: simbolos.INTEGER, Valor: simbvect.ValorLista.Len()}
	} else {
		t := time.Now()
		reportes.Nerror("No existe el vector", ent.Nombre_Entorno, strconv.Itoa(lenght.Linea), strconv.Itoa(lenght.Columna), t.Format("2006-01-02 15:04:05"))
	}
	//NO SE ACEPTO LOS TIPOS
	return simbolos.Valor{Tipo: simbolos.NULL, Valor: -1}
}

func (lenght LenVect) Compilar_Expresion(ent *entorno.Entorno, gen *generador.Generador_C3D) simbolos.ValoresC3D {
	gen.Agregar_Comentario("INICIANDO FUNCION NATIVA LEN VECTOR")
	if ent.Existe_ArreVect(lenght.Identificador) {
		simbvect := ent.Obtener_ArreVect(lenght.Identificador)
		nuevo_temporal := gen.Crear_temporal()
		gen.Agregar_Logica(nuevo_temporal + " = " + strconv.Itoa(simbvect.ValorLista.Len()) + ";")
		return simbolos.ValoresC3D{Valor: nuevo_temporal, EsTemporal: true, Tipo: simbolos.INTEGER, Label_verdadera: "", Label_false: ""}
	}
	gen.Agregar_Comentario("FINALIZANDO FUNCION NATIVA LEN VECTOR")
	return simbolos.ValoresC3D{Valor: "0", EsTemporal: false, Tipo: simbolos.INTEGER, Label_verdadera: "", Label_false: ""}
}
