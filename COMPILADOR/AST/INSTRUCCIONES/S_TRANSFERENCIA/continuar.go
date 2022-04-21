package stransferencia

import (
	entorno "OLC2-PROYECTO2/COMPILADOR/ENTORNO"
	generador "OLC2-PROYECTO2/COMPILADOR/GENERADOR"
	simbolos "OLC2-PROYECTO2/COMPILADOR/SIMBOLOS"
)

type Continuar struct {
	Tipo_continuar simbolos.TipoExpresion
}

func Ncontinuar(tip_rompe simbolos.TipoExpresion) Continuar {
	return Continuar{Tipo_continuar: tip_rompe}
}

func (r Continuar) Ejecutar_Instruccion(ent *entorno.Entorno, ent2 *entorno.Entorno) interface{} {
	if r.Tipo_continuar == simbolos.METODO {
		return simbolos.Valor{Tipo: simbolos.CONTINUE, Valor: "continue"}
	}
	return 0
}

func (r Continuar) Compilar_Instruccion(ent *entorno.Entorno, gen *generador.Generador_C3D) interface{} {

	return "continue"
}
