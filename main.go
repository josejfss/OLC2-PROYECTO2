package main

import (
	//reportando "OLC2-PROYECTO2/ANALIZADOR/REPORTANDO"
	funciones "OLC2-PROYECTO2/COMPILADOR/AST/INSTRUCCIONES/FUNCIONES"
	structs "OLC2-PROYECTO2/COMPILADOR/AST/INSTRUCCIONES/STRUCTS"
	entorno "OLC2-PROYECTO2/COMPILADOR/ENTORNO"
	generador "OLC2-PROYECTO2/COMPILADOR/GENERADOR"
	interfaces "OLC2-PROYECTO2/COMPILADOR/INTERFACES"
	reportes "OLC2-PROYECTO2/COMPILADOR/REPORTES"
	"OLC2-PROYECTO2/COMPILADOR/parser"
	opt "OLC2-PROYECTO2/OPTIMIZADOR/OPT"
	"fmt"
	"log"
	"os"
	"os/exec"
	"reflect"

	// "io/ioutil"

	// "fyne.io/fyne/dialog"
	// "fyne.io/fyne/storage"
	// "fyne.io/fyne/v2"
	// "fyne.io/fyne/v2/app"
	"io/ioutil"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/colegno/arraylist"
	// "fyne.io/fyne/v2/container"
	// "fyne.io/fyne/v2/widget"
)

type TreeShapeListener struct {
	*parser.Basedb_rustparserListener
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

var salida string

func (*TreeShapeListener) ExitStart(ctx *parser.StartContext) {
	result := ctx.GetLista()

	var gen *generador.Generador_C3D = generador.Ngenerador()
	salida = ""

	var entorno_global *entorno.Entorno = entorno.Nuevo_Entorno("global", nil)
	var EntGlobal *entorno.Entorno = entorno.Nuevo_Entorno("gobalc3d", nil)
	//var Ent_Global *entorno.Entorno = entorno.Nuevo_Entorno("gobalc3d", nil)
	for i := 0; i < result.Len(); i++ {
		instr := result.GetValue(i).(interfaces.Instruccion)
		instr.Ejecutar_Instruccion(entorno_global, EntGlobal)
	}

	fmt.Println("fin interpretacion")
	// for _, s := range result.ToArray() {
	// 	s.(interfaces.Instruccion).Compilar_Instruccion(entorno_global, gen)
	// }

	var fun_main funciones.Funciones
	lista_declas := arraylist.New()
	gen.Agregar_MetodoPrint()
	for i := 0; i < result.Len(); i++ {
		instr := result.GetValue(i).(interfaces.Instruccion)
		if instr != nil {
			if reflect.TypeOf(instr) == reflect.TypeOf(funciones.Funciones{}) {
				funca := instr.(funciones.Funciones)
				if funca.NombreFunca == "main" {
					//fmt.Println("main")
					fun_main = funca
				} else {
					instr.Compilar_Instruccion(EntGlobal, gen)
				}
				//result.RemoveAtIndex(i)
			} else if reflect.TypeOf(instr) != reflect.TypeOf(structs.StructRust{}) {
				lista_declas.Add(instr)
			}
		}
	}

	gen.Agregar_Comentario("------------------ENTRANDO FUNCION PRINCIPAL-------------------")
	gen.Agregar_Logica("int main() {\nSP=0;\t\t//INICIANDO PUNTERO STACK\nHP=0;\t\t//INICIANDO PUNTERO HEAP\n\n")
	for i := 0; i < lista_declas.Len(); i++ {
		instr := lista_declas.GetValue(i).(interfaces.Instruccion)
		instr.Compilar_Instruccion(EntGlobal, gen)
	}
	fun_main.Compilar_Instruccion(EntGlobal, gen)
	gen.Agregar_Logica("return 0;\n}")
	gen.Agregar_Comentario("------------------SALIENDO FUNCION PRINCIPAL-------------------")

	//Escribir salida
	f, err := os.Create("/home/iovana/go/src/OLC2-PROYECTO2/REP/salida.cpp")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	salida += "/*------ENACABEZADO------*/\n"
	salida += "#include <stdio.h>\n"
	salida += "/*------DECLARACION HEAP------*/\n"
	salida += "double HEAP[100000];\n"
	salida += "/*------DECLARACION STACK------*/\n"
	salida += "double STACK[100000];\n"
	salida += "/*------DECLARACION PUNTERO STACK------*/\n"
	salida += "double SP;\n"
	salida += "/*------DECLARACION PUNTERO HEAP------*/\n"
	salida += "double HP;\n"
	salida += "/*------DECLARACION TEMPORALES------*/\n"
	salida += "double "

	salida += fmt.Sprintf("%v", gen.Obtener_Ltemporales().GetValue(0))
	gen.Obtener_Ltemporales().RemoveAtIndex(0)

	for _, s := range gen.Obtener_Ltemporales().ToArray() {
		salida += ", "
		salida += fmt.Sprintf("%v", s)
	}

	salida += ";\n\n"
	//salida += "int main(){\n"

	for _, s := range gen.Obtener_codigo().ToArray() {
		salida += fmt.Sprintf("%v", s)
		salida += "\n"
	}

	//salida += "\nreturn 0;\n}\n"

	_, err2 := f.WriteString(salida)

	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println("fin compilacion")
}

//FUNCION PARA LEER LOS TOKENS Y QUE SI JALE LA GRAMATICA GG
func analizar(e string) {
	// is, _ := antlr.NewFileStream("pruebas.txt")
	is := antlr.NewInputStream(e)
	// Create the Lexer
	lexer := parser.Newdb_rustlexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.Newdb_rustparser(stream)

	p.BuildParseTrees = true
	tree := p.Start()

	antlr.ParseTreeWalkerDefault.Walk(NewTreeShapeListener(), tree)
}

var texto string = ""

//FUNCION MAIN
func main() {
	//PARA VERIFICAR QUE LA APLICACION YA INICIO
	fmt.Println("hola mundo")

	//SE CREA UNA APLICACION PARA INICIAR LA GUI
	applicacion := app.New()
	//SE CREA LA VENTANA MASTER DONDE IRA TODOS LOS COMPONENTES DEL PROYECTO
	ventana := applicacion.NewWindow("********** DB-RUST  OLC2-PROYECTO2 **********")
	//SE HACE ESTA VENTANA MASTER QUE CUANDO SE CIERRE ESTA VENTANA SE TERMINA LA EJECUCION
	ventana.SetMaster()
	//SE CAMBIA EL TAMAÑO DE LA VENTANA
	ventana.Resize(fyne.NewSize(1790, 650))
	//SE COLOCA QUE LA VENTANA SE ABRA EN EL CENTRO DE LA PANTALLA
	ventana.CenterOnScreen()

	//------------------------------------------------- CREACION DE COMPONENTES DEL PROYECTO -----------------------------------------------------
	//SE CREA UN LABEL PARA MOSTRAR EL TITULO
	lbl := widget.NewLabel("DB-RUST")
	//SE LE COLOCA UNA NUEVA POSICION
	lbl.Move(fyne.NewPos(870, 5))

	//SE CREA EL TEXTAREA PARA LA CONSOLA
	consola1 := widget.NewMultiLineEntry()
	consola1.SetPlaceHolder("Optimizacion 3 direcciones...")
	consola1.Resize(fyne.NewSize(550, 550))
	consola1.Move(fyne.NewPos(1200, 45))

	//SE CREA EL TEXTAREA PARA LA CONSOLA
	consola := widget.NewMultiLineEntry()
	consola.SetPlaceHolder("Codigo 3 direcciones...")
	consola.Resize(fyne.NewSize(550, 550))
	consola.Move(fyne.NewPos(620, 45))

	//SE CREA EL TEXTAREA PARA LA ENTRADA
	entrada := widget.NewMultiLineEntry()
	//SE LE COLOCA UN TEXTO POR DEFECTO
	entrada.SetPlaceHolder("Ingrese texto BD-RUST...")
	//SE CAMBIA LAS DIMENSIONES
	entrada.Resize(fyne.NewSize(550, 550))
	//NUEVA POSICION DEL TEXTAREA
	entrada.Move(fyne.NewPos(40, 45))
	//-------------------------------------------------------------------------------------- -----------------------------------------------------

	//------------------------------------------------- CREACION DEL MENU PRINCIPAL DEL PROYECTO -------------------------------------------------
	//SE CREA EL MENU PRINCIPAL DONDE ESTARAN LAS FUNCIONES DE
	//ABRIR ITEMS
	//var texto string
	menu_abrir := fyne.NewMenuItem("Abrir DR-Rust", func() {
		file_Dialog := dialog.NewFileOpen(

			func(r fyne.URIReadCloser, _ error) {
				// read files
				data, _ := ioutil.ReadAll(r)

				result := fyne.NewStaticResource("name", data)

				entry := widget.NewMultiLineEntry()

				entry.SetText(string(result.StaticContent))
				//fmt.Println(string(result.StaticContent))
				entrada.Text = string(result.StaticContent)
				entrada.Refresh()
				texto = string(result.StaticContent)

			}, ventana)
		file_Dialog.SetFilter(
			storage.NewExtensionFileFilter([]string{".rs"}))
		file_Dialog.Show()
	})

	//MENU ABRIR
	menu1 := fyne.NewMenu("Archivo", menu_abrir)
	//EJECUTAR ITEMS
	menu_ejecutar := fyne.NewMenuItem("Compilar", func() {
		analizar(texto)
		consola.Text = salida
		consola.Refresh()
	})
	//-------------------------------------------------------------------------------------- -----------------------------------------------------
	//MENU EJECUTAR
	menu2 := fyne.NewMenu("Compilar", menu_ejecutar)

	menu_optimizar := fyne.NewMenuItem("Optimizar", func() {
		texts := consola.Text
		opt.Optimizar(texts)
		consola1.Text = opt.OptimizadoC3D
		consola1.Refresh()
	})
	//-------------------------------------------------------------------------------------- -----------------------------------------------------
	//MENU EJECUTAR
	menu5 := fyne.NewMenu("Optimizar", menu_optimizar)

	//REPORTES ITEMS
	menu_reporte := fyne.NewMenuItem("Tabla de Simbolos", func() {
		//LLAMADA DEL METODO PARA LA CREACION DEL REPORTE
		reportes.CreandoReporteSimbolos()
		//ABRIR EL REPORTE CON COMANDOS
		prc := exec.Command("xdg-open", "/home/iovana/go/src/OLC2-PROYECTO2/REP/reporte_simbolos.pdf")
		err := prc.Run()
		if err != nil {
			fmt.Printf("error:%v \n", err)
		}
	})
	menu_reporte1 := fyne.NewMenuItem("Tabla de Errores", func() {
		reportes.Reporte_Errores()
		//ABRIR EL REPORTE CON COMANDOS
		prc := exec.Command("xdg-open", "/home/iovana/go/src/OLC2-PROYECTO2/REP/reporte_errores.pdf")
		err := prc.Run()
		if err != nil {
			fmt.Printf("error:%v \n", err)
		}
	})
	menu_reporte2 := fyne.NewMenuItem("Base de Datos Existente", func() {
		reportes.CreandoReporteBasesDatos()
		//ABRIR EL REPORTE CON COMANDOS
		prc := exec.Command("xdg-open", "/home/iovana/go/src/OLC2-PROYECTO2/REP/bases_datos.pdf")
		err := prc.Run()
		if err != nil {
			fmt.Printf("error:%v \n", err)
		}
	})
	menu_reporte3 := fyne.NewMenuItem("Optimizacion", func() {
		reportes.CreandoReporteOptimizacion()
		//ABRIR EL REPORTE CON COMANDOS
		prc := exec.Command("xdg-open", "/home/iovana/go/src/OLC2-PROYECTO2/REP/reporte_optimizacion.pdf")
		err := prc.Run()
		if err != nil {
			fmt.Printf("error:%v \n", err)
		}
	})
	//MENU REPORTE
	menu3 := fyne.NewMenu("Reportes", menu_reporte, menu_reporte1, menu_reporte2, menu_reporte3)
	//INFO ITEMS
	menu_acerca := fyne.NewMenuItem("Acerca De", func() {
		info_ventana := applicacion.NewWindow("********** ACERDA DE ESTUDIANTE **********")
		//INFORMACION A MOSTRAR
		informacion := "***********************************************\n" //46
		informacion += "************* INFORMACION DEL CURSO **********\n"
		informacion += "***********************************************\n"
		informacion += "* UNIVERSIDAD DE SAN CARLOS DE GUATEMALA          *\n"
		informacion += "* FACULTAD DE INGENIERIA                                               *\n"
		informacion += "* ESCUELA DE CIENCIAS Y SISTEMAS                                *\n"
		informacion += "* AREA DE CIENCIAS DE LA COMPUTACION                    *\n"
		informacion += "* ORGANIZACION DE LENGUAJES Y COMPILADORES 2 *\n"
		informacion += "* SECCION B                                                                          *\n"
		informacion += "***********************************************\n"
		informacion += "+++++++++++++++++++++++++++++++++++++++++++++\n"
		informacion += "+++++++++ INFORMACION DEL ESTUDIANTE ++++++++\n"
		informacion += "+++++++++++++++++++++++++++++++++++++++++++++\n"
		informacion += "+ CLAUDIA IOVANA MIRANDA ALVAREZ 			                 +\n"
		informacion += "+ 201700387 								                                         +\n"
		informacion += "+++++++++++++++++++++++++++++++++++++++++++++\n"
		//SE AGREGA LA INFORMACION A UN LABEL EN LA VENTANA
		info_ventana.SetContent(widget.NewLabel(informacion))
		//SE REDIMENSIONA LA VENTANA
		info_ventana.Resize(fyne.NewSize(600, 400))
		//LA VENTANA SE MUESTRA EN EL CENTRO DE LA PANTALLA
		info_ventana.CenterOnScreen()
		//MOSTRAR LA VENTANA
		info_ventana.Show()
	})
	//MENU ACERCA DE
	menu4 := fyne.NewMenu("Informacion", menu_acerca)
	//SE CREA EL MENU COMPLETO AGREGANDO TODOS LOS SUBMENUS
	todo_menu := fyne.NewMainMenu(menu1, menu2, menu5, menu3, menu4)
	//SE AGREGA EL MENU A LA VENTANA
	ventana.SetMainMenu(todo_menu)
	//--------------------------------------------------------------------------------------------------------------------------------------------

	//SE AGREGAN TODOS LOS COMPONENTES
	contendo_principal := container.NewWithoutLayout(lbl, entrada, consola, consola1)

	//SE AGREGA EL COMPONENTE A LA VENTANA
	ventana.SetContent(contendo_principal)
	//SE CORRE LA VENTANA
	ventana.ShowAndRun()
}
