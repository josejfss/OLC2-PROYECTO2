package primitivos

import objeto "OLC2-PROYECTO2/OPTIMIZADOR/OBJETO"

type Primitivo struct {
	Valor string
}

func NuevoPrimitivo(val string) Primitivo {
	datpri := Primitivo{val}
	return datpri
}

func (prim Primitivo) Optimizar_Expresion(block *objeto.Bloque) objeto.ObjetoBloque {
	return objeto.ObjetoBloque{
		Operacion: false,
		Opiz:      prim.Valor,
		Opde:      prim.Valor,
		Valor:     prim.Valor,
	}
}
