package simbolos

import "github.com/colegno/arraylist"

type Valor struct {
	Tipo  TipoExpresion
	Valor interface{}
}

type SimboloTodo struct {
	Tipo   string
	Nombre string
	Pos    int
}

type Simbolo_Vars struct {
	Mutable        bool
	TipoVariable   TipoExpresion
	NombreVariable string
	ValorVariable  interface{}
	PosicionTabla  int
	Linea          int
	Columna        int
}

var ValArreglo = *arraylist.New()

type Simbolo_ArreVect struct {
	Mutable          bool
	EsArreVect       int
	TipoVect         TipoExpresion
	Nombre           string
	Dimensiones      int
	DimensionesLista *arraylist.List
	ValorArreVect    []interface{}
	ValorLista       *arraylist.List
	Lintdim          *arraylist.List
	Lintcap          *arraylist.List
	PosicionTabla    int
	Linea            int
	Columna          int
}

type Simbolo_Struct struct {
	Identificador string
	L_Atributos   *arraylist.List
	Tamaño        int
	Linea         int
	Columna       int
}

type Simbolos_Modulos struct {
	Publico         bool
	Identificador   string
	L_Declaraciones *arraylist.List
	Linea           int
	Columna         int
}

type Simbolo_Funciones struct {
	Publico       bool
	Identificador string
	TipoFunca     TipoExpresion
	ValVectArre   interface{}
	L_Para        *arraylist.List
	L_Instr       *arraylist.List
	Tam           int
	Linea         int
	Columna       int
}
