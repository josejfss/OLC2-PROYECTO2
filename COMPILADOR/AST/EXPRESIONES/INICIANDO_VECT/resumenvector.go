package iniciandovect

import (
	entorno "OLC2-PROYECTO2/COMPILADOR/ENTORNO"
	generador "OLC2-PROYECTO2/COMPILADOR/GENERADOR"
	interfaces "OLC2-PROYECTO2/COMPILADOR/INTERFACES"
	simbolos "OLC2-PROYECTO2/COMPILADOR/SIMBOLOS"
	"fmt"

	"github.com/colegno/arraylist"
)

type ResumenVector struct {
	Exp1 interfaces.Expresion
	Exp2 interfaces.Expresion
}

func Nvalvect(e1 interfaces.Expresion, e2 interfaces.Expresion) ResumenVector {
	return ResumenVector{Exp1: e1, Exp2: e2}
}

func (ini ResumenVector) Ejecutar_Expresion(ent *entorno.Entorno) simbolos.Valor {
	data, tipo := ini.Obteniendo_Valores(ent)

	return simbolos.Valor{Valor: data, Tipo: tipo}
}

func (ini ResumenVector) Obteniendo_Valores(ent *entorno.Entorno) (interface{}, simbolos.TipoExpresion) {
	//tipo := interfaces.NULL
	data := arraylist.New()

	ex1 := ini.Exp1.Ejecutar_Expresion(ent)
	ex2 := ini.Exp2.Ejecutar_Expresion(ent)

	for i := 0; i < ex2.Valor.(int); i++ {
		//tipo = ex1.Tipo
		data.Add(ex1)
	}

	lista_dimensiones := arraylist.New()
	lista_dimensiones.Add(data.Len())
	lista_capacidades := arraylist.New()
	lista_capacidades.Add(data.Len() + 1)
	tipo_datos := simbolos.NULL

	fin := make([]interface{}, data.Len())

	for i := 0; i < data.Len(); i++ {
		dato := data.GetValue(i)
		valor_datico := dato.(simbolos.Valor)

		if valor_datico.Tipo == simbolos.NULL {
			fmt.Print("ERROR")
		}

		if valor_datico.Tipo != simbolos.VECTOR {
			if i == 0 {
				tipo_datos = valor_datico.Tipo
			} else {
				if tipo_datos != valor_datico.Tipo {
					fmt.Println("ERROR")
				}
			}
			fin[i] = valor_datico.Valor
			continue
		}

		valor_objeto := valor_datico.Valor.(simbolos.Valor)
		objeto_vect := valor_objeto.Valor.(simbolos.Simbolo_ArreVect)

		if i == 0 {
			tipo_datos = valor_objeto.Tipo
			lista_dimensiones.Add(objeto_vect.Lintdim.ToArray())
			lista_capacidades.Add(objeto_vect.Lintcap.ToArray())
		} else {
			if tipo_datos != valor_objeto.Tipo {
				fmt.Println("ERROR")
			}
			lista_dimensiones.Add(objeto_vect.Lintdim.ToArray())
			lista_capacidades.Add(objeto_vect.Lintcap.ToArray())
		}
		fin[i] = objeto_vect.ValorArreVect

	}

	vectorsito := simbolos.Simbolo_ArreVect{Mutable: false, TipoVect: tipo_datos, Nombre: "", ValorArreVect: fin, Lintdim: lista_dimensiones, Lintcap: lista_capacidades, Linea: 0, Columna: 0}
	objtval := simbolos.Valor{
		Valor: vectorsito,
		Tipo:  tipo_datos,
	}

	return objtval, simbolos.VECTOR
}

func (ini ResumenVector) Compilar_Expresion(ent *entorno.Entorno, gen *generador.Generador_C3D) simbolos.ValoresC3D {
	return simbolos.ValoresC3D{Valor: "0", EsTemporal: false, Tipo: simbolos.INTEGER, Label_verdadera: "", Label_false: ""}
}
