package stransferencia

import (
	entorno "OLC2-PROYECTO2/COMPILADOR/ENTORNO"
	generador "OLC2-PROYECTO2/COMPILADOR/GENERADOR"
	interfaces "OLC2-PROYECTO2/COMPILADOR/INTERFACES"
	simbolos "OLC2-PROYECTO2/COMPILADOR/SIMBOLOS"
)

type Romper struct {
	Tipo_rompe simbolos.TipoExpresion
	Salida     interfaces.Expresion
}

func Nromper(tip_rompe simbolos.TipoExpresion, sa interfaces.Expresion) Romper {
	if sa != nil {
		return Romper{Tipo_rompe: tip_rompe, Salida: sa}
	}
	return Romper{Tipo_rompe: tip_rompe, Salida: nil}
}

func (r Romper) Ejecutar_Instruccion(ent *entorno.Entorno, ent2 *entorno.Entorno) interface{} {
	if r.Tipo_rompe == simbolos.METODO {
		return simbolos.Valor{Tipo: simbolos.BREAK, Valor: "break"}
	}

	retorno_salida := r.Salida.Ejecutar_Expresion(ent)
	return simbolos.Valor{Tipo: simbolos.BREAK, Valor: retorno_salida}
}

func (r Romper) Compilar_Instruccion(ent *entorno.Entorno, gen *generador.Generador_C3D) interface{} {

	return "break"
}
