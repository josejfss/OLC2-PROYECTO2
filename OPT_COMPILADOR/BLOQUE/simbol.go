package bloque

type ObjetoBloque struct {
	Tipo       TipoObjeto
	Asignacion string
	OpeIz      string
	Ope        string
	OpeDe      string
	Valor      string
	Linea      int
}

type ObjetoLista struct {
	Declaracion string
	Valores     string
}

func Nobjetolista(d string, v string) ObjetoLista {
	objlist := ObjetoLista{Declaracion: d, Valores: v}
	return objlist
}
