parser grammar db_rustparser;

options {
    tokenVocab = db_rustlexer;
}

@header {
    import "OLC2-PROYECTO2/COMPILADOR/INTERFACES"
    import "OLC2-PROYECTO2/COMPILADOR/AST/EXPRESIONES/PRIMITIVO"
    import "OLC2-PROYECTO2/COMPILADOR/AST/EXPRESIONES/VARIABLES"
    import "OLC2-PROYECTO2/COMPILADOR/AST/EXPRESIONES/CASTEOS"
    import "OLC2-PROYECTO2/COMPILADOR/AST/EXPRESIONES/NATIVAS"
    import "OLC2-PROYECTO2/COMPILADOR/AST/EXPRESIONES/ARITMETICA"
    import "OLC2-PROYECTO2/COMPILADOR/AST/EXPRESIONES/LOGICA"
    import "OLC2-PROYECTO2/COMPILADOR/AST/EXPRESIONES/RELACIONAL"
    import "OLC2-PROYECTO2/COMPILADOR/AST/EXPRESIONES/INICIANDO_VECT"
    import "OLC2-PROYECTO2/COMPILADOR/AST/EXPRESIONES/INICIANDO_ARRE"
    import "OLC2-PROYECTO2/COMPILADOR/AST/EXPRESIONES/ACCESO_ARRE"
    import "OLC2-PROYECTO2/COMPILADOR/AST/EXPRESIONES/ACCESO_VECT"
    import "OLC2-PROYECTO2/COMPILADOR/AST/EXPRESIONES/ACCESO_STRUCT"
    import "OLC2-PROYECTO2/COMPILADOR/AST/EXPRESIONES/NATIVAS_VECT"
    import "OLC2-PROYECTO2/COMPILADOR/AST/INSTRUCCIONES/IMPRESION"
    import "OLC2-PROYECTO2/COMPILADOR/AST/INSTRUCCIONES/S_CONDICIONAL"
    import "OLC2-PROYECTO2/COMPILADOR/AST/INSTRUCCIONES/S_MATCH"
    import "OLC2-PROYECTO2/COMPILADOR/AST/INSTRUCCIONES/C_LOOP"
    import "OLC2-PROYECTO2/COMPILADOR/AST/INSTRUCCIONES/C_WHILE"
    import "OLC2-PROYECTO2/COMPILADOR/AST/INSTRUCCIONES/C_FORIN"
    import "OLC2-PROYECTO2/COMPILADOR/AST/INSTRUCCIONES/DECLARACIONES"
    import "OLC2-PROYECTO2/COMPILADOR/AST/INSTRUCCIONES/DECLARACIONVAR"
    import "OLC2-PROYECTO2/COMPILADOR/AST/INSTRUCCIONES/DECLARACIONVECT"
    import "OLC2-PROYECTO2/COMPILADOR/AST/INSTRUCCIONES/FUN_NAT_VECT"
    import "OLC2-PROYECTO2/COMPILADOR/AST/INSTRUCCIONES/DECLARACIONARRE"
    import "OLC2-PROYECTO2/COMPILADOR/AST/INSTRUCCIONES/S_TRANSFERENCIA"
    import "OLC2-PROYECTO2/COMPILADOR/AST/INSTRUCCIONES/ASIGNACIONES"
    import "OLC2-PROYECTO2/COMPILADOR/AST/INSTRUCCIONES/FUNCIONES/PARAMETROS"
    import "OLC2-PROYECTO2/COMPILADOR/AST/INSTRUCCIONES/FUNCIONES/TIPOS"
    import "OLC2-PROYECTO2/COMPILADOR/AST/INSTRUCCIONES/FUNCIONES"    
    import "OLC2-PROYECTO2/COMPILADOR/AST/INSTRUCCIONES/STRUCTS"     
    import "OLC2-PROYECTO2/COMPILADOR/AST/INSTRUCCIONES/MODULO"    
	  import "OLC2-PROYECTO2/COMPILADOR/SIMBOLOS"
    import arrayList "github.com/colegno/arraylist"
}

start returns[*arrayList.List lista]
  : instrucciones {$lista = $instrucciones.lis}
;

instrucciones returns[*arrayList.List lis]
  @init {
      $lis =  arrayList.New()
  }: 
  e +=instruccion* {
    listInt := localctx.(*InstruccionesContext).GetE()
    for _, e := range listInt {
      $lis.Add(e.GetInstr())
    }
  }
;

instruccion returns[interfaces.Instruccion instr]
  : declaracion       {$instr = $declaracion.instr}
  | funcas            {$instr = $funcas.instr}
  | stucts            {$instr = $stucts.instr}
  | modulos           {$instr = $modulos.instr}
;

modulos returns[interfaces.Instruccion instr]
  : publi TK_MOD TK_IDENTIFICADOR TK_LI l_modulos TK_LD {
    $instr = modulo.Nmods($publi.m, $TK_IDENTIFICADOR.text, $l_modulos.lmod, $TK_IDENTIFICADOR.line, localctx.(*ModulosContext).Get_TK_IDENTIFICADOR().GetColumn())
  }
;

l_modulos returns[*arrayList.List lmod]
  @init {
    $lmod =  arrayList.New()
  }
  : mods = l_modulos modulitos {
    $mods.lmod.Add($modulitos.instr)
    $lmod=$mods.lmod
  }
  | modulitos { $lmod.Add($modulitos.instr) }
;

modulitos returns[interfaces.Instruccion instr]
  : funcas        {$instr = $funcas.instr}
  | stucts        {$instr = $stucts.instr}
  | modulos       {$instr = $modulos.instr}
  | declaracion   {$instr = $declaracion.instr}
;

/* -------------------------------------------------------------------ESTRUCTURAS------------------------------------------------------------------------------- */
stucts returns[interfaces.Instruccion instr]
  : TK_STRUCT TK_IDENTIFICADOR TK_LI l_structs TK_LD {
    $instr = structs.Nstructrust($TK_IDENTIFICADOR.text, $l_structs.lstr, $TK_IDENTIFICADOR.line, localctx.(*StuctsContext).Get_TK_IDENTIFICADOR().GetColumn())
  }
;

l_structs returns[*arrayList.List lstr]
  @init {
    $lstr =  arrayList.New()
  }
  : lista = l_structs TK_COMA defstruct {
    $lista.lstr.Add($defstruct.stru)
    $lstr=$lista.lstr
  }
  | defstruct { $lstr.Add($defstruct.stru) }
;

defstruct returns[structs.DefStruct stru]
  : TK_IDENTIFICADOR tipado {
    $stru = structs.Ndefstruct($TK_IDENTIFICADOR.text, $tipado.tip)
  }
;

/* -------------------------------------------------------------------FUNCIONES--------------------------------------------------------------------------------- */
funcas returns[interfaces.Instruccion instr]
  @init {
    listaparametros :=  arrayList.New()
  }
  : publi TK_FUNCION TK_MAIN TK_PI TK_PD TK_LI l_bloque TK_LD {
    tf := tipos.Ntiposfunca("null", simbolos.NULL, nil, nil)
    $instr = funciones.Nfunciones($publi.m, $TK_MAIN.text,listaparametros,tf, $l_bloque.lbloque, $TK_MAIN.line, localctx.(*FuncasContext).Get_TK_MAIN().GetColumn())
  }
  | publi TK_FUNCION TK_IDENTIFICADOR TK_PI TK_PD tipos_funciones TK_LI l_bloque TK_LD {
    $instr = funciones.Nfunciones($publi.m, $TK_IDENTIFICADOR.text,listaparametros,$tipos_funciones.tf, $l_bloque.lbloque, $TK_IDENTIFICADOR.line, localctx.(*FuncasContext).Get_TK_IDENTIFICADOR().GetColumn())
  }
  | publi TK_FUNCION TK_IDENTIFICADOR TK_PI parametros TK_PD tipos_funciones TK_LI l_bloque TK_LD {
    $instr = funciones.Nfunciones($publi.m, $TK_IDENTIFICADOR.text,$parametros.lparame,$tipos_funciones.tf, $l_bloque.lbloque, $TK_IDENTIFICADOR.line, localctx.(*FuncasContext).Get_TK_IDENTIFICADOR().GetColumn())
  }
;

publi returns[bool m]
  : TK_PUBLICO  { $m = true }
  |             { $m = false }
;

parametros returns[*arrayList.List lparame]
  @init {
    $lparame =  arrayList.New()
  }
  : par = parametros TK_COMA definiciones {
    $par.lparame.Add($definiciones.pa)
    $lparame=$par.lparame
  }
  | definiciones { $lparame.Add($definiciones.pa)}
;

definiciones returns[parametros.Param pa]
  : TK_IDENTIFICADOR TK_DP valref tipos {
    $pa = parametros.Nuevo_parametro($TK_IDENTIFICADOR.text, $valref.val, "variable", $tipos.tip, nil, nil, $TK_IDENTIFICADOR.line, localctx.(*DefinicionesContext).Get_TK_IDENTIFICADOR().GetColumn() )
  }
  | TK_IDENTIFICADOR TK_DP valref TK_VECTOR tipos_vectorarre {
    $pa = parametros.Nuevo_parametro($TK_IDENTIFICADOR.text, $valref.val, "vector", simbolos.NULL, $tipos_vectorarre.p, nil, $TK_IDENTIFICADOR.line, localctx.(*DefinicionesContext).Get_TK_IDENTIFICADOR().GetColumn() )
  }
  | TK_IDENTIFICADOR TK_DP valref TK_CI tipos_vectorarre TK_PYC expression TK_CD {
    $pa = parametros.Nuevo_parametro($TK_IDENTIFICADOR.text, $valref.val, "arreglo", simbolos.NULL, $tipos_vectorarre.p, $expression.p, $TK_IDENTIFICADOR.line, localctx.(*DefinicionesContext).Get_TK_IDENTIFICADOR().GetColumn() )
  }
;

valref returns[int val]
  : TK_MUT      {$val=1}
  | TK_DIRMUT   {$val=2}
  |             {$val=3}
;

tipos_funciones returns[tipos.TiposFunca tf]
  : TK_IGUAL TK_MAYOR tipos {
    $tf = tipos.Ntiposfunca("variable", $tipos.tip, nil, nil)
  }
  | TK_IGUAL TK_MAYOR TK_VECTOR tipos_vectorarre {
    $tf = tipos.Ntiposfunca("vector", simbolos.NULL, $tipos_vectorarre.p, nil)
  }
  | TK_IGUAL TK_MAYOR TK_CI tipos_vectorarre TK_PYC expression TK_CD {
    $tf = tipos.Ntiposfunca("arreglo", simbolos.NULL, $tipos_vectorarre.p, $expression.p)
  }
  | {
    $tf = tipos.Ntiposfunca("null", simbolos.NULL, nil, nil)
  }
;

/* -------------------------------------------------------------------CICLOS Y SENTENCIAS---------------------------------------------------------------------- */
/* CONDICIONAL IF */ 
sent_if returns [interfaces.Instruccion instr, interfaces.Expresion p]
  @init {
    lista :=  arrayList.New()
  }
  : TK_IF expression TK_LI l_bloque TK_LD { 
    $instr = scondicional.Nsi($expression.p, $l_bloque.lbloque, lista)
    $p = scondicional.Nsi($expression.p, $l_bloque.lbloque, lista)
  }  
  | TK_IF expression TK_LI inst = l_bloque TK_LD TK_ELSE TK_LI els = l_bloque TK_LD { 
    $instr = scondicional.Nsi($expression.p, $inst.lbloque, $els.lbloque)
    $p = scondicional.Nsi($expression.p, $inst.lbloque, $els.lbloque)
  }
  | TK_IF expression TK_LI l_bloque TK_LD TK_ELSE sent_ifelse { 
    $instr = scondicional.Nsi($expression.p, $l_bloque.lbloque, $sent_ifelse.lia)
    $p = scondicional.Nsi($expression.p, $l_bloque.lbloque, $sent_ifelse.lia)
  }
;

/* ELSE IF */
sent_ifelse returns[*arrayList.List lia]
  @init {
    $lia =  arrayList.New()
  }
  : sent_if {
    $lia.Add($sent_if.instr)
  }
;

/* SENTENCIA MATCH */
sent_match returns [interfaces.Instruccion instr, interfaces.Expresion p]
  : TK_MATCH expression TK_LI l_matches TK_LD { 
    $instr = smatch.Nmatch($expression.p,$l_matches.lmatch)
    $p = smatch.Nmatch($expression.p,$l_matches.lmatch)
  }
;

/* LISTA MATCHES */
l_matches returns[*arrayList.List lmatch]
  @init {
    $lmatch =  arrayList.New()
  }
  : matchs = l_matches matches {
    $matchs.lmatch.Add($matches.instr)
    $lmatch=$matchs.lmatch
  }
  | matches { $lmatch.Add($matches.instr)}
;

/* MATCH */
matches returns [smatch.Matches instr]
  @init {
    listica :=  arrayList.New()
  }
  : expression TK_IGUAL TK_MAYOR TK_LI l_bloque TK_LD { 
    $instr = smatch.Nmatches($expression.p, $l_bloque.lbloque, false, listica) 
  }
  | l_mat_con TK_IGUAL TK_MAYOR TK_LI l_bloque TK_LD { 
    $instr = smatch.Nmatches(primitivo.Nuevo_Dato_Primitivo("holas",simbolos.YTEXTO), $l_bloque.lbloque, true, $l_mat_con.condis) 
  }
;

/* MATCH CON CONDICIONES */
l_mat_con returns[*arrayList.List condis]
  @init {
    $condis =  arrayList.New()
  }
  : e = l_mat_con TK_BARRA expression {
    $e.condis.Add($expression.p)
    $condis=$e.condis
  }
  | expression { $condis.Add($expression.p)}
;

/* CICLO LOOP */
loop returns[interfaces.Instruccion instr, interfaces.Expresion p]
  : TK_LOOP TK_LI l_bloque TK_LD {
    $instr = cloop.Nloop($l_bloque.lbloque, $TK_LOOP.line, localctx.(*LoopContext).Get_TK_LOOP().GetColumn())
    $p = cloop.Nloop($l_bloque.lbloque, $TK_LOOP.line, localctx.(*LoopContext).Get_TK_LOOP().GetColumn())
  }
;

/* CICLO WHILE */
c_while returns[interfaces.Instruccion instr]
  : TK_WHILE expression TK_LI l_bloque TK_LD { 
    $instr = cwhile.Nwhile($expression.p, $l_bloque.lbloque, $TK_WHILE.line, localctx.(*C_whileContext).Get_TK_WHILE().GetColumn() ) 
    }
;

/* CICLO FORIN */
for_in returns[interfaces.Instruccion instr]
  : TK_FOR TK_IDENTIFICADOR TK_IN inf = expression TK_LI l_bloque TK_LD {
    $instr = cforin.NforinVA($TK_IDENTIFICADOR.text, $inf.p, $l_bloque.lbloque, $TK_IDENTIFICADOR.line, localctx.(*For_inContext).Get_TK_IDENTIFICADOR().GetColumn() ) 
  }
  | TK_FOR TK_IDENTIFICADOR TK_IN inf = expression TK_PUNTO TK_PUNTO sup = expression TK_LI l_bloque TK_LD {
    $instr = cforin.Nfin($TK_IDENTIFICADOR.text, $inf.p, $sup.p ,$l_bloque.lbloque, $TK_IDENTIFICADOR.line, localctx.(*For_inContext).Get_TK_IDENTIFICADOR().GetColumn() ) 
  }
;

/* ------------------------------------------------------------IMPRESION EN CONSOLA------------------------------------------------------------------------ */
/* IMPRESION */
imprimir returns[interfaces.Instruccion instr]
  : TK_PRINTLN TK_PI expression TK_PD TK_PYC {
    $instr = impresion.Nimprimir($expression.p, $TK_PRINTLN.line, localctx.(*ImprimirContext).Get_TK_PRINTLN().GetColumn())
  }
  | TK_PRINTLN TK_PI expression TK_COMA list_expres TK_PD TK_PYC {
    $instr = impresion.Nimpres($expression.p, $list_expres.lis_expres, $TK_PRINTLN.line, localctx.(*ImprimirContext).Get_TK_PRINTLN().GetColumn())
  }
;

/* ----------------------------------------------------------DECLARAR VARIABLES, VECTORES Y ARREGLOS--------------------------------------------------------- */
declaracion returns[interfaces.Instruccion instr]
  : TK_LET mutable TK_IDENTIFICADOR TK_IGUAL noms = TK_IDENTIFICADOR TK_LI l_asigstruct TK_LD TK_PYC {
    $instr = declaraciones.Ndeclarar($mutable.m, $TK_IDENTIFICADOR.text, $noms.text, $l_asigstruct.otralista, $TK_IDENTIFICADOR.line, localctx.(*DeclaracionContext).Get_TK_IDENTIFICADOR().GetColumn())
  }
  | TK_LET mutable TK_IDENTIFICADOR tipado TK_IGUAL expression TK_PYC {
    $instr = declaracionvar.Ndeclaracion($mutable.m, $TK_IDENTIFICADOR.text, $tipado.tip, $expression.p, $TK_IDENTIFICADOR.line, localctx.(*DeclaracionContext).Get_TK_IDENTIFICADOR().GetColumn())
  }
  | TK_LET mutable TK_IDENTIFICADOR TK_DP tamarre TK_IGUAL expression TK_PYC {
    $instr = declaracionarre.Ndeclaarre($mutable.m, $TK_IDENTIFICADOR.text, $tamarre.tam_arre, $expression.p, $TK_IDENTIFICADOR.line, localctx.(*DeclaracionContext).Get_TK_IDENTIFICADOR().GetColumn())
  }
  | TK_LET mutable TK_IDENTIFICADOR TK_DP tamvector TK_IGUAL expression TK_PYC {
    $instr = declaracionvect.Ndeclavector($mutable.m, $TK_IDENTIFICADOR.text, $tamvector.tam_vect, $expression.p, $TK_IDENTIFICADOR.line, localctx.(*DeclaracionContext).Get_TK_IDENTIFICADOR().GetColumn())
  }
  | TK_LET mutable TK_IDENTIFICADOR TK_DP TK_VECTOR TK_MENOR tipos_vectorarre TK_MAYOR TK_IGUAL TK_VECTOR TK_NEW TK_PYC {
    $instr = funnatvect.Nnewvect($mutable.m, $TK_IDENTIFICADOR.text, $tipos_vectorarre.p, $TK_IDENTIFICADOR.line, localctx.(*DeclaracionContext).Get_TK_IDENTIFICADOR().GetColumn())
  }
  | TK_LET mutable TK_IDENTIFICADOR TK_DP TK_VECTOR TK_MENOR tipos_vectorarre TK_MAYOR TK_IGUAL TK_VECTOR TK_WCAPACITY TK_PI expression TK_PD TK_PYC {
    $instr = funnatvect.Nwcvect($mutable.m, $TK_IDENTIFICADOR.text, $tipos_vectorarre.p, $expression.p, $TK_IDENTIFICADOR.line, localctx.(*DeclaracionContext).Get_TK_IDENTIFICADOR().GetColumn())
  }
;

/* VARIIABLE, ARREGLO, VECTOR MUTABLE */
mutable returns[bool m]
  : TK_MUT  { $m = true }
  |         { $m = false }
;

/* DIMENSIONES DEL VECTOR */
tamvector returns[*arrayList.List tam_vect]
  @init {
    $tam_vect =  arrayList.New()
  }
  : TK_VECTOR TK_MENOR tv = tamvector TK_MAYOR {
    $tv.tam_vect.Add($TK_MAYOR.text)
    $tam_vect=$tv.tam_vect
  }
  | TK_VECTOR TK_MENOR tipos_vectorarre TK_MAYOR { 
    $tam_vect.Add($tipos_vectorarre.p)
  }
;

/* DIMENSIONES DEL ARREGLO */
tamarre returns[*arrayList.List tam_arre]
  @init {
    $tam_arre =  arrayList.New()
  }
  : TK_CI l = tamarre TK_PYC expression TK_CD {
    $l.tam_arre.Add($expression.p)
    $tam_arre=$l.tam_arre
  }
  | TK_CI tipos_vectorarre TK_PYC expression TK_CD {
    guardar := declaracionarre.Ntipodeclarre($tipos_vectorarre.p,$expression.p)
    $tam_arre.Add(guardar)
  }
;

/* TIPO DE LA VARIABLE */
tipado returns[simbolos.TipoExpresion tip]
  : TK_DP TK_TIPOINT      {$tip=simbolos.INTEGER}
  | TK_DP TK_TIPOFLOAT    {$tip=simbolos.FLOAT}
  | TK_DP TK_TIPOBOOL     {$tip=simbolos.BOOLEAN}
  | TK_DP TK_TIPOCHAR     {$tip=simbolos.CHAR}
  | TK_DP TK_TIPOSTRING   {$tip=simbolos.TEXTO}
  | TK_DP TK_DIRSTRING    {$tip=simbolos.YTEXTO}
  |                       {$tip=simbolos.NULL}
;

/* TIPO VECTOR */
tipos_vectorarre returns[interfaces.Expresion p]
  : TK_TIPOINT   {
    $p = primitivo.Nuevo_Dato_Primitivo (0,simbolos.INTEGER) 
  }
  | TK_TIPOFLOAT {
    $p = primitivo.Nuevo_Dato_Primitivo (0.0,simbolos.FLOAT)
  }
  | TK_TIPOBOOL  {
    $p = primitivo.Nuevo_Dato_Primitivo(true,simbolos.BOOLEAN)
  }
  | TK_TIPOCHAR  {
    $p = primitivo.Nuevo_Dato_Primitivo(0,simbolos.CHAR)
  }
  | TK_TIPOSTRING {
    $p = primitivo.Nuevo_Dato_Primitivo("",simbolos.TEXTO)
  }
  | TK_DIRSTRING  {
    $p = primitivo.Nuevo_Dato_Primitivo("",simbolos.YTEXTO)
  }
  | TK_IDENTIFICADOR {
    $p = primitivo.Nuevo_Dato_Primitivo($TK_IDENTIFICADOR.text,simbolos.STRUCT)
  }
  |  {
    $p = primitivo.Nuevo_Dato_Primitivo("",simbolos.NULL)
  }
;

/* TIPO */
tipos returns[simbolos.TipoExpresion tip]
  : TK_TIPOINT        {$tip=simbolos.INTEGER  }
  | TK_TIPOFLOAT      {$tip=simbolos.FLOAT}
  | TK_TIPOBOOL       {$tip=simbolos.BOOLEAN}
  | TK_TIPOCHAR       {$tip=simbolos.CHAR}
  | TK_TIPOSTRING     {$tip=simbolos.TEXTO}
  | TK_DIRSTRING      {$tip=simbolos.YTEXTO}
  | TK_IDENTIFICADOR  {$tip=simbolos.STRUCT}
  |                   {$tip=simbolos.NULL}
;

/* LISTA STRUCTS */
l_asigstruct returns[*arrayList.List otralista]
  @init {
    $otralista =  arrayList.New()
  }
  : la = l_asigstruct TK_COMA asignacionstruct {
    $la.otralista.Add($asignacionstruct.astru)
    $otralista=$la.otralista
  }
  | asignacionstruct { $otralista.Add($asignacionstruct.astru) }
;

asignacionstruct returns [declaraciones.AsigStruct astru]
  : TK_IDENTIFICADOR TK_DP expression {
    $astru = declaraciones.Nasigstruct($TK_IDENTIFICADOR.text, $expression.p)
  }
;

/* ------------------------------------------------------------------FUNCIONES NATIVAS VECTORES----------------------------------------------------------- */
fn_vector returns[interfaces.Instruccion instr]
  : TK_IDENTIFICADOR TK_PUNTO TK_PUSH TK_PI expression TK_PD TK_PYC {
    $instr = funnatvect.Npushvect($TK_IDENTIFICADOR.text, $expression.p, $TK_IDENTIFICADOR.line, localctx.(*Fn_vectorContext).Get_TK_IDENTIFICADOR().GetColumn())
  }
  | TK_IDENTIFICADOR TK_PUNTO TK_INSERT TK_PI pos = expression TK_COMA val = expression TK_PD TK_PYC {
    $instr = funnatvect.Ninsertvect($TK_IDENTIFICADOR.text, $pos.p, $val.p, $TK_IDENTIFICADOR.line, localctx.(*Fn_vectorContext).Get_TK_IDENTIFICADOR().GetColumn())
  }
  | TK_IDENTIFICADOR TK_PUNTO TK_REMOVE TK_PI expression TK_PD TK_PYC {
    $instr = funnatvect.Nremovevect($TK_IDENTIFICADOR.text, $expression.p, $TK_IDENTIFICADOR.line, localctx.(*Fn_vectorContext).Get_TK_IDENTIFICADOR().GetColumn())
  }
;

/* ------------------------------------------------------------------ASIGNACION DE VARIABLES-------------------------------------------------------------- */
asignacion returns[interfaces.Instruccion instr]
  : TK_IDENTIFICADOR TK_IGUAL expression TK_PYC {
      str:= $TK_IDENTIFICADOR.text[0:len($TK_IDENTIFICADOR.text)]
      $instr = asignaciones.Nasignacion(str, $expression.p, $TK_IDENTIFICADOR.line, localctx.(*AsignacionContext).Get_TK_IDENTIFICADOR().GetColumn()) 
    }
;

/* ------------------------------------------------------------------ASIGNACION DE STRUCTS---------------------------------------------------------------- */
asignar_struct_mutable returns[interfaces.Instruccion instr]
  : TK_IDENTIFICADOR TK_PUNTO nom= TK_IDENTIFICADOR TK_IGUAL expression TK_PYC {
    $instr = asignaciones.Nasigstruct($TK_IDENTIFICADOR.text, $nom.text, $expression.p, $TK_IDENTIFICADOR.line, localctx.(*Asignar_struct_mutableContext).Get_TK_IDENTIFICADOR().GetColumn()) 
  }
;

/* ------------------------------------------------------------------LLAMADA DE FUNCIONES---------------------------------------------------------------------- */
llamada returns[interfaces.Instruccion instr]
  : TK_IDENTIFICADOR TK_PI list_expres TK_PD TK_PYC {
    str:= $TK_IDENTIFICADOR.text[0:len($TK_IDENTIFICADOR.text)]
    $instr = funciones.Nllamada(str, $list_expres.lis_expres, $TK_IDENTIFICADOR.line, localctx.(*LlamadaContext).Get_TK_IDENTIFICADOR().GetColumn()) 
  }
;

/* ---------------------------------------------------------------SENTENCIAS DE TRANSFERENCIA------------------------------------------------------------------ */
/* TRANSFERENCIA */
transferencia returns[interfaces.Instruccion instr]
  : retorno        {$instr = $retorno.instr}
  | romper         {$instr = $romper.instr}
  | continuar      {$instr = $continuar.instr}
;

/* RETURN */
retorno returns[interfaces.Instruccion instr]
  : TK_RETURN TK_PYC {
    $instr = stransferencia.Nretorno(simbolos.METODO, nil)
  }
  | TK_RETURN expression TK_PYC {
    $instr = stransferencia.Nretorno(simbolos.NULL, $expression.p)
  }
  | expression TK_PYC {
    $instr = stransferencia.Nretorno(simbolos.NULL, $expression.p)
  }
;

/* BREAK */
romper returns[interfaces.Instruccion instr]
  : TK_BREAK TK_PYC {
    $instr = stransferencia.Nromper(simbolos.METODO, nil)
  }
  | TK_BREAK expression TK_PYC {
    $instr = stransferencia.Nromper(simbolos.NULL, $expression.p)
  }
;

/* CONTINUE */
continuar returns[interfaces.Instruccion instr]
  : TK_CONTINUE TK_PYC {
    $instr = stransferencia.Ncontinuar(simbolos.METODO)
  }
;

/* LISTA DE BLOQUES INSTRUCCIONES */
l_bloque returns[*arrayList.List lbloque]
  @init {
    $lbloque =  arrayList.New()
  }
  : block = l_bloque bloque {
    $block.lbloque.Add($bloque.instr)
    $lbloque=$block.lbloque
  }
  | bloque { $lbloque.Add($bloque.instr)}
;

/* BLOQUE INSTRUCCIONES */
bloque returns[interfaces.Instruccion instr]
  : imprimir                {$instr = $imprimir.instr}
  | sent_if                 {$instr = $sent_if.instr}
  | sent_match              {$instr = $sent_match.instr}
  | loop                    {$instr = $loop.instr}
  | c_while                 {$instr = $c_while.instr}
  | declaracion             {$instr = $declaracion.instr}
  | asignacion              {$instr = $asignacion.instr}
  | asignar_struct_mutable  {$instr = $asignar_struct_mutable.instr}
  | transferencia           {$instr = $transferencia.instr}
  | llamada                 {$instr = $llamada.instr}
  | fn_vector               {$instr = $fn_vector.instr}
;

/* EXPRESIONES */
expression returns[interfaces.Expresion p]
    : fun_nativas             {$p = $fun_nativas.p}
    | nativas_vector          {$p = $nativas_vector.p}
    | expre_logica            {$p = $expre_logica.p}
    | expre_relacional        {$p = $expre_relacional.p}
    | expre_aritmetica        {$p = $expre_aritmetica.p}
    | sent_if                 {$p = $sent_if.p}
    | sent_match              {$p = $sent_match.p}
    | loop                    {$p = $loop.p}
    | vectores_inicio         {$p = $vectores_inicio.p}
    | arreglos_inicio         {$p = $arreglos_inicio.p}
    | accesso_arreglo         {$p = $accesso_arreglo.p}
    | acceso_vector           {$p = $acceso_vector.p}
    | ingreso_struct          {$p = $ingreso_struct.p}
    | valores                 {$p = $valores.p}
;

/* LISTA DE EXPRESIONES */
list_expres returns[*arrayList.List lis_expres]
  @init {
    $lis_expres =  arrayList.New()
  }
  : list = list_expres TK_COMA expression {
    $list.lis_expres.Add($expression.p)
    $lis_expres=$list.lis_expres
  }
  | expression {$lis_expres.Add($expression.p)}
;

/* INICIANDO VECTORES */
vectores_inicio returns[interfaces.Expresion p]
  : inicio_vect     {$p = $inicio_vect.p}
  | resumen_vect    {$p = $resumen_vect.p}
;

inicio_vect returns[interfaces.Expresion p]
  : TK_VECT TK_CI list_expres TK_CD { 
    $p = iniciandovect.Nvalor_vector($list_expres.lis_expres) 
  }
;

resumen_vect returns[interfaces.Expresion p]
  : TK_VECT TK_CI ex1 = expression TK_PYC ex2 = expression TK_CD { 
    $p = iniciandovect.Nvalvect($ex1.p,$ex2.p) 
  }
;

/* INICIALIZANDO ARREGLOS */
arreglos_inicio returns[interfaces.Expresion p]
  : inicializando_arreglo {$p = $inicializando_arreglo.p}
  | resumen_arre {$p = $resumen_arre.p}
;

inicializando_arreglo returns[interfaces.Expresion p]
  : TK_CI list_expres TK_CD { $p = iniciandoarre.Nvalor_arreglo($list_expres.lis_expres) }
;

resumen_arre returns[interfaces.Expresion p]
  : TK_CI ex1 = expression TK_PYC ex2 = expression TK_CD { 
    $p = iniciandoarre.Nresumenarre($ex1.p,$ex2.p) 
  }
;

/* ACCESO A OBJETOS DEL ARREGLO */
accesso_arreglo returns[interfaces.Expresion p]
  : TK_IDENTIFICADOR lista_acceso {
    $p = accesoarre.Naccessarre($TK_IDENTIFICADOR.text,$lista_acceso.lacceso, $TK_IDENTIFICADOR.line, localctx.(*Accesso_arregloContext).Get_TK_IDENTIFICADOR().GetColumn())
  }
;

lista_acceso returns[*arrayList.List lacceso]
  @init {
    $lacceso =  arrayList.New()
  }
  : lacc = lista_acceso access {
    $lacc.lacceso.Add($access.p)
    $lacceso=$lacc.lacceso
  }
  | access {
    $lacceso.Add($access.p)
  }
;

access returns[interfaces.Expresion p]
  : TK_CI expression TK_CD { $p = $expression.p}
;

/* ACCESO A OBJETOS DEL VECTOR */
acceso_vector returns[interfaces.Expresion p]
  : TK_IDENTIFICADOR TK_MENOR expression TK_MAYOR {
    $p = accesovect.Naccessvect($TK_IDENTIFICADOR.text,$expression.p, $TK_IDENTIFICADOR.line, localctx.(*Acceso_vectorContext).Get_TK_IDENTIFICADOR().GetColumn())
  }
;

ingreso_struct returns[interfaces.Expresion p]
  : TK_IDENTIFICADOR TK_PUNTO noms = TK_IDENTIFICADOR {
    $p = accesostruct.NaccesoStruct($TK_IDENTIFICADOR.text,$noms.text, $TK_IDENTIFICADOR.line, localctx.(*Ingreso_structContext).Get_TK_IDENTIFICADOR().GetColumn())
  }
;

nativas_vector returns[interfaces.Expresion p]
  : TK_IDENTIFICADOR TK_PUNTO TK_LEN {
    $p = nativasvect.Nlenvect($TK_IDENTIFICADOR.text, $TK_IDENTIFICADOR.line, localctx.(*Nativas_vectorContext).Get_TK_IDENTIFICADOR().GetColumn())
  }
  | TK_IDENTIFICADOR TK_PUNTO TK_CAPACITY {
    $p = nativasvect.Ncapvect($TK_IDENTIFICADOR.text, $TK_IDENTIFICADOR.line, localctx.(*Nativas_vectorContext).Get_TK_IDENTIFICADOR().GetColumn())
  }
;

/* FUNCIONES NATIVAS */
fun_nativas returns[interfaces.Expresion p]
  : val = fun_nativas TK_PUNTO TK_ABS {
    $p = nativas.Nabsoluto($val.p,$TK_PUNTO.line,localctx.(*Fun_nativasContext).Get_TK_PUNTO().GetColumn())
  }
  | val = fun_nativas TK_PUNTO TK_SQRT {
    $p = nativas.Nraiz($val.p,$TK_PUNTO.line,localctx.(*Fun_nativasContext).Get_TK_PUNTO().GetColumn())
  }
  | val = fun_nativas TK_PUNTO TK_TOSTRING {
    $p = nativas.Ntostring($val.p,$TK_PUNTO.line,localctx.(*Fun_nativasContext).Get_TK_PUNTO().GetColumn())
  }
  | val = fun_nativas TK_PUNTO TK_TOOWNED {
    $p = nativas.Ntoowned($val.p,$TK_PUNTO.line,localctx.(*Fun_nativasContext).Get_TK_PUNTO().GetColumn())
  }
  | val = fun_nativas TK_PUNTO TK_CLONE {
    $p = nativas.Nclone($val.p,$TK_PUNTO.line,localctx.(*Fun_nativasContext).Get_TK_PUNTO().GetColumn())
  }
  | expre_logica  {$p = $expre_logica.p}
;

/* LOGICA */
expre_logica returns[interfaces.Expresion p]
  : op = TK_NOT opI = expression {
      $p = logica.Nopnot($opI.p, $op.line, localctx.(*Expre_logicaContext).op.GetColumn())
    }
  | opIz = expre_logica op = (TK_AND | TK_OR) opDe = expre_logica {
      if $op.text == "&&"{
        $p = logica.Nopand($opIz.p,$opDe.p, $op.line, localctx.(*Expre_logicaContext).op.GetColumn())
      }else if $op.text == "||"{
        $p = logica.Noor($opIz.p,$opDe.p, $op.line, localctx.(*Expre_logicaContext).op.GetColumn())
      }
    
    }
  | expre_relacional        {$p = $expre_relacional.p}
;

/* RELACIONAL */
expre_relacional returns[interfaces.Expresion p]
  : opIz = expre_relacional op = (TK_MAYORIGUAL| TK_MENORIGUAL| TK_IGUALDAD| TK_DESIGUALDAD| TK_MAYOR| TK_MENOR) opDe = expre_relacional {
      if $op.text == "<"{
        $p = relacional.Nopmenor($opIz.p,$opDe.p, $op.line, localctx.(*Expre_relacionalContext).op.GetColumn())
      }else if $op.text == "<="{
        $p = relacional.Nopmenorigual($opIz.p,$opDe.p, $op.line, localctx.(*Expre_relacionalContext).op.GetColumn())
      }else if $op.text == ">"{
        $p = relacional.Nopmayor($opIz.p,$opDe.p, $op.line, localctx.(*Expre_relacionalContext).op.GetColumn())
      }else if $op.text == ">="{
        $p = relacional.Nopmayorigaul($opIz.p,$opDe.p, $op.line, localctx.(*Expre_relacionalContext).op.GetColumn())
      }else if $op.text == "=="{
        $p = relacional.Nopigualdad($opIz.p,$opDe.p, $op.line, localctx.(*Expre_relacionalContext).op.GetColumn())
      }else if $op.text == "!="{
        $p = relacional.Nopdesigualdad($opIz.p,$opDe.p, $op.line, localctx.(*Expre_relacionalContext).op.GetColumn())
      }
  }
  | expre_aritmetica {$p = $expre_aritmetica.p}
;

/* ARITMETICA */
expre_aritmetica returns[interfaces.Expresion p]
  : opera = TK_RESTA opUn = expression {
      $p = aritmetica.Nopnegativo($opUn.p, $opera.line, localctx.(*Expre_aritmeticaContext).opera.GetColumn())
    }
  | TK_PI expression TK_PD  {$p = $expression.p}
  | tipito = TK_TIPOINT TK_POW TK_PI opIz = expre_aritmetica TK_COMA opDe = expre_aritmetica TK_PD  {
      $p = aritmetica.Npotencia($tipito.text,$opIz.p,$opDe.p, $TK_POW.line, localctx.(*Expre_aritmeticaContext).Get_TK_POW().GetColumn())
    }
  | tipito = TK_TIPOFLOAT TK_POWF TK_PI opIz = expre_aritmetica TK_COMA opDe = expre_aritmetica TK_PD   {
      $p = aritmetica.Npotencia($tipito.text,$opIz.p,$opDe.p, $TK_POWF.line, localctx.(*Expre_aritmeticaContext).Get_TK_POWF().GetColumn())
    } 
  | opIz = expre_aritmetica opera = (TK_MULTI|TK_DIVI|TK_MODULO) opDe = expre_aritmetica {
      if $opera.text == "*"{ 
        $p = aritmetica.Nopmultiplicacion($opIz.p,$opDe.p, $opera.line, localctx.(*Expre_aritmeticaContext).opera.GetColumn())
      }else if $opera.text == "/"{
        $p = aritmetica.Nopdivision($opIz.p,$opDe.p, $opera.line, localctx.(*Expre_aritmeticaContext).opera.GetColumn())
      }else if $opera.text == "%"{
        $p = aritmetica.Nopmodulo($opIz.p,$opDe.p, $opera.line, localctx.(*Expre_aritmeticaContext).opera.GetColumn())
      }
    }
  | opIz = expre_aritmetica opera = (TK_SUMA|TK_RESTA) opDe = expre_aritmetica  {        
      if $opera.text=="+"{
        $p = aritmetica.Nopsuma($opIz.p,$opDe.p, $opera.line, localctx.(*Expre_aritmeticaContext).opera.GetColumn())
      }else{ 
        $p = aritmetica.Nopresta($opIz.p,$opDe.p, $opera.line, localctx.(*Expre_aritmeticaContext).opera.GetColumn())
      }
    }
  | valores  {$p = $valores.p}
;

/* VALORES DONDE ESTA LOS PRIMITIVOS Y SUS CASTEOS */
valores returns[interfaces.Expresion p]
  : datos         {$p = $datos.p}
  | casteo_datos  {$p = $casteo_datos.p}
;

/* CASTEO EXPLICITO */
casteo_datos returns[interfaces.Expresion p]
  : TK_ENTERO TK_AS tipos {
      num,err := strconv.Atoi($TK_ENTERO.text)
      if err!= nil{
        fmt.Println(err)
      }
      val := primitivo.Nuevo_Dato_Primitivo (num,simbolos.INTEGER)
      $p = casteos.Ncasteentero(val,$tipos.tip)
    }
  | TK_FLOAT TK_AS tipos {
      num,err := strconv.ParseFloat($TK_FLOAT.text,64)
      if err!=nil{
        fmt.Println(err)
      }
      val := primitivo.Nuevo_Dato_Primitivo (num,simbolos.FLOAT)
      $p = casteos.Ncastefloat(val,$tipos.tip)
    }
  | TK_CADENA TK_AS tipos { 
      str:= $TK_CADENA.text[1:len($TK_CADENA.text)-1]
      val := primitivo.Nuevo_Dato_Primitivo(str,simbolos.YTEXTO)
      $p = casteos.Ncastetexto(val,$tipos.tip)
    }
  | TK_CARACTER TK_AS tipos { 
      str:= $TK_CARACTER.text[1:len($TK_CARACTER.text)-1]
      r := []rune(str)
      val := primitivo.Nuevo_Dato_Primitivo(r[0],simbolos.CHAR)
      $p = casteos.Ncastechar(val,$tipos.tip)
    }
  | TK_TRUE TK_AS tipos {
      val := primitivo.Nuevo_Dato_Primitivo(true,simbolos.BOOLEAN)
      $p = casteos.Ncastebool(val,$tipos.tip)
    }
  | TK_FALSE TK_AS tipos {
      val := primitivo.Nuevo_Dato_Primitivo(false,simbolos.BOOLEAN)
      $p = casteos.Ncastebool(val,$tipos.tip)
    }
;

/* DATO PRIMITIVO */
datos returns[interfaces.Expresion p]
  : TK_ENTERO {
      num,err := strconv.Atoi($TK_ENTERO.text)
      if err!= nil{
        fmt.Println(err)
      }
      $p = primitivo.Nuevo_Dato_Primitivo (num,simbolos.INTEGER)
    }
  | TK_FLOAT {
      num,err := strconv.ParseFloat($TK_FLOAT.text,64)
      if err!=nil{
        fmt.Println(err)
      }
      $p = primitivo.Nuevo_Dato_Primitivo (num,simbolos.FLOAT)
    }
  | TK_CADENA { 
      str:= $TK_CADENA.text[1:len($TK_CADENA.text)-1]
      $p = primitivo.Nuevo_Dato_Primitivo(str,simbolos.YTEXTO)
    }
  | TK_IDENTIFICADOR {
      str:= $TK_IDENTIFICADOR.text[0:len($TK_IDENTIFICADOR.text)]
      $p = variables.NVariable(str, $TK_IDENTIFICADOR.line, localctx.(*DatosContext).Get_TK_IDENTIFICADOR().GetColumn())
    }
  | TK_CARACTER { 
      str:= $TK_CARACTER.text[1:len($TK_CARACTER.text)-1]
      r := []rune(str)
      $p = primitivo.Nuevo_Dato_Primitivo(r[0],simbolos.CHAR)
    }
  | TK_TRUE {
      $p = primitivo.Nuevo_Dato_Primitivo(true,simbolos.BOOLEAN)
    }
  | TK_FALSE {
      $p = primitivo.Nuevo_Dato_Primitivo(false,simbolos.BOOLEAN)
    }
;