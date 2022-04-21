package stransferencia

import (
	entorno "OLC2-PROYECTO2/COMPILADOR/ENTORNO"
	generador "OLC2-PROYECTO2/COMPILADOR/GENERADOR"
	interfaces "OLC2-PROYECTO2/COMPILADOR/INTERFACES"
	simbolos "OLC2-PROYECTO2/COMPILADOR/SIMBOLOS"
)

type Retorno struct {
	Tipo_retorno simbolos.TipoExpresion
	Salida       interfaces.Expresion
}

func Nretorno(tipo_retorno simbolos.TipoExpresion, salida interfaces.Expresion) Retorno {
	if salida != nil {
		return Retorno{Tipo_retorno: tipo_retorno, Salida: salida}
	}
	return Retorno{Tipo_retorno: tipo_retorno, Salida: nil}
}

func (r Retorno) Ejecutar_Instruccion(ent *entorno.Entorno, ent2 *entorno.Entorno) interface{} {
	if r.Tipo_retorno == simbolos.METODO {
		return simbolos.Valor{Tipo: simbolos.NULL, Valor: 0}
	}

	retorno_salida := r.Salida.Ejecutar_Expresion(ent)
	return retorno_salida
}
func (r Retorno) Compilar_Instruccion(ent *entorno.Entorno, gen *generador.Generador_C3D) interface{} {

	return 0
}
