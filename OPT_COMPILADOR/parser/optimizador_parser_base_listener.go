// Code generated from optimizador_parser.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // optimizador_parser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// Baseoptimizador_parserListener is a complete listener for a parse tree produced by optimizador_parser.
type Baseoptimizador_parserListener struct{}

var _ optimizador_parserListener = &Baseoptimizador_parserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *Baseoptimizador_parserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *Baseoptimizador_parserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *Baseoptimizador_parserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *Baseoptimizador_parserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStart is called when production start is entered.
func (s *Baseoptimizador_parserListener) EnterStart(ctx *StartContext) {}

// ExitStart is called when production start is exited.
func (s *Baseoptimizador_parserListener) ExitStart(ctx *StartContext) {}

// EnterInstrucss is called when production instrucss is entered.
func (s *Baseoptimizador_parserListener) EnterInstrucss(ctx *InstrucssContext) {}

// ExitInstrucss is called when production instrucss is exited.
func (s *Baseoptimizador_parserListener) ExitInstrucss(ctx *InstrucssContext) {}

// EnterInstrs is called when production instrs is entered.
func (s *Baseoptimizador_parserListener) EnterInstrs(ctx *InstrsContext) {}

// ExitInstrs is called when production instrs is exited.
func (s *Baseoptimizador_parserListener) ExitInstrs(ctx *InstrsContext) {}

// EnterEncabezado is called when production encabezado is entered.
func (s *Baseoptimizador_parserListener) EnterEncabezado(ctx *EncabezadoContext) {}

// ExitEncabezado is called when production encabezado is exited.
func (s *Baseoptimizador_parserListener) ExitEncabezado(ctx *EncabezadoContext) {}

// EnterL_temporales is called when production l_temporales is entered.
func (s *Baseoptimizador_parserListener) EnterL_temporales(ctx *L_temporalesContext) {}

// ExitL_temporales is called when production l_temporales is exited.
func (s *Baseoptimizador_parserListener) ExitL_temporales(ctx *L_temporalesContext) {}

// EnterL_funcas is called when production l_funcas is entered.
func (s *Baseoptimizador_parserListener) EnterL_funcas(ctx *L_funcasContext) {}

// ExitL_funcas is called when production l_funcas is exited.
func (s *Baseoptimizador_parserListener) ExitL_funcas(ctx *L_funcasContext) {}

// EnterFuncas is called when production funcas is entered.
func (s *Baseoptimizador_parserListener) EnterFuncas(ctx *FuncasContext) {}

// ExitFuncas is called when production funcas is exited.
func (s *Baseoptimizador_parserListener) ExitFuncas(ctx *FuncasContext) {}

// EnterSent_if is called when production sent_if is entered.
func (s *Baseoptimizador_parserListener) EnterSent_if(ctx *Sent_ifContext) {}

// ExitSent_if is called when production sent_if is exited.
func (s *Baseoptimizador_parserListener) ExitSent_if(ctx *Sent_ifContext) {}

// EnterImprimir is called when production imprimir is entered.
func (s *Baseoptimizador_parserListener) EnterImprimir(ctx *ImprimirContext) {}

// ExitImprimir is called when production imprimir is exited.
func (s *Baseoptimizador_parserListener) ExitImprimir(ctx *ImprimirContext) {}

// EnterCasteo is called when production casteo is entered.
func (s *Baseoptimizador_parserListener) EnterCasteo(ctx *CasteoContext) {}

// ExitCasteo is called when production casteo is exited.
func (s *Baseoptimizador_parserListener) ExitCasteo(ctx *CasteoContext) {}

// EnterDeclaracion is called when production declaracion is entered.
func (s *Baseoptimizador_parserListener) EnterDeclaracion(ctx *DeclaracionContext) {}

// ExitDeclaracion is called when production declaracion is exited.
func (s *Baseoptimizador_parserListener) ExitDeclaracion(ctx *DeclaracionContext) {}

// EnterEtiquetas is called when production etiquetas is entered.
func (s *Baseoptimizador_parserListener) EnterEtiquetas(ctx *EtiquetasContext) {}

// ExitEtiquetas is called when production etiquetas is exited.
func (s *Baseoptimizador_parserListener) ExitEtiquetas(ctx *EtiquetasContext) {}

// EnterLlamada is called when production llamada is entered.
func (s *Baseoptimizador_parserListener) EnterLlamada(ctx *LlamadaContext) {}

// ExitLlamada is called when production llamada is exited.
func (s *Baseoptimizador_parserListener) ExitLlamada(ctx *LlamadaContext) {}

// EnterL_bloque is called when production l_bloque is entered.
func (s *Baseoptimizador_parserListener) EnterL_bloque(ctx *L_bloqueContext) {}

// ExitL_bloque is called when production l_bloque is exited.
func (s *Baseoptimizador_parserListener) ExitL_bloque(ctx *L_bloqueContext) {}

// EnterBloque is called when production bloque is entered.
func (s *Baseoptimizador_parserListener) EnterBloque(ctx *BloqueContext) {}

// ExitBloque is called when production bloque is exited.
func (s *Baseoptimizador_parserListener) ExitBloque(ctx *BloqueContext) {}

// EnterExpression is called when production expression is entered.
func (s *Baseoptimizador_parserListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *Baseoptimizador_parserListener) ExitExpression(ctx *ExpressionContext) {}

// EnterAsignaciones is called when production asignaciones is entered.
func (s *Baseoptimizador_parserListener) EnterAsignaciones(ctx *AsignacionesContext) {}

// ExitAsignaciones is called when production asignaciones is exited.
func (s *Baseoptimizador_parserListener) ExitAsignaciones(ctx *AsignacionesContext) {}

// EnterExpre_relacional is called when production expre_relacional is entered.
func (s *Baseoptimizador_parserListener) EnterExpre_relacional(ctx *Expre_relacionalContext) {}

// ExitExpre_relacional is called when production expre_relacional is exited.
func (s *Baseoptimizador_parserListener) ExitExpre_relacional(ctx *Expre_relacionalContext) {}

// EnterExpre_aritmetica is called when production expre_aritmetica is entered.
func (s *Baseoptimizador_parserListener) EnterExpre_aritmetica(ctx *Expre_aritmeticaContext) {}

// ExitExpre_aritmetica is called when production expre_aritmetica is exited.
func (s *Baseoptimizador_parserListener) ExitExpre_aritmetica(ctx *Expre_aritmeticaContext) {}

// EnterDatos is called when production datos is entered.
func (s *Baseoptimizador_parserListener) EnterDatos(ctx *DatosContext) {}

// ExitDatos is called when production datos is exited.
func (s *Baseoptimizador_parserListener) ExitDatos(ctx *DatosContext) {}
