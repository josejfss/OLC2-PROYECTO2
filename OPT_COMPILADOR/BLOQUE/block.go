package bloque

import "github.com/colegno/arraylist"

type BloquesOpt struct {
	Nombre_Bloque   string
	Encabezado      string
	TablaObjBloques map[string]ObjetoBloque
	ListaTemporales *arraylist.List
}

func Nuevo_Bloque(nombre string) *BloquesOpt {
	bloquesito := BloquesOpt{Nombre_Bloque: nombre}
	bloquesito.TablaObjBloques = make(map[string]ObjetoBloque)
	bloquesito.ListaTemporales = arraylist.New()
	return &bloquesito
}

func (block *BloquesOpt) Guardar(nom string, simb ObjetoBloque) {
	block.TablaObjBloques[nom] = simb
}

func (block *BloquesOpt) Regla1(val string) (bool, string) {
	for _, v := range block.TablaObjBloques {
		if v.Valor == val {
			return true, v.Asignacion
		}
	}
	return false, ""
}

func (block *BloquesOpt) Obtener_Temporal(nombre string) ObjetoBloque {
	if variable, ok := block.TablaObjBloques[nombre]; ok {
		return variable
	}
	return ObjetoBloque{Tipo: NULL, Asignacion: "", OpeIz: "", Ope: "", OpeDe: "", Linea: 0}
}

func (block *BloquesOpt) EliminarTemporalLista(nom string) {
	for i := 0; i < block.ListaTemporales.Len(); i++ {
		act := block.ListaTemporales.GetValue(i).(string)
		if act == nom {
			block.ListaTemporales.RemoveAtIndex(i)
		}
	}
}
