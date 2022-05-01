package accesostruct

import (
	entorno "OLC2-PROYECTO2/COMPILADOR/ENTORNO"
	generador "OLC2-PROYECTO2/COMPILADOR/GENERADOR"
	simbolos "OLC2-PROYECTO2/COMPILADOR/SIMBOLOS"
	"fmt"
	"strconv"
)

type AccesoStruct struct {
	NombreStruct string
	NombreId     string
	Linea        int
	Columna      int
}

func NaccesoStruct(nom1 string, nom2 string, lin int, col int) AccesoStruct {
	access := AccesoStruct{NombreStruct: nom1, NombreId: nom2, Linea: lin, Columna: col}
	return access
}

func (access AccesoStruct) Ejecutar_Expresion(ent *entorno.Entorno) simbolos.Valor {
	return simbolos.Valor{Tipo: simbolos.NULL, Valor: -1}
}

func (access AccesoStruct) Compilar_Expresion(ent *entorno.Entorno, gen *generador.Generador_C3D) simbolos.ValoresC3D {

	if ent.Existe_Variable(access.NombreStruct) {
		gen.Agregar_Comentario("INICIO ACCESO AL STRUCT")
		simbstruct := ent.Obtener_Variable(access.NombreStruct)
		temp1 := gen.Crear_temporal()
		temp2 := gen.Crear_temporal()
		gen.Agregar_Logica(temp1 + " = SP + " + strconv.Itoa(simbstruct.PosicionTabla) + ";")
		gen.Agregar_Logica(temp2 + " = STACK[(int)" + temp1 + "];")
		if ent.Existe_Structs(simbstruct.ValorVariable.(string)) {
			entstruct := ent.Entorno_Anterior.Retornar_Entorno("struct_" + simbstruct.ValorVariable.(string))
			if entstruct.Existe_Variable(access.NombreId) {
				varstruct := entstruct.Obtener_Variable(access.NombreId)
				temp3 := gen.Crear_temporal()
				gen.Agregar_Logica(temp3 + " = " + temp2 + " + " + strconv.Itoa(varstruct.PosicionTabla) + ";")
				gen.Agregar_Comentario("FIN ACCESO AL STRUCT")
				temp4 := gen.Crear_temporal()
				infoid := access.NombreId + ": "
				imprimir := ""
				for _, txt := range infoid {
					f := int(txt)
					imprimir += "printf(\"%c\", (char)" + fmt.Sprintf("%v", f) + "); //LETRA-> " + string(txt) + "\n"
					imprimir += "HP = HP + 1;\n"
				}
				if varstruct.TipoVariable == simbolos.TEXTO || varstruct.TipoVariable == simbolos.YTEXTO {
					temp5 := gen.Crear_temporal()
					etiqueta_entrada := gen.Crear_label()
					gen.Eliminar_label(etiqueta_entrada)
					etiqueta_salida := gen.Crear_label()
					gen.Eliminar_label(etiqueta_salida)
					imprimir += temp4 + " = HEAP[(int)" + temp3 + "];\n"
					imprimir += etiqueta_entrada + ":\n"
					imprimir += temp5 + " = HEAP[(int)" + temp4 + "];\n"
					imprimir += "if ( " + temp5 + " == -1) { goto " + etiqueta_salida + "; }\n"
					imprimir += "printf(\"%c\", (char)" + temp5 + ");\n"
					imprimir += temp4 + " = " + temp4 + " + 1;\ngoto " + etiqueta_entrada + ";\n"
					imprimir += etiqueta_salida + ":\n"
					return simbolos.ValoresC3D{Valor: imprimir, EsTemporal: false, Tipo: simbolos.STRUCT, Label_verdadera: "", Label_false: ""}
				} else if varstruct.TipoVariable == simbolos.INTEGER {
					imprimir += temp4 + " = HEAP[(int)" + temp3 + "];\n"
					imprimir += "printf(\"%d\",(int)" + temp4 + ");\n"
					return simbolos.ValoresC3D{Valor: imprimir, EsTemporal: false, Tipo: simbolos.STRUCT, Label_verdadera: "", Label_false: ""}
				} else if varstruct.TipoVariable == simbolos.FLOAT {
					imprimir += temp4 + " = HEAP[(int)" + temp3 + "];\n"
					imprimir += "printf(\"%f\"" + temp4 + ");\n"
					return simbolos.ValoresC3D{Valor: imprimir, EsTemporal: false, Tipo: simbolos.STRUCT, Label_verdadera: "", Label_false: ""}
				}
			}
		}
	}
	return simbolos.ValoresC3D{Valor: "0", EsTemporal: false, Tipo: simbolos.INTEGER, Label_verdadera: "", Label_false: ""}
}
