package reglas

import objeto "OLC2-PROYECTO2/OPTIMIZADOR/OBJETO"

func Regla4(block *objeto.Bloque) {
	for j := 0; j < block.ListaBloques.Len(); j++ {
		bloque := block.ListaBloques.GetValue(j).(*objeto.Bloque)
		for i := 0; i < len(bloque.Tabla_Reglas); i++ {
			nom_temp := bloque.ListaTemporales.GetValue(i).(string)
			simb := bloque.Tabla_Reglas[nom_temp]
			if simb.Operacion {
				tiz := simb.Opiz
				tde := simb.Opde
				simbiz := bloque.Tabla_Reglas[tiz]
				simbde := bloque.Tabla_Reglas[tde]
				if simbiz.Constante {
					simb.Opiz = simbiz.Valor
					AgregarTempEliminar(tiz)
				}
				if simbde.Constante {
					simb.Opde = simbde.Valor
					AgregarTempEliminar(tde)
				}
				simb.Valor = simb.Opiz + simb.Ope + simb.Opde
				bloque.Modificar_Declaracion(nom_temp, simb)
			}
		}
	}
}
