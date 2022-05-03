package reglas

import objeto "OLC2-PROYECTO2/OPTIMIZADOR/OBJETO"

func Regla3(block *objeto.Bloque) {
	for i := 0; i < ListaBorrarTemporales.Len(); i++ {
		act := ListaBorrarTemporales.GetValue(i).(string)
		for j := 0; j < block.ListaBloques.Len(); j++ {
			bloque := block.ListaBloques.GetValue(j).(*objeto.Bloque)
			delete(bloque.Tabla_Bloques, act)
			delete(bloque.Tabla_BloquesOpt, act)
			delete(bloque.Tabla_Reglas, act)
			delete(Bloque_regla1, act)
			block.Eliminar_Temporal(act)
			block.EliminarTempTodo(act)
			bloque.Eliminar_Temporal(act)
		}
	}
}
