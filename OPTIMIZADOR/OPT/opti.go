package opt

import (
	interfaz "OLC2-PROYECTO2/OPTIMIZADOR/INTERFAZ"
	objeto "OLC2-PROYECTO2/OPTIMIZADOR/OBJETO"
	reglas "OLC2-PROYECTO2/OPTIMIZADOR/REGLAS"
	"OLC2-PROYECTO2/OPTIMIZADOR/parser"
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

	var bloque_global *objeto.Bloque = objeto.Nuevo_Bloque()

	for i := 0; i < result.Len(); i++ {
		instr := result.GetValue(i).(interfaz.Instruccion)
		instr.Optimizar_Instruccion(bloque_global)
	}

	fmt.Println("fin optimizacion")

	reglas.Regla1(bloque_global)
	fmt.Println("fin regla1")
	reglas.Regla2(bloque_global)
	fmt.Println("fin regla2")
	// reglas.Regla4(bloque_global)
	// fmt.Println("fin regla4")
	// reglas.Regla3(bloque_global)
	// fmt.Println("fin regla3")
	Escribir(bloque_global)

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

func Escribir(block *objeto.Bloque) {
	contando := 0
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

	for j := 0; j < block.ListaBloques.Len(); j++ {
		bloque := block.ListaBloques.GetValue(j).(*objeto.Bloque)
		for i := 0; i < len(bloque.Tabla_Bloques); i++ {
			act := block.ListaTodo.GetValue(contando).(string)
			simbs := bloque.Tabla_Bloques[act]
			if simbs.Tipo == 1 {
				concatenando = concatenando + simbs.Temporal + "=" + simbs.Valor + ";\n"
				contando = contando + 1
			} else if simbs.Tipo == 2 {
				concatenando = concatenando + "if(" + simbs.Opiz + simbs.Ope + simbs.Opde + ") goto " + simbs.Temporal + ";\n"
				contando = contando + 1
			} else if simbs.Tipo == 3 {
				concatenando = concatenando + "printf(" + simbs.Temporal + ", " + simbs.Opde + simbs.Opiz + ");\n"
				contando = contando + 1
			} else {
				//fmt.Println(simbs.Valor)
				concatenando = concatenando + simbs.Valor + "\n"
				contando = contando + 1
			}
		}
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
