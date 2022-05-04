package enviar

import (
	objeto "OLC2-PROYECTO2/OPT_COMPILADOR/BLOQUE"
	generando "OLC2-PROYECTO2/OPT_COMPILADOR/GENERANDO"
	inter "OLC2-PROYECTO2/OPT_COMPILADOR/INTER"
	"OLC2-PROYECTO2/OPT_COMPILADOR/parser"
	"fmt"
	"log"
	"os"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/colegno/arraylist"
)

type TreeShapeListener struct {
	*parser.Baseoptimizador_parserListener
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

var OptimizadoC3D string = ""

var ListaBorrarTemporales *arraylist.List = arraylist.New()

func (*TreeShapeListener) ExitStart(ctx *parser.StartContext) {
	fmt.Println("INICIO optimizacion")
	result := ctx.GetLista()

	var bloque_global *objeto.BloquesOpt = objeto.Nuevo_Bloque("bloque_global")
	var generador_global *generando.GenerandoOpti = generando.Ngenerador()

	for i := 0; i < result.Len(); i++ {
		instr := result.GetValue(i).(inter.Instruccion)
		instr.Opti_Instruccion(bloque_global, generador_global)
	}

	fmt.Println("fin optimizacion")

	// reglas.Regla1(bloque_global)
	// fmt.Println("fin regla1")
	// reglas.Regla2(bloque_global)
	// fmt.Println("fin regla2")
	// // reglas.Regla4(bloque_global)
	// // fmt.Println("fin regla4")
	// // reglas.Regla3(bloque_global)
	// // fmt.Println("fin regla3")
	Escribir(bloque_global, generador_global)

}

//FUNCION PARA LEER LOS TOKENS Y QUE SI JALE LA GRAMATICA GG
func Optimizar(e string) {
	// is, _ := antlr.NewFileStream("pruebas.txt")
	is := antlr.NewInputStream(e)
	// Create the Lexer
	lexer := parser.Newoptimizador_lexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.Newoptimizador_parser(stream)

	p.BuildParseTrees = true
	tree := p.Start()

	antlr.ParseTreeWalkerDefault.Walk(NewTreeShapeListener(), tree)
}

func Escribir(block *objeto.BloquesOpt, genopt *generando.GenerandoOpti) {
	// contando := 0
	concatenando := block.Encabezado + "\n"
	concatenando = concatenando + "double "
	for i := 0; i < block.ListaTemporales.Len(); i++ {
		act := block.ListaTemporales.GetValue(i).(string)
		concatenando = concatenando + act
		if i == block.ListaTemporales.Len()-1 {
			concatenando = concatenando + ";\n"
		} else {
			concatenando = concatenando + ","
		}
	}

	for j := 0; j < genopt.CodigoOpt.Len(); j++ {
		bloque := genopt.CodigoOpt.GetValue(j).(objeto.ObjetoLista)
		concatenando = concatenando + bloque.Valores + "\n"
	}
	fmt.Println(concatenando)
	OptimizadoC3D = concatenando
	//Escribir salida
	f, err := os.Create("/home/iovana/go/src/OLC2-PROYECTO2/REP/optimizacion.cpp")
	if err != nil {
		log.Fatal(err)
	}

	_, err2 := f.WriteString(concatenando)

	if err2 != nil {
		log.Fatal(err2)
	}
	defer f.Close()
}
