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

type OpModulo struct {
	OpIz    interfaces.Expresion
	OpDe    interfaces.Expresion
	Linea   int
	Columna int
}

func Nopmodulo(iz interfaces.Expresion, der interfaces.Expresion, l int, c int) OpModulo {
	modulo := OpModulo{OpIz: iz, OpDe: der, Linea: l, Columna: c}
	return modulo
}

func (modulo OpModulo) Ejecutar_Expresion(ent *entorno.Entorno) simbolos.Valor {
	//INTERGER			FLOAT			BOOLEAN			CHAR		TEXTO			YTEXTO			NULL
	modulo_dominante := [7][7]simbolos.TipoExpresion{
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

	var operador_izquierda simbolos.Valor = modulo.OpIz.Ejecutar_Expresion(ent)
	var operador_derecha simbolos.Valor = modulo.OpDe.Ejecutar_Expresion(ent)

	var tipo_dominante simbolos.TipoExpresion = modulo_dominante[operador_izquierda.Tipo][operador_derecha.Tipo]
	//SE VERIFICA SI ES ENTERO
	if tipo_dominante == simbolos.INTEGER {
		//SE HACE CONVERSIONES
		val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", operador_izquierda.Valor), 64)
		val2, _ := strconv.ParseFloat(fmt.Sprintf("%v", operador_derecha.Valor), 64)
		i := int(val1)
		j := int(val2)
		if j != 0 {
			//SE RETORNA SIMBOLO
			return simbolos.Valor{Tipo: tipo_dominante, Valor: i % j}
		} else {
			t := time.Now()
			reportes.Nerror("No es posible dividir entre 0", ent.Nombre_Entorno, strconv.Itoa(modulo.Linea), strconv.Itoa(modulo.Columna), t.Format("2006-01-02 15:04:05"))
		}
		//SE VERIFICA SI ES DECIMAL
	} else if tipo_dominante == simbolos.FLOAT {
		//SE HACE CONVERSIONES
		val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", operador_izquierda.Valor), 64)
		val2, _ := strconv.ParseFloat(fmt.Sprintf("%v", operador_derecha.Valor), 64)
		if val2 != 0 {
			//SE RETORNA SIMBOLO
			return simbolos.Valor{Tipo: tipo_dominante, Valor: math.Mod(val1, val2)}
		} else {
			t := time.Now()
			reportes.Nerror("No es posible dividir entre 0", ent.Nombre_Entorno, strconv.Itoa(modulo.Linea), strconv.Itoa(modulo.Columna), t.Format("2006-01-02 15:04:05"))
		}
	} else {
		//ERROR
		t := time.Now()
		reportes.Nerror("No es posible realizar la opreacion: Modulo", ent.Nombre_Entorno, strconv.Itoa(modulo.Linea), strconv.Itoa(modulo.Columna), t.Format("2006-01-02 15:04:05"))
	} //SE VERIFICA SI ES ENTERO
	if tipo_dominante == simbolos.INTEGER {
		//SE HACE CONVERSIONES
		val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", operador_izquierda.Valor), 64)
		val2, _ := strconv.ParseFloat(fmt.Sprintf("%v", operador_derecha.Valor), 64)
		i := int(val1)
		j := int(val2)
		if j != 0 {
			//SE RETORNA SIMBOLO
			return simbolos.Valor{Tipo: tipo_dominante, Valor: i % j}
		} else {
			t := time.Now()
			reportes.Nerror("No es posible dividir entre 0", ent.Nombre_Entorno, strconv.Itoa(modulo.Linea), strconv.Itoa(modulo.Columna), t.Format("2006-01-02 15:04:05"))
		}
		//SE VERIFICA SI ES DECIMAL
	} else if tipo_dominante == simbolos.FLOAT {
		//SE HACE CONVERSIONES
		val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", operador_izquierda.Valor), 64)
		val2, _ := strconv.ParseFloat(fmt.Sprintf("%v", operador_derecha.Valor), 64)
		if val2 != 0 {
			//SE RETORNA SIMBOLO
			return simbolos.Valor{Tipo: tipo_dominante, Valor: math.Mod(val1, val2)}
		} else {
			t := time.Now()
			reportes.Nerror("No es posible dividir entre 0", ent.Nombre_Entorno, strconv.Itoa(modulo.Linea), strconv.Itoa(modulo.Columna), t.Format("2006-01-02 15:04:05"))
		}
	} else {
		//ERROR
		t := time.Now()
		reportes.Nerror("No es posible realizar la opreacion: Modulo", ent.Nombre_Entorno, strconv.Itoa(modulo.Linea), strconv.Itoa(modulo.Columna), t.Format("2006-01-02 15:04:05"))
	}
	//NO SE ACEPTO LOS TIPOS
	return simbolos.Valor{Tipo: simbolos.NULL, Valor: -1}
}

func (modulo OpModulo) Compilar_Expresion(ent *entorno.Entorno, gen *generador.Generador_C3D) simbolos.ValoresC3D {
	//INTERGER			FLOAT			BOOLEAN			CHAR		TEXTO			YTEXTO			NULL
	modulo_dominante := [7][7]simbolos.TipoExpresion{
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

	var operador_izquierda simbolos.ValoresC3D = modulo.OpIz.Compilar_Expresion(ent, gen)
	var operador_derecha simbolos.ValoresC3D = modulo.OpDe.Compilar_Expresion(ent, gen)

	var tipo_dominante simbolos.TipoExpresion = modulo_dominante[operador_izquierda.Tipo][operador_derecha.Tipo]

	nuevo_temporal := gen.Crear_temporal()

	if tipo_dominante == simbolos.INTEGER {
		gen.Agregar_Expression(nuevo_temporal, operador_izquierda.Valor, "%", operador_derecha.Valor)
		return simbolos.ValoresC3D{Valor: nuevo_temporal, EsTemporal: true, Tipo: tipo_dominante, Label_verdadera: "", Label_false: ""}
	} else if tipo_dominante == simbolos.FLOAT {
		gen.Agregar_Expression(nuevo_temporal, operador_izquierda.Valor, "%", operador_derecha.Valor)
		return simbolos.ValoresC3D{Valor: nuevo_temporal, EsTemporal: true, Tipo: tipo_dominante, Label_verdadera: "", Label_false: ""}
	} else {
		fmt.Println("ERROR, tipos")
	}
	return simbolos.ValoresC3D{Valor: "0", EsTemporal: false, Tipo: simbolos.INTEGER, Label_verdadera: "", Label_false: ""}
}
