package reglas

import objeto "OLC2-PROYECTO2/OPTIMIZADOR/OBJETO"

var Bloque_regla1 map[string]objeto.ObjetoBloque = make(map[string]objeto.ObjetoBloque)

func Regla1(block *objeto.Bloque) {
	for i := 0; i < block.ListaBloques.Len(); i++ {
		bloque := block.ListaBloques.GetValue(i).(*objeto.Bloque)
		for j := 0; j < len(bloque.Tabla_BloquesOpt); j++ {
			nombre_temporal := bloque.ListaTemporales.GetValue(j).(string)
			simb_bloque := bloque.Obtener_Declaracion1(nombre_temporal)
			if len(Bloque_regla1) == 0 {
				Bloque_regla1[nombre_temporal] = simb_bloque
			} else {
				for k := 0; k < len(Bloque_regla1); k++ {
					nom_temp := bloque.ListaTemporales.GetValue(k).(string)
					bloquesito := Bloque_regla1[nom_temp]
					if simb_bloque.Valor == bloquesito.Valor {
						simb_bloque.Opiz = bloquesito.Temporal
						simb_bloque.Opde = bloquesito.Temporal
						simb_bloque.Valor = bloquesito.Temporal
						simb_bloque.EsTemporal = true
						simb_bloque.Operacion = false
						bloque.Modificar_Declaracion(nombre_temporal, simb_bloque)
						//bloque_regla1[nombre_temporal] = simb_bloque
					}
				}
				Bloque_regla1[nombre_temporal] = simb_bloque
			}
		}
		bloque.Tabla_Reglas = Bloque_regla1
		Bloque_regla1 = make(map[string]objeto.ObjetoBloque)

	}
}
