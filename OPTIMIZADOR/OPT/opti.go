package opt

import (
	interfaz "OLC2-PROYECTO2/OPTIMIZADOR/INTERFAZ"
	objeto "OLC2-PROYECTO2/OPTIMIZADOR/OBJETO"
	"OLC2-PROYECTO2/OPTIMIZADOR/parser"
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type TreeShapeListener struct {
	*parser.Baseoptimizador_parserListener
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

//var salida1 string

func (*TreeShapeListener) ExitStart(ctx *parser.StartContext) {
	fmt.Println("INICIO optimizacion")
	result := ctx.GetLista()

	var bloque_global *objeto.Bloque = objeto.Nuevo_Bloque()

	for i := 0; i < result.Len(); i++ {
		instr := result.GetValue(i).(interfaz.Instruccion)
		instr.Optimizar_Instruccion(bloque_global)
	}

	fmt.Println("fin optimizacion")
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
