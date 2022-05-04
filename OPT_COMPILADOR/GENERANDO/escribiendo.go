package generando

import "github.com/colegno/arraylist"

type GenerandoOpti struct {
	CodigoOpt *arraylist.List
}

//CONSTRUCTOR DE LA CLASE GENERADOR DE OPT
func Ngenerador() *GenerandoOpti {
	gc3d := GenerandoOpti{CodigoOpt: arraylist.New()}
	return &gc3d
}

//METODO PARA OBTENER EL CODIGO YA OPTIMIZADO
func (g *GenerandoOpti) Obtener_codigo() *arraylist.List {
	//SE DEVUELVE LA LISTA DE CODIGO
	return g.CodigoOpt
}

//METODO PARA AGREGAR EXPRESIONES LOGICAS A LA LISTA CODIGO
func (g *GenerandoOpti) AgegarCodigoOpt(entra string) {
	g.CodigoOpt.Add(entra)
}
