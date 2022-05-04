package primitiivos

import bloque "OLC2-PROYECTO2/OPT_COMPILADOR/BLOQUE"

type Primitivo struct {
	Valor string
	TipoO bloque.TipoObjeto
}

func NuevoPrimitivo(val string, tip bloque.TipoObjeto) Primitivo {
	datpri := Primitivo{val, tip}
	return datpri
}

func (prim Primitivo) Opti_Expresion(block *bloque.BloquesOpt) bloque.ObjetoBloque {
	if prim.TipoO == bloque.INTEGER {
		return bloque.ObjetoBloque{
			Tipo:       prim.TipoO,
			Asignacion: prim.Valor,
			OpeIz:      prim.Valor,
			OpeDe:      prim.Valor,
			Ope:        prim.Valor,
			Valor:      prim.Valor,
		}
	} else if prim.TipoO == bloque.FLOAT {
		return bloque.ObjetoBloque{
			Tipo:       prim.TipoO,
			Asignacion: prim.Valor,
			OpeIz:      prim.Valor,
			OpeDe:      prim.Valor,
			Ope:        prim.Valor,
			Valor:      prim.Valor,
		}
	} else if prim.TipoO == bloque.TEMPORAL {
		return bloque.ObjetoBloque{
			Tipo:       prim.TipoO,
			Asignacion: prim.Valor,
			OpeIz:      prim.Valor,
			OpeDe:      prim.Valor,
			Ope:        prim.Valor,
			Valor:      prim.Valor,
		}
	} else {
		return bloque.ObjetoBloque{
			Tipo:       prim.TipoO,
			Asignacion: prim.Valor,
			OpeIz:      prim.Valor,
			OpeDe:      prim.Valor,
			Ope:        prim.Valor,
			Valor:      prim.Valor,
		}
	}

}
