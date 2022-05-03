package reglas

import (
	objeto "OLC2-PROYECTO2/OPTIMIZADOR/OBJETO"

	"github.com/colegno/arraylist"
)

var Bloque_regla2 map[string]objeto.ObjetoBloque = make(map[string]objeto.ObjetoBloque)
var ListaBorrarTemporales *arraylist.List = arraylist.New()

func Regla2(block *objeto.Bloque) {
	for j := 0; j < block.ListaBloques.Len(); j++ {
		bloque := block.ListaBloques.GetValue(j).(*objeto.Bloque)
		for i := 0; i < len(bloque.Tabla_BloquesOpt); i++ {
			nom_temp := bloque.ListaTemporales.GetValue(i).(string)
			simb := bloque.Tabla_BloquesOpt[nom_temp]
			if simb.Operacion {
				tiz := simb.Opiz
				tde := simb.Opde
				simbiz := block.Obtener_Tempo(tiz)
				simbde := block.Obtener_Tempo(tde)
				if simbiz.EsTemporal {
					simb.Opiz = simbiz.Valor
					AgregarTempEliminar(tiz)
				} else if simbde.EsTemporal {
					simb.Opde = simbde.Valor
					AgregarTempEliminar(tde)
				} else if simbiz.Constante {
					simb.Opiz = simbiz.Valor
					AgregarTempEliminar(tiz)
				} else if simbde.Constante {
					simb.Opde = simbde.Valor
					AgregarTempEliminar(tde)
				}
				simb.Valor = simb.Opiz + simb.Ope + simb.Opde
				bloque.Modificar_Declaracion(nom_temp, simb)
			}
			// if i != 0 {

			// }
		}
	}

	for i := 0; i < ListaBorrarTemporales.Len(); i++ {
		act := ListaBorrarTemporales.GetValue(i).(string)
		for j := 0; j < block.ListaBloques.Len(); j++ {
			bloque := block.ListaBloques.GetValue(j).(*objeto.Bloque)
			delete(bloque.Tabla_Bloques, act)
			block.EliminarTempTodo(act)
			bloque.Eliminar_Temporal(act)
		}
	}

}

/*

for j := 0; j < block.ListaBloques.Len(); j++ {
		bloque := block.ListaBloques.GetValue(j).(*objeto.Bloque)
		for i := 0; i < len(bloque.Tabla_Bloques); i++ {
			nom_temp := bloque.ListaTemporales.GetValue(i).(string)
			simb := bloque.Tabla_Bloques[nom_temp]
			if i != 0 {
				if simb.Operacion {
					tiz := simb.Opiz
					tde := simb.Opde
					simbiz := bloque.Tabla_Reglas[tiz]
					simbde := bloque.Tabla_Reglas[tde]
					if simbiz.EsTemporal {
						simb.Opiz = simbiz.Valor
						AgregarTempEliminar(tiz)
					}
					if simbde.EsTemporal {
						simb.Opde = simbde.Valor
						AgregarTempEliminar(tde)
					}
					simb.Valor = simb.Opiz + simb.Ope + simb.Opde
					bloque.Modificar_Declaracion(nom_temp, simb)
				}
			}
		}
	}

*/

func AgregarTempEliminar(borrartemp string) {
	for i := 0; i < ListaBorrarTemporales.Len(); i++ {
		actu := ListaBorrarTemporales.GetValue(i).(string)
		if actu == borrartemp {
			return
		}
	}
	ListaBorrarTemporales.Add(borrartemp)
}
