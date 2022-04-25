package aritmetica

import (
	entorno "OLC2-PROYECTO2/COMPILADOR/ENTORNO"
	generador "OLC2-PROYECTO2/COMPILADOR/GENERADOR"
	interfaces "OLC2-PROYECTO2/COMPILADOR/INTERFACES"
	reportes "OLC2-PROYECTO2/COMPILADOR/REPORTES"
	simbolos "OLC2-PROYECTO2/COMPILADOR/SIMBOLOS"
	"fmt"
	"strconv"
	"time"
)

type OpSuma struct {
	OpIzquierda     interfaces.Expresion
	OperadorDerecha interfaces.Expresion
	Linea           int
	Columna         int
}

func Nopsuma(oi interfaces.Expresion, od interfaces.Expresion, linea int, colu int) OpSuma {
	suma := OpSuma{OpIzquierda: oi, OperadorDerecha: od, Linea: linea, Columna: colu}
	return suma
}

func (suma OpSuma) Ejecutar_Expresion(ent *entorno.Entorno) simbolos.Valor {
	//INTERGER			FLOAT			BOOLEAN			CHAR		TEXTO			YTEXTO			NULL
	suma_dominante := [7][7]simbolos.TipoExpresion{
		//INTEGER
		{simbolos.INTEGER, simbolos.FLOAT, simbolos.NULL, simbolos.INTEGER, simbolos.NULL, simbolos.NULL, simbolos.NULL},
		//FLOAT
		{simbolos.FLOAT, simbolos.FLOAT, simbolos.NULL, simbolos.FLOAT, simbolos.NULL, simbolos.NULL, simbolos.NULL},
		//BOOLEAN
		{simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL},
		//CHAR
		{simbolos.INTEGER, simbolos.FLOAT, simbolos.NULL, simbolos.CHAR, simbolos.NULL, simbolos.NULL, simbolos.NULL},
		//TEXTO
		{simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.TEXTO, simbolos.YTEXTO, simbolos.NULL},
		//&TEXTO
		{simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.YTEXTO, simbolos.NULL},
		//NULL
		{simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL},
	}

	var operador_izquierda simbolos.Valor = suma.OpIzquierda.Ejecutar_Expresion(ent)
	var operador_derecha simbolos.Valor = suma.OperadorDerecha.Ejecutar_Expresion(ent)

	var tipo_dominante simbolos.TipoExpresion = suma_dominante[operador_izquierda.Tipo][operador_derecha.Tipo]

	//SE VERIFICA SI ES ENTERO
	if tipo_dominante == simbolos.INTEGER {
		//SE HACE LAS CONVERSIONES
		val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", operador_izquierda.Valor), 64)
		val2, _ := strconv.ParseFloat(fmt.Sprintf("%v", operador_derecha.Valor), 64)
		i := int(val1)
		j := int(val2)
		//SE RETORNA SIMBOLO
		return simbolos.Valor{Tipo: tipo_dominante, Valor: i + j}
		//SE VERIFICA SI ES DECIMAL
	} else if tipo_dominante == simbolos.FLOAT {
		//SE HACE LAS CONVERSIONES
		val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", operador_izquierda.Valor), 64)
		val2, _ := strconv.ParseFloat(fmt.Sprintf("%v", operador_derecha.Valor), 64)
		//SE RETORNA SIMBOLO
		return simbolos.Valor{Tipo: tipo_dominante, Valor: val1 + val2}
		//SE VERIFICA SI ES STRING
	} else if tipo_dominante == simbolos.TEXTO {
		//SE HACE LAS CONVERSIONES
		l1 := fmt.Sprintf("%v", operador_izquierda.Valor)
		l2 := fmt.Sprintf("%v", operador_derecha.Valor)
		//SE RETORNA SIMBOLO
		return simbolos.Valor{Tipo: tipo_dominante, Valor: l1 + l2}
		//SE VERIFIA SI EL OPIZ ES CHAR Y EL OPDE ES STRING
	} else if operador_izquierda.Tipo == simbolos.CHAR && operador_derecha.Tipo == simbolos.TEXTO {
		//SE HACE LAS CONVERSIONES
		j := string(rune(operador_izquierda.Valor.(int32)))
		h := fmt.Sprintf("%v", operador_derecha.Valor)
		tipo_dominante = simbolos.TEXTO
		//SE RETORNA SIMBOLO
		return simbolos.Valor{Tipo: tipo_dominante, Valor: j + h}
		//SE VERIFICA SI OPIZ ES STRING Y EL OPDE ES CHAR
	} else if operador_izquierda.Tipo == simbolos.TEXTO && operador_derecha.Tipo == simbolos.CHAR {
		//SE HACE LAS CONVERSIONES
		j := fmt.Sprintf("%v", operador_izquierda.Valor)
		h := string(rune(operador_derecha.Valor.(int32)))
		tipo_dominante = simbolos.TEXTO
		//SE RETORNA SIMBOLO
		return simbolos.Valor{Tipo: tipo_dominante, Valor: j + h}
	} else {
		//ERROR
		t := time.Now()
		reportes.Nerror("No es posible realizar la operacion: Suma", ent.Nombre_Entorno, strconv.Itoa(suma.Linea), strconv.Itoa(suma.Columna), t.Format("2006-01-02 15:04:05"))
	}

	return simbolos.Valor{Tipo: simbolos.NULL, Valor: -1}
}

func (suma OpSuma) Compilar_Expresion(ent *entorno.Entorno, gen *generador.Generador_C3D) simbolos.ValoresC3D {
	//INTERGER			FLOAT			BOOLEAN			CHAR		TEXTO			YTEXTO			NULL
	suma_dominante := [7][7]simbolos.TipoExpresion{
		//INTEGER
		{simbolos.INTEGER, simbolos.FLOAT, simbolos.NULL, simbolos.INTEGER, simbolos.NULL, simbolos.NULL, simbolos.NULL},
		//FLOAT
		{simbolos.FLOAT, simbolos.FLOAT, simbolos.NULL, simbolos.FLOAT, simbolos.NULL, simbolos.NULL, simbolos.NULL},
		//BOOLEAN
		{simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL},
		//CHAR
		{simbolos.INTEGER, simbolos.FLOAT, simbolos.NULL, simbolos.CHAR, simbolos.NULL, simbolos.NULL, simbolos.NULL},
		//TEXTO
		{simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.YTEXTO, simbolos.NULL},
		//&TEXTO
		{simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.YTEXTO, simbolos.NULL},
		//NULL
		{simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL},
	}

	var operador_izquierda simbolos.ValoresC3D = suma.OpIzquierda.Compilar_Expresion(ent, gen)
	var operador_derecha simbolos.ValoresC3D = suma.OperadorDerecha.Compilar_Expresion(ent, gen)

	var tipo_dominante simbolos.TipoExpresion = suma_dominante[operador_izquierda.Tipo][operador_derecha.Tipo]

	nuevo_temporal := gen.Crear_temporal()

	if tipo_dominante == simbolos.INTEGER {
		gen.Agregar_Expression(nuevo_temporal, operador_izquierda.Valor, "+", operador_derecha.Valor)
		return simbolos.ValoresC3D{Valor: nuevo_temporal, EsTemporal: true, Tipo: tipo_dominante, Label_verdadera: "", Label_false: ""}
	} else if tipo_dominante == simbolos.FLOAT {
		gen.Agregar_Expression(nuevo_temporal, operador_izquierda.Valor, "+", operador_derecha.Valor)
		return simbolos.ValoresC3D{Valor: nuevo_temporal, EsTemporal: true, Tipo: tipo_dominante, Label_verdadera: "", Label_false: ""}
	} else {
		fmt.Println("ERROR, tipos")
		gen.AgregarError("ERROR-TIPOS--SUMA", strconv.Itoa(suma.Linea), strconv.Itoa(suma.Columna))
	}
	return simbolos.ValoresC3D{Valor: "0", EsTemporal: false, Tipo: simbolos.INTEGER, Label_verdadera: "", Label_false: ""}
}
