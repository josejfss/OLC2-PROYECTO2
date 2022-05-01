// Code generated from db_rustparser.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // db_rustparser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// db_rustparserListener is a complete listener for a parse tree produced by db_rustparser.
type db_rustparserListener interface {
	antlr.ParseTreeListener

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterInstrucciones is called when entering the instrucciones production.
	EnterInstrucciones(c *InstruccionesContext)

	// EnterInstruccion is called when entering the instruccion production.
	EnterInstruccion(c *InstruccionContext)

	// EnterStucts is called when entering the stucts production.
	EnterStucts(c *StuctsContext)

	// EnterL_structs is called when entering the l_structs production.
	EnterL_structs(c *L_structsContext)

	// EnterDefstruct is called when entering the defstruct production.
	EnterDefstruct(c *DefstructContext)

	// EnterFuncas is called when entering the funcas production.
	EnterFuncas(c *FuncasContext)

	// EnterPubli is called when entering the publi production.
	EnterPubli(c *PubliContext)

	// EnterParametros is called when entering the parametros production.
	EnterParametros(c *ParametrosContext)

	// EnterDefiniciones is called when entering the definiciones production.
	EnterDefiniciones(c *DefinicionesContext)

	// EnterValref is called when entering the valref production.
	EnterValref(c *ValrefContext)

	// EnterTipos_funciones is called when entering the tipos_funciones production.
	EnterTipos_funciones(c *Tipos_funcionesContext)

	// EnterSent_if is called when entering the sent_if production.
	EnterSent_if(c *Sent_ifContext)

	// EnterSent_ifelse is called when entering the sent_ifelse production.
	EnterSent_ifelse(c *Sent_ifelseContext)

	// EnterSent_match is called when entering the sent_match production.
	EnterSent_match(c *Sent_matchContext)

	// EnterL_matches is called when entering the l_matches production.
	EnterL_matches(c *L_matchesContext)

	// EnterMatches is called when entering the matches production.
	EnterMatches(c *MatchesContext)

	// EnterL_mat_con is called when entering the l_mat_con production.
	EnterL_mat_con(c *L_mat_conContext)

	// EnterLoop is called when entering the loop production.
	EnterLoop(c *LoopContext)

	// EnterC_while is called when entering the c_while production.
	EnterC_while(c *C_whileContext)

	// EnterFor_in is called when entering the for_in production.
	EnterFor_in(c *For_inContext)

	// EnterImprimir is called when entering the imprimir production.
	EnterImprimir(c *ImprimirContext)

	// EnterDeclaracion is called when entering the declaracion production.
	EnterDeclaracion(c *DeclaracionContext)

	// EnterMutable is called when entering the mutable production.
	EnterMutable(c *MutableContext)

	// EnterTamvector is called when entering the tamvector production.
	EnterTamvector(c *TamvectorContext)

	// EnterTamarre is called when entering the tamarre production.
	EnterTamarre(c *TamarreContext)

	// EnterTipado is called when entering the tipado production.
	EnterTipado(c *TipadoContext)

	// EnterTipos_vectorarre is called when entering the tipos_vectorarre production.
	EnterTipos_vectorarre(c *Tipos_vectorarreContext)

	// EnterTipos is called when entering the tipos production.
	EnterTipos(c *TiposContext)

	// EnterL_asigstruct is called when entering the l_asigstruct production.
	EnterL_asigstruct(c *L_asigstructContext)

	// EnterAsignacionstruct is called when entering the asignacionstruct production.
	EnterAsignacionstruct(c *AsignacionstructContext)

	// EnterFn_vector is called when entering the fn_vector production.
	EnterFn_vector(c *Fn_vectorContext)

	// EnterAsignacion is called when entering the asignacion production.
	EnterAsignacion(c *AsignacionContext)

	// EnterLlamada is called when entering the llamada production.
	EnterLlamada(c *LlamadaContext)

	// EnterTransferencia is called when entering the transferencia production.
	EnterTransferencia(c *TransferenciaContext)

	// EnterRetorno is called when entering the retorno production.
	EnterRetorno(c *RetornoContext)

	// EnterRomper is called when entering the romper production.
	EnterRomper(c *RomperContext)

	// EnterContinuar is called when entering the continuar production.
	EnterContinuar(c *ContinuarContext)

	// EnterL_bloque is called when entering the l_bloque production.
	EnterL_bloque(c *L_bloqueContext)

	// EnterBloque is called when entering the bloque production.
	EnterBloque(c *BloqueContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterList_expres is called when entering the list_expres production.
	EnterList_expres(c *List_expresContext)

	// EnterVectores_inicio is called when entering the vectores_inicio production.
	EnterVectores_inicio(c *Vectores_inicioContext)

	// EnterInicio_vect is called when entering the inicio_vect production.
	EnterInicio_vect(c *Inicio_vectContext)

	// EnterResumen_vect is called when entering the resumen_vect production.
	EnterResumen_vect(c *Resumen_vectContext)

	// EnterArreglos_inicio is called when entering the arreglos_inicio production.
	EnterArreglos_inicio(c *Arreglos_inicioContext)

	// EnterInicializando_arreglo is called when entering the inicializando_arreglo production.
	EnterInicializando_arreglo(c *Inicializando_arregloContext)

	// EnterResumen_arre is called when entering the resumen_arre production.
	EnterResumen_arre(c *Resumen_arreContext)

	// EnterAccesso_arreglo is called when entering the accesso_arreglo production.
	EnterAccesso_arreglo(c *Accesso_arregloContext)

	// EnterLista_acceso is called when entering the lista_acceso production.
	EnterLista_acceso(c *Lista_accesoContext)

	// EnterAccess is called when entering the access production.
	EnterAccess(c *AccessContext)

	// EnterAcceso_vector is called when entering the acceso_vector production.
	EnterAcceso_vector(c *Acceso_vectorContext)

	// EnterIngreso_struct is called when entering the ingreso_struct production.
	EnterIngreso_struct(c *Ingreso_structContext)

	// EnterNativas_vector is called when entering the nativas_vector production.
	EnterNativas_vector(c *Nativas_vectorContext)

	// EnterFun_nativas is called when entering the fun_nativas production.
	EnterFun_nativas(c *Fun_nativasContext)

	// EnterExpre_logica is called when entering the expre_logica production.
	EnterExpre_logica(c *Expre_logicaContext)

	// EnterExpre_relacional is called when entering the expre_relacional production.
	EnterExpre_relacional(c *Expre_relacionalContext)

	// EnterExpre_aritmetica is called when entering the expre_aritmetica production.
	EnterExpre_aritmetica(c *Expre_aritmeticaContext)

	// EnterValores is called when entering the valores production.
	EnterValores(c *ValoresContext)

	// EnterCasteo_datos is called when entering the casteo_datos production.
	EnterCasteo_datos(c *Casteo_datosContext)

	// EnterDatos is called when entering the datos production.
	EnterDatos(c *DatosContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitInstrucciones is called when exiting the instrucciones production.
	ExitInstrucciones(c *InstruccionesContext)

	// ExitInstruccion is called when exiting the instruccion production.
	ExitInstruccion(c *InstruccionContext)

	// ExitStucts is called when exiting the stucts production.
	ExitStucts(c *StuctsContext)

	// ExitL_structs is called when exiting the l_structs production.
	ExitL_structs(c *L_structsContext)

	// ExitDefstruct is called when exiting the defstruct production.
	ExitDefstruct(c *DefstructContext)

	// ExitFuncas is called when exiting the funcas production.
	ExitFuncas(c *FuncasContext)

	// ExitPubli is called when exiting the publi production.
	ExitPubli(c *PubliContext)

	// ExitParametros is called when exiting the parametros production.
	ExitParametros(c *ParametrosContext)

	// ExitDefiniciones is called when exiting the definiciones production.
	ExitDefiniciones(c *DefinicionesContext)

	// ExitValref is called when exiting the valref production.
	ExitValref(c *ValrefContext)

	// ExitTipos_funciones is called when exiting the tipos_funciones production.
	ExitTipos_funciones(c *Tipos_funcionesContext)

	// ExitSent_if is called when exiting the sent_if production.
	ExitSent_if(c *Sent_ifContext)

	// ExitSent_ifelse is called when exiting the sent_ifelse production.
	ExitSent_ifelse(c *Sent_ifelseContext)

	// ExitSent_match is called when exiting the sent_match production.
	ExitSent_match(c *Sent_matchContext)

	// ExitL_matches is called when exiting the l_matches production.
	ExitL_matches(c *L_matchesContext)

	// ExitMatches is called when exiting the matches production.
	ExitMatches(c *MatchesContext)

	// ExitL_mat_con is called when exiting the l_mat_con production.
	ExitL_mat_con(c *L_mat_conContext)

	// ExitLoop is called when exiting the loop production.
	ExitLoop(c *LoopContext)

	// ExitC_while is called when exiting the c_while production.
	ExitC_while(c *C_whileContext)

	// ExitFor_in is called when exiting the for_in production.
	ExitFor_in(c *For_inContext)

	// ExitImprimir is called when exiting the imprimir production.
	ExitImprimir(c *ImprimirContext)

	// ExitDeclaracion is called when exiting the declaracion production.
	ExitDeclaracion(c *DeclaracionContext)

	// ExitMutable is called when exiting the mutable production.
	ExitMutable(c *MutableContext)

	// ExitTamvector is called when exiting the tamvector production.
	ExitTamvector(c *TamvectorContext)

	// ExitTamarre is called when exiting the tamarre production.
	ExitTamarre(c *TamarreContext)

	// ExitTipado is called when exiting the tipado production.
	ExitTipado(c *TipadoContext)

	// ExitTipos_vectorarre is called when exiting the tipos_vectorarre production.
	ExitTipos_vectorarre(c *Tipos_vectorarreContext)

	// ExitTipos is called when exiting the tipos production.
	ExitTipos(c *TiposContext)

	// ExitL_asigstruct is called when exiting the l_asigstruct production.
	ExitL_asigstruct(c *L_asigstructContext)

	// ExitAsignacionstruct is called when exiting the asignacionstruct production.
	ExitAsignacionstruct(c *AsignacionstructContext)

	// ExitFn_vector is called when exiting the fn_vector production.
	ExitFn_vector(c *Fn_vectorContext)

	// ExitAsignacion is called when exiting the asignacion production.
	ExitAsignacion(c *AsignacionContext)

	// ExitLlamada is called when exiting the llamada production.
	ExitLlamada(c *LlamadaContext)

	// ExitTransferencia is called when exiting the transferencia production.
	ExitTransferencia(c *TransferenciaContext)

	// ExitRetorno is called when exiting the retorno production.
	ExitRetorno(c *RetornoContext)

	// ExitRomper is called when exiting the romper production.
	ExitRomper(c *RomperContext)

	// ExitContinuar is called when exiting the continuar production.
	ExitContinuar(c *ContinuarContext)

	// ExitL_bloque is called when exiting the l_bloque production.
	ExitL_bloque(c *L_bloqueContext)

	// ExitBloque is called when exiting the bloque production.
	ExitBloque(c *BloqueContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitList_expres is called when exiting the list_expres production.
	ExitList_expres(c *List_expresContext)

	// ExitVectores_inicio is called when exiting the vectores_inicio production.
	ExitVectores_inicio(c *Vectores_inicioContext)

	// ExitInicio_vect is called when exiting the inicio_vect production.
	ExitInicio_vect(c *Inicio_vectContext)

	// ExitResumen_vect is called when exiting the resumen_vect production.
	ExitResumen_vect(c *Resumen_vectContext)

	// ExitArreglos_inicio is called when exiting the arreglos_inicio production.
	ExitArreglos_inicio(c *Arreglos_inicioContext)

	// ExitInicializando_arreglo is called when exiting the inicializando_arreglo production.
	ExitInicializando_arreglo(c *Inicializando_arregloContext)

	// ExitResumen_arre is called when exiting the resumen_arre production.
	ExitResumen_arre(c *Resumen_arreContext)

	// ExitAccesso_arreglo is called when exiting the accesso_arreglo production.
	ExitAccesso_arreglo(c *Accesso_arregloContext)

	// ExitLista_acceso is called when exiting the lista_acceso production.
	ExitLista_acceso(c *Lista_accesoContext)

	// ExitAccess is called when exiting the access production.
	ExitAccess(c *AccessContext)

	// ExitAcceso_vector is called when exiting the acceso_vector production.
	ExitAcceso_vector(c *Acceso_vectorContext)

	// ExitIngreso_struct is called when exiting the ingreso_struct production.
	ExitIngreso_struct(c *Ingreso_structContext)

	// ExitNativas_vector is called when exiting the nativas_vector production.
	ExitNativas_vector(c *Nativas_vectorContext)

	// ExitFun_nativas is called when exiting the fun_nativas production.
	ExitFun_nativas(c *Fun_nativasContext)

	// ExitExpre_logica is called when exiting the expre_logica production.
	ExitExpre_logica(c *Expre_logicaContext)

	// ExitExpre_relacional is called when exiting the expre_relacional production.
	ExitExpre_relacional(c *Expre_relacionalContext)

	// ExitExpre_aritmetica is called when exiting the expre_aritmetica production.
	ExitExpre_aritmetica(c *Expre_aritmeticaContext)

	// ExitValores is called when exiting the valores production.
	ExitValores(c *ValoresContext)

	// ExitCasteo_datos is called when exiting the casteo_datos production.
	ExitCasteo_datos(c *Casteo_datosContext)

	// ExitDatos is called when exiting the datos production.
	ExitDatos(c *DatosContext)
}
