// Generated from /Users/xiaozhen/code/huabaobao/go-zero/tools/goctl/api/parser/g4/ApiLexer.g4 by ANTLR 4.10.1
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class ApiLexer extends Lexer {
	static { RuntimeMetaData.checkVersion("4.10.1", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		ATDOC=1, ATHANDLER=2, INTERFACE=3, ATSERVER=4, WS=5, COMMENT=6, LINE_COMMENT=7, 
		STRING=8, RAW_STRING=9, LINE_VALUE=10, ID=11;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"ATDOC", "ATHANDLER", "INTERFACE", "ATSERVER", "WS", "COMMENT", "LINE_COMMENT", 
			"STRING", "RAW_STRING", "LINE_VALUE", "ID", "ExponentPart", "EscapeSequence", 
			"HexDigits", "HexDigit", "Digits", "LetterOrDigit", "Letter"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'@doc'", "'@handler'", "'interface{}'", "'@server'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "ATDOC", "ATHANDLER", "INTERFACE", "ATSERVER", "WS", "COMMENT", 
			"LINE_COMMENT", "STRING", "RAW_STRING", "LINE_VALUE", "ID"
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


	public ApiLexer(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "ApiLexer.g4"; }

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
		"\u0004\u0000\u000b\u00cf\u0006\uffff\uffff\u0002\u0000\u0007\u0000\u0002"+
		"\u0001\u0007\u0001\u0002\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002"+
		"\u0004\u0007\u0004\u0002\u0005\u0007\u0005\u0002\u0006\u0007\u0006\u0002"+
		"\u0007\u0007\u0007\u0002\b\u0007\b\u0002\t\u0007\t\u0002\n\u0007\n\u0002"+
		"\u000b\u0007\u000b\u0002\f\u0007\f\u0002\r\u0007\r\u0002\u000e\u0007\u000e"+
		"\u0002\u000f\u0007\u000f\u0002\u0010\u0007\u0010\u0002\u0011\u0007\u0011"+
		"\u0001\u0000\u0001\u0000\u0001\u0000\u0001\u0000\u0001\u0000\u0001\u0001"+
		"\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001"+
		"\u0001\u0001\u0001\u0001\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002"+
		"\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002"+
		"\u0001\u0002\u0001\u0002\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003"+
		"\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0004\u0004\u0004"+
		"I\b\u0004\u000b\u0004\f\u0004J\u0001\u0004\u0001\u0004\u0001\u0005\u0001"+
		"\u0005\u0001\u0005\u0001\u0005\u0005\u0005S\b\u0005\n\u0005\f\u0005V\t"+
		"\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001"+
		"\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0005\u0006a\b\u0006\n\u0006"+
		"\f\u0006d\t\u0006\u0001\u0006\u0001\u0006\u0001\u0007\u0001\u0007\u0001"+
		"\u0007\u0005\u0007k\b\u0007\n\u0007\f\u0007n\t\u0007\u0001\u0007\u0001"+
		"\u0007\u0001\b\u0001\b\u0001\b\u0004\bu\b\b\u000b\b\f\bv\u0001\b\u0001"+
		"\b\u0001\t\u0001\t\u0005\t}\b\t\n\t\f\t\u0080\t\t\u0001\t\u0001\t\u0005"+
		"\t\u0084\b\t\n\t\f\t\u0087\t\t\u0003\t\u0089\b\t\u0001\n\u0001\n\u0005"+
		"\n\u008d\b\n\n\n\f\n\u0090\t\n\u0001\u000b\u0001\u000b\u0003\u000b\u0094"+
		"\b\u000b\u0001\u000b\u0001\u000b\u0001\f\u0001\f\u0001\f\u0001\f\u0003"+
		"\f\u009c\b\f\u0001\f\u0003\f\u009f\b\f\u0001\f\u0001\f\u0001\f\u0004\f"+
		"\u00a4\b\f\u000b\f\f\f\u00a5\u0001\f\u0001\f\u0001\f\u0001\f\u0001\f\u0003"+
		"\f\u00ad\b\f\u0001\r\u0001\r\u0001\r\u0005\r\u00b2\b\r\n\r\f\r\u00b5\t"+
		"\r\u0001\r\u0003\r\u00b8\b\r\u0001\u000e\u0001\u000e\u0001\u000f\u0001"+
		"\u000f\u0005\u000f\u00be\b\u000f\n\u000f\f\u000f\u00c1\t\u000f\u0001\u000f"+
		"\u0003\u000f\u00c4\b\u000f\u0001\u0010\u0001\u0010\u0003\u0010\u00c8\b"+
		"\u0010\u0001\u0011\u0001\u0011\u0001\u0011\u0001\u0011\u0003\u0011\u00ce"+
		"\b\u0011\u0001T\u0000\u0012\u0001\u0001\u0003\u0002\u0005\u0003\u0007"+
		"\u0004\t\u0005\u000b\u0006\r\u0007\u000f\b\u0011\t\u0013\n\u0015\u000b"+
		"\u0017\u0000\u0019\u0000\u001b\u0000\u001d\u0000\u001f\u0000!\u0000#\u0000"+
		"\u0001\u0000\u0012\u0003\u0000\t\n\f\r  \u0002\u0000\n\n\r\r\u0002\u0000"+
		"\"\"\\\\\u0004\u0000\n\n\r\r\\\\``\u0002\u0000\t\t  \u0004\u0000\n\n\r"+
		"\r\"\"``\u0002\u0000EEee\u0002\u0000++--\b\u0000\"\"\'\'\\\\bbffnnrrt"+
		"t\u0001\u000003\u0001\u000007\u0003\u000009AFaf\u0001\u000009\u0002\u0000"+
		"09__\u0004\u0000$$AZ__az\u0002\u0000\u0000\u007f\u8000\ud800\u8000\udbff"+
		"\u0001\u0000\u8000\ud800\u8000\udbff\u0001\u0000\u8000\udc00\u8000\udfff"+
		"\u00e0\u0000\u0001\u0001\u0000\u0000\u0000\u0000\u0003\u0001\u0000\u0000"+
		"\u0000\u0000\u0005\u0001\u0000\u0000\u0000\u0000\u0007\u0001\u0000\u0000"+
		"\u0000\u0000\t\u0001\u0000\u0000\u0000\u0000\u000b\u0001\u0000\u0000\u0000"+
		"\u0000\r\u0001\u0000\u0000\u0000\u0000\u000f\u0001\u0000\u0000\u0000\u0000"+
		"\u0011\u0001\u0000\u0000\u0000\u0000\u0013\u0001\u0000\u0000\u0000\u0000"+
		"\u0015\u0001\u0000\u0000\u0000\u0001%\u0001\u0000\u0000\u0000\u0003*\u0001"+
		"\u0000\u0000\u0000\u00053\u0001\u0000\u0000\u0000\u0007?\u0001\u0000\u0000"+
		"\u0000\tH\u0001\u0000\u0000\u0000\u000bN\u0001\u0000\u0000\u0000\r\\\u0001"+
		"\u0000\u0000\u0000\u000fg\u0001\u0000\u0000\u0000\u0011q\u0001\u0000\u0000"+
		"\u0000\u0013z\u0001\u0000\u0000\u0000\u0015\u008a\u0001\u0000\u0000\u0000"+
		"\u0017\u0091\u0001\u0000\u0000\u0000\u0019\u00ac\u0001\u0000\u0000\u0000"+
		"\u001b\u00ae\u0001\u0000\u0000\u0000\u001d\u00b9\u0001\u0000\u0000\u0000"+
		"\u001f\u00bb\u0001\u0000\u0000\u0000!\u00c7\u0001\u0000\u0000\u0000#\u00cd"+
		"\u0001\u0000\u0000\u0000%&\u0005@\u0000\u0000&\'\u0005d\u0000\u0000\'"+
		"(\u0005o\u0000\u0000()\u0005c\u0000\u0000)\u0002\u0001\u0000\u0000\u0000"+
		"*+\u0005@\u0000\u0000+,\u0005h\u0000\u0000,-\u0005a\u0000\u0000-.\u0005"+
		"n\u0000\u0000./\u0005d\u0000\u0000/0\u0005l\u0000\u000001\u0005e\u0000"+
		"\u000012\u0005r\u0000\u00002\u0004\u0001\u0000\u0000\u000034\u0005i\u0000"+
		"\u000045\u0005n\u0000\u000056\u0005t\u0000\u000067\u0005e\u0000\u0000"+
		"78\u0005r\u0000\u000089\u0005f\u0000\u00009:\u0005a\u0000\u0000:;\u0005"+
		"c\u0000\u0000;<\u0005e\u0000\u0000<=\u0005{\u0000\u0000=>\u0005}\u0000"+
		"\u0000>\u0006\u0001\u0000\u0000\u0000?@\u0005@\u0000\u0000@A\u0005s\u0000"+
		"\u0000AB\u0005e\u0000\u0000BC\u0005r\u0000\u0000CD\u0005v\u0000\u0000"+
		"DE\u0005e\u0000\u0000EF\u0005r\u0000\u0000F\b\u0001\u0000\u0000\u0000"+
		"GI\u0007\u0000\u0000\u0000HG\u0001\u0000\u0000\u0000IJ\u0001\u0000\u0000"+
		"\u0000JH\u0001\u0000\u0000\u0000JK\u0001\u0000\u0000\u0000KL\u0001\u0000"+
		"\u0000\u0000LM\u0006\u0004\u0000\u0000M\n\u0001\u0000\u0000\u0000NO\u0005"+
		"/\u0000\u0000OP\u0005*\u0000\u0000PT\u0001\u0000\u0000\u0000QS\t\u0000"+
		"\u0000\u0000RQ\u0001\u0000\u0000\u0000SV\u0001\u0000\u0000\u0000TU\u0001"+
		"\u0000\u0000\u0000TR\u0001\u0000\u0000\u0000UW\u0001\u0000\u0000\u0000"+
		"VT\u0001\u0000\u0000\u0000WX\u0005*\u0000\u0000XY\u0005/\u0000\u0000Y"+
		"Z\u0001\u0000\u0000\u0000Z[\u0006\u0005\u0001\u0000[\f\u0001\u0000\u0000"+
		"\u0000\\]\u0005/\u0000\u0000]^\u0005/\u0000\u0000^b\u0001\u0000\u0000"+
		"\u0000_a\b\u0001\u0000\u0000`_\u0001\u0000\u0000\u0000ad\u0001\u0000\u0000"+
		"\u0000b`\u0001\u0000\u0000\u0000bc\u0001\u0000\u0000\u0000ce\u0001\u0000"+
		"\u0000\u0000db\u0001\u0000\u0000\u0000ef\u0006\u0006\u0001\u0000f\u000e"+
		"\u0001\u0000\u0000\u0000gl\u0005\"\u0000\u0000hk\b\u0002\u0000\u0000i"+
		"k\u0003\u0019\f\u0000jh\u0001\u0000\u0000\u0000ji\u0001\u0000\u0000\u0000"+
		"kn\u0001\u0000\u0000\u0000lj\u0001\u0000\u0000\u0000lm\u0001\u0000\u0000"+
		"\u0000mo\u0001\u0000\u0000\u0000nl\u0001\u0000\u0000\u0000op\u0005\"\u0000"+
		"\u0000p\u0010\u0001\u0000\u0000\u0000qt\u0005`\u0000\u0000ru\b\u0003\u0000"+
		"\u0000su\u0003\u0019\f\u0000tr\u0001\u0000\u0000\u0000ts\u0001\u0000\u0000"+
		"\u0000uv\u0001\u0000\u0000\u0000vt\u0001\u0000\u0000\u0000vw\u0001\u0000"+
		"\u0000\u0000wx\u0001\u0000\u0000\u0000xy\u0005`\u0000\u0000y\u0012\u0001"+
		"\u0000\u0000\u0000z~\u0005:\u0000\u0000{}\u0007\u0004\u0000\u0000|{\u0001"+
		"\u0000\u0000\u0000}\u0080\u0001\u0000\u0000\u0000~|\u0001\u0000\u0000"+
		"\u0000~\u007f\u0001\u0000\u0000\u0000\u007f\u0088\u0001\u0000\u0000\u0000"+
		"\u0080~\u0001\u0000\u0000\u0000\u0081\u0089\u0003\u000f\u0007\u0000\u0082"+
		"\u0084\b\u0005\u0000\u0000\u0083\u0082\u0001\u0000\u0000\u0000\u0084\u0087"+
		"\u0001\u0000\u0000\u0000\u0085\u0083\u0001\u0000\u0000\u0000\u0085\u0086"+
		"\u0001\u0000\u0000\u0000\u0086\u0089\u0001\u0000\u0000\u0000\u0087\u0085"+
		"\u0001\u0000\u0000\u0000\u0088\u0081\u0001\u0000\u0000\u0000\u0088\u0085"+
		"\u0001\u0000\u0000\u0000\u0089\u0014\u0001\u0000\u0000\u0000\u008a\u008e"+
		"\u0003#\u0011\u0000\u008b\u008d\u0003!\u0010\u0000\u008c\u008b\u0001\u0000"+
		"\u0000\u0000\u008d\u0090\u0001\u0000\u0000\u0000\u008e\u008c\u0001\u0000"+
		"\u0000\u0000\u008e\u008f\u0001\u0000\u0000\u0000\u008f\u0016\u0001\u0000"+
		"\u0000\u0000\u0090\u008e\u0001\u0000\u0000\u0000\u0091\u0093\u0007\u0006"+
		"\u0000\u0000\u0092\u0094\u0007\u0007\u0000\u0000\u0093\u0092\u0001\u0000"+
		"\u0000\u0000\u0093\u0094\u0001\u0000\u0000\u0000\u0094\u0095\u0001\u0000"+
		"\u0000\u0000\u0095\u0096\u0003\u001f\u000f\u0000\u0096\u0018\u0001\u0000"+
		"\u0000\u0000\u0097\u0098\u0005\\\u0000\u0000\u0098\u00ad\u0007\b\u0000"+
		"\u0000\u0099\u009e\u0005\\\u0000\u0000\u009a\u009c\u0007\t\u0000\u0000"+
		"\u009b\u009a\u0001\u0000\u0000\u0000\u009b\u009c\u0001\u0000\u0000\u0000"+
		"\u009c\u009d\u0001\u0000\u0000\u0000\u009d\u009f\u0007\n\u0000\u0000\u009e"+
		"\u009b\u0001\u0000\u0000\u0000\u009e\u009f\u0001\u0000\u0000\u0000\u009f"+
		"\u00a0\u0001\u0000\u0000\u0000\u00a0\u00ad\u0007\n\u0000\u0000\u00a1\u00a3"+
		"\u0005\\\u0000\u0000\u00a2\u00a4\u0005u\u0000\u0000\u00a3\u00a2\u0001"+
		"\u0000\u0000\u0000\u00a4\u00a5\u0001\u0000\u0000\u0000\u00a5\u00a3\u0001"+
		"\u0000\u0000\u0000\u00a5\u00a6\u0001\u0000\u0000\u0000\u00a6\u00a7\u0001"+
		"\u0000\u0000\u0000\u00a7\u00a8\u0003\u001d\u000e\u0000\u00a8\u00a9\u0003"+
		"\u001d\u000e\u0000\u00a9\u00aa\u0003\u001d\u000e\u0000\u00aa\u00ab\u0003"+
		"\u001d\u000e\u0000\u00ab\u00ad\u0001\u0000\u0000\u0000\u00ac\u0097\u0001"+
		"\u0000\u0000\u0000\u00ac\u0099\u0001\u0000\u0000\u0000\u00ac\u00a1\u0001"+
		"\u0000\u0000\u0000\u00ad\u001a\u0001\u0000\u0000\u0000\u00ae\u00b7\u0003"+
		"\u001d\u000e\u0000\u00af\u00b2\u0003\u001d\u000e\u0000\u00b0\u00b2\u0005"+
		"_\u0000\u0000\u00b1\u00af\u0001\u0000\u0000\u0000\u00b1\u00b0\u0001\u0000"+
		"\u0000\u0000\u00b2\u00b5\u0001\u0000\u0000\u0000\u00b3\u00b1\u0001\u0000"+
		"\u0000\u0000\u00b3\u00b4\u0001\u0000\u0000\u0000\u00b4\u00b6\u0001\u0000"+
		"\u0000\u0000\u00b5\u00b3\u0001\u0000\u0000\u0000\u00b6\u00b8\u0003\u001d"+
		"\u000e\u0000\u00b7\u00b3\u0001\u0000\u0000\u0000\u00b7\u00b8\u0001\u0000"+
		"\u0000\u0000\u00b8\u001c\u0001\u0000\u0000\u0000\u00b9\u00ba\u0007\u000b"+
		"\u0000\u0000\u00ba\u001e\u0001\u0000\u0000\u0000\u00bb\u00c3\u0007\f\u0000"+
		"\u0000\u00bc\u00be\u0007\r\u0000\u0000\u00bd\u00bc\u0001\u0000\u0000\u0000"+
		"\u00be\u00c1\u0001\u0000\u0000\u0000\u00bf\u00bd\u0001\u0000\u0000\u0000"+
		"\u00bf\u00c0\u0001\u0000\u0000\u0000\u00c0\u00c2\u0001\u0000\u0000\u0000"+
		"\u00c1\u00bf\u0001\u0000\u0000\u0000\u00c2\u00c4\u0007\f\u0000\u0000\u00c3"+
		"\u00bf\u0001\u0000\u0000\u0000\u00c3\u00c4\u0001\u0000\u0000\u0000\u00c4"+
		" \u0001\u0000\u0000\u0000\u00c5\u00c8\u0003#\u0011\u0000\u00c6\u00c8\u0007"+
		"\f\u0000\u0000\u00c7\u00c5\u0001\u0000\u0000\u0000\u00c7\u00c6\u0001\u0000"+
		"\u0000\u0000\u00c8\"\u0001\u0000\u0000\u0000\u00c9\u00ce\u0007\u000e\u0000"+
		"\u0000\u00ca\u00ce\b\u000f\u0000\u0000\u00cb\u00cc\u0007\u0010\u0000\u0000"+
		"\u00cc\u00ce\u0007\u0011\u0000\u0000\u00cd\u00c9\u0001\u0000\u0000\u0000"+
		"\u00cd\u00ca\u0001\u0000\u0000\u0000\u00cd\u00cb\u0001\u0000\u0000\u0000"+
		"\u00ce$\u0001\u0000\u0000\u0000\u0018\u0000JTbjltv~\u0085\u0088\u008e"+
		"\u0093\u009b\u009e\u00a5\u00ac\u00b1\u00b3\u00b7\u00bf\u00c3\u00c7\u00cd"+
		"\u0002\u0000\u0001\u0000\u0000X\u0000";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}