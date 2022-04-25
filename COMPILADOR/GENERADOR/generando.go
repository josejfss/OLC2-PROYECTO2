package generador

import (
	"fmt"

	"github.com/colegno/arraylist"
)

//CLASE GENERADOR DE C3D
type Generador_C3D struct {
	Temporal       int
	Label          int
	Codigo         *arraylist.List
	Lista_temporal *arraylist.List
	Lista_etiqueta *arraylist.List
}

//CONSTRUCTOR DE LA CLASE GENERADOR DE C3D
func Ngenerador() *Generador_C3D {
	gc3d := Generador_C3D{Temporal: 0, Label: 0, Codigo: arraylist.New(), Lista_temporal: arraylist.New(), Lista_etiqueta: arraylist.New()}
	return &gc3d
}

//METODO PARA OBTENER EL CODIGO YA TRADUCIDO
func (g *Generador_C3D) Obtener_codigo() *arraylist.List {
	//SE DEVUELVE LA LISTA DE CODIGO
	return g.Codigo
}

func (g *Generador_C3D) Obtener_TamCodigo() int {
	return g.Codigo.Len()
}

//METODO PARA OBTENER LA LISTA DE TEMPORALES
func (g *Generador_C3D) Obtener_Ltemporales() *arraylist.List {
	//SE DEVUELVE LA LISTA DE TEMPORALES
	return g.Lista_temporal
}

//METODO PARA CREAR TEMPORALES Y AGREGARLOS A LA LISTA
func (g *Generador_C3D) Crear_temporal() string {
	//SE CREA EL TEMPORAL
	tempo := "T" + fmt.Sprintf("%v", g.Temporal)
	//SE AUMENTA EN UNO AL CONTADOR
	g.Temporal = g.Temporal + 1
	//SE VERIFICA SI EL TEMPORAL YA EXISTE
	//BANDERA PARA VERIFICAR QUE EXISTA EL TEMPORAL EN LA LISTA
	bandera := false
	//SE RECORRE LA LISTA DE TEMPORALES
	for i := 0; i < g.Lista_temporal.Len(); i++ {
		//SE OBTIENE EL TEMPORAL ACTUAL DE LA LISTA
		acttemp := g.Lista_temporal.GetValue(i).(string)
		//SE VERIFICA SI EL TEMPORAL YA EXISTE EN LA LISTA
		if tempo == acttemp {
			//DE EXISTIR LA BANDERA SE VUELVE VERDADERA
			bandera = true
		}
	}
	//SI LA BANDERA ES FALSA ES PORQUE NO EXISTE POR ESO SE GUARDA
	if !bandera {
		//SE GUARDA EN LA LISTA DE TEMPORALES
		g.Lista_temporal.Add(tempo)
	}
	//SE DEVUELVE EL TEMPORAL
	return tempo
}

func (g *Generador_C3D) LiberarTodosTemporales() {
	for i := 0; i < g.Lista_temporal.Len(); i++ {
		if g.Temporal != 0 {
			g.Temporal = g.Temporal - 1
		}
	}
}

//METODO PARA CREAR LABELS
func (g *Generador_C3D) Crear_label() string {
	//SE CREA LABEL
	lab := "L" + fmt.Sprintf("%v", g.Label)
	//SE AUMENTA EN UNO AL CONTADOR
	g.Label = g.Label + 1
	//SE GUARDA EN LA LISTA DE ETIQUETAS
	g.Lista_etiqueta.Add(lab)
	//SE DEVUELVE EL LABEL
	return lab
}

//METODO PARA ELIMINAR LABELS
func (g *Generador_C3D) Eliminar_label(borrar string) {
	for i := 0; i < g.Lista_etiqueta.Len(); i++ {
		eti_act := g.Lista_etiqueta.GetValue(i)
		if eti_act == borrar {
			g.Lista_etiqueta.RemoveAtIndex(i)
		}
	}
}

//METODO PARA AGREGAR EL LABEL A LA LISTA CODIGO
func (g *Generador_C3D) Agregar_label(lab string) {
	//SE CREA LABEL A GUARDAR
	labs := lab + ":"
	//SE AÑADE A LA LISTA DEL CODIGO
	g.Codigo.Add(labs)
}

//METODO PARA AGREGAR EXPRESIONES ARITMETICAS A LA LISTA CODIGO
func (g *Generador_C3D) Agregar_Expression(resultado string, opizq string, op string, opder string) {
	expres := resultado + " = " + opizq + " " + op + " " + opder + ";"
	if op == "+" {
		expres += "\t\t//OPERACION SUMA"
	} else if op == "-" {
		expres += "\t\t//OPERACION RESTA"
	} else if op == "*" {
		expres += "\t\t//OPERACION MULTIPLICACION"
	} else if op == "/" {
		expres += "\t\t//OPERACION DIVISION"
	} else if op == "%" {
		expres += "\t\t//OPERACION MODULO"
	}
	g.Codigo.Add(expres)
}

//METODO PARA AGREGAR EXPRESIONES RELACIONALES A LA LISTA CODIGO
func (g *Generador_C3D) Agregar_Relacional(opizq string, op string, opder string, ev string, ef string) {
	rela := "if (" + opizq + " " + op + " " + opder + ") { goto " + ev + "; }\ngoto " + ef + ";\n"
	g.Codigo.Add(rela)
}

//METODO PARA AGREGAR EXPRESIONES LOGICAS A LA LISTA CODIGO
func (g *Generador_C3D) Agregar_Logica(entra string) {
	g.Codigo.Add(entra)
}

//METODO PARA AGREGAR COMENTARIO A LA LISTA CODIGO
func (g *Generador_C3D) Agregar_Comentario(coment string) {
	com := "/*" + coment + "*/"
	g.Codigo.Add(com)
}

//METODO PARA AGREGAR COSAS AL STACK
func (g *Generador_C3D) Agregar_Stack(posi string, val string) {
	pila := "STACK[(int)" + posi + "] = " + val + ";"
	g.Codigo.Add(pila)
}

func (g *Generador_C3D) AgregarErrorMate(lin string, col string) {
	errors := "printf(\"%c\", 77);	//LETRA M\n"
	errors += "printf(\"%c\", 65);	//LETRA A\n"
	errors += "printf(\"%c\", 84);	//LETRA T\n"
	errors += "printf(\"%c\", 72);	//LETRA H\n"
	errors += "printf(\"%c\", 32);	//ESPACIO\n"
	errors += "printf(\"%c\", 69);	//LETRA E\n"
	errors += "printf(\"%c\", 82);	//LETRA R\n"
	errors += "printf(\"%c\", 82);	//LETRA R\n"
	errors += "printf(\"%c\", 79);	//LETRA O\n"
	errors += "printf(\"%c\", 82);	//LETRA R\n"
	errors += "printf(\"%c\", 46);	//PUNTO\n"
	errors += "printf(\"%c\", 32);	//ESPACIO\n"
	errors += "printf(\"%c\", 76);	//LETRA L\n"
	errors += "printf(\"%c\", 73);	//LETRA I\n"
	errors += "printf(\"%c\", 78);	//LETRA N\n"
	errors += "printf(\"%c\", 46);	//PUNTO\n"
	errors += "printf(\"%c\", 32);	//ESPACIO\n"
	errors += "printf(\"%o\", int(" + lin + "));	//NUM\n"
	errors += "printf(\"%c\", 32);	//ESPACIO\n"
	errors += "printf(\"%c\", 67);	//LETRA C\n"
	errors += "printf(\"%c\", 79);	//LETRA O\n"
	errors += "printf(\"%c\", 76);	//LETRA L\n"
	errors += "printf(\"%c\", 46);	//PUNTO\n"
	errors += "printf(\"%c\", 32);	//ESPACIO\n"
	errors += "printf(\"%d\", int(" + col + "));	//NUM\n"
	errors += "printf(\"%c\", 10);	//SALTO LINEA\n"
	g.Codigo.Add(errors)
}

func (g *Generador_C3D) AgregarError(mensaje string, lin string, col string) {
	enlinea := " En la Linea: "
	encolum := " En la columna: "
	errores := ""
	for _, txt := range mensaje {
		f := int(txt)
		errores += "printf(\"%c\", " + fmt.Sprintf("%v", f) + "); //LETRA-> " + string(txt) + "\n"
	}
	errores += "printf(\"%c\", 46);	//PUNTO\n"
	for _, txt := range enlinea {
		f := int(txt)
		errores += "printf(\"%c\", " + fmt.Sprintf("%v", f) + "); //LETRA-> " + string(txt) + "\n"
	}
	errores += "printf(\"%o\", int(" + lin + "));	//NUM\n"
	errores += "printf(\"%c\", 46);	//PUNTO\n"
	errores += "printf(\"%c\", 32);	//ESPACIO\n"
	for _, txt := range encolum {
		f := int(txt)
		errores += "printf(\"%c\", " + fmt.Sprintf("%v", f) + "); //LETRA-> " + string(txt) + "\n"
	}
	errores += "printf(\"%d\", int(" + col + "));	//NUM\n"
	errores += "printf(\"%c\", 46);	//PUNTO\n"
	errores += "printf(\"%c\", 10);	//SALTO LINEA\n"
	g.Codigo.Add(errores)
}

func (g *Generador_C3D) Agregar_MetodoPrint() {
	g.Agregar_Comentario("INICIO DECLARACION FUNCION NATIVA PRINT")
	temp_puntero := g.Crear_temporal()
	temp_heap := g.Crear_temporal()
	temp_conca := g.Crear_temporal()
	etiqueta_entrada := g.Crear_label()
	g.Eliminar_label(etiqueta_entrada)
	etiqueta_salida := g.Crear_label()
	g.Eliminar_label(etiqueta_salida)

	prints := "void print() {\n"
	prints += temp_puntero + " = SP;\n"
	prints += temp_heap + " = STACK[(int)" + temp_puntero + "];\n"
	prints += etiqueta_entrada + ":\n"
	prints += temp_conca + " = HEAP[(int)" + temp_heap + "];\n"
	prints += "if ( " + temp_conca + " == -1) { goto " + etiqueta_salida + "; }\n"
	prints += "printf(\"%c\", (char)" + temp_conca + ");\n"
	prints += temp_heap + " = " + temp_heap + " + 1;\n"
	prints += "goto " + etiqueta_entrada + ";\n"
	prints += etiqueta_salida + ":\n"
	prints += "return;\n}\n"
	g.Agregar_Logica(prints)
	g.Agregar_Comentario("FIN DECLARACION FUNCION NATIVA PRINT")
	g.LiberarTodosTemporales()
}
