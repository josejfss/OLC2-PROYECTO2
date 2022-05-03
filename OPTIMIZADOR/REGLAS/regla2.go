package reglas

import (
	reportes "OLC2-PROYECTO2/COMPILADOR/REPORTES"
	objeto "OLC2-PROYECTO2/OPTIMIZADOR/OBJETO"
	"strconv"

	"github.com/colegno/arraylist"
)

var Bloque_regla2 map[string]objeto.ObjetoBloque = make(map[string]objeto.ObjetoBloque)
var ListaBorrarTemporales *arraylist.List = arraylist.New()
var cont_regla int = 0

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
					cont_regla = 2
					AgregarTempEliminar(tiz)
					vals := simb.Opiz + simb.Ope + simb.Opde
					reportes.ReporteOpti("Bloques", "Regla "+strconv.Itoa(cont_regla), simb.Valor, vals, strconv.Itoa(simb.Linea))
				} else if simbde.EsTemporal {
					simb.Opde = simbde.Valor
					cont_regla = 2
					AgregarTempEliminar(tde)
					vals := simb.Opiz + simb.Ope + simb.Opde
					reportes.ReporteOpti("Bloques", "Regla "+strconv.Itoa(cont_regla), simb.Valor, vals, strconv.Itoa(simb.Linea))
				} else if simbiz.Constante {
					simb.Opiz = simbiz.Valor
					cont_regla = 4
					AgregarTempEliminar(tiz)
					vals := simb.Opiz + simb.Ope + simb.Opde
					reportes.ReporteOpti("Bloques", "Regla "+strconv.Itoa(cont_regla), simb.Valor, vals, strconv.Itoa(simb.Linea))
				} else if simbde.Constante {
					simb.Opde = simbde.Valor
					cont_regla = 4
					AgregarTempEliminar(tde)
					vals := simb.Opiz + simb.Ope + simb.Opde
					reportes.ReporteOpti("Bloques", "Regla "+strconv.Itoa(cont_regla), simb.Valor, vals, strconv.Itoa(simb.Linea))
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

	for j := 0; j < ListaBorrarTemporales.Len(); j++ {
		act := ListaBorrarTemporales.GetValue(j).(string)
		simbbs := block.Obtener_Tempo(act)
		reportes.ReporteOpti("Bloques", "Regla 3", simbbs.Valor, "-", strconv.Itoa(simbbs.Linea))
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
