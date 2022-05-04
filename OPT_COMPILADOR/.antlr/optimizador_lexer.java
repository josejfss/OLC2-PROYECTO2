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
		TK_IF=18, TK_DOUBLE=19, TK_INT=20, TK_FLOAT=21, TK_ENTERO=22, TK_CADENA=23, 
		TK_TEMPORAL=24, TK_ETIQUETA=25, TK_IDENTIFICADOR=26, TK_SUMA=27, TK_RESTA=28, 
		TK_MULTI=29, TK_DIVI=30, TK_MODULO=31, TK_MENORIGUAL=32, TK_MAYORIGUAL=33, 
		TK_IGUALDAD=34, TK_DESIGUALDAD=35, TK_MENOR=36, TK_MAYOR=37, TK_LI=38, 
		TK_LD=39, TK_CI=40, TK_CD=41, TK_PI=42, TK_PD=43, TK_PYC=44, TK_DP=45, 
		TK_COMA=46, TK_IGUAL=47, COMENTARIO_MUL=48, COMENTARIO_LIN=49, WHITESPACE=50;
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
			"TK_DOUBLE", "TK_INT", "TK_FLOAT", "TK_ENTERO", "TK_CADENA", "TK_TEMPORAL", 
			"TK_ETIQUETA", "TK_IDENTIFICADOR", "TK_SUMA", "TK_RESTA", "TK_MULTI", 
			"TK_DIVI", "TK_MODULO", "TK_MENORIGUAL", "TK_MAYORIGUAL", "TK_IGUALDAD", 
			"TK_DESIGUALDAD", "TK_MENOR", "TK_MAYOR", "TK_LI", "TK_LD", "TK_CI", 
			"TK_CD", "TK_PI", "TK_PD", "TK_PYC", "TK_DP", "TK_COMA", "TK_IGUAL", 
			"COMENTARIO_MUL", "COMENTARIO_LIN", "WHITESPACE", "ESC_SEQ"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'#include <stdio.h>'", "'double HEAP[100000];'", "'double STACK[100000];'", 
			"'double SP;'", "'double HP;'", "'void'", "'goto'", "'return;'", "'HEAP'", 
			"'STACK'", "'(int)'", "'(char)'", "'printf'", "'main'", "'return 0;'", 
			"'SP'", "'HP'", "'if'", "'double'", "'int'", null, null, null, null, 
			null, null, "'+'", "'-'", "'*'", "'/'", "'%'", "'<='", "'>='", "'=='", 
			"'!='", "'<'", "'>'", "'{'", "'}'", "'['", "']'", "'('", "')'", "';'", 
			"':'", "','", "'='"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "TK_STDIOH", "TK_DHEAP", "TK_DSTACK", "TK_PSTAKC", "TK_PHEAP", 
			"TK_VOID", "TK_GOTO", "TK_RETURN", "TK_HEAP", "TK_STACK", "TK_CASTINT", 
			"TK_CASTCHAR", "TK_PRINTF", "TK_MAIN", "TK_RETMAIN", "TK_PUNSTACK", "TK_PUNHEAP", 
			"TK_IF", "TK_DOUBLE", "TK_INT", "TK_FLOAT", "TK_ENTERO", "TK_CADENA", 
			"TK_TEMPORAL", "TK_ETIQUETA", "TK_IDENTIFICADOR", "TK_SUMA", "TK_RESTA", 
			"TK_MULTI", "TK_DIVI", "TK_MODULO", "TK_MENORIGUAL", "TK_MAYORIGUAL", 
			"TK_IGUALDAD", "TK_DESIGUALDAD", "TK_MENOR", "TK_MAYOR", "TK_LI", "TK_LD", 
			"TK_CI", "TK_CD", "TK_PI", "TK_PD", "TK_PYC", "TK_DP", "TK_COMA", "TK_IGUAL", 
			"COMENTARIO_MUL", "COMENTARIO_LIN", "WHITESPACE"
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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2\64\u018c\b\1\4\2"+
		"\t\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4"+
		"\13\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22"+
		"\t\22\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31"+
		"\t\31\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t"+
		" \4!\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\4)\t)\4*\t*\4+\t"+
		"+\4,\t,\4-\t-\4.\t.\4/\t/\4\60\t\60\4\61\t\61\4\62\t\62\4\63\t\63\4\64"+
		"\t\64\3\2\3\2\3\2\3\2\3\2\3\2\3\2\3\2\3\2\3\2\3\2\3\2\3\2\3\2\3\2\3\2"+
		"\3\2\3\2\3\2\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3"+
		"\3\3\3\3\3\3\3\3\3\3\3\3\3\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4"+
		"\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\5\3\5\3\5\3\5\3\5\3\5\3"+
		"\5\3\5\3\5\3\5\3\5\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\7\3\7"+
		"\3\7\3\7\3\7\3\b\3\b\3\b\3\b\3\b\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\n\3"+
		"\n\3\n\3\n\3\n\3\13\3\13\3\13\3\13\3\13\3\13\3\f\3\f\3\f\3\f\3\f\3\f\3"+
		"\r\3\r\3\r\3\r\3\r\3\r\3\r\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\17\3\17"+
		"\3\17\3\17\3\17\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\21"+
		"\3\21\3\21\3\22\3\22\3\22\3\23\3\23\3\23\3\24\3\24\3\24\3\24\3\24\3\24"+
		"\3\24\3\25\3\25\3\25\3\25\3\26\6\26\u0113\n\26\r\26\16\26\u0114\3\26\3"+
		"\26\6\26\u0119\n\26\r\26\16\26\u011a\3\27\6\27\u011e\n\27\r\27\16\27\u011f"+
		"\3\30\3\30\7\30\u0124\n\30\f\30\16\30\u0127\13\30\3\30\3\30\3\31\3\31"+
		"\6\31\u012d\n\31\r\31\16\31\u012e\3\32\3\32\6\32\u0133\n\32\r\32\16\32"+
		"\u0134\3\33\3\33\7\33\u0139\n\33\f\33\16\33\u013c\13\33\3\34\3\34\3\35"+
		"\3\35\3\36\3\36\3\37\3\37\3 \3 \3!\3!\3!\3\"\3\"\3\"\3#\3#\3#\3$\3$\3"+
		"$\3%\3%\3&\3&\3\'\3\'\3(\3(\3)\3)\3*\3*\3+\3+\3,\3,\3-\3-\3.\3.\3/\3/"+
		"\3\60\3\60\3\61\3\61\3\61\3\61\6\61\u0170\n\61\r\61\16\61\u0171\3\61\3"+
		"\61\3\61\3\61\3\61\3\62\3\62\3\62\3\62\6\62\u017d\n\62\r\62\16\62\u017e"+
		"\3\62\3\62\3\63\6\63\u0184\n\63\r\63\16\63\u0185\3\63\3\63\3\64\3\64\3"+
		"\64\2\2\65\3\3\5\4\7\5\t\6\13\7\r\b\17\t\21\n\23\13\25\f\27\r\31\16\33"+
		"\17\35\20\37\21!\22#\23%\24\'\25)\26+\27-\30/\31\61\32\63\33\65\34\67"+
		"\359\36;\37= ?!A\"C#E$G%I&K\'M(O)Q*S+U,W-Y.[/]\60_\61a\62c\63e\64g\2\3"+
		"\2\n\3\2\62;\3\2$$\7\2C\\aac|\u00d3\u00d3\u00f3\u00f3\b\2\62;C\\aac|\u00d3"+
		"\u00d3\u00f3\u00f3\3\2+,\3\2\f\f\6\2\13\f\17\17\"\"^^\t\2\"#%%--/\60<"+
		"<BB]_\2\u0194\2\3\3\2\2\2\2\5\3\2\2\2\2\7\3\2\2\2\2\t\3\2\2\2\2\13\3\2"+
		"\2\2\2\r\3\2\2\2\2\17\3\2\2\2\2\21\3\2\2\2\2\23\3\2\2\2\2\25\3\2\2\2\2"+
		"\27\3\2\2\2\2\31\3\2\2\2\2\33\3\2\2\2\2\35\3\2\2\2\2\37\3\2\2\2\2!\3\2"+
		"\2\2\2#\3\2\2\2\2%\3\2\2\2\2\'\3\2\2\2\2)\3\2\2\2\2+\3\2\2\2\2-\3\2\2"+
		"\2\2/\3\2\2\2\2\61\3\2\2\2\2\63\3\2\2\2\2\65\3\2\2\2\2\67\3\2\2\2\29\3"+
		"\2\2\2\2;\3\2\2\2\2=\3\2\2\2\2?\3\2\2\2\2A\3\2\2\2\2C\3\2\2\2\2E\3\2\2"+
		"\2\2G\3\2\2\2\2I\3\2\2\2\2K\3\2\2\2\2M\3\2\2\2\2O\3\2\2\2\2Q\3\2\2\2\2"+
		"S\3\2\2\2\2U\3\2\2\2\2W\3\2\2\2\2Y\3\2\2\2\2[\3\2\2\2\2]\3\2\2\2\2_\3"+
		"\2\2\2\2a\3\2\2\2\2c\3\2\2\2\2e\3\2\2\2\3i\3\2\2\2\5|\3\2\2\2\7\u0091"+
		"\3\2\2\2\t\u00a7\3\2\2\2\13\u00b2\3\2\2\2\r\u00bd\3\2\2\2\17\u00c2\3\2"+
		"\2\2\21\u00c7\3\2\2\2\23\u00cf\3\2\2\2\25\u00d4\3\2\2\2\27\u00da\3\2\2"+
		"\2\31\u00e0\3\2\2\2\33\u00e7\3\2\2\2\35\u00ee\3\2\2\2\37\u00f3\3\2\2\2"+
		"!\u00fd\3\2\2\2#\u0100\3\2\2\2%\u0103\3\2\2\2\'\u0106\3\2\2\2)\u010d\3"+
		"\2\2\2+\u0112\3\2\2\2-\u011d\3\2\2\2/\u0121\3\2\2\2\61\u012a\3\2\2\2\63"+
		"\u0130\3\2\2\2\65\u0136\3\2\2\2\67\u013d\3\2\2\29\u013f\3\2\2\2;\u0141"+
		"\3\2\2\2=\u0143\3\2\2\2?\u0145\3\2\2\2A\u0147\3\2\2\2C\u014a\3\2\2\2E"+
		"\u014d\3\2\2\2G\u0150\3\2\2\2I\u0153\3\2\2\2K\u0155\3\2\2\2M\u0157\3\2"+
		"\2\2O\u0159\3\2\2\2Q\u015b\3\2\2\2S\u015d\3\2\2\2U\u015f\3\2\2\2W\u0161"+
		"\3\2\2\2Y\u0163\3\2\2\2[\u0165\3\2\2\2]\u0167\3\2\2\2_\u0169\3\2\2\2a"+
		"\u016b\3\2\2\2c\u0178\3\2\2\2e\u0183\3\2\2\2g\u0189\3\2\2\2ij\7%\2\2j"+
		"k\7k\2\2kl\7p\2\2lm\7e\2\2mn\7n\2\2no\7w\2\2op\7f\2\2pq\7g\2\2qr\7\"\2"+
		"\2rs\7>\2\2st\7u\2\2tu\7v\2\2uv\7f\2\2vw\7k\2\2wx\7q\2\2xy\7\60\2\2yz"+
		"\7j\2\2z{\7@\2\2{\4\3\2\2\2|}\7f\2\2}~\7q\2\2~\177\7w\2\2\177\u0080\7"+
		"d\2\2\u0080\u0081\7n\2\2\u0081\u0082\7g\2\2\u0082\u0083\7\"\2\2\u0083"+
		"\u0084\7J\2\2\u0084\u0085\7G\2\2\u0085\u0086\7C\2\2\u0086\u0087\7R\2\2"+
		"\u0087\u0088\7]\2\2\u0088\u0089\7\63\2\2\u0089\u008a\7\62\2\2\u008a\u008b"+
		"\7\62\2\2\u008b\u008c\7\62\2\2\u008c\u008d\7\62\2\2\u008d\u008e\7\62\2"+
		"\2\u008e\u008f\7_\2\2\u008f\u0090\7=\2\2\u0090\6\3\2\2\2\u0091\u0092\7"+
		"f\2\2\u0092\u0093\7q\2\2\u0093\u0094\7w\2\2\u0094\u0095\7d\2\2\u0095\u0096"+
		"\7n\2\2\u0096\u0097\7g\2\2\u0097\u0098\7\"\2\2\u0098\u0099\7U\2\2\u0099"+
		"\u009a\7V\2\2\u009a\u009b\7C\2\2\u009b\u009c\7E\2\2\u009c\u009d\7M\2\2"+
		"\u009d\u009e\7]\2\2\u009e\u009f\7\63\2\2\u009f\u00a0\7\62\2\2\u00a0\u00a1"+
		"\7\62\2\2\u00a1\u00a2\7\62\2\2\u00a2\u00a3\7\62\2\2\u00a3\u00a4\7\62\2"+
		"\2\u00a4\u00a5\7_\2\2\u00a5\u00a6\7=\2\2\u00a6\b\3\2\2\2\u00a7\u00a8\7"+
		"f\2\2\u00a8\u00a9\7q\2\2\u00a9\u00aa\7w\2\2\u00aa\u00ab\7d\2\2\u00ab\u00ac"+
		"\7n\2\2\u00ac\u00ad\7g\2\2\u00ad\u00ae\7\"\2\2\u00ae\u00af\7U\2\2\u00af"+
		"\u00b0\7R\2\2\u00b0\u00b1\7=\2\2\u00b1\n\3\2\2\2\u00b2\u00b3\7f\2\2\u00b3"+
		"\u00b4\7q\2\2\u00b4\u00b5\7w\2\2\u00b5\u00b6\7d\2\2\u00b6\u00b7\7n\2\2"+
		"\u00b7\u00b8\7g\2\2\u00b8\u00b9\7\"\2\2\u00b9\u00ba\7J\2\2\u00ba\u00bb"+
		"\7R\2\2\u00bb\u00bc\7=\2\2\u00bc\f\3\2\2\2\u00bd\u00be\7x\2\2\u00be\u00bf"+
		"\7q\2\2\u00bf\u00c0\7k\2\2\u00c0\u00c1\7f\2\2\u00c1\16\3\2\2\2\u00c2\u00c3"+
		"\7i\2\2\u00c3\u00c4\7q\2\2\u00c4\u00c5\7v\2\2\u00c5\u00c6\7q\2\2\u00c6"+
		"\20\3\2\2\2\u00c7\u00c8\7t\2\2\u00c8\u00c9\7g\2\2\u00c9\u00ca\7v\2\2\u00ca"+
		"\u00cb\7w\2\2\u00cb\u00cc\7t\2\2\u00cc\u00cd\7p\2\2\u00cd\u00ce\7=\2\2"+
		"\u00ce\22\3\2\2\2\u00cf\u00d0\7J\2\2\u00d0\u00d1\7G\2\2\u00d1\u00d2\7"+
		"C\2\2\u00d2\u00d3\7R\2\2\u00d3\24\3\2\2\2\u00d4\u00d5\7U\2\2\u00d5\u00d6"+
		"\7V\2\2\u00d6\u00d7\7C\2\2\u00d7\u00d8\7E\2\2\u00d8\u00d9\7M\2\2\u00d9"+
		"\26\3\2\2\2\u00da\u00db\7*\2\2\u00db\u00dc\7k\2\2\u00dc\u00dd\7p\2\2\u00dd"+
		"\u00de\7v\2\2\u00de\u00df\7+\2\2\u00df\30\3\2\2\2\u00e0\u00e1\7*\2\2\u00e1"+
		"\u00e2\7e\2\2\u00e2\u00e3\7j\2\2\u00e3\u00e4\7c\2\2\u00e4\u00e5\7t\2\2"+
		"\u00e5\u00e6\7+\2\2\u00e6\32\3\2\2\2\u00e7\u00e8\7r\2\2\u00e8\u00e9\7"+
		"t\2\2\u00e9\u00ea\7k\2\2\u00ea\u00eb\7p\2\2\u00eb\u00ec\7v\2\2\u00ec\u00ed"+
		"\7h\2\2\u00ed\34\3\2\2\2\u00ee\u00ef\7o\2\2\u00ef\u00f0\7c\2\2\u00f0\u00f1"+
		"\7k\2\2\u00f1\u00f2\7p\2\2\u00f2\36\3\2\2\2\u00f3\u00f4\7t\2\2\u00f4\u00f5"+
		"\7g\2\2\u00f5\u00f6\7v\2\2\u00f6\u00f7\7w\2\2\u00f7\u00f8\7t\2\2\u00f8"+
		"\u00f9\7p\2\2\u00f9\u00fa\7\"\2\2\u00fa\u00fb\7\62\2\2\u00fb\u00fc\7="+
		"\2\2\u00fc \3\2\2\2\u00fd\u00fe\7U\2\2\u00fe\u00ff\7R\2\2\u00ff\"\3\2"+
		"\2\2\u0100\u0101\7J\2\2\u0101\u0102\7R\2\2\u0102$\3\2\2\2\u0103\u0104"+
		"\7k\2\2\u0104\u0105\7h\2\2\u0105&\3\2\2\2\u0106\u0107\7f\2\2\u0107\u0108"+
		"\7q\2\2\u0108\u0109\7w\2\2\u0109\u010a\7d\2\2\u010a\u010b\7n\2\2\u010b"+
		"\u010c\7g\2\2\u010c(\3\2\2\2\u010d\u010e\7k\2\2\u010e\u010f\7p\2\2\u010f"+
		"\u0110\7v\2\2\u0110*\3\2\2\2\u0111\u0113\t\2\2\2\u0112\u0111\3\2\2\2\u0113"+
		"\u0114\3\2\2\2\u0114\u0112\3\2\2\2\u0114\u0115\3\2\2\2\u0115\u0116\3\2"+
		"\2\2\u0116\u0118\7\60\2\2\u0117\u0119\t\2\2\2\u0118\u0117\3\2\2\2\u0119"+
		"\u011a\3\2\2\2\u011a\u0118\3\2\2\2\u011a\u011b\3\2\2\2\u011b,\3\2\2\2"+
		"\u011c\u011e\t\2\2\2\u011d\u011c\3\2\2\2\u011e\u011f\3\2\2\2\u011f\u011d"+
		"\3\2\2\2\u011f\u0120\3\2\2\2\u0120.\3\2\2\2\u0121\u0125\7$\2\2\u0122\u0124"+
		"\n\3\2\2\u0123\u0122\3\2\2\2\u0124\u0127\3\2\2\2\u0125\u0123\3\2\2\2\u0125"+
		"\u0126\3\2\2\2\u0126\u0128\3\2\2\2\u0127\u0125\3\2\2\2\u0128\u0129\7$"+
		"\2\2\u0129\60\3\2\2\2\u012a\u012c\7V\2\2\u012b\u012d\t\2\2\2\u012c\u012b"+
		"\3\2\2\2\u012d\u012e\3\2\2\2\u012e\u012c\3\2\2\2\u012e\u012f\3\2\2\2\u012f"+
		"\62\3\2\2\2\u0130\u0132\7N\2\2\u0131\u0133\t\2\2\2\u0132\u0131\3\2\2\2"+
		"\u0133\u0134\3\2\2\2\u0134\u0132\3\2\2\2\u0134\u0135\3\2\2\2\u0135\64"+
		"\3\2\2\2\u0136\u013a\t\4\2\2\u0137\u0139\t\5\2\2\u0138\u0137\3\2\2\2\u0139"+
		"\u013c\3\2\2\2\u013a\u0138\3\2\2\2\u013a\u013b\3\2\2\2\u013b\66\3\2\2"+
		"\2\u013c\u013a\3\2\2\2\u013d\u013e\7-\2\2\u013e8\3\2\2\2\u013f\u0140\7"+
		"/\2\2\u0140:\3\2\2\2\u0141\u0142\7,\2\2\u0142<\3\2\2\2\u0143\u0144\7\61"+
		"\2\2\u0144>\3\2\2\2\u0145\u0146\7\'\2\2\u0146@\3\2\2\2\u0147\u0148\7>"+
		"\2\2\u0148\u0149\7?\2\2\u0149B\3\2\2\2\u014a\u014b\7@\2\2\u014b\u014c"+
		"\7?\2\2\u014cD\3\2\2\2\u014d\u014e\7?\2\2\u014e\u014f\7?\2\2\u014fF\3"+
		"\2\2\2\u0150\u0151\7#\2\2\u0151\u0152\7?\2\2\u0152H\3\2\2\2\u0153\u0154"+
		"\7>\2\2\u0154J\3\2\2\2\u0155\u0156\7@\2\2\u0156L\3\2\2\2\u0157\u0158\7"+
		"}\2\2\u0158N\3\2\2\2\u0159\u015a\7\177\2\2\u015aP\3\2\2\2\u015b\u015c"+
		"\7]\2\2\u015cR\3\2\2\2\u015d\u015e\7_\2\2\u015eT\3\2\2\2\u015f\u0160\7"+
		"*\2\2\u0160V\3\2\2\2\u0161\u0162\7+\2\2\u0162X\3\2\2\2\u0163\u0164\7="+
		"\2\2\u0164Z\3\2\2\2\u0165\u0166\7<\2\2\u0166\\\3\2\2\2\u0167\u0168\7."+
		"\2\2\u0168^\3\2\2\2\u0169\u016a\7?\2\2\u016a`\3\2\2\2\u016b\u016c\7\61"+
		"\2\2\u016c\u016d\7,\2\2\u016d\u016f\3\2\2\2\u016e\u0170\n\6\2\2\u016f"+
		"\u016e\3\2\2\2\u0170\u0171\3\2\2\2\u0171\u016f\3\2\2\2\u0171\u0172\3\2"+
		"\2\2\u0172\u0173\3\2\2\2\u0173\u0174\7,\2\2\u0174\u0175\7\61\2\2\u0175"+
		"\u0176\3\2\2\2\u0176\u0177\b\61\2\2\u0177b\3\2\2\2\u0178\u0179\7\61\2"+
		"\2\u0179\u017a\7\61\2\2\u017a\u017c\3\2\2\2\u017b\u017d\n\7\2\2\u017c"+
		"\u017b\3\2\2\2\u017d\u017e\3\2\2\2\u017e\u017c\3\2\2\2\u017e\u017f\3\2"+
		"\2\2\u017f\u0180\3\2\2\2\u0180\u0181\b\62\2\2\u0181d\3\2\2\2\u0182\u0184"+
		"\t\b\2\2\u0183\u0182\3\2\2\2\u0184\u0185\3\2\2\2\u0185\u0183\3\2\2\2\u0185"+
		"\u0186\3\2\2\2\u0186\u0187\3\2\2\2\u0187\u0188\b\63\2\2\u0188f\3\2\2\2"+
		"\u0189\u018a\7^\2\2\u018a\u018b\t\t\2\2\u018bh\3\2\2\2\r\2\u0114\u011a"+
		"\u011f\u0125\u012e\u0134\u013a\u0171\u017e\u0185\3\b\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}