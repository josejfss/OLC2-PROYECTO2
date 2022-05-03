package primitivos

import objeto "OLC2-PROYECTO2/OPTIMIZADOR/OBJETO"

type Primitivo struct {
	Valor string
	TipoO objeto.TipoObjeto
}

func NuevoPrimitivo(val string, tip objeto.TipoObjeto) Primitivo {
	datpri := Primitivo{val, tip}
	return datpri
}

func (prim Primitivo) Optimizar_Expresion(block *objeto.Bloque) objeto.ObjetoBloque {
	if prim.TipoO == objeto.INTEGER {
		return objeto.ObjetoBloque{
			Operacion:  false,
			EsTemporal: false,
			Constante:  true,
			Opiz:       prim.Valor,
			Opde:       prim.Valor,
			Valor:      prim.Valor,
		}
	} else if prim.TipoO == objeto.FLOAT {
		return objeto.ObjetoBloque{
			Operacion:  false,
			EsTemporal: false,
			Constante:  true,
			Opiz:       prim.Valor,
			Opde:       prim.Valor,
			Valor:      prim.Valor,
		}
	} else if prim.TipoO == objeto.TEMPORAL {
		return objeto.ObjetoBloque{
			Operacion:  false,
			EsTemporal: true,
			Constante:  false,
			Opiz:       prim.Valor,
			Opde:       prim.Valor,
			Valor:      prim.Valor,
		}
	} else {
		return objeto.ObjetoBloque{
			Operacion:  false,
			EsTemporal: false,
			Constante:  false,
			Opiz:       prim.Valor,
			Opde:       prim.Valor,
			Valor:      prim.Valor,
		}
	}

}
