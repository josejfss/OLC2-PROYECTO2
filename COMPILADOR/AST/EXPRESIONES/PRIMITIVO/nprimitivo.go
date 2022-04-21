package primitivo

import (
	entorno "OLC2-PROYECTO2/COMPILADOR/ENTORNO"
	generador "OLC2-PROYECTO2/COMPILADOR/GENERADOR"
	simbolos "OLC2-PROYECTO2/COMPILADOR/SIMBOLOS"
	"fmt"
)

type Dato_Primitivo struct {
	Valor interface{}
	Tipo  simbolos.TipoExpresion
}

func Nuevo_Dato_Primitivo(val interface{}, tipo simbolos.TipoExpresion) Dato_Primitivo {
	datpri := Dato_Primitivo{val, tipo}
	return datpri
}

func (dp Dato_Primitivo) Ejecutar_Expresion(ent *entorno.Entorno) simbolos.Valor {
	return simbolos.Valor{
		Tipo:  dp.Tipo,
		Valor: dp.Valor,
	}
}

func (dp Dato_Primitivo) Compilar_Expresion(ent *entorno.Entorno, gen *generador.Generador_C3D) simbolos.ValoresC3D {
	if dp.Valor == true {
		temp := gen.Crear_temporal()
		verdad := gen.Crear_label()
		falso := gen.Crear_label()
		boolean_true := temp + " = 1;\n"
		boolean_true += "if (" + temp + " == 1) { goto " + verdad + "; }\ngoto " + falso + ";"
		gen.Agregar_Logica(boolean_true)
		return simbolos.ValoresC3D{Valor: boolean_true, EsTemporal: false, Tipo: dp.Tipo, Label_verdadera: verdad, Label_false: falso}
	} else if dp.Valor == false {
		temp := gen.Crear_temporal()
		verdad := gen.Crear_label()
		falso := gen.Crear_label()
		boolean_false := temp + " = 0;\n"
		boolean_false += "if (" + temp + " == 0) { goto " + verdad + "; }\ngoto " + falso + ";"
		gen.Agregar_Logica(boolean_false)
		return simbolos.ValoresC3D{Valor: boolean_false, EsTemporal: false, Tipo: dp.Tipo, Label_verdadera: verdad, Label_false: falso}
	} else {
		return simbolos.ValoresC3D{
			Valor:           fmt.Sprintf("%v", dp.Valor),
			EsTemporal:      false,
			Tipo:            dp.Tipo,
			Label_verdadera: "",
			Label_false:     "",
		}
	}
	// } else if dp.Tipo == simbolos.YTEXTO {
	// 	tempo := gen.Crear_temporal()
	// 	gen.Agregar_Logica(tempo + " = HP;")
	// 	vals := fmt.Sprintf("%v", dp.Valor)
	// 	for _, txt := range vals {
	// 		f := int(txt)
	// 		gen.Agregar_Logica("HEAP[int(HP)] = " + fmt.Sprintf("%v", f) + "; //LETRA-> " + string(txt) + "\n")
	// 		gen.Agregar_Logica("HP = HP + 1;\n")
	// 	}
	// 	gen.Agregar_Logica("HEAP[int(HP)] = 10;\nHP = HP + 1;\n")
	// 	gen.Agregar_Logica("HEAP[int(HP)] = -1;\nHP = HP + 1;\n")
	// 	return simbolos.ValoresC3D{Valor: tempo, EsTemporal: true, Tipo: dp.Tipo, Label_verdadera: "", Label_false: ""}
	// }

}
