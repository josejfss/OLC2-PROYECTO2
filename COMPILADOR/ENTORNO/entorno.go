package entorno

import (
	simbolos "OLC2-PROYECTO2/COMPILADOR/SIMBOLOS"

	"github.com/colegno/arraylist"
)

type Entorno struct {
	Nombre_Entorno   string
	Entorno_Anterior *Entorno
	Lista_Entorno    *arraylist.List
	Tabla_Variables  map[string]simbolos.Simbolo_Vars
	Tabla_Variables2 map[string]simbolos.Simbolo_Vars
	Tabla_ArreVect   map[string]simbolos.Simbolo_ArreVect
	Tabla_Funcines   map[string]simbolos.Simbolo_Funciones
	Tabla_Struct     map[string]simbolos.Simbolo_Struct
	Tabla_Modulos    map[string]simbolos.Simbolos_Modulos
	ListaTodo        *arraylist.List
	ContadorTodo     int
	Posicion         int
	Tamaño           int
}

func Nuevo_Entorno(nombre string, anterior *Entorno) *Entorno {
	ent := Entorno{Nombre_Entorno: nombre, Entorno_Anterior: anterior}
	ent.Lista_Entorno = arraylist.New()
	ent.Tabla_Variables = make(map[string]simbolos.Simbolo_Vars)
	ent.Tabla_Variables2 = make(map[string]simbolos.Simbolo_Vars)
	ent.Tabla_ArreVect = make(map[string]simbolos.Simbolo_ArreVect)
	ent.Tabla_Funcines = make(map[string]simbolos.Simbolo_Funciones)
	ent.Tabla_Struct = make(map[string]simbolos.Simbolo_Struct)
	ent.Tabla_Modulos = make(map[string]simbolos.Simbolos_Modulos)
	ent.ListaTodo = arraylist.New()
	ent.ContadorTodo = 0
	ent.Posicion = 0
	ent.Tamaño = 0
	// if anterior != nil {
	// 	anterior.Lista_Entorno.Add(ent)
	// }
	return &ent
}

func (ent *Entorno) Guardar_Entorno(ento *Entorno) {
	ent.Lista_Entorno.Add(ento)
}

func (ent *Entorno) Retornar_Entorno(nom string) *Entorno {
	for i := 0; i < ent.Lista_Entorno.Len(); i++ {
		entact := ent.Lista_Entorno.GetValue(i).(*Entorno)
		if entact.Nombre_Entorno == nom {
			return entact
		}
	}
	return ent
}

func (ent *Entorno) GuardaLTodo(tod simbolos.SimboloTodo) {
	ent.ListaTodo.Add(tod)
}

func (ent *Entorno) EliminarLTodo() {
	ent.ListaTodo.RemoveAtIndex(0)
}

func (ent *Entorno) AumentarTodo() {
	ent.ContadorTodo += 1
}

func (ent *Entorno) ObtFunca(nom string) simbolos.SimboloTodo {
	for i := 0; i < ent.ListaTodo.Len(); i++ {
		act := ent.ListaTodo.GetValue(i).(simbolos.SimboloTodo)
		if act.Nombre == nom {
			ent.ListaTodo.RemoveAtIndex(i)
			return act
		}
	}
	return simbolos.SimboloTodo{}
}

func (ent *Entorno) ObtTamLTodo() int {
	return ent.ListaTodo.Len()
}

/*----------------------------------------------------------------------------------------------------------------------------------------------------------------*/
/*----------------------------------------------------------------------METODOS PARA LA TABLA DE VARIABLES--------------------------------------------------------*/
/* METODO PARA GUADAR VARIABLES */
func (ent *Entorno) Guardar_Variable(nombre string, simvar simbolos.Simbolo_Vars) {
	ent.Tabla_Variables[nombre] = simvar
	ent.Posicion = ent.Posicion + 1
	ent.Tamaño = ent.Tamaño + 1
}

/* METODO PARA GUADAR VARIABLES */
func (ent *Entorno) Guardar_Variable2(nombre string, simvar simbolos.Simbolo_Vars) {
	ent.Tabla_Variables2[nombre] = simvar
	ent.Posicion = ent.Posicion + 1
	ent.Tamaño = ent.Tamaño + 1
}

/* METODO PARA OBTENER VARIABLES */
func (ent *Entorno) Obtener_Variable(nom string) simbolos.Simbolo_Vars {
	var tmpEnv *Entorno = ent

	for {
		if variable, ok := tmpEnv.Tabla_Variables[nom]; ok {
			return variable
		}

		if tmpEnv.Entorno_Anterior == nil {
			break
		} else {
			tmpEnv = tmpEnv.Entorno_Anterior
		}
	}
	return simbolos.Simbolo_Vars{Mutable: false, NombreVariable: "", TipoVariable: simbolos.NULL, ValorVariable: 0, Linea: 0, Columna: 0}
}

/* MODIFICAR VARIABLE POR VALOR */
func (ent *Entorno) Modificar_Variable(nombre string, sim simbolos.Simbolo_Vars) {
	ent.Tabla_Variables[nombre] = sim
}

/* METODO PARA VERIFICAR SI EXISTE VARIABLES */
func (ent *Entorno) Existe_Variable(nom string) bool {
	var entemp *Entorno
	for entemp = ent; entemp != nil; entemp = entemp.Entorno_Anterior {
		if _, ok := entemp.Tabla_Variables[nom]; ok {
			return true
		}
	}
	return false
}

/* METODO PARA VERIFICAR SI EXISTE EN EL ENT ACTUAL VARIABLES */
func (ent *Entorno) ExisteAcual_Variable(nom string) bool {
	if _, ok := ent.Tabla_Variables[nom]; ok {
		return true
	}
	return false
}

/*----------------------------------------------------------------------------------------------------------------------------------------------------------------*/

/*----------------------------------------------------------------------------------------------------------------------------------------------------------------*/
/*----------------------------------------------------------------------METODOS PARA LA TABLA DE ARREVECT---------------------------------------------------------*/
/* METODO PARA GUADAR ARREGLOS Y VECTORES */
func (ent *Entorno) Guardar_ArregloVector(nom string, arrevect simbolos.Simbolo_ArreVect) {
	ent.Tabla_ArreVect[nom] = arrevect
	ent.Posicion = ent.Posicion + 1
	ent.Tamaño = ent.Tamaño + 1
}

/* METODO PARA OBTENER ARREGLOS Y VECTORES */
func (ent *Entorno) Obtener_ArreVect(nom string) simbolos.Simbolo_ArreVect {
	var tmpEnv *Entorno = ent

	for {
		if variable, ok := tmpEnv.Tabla_ArreVect[nom]; ok {
			return variable
		}

		if tmpEnv.Entorno_Anterior == nil {
			break
		} else {
			tmpEnv = tmpEnv.Entorno_Anterior
		}
	}
	return simbolos.Simbolo_ArreVect{Mutable: false, Nombre: "", TipoVect: simbolos.NULL, Linea: 0, Columna: 0}
}

/* MODIFICAR VECTOR POR VALOR */
func (ent *Entorno) Modificar_Vector(nombre string, sim simbolos.Simbolo_ArreVect) {
	ent.Tabla_ArreVect[nombre] = sim
}

/* METODO PARA VERIFICAR SI EXISTE ARREGLOS Y VECTORES */
func (ent *Entorno) Existe_ArreVect(nom string) bool {
	var entemp *Entorno
	for entemp = ent; entemp != nil; entemp = entemp.Entorno_Anterior {
		if _, ok := entemp.Tabla_ArreVect[nom]; ok {
			return true
		}
	}
	return false
}

/* METODO PARA VERIFICAR SI EXISTE EN EL ENT ACTUAL ARREGLOS Y VECTORES */
func (ent *Entorno) ExisteAcual_ArreVect(nom string) bool {
	if _, ok := ent.Tabla_ArreVect[nom]; ok {
		return true
	}
	return false
}

/*----------------------------------------------------------------------------------------------------------------------------------------------------------------*/

/*----------------------------------------------------------------------------------------------------------------------------------------------------------------*/
/*----------------------------------------------------------------------METODOS PARA LA TABLA DE FUNCIONES--------------------------------------------------------*/
/* GUARDAR FUNCIONES */
func (ent *Entorno) Guardar_Funciones(pub bool, id string, tfun simbolos.TipoExpresion, vva interface{}, lp *arraylist.List, li *arraylist.List, lin int, col int) {
	ent.Tabla_Funcines[id] = simbolos.Simbolo_Funciones{Publico: pub, Identificador: id, TipoFunca: tfun, ValVectArre: vva, L_Para: lp, L_Instr: li, Linea: lin, Columna: col}
	ent.Posicion = ent.Posicion + 1
	ent.Tamaño = ent.Tamaño + 1
}

/* OBTENER FUNCIONES */
func (ent *Entorno) Obtener_Funciones(nom string) simbolos.Simbolo_Funciones {
	var tmpEnv *Entorno = ent

	for {
		if variable, ok := tmpEnv.Tabla_Funcines[nom]; ok {
			return variable
		}

		if tmpEnv.Entorno_Anterior == nil {
			break
		} else {
			tmpEnv = tmpEnv.Entorno_Anterior
		}
	}
	return simbolos.Simbolo_Funciones{Publico: false, Identificador: "", TipoFunca: simbolos.NULL, Linea: 0, Columna: 0}
}

/* MODIFICAR FUNCION */
func (ent *Entorno) Modificar_Funcion(nombre string, sim simbolos.Simbolo_Funciones) {
	ent.Tabla_Funcines[nombre] = sim
}

/* EXISTE FUNCION EN LOS ENTORNOS */
func (ent *Entorno) Existe_Funciones(nom string) bool {
	var entemp *Entorno
	for entemp = ent; entemp != nil; entemp = entemp.Entorno_Anterior {
		if _, ok := entemp.Tabla_Funcines[nom]; ok {
			return true
		}
	}
	return false
}

func (ent *Entorno) Guardar_Retorno(nombre string, sim simbolos.Simbolo_Funciones) {
	ent.Tabla_Funcines[nombre] = sim
}

/* EXISTE FUNCION EN LA ACTUAL */
func (ent *Entorno) ExisteAcual_Funciones(nom string) bool {
	if _, ok := ent.Tabla_Funcines[nom]; ok {
		return true
	}
	return false
}

/*----------------------------------------------------------------------------------------------------------------------------------------------------------------*/

/*----------------------------------------------------------------------------------------------------------------------------------------------------------------*/
/*----------------------------------------------------------------------METODOS PARA LA TABLA DE STRUCTS----------------------------------------------------------*/
/* EXISTE FUNCION EN LOS ENTORNOS */
func (ent *Entorno) Existe_Structs(nom string) bool {
	var entemp *Entorno
	for entemp = ent; entemp != nil; entemp = entemp.Entorno_Anterior {
		if _, ok := entemp.Tabla_Struct[nom]; ok {
			return true
		}
	}
	return false
}

func (ent *Entorno) Guardar_Struct(nom string, stru simbolos.Simbolo_Struct) {
	ent.Tabla_Struct[nom] = stru
}

func (ent *Entorno) Obtener_Struct(nom string) simbolos.Simbolo_Struct {
	var tmpEnv *Entorno = ent

	for {
		if variable, ok := tmpEnv.Tabla_Struct[nom]; ok {
			return variable
		}

		if tmpEnv.Entorno_Anterior == nil {
			break
		} else {
			tmpEnv = tmpEnv.Entorno_Anterior
		}
	}
	return simbolos.Simbolo_Struct{Identificador: "", L_Atributos: arraylist.New(), Linea: 0, Columna: 0}
}

/*----------------------------------------------------------------------------------------------------------------------------------------------------------------*/

/*----------------------------------------------------------------------------------------------------------------------------------------------------------------*/
/*----------------------------------------------------------------------METODOS PARA LA TABLA DE MODULOS----------------------------------------------------------*/
func (ent *Entorno) Existe_Modulos(nom string) bool {
	var entemp *Entorno
	for entemp = ent; entemp != nil; entemp = entemp.Entorno_Anterior {
		if _, ok := entemp.Tabla_Modulos[nom]; ok {
			return true
		}
	}
	return false
}

func (ent *Entorno) Guardar_Modulos(nom string, stru simbolos.Simbolos_Modulos) {
	ent.Tabla_Modulos[nom] = stru
}

func (ent *Entorno) Obtener_Modulos(nom string) simbolos.Simbolos_Modulos {
	var tmpEnv *Entorno = ent

	for {
		if variable, ok := tmpEnv.Tabla_Modulos[nom]; ok {
			return variable
		}

		if tmpEnv.Entorno_Anterior == nil {
			break
		} else {
			tmpEnv = tmpEnv.Entorno_Anterior
		}
	}
	return simbolos.Simbolos_Modulos{Identificador: "", L_Declaraciones: arraylist.New(), Linea: 0, Columna: 0}
}

/*----------------------------------------------------------------------------------------------------------------------------------------------------------------*/
