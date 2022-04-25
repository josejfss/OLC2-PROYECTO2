package aritmetica

import (
	entorno "OLC2-PROYECTO2/COMPILADOR/ENTORNO"
	generador "OLC2-PROYECTO2/COMPILADOR/GENERADOR"
	interfaces "OLC2-PROYECTO2/COMPILADOR/INTERFACES"
	reportes "OLC2-PROYECTO2/COMPILADOR/REPORTES"
	simbolos "OLC2-PROYECTO2/COMPILADOR/SIMBOLOS"
	"fmt"
	"math"
	"strconv"
	"time"
)

type Potencia struct {
	Tipo    string
	Oi      interfaces.Expresion
	Od      interfaces.Expresion
	Linea   int
	Columna int
}

func Npotencia(t string, opi interfaces.Expresion, opd interfaces.Expresion, li int, col int) Potencia {
	pot := Potencia{Tipo: t, Oi: opi, Od: opd, Linea: li, Columna: col}
	return pot
}

func (pot Potencia) Ejecutar_Expresion(ent *entorno.Entorno) simbolos.Valor {

	//MATRIZ DOMINANTE POTENCIA
	//			INTEGER				FLOAT			  STRING		    CHAR		      BOOL			 &STRING		NULL
	potencia_dominante := [7][7]simbolos.TipoExpresion{
		//INTEGER
		{simbolos.INTEGER, simbolos.FLOAT, simbolos.NULL, simbolos.INTEGER, simbolos.NULL, simbolos.NULL, simbolos.NULL},
		//FLOAT
		{simbolos.FLOAT, simbolos.FLOAT, simbolos.NULL, simbolos.FLOAT, simbolos.NULL, simbolos.NULL, simbolos.NULL},
		//BOOLEAN
		{simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL},
		//CHAR
		{simbolos.INTEGER, simbolos.FLOAT, simbolos.NULL, simbolos.CHAR, simbolos.NULL, simbolos.NULL, simbolos.NULL},
		//TEXTO
		{simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL},
		//&TEXTO
		{simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL},
		//NULL
		{simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL},
	}

	//SE OBTIENE EL VALOR DEL LOS OPERADOES IZQUIERDA Y DERECHO
	var operador_izq simbolos.Valor = pot.Oi.Ejecutar_Expresion(ent)
	var operador_der simbolos.Valor = pot.Od.Ejecutar_Expresion(ent)

	//SE VALIDA EL TIPO DE LA POTENCIA
	var tipo_potencia simbolos.TipoExpresion

	//SE VALIDA QUE TIPO DE POTENCIA ES SI ENTERA O DECIMAL
	switch pot.Tipo {
	case "i64":
		{
			//SE OBTIENE EL TIPO DE LA POTENCIA
			tipo_potencia = potencia_dominante[operador_izq.Tipo][operador_der.Tipo]
			//SE VERIFICA QUE SEA ENTERA
			if tipo_potencia == simbolos.INTEGER {
				//SE HACE LAS CONVERSIONES
				val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", operador_izq.Valor), 64)
				val2, _ := strconv.ParseFloat(fmt.Sprintf("%v", operador_der.Valor), 64)
				var resu int = int(math.Pow(val1, val2))
				//SE RETORNA SIMBOLO
				return simbolos.Valor{Tipo: simbolos.INTEGER, Valor: resu}
			} else {
				//ERROR
				t := time.Now()
				reportes.Nerror("No es posible realizar la potencia, tipos no compatibles", ent.Nombre_Entorno, strconv.Itoa(pot.Linea), strconv.Itoa(pot.Columna), t.Format("2006-01-02 15:04:05"))
			}
		}
	case "f64":
		{
			//SE OBTIENE EL TIPO DE LA POTENCIA
			tipo_potencia = potencia_dominante[operador_izq.Tipo][operador_der.Tipo]
			//SE VERIFICA QUE SEA DECIMAL
			if tipo_potencia == simbolos.FLOAT {
				//SE HACE LAS CONVERSIONES
				val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", operador_izq.Valor), 64)
				val2, _ := strconv.ParseFloat(fmt.Sprintf("%v", operador_der.Valor), 64)
				//SE RETORNA SIMBOLO
				return simbolos.Valor{Tipo: simbolos.FLOAT, Valor: math.Pow(val1, val2)}
			} else {
				//ERROR
				t := time.Now()
				reportes.Nerror("No es posible realizar la potencia, tipos no compatibles", ent.Nombre_Entorno, strconv.Itoa(pot.Linea), strconv.Itoa(pot.Columna), t.Format("2006-01-02 15:04:05"))
			}
		}
	}

	//SINO CUMPLE SE RETORNA SIMBOLO DE TIPO NULL CON VALOR NULL
	return simbolos.Valor{Tipo: simbolos.NULL, Valor: -1}
}

func (pot Potencia) Compilar_Expresion(ent *entorno.Entorno, gen *generador.Generador_C3D) simbolos.ValoresC3D {

	//MATRIZ DOMINANTE POTENCIA
	//			INTEGER				FLOAT			  STRING		    CHAR		      BOOL			 &STRING		NULL
	potencia_dominante := [7][7]simbolos.TipoExpresion{
		//INTEGER
		{simbolos.INTEGER, simbolos.FLOAT, simbolos.NULL, simbolos.INTEGER, simbolos.NULL, simbolos.NULL, simbolos.NULL},
		//FLOAT
		{simbolos.FLOAT, simbolos.FLOAT, simbolos.NULL, simbolos.FLOAT, simbolos.NULL, simbolos.NULL, simbolos.NULL},
		//BOOLEAN
		{simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL},
		//CHAR
		{simbolos.INTEGER, simbolos.FLOAT, simbolos.NULL, simbolos.CHAR, simbolos.NULL, simbolos.NULL, simbolos.NULL},
		//TEXTO
		{simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL},
		//&TEXTO
		{simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL},
		//NULL
		{simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL},
	}

	//SE OBTIENE EL VALOR DEL LOS OPERADOES IZQUIERDA Y DERECHO
	var operador_izq simbolos.ValoresC3D = pot.Oi.Compilar_Expresion(ent, gen)
	var operador_der simbolos.ValoresC3D = pot.Od.Compilar_Expresion(ent, gen)

	//SE VALIDA EL TIPO DE LA POTENCIA
	var tipo_potencia simbolos.TipoExpresion
	nuevo_temporal := gen.Crear_temporal()
	//SE VALIDA QUE TIPO DE POTENCIA ES SI ENTERA O DECIMAL
	switch pot.Tipo {
	case "i64":
		{
			//SE OBTIENE EL TIPO DE LA POTENCIA
			tipo_potencia = potencia_dominante[operador_izq.Tipo][operador_der.Tipo]
			//SE VERIFICA QUE SEA ENTERA
			if tipo_potencia == simbolos.INTEGER {
				gen.Agregar_Logica(nuevo_temporal + " = pow(" + operador_izq.Valor + ", " + operador_der.Valor + "); \t\t//OPERACION POTENCIA")
				return simbolos.ValoresC3D{Valor: nuevo_temporal, EsTemporal: true, Tipo: simbolos.FLOAT, Label_verdadera: "", Label_false: ""}
			} else {
				//ERROR
				gen.AgregarError("ERROR-TIPOS--POTENCIA", strconv.Itoa(pot.Linea), strconv.Itoa(pot.Columna))
			}
		}
	case "f64":
		{
			//SE OBTIENE EL TIPO DE LA POTENCIA
			tipo_potencia = potencia_dominante[operador_izq.Tipo][operador_der.Tipo]
			//SE VERIFICA QUE SEA DECIMAL
			if tipo_potencia == simbolos.FLOAT {
				gen.Agregar_Logica(nuevo_temporal + " = pow(" + operador_izq.Valor + ", " + operador_der.Valor + "); \t\t//OPERACION POTENCIA")
				return simbolos.ValoresC3D{Valor: nuevo_temporal, EsTemporal: true, Tipo: tipo_potencia, Label_verdadera: "", Label_false: ""}
			} else {
				//ERROR
				gen.AgregarError("ERROR-TIPOS--POTENCIA", strconv.Itoa(pot.Linea), strconv.Itoa(pot.Columna))
			}
		}
	}

	//SINO CUMPLE SE RETORNA SIMBOLO DE TIPO NULL CON VALOR NULL
	return simbolos.ValoresC3D{Valor: "0", EsTemporal: false, Tipo: simbolos.INTEGER, Label_verdadera: "", Label_false: ""}
}
