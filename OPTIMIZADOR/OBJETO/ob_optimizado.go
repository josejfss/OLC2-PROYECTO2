package objeto

import (
	"strconv"

	"github.com/colegno/arraylist"
)

var ContadorBloque int = 0
var ContadorIf int = 0
var ContadorHeap int = 0
var ContadorStack int = 0
var ContadorPrint int = 0
var ContadorGoto int = 0
var ContadorReturn int = 0

type ObjetoBloque struct {
	Operacion  bool
	EsTemporal bool
	Constante  bool
	Temporal   string
	Opiz       string
	Ope        string
	Opde       string
	Valor      string
	Tipo       int
}

type Bloque struct {
	NombreBloque     string
	Encabezado       string
	NombreFuncion    string
	ListaBloques     *arraylist.List
	ListaTemporales  *arraylist.List
	ListaTodo        *arraylist.List
	Tabla_Bloques    map[string]ObjetoBloque
	Tabla_BloquesOpt map[string]ObjetoBloque
	Tabla_Reglas     map[string]ObjetoBloque
	Tabla_Junto      map[string]ObjetoBloque
}

func Nuevo_Bloque() *Bloque {
	blo := Bloque{NombreBloque: "Bloque" + strconv.Itoa(ContadorBloque)}
	ContadorBloque = ContadorBloque + 1
	blo.ListaBloques = arraylist.New()
	blo.ListaTemporales = arraylist.New()
	blo.ListaTodo = arraylist.New()
	blo.Tabla_Bloques = make(map[string]ObjetoBloque)
	blo.Tabla_BloquesOpt = make(map[string]ObjetoBloque)
	blo.Tabla_Reglas = make(map[string]ObjetoBloque)
	blo.Tabla_Junto = make(map[string]ObjetoBloque)
	return &blo
}

func (block *Bloque) Guardar_Declaracion(nombre string, simvar ObjetoBloque) {
	block.Tabla_Bloques[nombre] = simvar
}

func (block *Bloque) Guardar_Declaracion1(nombre string, simvar ObjetoBloque) {
	block.Tabla_BloquesOpt[nombre] = simvar
}

func (block *Bloque) Guardar_Bloque(bloque *Bloque) {
	block.ListaBloques.Add(bloque)
}

func (block *Bloque) Obtener_Declaracion(nombre string) ObjetoBloque {
	if variable, ok := block.Tabla_Bloques[nombre]; ok {
		return variable
	}
	return ObjetoBloque{Operacion: false, Temporal: "", Opiz: "", Opde: "", Valor: ""}
}

func (block *Bloque) Obtener_Declaracion1(nombre string) ObjetoBloque {
	if variable, ok := block.Tabla_BloquesOpt[nombre]; ok {
		return variable
	}
	return ObjetoBloque{Operacion: false, Temporal: "", Opiz: "", Opde: "", Valor: ""}
}

func (block *Bloque) Modificar_Declaracion(nombre string, simb ObjetoBloque) {
	block.Tabla_Bloques[nombre] = simb
	block.Tabla_BloquesOpt[nombre] = simb
}

func (block *Bloque) Eliminar_Temporal(borrar string) {
	for i := 0; i < block.ListaTemporales.Len(); i++ {
		eti_act := block.ListaTemporales.GetValue(i)
		if eti_act == borrar {
			block.ListaTemporales.RemoveAtIndex(i)
		}
	}
}

func (block *Bloque) EliminarTempTodo(borrar string) {
	for i := 0; i < block.ListaTodo.Len(); i++ {
		eti_act := block.ListaTodo.GetValue(i)
		if eti_act == borrar {
			block.ListaTodo.RemoveAtIndex(i)
		}
	}
}

func (block *Bloque) Obtener_Tempo(nombre string) ObjetoBloque {
	for i := 0; i < block.ListaBloques.Len(); i++ {
		bloques := block.ListaBloques.GetValue(i).(*Bloque)
		if variable, ok := bloques.Tabla_Reglas[nombre]; ok {
			return variable
		}
	}
	return ObjetoBloque{Operacion: false, Temporal: "", Opiz: "", Opde: "", Valor: ""}
}
