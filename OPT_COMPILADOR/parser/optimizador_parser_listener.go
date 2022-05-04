// Code generated from optimizador_parser.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // optimizador_parser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// optimizador_parserListener is a complete listener for a parse tree produced by optimizador_parser.
type optimizador_parserListener interface {
	antlr.ParseTreeListener

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterInstrucss is called when entering the instrucss production.
	EnterInstrucss(c *InstrucssContext)

	// EnterInstrs is called when entering the instrs production.
	EnterInstrs(c *InstrsContext)

	// EnterEncabezado is called when entering the encabezado production.
	EnterEncabezado(c *EncabezadoContext)

	// EnterL_temporales is called when entering the l_temporales production.
	EnterL_temporales(c *L_temporalesContext)

	// EnterL_funcas is called when entering the l_funcas production.
	EnterL_funcas(c *L_funcasContext)

	// EnterFuncas is called when entering the funcas production.
	EnterFuncas(c *FuncasContext)

	// EnterSent_if is called when entering the sent_if production.
	EnterSent_if(c *Sent_ifContext)

	// EnterImprimir is called when entering the imprimir production.
	EnterImprimir(c *ImprimirContext)

	// EnterCasteo is called when entering the casteo production.
	EnterCasteo(c *CasteoContext)

	// EnterDeclaracion is called when entering the declaracion production.
	EnterDeclaracion(c *DeclaracionContext)

	// EnterEtiquetas is called when entering the etiquetas production.
	EnterEtiquetas(c *EtiquetasContext)

	// EnterLlamada is called when entering the llamada production.
	EnterLlamada(c *LlamadaContext)

	// EnterL_bloque is called when entering the l_bloque production.
	EnterL_bloque(c *L_bloqueContext)

	// EnterBloque is called when entering the bloque production.
	EnterBloque(c *BloqueContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterAsignaciones is called when entering the asignaciones production.
	EnterAsignaciones(c *AsignacionesContext)

	// EnterExpre_relacional is called when entering the expre_relacional production.
	EnterExpre_relacional(c *Expre_relacionalContext)

	// EnterExpre_aritmetica is called when entering the expre_aritmetica production.
	EnterExpre_aritmetica(c *Expre_aritmeticaContext)

	// EnterDatos is called when entering the datos production.
	EnterDatos(c *DatosContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitInstrucss is called when exiting the instrucss production.
	ExitInstrucss(c *InstrucssContext)

	// ExitInstrs is called when exiting the instrs production.
	ExitInstrs(c *InstrsContext)

	// ExitEncabezado is called when exiting the encabezado production.
	ExitEncabezado(c *EncabezadoContext)

	// ExitL_temporales is called when exiting the l_temporales production.
	ExitL_temporales(c *L_temporalesContext)

	// ExitL_funcas is called when exiting the l_funcas production.
	ExitL_funcas(c *L_funcasContext)

	// ExitFuncas is called when exiting the funcas production.
	ExitFuncas(c *FuncasContext)

	// ExitSent_if is called when exiting the sent_if production.
	ExitSent_if(c *Sent_ifContext)

	// ExitImprimir is called when exiting the imprimir production.
	ExitImprimir(c *ImprimirContext)

	// ExitCasteo is called when exiting the casteo production.
	ExitCasteo(c *CasteoContext)

	// ExitDeclaracion is called when exiting the declaracion production.
	ExitDeclaracion(c *DeclaracionContext)

	// ExitEtiquetas is called when exiting the etiquetas production.
	ExitEtiquetas(c *EtiquetasContext)

	// ExitLlamada is called when exiting the llamada production.
	ExitLlamada(c *LlamadaContext)

	// ExitL_bloque is called when exiting the l_bloque production.
	ExitL_bloque(c *L_bloqueContext)

	// ExitBloque is called when exiting the bloque production.
	ExitBloque(c *BloqueContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitAsignaciones is called when exiting the asignaciones production.
	ExitAsignaciones(c *AsignacionesContext)

	// ExitExpre_relacional is called when exiting the expre_relacional production.
	ExitExpre_relacional(c *Expre_relacionalContext)

	// ExitExpre_aritmetica is called when exiting the expre_aritmetica production.
	ExitExpre_aritmetica(c *Expre_aritmeticaContext)

	// ExitDatos is called when exiting the datos production.
	ExitDatos(c *DatosContext)
}
