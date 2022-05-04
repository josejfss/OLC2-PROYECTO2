// Generated from /home/iovana/go/src/OLC2-PROYECTO2/OPT_COMPILADOR/optimizador_lexer.g4 by ANTLR 4.8
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class optimizador_lexer extends Lexer {
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
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"TK_STDIOH", "TK_DHEAP", "TK_DSTACK", "TK_PSTAKC", "TK_PHEAP", "TK_VOID", 
			"TK_GOTO", "TK_RETURN", "TK_HEAP", "TK_STACK", "TK_CASTINT", "TK_CASTCHAR", 
			"TK_PRINTF", "TK_MAIN", "TK_RETMAIN", "TK_PUNSTACK", "TK_PUNHEAP", "TK_IF", 
			"TK_DOUBLE", "TK_INT", "TK_ASIGPSTACK", "TK_ASIGHEAP", "TK_FLOAT", "TK_ENTERO", 
			"TK_CADENA", "TK_TEMPORAL", "TK_ETIQUETA", "TK_IDENTIFICADOR", "TK_SUMA", 
			"TK_RESTA", "TK_MULTI", "TK_DIVI", "TK_MODULO", "TK_MENORIGUAL", "TK_MAYORIGUAL", 
			"TK_IGUALDAD", "TK_DESIGUALDAD", "TK_MENOR", "TK_MAYOR", "TK_LI", "TK_LD", 
			"TK_CI", "TK_CD", "TK_PI", "TK_PD", "TK_PYC", "TK_DP", "TK_COMA", "TK_IGUAL", 
			"COMENTARIO_MUL", "COMENTARIO_LIN", "WHITESPACE", "ESC_SEQ"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'#include <stdio.h>'", "'double HEAP[100000];'", "'double STACK[100000];'", 
			"'double SP;'", "'double HP;'", "'void'", "'goto'", "'return;'", "'HEAP'", 
			"'STACK'", "'(int)'", "'(char)'", "'printf'", "'main'", "'return 0;'", 
			"'SP'", "'HP'", "'if'", "'double'", "'int'", "'SP=0;'", "'HP=0;'", null, 
			null, null, null, null, null, "'+'", "'-'", "'*'", "'/'", "'%'", "'<='", 
			"'>='", "'=='", "'!='", "'<'", "'>'", "'{'", "'}'", "'['", "']'", "'('", 
			"')'", "';'", "':'", "','", "'='"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "TK_STDIOH", "TK_DHEAP", "TK_DSTACK", "TK_PSTAKC", "TK_PHEAP", 
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


	public optimizador_lexer(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "optimizador_lexer.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public String[] getChannelNames() { return channelNames; }

	@Override
	public String[] getModeNames() { return modeNames; }

	@Override
	public ATN getATN() { return _ATN; }

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2\66\u019c\b\1\4\2"+
		"\t\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4"+
		"\13\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22"+
		"\t\22\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31"+
		"\t\31\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t"+
		" \4!\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\4)\t)\4*\t*\4+\t"+
		"+\4,\t,\4-\t-\4.\t.\4/\t/\4\60\t\60\4\61\t\61\4\62\t\62\4\63\t\63\4\64"+
		"\t\64\4\65\t\65\4\66\t\66\3\2\3\2\3\2\3\2\3\2\3\2\3\2\3\2\3\2\3\2\3\2"+
		"\3\2\3\2\3\2\3\2\3\2\3\2\3\2\3\2\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3"+
		"\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\4\3\4\3\4\3\4\3\4\3\4"+
		"\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\5\3"+
		"\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6"+
		"\3\6\3\6\3\6\3\7\3\7\3\7\3\7\3\7\3\b\3\b\3\b\3\b\3\b\3\t\3\t\3\t\3\t\3"+
		"\t\3\t\3\t\3\t\3\n\3\n\3\n\3\n\3\n\3\13\3\13\3\13\3\13\3\13\3\13\3\f\3"+
		"\f\3\f\3\f\3\f\3\f\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\16\3\16\3\16\3\16\3\16"+
		"\3\16\3\16\3\17\3\17\3\17\3\17\3\17\3\20\3\20\3\20\3\20\3\20\3\20\3\20"+
		"\3\20\3\20\3\20\3\21\3\21\3\21\3\22\3\22\3\22\3\23\3\23\3\23\3\24\3\24"+
		"\3\24\3\24\3\24\3\24\3\24\3\25\3\25\3\25\3\25\3\26\3\26\3\26\3\26\3\26"+
		"\3\26\3\27\3\27\3\27\3\27\3\27\3\27\3\30\6\30\u0123\n\30\r\30\16\30\u0124"+
		"\3\30\3\30\6\30\u0129\n\30\r\30\16\30\u012a\3\31\6\31\u012e\n\31\r\31"+
		"\16\31\u012f\3\32\3\32\7\32\u0134\n\32\f\32\16\32\u0137\13\32\3\32\3\32"+
		"\3\33\3\33\6\33\u013d\n\33\r\33\16\33\u013e\3\34\3\34\6\34\u0143\n\34"+
		"\r\34\16\34\u0144\3\35\3\35\7\35\u0149\n\35\f\35\16\35\u014c\13\35\3\36"+
		"\3\36\3\37\3\37\3 \3 \3!\3!\3\"\3\"\3#\3#\3#\3$\3$\3$\3%\3%\3%\3&\3&\3"+
		"&\3\'\3\'\3(\3(\3)\3)\3*\3*\3+\3+\3,\3,\3-\3-\3.\3.\3/\3/\3\60\3\60\3"+
		"\61\3\61\3\62\3\62\3\63\3\63\3\63\3\63\6\63\u0180\n\63\r\63\16\63\u0181"+
		"\3\63\3\63\3\63\3\63\3\63\3\64\3\64\3\64\3\64\6\64\u018d\n\64\r\64\16"+
		"\64\u018e\3\64\3\64\3\65\6\65\u0194\n\65\r\65\16\65\u0195\3\65\3\65\3"+
		"\66\3\66\3\66\2\2\67\3\3\5\4\7\5\t\6\13\7\r\b\17\t\21\n\23\13\25\f\27"+
		"\r\31\16\33\17\35\20\37\21!\22#\23%\24\'\25)\26+\27-\30/\31\61\32\63\33"+
		"\65\34\67\359\36;\37= ?!A\"C#E$G%I&K\'M(O)Q*S+U,W-Y.[/]\60_\61a\62c\63"+
		"e\64g\65i\66k\2\3\2\n\3\2\62;\3\2$$\7\2C\\aac|\u00d3\u00d3\u00f3\u00f3"+
		"\b\2\62;C\\aac|\u00d3\u00d3\u00f3\u00f3\3\2+,\3\2\f\f\6\2\13\f\17\17\""+
		"\"^^\t\2\"#%%--/\60<<BB]_\2\u01a4\2\3\3\2\2\2\2\5\3\2\2\2\2\7\3\2\2\2"+
		"\2\t\3\2\2\2\2\13\3\2\2\2\2\r\3\2\2\2\2\17\3\2\2\2\2\21\3\2\2\2\2\23\3"+
		"\2\2\2\2\25\3\2\2\2\2\27\3\2\2\2\2\31\3\2\2\2\2\33\3\2\2\2\2\35\3\2\2"+
		"\2\2\37\3\2\2\2\2!\3\2\2\2\2#\3\2\2\2\2%\3\2\2\2\2\'\3\2\2\2\2)\3\2\2"+
		"\2\2+\3\2\2\2\2-\3\2\2\2\2/\3\2\2\2\2\61\3\2\2\2\2\63\3\2\2\2\2\65\3\2"+
		"\2\2\2\67\3\2\2\2\29\3\2\2\2\2;\3\2\2\2\2=\3\2\2\2\2?\3\2\2\2\2A\3\2\2"+
		"\2\2C\3\2\2\2\2E\3\2\2\2\2G\3\2\2\2\2I\3\2\2\2\2K\3\2\2\2\2M\3\2\2\2\2"+
		"O\3\2\2\2\2Q\3\2\2\2\2S\3\2\2\2\2U\3\2\2\2\2W\3\2\2\2\2Y\3\2\2\2\2[\3"+
		"\2\2\2\2]\3\2\2\2\2_\3\2\2\2\2a\3\2\2\2\2c\3\2\2\2\2e\3\2\2\2\2g\3\2\2"+
		"\2\2i\3\2\2\2\3m\3\2\2\2\5\u0080\3\2\2\2\7\u0095\3\2\2\2\t\u00ab\3\2\2"+
		"\2\13\u00b6\3\2\2\2\r\u00c1\3\2\2\2\17\u00c6\3\2\2\2\21\u00cb\3\2\2\2"+
		"\23\u00d3\3\2\2\2\25\u00d8\3\2\2\2\27\u00de\3\2\2\2\31\u00e4\3\2\2\2\33"+
		"\u00eb\3\2\2\2\35\u00f2\3\2\2\2\37\u00f7\3\2\2\2!\u0101\3\2\2\2#\u0104"+
		"\3\2\2\2%\u0107\3\2\2\2\'\u010a\3\2\2\2)\u0111\3\2\2\2+\u0115\3\2\2\2"+
		"-\u011b\3\2\2\2/\u0122\3\2\2\2\61\u012d\3\2\2\2\63\u0131\3\2\2\2\65\u013a"+
		"\3\2\2\2\67\u0140\3\2\2\29\u0146\3\2\2\2;\u014d\3\2\2\2=\u014f\3\2\2\2"+
		"?\u0151\3\2\2\2A\u0153\3\2\2\2C\u0155\3\2\2\2E\u0157\3\2\2\2G\u015a\3"+
		"\2\2\2I\u015d\3\2\2\2K\u0160\3\2\2\2M\u0163\3\2\2\2O\u0165\3\2\2\2Q\u0167"+
		"\3\2\2\2S\u0169\3\2\2\2U\u016b\3\2\2\2W\u016d\3\2\2\2Y\u016f\3\2\2\2["+
		"\u0171\3\2\2\2]\u0173\3\2\2\2_\u0175\3\2\2\2a\u0177\3\2\2\2c\u0179\3\2"+
		"\2\2e\u017b\3\2\2\2g\u0188\3\2\2\2i\u0193\3\2\2\2k\u0199\3\2\2\2mn\7%"+
		"\2\2no\7k\2\2op\7p\2\2pq\7e\2\2qr\7n\2\2rs\7w\2\2st\7f\2\2tu\7g\2\2uv"+
		"\7\"\2\2vw\7>\2\2wx\7u\2\2xy\7v\2\2yz\7f\2\2z{\7k\2\2{|\7q\2\2|}\7\60"+
		"\2\2}~\7j\2\2~\177\7@\2\2\177\4\3\2\2\2\u0080\u0081\7f\2\2\u0081\u0082"+
		"\7q\2\2\u0082\u0083\7w\2\2\u0083\u0084\7d\2\2\u0084\u0085\7n\2\2\u0085"+
		"\u0086\7g\2\2\u0086\u0087\7\"\2\2\u0087\u0088\7J\2\2\u0088\u0089\7G\2"+
		"\2\u0089\u008a\7C\2\2\u008a\u008b\7R\2\2\u008b\u008c\7]\2\2\u008c\u008d"+
		"\7\63\2\2\u008d\u008e\7\62\2\2\u008e\u008f\7\62\2\2\u008f\u0090\7\62\2"+
		"\2\u0090\u0091\7\62\2\2\u0091\u0092\7\62\2\2\u0092\u0093\7_\2\2\u0093"+
		"\u0094\7=\2\2\u0094\6\3\2\2\2\u0095\u0096\7f\2\2\u0096\u0097\7q\2\2\u0097"+
		"\u0098\7w\2\2\u0098\u0099\7d\2\2\u0099\u009a\7n\2\2\u009a\u009b\7g\2\2"+
		"\u009b\u009c\7\"\2\2\u009c\u009d\7U\2\2\u009d\u009e\7V\2\2\u009e\u009f"+
		"\7C\2\2\u009f\u00a0\7E\2\2\u00a0\u00a1\7M\2\2\u00a1\u00a2\7]\2\2\u00a2"+
		"\u00a3\7\63\2\2\u00a3\u00a4\7\62\2\2\u00a4\u00a5\7\62\2\2\u00a5\u00a6"+
		"\7\62\2\2\u00a6\u00a7\7\62\2\2\u00a7\u00a8\7\62\2\2\u00a8\u00a9\7_\2\2"+
		"\u00a9\u00aa\7=\2\2\u00aa\b\3\2\2\2\u00ab\u00ac\7f\2\2\u00ac\u00ad\7q"+
		"\2\2\u00ad\u00ae\7w\2\2\u00ae\u00af\7d\2\2\u00af\u00b0\7n\2\2\u00b0\u00b1"+
		"\7g\2\2\u00b1\u00b2\7\"\2\2\u00b2\u00b3\7U\2\2\u00b3\u00b4\7R\2\2\u00b4"+
		"\u00b5\7=\2\2\u00b5\n\3\2\2\2\u00b6\u00b7\7f\2\2\u00b7\u00b8\7q\2\2\u00b8"+
		"\u00b9\7w\2\2\u00b9\u00ba\7d\2\2\u00ba\u00bb\7n\2\2\u00bb\u00bc\7g\2\2"+
		"\u00bc\u00bd\7\"\2\2\u00bd\u00be\7J\2\2\u00be\u00bf\7R\2\2\u00bf\u00c0"+
		"\7=\2\2\u00c0\f\3\2\2\2\u00c1\u00c2\7x\2\2\u00c2\u00c3\7q\2\2\u00c3\u00c4"+
		"\7k\2\2\u00c4\u00c5\7f\2\2\u00c5\16\3\2\2\2\u00c6\u00c7\7i\2\2\u00c7\u00c8"+
		"\7q\2\2\u00c8\u00c9\7v\2\2\u00c9\u00ca\7q\2\2\u00ca\20\3\2\2\2\u00cb\u00cc"+
		"\7t\2\2\u00cc\u00cd\7g\2\2\u00cd\u00ce\7v\2\2\u00ce\u00cf\7w\2\2\u00cf"+
		"\u00d0\7t\2\2\u00d0\u00d1\7p\2\2\u00d1\u00d2\7=\2\2\u00d2\22\3\2\2\2\u00d3"+
		"\u00d4\7J\2\2\u00d4\u00d5\7G\2\2\u00d5\u00d6\7C\2\2\u00d6\u00d7\7R\2\2"+
		"\u00d7\24\3\2\2\2\u00d8\u00d9\7U\2\2\u00d9\u00da\7V\2\2\u00da\u00db\7"+
		"C\2\2\u00db\u00dc\7E\2\2\u00dc\u00dd\7M\2\2\u00dd\26\3\2\2\2\u00de\u00df"+
		"\7*\2\2\u00df\u00e0\7k\2\2\u00e0\u00e1\7p\2\2\u00e1\u00e2\7v\2\2\u00e2"+
		"\u00e3\7+\2\2\u00e3\30\3\2\2\2\u00e4\u00e5\7*\2\2\u00e5\u00e6\7e\2\2\u00e6"+
		"\u00e7\7j\2\2\u00e7\u00e8\7c\2\2\u00e8\u00e9\7t\2\2\u00e9\u00ea\7+\2\2"+
		"\u00ea\32\3\2\2\2\u00eb\u00ec\7r\2\2\u00ec\u00ed\7t\2\2\u00ed\u00ee\7"+
		"k\2\2\u00ee\u00ef\7p\2\2\u00ef\u00f0\7v\2\2\u00f0\u00f1\7h\2\2\u00f1\34"+
		"\3\2\2\2\u00f2\u00f3\7o\2\2\u00f3\u00f4\7c\2\2\u00f4\u00f5\7k\2\2\u00f5"+
		"\u00f6\7p\2\2\u00f6\36\3\2\2\2\u00f7\u00f8\7t\2\2\u00f8\u00f9\7g\2\2\u00f9"+
		"\u00fa\7v\2\2\u00fa\u00fb\7w\2\2\u00fb\u00fc\7t\2\2\u00fc\u00fd\7p\2\2"+
		"\u00fd\u00fe\7\"\2\2\u00fe\u00ff\7\62\2\2\u00ff\u0100\7=\2\2\u0100 \3"+
		"\2\2\2\u0101\u0102\7U\2\2\u0102\u0103\7R\2\2\u0103\"\3\2\2\2\u0104\u0105"+
		"\7J\2\2\u0105\u0106\7R\2\2\u0106$\3\2\2\2\u0107\u0108\7k\2\2\u0108\u0109"+
		"\7h\2\2\u0109&\3\2\2\2\u010a\u010b\7f\2\2\u010b\u010c\7q\2\2\u010c\u010d"+
		"\7w\2\2\u010d\u010e\7d\2\2\u010e\u010f\7n\2\2\u010f\u0110\7g\2\2\u0110"+
		"(\3\2\2\2\u0111\u0112\7k\2\2\u0112\u0113\7p\2\2\u0113\u0114\7v\2\2\u0114"+
		"*\3\2\2\2\u0115\u0116\7U\2\2\u0116\u0117\7R\2\2\u0117\u0118\7?\2\2\u0118"+
		"\u0119\7\62\2\2\u0119\u011a\7=\2\2\u011a,\3\2\2\2\u011b\u011c\7J\2\2\u011c"+
		"\u011d\7R\2\2\u011d\u011e\7?\2\2\u011e\u011f\7\62\2\2\u011f\u0120\7=\2"+
		"\2\u0120.\3\2\2\2\u0121\u0123\t\2\2\2\u0122\u0121\3\2\2\2\u0123\u0124"+
		"\3\2\2\2\u0124\u0122\3\2\2\2\u0124\u0125\3\2\2\2\u0125\u0126\3\2\2\2\u0126"+
		"\u0128\7\60\2\2\u0127\u0129\t\2\2\2\u0128\u0127\3\2\2\2\u0129\u012a\3"+
		"\2\2\2\u012a\u0128\3\2\2\2\u012a\u012b\3\2\2\2\u012b\60\3\2\2\2\u012c"+
		"\u012e\t\2\2\2\u012d\u012c\3\2\2\2\u012e\u012f\3\2\2\2\u012f\u012d\3\2"+
		"\2\2\u012f\u0130\3\2\2\2\u0130\62\3\2\2\2\u0131\u0135\7$\2\2\u0132\u0134"+
		"\n\3\2\2\u0133\u0132\3\2\2\2\u0134\u0137\3\2\2\2\u0135\u0133\3\2\2\2\u0135"+
		"\u0136\3\2\2\2\u0136\u0138\3\2\2\2\u0137\u0135\3\2\2\2\u0138\u0139\7$"+
		"\2\2\u0139\64\3\2\2\2\u013a\u013c\7V\2\2\u013b\u013d\t\2\2\2\u013c\u013b"+
		"\3\2\2\2\u013d\u013e\3\2\2\2\u013e\u013c\3\2\2\2\u013e\u013f\3\2\2\2\u013f"+
		"\66\3\2\2\2\u0140\u0142\7N\2\2\u0141\u0143\t\2\2\2\u0142\u0141\3\2\2\2"+
		"\u0143\u0144\3\2\2\2\u0144\u0142\3\2\2\2\u0144\u0145\3\2\2\2\u01458\3"+
		"\2\2\2\u0146\u014a\t\4\2\2\u0147\u0149\t\5\2\2\u0148\u0147\3\2\2\2\u0149"+
		"\u014c\3\2\2\2\u014a\u0148\3\2\2\2\u014a\u014b\3\2\2\2\u014b:\3\2\2\2"+
		"\u014c\u014a\3\2\2\2\u014d\u014e\7-\2\2\u014e<\3\2\2\2\u014f\u0150\7/"+
		"\2\2\u0150>\3\2\2\2\u0151\u0152\7,\2\2\u0152@\3\2\2\2\u0153\u0154\7\61"+
		"\2\2\u0154B\3\2\2\2\u0155\u0156\7\'\2\2\u0156D\3\2\2\2\u0157\u0158\7>"+
		"\2\2\u0158\u0159\7?\2\2\u0159F\3\2\2\2\u015a\u015b\7@\2\2\u015b\u015c"+
		"\7?\2\2\u015cH\3\2\2\2\u015d\u015e\7?\2\2\u015e\u015f\7?\2\2\u015fJ\3"+
		"\2\2\2\u0160\u0161\7#\2\2\u0161\u0162\7?\2\2\u0162L\3\2\2\2\u0163\u0164"+
		"\7>\2\2\u0164N\3\2\2\2\u0165\u0166\7@\2\2\u0166P\3\2\2\2\u0167\u0168\7"+
		"}\2\2\u0168R\3\2\2\2\u0169\u016a\7\177\2\2\u016aT\3\2\2\2\u016b\u016c"+
		"\7]\2\2\u016cV\3\2\2\2\u016d\u016e\7_\2\2\u016eX\3\2\2\2\u016f\u0170\7"+
		"*\2\2\u0170Z\3\2\2\2\u0171\u0172\7+\2\2\u0172\\\3\2\2\2\u0173\u0174\7"+
		"=\2\2\u0174^\3\2\2\2\u0175\u0176\7<\2\2\u0176`\3\2\2\2\u0177\u0178\7."+
		"\2\2\u0178b\3\2\2\2\u0179\u017a\7?\2\2\u017ad\3\2\2\2\u017b\u017c\7\61"+
		"\2\2\u017c\u017d\7,\2\2\u017d\u017f\3\2\2\2\u017e\u0180\n\6\2\2\u017f"+
		"\u017e\3\2\2\2\u0180\u0181\3\2\2\2\u0181\u017f\3\2\2\2\u0181\u0182\3\2"+
		"\2\2\u0182\u0183\3\2\2\2\u0183\u0184\7,\2\2\u0184\u0185\7\61\2\2\u0185"+
		"\u0186\3\2\2\2\u0186\u0187\b\63\2\2\u0187f\3\2\2\2\u0188\u0189\7\61\2"+
		"\2\u0189\u018a\7\61\2\2\u018a\u018c\3\2\2\2\u018b\u018d\n\7\2\2\u018c"+
		"\u018b\3\2\2\2\u018d\u018e\3\2\2\2\u018e\u018c\3\2\2\2\u018e\u018f\3\2"+
		"\2\2\u018f\u0190\3\2\2\2\u0190\u0191\b\64\2\2\u0191h\3\2\2\2\u0192\u0194"+
		"\t\b\2\2\u0193\u0192\3\2\2\2\u0194\u0195\3\2\2\2\u0195\u0193\3\2\2\2\u0195"+
		"\u0196\3\2\2\2\u0196\u0197\3\2\2\2\u0197\u0198\b\65\2\2\u0198j\3\2\2\2"+
		"\u0199\u019a\7^\2\2\u019a\u019b\t\t\2\2\u019bl\3\2\2\2\r\2\u0124\u012a"+
		"\u012f\u0135\u013e\u0144\u014a\u0181\u018e\u0195\3\b\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}