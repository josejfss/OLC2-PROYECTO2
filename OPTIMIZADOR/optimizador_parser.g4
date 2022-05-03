parser grammar optimizador_parser;

options {
    tokenVocab = optimizador_lexer;
}

@header {
    import "OLC2-PROYECTO2/OPTIMIZADOR/INTERFAZ"
    import "OLC2-PROYECTO2/OPTIMIZADOR/OBJETO"
    import "OLC2-PROYECTO2/OPTIMIZADOR/AST/EXPRESION/OPERACION/ARITMETICA"
    import "OLC2-PROYECTO2/OPTIMIZADOR/AST/EXPRESION/OPERACION/RELACIONAL"
    import "OLC2-PROYECTO2/OPTIMIZADOR/AST/EXPRESION/PRIMITIVOS"
    import "OLC2-PROYECTO2/OPTIMIZADOR/AST/EXPRESION/ASIGNACIONES"
    import "OLC2-PROYECTO2/OPTIMIZADOR/AST/INSTRUCCION/ENCABEZADO"
    import "OLC2-PROYECTO2/OPTIMIZADOR/AST/INSTRUCCION/FUNCION"
    import "OLC2-PROYECTO2/OPTIMIZADOR/AST/INSTRUCCION/LLAMADA"
    import "OLC2-PROYECTO2/OPTIMIZADOR/AST/INSTRUCCION/DECLARAR"
    import "OLC2-PROYECTO2/OPTIMIZADOR/AST/INSTRUCCION/SIF"
    import "OLC2-PROYECTO2/OPTIMIZADOR/AST/INSTRUCCION/IMPRIMIR"
    import "OLC2-PROYECTO2/OPTIMIZADOR/AST/INSTRUCCION/ETIQUETAS"
    import "OLC2-PROYECTO2/OPTIMIZADOR/AST/INSTRUCCION/RETORNO"
    import arrayList "github.com/colegno/arraylist"
}

start returns[*arrayList.List lista]
  : instrucss {$lista = $instrucss.lis}
;

instrucss returns[*arrayList.List lis]
  @init {
      $lis =  arrayList.New()
  }: 
  e +=instrs* {
    listInt := localctx.(*InstrucssContext).GetE()
    for _, e := range listInt {
      $lis.Add(e.GetInstr())
    }
  }
;

instrs returns[interfaz.Instruccion instr]
  : encabezado            {$instr = $encabezado.instr}
;

encabezado returns[interfaz.Instruccion instr]
  : TK_STDIOH TK_DHEAP TK_DSTACK TK_PSTAKC TK_PHEAP TK_DOUBLE l_temporales TK_PYC l_funcas {
    $instr = encabezado.Nenca($TK_STDIOH.text,$TK_DHEAP.text,$TK_DSTACK.text,$TK_PSTAKC.text,$TK_PHEAP.text,$l_temporales.ltemps,$l_funcas.lfuncas)
  }
;

l_temporales returns[*arrayList.List ltemps]
  @init {
      $ltemps =  arrayList.New()
  }
  : temps = l_temporales TK_COMA TK_TEMPORAL {
    $temps.ltemps.Add($TK_TEMPORAL.text)
    $ltemps=$temps.ltemps
  }
  | TK_TEMPORAL { $ltemps.Add($TK_TEMPORAL.text)}
;

l_funcas returns[*arrayList.List lfuncas]
  @init {
      $lfuncas =  arrayList.New()
  }
  : funcs=l_funcas funcas {
    $funcs.lfuncas.Add($funcas.instr)
    $lfuncas=$funcs.lfuncas
  }
  | funcas { $lfuncas.Add($funcas.instr)}
;

/*

  | funcas instrs         {$instr = $funcas.instr}
  | funcas_main instrs    {$instr = $funcas_main.instr} */

/* -------------------------------------------------------------------FUNCIONES--------------------------------------------------------------------------------- */
funcas returns[interfaz.Instruccion instr]
  : TK_VOID TK_IDENTIFICADOR TK_PI TK_PD TK_LI l_bloque TK_LD {
    $instr = funcion.Nfunca($TK_IDENTIFICADOR.text, $l_bloque.lbloque)
  }
  | TK_INT TK_MAIN TK_PI TK_PD TK_LI l_bloque TK_LD {
    $instr = funcion.Nfunca($TK_MAIN.text, $l_bloque.lbloque)
  } 
;

// funcas_main returns[interfaz.Instruccion instr]
//   : 
// ;

/* -------------------------------------------------------------------CICLOS Y SENTENCIAS---------------------------------------------------------------------- */
/* CONDICIONAL IF */ 
sent_if returns [interfaz.Instruccion instr]
  : TK_IF TK_PI expression TK_PD TK_GOTO TK_ETIQUETA TK_PYC {
    $instr = sif.Nsenteif($expression.p,$TK_ETIQUETA.text)
  }
;

/* ------------------------------------------------------------IMPRESION EN CONSOLA------------------------------------------------------------------------ */
/* IMPRESION */
imprimir returns[interfaz.Instruccion instr]
  : TK_PRINTF TK_PI TK_CADENA TK_COMA casteo expression TK_PD TK_PYC {
    $instr = imprimir.Nimprimir($TK_CADENA.text, $casteo.cast, $expression.p)
  }
;

casteo returns[string cast]
  : TK_CASTCHAR   {$cast = $TK_CASTCHAR.text}
  | TK_CASTINT    {$cast = $TK_CASTINT.text}
  | {$cast = ""}
;

/* ----------------------------------------------------------DECLARAR VARIABLES, VECTORES Y ARREGLOS--------------------------------------------------------- */
declaracion returns[interfaz.Instruccion instr]
  : TK_TEMPORAL TK_IGUAL expression TK_PYC {
    $instr = declarar.Ndeclaracion($TK_TEMPORAL.text, $expression.p)
  }
  | TK_STACK TK_CI TK_CASTINT pos=expression TK_CD TK_IGUAL exp=expression TK_PYC {
    $instr = declarar.Ndeclastack($pos.p, $exp.p)
  }
  | TK_HEAP TK_CI TK_CASTINT pos=expression TK_CD TK_IGUAL exp=expression TK_PYC {
    $instr = declarar.Ndeclaheap($pos.p, $exp.p)
  }
  | TK_PUNSTACK TK_IGUAL expression TK_PYC {
    $instr = declarar.Ndeclapstack($TK_PUNSTACK.text, $expression.p)
  }
  | TK_PUNHEAP TK_IGUAL expression TK_PYC {
    $instr = declarar.Ndeclapheap($TK_PUNHEAP.text, $expression.p)
  }
;

etiquetas returns[interfaz.Instruccion instr]
  : TK_ETIQUETA TK_DP {
    $instr = etiquetas.Netiqueta($TK_ETIQUETA.text)
  }
  | TK_GOTO TK_ETIQUETA TK_PYC{
    $instr = etiquetas.Ngoto($TK_ETIQUETA.text)
  }
;

retorno returns[interfaz.Instruccion instr]
  : TK_RETMAIN  {$instr= retorno.Nreturns($TK_RETMAIN.text)}
  | TK_RETURN   {$instr= retorno.Nreturns($TK_RETURN.text)}
;

llamada returns[interfaz.Instruccion instr]
  : TK_IDENTIFICADOR TK_PI TK_PD TK_PYC {
    $instr= llamada.Nllama($TK_IDENTIFICADOR.text)
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
bloque returns[interfaz.Instruccion instr]
  : imprimir                {$instr = $imprimir.instr}
  | sent_if                 {$instr = $sent_if.instr}
  | declaracion             {$instr = $declaracion.instr}
  | etiquetas               {$instr = $etiquetas.instr}
  | retorno                 {$instr = $retorno.instr}
  | llamada                 {$instr = $llamada.instr}
;

/* EXPRESIONES */
expression returns[interfaz.Expresion p]
    : expre_relacional        {$p = $expre_relacional.p}
    | expre_aritmetica        {$p = $expre_aritmetica.p}
    | asignaciones            {$p = $asignaciones.p}
;

asignaciones returns[interfaz.Expresion p]
  : TK_HEAP TK_CI TK_CASTINT expression TK_CD {
    $p = asignaciones.Nasiheap($expression.p)
  }
  | TK_STACK TK_CI TK_CASTINT expression TK_CD  {
    $p = asignaciones.Nasistack($expression.p)
  }
;

/* RELACIONAL */
expre_relacional returns[interfaz.Expresion p]
  : opIz = expre_relacional op = (TK_MAYORIGUAL| TK_MENORIGUAL| TK_IGUALDAD| TK_DESIGUALDAD| TK_MAYOR| TK_MENOR) opDe = expre_relacional {
      if $op.text == "<"{
        $p = relacional.Noperacionmenor($opIz.p,$opDe.p)
      }else if $op.text == "<="{
        $p = relacional.Noperacionmenorigual($opIz.p,$opDe.p)
      }else if $op.text == ">"{
        $p = relacional.Noperacionmayor($opIz.p,$opDe.p)
      }else if $op.text == ">="{
        $p = relacional.Noperacionmayorigual($opIz.p,$opDe.p)
      }else if $op.text == "=="{
        $p = relacional.Noperacionigualdad($opIz.p,$opDe.p)
      }else if $op.text == "!="{
        $p = relacional.Noperaciondesigualdad($opIz.p,$opDe.p)
      }
  }
  | expre_aritmetica {$p = $expre_aritmetica.p}
;

/* ARITMETICA */
expre_aritmetica returns[interfaz.Expresion p]
  : opera = TK_RESTA opUn = expre_aritmetica {
      $p = aritmetica.Nnegativo($opUn.p)
    }
  | opIz = expre_aritmetica opera = (TK_MULTI|TK_DIVI|TK_MODULO) opDe = expre_aritmetica {
      if $opera.text == "*"{ 
        $p = aritmetica.Noperacionmultiplicacion($opIz.p,$opDe.p)
      }else if $opera.text == "/"{
        $p = aritmetica.Noperaciondivision($opIz.p,$opDe.p)
      }else if $opera.text == "%"{
        $p = aritmetica.Noperacionmodulo($opIz.p,$opDe.p)
      }
    }
  | opIz = expre_aritmetica opera = (TK_SUMA|TK_RESTA) opDe = expre_aritmetica  {        
      if $opera.text=="+"{
        $p = aritmetica.Noperacionsuma($opIz.p,$opDe.p)
      }else{ 
        $p = aritmetica.Noperacionresta($opIz.p,$opDe.p)
      }
    }
  | datos  {$p = $datos.p}
;


/* DATO PRIMITIVO */
datos returns[interfaz.Expresion p]
  : TK_ENTERO {
      $p = primitivos.NuevoPrimitivo ($TK_ENTERO.text,objeto.INTEGER)
    }
  | TK_FLOAT {
      $p = primitivos.NuevoPrimitivo ($TK_FLOAT.text,objeto.FLOAT)
    }
  | TK_CADENA { 
      $p = primitivos.NuevoPrimitivo($TK_CADENA.text,objeto.NULL)
    }
  | TK_IDENTIFICADOR {
      $p = primitivos.NuevoPrimitivo($TK_IDENTIFICADOR.text,objeto.NULL)
    }
  | TK_TEMPORAL {
      $p = primitivos.NuevoPrimitivo($TK_TEMPORAL.text,objeto.TEMPORAL)
    }
  | TK_PUNSTACK {
      $p = primitivos.NuevoPrimitivo($TK_PUNSTACK.text,objeto.NULL)
    }
  | TK_PUNHEAP {
      $p = primitivos.NuevoPrimitivo($TK_PUNHEAP.text,objeto.NULL)
    }
;