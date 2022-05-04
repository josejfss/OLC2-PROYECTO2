package funcion

import (
	bloque "OLC2-PROYECTO2/OPT_COMPILADOR/BLOQUE"
	generando "OLC2-PROYECTO2/OPT_COMPILADOR/GENERANDO"
	inter "OLC2-PROYECTO2/OPT_COMPILADOR/INTER"

	"github.com/colegno/arraylist"
)

type Funca struct {
	Nombre string
	Lista  *arraylist.List
}

func Nfunca(nom string, list *arraylist.List) Funca {
	fun := Funca{Nombre: nom, Lista: list}
	return fun
}

func (mains Funca) Opti_Instruccion(block *bloque.BloquesOpt, genopt *generando.GenerandoOpti) interface{} {
	guardar := bloque.Nobjetolista(mains.Nombre, "void "+mains.Nombre+"() {")
	genopt.AgegarCodigoOpt(guardar)
	for i := 0; i < mains.Lista.Len(); i++ {
		mains.Lista.GetValue(i).(inter.Instruccion).Opti_Instruccion(block, genopt)
	}
	guardar1 := bloque.Nobjetolista("return", "return;\n}\n")
	genopt.AgegarCodigoOpt(guardar1)
	return 0
}
