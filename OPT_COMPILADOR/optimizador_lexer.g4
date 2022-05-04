lexer grammar optimizador_lexer;

// ********** TOKENS ***********
// ********** ENCABEZADO **********
TK_STDIOH:         '#include <stdio.h>';
TK_DHEAP:          'double HEAP[100000];';
TK_DSTACK:         'double STACK[100000];';
TK_PSTAKC:         'double SP;';
TK_PHEAP:          'double HP;';
// ********** PALABRAS RESERVADAS ***********
TK_VOID:        'void';
TK_GOTO:        'goto';
TK_RETURN:      'return;';
TK_HEAP:        'HEAP';
TK_STACK:       'STACK';
TK_CASTINT:     '(int)';
TK_CASTCHAR:    '(char)';
TK_PRINTF:      'printf';
TK_MAIN:        'main';
TK_RETMAIN:     'return 0;';
TK_PUNSTACK:    'SP';
TK_PUNHEAP:     'HP';
TK_IF:          'if';
TK_DOUBLE:      'double';
TK_INT:         'int';
TK_ASIGPSTACK:  'SP=0;';
TK_ASIGHEAP:    'HP=0;';
// ********** EXPRESIONES REGUALES **********
TK_FLOAT:           [0-9]+('.'[0-9]+);
TK_ENTERO:          [0-9]+;
TK_CADENA:          '"'~["]*'"';
TK_TEMPORAL:        'T'[0-9]+;
TK_ETIQUETA:        'L'[0-9]+;
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
TK_IGUAL:   '=';
// ********** IGNORAR **********
COMENTARIO_MUL: '/*' (~[*)])+ '*/' -> skip;
COMENTARIO_LIN: '//'(~[\n])+ -> skip;
WHITESPACE: [ \\\r\n\t]+ -> skip;
fragment
ESC_SEQ
    :   '\\' ('\\'|'@'|'['|']'|'.'|'#'|'+'|'-'|'!'|':'|' ')
    ;



