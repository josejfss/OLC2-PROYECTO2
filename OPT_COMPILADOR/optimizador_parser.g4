parser grammar optimizador_parser;

options {
    tokenVocab = optimizador_lexer;
}

@header {
    import "OLC2-PROYECTO2/OPT_COMPILADOR/INTER"
    import "OLC2-PROYECTO2/OPT_COMPILADOR/BLOQUE"
    /* EXPRESIONES */
    /* OPERACION ARITMETICA */
    import "OLC2-PROYECTO2/OPT_COMPILADOR/AST/EXPRESION/OPERACION/ARITMETICA/SUMA"
    import "OLC2-PROYECTO2/OPT_COMPILADOR/AST/EXPRESION/OPERACION/ARITMETICA/RESTA"
    import "OLC2-PROYECTO2/OPT_COMPILADOR/AST/EXPRESION/OPERACION/ARITMETICA/MULTIPLICACION"
    import "OLC2-PROYECTO2/OPT_COMPILADOR/AST/EXPRESION/OPERACION/ARITMETICA/DIVISION"
    import "OLC2-PROYECTO2/OPT_COMPILADOR/AST/EXPRESION/OPERACION/ARITMETICA/MODULO"
    import "OLC2-PROYECTO2/OPT_COMPILADOR/AST/EXPRESION/OPERACION/ARITMETICA/NEGATIVO"
    /* OPERACION RELACIONAL */
    import "OLC2-PROYECTO2/OPT_COMPILADOR/AST/EXPRESION/OPERACION/RELACIONAL/DESIGUALDAD"
    import "OLC2-PROYECTO2/OPT_COMPILADOR/AST/EXPRESION/OPERACION/RELACIONAL/IGUALDAD"
    import "OLC2-PROYECTO2/OPT_COMPILADOR/AST/EXPRESION/OPERACION/RELACIONAL/MAYOR"
    import "OLC2-PROYECTO2/OPT_COMPILADOR/AST/EXPRESION/OPERACION/RELACIONAL/MAYORIGUAL"
    import "OLC2-PROYECTO2/OPT_COMPILADOR/AST/EXPRESION/OPERACION/RELACIONAL/MENOR"
    import "OLC2-PROYECTO2/OPT_COMPILADOR/AST/EXPRESION/OPERACION/RELACIONAL/MENORIGUAL"
    /* PRIMITIVOS /home/iovana/go/src/  */
    import "OLC2-PROYECTO2/OPT_COMPILADOR/AST/EXPRESION/PRIMITIIVOS"
    /* ASIGNACIONES */
    import "OLC2-PROYECTO2/OPT_COMPILADOR/AST/EXPRESION/ASIGNACIONES"
    /* FUNCIONES */
    /* ENCABEZADO */
    import "OLC2-PROYECTO2/OPT_COMPILADOR/AST/INSTRUCCION/ENCABEZADO"
    import "OLC2-PROYECTO2/OPT_COMPILADOR/AST/INSTRUCCION/FUNCION"
    import "OLC2-PROYECTO2/OPT_COMPILADOR/AST/INSTRUCCION/DECLARAR"
    import "OLC2-PROYECTO2/OPT_COMPILADOR/AST/INSTRUCCION/SIF"
    import "OLC2-PROYECTO2/OPT_COMPILADOR/AST/INSTRUCCION/IMPRIMIR"
    import "OLC2-PROYECTO2/OPT_COMPILADOR/AST/INSTRUCCION/ETIQUETAS"
    import "OLC2-PROYECTO2/OPT_COMPILADOR/AST/INSTRUCCION/LLAMADA"
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

instrs returns[inter.Instruccion instr]
  : encabezado            {$instr = $encabezado.instr}
;

encabezado returns[inter.Instruccion instr]
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

/* -------------------------------------------------------------------FUNCIONES--------------------------------------------------------------------------------- */
funcas returns[inter.Instruccion instr]
  : TK_VOID TK_IDENTIFICADOR TK_PI TK_PD TK_LI l_bloque TK_RETURN TK_LD {
    $instr = funcion.Nfunca($TK_IDENTIFICADOR.text, $l_bloque.lbloque)
  }
  | TK_INT TK_MAIN TK_PI TK_PD TK_LI TK_ASIGPSTACK TK_ASIGHEAP l_bloque TK_RETMAIN TK_LD {
    $instr = funcion.Nfuncamain($TK_MAIN.text, $l_bloque.lbloque)
  } 
;

/* -------------------------------------------------------------------CICLOS Y SENTENCIAS---------------------------------------------------------------------- */
/* CONDICIONAL IF */ 
sent_if returns [inter.Instruccion instr]
  : TK_IF TK_PI expression TK_PD TK_GOTO TK_ETIQUETA TK_PYC {
    $instr = sif.Nsenteif($expression.p,$TK_ETIQUETA.text, $TK_IF.line)
  }
;

/* ------------------------------------------------------------IMPRESION EN CONSOLA------------------------------------------------------------------------ */
/* IMPRESION */
imprimir returns[inter.Instruccion instr]
  : TK_PRINTF TK_PI TK_CADENA TK_COMA casteo expression TK_PD TK_PYC {
    $instr = imprimir.Nimprimir($TK_CADENA.text, $casteo.cast, $expression.p,$TK_PRINTF.line)
  }
;

casteo returns[string cast]
  : TK_CASTCHAR   {$cast = $TK_CASTCHAR.text}
  | TK_CASTINT    {$cast = $TK_CASTINT.text}
  | {$cast = ""}
;

/* ----------------------------------------------------------DECLARAR VARIABLES, VECTORES Y ARREGLOS--------------------------------------------------------- */
declaracion returns[inter.Instruccion instr]
  : TK_TEMPORAL TK_IGUAL expression TK_PYC {
    $instr = declarar.Ndeclaracion($TK_TEMPORAL.text, $expression.p,$TK_TEMPORAL.line)
  }
  | nom=(TK_STACK|TK_HEAP) TK_CI TK_CASTINT pos=expression TK_CD TK_IGUAL exp=expression TK_PYC {
    $instr = declarar.Ndeclapunteros($nom.text, $pos.p, $exp.p, $nom.line)
  }
  | nom=(TK_PUNSTACK|TK_PUNHEAP) TK_IGUAL expression TK_PYC {
    $instr = declarar.Ndeclapstack($nom.text, $expression.p,$nom.line)
  }
;

etiquetas returns[inter.Instruccion instr]
  : TK_ETIQUETA TK_DP {
    $instr = etiquetas.Netiqueta($TK_ETIQUETA.text)
  }
  | TK_GOTO TK_ETIQUETA TK_PYC{
    $instr = etiquetas.Ngoto($TK_ETIQUETA.text)
  }
;

llamada returns[inter.Instruccion instr]
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
bloque returns[inter.Instruccion instr]
  : imprimir                {$instr = $imprimir.instr}
  | sent_if                 {$instr = $sent_if.instr}
  | declaracion             {$instr = $declaracion.instr}
  | etiquetas               {$instr = $etiquetas.instr}
  | llamada                 {$instr = $llamada.instr}
;

/* EXPRESIONES */
expression returns[inter.Expresion p]
    : expre_relacional        {$p = $expre_relacional.p}
    | expre_aritmetica        {$p = $expre_aritmetica.p}
    | asignaciones            {$p = $asignaciones.p}
;

asignaciones returns[inter.Expresion p]
  : TK_HEAP TK_CI TK_CASTINT expression TK_CD {
    $p = asignaciones.Nasiheap($expression.p)
  }
  | TK_STACK TK_CI TK_CASTINT expression TK_CD  {
    $p = asignaciones.Nasistack($expression.p)
  }
;

/* RELACIONAL */
expre_relacional returns[inter.Expresion p]
  : opIz = expre_relacional op = (TK_MAYORIGUAL| TK_MENORIGUAL| TK_IGUALDAD| TK_DESIGUALDAD| TK_MAYOR| TK_MENOR) opDe = expre_relacional {
      if $op.text == "<"{
        $p = menor.Noperacionmenor($opIz.p,$opDe.p)
      }else if $op.text == "<="{
        $p = menorigual.Noperacionmenorigual($opIz.p,$opDe.p)
      }else if $op.text == ">"{
        $p = mayor.Noperacionmayor($opIz.p,$opDe.p)
      }else if $op.text == ">="{
        $p = mayorigual.Noperacionmayorigual($opIz.p,$opDe.p)
      }else if $op.text == "=="{
        $p = igualdad.Noperacionigualdad($opIz.p,$opDe.p)
      }else if $op.text == "!="{
        $p = desigualdad.Noperaciondesigualdad($opIz.p,$opDe.p)
      }
  }
  | expre_aritmetica {$p = $expre_aritmetica.p}
;

/* ARITMETICA */
expre_aritmetica returns[inter.Expresion p]
  : opera = TK_RESTA opUn = expre_aritmetica {
      $p = negativo.Nnegativo($opUn.p)
    }
  | opIz = expre_aritmetica opera = (TK_MULTI|TK_DIVI|TK_MODULO) opDe = expre_aritmetica {
      if $opera.text == "*"{ 
        $p = multiplicacion.Noperacionmultiplicacion($opIz.p,$opDe.p)
      }else if $opera.text == "/"{
        $p = division.Noperaciondivision($opIz.p,$opDe.p)
      }else if $opera.text == "%"{
        $p = modulo.Noperacionmodulo($opIz.p,$opDe.p)
      }
    }
  | opIz = expre_aritmetica opera = (TK_SUMA|TK_RESTA) opDe = expre_aritmetica  {        
      if $opera.text=="+"{
        $p = suma.Noperacionsuma($opIz.p,$opDe.p)
      }else{ 
        $p = resta.Noperacionresta($opIz.p,$opDe.p)
      }
    }
  | datos  {$p = $datos.p}
;


/* DATO PRIMITIVO */
datos returns[inter.Expresion p]
  : TK_ENTERO {
      $p = primitiivos.NuevoPrimitivo ($TK_ENTERO.text,bloque.INTEGER)
    }
  | TK_FLOAT {
      $p = primitiivos.NuevoPrimitivo ($TK_FLOAT.text,bloque.FLOAT)
    }
  | TK_IDENTIFICADOR {
      $p = primitiivos.NuevoPrimitivo($TK_IDENTIFICADOR.text,bloque.NULL)
    }
  | TK_CADENA { 
      $p = primitiivos.NuevoPrimitivo($TK_CADENA.text,bloque.NULL)
    }
  | TK_TEMPORAL {
      $p = primitiivos.NuevoPrimitivo($TK_TEMPORAL.text,bloque.TEMPORAL)
    }
  | TK_PUNSTACK {
      $p = primitiivos.NuevoPrimitivo($TK_PUNSTACK.text,bloque.NULL)
    }
  | TK_PUNHEAP {
      $p = primitiivos.NuevoPrimitivo($TK_PUNHEAP.text,bloque.NULL)
    }
;