package bloque

type BloquesOpt struct {
	Nombre_Bloque   string
	TablaObjBloques map[string]ObjetoBloque
}

func Nuevo_Bloque(nombre string) *BloquesOpt {
	bloquesito := BloquesOpt{Nombre_Bloque: nombre}
	bloquesito.TablaObjBloques = make(map[string]ObjetoBloque)
	return &bloquesito
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
