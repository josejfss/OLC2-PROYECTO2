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

type ObjetoBloque struct {
	Operacion bool
	Temporal  string
	Opiz      string
	Opde      string
	Valor     string
}

type Bloque struct {
	NombreBloque    string
	NombreFuncion   string
	ListaBloques    *arraylist.List
	ListaTemporales *arraylist.List
	Tabla_Bloques   map[string]ObjetoBloque
}

func Nuevo_Bloque() *Bloque {
	blo := Bloque{NombreBloque: "Bloque" + strconv.Itoa(ContadorBloque)}
	ContadorBloque = ContadorBloque + 1
	blo.ListaBloques = arraylist.New()
	blo.ListaTemporales = arraylist.New()
	blo.Tabla_Bloques = make(map[string]ObjetoBloque)
	return &blo
}

func (block *Bloque) Guardar_Declaracion(nombre string, simvar ObjetoBloque) {
	block.Tabla_Bloques[nombre] = simvar
}

func (block *Bloque) Guardar_Bloque(bloque *Bloque) {
	block.ListaBloques.Add(bloque)
}
