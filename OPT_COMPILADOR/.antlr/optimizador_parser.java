// Generated from /home/iovana/go/src/OLC2-PROYECTO2/OPT_COMPILADOR/optimizador_parser.g4 by ANTLR 4.8

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

import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class optimizador_parser extends Parser {
	static { RuntimeMetaData.checkVersion("4.8", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		TK_STDIOH=1, TK_DHEAP=2, TK_DSTACK=3, TK_PSTAKC=4, TK_PHEAP=5, TK_VOID=6, 
		TK_GOTO=7, TK_RETURN=8, TK_HEAP=9, TK_STACK=10, TK_CASTINT=11, TK_CASTCHAR=12, 
		TK_PRINTF=13, TK_MAIN=14, TK_RETMAIN=15, TK_PUNSTACK=16, TK_PUNHEAP=17, 
		TK_IF=18, TK_DOUBLE=19, TK_INT=20, TK_ASIGPSTACK=21, TK_ASIGHEAP=22, TK_FLOAT=23, 
		TK_ENTERO=24, TK_CADENA=25, TK_TEMPORAL=26, TK_ETIQUETA=27, TK_IDENTIFICADOR=28, 
		TK_SUMA=29, TK_RESTA=30, TK_MULTI=31, TK_DIVI=32, TK_MODULO=33, TK_MENORIGUAL=34, 
		TK_MAYORIGUAL=35, TK_IGUALDAD=36, TK_DESIGUALDAD=37, TK_MENOR=38, TK_MAYOR=39, 
		TK_LI=40, TK_LD=41, TK_CI=42, TK_CD=43, TK_PI=44, TK_PD=45, TK_PYC=46, 
		TK_DP=47, TK_COMA=48, TK_IGUAL=49, COMENTARIO_MUL=50, COMENTARIO_LIN=51, 
		WHITESPACE=52;
	public static final int
		RULE_start = 0, RULE_instrucss = 1, RULE_instrs = 2, RULE_encabezado = 3, 
		RULE_l_temporales = 4, RULE_l_funcas = 5, RULE_funcas = 6, RULE_sent_if = 7, 
		RULE_imprimir = 8, RULE_casteo = 9, RULE_declaracion = 10, RULE_etiquetas = 11, 
		RULE_llamada = 12, RULE_l_bloque = 13, RULE_bloque = 14, RULE_expression = 15, 
		RULE_asignaciones = 16, RULE_expre_relacional = 17, RULE_expre_aritmetica = 18, 
		RULE_datos = 19;
	private static String[] makeRuleNames() {
		return new String[] {
			"start", "instrucss", "instrs", "encabezado", "l_temporales", "l_funcas", 
			"funcas", "sent_if", "imprimir", "casteo", "declaracion", "etiquetas", 
			"llamada", "l_bloque", "bloque", "expression", "asignaciones", "expre_relacional", 
			"expre_aritmetica", "datos"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			"'H'", "'#include <stdio.h>'", "'double HEAP[100000];'", "'double STACK[100000];'", 
			"'double SP;'", "'double HP;'", "'void'", "'goto'", "'return;'", "'HEAP'", 
			"'STACK'", "'(int)'", "'(char)'", "'printf'", "'main'", "'return 0;'", 
			"'SP'", "'HP'", "'if'", "'double'", "'int'", null, null, null, null, 
			null, null, null, null, "'+'", "'-'", "'*'", "'/'", "'%'", "'<='", "'>='", 
			"'=='", "'!='", "'<'", "'>'", "'{'", "'}'", "'['", "']'", "'('", "')'", 
			"';'", "':'", "','", "'='"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			"'H'", "TK_STDIOH", "TK_DHEAP", "TK_DSTACK", "TK_PSTAKC", "TK_PHEAP", 
			"TK_VOID", "TK_GOTO", "TK_RETURN", "TK_HEAP", "TK_STACK", "TK_CASTINT", 
			"TK_CASTCHAR", "TK_PRINTF", "TK_MAIN", "TK_RETMAIN", "TK_PUNSTACK", "TK_PUNHEAP", 
			"TK_IF", "TK_DOUBLE", "TK_INT", "TK_ASIGPSTACK", "TK_ASIGHEAP", "TK_FLOAT", 
			"TK_ENTERO", "TK_CADENA", "TK_TEMPORAL", "TK_ETIQUETA", "TK_IDENTIFICADOR", 
			"TK_SUMA", "TK_RESTA", "TK_MULTI", "TK_DIVI", "TK_MODULO", "TK_MENORIGUAL", 
			"TK_MAYORIGUAL", "TK_IGUALDAD", "TK_DESIGUALDAD", "TK_MENOR", "TK_MAYOR", 
			"TK_LI", "TK_LD", "TK_CI", "TK_CD", "TK_PI", "TK_PD", "TK_PYC", "TK_DP", 
			"TK_COMA", "TK_IGUAL", "COMENTARIO_MUL", "COMENTARIO_LIN", "WHITESPACE"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}

	@Override
	public String getGrammarFileName() { return "optimizador_parser.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public optimizador_parser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	public static class StartContext extends ParserRuleContext {
		public *arrayList.List lista;
		public InstrucssContext instrucss;
		public InstrucssContext instrucss() {
			return getRuleContext(InstrucssContext.class,0);
		}
		public StartContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_start; }
	}

	public final StartContext start() throws RecognitionException {
		StartContext _localctx = new StartContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_start);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(40);
			((StartContext)_localctx).instrucss = instrucss();
			_localctx.lista = ((StartContext)_localctx).instrucss.lis
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class InstrucssContext extends ParserRuleContext {
		public *arrayList.List lis;
		public InstrsContext instrs;
		public List<InstrsContext> e = new ArrayList<InstrsContext>();
		public List<InstrsContext> instrs() {
			return getRuleContexts(InstrsContext.class);
		}
		public InstrsContext instrs(int i) {
			return getRuleContext(InstrsContext.class,i);
		}
		public InstrucssContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_instrucss; }
	}

	public final InstrucssContext instrucss() throws RecognitionException {
		InstrucssContext _localctx = new InstrucssContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_instrucss);

		      _localctx.lis =  arrayList.New()
		  
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(46);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==TK_STDIOH) {
				{
				{
				setState(43);
				((InstrucssContext)_localctx).instrs = instrs();
				((InstrucssContext)_localctx).e.add(((InstrucssContext)_localctx).instrs);
				}
				}
				setState(48);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}

			    listInt := localctx.(*InstrucssContext).GetE()
			    for _, e := range listInt {
			      _localctx.lis.Add(e.GetInstr())
			    }
			  
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class InstrsContext extends ParserRuleContext {
		public inter.Instruccion instr;
		public EncabezadoContext encabezado;
		public EncabezadoContext encabezado() {
			return getRuleContext(EncabezadoContext.class,0);
		}
		public InstrsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_instrs; }
	}

	public final InstrsContext instrs() throws RecognitionException {
		InstrsContext _localctx = new InstrsContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_instrs);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(51);
			((InstrsContext)_localctx).encabezado = encabezado();
			_localctx.instr = ((InstrsContext)_localctx).encabezado.instr
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class EncabezadoContext extends ParserRuleContext {
		public inter.Instruccion instr;
		public Token TK_STDIOH;
		public Token TK_DHEAP;
		public Token TK_DSTACK;
		public Token TK_PSTAKC;
		public Token TK_PHEAP;
		public L_temporalesContext l_temporales;
		public L_funcasContext l_funcas;
		public TerminalNode TK_STDIOH() { return getToken(optimizador_parser.TK_STDIOH, 0); }
		public TerminalNode TK_DHEAP() { return getToken(optimizador_parser.TK_DHEAP, 0); }
		public TerminalNode TK_DSTACK() { return getToken(optimizador_parser.TK_DSTACK, 0); }
		public TerminalNode TK_PSTAKC() { return getToken(optimizador_parser.TK_PSTAKC, 0); }
		public TerminalNode TK_PHEAP() { return getToken(optimizador_parser.TK_PHEAP, 0); }
		public TerminalNode TK_DOUBLE() { return getToken(optimizador_parser.TK_DOUBLE, 0); }
		public L_temporalesContext l_temporales() {
			return getRuleContext(L_temporalesContext.class,0);
		}
		public TerminalNode TK_PYC() { return getToken(optimizador_parser.TK_PYC, 0); }
		public L_funcasContext l_funcas() {
			return getRuleContext(L_funcasContext.class,0);
		}
		public EncabezadoContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_encabezado; }
	}

	public final EncabezadoContext encabezado() throws RecognitionException {
		EncabezadoContext _localctx = new EncabezadoContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_encabezado);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(54);
			((EncabezadoContext)_localctx).TK_STDIOH = match(TK_STDIOH);
			setState(55);
			((EncabezadoContext)_localctx).TK_DHEAP = match(TK_DHEAP);
			setState(56);
			((EncabezadoContext)_localctx).TK_DSTACK = match(TK_DSTACK);
			setState(57);
			((EncabezadoContext)_localctx).TK_PSTAKC = match(TK_PSTAKC);
			setState(58);
			((EncabezadoContext)_localctx).TK_PHEAP = match(TK_PHEAP);
			setState(59);
			match(TK_DOUBLE);
			setState(60);
			((EncabezadoContext)_localctx).l_temporales = l_temporales(0);
			setState(61);
			match(TK_PYC);
			setState(62);
			((EncabezadoContext)_localctx).l_funcas = l_funcas(0);

			    _localctx.instr = encabezado.Nenca((((EncabezadoContext)_localctx).TK_STDIOH!=null?((EncabezadoContext)_localctx).TK_STDIOH.getText():null),(((EncabezadoContext)_localctx).TK_DHEAP!=null?((EncabezadoContext)_localctx).TK_DHEAP.getText():null),(((EncabezadoContext)_localctx).TK_DSTACK!=null?((EncabezadoContext)_localctx).TK_DSTACK.getText():null),(((EncabezadoContext)_localctx).TK_PSTAKC!=null?((EncabezadoContext)_localctx).TK_PSTAKC.getText():null),(((EncabezadoContext)_localctx).TK_PHEAP!=null?((EncabezadoContext)_localctx).TK_PHEAP.getText():null),((EncabezadoContext)_localctx).l_temporales.ltemps,((EncabezadoContext)_localctx).l_funcas.lfuncas)
			  
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class L_temporalesContext extends ParserRuleContext {
		public *arrayList.List ltemps;
		public L_temporalesContext temps;
		public Token TK_TEMPORAL;
		public TerminalNode TK_TEMPORAL() { return getToken(optimizador_parser.TK_TEMPORAL, 0); }
		public TerminalNode TK_COMA() { return getToken(optimizador_parser.TK_COMA, 0); }
		public L_temporalesContext l_temporales() {
			return getRuleContext(L_temporalesContext.class,0);
		}
		public L_temporalesContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_l_temporales; }
	}

	public final L_temporalesContext l_temporales() throws RecognitionException {
		return l_temporales(0);
	}

	private L_temporalesContext l_temporales(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		L_temporalesContext _localctx = new L_temporalesContext(_ctx, _parentState);
		L_temporalesContext _prevctx = _localctx;
		int _startState = 8;
		enterRecursionRule(_localctx, 8, RULE_l_temporales, _p);

		      _localctx.ltemps =  arrayList.New()
		  
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(66);
			((L_temporalesContext)_localctx).TK_TEMPORAL = match(TK_TEMPORAL);
			 _localctx.ltemps.Add((((L_temporalesContext)_localctx).TK_TEMPORAL!=null?((L_temporalesContext)_localctx).TK_TEMPORAL.getText():null))
			}
			_ctx.stop = _input.LT(-1);
			setState(75);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,1,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new L_temporalesContext(_parentctx, _parentState);
					_localctx.temps = _prevctx;
					_localctx.temps = _prevctx;
					pushNewRecursionContext(_localctx, _startState, RULE_l_temporales);
					setState(69);
					if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
					setState(70);
					match(TK_COMA);
					setState(71);
					((L_temporalesContext)_localctx).TK_TEMPORAL = match(TK_TEMPORAL);

					              ((L_temporalesContext)_localctx).temps.ltemps.Add((((L_temporalesContext)_localctx).TK_TEMPORAL!=null?((L_temporalesContext)_localctx).TK_TEMPORAL.getText():null))
					              _localctx.ltemps=((L_temporalesContext)_localctx).temps.ltemps
					            
					}
					} 
				}
				setState(77);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,1,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	public static class L_funcasContext extends ParserRuleContext {
		public *arrayList.List lfuncas;
		public L_funcasContext funcs;
		public FuncasContext funcas;
		public FuncasContext funcas() {
			return getRuleContext(FuncasContext.class,0);
		}
		public L_funcasContext l_funcas() {
			return getRuleContext(L_funcasContext.class,0);
		}
		public L_funcasContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_l_funcas; }
	}

	public final L_funcasContext l_funcas() throws RecognitionException {
		return l_funcas(0);
	}

	private L_funcasContext l_funcas(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		L_funcasContext _localctx = new L_funcasContext(_ctx, _parentState);
		L_funcasContext _prevctx = _localctx;
		int _startState = 10;
		enterRecursionRule(_localctx, 10, RULE_l_funcas, _p);

		      _localctx.lfuncas =  arrayList.New()
		  
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(79);
			((L_funcasContext)_localctx).funcas = funcas();
			 _localctx.lfuncas.Add(((L_funcasContext)_localctx).funcas.instr)
			}
			_ctx.stop = _input.LT(-1);
			setState(88);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,2,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new L_funcasContext(_parentctx, _parentState);
					_localctx.funcs = _prevctx;
					_localctx.funcs = _prevctx;
					pushNewRecursionContext(_localctx, _startState, RULE_l_funcas);
					setState(82);
					if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
					setState(83);
					((L_funcasContext)_localctx).funcas = funcas();

					              ((L_funcasContext)_localctx).funcs.lfuncas.Add(((L_funcasContext)_localctx).funcas.instr)
					              _localctx.lfuncas=((L_funcasContext)_localctx).funcs.lfuncas
					            
					}
					} 
				}
				setState(90);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,2,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	public static class FuncasContext extends ParserRuleContext {
		public inter.Instruccion instr;
		public Token TK_IDENTIFICADOR;
		public L_bloqueContext l_bloque;
		public Token TK_MAIN;
		public TerminalNode TK_VOID() { return getToken(optimizador_parser.TK_VOID, 0); }
		public TerminalNode TK_IDENTIFICADOR() { return getToken(optimizador_parser.TK_IDENTIFICADOR, 0); }
		public TerminalNode TK_PI() { return getToken(optimizador_parser.TK_PI, 0); }
		public TerminalNode TK_PD() { return getToken(optimizador_parser.TK_PD, 0); }
		public TerminalNode TK_LI() { return getToken(optimizador_parser.TK_LI, 0); }
		public L_bloqueContext l_bloque() {
			return getRuleContext(L_bloqueContext.class,0);
		}
		public TerminalNode TK_RETURN() { return getToken(optimizador_parser.TK_RETURN, 0); }
		public TerminalNode TK_LD() { return getToken(optimizador_parser.TK_LD, 0); }
		public TerminalNode TK_INT() { return getToken(optimizador_parser.TK_INT, 0); }
		public TerminalNode TK_MAIN() { return getToken(optimizador_parser.TK_MAIN, 0); }
		public TerminalNode TK_ASIGPSTACK() { return getToken(optimizador_parser.TK_ASIGPSTACK, 0); }
		public TerminalNode TK_ASIGHEAP() { return getToken(optimizador_parser.TK_ASIGHEAP, 0); }
		public TerminalNode TK_RETMAIN() { return getToken(optimizador_parser.TK_RETMAIN, 0); }
		public FuncasContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_funcas; }
	}

	public final FuncasContext funcas() throws RecognitionException {
		FuncasContext _localctx = new FuncasContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_funcas);
		try {
			setState(113);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case TK_VOID:
				enterOuterAlt(_localctx, 1);
				{
				setState(91);
				match(TK_VOID);
				setState(92);
				((FuncasContext)_localctx).TK_IDENTIFICADOR = match(TK_IDENTIFICADOR);
				setState(93);
				match(TK_PI);
				setState(94);
				match(TK_PD);
				setState(95);
				match(TK_LI);
				setState(96);
				((FuncasContext)_localctx).l_bloque = l_bloque(0);
				setState(97);
				match(TK_RETURN);
				setState(98);
				match(TK_LD);

				    _localctx.instr = funcion.Nfunca((((FuncasContext)_localctx).TK_IDENTIFICADOR!=null?((FuncasContext)_localctx).TK_IDENTIFICADOR.getText():null), ((FuncasContext)_localctx).l_bloque.lbloque)
				  
				}
				break;
			case TK_INT:
				enterOuterAlt(_localctx, 2);
				{
				setState(101);
				match(TK_INT);
				setState(102);
				((FuncasContext)_localctx).TK_MAIN = match(TK_MAIN);
				setState(103);
				match(TK_PI);
				setState(104);
				match(TK_PD);
				setState(105);
				match(TK_LI);
				setState(106);
				match(TK_ASIGPSTACK);
				setState(107);
				match(TK_ASIGHEAP);
				setState(108);
				((FuncasContext)_localctx).l_bloque = l_bloque(0);
				setState(109);
				match(TK_RETMAIN);
				setState(110);
				match(TK_LD);

				    _localctx.instr = funcion.Nfuncamain((((FuncasContext)_localctx).TK_MAIN!=null?((FuncasContext)_localctx).TK_MAIN.getText():null), ((FuncasContext)_localctx).l_bloque.lbloque)
				  
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Sent_ifContext extends ParserRuleContext {
		public inter.Instruccion instr;
		public Token TK_IF;
		public ExpressionContext expression;
		public Token TK_ETIQUETA;
		public TerminalNode TK_IF() { return getToken(optimizador_parser.TK_IF, 0); }
		public TerminalNode TK_PI() { return getToken(optimizador_parser.TK_PI, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode TK_PD() { return getToken(optimizador_parser.TK_PD, 0); }
		public TerminalNode TK_GOTO() { return getToken(optimizador_parser.TK_GOTO, 0); }
		public TerminalNode TK_ETIQUETA() { return getToken(optimizador_parser.TK_ETIQUETA, 0); }
		public TerminalNode TK_PYC() { return getToken(optimizador_parser.TK_PYC, 0); }
		public Sent_ifContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_sent_if; }
	}

	public final Sent_ifContext sent_if() throws RecognitionException {
		Sent_ifContext _localctx = new Sent_ifContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_sent_if);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(115);
			((Sent_ifContext)_localctx).TK_IF = match(TK_IF);
			setState(116);
			match(TK_PI);
			setState(117);
			((Sent_ifContext)_localctx).expression = expression();
			setState(118);
			match(TK_PD);
			setState(119);
			match(TK_GOTO);
			setState(120);
			((Sent_ifContext)_localctx).TK_ETIQUETA = match(TK_ETIQUETA);
			setState(121);
			match(TK_PYC);

			    _localctx.instr = sif.Nsenteif(((Sent_ifContext)_localctx).expression.p,(((Sent_ifContext)_localctx).TK_ETIQUETA!=null?((Sent_ifContext)_localctx).TK_ETIQUETA.getText():null), (((Sent_ifContext)_localctx).TK_IF!=null?((Sent_ifContext)_localctx).TK_IF.getLine():0))
			  
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ImprimirContext extends ParserRuleContext {
		public inter.Instruccion instr;
		public Token TK_PRINTF;
		public Token TK_CADENA;
		public CasteoContext casteo;
		public ExpressionContext expression;
		public TerminalNode TK_PRINTF() { return getToken(optimizador_parser.TK_PRINTF, 0); }
		public TerminalNode TK_PI() { return getToken(optimizador_parser.TK_PI, 0); }
		public TerminalNode TK_CADENA() { return getToken(optimizador_parser.TK_CADENA, 0); }
		public TerminalNode TK_COMA() { return getToken(optimizador_parser.TK_COMA, 0); }
		public CasteoContext casteo() {
			return getRuleContext(CasteoContext.class,0);
		}
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode TK_PD() { return getToken(optimizador_parser.TK_PD, 0); }
		public TerminalNode TK_PYC() { return getToken(optimizador_parser.TK_PYC, 0); }
		public ImprimirContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_imprimir; }
	}

	public final ImprimirContext imprimir() throws RecognitionException {
		ImprimirContext _localctx = new ImprimirContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_imprimir);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(124);
			((ImprimirContext)_localctx).TK_PRINTF = match(TK_PRINTF);
			setState(125);
			match(TK_PI);
			setState(126);
			((ImprimirContext)_localctx).TK_CADENA = match(TK_CADENA);
			setState(127);
			match(TK_COMA);
			setState(128);
			((ImprimirContext)_localctx).casteo = casteo();
			setState(129);
			((ImprimirContext)_localctx).expression = expression();
			setState(130);
			match(TK_PD);
			setState(131);
			match(TK_PYC);

			    _localctx.instr = imprimir.Nimprimir((((ImprimirContext)_localctx).TK_CADENA!=null?((ImprimirContext)_localctx).TK_CADENA.getText():null), ((ImprimirContext)_localctx).casteo.cast, ((ImprimirContext)_localctx).expression.p,(((ImprimirContext)_localctx).TK_PRINTF!=null?((ImprimirContext)_localctx).TK_PRINTF.getLine():0))
			  
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class CasteoContext extends ParserRuleContext {
		public string cast;
		public Token TK_CASTCHAR;
		public Token TK_CASTINT;
		public TerminalNode TK_CASTCHAR() { return getToken(optimizador_parser.TK_CASTCHAR, 0); }
		public TerminalNode TK_CASTINT() { return getToken(optimizador_parser.TK_CASTINT, 0); }
		public CasteoContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_casteo; }
	}

	public final CasteoContext casteo() throws RecognitionException {
		CasteoContext _localctx = new CasteoContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_casteo);
		try {
			setState(139);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case TK_CASTCHAR:
				enterOuterAlt(_localctx, 1);
				{
				setState(134);
				((CasteoContext)_localctx).TK_CASTCHAR = match(TK_CASTCHAR);
				_localctx.cast = (((CasteoContext)_localctx).TK_CASTCHAR!=null?((CasteoContext)_localctx).TK_CASTCHAR.getText():null)
				}
				break;
			case TK_CASTINT:
				enterOuterAlt(_localctx, 2);
				{
				setState(136);
				((CasteoContext)_localctx).TK_CASTINT = match(TK_CASTINT);
				_localctx.cast = (((CasteoContext)_localctx).TK_CASTINT!=null?((CasteoContext)_localctx).TK_CASTINT.getText():null)
				}
				break;
			case TK_HEAP:
			case TK_STACK:
			case TK_PUNSTACK:
			case TK_PUNHEAP:
			case TK_FLOAT:
			case TK_ENTERO:
			case TK_CADENA:
			case TK_TEMPORAL:
			case TK_IDENTIFICADOR:
			case TK_RESTA:
				enterOuterAlt(_localctx, 3);
				{
				_localctx.cast = ""
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class DeclaracionContext extends ParserRuleContext {
		public inter.Instruccion instr;
		public Token TK_TEMPORAL;
		public ExpressionContext expression;
		public Token nom;
		public ExpressionContext pos;
		public ExpressionContext exp;
		public TerminalNode TK_TEMPORAL() { return getToken(optimizador_parser.TK_TEMPORAL, 0); }
		public TerminalNode TK_IGUAL() { return getToken(optimizador_parser.TK_IGUAL, 0); }
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public TerminalNode TK_PYC() { return getToken(optimizador_parser.TK_PYC, 0); }
		public TerminalNode TK_CI() { return getToken(optimizador_parser.TK_CI, 0); }
		public TerminalNode TK_CASTINT() { return getToken(optimizador_parser.TK_CASTINT, 0); }
		public TerminalNode TK_CD() { return getToken(optimizador_parser.TK_CD, 0); }
		public TerminalNode TK_STACK() { return getToken(optimizador_parser.TK_STACK, 0); }
		public TerminalNode TK_HEAP() { return getToken(optimizador_parser.TK_HEAP, 0); }
		public TerminalNode TK_PUNSTACK() { return getToken(optimizador_parser.TK_PUNSTACK, 0); }
		public TerminalNode TK_PUNHEAP() { return getToken(optimizador_parser.TK_PUNHEAP, 0); }
		public DeclaracionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_declaracion; }
	}

	public final DeclaracionContext declaracion() throws RecognitionException {
		DeclaracionContext _localctx = new DeclaracionContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_declaracion);
		int _la;
		try {
			setState(163);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case TK_TEMPORAL:
				enterOuterAlt(_localctx, 1);
				{
				setState(141);
				((DeclaracionContext)_localctx).TK_TEMPORAL = match(TK_TEMPORAL);
				setState(142);
				match(TK_IGUAL);
				setState(143);
				((DeclaracionContext)_localctx).expression = expression();
				setState(144);
				match(TK_PYC);

				    _localctx.instr = declarar.Ndeclaracion((((DeclaracionContext)_localctx).TK_TEMPORAL!=null?((DeclaracionContext)_localctx).TK_TEMPORAL.getText():null), ((DeclaracionContext)_localctx).expression.p,(((DeclaracionContext)_localctx).TK_TEMPORAL!=null?((DeclaracionContext)_localctx).TK_TEMPORAL.getLine():0))
				  
				}
				break;
			case TK_HEAP:
			case TK_STACK:
				enterOuterAlt(_localctx, 2);
				{
				setState(147);
				((DeclaracionContext)_localctx).nom = _input.LT(1);
				_la = _input.LA(1);
				if ( !(_la==TK_HEAP || _la==TK_STACK) ) {
					((DeclaracionContext)_localctx).nom = (Token)_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				setState(148);
				match(TK_CI);
				setState(149);
				match(TK_CASTINT);
				setState(150);
				((DeclaracionContext)_localctx).pos = expression();
				setState(151);
				match(TK_CD);
				setState(152);
				match(TK_IGUAL);
				setState(153);
				((DeclaracionContext)_localctx).exp = expression();
				setState(154);
				match(TK_PYC);

				    _localctx.instr = declarar.Ndeclapunteros((((DeclaracionContext)_localctx).nom!=null?((DeclaracionContext)_localctx).nom.getText():null), ((DeclaracionContext)_localctx).pos.p, ((DeclaracionContext)_localctx).exp.p, (((DeclaracionContext)_localctx).nom!=null?((DeclaracionContext)_localctx).nom.getLine():0))
				  
				}
				break;
			case TK_PUNSTACK:
			case TK_PUNHEAP:
				enterOuterAlt(_localctx, 3);
				{
				setState(157);
				((DeclaracionContext)_localctx).nom = _input.LT(1);
				_la = _input.LA(1);
				if ( !(_la==TK_PUNSTACK || _la==TK_PUNHEAP) ) {
					((DeclaracionContext)_localctx).nom = (Token)_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				setState(158);
				match(TK_IGUAL);
				setState(159);
				((DeclaracionContext)_localctx).expression = expression();
				setState(160);
				match(TK_PYC);

				    _localctx.instr = declarar.Ndeclapstack((((DeclaracionContext)_localctx).nom!=null?((DeclaracionContext)_localctx).nom.getText():null), ((DeclaracionContext)_localctx).expression.p,(((DeclaracionContext)_localctx).nom!=null?((DeclaracionContext)_localctx).nom.getLine():0))
				  
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class EtiquetasContext extends ParserRuleContext {
		public inter.Instruccion instr;
		public Token TK_ETIQUETA;
		public TerminalNode TK_ETIQUETA() { return getToken(optimizador_parser.TK_ETIQUETA, 0); }
		public TerminalNode TK_DP() { return getToken(optimizador_parser.TK_DP, 0); }
		public TerminalNode TK_GOTO() { return getToken(optimizador_parser.TK_GOTO, 0); }
		public TerminalNode TK_PYC() { return getToken(optimizador_parser.TK_PYC, 0); }
		public EtiquetasContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_etiquetas; }
	}

	public final EtiquetasContext etiquetas() throws RecognitionException {
		EtiquetasContext _localctx = new EtiquetasContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_etiquetas);
		try {
			setState(172);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case TK_ETIQUETA:
				enterOuterAlt(_localctx, 1);
				{
				setState(165);
				((EtiquetasContext)_localctx).TK_ETIQUETA = match(TK_ETIQUETA);
				setState(166);
				match(TK_DP);

				    _localctx.instr = etiquetas.Netiqueta((((EtiquetasContext)_localctx).TK_ETIQUETA!=null?((EtiquetasContext)_localctx).TK_ETIQUETA.getText():null))
				  
				}
				break;
			case TK_GOTO:
				enterOuterAlt(_localctx, 2);
				{
				setState(168);
				match(TK_GOTO);
				setState(169);
				((EtiquetasContext)_localctx).TK_ETIQUETA = match(TK_ETIQUETA);
				setState(170);
				match(TK_PYC);

				    _localctx.instr = etiquetas.Ngoto((((EtiquetasContext)_localctx).TK_ETIQUETA!=null?((EtiquetasContext)_localctx).TK_ETIQUETA.getText():null))
				  
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class LlamadaContext extends ParserRuleContext {
		public inter.Instruccion instr;
		public Token TK_IDENTIFICADOR;
		public TerminalNode TK_IDENTIFICADOR() { return getToken(optimizador_parser.TK_IDENTIFICADOR, 0); }
		public TerminalNode TK_PI() { return getToken(optimizador_parser.TK_PI, 0); }
		public TerminalNode TK_PD() { return getToken(optimizador_parser.TK_PD, 0); }
		public TerminalNode TK_PYC() { return getToken(optimizador_parser.TK_PYC, 0); }
		public LlamadaContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_llamada; }
	}

	public final LlamadaContext llamada() throws RecognitionException {
		LlamadaContext _localctx = new LlamadaContext(_ctx, getState());
		enterRule(_localctx, 24, RULE_llamada);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(174);
			((LlamadaContext)_localctx).TK_IDENTIFICADOR = match(TK_IDENTIFICADOR);
			setState(175);
			match(TK_PI);
			setState(176);
			match(TK_PD);
			setState(177);
			match(TK_PYC);

			    _localctx.instr= llamada.Nllama((((LlamadaContext)_localctx).TK_IDENTIFICADOR!=null?((LlamadaContext)_localctx).TK_IDENTIFICADOR.getText():null))
			  
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class L_bloqueContext extends ParserRuleContext {
		public *arrayList.List lbloque;
		public L_bloqueContext block;
		public BloqueContext bloque;
		public BloqueContext bloque() {
			return getRuleContext(BloqueContext.class,0);
		}
		public L_bloqueContext l_bloque() {
			return getRuleContext(L_bloqueContext.class,0);
		}
		public L_bloqueContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_l_bloque; }
	}

	public final L_bloqueContext l_bloque() throws RecognitionException {
		return l_bloque(0);
	}

	private L_bloqueContext l_bloque(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		L_bloqueContext _localctx = new L_bloqueContext(_ctx, _parentState);
		L_bloqueContext _prevctx = _localctx;
		int _startState = 26;
		enterRecursionRule(_localctx, 26, RULE_l_bloque, _p);

		    _localctx.lbloque =  arrayList.New()
		  
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(181);
			((L_bloqueContext)_localctx).bloque = bloque();
			 _localctx.lbloque.Add(((L_bloqueContext)_localctx).bloque.instr)
			}
			_ctx.stop = _input.LT(-1);
			setState(190);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,7,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new L_bloqueContext(_parentctx, _parentState);
					_localctx.block = _prevctx;
					_localctx.block = _prevctx;
					pushNewRecursionContext(_localctx, _startState, RULE_l_bloque);
					setState(184);
					if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
					setState(185);
					((L_bloqueContext)_localctx).bloque = bloque();

					              ((L_bloqueContext)_localctx).block.lbloque.Add(((L_bloqueContext)_localctx).bloque.instr)
					              _localctx.lbloque=((L_bloqueContext)_localctx).block.lbloque
					            
					}
					} 
				}
				setState(192);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,7,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	public static class BloqueContext extends ParserRuleContext {
		public inter.Instruccion instr;
		public ImprimirContext imprimir;
		public Sent_ifContext sent_if;
		public DeclaracionContext declaracion;
		public EtiquetasContext etiquetas;
		public LlamadaContext llamada;
		public ImprimirContext imprimir() {
			return getRuleContext(ImprimirContext.class,0);
		}
		public Sent_ifContext sent_if() {
			return getRuleContext(Sent_ifContext.class,0);
		}
		public DeclaracionContext declaracion() {
			return getRuleContext(DeclaracionContext.class,0);
		}
		public EtiquetasContext etiquetas() {
			return getRuleContext(EtiquetasContext.class,0);
		}
		public LlamadaContext llamada() {
			return getRuleContext(LlamadaContext.class,0);
		}
		public BloqueContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_bloque; }
	}

	public final BloqueContext bloque() throws RecognitionException {
		BloqueContext _localctx = new BloqueContext(_ctx, getState());
		enterRule(_localctx, 28, RULE_bloque);
		try {
			setState(208);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case TK_PRINTF:
				enterOuterAlt(_localctx, 1);
				{
				setState(193);
				((BloqueContext)_localctx).imprimir = imprimir();
				_localctx.instr = ((BloqueContext)_localctx).imprimir.instr
				}
				break;
			case TK_IF:
				enterOuterAlt(_localctx, 2);
				{
				setState(196);
				((BloqueContext)_localctx).sent_if = sent_if();
				_localctx.instr = ((BloqueContext)_localctx).sent_if.instr
				}
				break;
			case TK_HEAP:
			case TK_STACK:
			case TK_PUNSTACK:
			case TK_PUNHEAP:
			case TK_TEMPORAL:
				enterOuterAlt(_localctx, 3);
				{
				setState(199);
				((BloqueContext)_localctx).declaracion = declaracion();
				_localctx.instr = ((BloqueContext)_localctx).declaracion.instr
				}
				break;
			case TK_GOTO:
			case TK_ETIQUETA:
				enterOuterAlt(_localctx, 4);
				{
				setState(202);
				((BloqueContext)_localctx).etiquetas = etiquetas();
				_localctx.instr = ((BloqueContext)_localctx).etiquetas.instr
				}
				break;
			case TK_IDENTIFICADOR:
				enterOuterAlt(_localctx, 5);
				{
				setState(205);
				((BloqueContext)_localctx).llamada = llamada();
				_localctx.instr = ((BloqueContext)_localctx).llamada.instr
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ExpressionContext extends ParserRuleContext {
		public inter.Expresion p;
		public Expre_relacionalContext expre_relacional;
		public Expre_aritmeticaContext expre_aritmetica;
		public AsignacionesContext asignaciones;
		public Expre_relacionalContext expre_relacional() {
			return getRuleContext(Expre_relacionalContext.class,0);
		}
		public Expre_aritmeticaContext expre_aritmetica() {
			return getRuleContext(Expre_aritmeticaContext.class,0);
		}
		public AsignacionesContext asignaciones() {
			return getRuleContext(AsignacionesContext.class,0);
		}
		public ExpressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expression; }
	}

	public final ExpressionContext expression() throws RecognitionException {
		ExpressionContext _localctx = new ExpressionContext(_ctx, getState());
		enterRule(_localctx, 30, RULE_expression);
		try {
			setState(219);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,9,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(210);
				((ExpressionContext)_localctx).expre_relacional = expre_relacional(0);
				_localctx.p = ((ExpressionContext)_localctx).expre_relacional.p
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(213);
				((ExpressionContext)_localctx).expre_aritmetica = expre_aritmetica(0);
				_localctx.p = ((ExpressionContext)_localctx).expre_aritmetica.p
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(216);
				((ExpressionContext)_localctx).asignaciones = asignaciones();
				_localctx.p = ((ExpressionContext)_localctx).asignaciones.p
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class AsignacionesContext extends ParserRuleContext {
		public inter.Expresion p;
		public ExpressionContext expression;
		public TerminalNode TK_HEAP() { return getToken(optimizador_parser.TK_HEAP, 0); }
		public TerminalNode TK_CI() { return getToken(optimizador_parser.TK_CI, 0); }
		public TerminalNode TK_CASTINT() { return getToken(optimizador_parser.TK_CASTINT, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode TK_CD() { return getToken(optimizador_parser.TK_CD, 0); }
		public TerminalNode TK_STACK() { return getToken(optimizador_parser.TK_STACK, 0); }
		public AsignacionesContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_asignaciones; }
	}

	public final AsignacionesContext asignaciones() throws RecognitionException {
		AsignacionesContext _localctx = new AsignacionesContext(_ctx, getState());
		enterRule(_localctx, 32, RULE_asignaciones);
		try {
			setState(235);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case TK_HEAP:
				enterOuterAlt(_localctx, 1);
				{
				setState(221);
				match(TK_HEAP);
				setState(222);
				match(TK_CI);
				setState(223);
				match(TK_CASTINT);
				setState(224);
				((AsignacionesContext)_localctx).expression = expression();
				setState(225);
				match(TK_CD);

				    _localctx.p = asignaciones.Nasiheap(((AsignacionesContext)_localctx).expression.p)
				  
				}
				break;
			case TK_STACK:
				enterOuterAlt(_localctx, 2);
				{
				setState(228);
				match(TK_STACK);
				setState(229);
				match(TK_CI);
				setState(230);
				match(TK_CASTINT);
				setState(231);
				((AsignacionesContext)_localctx).expression = expression();
				setState(232);
				match(TK_CD);

				    _localctx.p = asignaciones.Nasistack(((AsignacionesContext)_localctx).expression.p)
				  
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Expre_relacionalContext extends ParserRuleContext {
		public inter.Expresion p;
		public Expre_relacionalContext opIz;
		public Expre_aritmeticaContext expre_aritmetica;
		public Token op;
		public Expre_relacionalContext opDe;
		public Expre_aritmeticaContext expre_aritmetica() {
			return getRuleContext(Expre_aritmeticaContext.class,0);
		}
		public List<Expre_relacionalContext> expre_relacional() {
			return getRuleContexts(Expre_relacionalContext.class);
		}
		public Expre_relacionalContext expre_relacional(int i) {
			return getRuleContext(Expre_relacionalContext.class,i);
		}
		public TerminalNode TK_MAYORIGUAL() { return getToken(optimizador_parser.TK_MAYORIGUAL, 0); }
		public TerminalNode TK_MENORIGUAL() { return getToken(optimizador_parser.TK_MENORIGUAL, 0); }
		public TerminalNode TK_IGUALDAD() { return getToken(optimizador_parser.TK_IGUALDAD, 0); }
		public TerminalNode TK_DESIGUALDAD() { return getToken(optimizador_parser.TK_DESIGUALDAD, 0); }
		public TerminalNode TK_MAYOR() { return getToken(optimizador_parser.TK_MAYOR, 0); }
		public TerminalNode TK_MENOR() { return getToken(optimizador_parser.TK_MENOR, 0); }
		public Expre_relacionalContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expre_relacional; }
	}

	public final Expre_relacionalContext expre_relacional() throws RecognitionException {
		return expre_relacional(0);
	}

	private Expre_relacionalContext expre_relacional(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		Expre_relacionalContext _localctx = new Expre_relacionalContext(_ctx, _parentState);
		Expre_relacionalContext _prevctx = _localctx;
		int _startState = 34;
		enterRecursionRule(_localctx, 34, RULE_expre_relacional, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(238);
			((Expre_relacionalContext)_localctx).expre_aritmetica = expre_aritmetica(0);
			_localctx.p = ((Expre_relacionalContext)_localctx).expre_aritmetica.p
			}
			_ctx.stop = _input.LT(-1);
			setState(248);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,11,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new Expre_relacionalContext(_parentctx, _parentState);
					_localctx.opIz = _prevctx;
					_localctx.opIz = _prevctx;
					pushNewRecursionContext(_localctx, _startState, RULE_expre_relacional);
					setState(241);
					if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
					setState(242);
					((Expre_relacionalContext)_localctx).op = _input.LT(1);
					_la = _input.LA(1);
					if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << TK_MENORIGUAL) | (1L << TK_MAYORIGUAL) | (1L << TK_IGUALDAD) | (1L << TK_DESIGUALDAD) | (1L << TK_MENOR) | (1L << TK_MAYOR))) != 0)) ) {
						((Expre_relacionalContext)_localctx).op = (Token)_errHandler.recoverInline(this);
					}
					else {
						if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
						_errHandler.reportMatch(this);
						consume();
					}
					setState(243);
					((Expre_relacionalContext)_localctx).opDe = expre_relacional(3);

					                if (((Expre_relacionalContext)_localctx).op!=null?((Expre_relacionalContext)_localctx).op.getText():null) == "<"{
					                  _localctx.p = menor.Noperacionmenor(((Expre_relacionalContext)_localctx).opIz.p,((Expre_relacionalContext)_localctx).opDe.p)
					                }else if (((Expre_relacionalContext)_localctx).op!=null?((Expre_relacionalContext)_localctx).op.getText():null) == "<="{
					                  _localctx.p = menorigual.Noperacionmenorigual(((Expre_relacionalContext)_localctx).opIz.p,((Expre_relacionalContext)_localctx).opDe.p)
					                }else if (((Expre_relacionalContext)_localctx).op!=null?((Expre_relacionalContext)_localctx).op.getText():null) == ">"{
					                  _localctx.p = mayor.Noperacionmayor(((Expre_relacionalContext)_localctx).opIz.p,((Expre_relacionalContext)_localctx).opDe.p)
					                }else if (((Expre_relacionalContext)_localctx).op!=null?((Expre_relacionalContext)_localctx).op.getText():null) == ">="{
					                  _localctx.p = mayorigual.Noperacionmayorigual(((Expre_relacionalContext)_localctx).opIz.p,((Expre_relacionalContext)_localctx).opDe.p)
					                }else if (((Expre_relacionalContext)_localctx).op!=null?((Expre_relacionalContext)_localctx).op.getText():null) == "=="{
					                  _localctx.p = igualdad.Noperacionigualdad(((Expre_relacionalContext)_localctx).opIz.p,((Expre_relacionalContext)_localctx).opDe.p)
					                }else if (((Expre_relacionalContext)_localctx).op!=null?((Expre_relacionalContext)_localctx).op.getText():null) == "!="{
					                  _localctx.p = desigualdad.Noperaciondesigualdad(((Expre_relacionalContext)_localctx).opIz.p,((Expre_relacionalContext)_localctx).opDe.p)
					                }
					            
					}
					} 
				}
				setState(250);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,11,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	public static class Expre_aritmeticaContext extends ParserRuleContext {
		public inter.Expresion p;
		public Expre_aritmeticaContext opIz;
		public Token opera;
		public Expre_aritmeticaContext opUn;
		public DatosContext datos;
		public Expre_aritmeticaContext opDe;
		public TerminalNode TK_RESTA() { return getToken(optimizador_parser.TK_RESTA, 0); }
		public List<Expre_aritmeticaContext> expre_aritmetica() {
			return getRuleContexts(Expre_aritmeticaContext.class);
		}
		public Expre_aritmeticaContext expre_aritmetica(int i) {
			return getRuleContext(Expre_aritmeticaContext.class,i);
		}
		public DatosContext datos() {
			return getRuleContext(DatosContext.class,0);
		}
		public TerminalNode TK_MULTI() { return getToken(optimizador_parser.TK_MULTI, 0); }
		public TerminalNode TK_DIVI() { return getToken(optimizador_parser.TK_DIVI, 0); }
		public TerminalNode TK_MODULO() { return getToken(optimizador_parser.TK_MODULO, 0); }
		public TerminalNode TK_SUMA() { return getToken(optimizador_parser.TK_SUMA, 0); }
		public Expre_aritmeticaContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expre_aritmetica; }
	}

	public final Expre_aritmeticaContext expre_aritmetica() throws RecognitionException {
		return expre_aritmetica(0);
	}

	private Expre_aritmeticaContext expre_aritmetica(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		Expre_aritmeticaContext _localctx = new Expre_aritmeticaContext(_ctx, _parentState);
		Expre_aritmeticaContext _prevctx = _localctx;
		int _startState = 36;
		enterRecursionRule(_localctx, 36, RULE_expre_aritmetica, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(259);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case TK_RESTA:
				{
				setState(252);
				((Expre_aritmeticaContext)_localctx).opera = match(TK_RESTA);
				setState(253);
				((Expre_aritmeticaContext)_localctx).opUn = expre_aritmetica(4);

				      _localctx.p = negativo.Nnegativo(((Expre_aritmeticaContext)_localctx).opUn.p)
				    
				}
				break;
			case TK_PUNSTACK:
			case TK_PUNHEAP:
			case TK_FLOAT:
			case TK_ENTERO:
			case TK_CADENA:
			case TK_TEMPORAL:
			case TK_IDENTIFICADOR:
				{
				setState(256);
				((Expre_aritmeticaContext)_localctx).datos = datos();
				_localctx.p = ((Expre_aritmeticaContext)_localctx).datos.p
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			_ctx.stop = _input.LT(-1);
			setState(273);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,14,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(271);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,13,_ctx) ) {
					case 1:
						{
						_localctx = new Expre_aritmeticaContext(_parentctx, _parentState);
						_localctx.opIz = _prevctx;
						_localctx.opIz = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expre_aritmetica);
						setState(261);
						if (!(precpred(_ctx, 3))) throw new FailedPredicateException(this, "precpred(_ctx, 3)");
						setState(262);
						((Expre_aritmeticaContext)_localctx).opera = _input.LT(1);
						_la = _input.LA(1);
						if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << TK_MULTI) | (1L << TK_DIVI) | (1L << TK_MODULO))) != 0)) ) {
							((Expre_aritmeticaContext)_localctx).opera = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(263);
						((Expre_aritmeticaContext)_localctx).opDe = expre_aritmetica(4);

						                if (((Expre_aritmeticaContext)_localctx).opera!=null?((Expre_aritmeticaContext)_localctx).opera.getText():null) == "*"{ 
						                  _localctx.p = multiplicacion.Noperacionmultiplicacion(((Expre_aritmeticaContext)_localctx).opIz.p,((Expre_aritmeticaContext)_localctx).opDe.p)
						                }else if (((Expre_aritmeticaContext)_localctx).opera!=null?((Expre_aritmeticaContext)_localctx).opera.getText():null) == "/"{
						                  _localctx.p = division.Noperaciondivision(((Expre_aritmeticaContext)_localctx).opIz.p,((Expre_aritmeticaContext)_localctx).opDe.p)
						                }else if (((Expre_aritmeticaContext)_localctx).opera!=null?((Expre_aritmeticaContext)_localctx).opera.getText():null) == "%"{
						                  _localctx.p = modulo.Noperacionmodulo(((Expre_aritmeticaContext)_localctx).opIz.p,((Expre_aritmeticaContext)_localctx).opDe.p)
						                }
						              
						}
						break;
					case 2:
						{
						_localctx = new Expre_aritmeticaContext(_parentctx, _parentState);
						_localctx.opIz = _prevctx;
						_localctx.opIz = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expre_aritmetica);
						setState(266);
						if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
						setState(267);
						((Expre_aritmeticaContext)_localctx).opera = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==TK_SUMA || _la==TK_RESTA) ) {
							((Expre_aritmeticaContext)_localctx).opera = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(268);
						((Expre_aritmeticaContext)_localctx).opDe = expre_aritmetica(3);
						        
						                if (((Expre_aritmeticaContext)_localctx).opera!=null?((Expre_aritmeticaContext)_localctx).opera.getText():null)=="+"{
						                  _localctx.p = suma.Noperacionsuma(((Expre_aritmeticaContext)_localctx).opIz.p,((Expre_aritmeticaContext)_localctx).opDe.p)
						                }else{ 
						                  _localctx.p = resta.Noperacionresta(((Expre_aritmeticaContext)_localctx).opIz.p,((Expre_aritmeticaContext)_localctx).opDe.p)
						                }
						              
						}
						break;
					}
					} 
				}
				setState(275);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,14,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	public static class DatosContext extends ParserRuleContext {
		public inter.Expresion p;
		public Token TK_ENTERO;
		public Token TK_FLOAT;
		public Token TK_IDENTIFICADOR;
		public Token TK_CADENA;
		public Token TK_TEMPORAL;
		public Token TK_PUNSTACK;
		public Token TK_PUNHEAP;
		public TerminalNode TK_ENTERO() { return getToken(optimizador_parser.TK_ENTERO, 0); }
		public TerminalNode TK_FLOAT() { return getToken(optimizador_parser.TK_FLOAT, 0); }
		public TerminalNode TK_IDENTIFICADOR() { return getToken(optimizador_parser.TK_IDENTIFICADOR, 0); }
		public TerminalNode TK_CADENA() { return getToken(optimizador_parser.TK_CADENA, 0); }
		public TerminalNode TK_TEMPORAL() { return getToken(optimizador_parser.TK_TEMPORAL, 0); }
		public TerminalNode TK_PUNSTACK() { return getToken(optimizador_parser.TK_PUNSTACK, 0); }
		public TerminalNode TK_PUNHEAP() { return getToken(optimizador_parser.TK_PUNHEAP, 0); }
		public DatosContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_datos; }
	}

	public final DatosContext datos() throws RecognitionException {
		DatosContext _localctx = new DatosContext(_ctx, getState());
		enterRule(_localctx, 38, RULE_datos);
		try {
			setState(290);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case TK_ENTERO:
				enterOuterAlt(_localctx, 1);
				{
				setState(276);
				((DatosContext)_localctx).TK_ENTERO = match(TK_ENTERO);

				      _localctx.p = primitiivos.NuevoPrimitivo ((((DatosContext)_localctx).TK_ENTERO!=null?((DatosContext)_localctx).TK_ENTERO.getText():null),bloque.INTEGER)
				    
				}
				break;
			case TK_FLOAT:
				enterOuterAlt(_localctx, 2);
				{
				setState(278);
				((DatosContext)_localctx).TK_FLOAT = match(TK_FLOAT);

				      _localctx.p = primitiivos.NuevoPrimitivo ((((DatosContext)_localctx).TK_FLOAT!=null?((DatosContext)_localctx).TK_FLOAT.getText():null),bloque.FLOAT)
				    
				}
				break;
			case TK_IDENTIFICADOR:
				enterOuterAlt(_localctx, 3);
				{
				setState(280);
				((DatosContext)_localctx).TK_IDENTIFICADOR = match(TK_IDENTIFICADOR);

				      _localctx.p = primitiivos.NuevoPrimitivo((((DatosContext)_localctx).TK_IDENTIFICADOR!=null?((DatosContext)_localctx).TK_IDENTIFICADOR.getText():null),bloque.NULL)
				    
				}
				break;
			case TK_CADENA:
				enterOuterAlt(_localctx, 4);
				{
				setState(282);
				((DatosContext)_localctx).TK_CADENA = match(TK_CADENA);
				 
				      _localctx.p = primitiivos.NuevoPrimitivo((((DatosContext)_localctx).TK_CADENA!=null?((DatosContext)_localctx).TK_CADENA.getText():null),bloque.NULL)
				    
				}
				break;
			case TK_TEMPORAL:
				enterOuterAlt(_localctx, 5);
				{
				setState(284);
				((DatosContext)_localctx).TK_TEMPORAL = match(TK_TEMPORAL);

				      _localctx.p = primitiivos.NuevoPrimitivo((((DatosContext)_localctx).TK_TEMPORAL!=null?((DatosContext)_localctx).TK_TEMPORAL.getText():null),bloque.TEMPORAL)
				    
				}
				break;
			case TK_PUNSTACK:
				enterOuterAlt(_localctx, 6);
				{
				setState(286);
				((DatosContext)_localctx).TK_PUNSTACK = match(TK_PUNSTACK);

				      _localctx.p = primitiivos.NuevoPrimitivo((((DatosContext)_localctx).TK_PUNSTACK!=null?((DatosContext)_localctx).TK_PUNSTACK.getText():null),bloque.NULL)
				    
				}
				break;
			case TK_PUNHEAP:
				enterOuterAlt(_localctx, 7);
				{
				setState(288);
				((DatosContext)_localctx).TK_PUNHEAP = match(TK_PUNHEAP);

				      _localctx.p = primitiivos.NuevoPrimitivo((((DatosContext)_localctx).TK_PUNHEAP!=null?((DatosContext)_localctx).TK_PUNHEAP.getText():null),bloque.NULL)
				    
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public boolean sempred(RuleContext _localctx, int ruleIndex, int predIndex) {
		switch (ruleIndex) {
		case 4:
			return l_temporales_sempred((L_temporalesContext)_localctx, predIndex);
		case 5:
			return l_funcas_sempred((L_funcasContext)_localctx, predIndex);
		case 13:
			return l_bloque_sempred((L_bloqueContext)_localctx, predIndex);
		case 17:
			return expre_relacional_sempred((Expre_relacionalContext)_localctx, predIndex);
		case 18:
			return expre_aritmetica_sempred((Expre_aritmeticaContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean l_temporales_sempred(L_temporalesContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 2);
		}
		return true;
	}
	private boolean l_funcas_sempred(L_funcasContext _localctx, int predIndex) {
		switch (predIndex) {
		case 1:
			return precpred(_ctx, 2);
		}
		return true;
	}
	private boolean l_bloque_sempred(L_bloqueContext _localctx, int predIndex) {
		switch (predIndex) {
		case 2:
			return precpred(_ctx, 2);
		}
		return true;
	}
	private boolean expre_relacional_sempred(Expre_relacionalContext _localctx, int predIndex) {
		switch (predIndex) {
		case 3:
			return precpred(_ctx, 2);
		}
		return true;
	}
	private boolean expre_aritmetica_sempred(Expre_aritmeticaContext _localctx, int predIndex) {
		switch (predIndex) {
		case 4:
			return precpred(_ctx, 3);
		case 5:
			return precpred(_ctx, 2);
		}
		return true;
	}

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3\66\u0127\4\2\t\2"+
		"\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13"+
		"\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\3\2\3\2\3\2\3\3\7\3/\n\3\f\3\16\3\62\13"+
		"\3\3\3\3\3\3\4\3\4\3\4\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\6"+
		"\3\6\3\6\3\6\3\6\3\6\3\6\3\6\7\6L\n\6\f\6\16\6O\13\6\3\7\3\7\3\7\3\7\3"+
		"\7\3\7\3\7\3\7\7\7Y\n\7\f\7\16\7\\\13\7\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3"+
		"\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\5\bt\n\b\3"+
		"\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n"+
		"\3\n\3\13\3\13\3\13\3\13\3\13\5\13\u008e\n\13\3\f\3\f\3\f\3\f\3\f\3\f"+
		"\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\5\f\u00a6"+
		"\n\f\3\r\3\r\3\r\3\r\3\r\3\r\3\r\5\r\u00af\n\r\3\16\3\16\3\16\3\16\3\16"+
		"\3\16\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\7\17\u00bf\n\17\f\17\16"+
		"\17\u00c2\13\17\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20"+
		"\3\20\3\20\3\20\3\20\5\20\u00d3\n\20\3\21\3\21\3\21\3\21\3\21\3\21\3\21"+
		"\3\21\3\21\5\21\u00de\n\21\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22"+
		"\3\22\3\22\3\22\3\22\3\22\5\22\u00ee\n\22\3\23\3\23\3\23\3\23\3\23\3\23"+
		"\3\23\3\23\3\23\7\23\u00f9\n\23\f\23\16\23\u00fc\13\23\3\24\3\24\3\24"+
		"\3\24\3\24\3\24\3\24\3\24\5\24\u0106\n\24\3\24\3\24\3\24\3\24\3\24\3\24"+
		"\3\24\3\24\3\24\3\24\7\24\u0112\n\24\f\24\16\24\u0115\13\24\3\25\3\25"+
		"\3\25\3\25\3\25\3\25\3\25\3\25\3\25\3\25\3\25\3\25\3\25\3\25\5\25\u0125"+
		"\n\25\3\25\2\7\n\f\34$&\26\2\4\6\b\n\f\16\20\22\24\26\30\32\34\36 \"$"+
		"&(\2\7\3\2\13\f\3\2\22\23\3\2$)\3\2!#\3\2\37 \2\u012d\2*\3\2\2\2\4\60"+
		"\3\2\2\2\6\65\3\2\2\2\b8\3\2\2\2\nC\3\2\2\2\fP\3\2\2\2\16s\3\2\2\2\20"+
		"u\3\2\2\2\22~\3\2\2\2\24\u008d\3\2\2\2\26\u00a5\3\2\2\2\30\u00ae\3\2\2"+
		"\2\32\u00b0\3\2\2\2\34\u00b6\3\2\2\2\36\u00d2\3\2\2\2 \u00dd\3\2\2\2\""+
		"\u00ed\3\2\2\2$\u00ef\3\2\2\2&\u0105\3\2\2\2(\u0124\3\2\2\2*+\5\4\3\2"+
		"+,\b\2\1\2,\3\3\2\2\2-/\5\6\4\2.-\3\2\2\2/\62\3\2\2\2\60.\3\2\2\2\60\61"+
		"\3\2\2\2\61\63\3\2\2\2\62\60\3\2\2\2\63\64\b\3\1\2\64\5\3\2\2\2\65\66"+
		"\5\b\5\2\66\67\b\4\1\2\67\7\3\2\2\289\7\3\2\29:\7\4\2\2:;\7\5\2\2;<\7"+
		"\6\2\2<=\7\7\2\2=>\7\25\2\2>?\5\n\6\2?@\7\60\2\2@A\5\f\7\2AB\b\5\1\2B"+
		"\t\3\2\2\2CD\b\6\1\2DE\7\34\2\2EF\b\6\1\2FM\3\2\2\2GH\f\4\2\2HI\7\62\2"+
		"\2IJ\7\34\2\2JL\b\6\1\2KG\3\2\2\2LO\3\2\2\2MK\3\2\2\2MN\3\2\2\2N\13\3"+
		"\2\2\2OM\3\2\2\2PQ\b\7\1\2QR\5\16\b\2RS\b\7\1\2SZ\3\2\2\2TU\f\4\2\2UV"+
		"\5\16\b\2VW\b\7\1\2WY\3\2\2\2XT\3\2\2\2Y\\\3\2\2\2ZX\3\2\2\2Z[\3\2\2\2"+
		"[\r\3\2\2\2\\Z\3\2\2\2]^\7\b\2\2^_\7\36\2\2_`\7.\2\2`a\7/\2\2ab\7*\2\2"+
		"bc\5\34\17\2cd\7\n\2\2de\7+\2\2ef\b\b\1\2ft\3\2\2\2gh\7\26\2\2hi\7\20"+
		"\2\2ij\7.\2\2jk\7/\2\2kl\7*\2\2lm\7\27\2\2mn\7\30\2\2no\5\34\17\2op\7"+
		"\21\2\2pq\7+\2\2qr\b\b\1\2rt\3\2\2\2s]\3\2\2\2sg\3\2\2\2t\17\3\2\2\2u"+
		"v\7\24\2\2vw\7.\2\2wx\5 \21\2xy\7/\2\2yz\7\t\2\2z{\7\35\2\2{|\7\60\2\2"+
		"|}\b\t\1\2}\21\3\2\2\2~\177\7\17\2\2\177\u0080\7.\2\2\u0080\u0081\7\33"+
		"\2\2\u0081\u0082\7\62\2\2\u0082\u0083\5\24\13\2\u0083\u0084\5 \21\2\u0084"+
		"\u0085\7/\2\2\u0085\u0086\7\60\2\2\u0086\u0087\b\n\1\2\u0087\23\3\2\2"+
		"\2\u0088\u0089\7\16\2\2\u0089\u008e\b\13\1\2\u008a\u008b\7\r\2\2\u008b"+
		"\u008e\b\13\1\2\u008c\u008e\b\13\1\2\u008d\u0088\3\2\2\2\u008d\u008a\3"+
		"\2\2\2\u008d\u008c\3\2\2\2\u008e\25\3\2\2\2\u008f\u0090\7\34\2\2\u0090"+
		"\u0091\7\63\2\2\u0091\u0092\5 \21\2\u0092\u0093\7\60\2\2\u0093\u0094\b"+
		"\f\1\2\u0094\u00a6\3\2\2\2\u0095\u0096\t\2\2\2\u0096\u0097\7,\2\2\u0097"+
		"\u0098\7\r\2\2\u0098\u0099\5 \21\2\u0099\u009a\7-\2\2\u009a\u009b\7\63"+
		"\2\2\u009b\u009c\5 \21\2\u009c\u009d\7\60\2\2\u009d\u009e\b\f\1\2\u009e"+
		"\u00a6\3\2\2\2\u009f\u00a0\t\3\2\2\u00a0\u00a1\7\63\2\2\u00a1\u00a2\5"+
		" \21\2\u00a2\u00a3\7\60\2\2\u00a3\u00a4\b\f\1\2\u00a4\u00a6\3\2\2\2\u00a5"+
		"\u008f\3\2\2\2\u00a5\u0095\3\2\2\2\u00a5\u009f\3\2\2\2\u00a6\27\3\2\2"+
		"\2\u00a7\u00a8\7\35\2\2\u00a8\u00a9\7\61\2\2\u00a9\u00af\b\r\1\2\u00aa"+
		"\u00ab\7\t\2\2\u00ab\u00ac\7\35\2\2\u00ac\u00ad\7\60\2\2\u00ad\u00af\b"+
		"\r\1\2\u00ae\u00a7\3\2\2\2\u00ae\u00aa\3\2\2\2\u00af\31\3\2\2\2\u00b0"+
		"\u00b1\7\36\2\2\u00b1\u00b2\7.\2\2\u00b2\u00b3\7/\2\2\u00b3\u00b4\7\60"+
		"\2\2\u00b4\u00b5\b\16\1\2\u00b5\33\3\2\2\2\u00b6\u00b7\b\17\1\2\u00b7"+
		"\u00b8\5\36\20\2\u00b8\u00b9\b\17\1\2\u00b9\u00c0\3\2\2\2\u00ba\u00bb"+
		"\f\4\2\2\u00bb\u00bc\5\36\20\2\u00bc\u00bd\b\17\1\2\u00bd\u00bf\3\2\2"+
		"\2\u00be\u00ba\3\2\2\2\u00bf\u00c2\3\2\2\2\u00c0\u00be\3\2\2\2\u00c0\u00c1"+
		"\3\2\2\2\u00c1\35\3\2\2\2\u00c2\u00c0\3\2\2\2\u00c3\u00c4\5\22\n\2\u00c4"+
		"\u00c5\b\20\1\2\u00c5\u00d3\3\2\2\2\u00c6\u00c7\5\20\t\2\u00c7\u00c8\b"+
		"\20\1\2\u00c8\u00d3\3\2\2\2\u00c9\u00ca\5\26\f\2\u00ca\u00cb\b\20\1\2"+
		"\u00cb\u00d3\3\2\2\2\u00cc\u00cd\5\30\r\2\u00cd\u00ce\b\20\1\2\u00ce\u00d3"+
		"\3\2\2\2\u00cf\u00d0\5\32\16\2\u00d0\u00d1\b\20\1\2\u00d1\u00d3\3\2\2"+
		"\2\u00d2\u00c3\3\2\2\2\u00d2\u00c6\3\2\2\2\u00d2\u00c9\3\2\2\2\u00d2\u00cc"+
		"\3\2\2\2\u00d2\u00cf\3\2\2\2\u00d3\37\3\2\2\2\u00d4\u00d5\5$\23\2\u00d5"+
		"\u00d6\b\21\1\2\u00d6\u00de\3\2\2\2\u00d7\u00d8\5&\24\2\u00d8\u00d9\b"+
		"\21\1\2\u00d9\u00de\3\2\2\2\u00da\u00db\5\"\22\2\u00db\u00dc\b\21\1\2"+
		"\u00dc\u00de\3\2\2\2\u00dd\u00d4\3\2\2\2\u00dd\u00d7\3\2\2\2\u00dd\u00da"+
		"\3\2\2\2\u00de!\3\2\2\2\u00df\u00e0\7\13\2\2\u00e0\u00e1\7,\2\2\u00e1"+
		"\u00e2\7\r\2\2\u00e2\u00e3\5 \21\2\u00e3\u00e4\7-\2\2\u00e4\u00e5\b\22"+
		"\1\2\u00e5\u00ee\3\2\2\2\u00e6\u00e7\7\f\2\2\u00e7\u00e8\7,\2\2\u00e8"+
		"\u00e9\7\r\2\2\u00e9\u00ea\5 \21\2\u00ea\u00eb\7-\2\2\u00eb\u00ec\b\22"+
		"\1\2\u00ec\u00ee\3\2\2\2\u00ed\u00df\3\2\2\2\u00ed\u00e6\3\2\2\2\u00ee"+
		"#\3\2\2\2\u00ef\u00f0\b\23\1\2\u00f0\u00f1\5&\24\2\u00f1\u00f2\b\23\1"+
		"\2\u00f2\u00fa\3\2\2\2\u00f3\u00f4\f\4\2\2\u00f4\u00f5\t\4\2\2\u00f5\u00f6"+
		"\5$\23\5\u00f6\u00f7\b\23\1\2\u00f7\u00f9\3\2\2\2\u00f8\u00f3\3\2\2\2"+
		"\u00f9\u00fc\3\2\2\2\u00fa\u00f8\3\2\2\2\u00fa\u00fb\3\2\2\2\u00fb%\3"+
		"\2\2\2\u00fc\u00fa\3\2\2\2\u00fd\u00fe\b\24\1\2\u00fe\u00ff\7 \2\2\u00ff"+
		"\u0100\5&\24\6\u0100\u0101\b\24\1\2\u0101\u0106\3\2\2\2\u0102\u0103\5"+
		"(\25\2\u0103\u0104\b\24\1\2\u0104\u0106\3\2\2\2\u0105\u00fd\3\2\2\2\u0105"+
		"\u0102\3\2\2\2\u0106\u0113\3\2\2\2\u0107\u0108\f\5\2\2\u0108\u0109\t\5"+
		"\2\2\u0109\u010a\5&\24\6\u010a\u010b\b\24\1\2\u010b\u0112\3\2\2\2\u010c"+
		"\u010d\f\4\2\2\u010d\u010e\t\6\2\2\u010e\u010f\5&\24\5\u010f\u0110\b\24"+
		"\1\2\u0110\u0112\3\2\2\2\u0111\u0107\3\2\2\2\u0111\u010c\3\2\2\2\u0112"+
		"\u0115\3\2\2\2\u0113\u0111\3\2\2\2\u0113\u0114\3\2\2\2\u0114\'\3\2\2\2"+
		"\u0115\u0113\3\2\2\2\u0116\u0117\7\32\2\2\u0117\u0125\b\25\1\2\u0118\u0119"+
		"\7\31\2\2\u0119\u0125\b\25\1\2\u011a\u011b\7\36\2\2\u011b\u0125\b\25\1"+
		"\2\u011c\u011d\7\33\2\2\u011d\u0125\b\25\1\2\u011e\u011f\7\34\2\2\u011f"+
		"\u0125\b\25\1\2\u0120\u0121\7\22\2\2\u0121\u0125\b\25\1\2\u0122\u0123"+
		"\7\23\2\2\u0123\u0125\b\25\1\2\u0124\u0116\3\2\2\2\u0124\u0118\3\2\2\2"+
		"\u0124\u011a\3\2\2\2\u0124\u011c\3\2\2\2\u0124\u011e\3\2\2\2\u0124\u0120"+
		"\3\2\2\2\u0124\u0122\3\2\2\2\u0125)\3\2\2\2\22\60MZs\u008d\u00a5\u00ae"+
		"\u00c0\u00d2\u00dd\u00ed\u00fa\u0105\u0111\u0113\u0124";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}