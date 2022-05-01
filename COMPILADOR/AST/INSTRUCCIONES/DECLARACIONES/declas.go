package declaraciones

import (
	entorno "OLC2-PROYECTO2/COMPILADOR/ENTORNO"
	generador "OLC2-PROYECTO2/COMPILADOR/GENERADOR"
	interfaces "OLC2-PROYECTO2/COMPILADOR/INTERFACES"
	reportes "OLC2-PROYECTO2/COMPILADOR/REPORTES"
	simbolos "OLC2-PROYECTO2/COMPILADOR/SIMBOLOS"
	"fmt"
	"strconv"
	"time"

	"github.com/colegno/arraylist"
)

type AsigStruct struct {
	Nombre string
	Valor  interfaces.Expresion
}

func Nasigstruct(nom string, val interfaces.Expresion) AsigStruct {
	as := AsigStruct{Nombre: nom, Valor: val}
	return as
}

type Declarar struct {
	Mutable             bool
	Identificador       string
	IdentificadorStruct string
	Valor_exp           *arraylist.List
	Linea               int
	Columna             int
}

func Ndeclarar(mut bool, id string, id2 string, vexp *arraylist.List, l int, c int) Declarar {
	decla := Declarar{Mutable: mut, Identificador: id, IdentificadorStruct: id2, Valor_exp: vexp, Linea: l, Columna: c}
	return decla
}

func (decla Declarar) Ejecutar_Instruccion(ent *entorno.Entorno, ent2 *entorno.Entorno) interface{} {
	if !ent.Existe_ArreVect(decla.Identificador) {
		if !ent.Existe_Variable(decla.Identificador) {
			if ent.Existe_Structs(decla.IdentificadorStruct) {
				simbstruct := ent.Obtener_Struct(decla.IdentificadorStruct)
				if decla.Valor_exp.Len() == simbstruct.L_Atributos.Len() {
					entostruct := ent2.Entorno_Anterior.Retornar_Entorno("struct_" + decla.IdentificadorStruct)
					conta := 0
					for i := 0; i < decla.Valor_exp.Len(); i++ {
						act := decla.Valor_exp.GetValue(i).(AsigStruct)
						if entostruct.Existe_Variable(act.Nombre) {
							varact := entostruct.Obtener_Variable(act.Nombre)
							valor := act.Valor.Ejecutar_Expresion(ent)
							if varact.TipoVariable == valor.Tipo {
								conta = conta + 1
							} else {
								t := time.Now()
								reportes.Nerror("No coinciden los tipos", ent.Nombre_Entorno, strconv.Itoa(decla.Linea), strconv.Itoa(decla.Columna), t.Format("2006-01-02 15:04:05"))
							}
						} else {
							t := time.Now()
							reportes.Nerror("Atributo no declarado en Struct.", ent.Nombre_Entorno, strconv.Itoa(decla.Linea), strconv.Itoa(decla.Columna), t.Format("2006-01-02 15:04:05"))
						}
					}

					if conta == decla.Valor_exp.Len() {
						simbvar := simbolos.Simbolo_Vars{Mutable: decla.Mutable,
							TipoVariable:   simbolos.STRUCT,
							NombreVariable: decla.Identificador,
							ValorVariable:  simbstruct.Identificador,
							PosicionTabla:  ent.Posicion,
							Linea:          decla.Linea,
							Columna:        decla.Columna}
						if ent.Nombre_Entorno != ent2.Nombre_Entorno {
							ent2.Posicion = ent2.Posicion - 1
							ent.Guardar_Variable(decla.Identificador, simbvar)
							ent2.Guardar_Variable(decla.Identificador, simbvar)
						} else {
							ent.Guardar_Variable(decla.Identificador, simbvar)
							ent2.Guardar_Variable(decla.Identificador, simbvar)
						}

						reportes.ReporteSimbolos(decla.Identificador, "variable--"+reportes.ReportObteniendoSimbolos(simbolos.STRUCT), ent.Nombre_Entorno, "--", strconv.Itoa(decla.Linea), strconv.Itoa(decla.Columna))
					}
				} else {
					t := time.Now()
					reportes.Nerror("La cant. de atributos no coincide.", ent.Nombre_Entorno, strconv.Itoa(decla.Linea), strconv.Itoa(decla.Columna), t.Format("2006-01-02 15:04:05"))
				}
			} else {
				t := time.Now()
				reportes.Nerror("No existe struct con ese nombre.", ent.Nombre_Entorno, strconv.Itoa(decla.Linea), strconv.Itoa(decla.Columna), t.Format("2006-01-02 15:04:05"))
			}
		} else {
			t := time.Now()
			reportes.Nerror("No se puede declarar variable con el mismo nombre.", ent.Nombre_Entorno, strconv.Itoa(decla.Linea), strconv.Itoa(decla.Columna), t.Format("2006-01-02 15:04:05"))
		}
	} else {
		t := time.Now()
		reportes.Nerror("Vectores/Arreglos con el mismo nombre en el entorno.", ent.Nombre_Entorno, strconv.Itoa(decla.Linea), strconv.Itoa(decla.Columna), t.Format("2006-01-02 15:04:05"))
	}
	return 0
}

func (decla Declarar) Compilar_Instruccion(ent *entorno.Entorno, gen *generador.Generador_C3D) interface{} {
	gen.Agregar_Comentario("COMENZANDO DECLARACION")
	if ent.Existe_Variable(decla.Identificador) {
		simbdecla := ent.Obtener_Variable(decla.Identificador)
		if ent.Existe_Structs(decla.IdentificadorStruct) {
			simbstruct := ent.Obtener_Struct(decla.IdentificadorStruct)
			pos := simbdecla.PosicionTabla
			tam := simbstruct.Tamaño
			tempo1 := gen.Crear_temporal()
			tempo2 := gen.Crear_temporal()
			gen.Agregar_Comentario("DECLARACION STRUCT " + decla.IdentificadorStruct)
			gen.Agregar_Logica(tempo1 + " = SP + " + strconv.Itoa(pos) + ";\t\t// POSICION:" + decla.Identificador)
			gen.Agregar_Stack(tempo1, "HP")
			gen.Agregar_Logica(tempo2 + " = HP;")
			gen.Agregar_Logica("HP = HP + " + strconv.Itoa(tam) + ";")
			gen.Agregar_Comentario("ASIGNACION STRUCT " + decla.IdentificadorStruct)
			entostruct := ent.Entorno_Anterior.Retornar_Entorno("struct_" + decla.IdentificadorStruct)
			tempo3 := gen.Crear_temporal()
			for i := 0; i < decla.Valor_exp.Len(); i++ {
				act := decla.Valor_exp.GetValue(i).(AsigStruct)
				if entostruct.Existe_Variable(act.Nombre) {
					varact := entostruct.Obtener_Variable(act.Nombre)
					valor := act.Valor.Compilar_Expresion(ent, gen)
					if varact.TipoVariable == simbolos.TEXTO || varact.TipoVariable == simbolos.YTEXTO {
						gen.Agregar_Logica(tempo3 + " = " + tempo2 + " + " + strconv.Itoa(varact.PosicionTabla) + ";")
						gen.Agregar_Logica("HEAP[(int)" + tempo3 + "] = HP;")
						for _, txt := range valor.Valor {
							f := int(txt)
							gen.Agregar_Logica("HEAP[(int)HP] = " + fmt.Sprintf("%v", f) + "; //LETRA-> " + string(txt) + ";")
							gen.Agregar_Logica("HP = HP + 1;")
						}
						gen.Agregar_Logica("HEAP[(int)HP] = -1;\nHP = HP + 1;")
					} else {
						gen.Agregar_Logica(tempo3 + " = " + tempo2 + " + " + strconv.Itoa(varact.PosicionTabla) + ";")
						gen.Agregar_Logica("HEAP[(int)" + tempo3 + "] = " + valor.Valor + ";")
					}
				}
			}
		}
	}
	gen.Agregar_Comentario("FINALIZANDO DECLARACION")
	return 0
}
