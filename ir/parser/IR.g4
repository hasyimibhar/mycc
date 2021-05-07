grammar IR;

COMMENT : SEMICOLON ~[\r\n]* -> skip ;
WHITESPACE : [ \t\r\n]+ -> skip ;

ir 
    : statSeq
    ;

// Statement

statSeq
    : stat statSeqTail
    ;

statSeqTail
    : statSeq
    |
    ;

stat
    : statGlobalDef
    | statFunctDecl
    | statFunctDef
    ;


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
RPAREN : ')' ;
LBRACK : '[' ;
RBRACK : ']' ;
LBRACE : '{' ;
RBRACE : '}' ;
EQUAL : '=' ;
DOUBLEQUOTE : '"' ;
COMMA : ',' ;
STAR : '*' ;
PERCENT : '%' ;
SEMICOLON : ';' ;
EXCLAMATION : '!' ;
AT : '@' ;
DOT : '.' ;
MULTIPLY : 'x' ;
MINUS : '-' ;

// Type

type_
    : singleValueType
    | arrayType
    | pointerType
    ;

singleValueType
    : TYPE_I8
    | TYPE_I32
    | TYPE_I64
    ;

arrayType
    : LBRACK DECIMALS MULTIPLY type_ RBRACK
    ;

pointerType
    : (singleValueType | arrayType) STAR
    ;

TYPE_I8 : 'i8' ;
TYPE_I32 : 'i32' ;
TYPE_I64 : 'i64' ;

visibility
    : PRIVATE
    ;

// Identifiers

NAME
    : LETTER ( LETTER | DECIMAL_DIGIT ) *
    ;

// Literals

DECIMALS
    : DECIMAL_DIGIT +
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



// Local id

localIdent
    : localName
    ;

// LOCAL_ID
//     : PERCENT ID
//     ;

localName
    : PERCENT NAME
    ;

// Global variable

globalIdent
    : globalName
    ;

globalName
    : AT DOT? NAME
    ;

statGlobalDef
    : globalIdent EQUAL visibility UNNAMMED_ADDR? (GLOBAL | CONSTANT)? type_ exprConstant
    ;

// Function

statFunctDecl
    : DECLARE functSignature
    ;

statFunctDef
    : DEFINE functSignature LBRACE instructionSeq RBRACE
    ;

functSignature
    : type_ globalName LPAREN functDeclArgsList RPAREN functAttribsList
    ;
 
functDeclArgsList
    : functDeclArg functDeclArgsListTail
    |
    ;

functDeclArgsListTail
    : COMMA functDeclArgsList
    |
    ;

functDeclArg
    : type_ parameterAttrib?
    ;

functAttribsList
    : functAttrib functAttribsList
    |
    ;

functAttrib
    : NOUNWIND
    ;

parameterAttrib
    : NOCAPTURE
    ;

functCallArgsList
    : functCallArg functCallArgsListTail
    |
    ;

functCallArgsListTail
    : COMMA functCallArgsList
    |
    ;

functCallArg
    : type_ localIdent
    ;

// Instruction

instructionSeq
    : instruction instructionSeqTail
    ;

instructionSeqTail
    : instructionSeq
    |
    ;

instruction
    : instructionRet
    | instructionCall
    | instructionGetelementptr
    ;

instructionRet
    : 'ret' type_ value
    ;

instructionCall
    : 'call' type_ globalName LPAREN functCallArgsList RPAREN
    ;

instructionGetelementptr
    : localIdent EQUAL 'getelementptr' type_ COMMA pointerType (globalIdent | localIdent) instructionGetelementptrTail
    ;

instructionGetelementptrTail
    : COMMA type_ DECIMALS instructionGetelementptrTail
    |
    ;

// Metadata

// We ignore metadata for now
METADATA : EXCLAMATION ~[\r\n]* -> skip ;
 
// METADATA_IDENT
//     : METADATA_NAME
//     | METADATA_ID
//     ;
// 
// METADATA_ID
//     : '!' ID
//     ;
// 
// METADATA_NAME
//     : '!' NAME
//     ;
// 

// Expression

value
    : DECIMALS
    ;
 
exprConstant
    : STRING_LIT
    ;

fragment ESCAPED_QUOTE : '\\"';
QUOTED_STRING :   '"' ( ESCAPED_QUOTE | ~('\n'|'\r') )*? '"';

STRING_LIT
    : 'c' QUOTED_STRING
    ;
