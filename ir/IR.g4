grammar IR;

// Statement

STAT_SEQ
    : STAT STAT_SEQ_TAIL
    ;

STAT_SEQ_TAIL
    : STAT_SEQ
    |
    ;

STAT
    : STAT_GLOBAL
    | STAT_DECLARE_FUNCT
    | STAT_DEFINE_FUNCT
    | STAT_METADATA
    ;

// Global variable

GLOBAL_IDENT
    : GLOBAL_NAME
    | GLOBAL_ID
;

GLOBAL_NAME
    : AT ( NAME | QUOTED_NAME )
;

GLOBAL_ID
    : AT ID
;

STAT_GLOBAL
    : GLOBAL_ID EQUAL VISIBILITY (UNNAMMED_ADDR)? (GLOBAL | CONSTANT)? TYPE EXPR_CONSTANT
    ;

VISIBILITY
    : PRIVATE
    ;

// Function

STAT_DECLARE_FUNCT
    : DECLARE FUNCT_SIGNATURE
    ;

STAT_DEFINE_FUNCT
    : DEFINE FUNCT_SIGNATURE LBRACE INSTRUCTION_SEQ RBRACE
    ;

FUNCT_SIGNATURE
    : TYPE GLOBAL_IDENT LPAREN FUNCT_ARG_LIST RPAREN FUNCT_ATTR_LIST
    ;

FUNCT_ARG_LIST
    : FUNCT_ARG FUNCT_ARG_LIST
    |
    ;

FUNCT_ARG
    : TYPE
    ;

FUNCT_ATTR_LIST
    : FUNCT_ATTR FUNCT_ATTR_LIST
    |
    ;

FUNCT_ATTR
    : NOUNWIND
    | NOCAPTURE
    ;

// Instruction

INSTRUCTION_SEQ
    : INSTRUCTION INSTRUCTION_SEQ_TAIL
    ;

INSTRUCTION_SEQ_TAIL
    : INSTRUCTION_SEQ
    |
    ;

INSTRUCTION
    :
    ;

// Expression

EXPR_CONSTANT
    : EXPR_CONSTANT_STR
    ;

EXPR_CONSTANT_STR
    : 'c' DOUBLEQUOTE .* DOUBLEQUOTE
    ;

// Identifiers

NAME
    : LETTER ( LETTER | DECIMAL_DIGIT ) *
;

ESCAPE_NAME
    : ESCAPE_LETTER ( ESCAPE_LETTER | DECIMAL_DIGIT ) *
;

QUOTED_NAME
    : QUOTED_STRING
;

ID
    : DECIMALS
;

// Comment and whitespace

fragment ASCII_LETTER_UPPER
    : 'A'..'Z'
;

fragment ASCII_LETTER_LOWER
    : 'a'..'z'
;

fragment ASCII_LETTER
    : ASCII_LETTER_UPPER
    | ASCII_LETTER_LOWER
;

fragment LETTER
    : ASCII_LETTER
    | '$'
    | '-'
    | '.'
    | '_'
;

fragment ESCAPE_LETTER
    : LETTER
    | '\\'
;

fragment DECIMAL_DIGIT
    : '0'..'9'
;

fragment HEX_DIGIT
    : DECIMAL_DIGIT
    | 'A'..'F'
    | 'a'..'f'
;

COMMENT : SEMICOLON .*? '\r' ? '\n' -> skip ;
WHITESPACE : [ \t\r\n]+ -> skip ;

// Reserved words

DECLARE : 'declare' ;
DEFINE : 'define' ;
PRIVATE : 'private' ;
NOUNWIND : 'nounwind' ;
NOCAPTURE : 'nocapture' ;
UNNAMMED_ADDR : 'unnamed_addr' ;
NULL : 'null' ;
GLOBAL : 'global' ;
CONSTANT: 'constant' ;

// Symbols

LPAREN : '(' ;
RPAREN : '(' ;
LBRACK : '[' ;
RBRACK : ']' ;
LBRACE : '{' ;
RBRACE : '}' ;
EQUAL : '=' ;
DOUBLEQUOTE : '"' ;
COMMA : ',' ;
STAR : '*' ;
PERCENT : '%' ;
SEMICOLON : ' ;' ;
EXCLAMATION : '!' ;
AT : '@' ;
DOT : '.' ;
MULTIPLY : 'x' ;
