package generando

import (
	reportes "OLC2-PROYECTO2/COMPILADOR/REPORTES"
	bloque "OLC2-PROYECTO2/OPT_COMPILADOR/BLOQUE"

	"github.com/colegno/arraylist"
)

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
func (g *GenerandoOpti) AgegarCodigoOpt(entra bloque.ObjetoLista) {
	g.CodigoOpt.Add(entra)
}

func (g *GenerandoOpti) EliminarTempo(nom string, val string, lin string) {
	for i := 0; i < g.CodigoOpt.Len(); i++ {
		act := g.CodigoOpt.GetValue(i).(bloque.ObjetoLista)
		if act.Declaracion == nom {
			g.CodigoOpt.RemoveAtIndex(i)
			reportes.ReporteOpti("Bloques", "Regla 3", nom+"="+val, "--", lin)
		}
	}
}
