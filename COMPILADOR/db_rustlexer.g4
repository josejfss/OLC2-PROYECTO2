lexer grammar db_rustlexer;

// ********** TOKENS ***********
// ********** PALABRAS RESERVADAS **********
// ACCESO
TK_PUBLICO:     'pub';
//TIPOS
TK_TIPOINT:     'i64';
TK_TIPOFLOAT:   'f64';
TK_TIPOBOOL:    'bool';
TK_TIPOCHAR:    'char';
TK_TIPOSTRING:  'String';
TK_DIRSTRING:   '&str';
//CASTEO
TK_AS:          'as';
//IMPRIMIR
TK_PRINTLN:     'println!';
//DECLARAR
TK_LET:         'let';
TK_MUT:         'mut';
TK_DIRMUT:      '&mut';
//FUNCIONES
TK_FUNCION:     'fn';
TK_MAIN:        'main';
//FUNCIONES NATIVAS
TK_ABS:         'abs';
TK_SQRT:        'sqrt';
TK_TOSTRING:    'to_string()';
TK_TOOWNED:     'to_owned()';
TK_CLONE:       'clone()';
//VECTORES
TK_VECTOR:      'Vec';
TK_VECT:        'vec!';      
//FUNCIONES NATIVAS VECTORES
TK_NEW:         '::new()';
TK_LEN:         'len()';
TK_PUSH:        'push';
TK_REMOVE:      'remove';
TK_CONTAINS:    '&contains';
TK_INSERT:      'insert';
TK_CAPACITY:    'capacity()';
TK_WCAPACITY:   '::with_capacity';
//SENTENCIAS 
TK_IF:          'if';
TK_ELSE:        'else';
TK_MATCH:       'match';
//CICLOS
TK_LOOP:        'loop';
TK_WHILE:       'while';
TK_FOR:         'for';
TK_IN:          'in';
//SENTENCIAS DE RETORNO
TK_BREAK:       'break';
TK_CONTINUE:    'continue';
TK_RETURN:      'return';
//STRUCT
TK_STRUCT:      'Struct';
//MODULO
TK_MOD:         'mod';
//POTENCIAS
TK_POW:         '::pow';
TK_POWF:        '::powf';
//BOOLS
TK_TRUE:        'true';
TK_FALSE:       'false';
// ********** EXPRESIONES REGUALES **********
TK_FLOAT:           [0-9]+('.'[0-9]+);
TK_ENTERO:          [0-9]+;
TK_CADENA:          '"'~["]*'"';
TK_CARACTER:        '\''~[']*'\'';
TK_IDENTIFICADOR:   [a-zñA-ZÑ_][_a-zñA-ZÑ0-9]*;
// ********** OPERACION MATEMATICA **********
TK_SUMA:    '+';
TK_RESTA:   '-';
TK_MULTI:   '*';
TK_DIVI:    '/';
TK_MODULO:  '%';
// ********** OPERACION RELACIONAL **********
TK_MENORIGUAL:  '<=';
TK_MAYORIGUAL:  '>=';
TK_IGUALDAD:    '==';
TK_DESIGUALDAD: '!=';
TK_MENOR:   '<';
TK_MAYOR:   '>';
// ********** OPERACION LOGICA **********
TK_OR:  '||';
TK_AND: '&&';
TK_NOT: '!';
// ********** SIMBOLOS LENGUAJE **********
TK_LI:      '{';
TK_LD:      '}';
TK_CI:      '[';
TK_CD:      ']';
TK_PI:      '(';
TK_PD:      ')';
TK_PYC:     ';';
TK_DP:      ':';
TK_COMA:    ',';
TK_PUNTO:   '.'; 
TK_IGUAL:   '=';
TK_BARRA:   '|';
TK_GBAJO:   '_';
TK_REFER:   '&';
// ********** IGNORAR **********
COMENTARIO_MUL: '/*' (~[*)])+ '*/' -> skip;
COMENTARIO_LIN: '//'(~[\n])+ -> skip;
WHITESPACE: [ \\\r\n\t]+ -> skip;
fragment
ESC_SEQ
    :   '\\' ('\\'|'@'|'['|']'|'.'|'#'|'+'|'-'|'!'|':'|' ')
    ;



