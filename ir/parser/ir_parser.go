// Code generated from IR.g4 by ANTLR 4.9. DO NOT EDIT.

package parser // IR

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 41, 235,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 4,
	3, 4, 5, 4, 80, 10, 4, 3, 5, 3, 5, 3, 5, 5, 5, 85, 10, 5, 3, 6, 3, 6, 3,
	6, 5, 6, 90, 10, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3,
	9, 3, 9, 5, 9, 102, 10, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3, 11, 3, 12,
	3, 12, 3, 12, 3, 13, 3, 13, 3, 14, 3, 14, 5, 14, 117, 10, 14, 3, 14, 3,
	14, 3, 15, 3, 15, 3, 15, 3, 15, 5, 15, 125, 10, 15, 3, 15, 5, 15, 128,
	10, 15, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17,
	3, 17, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3,
	19, 3, 19, 3, 19, 3, 19, 5, 19, 153, 10, 19, 3, 20, 3, 20, 3, 20, 5, 20,
	158, 10, 20, 3, 21, 3, 21, 5, 21, 162, 10, 21, 3, 22, 3, 22, 3, 22, 3,
	22, 5, 22, 168, 10, 22, 3, 23, 3, 23, 3, 24, 3, 24, 3, 25, 3, 25, 3, 25,
	3, 25, 5, 25, 178, 10, 25, 3, 26, 3, 26, 3, 26, 5, 26, 183, 10, 26, 3,
	27, 3, 27, 3, 27, 3, 28, 3, 28, 3, 28, 3, 29, 3, 29, 5, 29, 193, 10, 29,
	3, 30, 3, 30, 3, 30, 5, 30, 198, 10, 30, 3, 31, 3, 31, 3, 31, 3, 31, 3,
	32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 33, 3, 33, 3, 33, 3, 33,
	3, 33, 3, 33, 3, 33, 3, 33, 5, 33, 219, 10, 33, 3, 33, 3, 33, 3, 34, 3,
	34, 3, 34, 3, 34, 3, 34, 3, 34, 5, 34, 229, 10, 34, 3, 35, 3, 35, 3, 36,
	3, 36, 3, 36, 2, 2, 37, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26,
	28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62,
	64, 66, 68, 70, 2, 4, 3, 2, 34, 36, 3, 2, 15, 16, 2, 219, 2, 72, 3, 2,
	2, 2, 4, 74, 3, 2, 2, 2, 6, 79, 3, 2, 2, 2, 8, 84, 3, 2, 2, 2, 10, 89,
	3, 2, 2, 2, 12, 91, 3, 2, 2, 2, 14, 93, 3, 2, 2, 2, 16, 101, 3, 2, 2, 2,
	18, 105, 3, 2, 2, 2, 20, 107, 3, 2, 2, 2, 22, 109, 3, 2, 2, 2, 24, 112,
	3, 2, 2, 2, 26, 114, 3, 2, 2, 2, 28, 120, 3, 2, 2, 2, 30, 132, 3, 2, 2,
	2, 32, 135, 3, 2, 2, 2, 34, 141, 3, 2, 2, 2, 36, 152, 3, 2, 2, 2, 38, 157,
	3, 2, 2, 2, 40, 159, 3, 2, 2, 2, 42, 167, 3, 2, 2, 2, 44, 169, 3, 2, 2,
	2, 46, 171, 3, 2, 2, 2, 48, 177, 3, 2, 2, 2, 50, 182, 3, 2, 2, 2, 52, 184,
	3, 2, 2, 2, 54, 187, 3, 2, 2, 2, 56, 192, 3, 2, 2, 2, 58, 197, 3, 2, 2,
	2, 60, 199, 3, 2, 2, 2, 62, 203, 3, 2, 2, 2, 64, 210, 3, 2, 2, 2, 66, 228,
	3, 2, 2, 2, 68, 230, 3, 2, 2, 2, 70, 232, 3, 2, 2, 2, 72, 73, 5, 4, 3,
	2, 73, 3, 3, 2, 2, 2, 74, 75, 5, 8, 5, 2, 75, 76, 5, 6, 4, 2, 76, 5, 3,
	2, 2, 2, 77, 80, 5, 4, 3, 2, 78, 80, 3, 2, 2, 2, 79, 77, 3, 2, 2, 2, 79,
	78, 3, 2, 2, 2, 80, 7, 3, 2, 2, 2, 81, 85, 5, 28, 15, 2, 82, 85, 5, 30,
	16, 2, 83, 85, 5, 32, 17, 2, 84, 81, 3, 2, 2, 2, 84, 82, 3, 2, 2, 2, 84,
	83, 3, 2, 2, 2, 85, 9, 3, 2, 2, 2, 86, 90, 5, 12, 7, 2, 87, 90, 5, 14,
	8, 2, 88, 90, 5, 16, 9, 2, 89, 86, 3, 2, 2, 2, 89, 87, 3, 2, 2, 2, 89,
	88, 3, 2, 2, 2, 90, 11, 3, 2, 2, 2, 91, 92, 9, 2, 2, 2, 92, 13, 3, 2, 2,
	2, 93, 94, 7, 19, 2, 2, 94, 95, 7, 38, 2, 2, 95, 96, 7, 32, 2, 2, 96, 97,
	5, 10, 6, 2, 97, 98, 7, 20, 2, 2, 98, 15, 3, 2, 2, 2, 99, 102, 5, 12, 7,
	2, 100, 102, 5, 14, 8, 2, 101, 99, 3, 2, 2, 2, 101, 100, 3, 2, 2, 2, 102,
	103, 3, 2, 2, 2, 103, 104, 7, 26, 2, 2, 104, 17, 3, 2, 2, 2, 105, 106,
	7, 10, 2, 2, 106, 19, 3, 2, 2, 2, 107, 108, 5, 22, 12, 2, 108, 21, 3, 2,
	2, 2, 109, 110, 7, 27, 2, 2, 110, 111, 7, 37, 2, 2, 111, 23, 3, 2, 2, 2,
	112, 113, 5, 26, 14, 2, 113, 25, 3, 2, 2, 2, 114, 116, 7, 30, 2, 2, 115,
	117, 7, 31, 2, 2, 116, 115, 3, 2, 2, 2, 116, 117, 3, 2, 2, 2, 117, 118,
	3, 2, 2, 2, 118, 119, 7, 37, 2, 2, 119, 27, 3, 2, 2, 2, 120, 121, 5, 24,
	13, 2, 121, 122, 7, 23, 2, 2, 122, 124, 5, 18, 10, 2, 123, 125, 7, 13,
	2, 2, 124, 123, 3, 2, 2, 2, 124, 125, 3, 2, 2, 2, 125, 127, 3, 2, 2, 2,
	126, 128, 9, 3, 2, 2, 127, 126, 3, 2, 2, 2, 127, 128, 3, 2, 2, 2, 128,
	129, 3, 2, 2, 2, 129, 130, 5, 10, 6, 2, 130, 131, 5, 70, 36, 2, 131, 29,
	3, 2, 2, 2, 132, 133, 7, 8, 2, 2, 133, 134, 5, 34, 18, 2, 134, 31, 3, 2,
	2, 2, 135, 136, 7, 9, 2, 2, 136, 137, 5, 34, 18, 2, 137, 138, 7, 21, 2,
	2, 138, 139, 5, 54, 28, 2, 139, 140, 7, 22, 2, 2, 140, 33, 3, 2, 2, 2,
	141, 142, 5, 10, 6, 2, 142, 143, 5, 26, 14, 2, 143, 144, 7, 17, 2, 2, 144,
	145, 5, 36, 19, 2, 145, 146, 7, 18, 2, 2, 146, 147, 5, 42, 22, 2, 147,
	35, 3, 2, 2, 2, 148, 149, 5, 40, 21, 2, 149, 150, 5, 38, 20, 2, 150, 153,
	3, 2, 2, 2, 151, 153, 3, 2, 2, 2, 152, 148, 3, 2, 2, 2, 152, 151, 3, 2,
	2, 2, 153, 37, 3, 2, 2, 2, 154, 155, 7, 25, 2, 2, 155, 158, 5, 36, 19,
	2, 156, 158, 3, 2, 2, 2, 157, 154, 3, 2, 2, 2, 157, 156, 3, 2, 2, 2, 158,
	39, 3, 2, 2, 2, 159, 161, 5, 10, 6, 2, 160, 162, 5, 46, 24, 2, 161, 160,
	3, 2, 2, 2, 161, 162, 3, 2, 2, 2, 162, 41, 3, 2, 2, 2, 163, 164, 5, 44,
	23, 2, 164, 165, 5, 42, 22, 2, 165, 168, 3, 2, 2, 2, 166, 168, 3, 2, 2,
	2, 167, 163, 3, 2, 2, 2, 167, 166, 3, 2, 2, 2, 168, 43, 3, 2, 2, 2, 169,
	170, 7, 11, 2, 2, 170, 45, 3, 2, 2, 2, 171, 172, 7, 12, 2, 2, 172, 47,
	3, 2, 2, 2, 173, 174, 5, 52, 27, 2, 174, 175, 5, 50, 26, 2, 175, 178, 3,
	2, 2, 2, 176, 178, 3, 2, 2, 2, 177, 173, 3, 2, 2, 2, 177, 176, 3, 2, 2,
	2, 178, 49, 3, 2, 2, 2, 179, 180, 7, 25, 2, 2, 180, 183, 5, 48, 25, 2,
	181, 183, 3, 2, 2, 2, 182, 179, 3, 2, 2, 2, 182, 181, 3, 2, 2, 2, 183,
	51, 3, 2, 2, 2, 184, 185, 5, 10, 6, 2, 185, 186, 5, 20, 11, 2, 186, 53,
	3, 2, 2, 2, 187, 188, 5, 58, 30, 2, 188, 189, 5, 56, 29, 2, 189, 55, 3,
	2, 2, 2, 190, 193, 5, 54, 28, 2, 191, 193, 3, 2, 2, 2, 192, 190, 3, 2,
	2, 2, 192, 191, 3, 2, 2, 2, 193, 57, 3, 2, 2, 2, 194, 198, 5, 60, 31, 2,
	195, 198, 5, 62, 32, 2, 196, 198, 5, 64, 33, 2, 197, 194, 3, 2, 2, 2, 197,
	195, 3, 2, 2, 2, 197, 196, 3, 2, 2, 2, 198, 59, 3, 2, 2, 2, 199, 200, 7,
	3, 2, 2, 200, 201, 5, 10, 6, 2, 201, 202, 5, 68, 35, 2, 202, 61, 3, 2,
	2, 2, 203, 204, 7, 4, 2, 2, 204, 205, 5, 10, 6, 2, 205, 206, 5, 26, 14,
	2, 206, 207, 7, 17, 2, 2, 207, 208, 5, 48, 25, 2, 208, 209, 7, 18, 2, 2,
	209, 63, 3, 2, 2, 2, 210, 211, 5, 20, 11, 2, 211, 212, 7, 23, 2, 2, 212,
	213, 7, 5, 2, 2, 213, 214, 5, 10, 6, 2, 214, 215, 7, 25, 2, 2, 215, 218,
	5, 16, 9, 2, 216, 219, 5, 24, 13, 2, 217, 219, 5, 20, 11, 2, 218, 216,
	3, 2, 2, 2, 218, 217, 3, 2, 2, 2, 219, 220, 3, 2, 2, 2, 220, 221, 5, 66,
	34, 2, 221, 65, 3, 2, 2, 2, 222, 223, 7, 25, 2, 2, 223, 224, 5, 10, 6,
	2, 224, 225, 7, 38, 2, 2, 225, 226, 5, 66, 34, 2, 226, 229, 3, 2, 2, 2,
	227, 229, 3, 2, 2, 2, 228, 222, 3, 2, 2, 2, 228, 227, 3, 2, 2, 2, 229,
	67, 3, 2, 2, 2, 230, 231, 7, 38, 2, 2, 231, 69, 3, 2, 2, 2, 232, 233, 7,
	41, 2, 2, 233, 71, 3, 2, 2, 2, 19, 79, 84, 89, 101, 116, 124, 127, 152,
	157, 161, 167, 177, 182, 192, 197, 218, 228,
}
var literalNames = []string{
	"", "'ret'", "'call'", "'getelementptr'", "", "", "'declare'", "'define'",
	"'private'", "'nounwind'", "'nocapture'", "'unnamed_addr'", "'null'", "'global'",
	"'constant'", "'('", "')'", "'['", "']'", "'{'", "'}'", "'='", "'\"'",
	"','", "'*'", "'%'", "';'", "'!'", "'@'", "'.'", "'x'", "'-'", "'i8'",
	"'i32'", "'i64'",
}
var symbolicNames = []string{
	"", "", "", "", "COMMENT", "WHITESPACE", "DECLARE", "DEFINE", "PRIVATE",
	"NOUNWIND", "NOCAPTURE", "UNNAMMED_ADDR", "NULL", "GLOBAL", "CONSTANT",
	"LPAREN", "RPAREN", "LBRACK", "RBRACK", "LBRACE", "RBRACE", "EQUAL", "DOUBLEQUOTE",
	"COMMA", "STAR", "PERCENT", "SEMICOLON", "EXCLAMATION", "AT", "DOT", "MULTIPLY",
	"MINUS", "TYPE_I8", "TYPE_I32", "TYPE_I64", "NAME", "DECIMALS", "METADATA",
	"QUOTED_STRING", "STRING_LIT",
}

var ruleNames = []string{
	"ir", "statSeq", "statSeqTail", "stat", "type_", "singleValueType", "arrayType",
	"pointerType", "visibility", "localIdent", "localName", "globalIdent",
	"globalName", "statGlobalDef", "statFunctDecl", "statFunctDef", "functSignature",
	"functDeclArgsList", "functDeclArgsListTail", "functDeclArg", "functAttribsList",
	"functAttrib", "parameterAttrib", "functCallArgsList", "functCallArgsListTail",
	"functCallArg", "instructionSeq", "instructionSeqTail", "instruction",
	"instructionRet", "instructionCall", "instructionGetelementptr", "instructionGetelementptrTail",
	"value", "exprConstant",
}

type IRParser struct {
	*antlr.BaseParser
}

// NewIRParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *IRParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewIRParser(input antlr.TokenStream) *IRParser {
	this := new(IRParser)
	deserializer := antlr.NewATNDeserializer(nil)
	deserializedATN := deserializer.DeserializeFromUInt16(parserATN)
	decisionToDFA := make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "IR.g4"

	return this
}

// IRParser tokens.
const (
	IRParserEOF           = antlr.TokenEOF
	IRParserT__0          = 1
	IRParserT__1          = 2
	IRParserT__2          = 3
	IRParserCOMMENT       = 4
	IRParserWHITESPACE    = 5
	IRParserDECLARE       = 6
	IRParserDEFINE        = 7
	IRParserPRIVATE       = 8
	IRParserNOUNWIND      = 9
	IRParserNOCAPTURE     = 10
	IRParserUNNAMMED_ADDR = 11
	IRParserNULL          = 12
	IRParserGLOBAL        = 13
	IRParserCONSTANT      = 14
	IRParserLPAREN        = 15
	IRParserRPAREN        = 16
	IRParserLBRACK        = 17
	IRParserRBRACK        = 18
	IRParserLBRACE        = 19
	IRParserRBRACE        = 20
	IRParserEQUAL         = 21
	IRParserDOUBLEQUOTE   = 22
	IRParserCOMMA         = 23
	IRParserSTAR          = 24
	IRParserPERCENT       = 25
	IRParserSEMICOLON     = 26
	IRParserEXCLAMATION   = 27
	IRParserAT            = 28
	IRParserDOT           = 29
	IRParserMULTIPLY      = 30
	IRParserMINUS         = 31
	IRParserTYPE_I8       = 32
	IRParserTYPE_I32      = 33
	IRParserTYPE_I64      = 34
	IRParserNAME          = 35
	IRParserDECIMALS      = 36
	IRParserMETADATA      = 37
	IRParserQUOTED_STRING = 38
	IRParserSTRING_LIT    = 39
)

// IRParser rules.
const (
	IRParserRULE_ir                           = 0
	IRParserRULE_statSeq                      = 1
	IRParserRULE_statSeqTail                  = 2
	IRParserRULE_stat                         = 3
	IRParserRULE_type_                        = 4
	IRParserRULE_singleValueType              = 5
	IRParserRULE_arrayType                    = 6
	IRParserRULE_pointerType                  = 7
	IRParserRULE_visibility                   = 8
	IRParserRULE_localIdent                   = 9
	IRParserRULE_localName                    = 10
	IRParserRULE_globalIdent                  = 11
	IRParserRULE_globalName                   = 12
	IRParserRULE_statGlobalDef                = 13
	IRParserRULE_statFunctDecl                = 14
	IRParserRULE_statFunctDef                 = 15
	IRParserRULE_functSignature               = 16
	IRParserRULE_functDeclArgsList            = 17
	IRParserRULE_functDeclArgsListTail        = 18
	IRParserRULE_functDeclArg                 = 19
	IRParserRULE_functAttribsList             = 20
	IRParserRULE_functAttrib                  = 21
	IRParserRULE_parameterAttrib              = 22
	IRParserRULE_functCallArgsList            = 23
	IRParserRULE_functCallArgsListTail        = 24
	IRParserRULE_functCallArg                 = 25
	IRParserRULE_instructionSeq               = 26
	IRParserRULE_instructionSeqTail           = 27
	IRParserRULE_instruction                  = 28
	IRParserRULE_instructionRet               = 29
	IRParserRULE_instructionCall              = 30
	IRParserRULE_instructionGetelementptr     = 31
	IRParserRULE_instructionGetelementptrTail = 32
	IRParserRULE_value                        = 33
	IRParserRULE_exprConstant                 = 34
)

// IIrContext is an interface to support dynamic dispatch.
type IIrContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIrContext differentiates from other interfaces.
	IsIrContext()
}

type IrContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIrContext() *IrContext {
	var p = new(IrContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IRParserRULE_ir
	return p
}

func (*IrContext) IsIrContext() {}

func NewIrContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IrContext {
	var p = new(IrContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IRParserRULE_ir

	return p
}

func (s *IrContext) GetParser() antlr.Parser { return s.parser }

func (s *IrContext) StatSeq() IStatSeqContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatSeqContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatSeqContext)
}

func (s *IrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IrContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRListener); ok {
		listenerT.EnterIr(s)
	}
}

func (s *IrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRListener); ok {
		listenerT.ExitIr(s)
	}
}

func (p *IRParser) Ir() (localctx IIrContext) {
	localctx = NewIrContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, IRParserRULE_ir)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(70)
		p.StatSeq()
	}

	return localctx
}

// IStatSeqContext is an interface to support dynamic dispatch.
type IStatSeqContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatSeqContext differentiates from other interfaces.
	IsStatSeqContext()
}

type StatSeqContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatSeqContext() *StatSeqContext {
	var p = new(StatSeqContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IRParserRULE_statSeq
	return p
}

func (*StatSeqContext) IsStatSeqContext() {}

func NewStatSeqContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatSeqContext {
	var p = new(StatSeqContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IRParserRULE_statSeq

	return p
}

func (s *StatSeqContext) GetParser() antlr.Parser { return s.parser }

func (s *StatSeqContext) Stat() IStatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatContext)
}

func (s *StatSeqContext) StatSeqTail() IStatSeqTailContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatSeqTailContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatSeqTailContext)
}

func (s *StatSeqContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatSeqContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatSeqContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRListener); ok {
		listenerT.EnterStatSeq(s)
	}
}

func (s *StatSeqContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRListener); ok {
		listenerT.ExitStatSeq(s)
	}
}

func (p *IRParser) StatSeq() (localctx IStatSeqContext) {
	localctx = NewStatSeqContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, IRParserRULE_statSeq)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(72)
		p.Stat()
	}
	{
		p.SetState(73)
		p.StatSeqTail()
	}

	return localctx
}

// IStatSeqTailContext is an interface to support dynamic dispatch.
type IStatSeqTailContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatSeqTailContext differentiates from other interfaces.
	IsStatSeqTailContext()
}

type StatSeqTailContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatSeqTailContext() *StatSeqTailContext {
	var p = new(StatSeqTailContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IRParserRULE_statSeqTail
	return p
}

func (*StatSeqTailContext) IsStatSeqTailContext() {}

func NewStatSeqTailContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatSeqTailContext {
	var p = new(StatSeqTailContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IRParserRULE_statSeqTail

	return p
}

func (s *StatSeqTailContext) GetParser() antlr.Parser { return s.parser }

func (s *StatSeqTailContext) StatSeq() IStatSeqContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatSeqContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatSeqContext)
}

func (s *StatSeqTailContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatSeqTailContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatSeqTailContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRListener); ok {
		listenerT.EnterStatSeqTail(s)
	}
}

func (s *StatSeqTailContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRListener); ok {
		listenerT.ExitStatSeqTail(s)
	}
}

func (p *IRParser) StatSeqTail() (localctx IStatSeqTailContext) {
	localctx = NewStatSeqTailContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, IRParserRULE_statSeqTail)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(77)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case IRParserDECLARE, IRParserDEFINE, IRParserAT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(75)
			p.StatSeq()
		}

	case IRParserEOF:
		p.EnterOuterAlt(localctx, 2)

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IStatContext is an interface to support dynamic dispatch.
type IStatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatContext differentiates from other interfaces.
	IsStatContext()
}

type StatContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatContext() *StatContext {
	var p = new(StatContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IRParserRULE_stat
	return p
}

func (*StatContext) IsStatContext() {}

func NewStatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatContext {
	var p = new(StatContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IRParserRULE_stat

	return p
}

func (s *StatContext) GetParser() antlr.Parser { return s.parser }

func (s *StatContext) StatGlobalDef() IStatGlobalDefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatGlobalDefContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatGlobalDefContext)
}

func (s *StatContext) StatFunctDecl() IStatFunctDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatFunctDeclContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatFunctDeclContext)
}

func (s *StatContext) StatFunctDef() IStatFunctDefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatFunctDefContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatFunctDefContext)
}

func (s *StatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRListener); ok {
		listenerT.EnterStat(s)
	}
}

func (s *StatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRListener); ok {
		listenerT.ExitStat(s)
	}
}

func (p *IRParser) Stat() (localctx IStatContext) {
	localctx = NewStatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, IRParserRULE_stat)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(82)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case IRParserAT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(79)
			p.StatGlobalDef()
		}

	case IRParserDECLARE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(80)
			p.StatFunctDecl()
		}

	case IRParserDEFINE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(81)
			p.StatFunctDef()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IType_Context is an interface to support dynamic dispatch.
type IType_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsType_Context differentiates from other interfaces.
	IsType_Context()
}

type Type_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyType_Context() *Type_Context {
	var p = new(Type_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IRParserRULE_type_
	return p
}

func (*Type_Context) IsType_Context() {}

func NewType_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Type_Context {
	var p = new(Type_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IRParserRULE_type_

	return p
}

func (s *Type_Context) GetParser() antlr.Parser { return s.parser }

func (s *Type_Context) SingleValueType() ISingleValueTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleValueTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISingleValueTypeContext)
}

func (s *Type_Context) ArrayType() IArrayTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArrayTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArrayTypeContext)
}

func (s *Type_Context) PointerType() IPointerTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPointerTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPointerTypeContext)
}

func (s *Type_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Type_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Type_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRListener); ok {
		listenerT.EnterType_(s)
	}
}

func (s *Type_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRListener); ok {
		listenerT.ExitType_(s)
	}
}

func (p *IRParser) Type_() (localctx IType_Context) {
	localctx = NewType_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, IRParserRULE_type_)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(87)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(84)
			p.SingleValueType()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(85)
			p.ArrayType()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(86)
			p.PointerType()
		}

	}

	return localctx
}

// ISingleValueTypeContext is an interface to support dynamic dispatch.
type ISingleValueTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSingleValueTypeContext differentiates from other interfaces.
	IsSingleValueTypeContext()
}

type SingleValueTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySingleValueTypeContext() *SingleValueTypeContext {
	var p = new(SingleValueTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IRParserRULE_singleValueType
	return p
}

func (*SingleValueTypeContext) IsSingleValueTypeContext() {}

func NewSingleValueTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SingleValueTypeContext {
	var p = new(SingleValueTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IRParserRULE_singleValueType

	return p
}

func (s *SingleValueTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *SingleValueTypeContext) TYPE_I8() antlr.TerminalNode {
	return s.GetToken(IRParserTYPE_I8, 0)
}

func (s *SingleValueTypeContext) TYPE_I32() antlr.TerminalNode {
	return s.GetToken(IRParserTYPE_I32, 0)
}

func (s *SingleValueTypeContext) TYPE_I64() antlr.TerminalNode {
	return s.GetToken(IRParserTYPE_I64, 0)
}

func (s *SingleValueTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SingleValueTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SingleValueTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRListener); ok {
		listenerT.EnterSingleValueType(s)
	}
}

func (s *SingleValueTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRListener); ok {
		listenerT.ExitSingleValueType(s)
	}
}

func (p *IRParser) SingleValueType() (localctx ISingleValueTypeContext) {
	localctx = NewSingleValueTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, IRParserRULE_singleValueType)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(89)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(IRParserTYPE_I8-32))|(1<<(IRParserTYPE_I32-32))|(1<<(IRParserTYPE_I64-32)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IArrayTypeContext is an interface to support dynamic dispatch.
type IArrayTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArrayTypeContext differentiates from other interfaces.
	IsArrayTypeContext()
}

type ArrayTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayTypeContext() *ArrayTypeContext {
	var p = new(ArrayTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IRParserRULE_arrayType
	return p
}

func (*ArrayTypeContext) IsArrayTypeContext() {}

func NewArrayTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayTypeContext {
	var p = new(ArrayTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IRParserRULE_arrayType

	return p
}

func (s *ArrayTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayTypeContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(IRParserLBRACK, 0)
}

func (s *ArrayTypeContext) DECIMALS() antlr.TerminalNode {
	return s.GetToken(IRParserDECIMALS, 0)
}

func (s *ArrayTypeContext) MULTIPLY() antlr.TerminalNode {
	return s.GetToken(IRParserMULTIPLY, 0)
}

func (s *ArrayTypeContext) Type_() IType_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *ArrayTypeContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(IRParserRBRACK, 0)
}

func (s *ArrayTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRListener); ok {
		listenerT.EnterArrayType(s)
	}
}

func (s *ArrayTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRListener); ok {
		listenerT.ExitArrayType(s)
	}
}

func (p *IRParser) ArrayType() (localctx IArrayTypeContext) {
	localctx = NewArrayTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, IRParserRULE_arrayType)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(91)
		p.Match(IRParserLBRACK)
	}
	{
		p.SetState(92)
		p.Match(IRParserDECIMALS)
	}
	{
		p.SetState(93)
		p.Match(IRParserMULTIPLY)
	}
	{
		p.SetState(94)
		p.Type_()
	}
	{
		p.SetState(95)
		p.Match(IRParserRBRACK)
	}

	return localctx
}

// IPointerTypeContext is an interface to support dynamic dispatch.
type IPointerTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPointerTypeContext differentiates from other interfaces.
	IsPointerTypeContext()
}

type PointerTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPointerTypeContext() *PointerTypeContext {
	var p = new(PointerTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IRParserRULE_pointerType
	return p
}

func (*PointerTypeContext) IsPointerTypeContext() {}

func NewPointerTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PointerTypeContext {
	var p = new(PointerTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IRParserRULE_pointerType

	return p
}

func (s *PointerTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *PointerTypeContext) STAR() antlr.TerminalNode {
	return s.GetToken(IRParserSTAR, 0)
}

func (s *PointerTypeContext) SingleValueType() ISingleValueTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleValueTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISingleValueTypeContext)
}

func (s *PointerTypeContext) ArrayType() IArrayTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArrayTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArrayTypeContext)
}

func (s *PointerTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PointerTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PointerTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRListener); ok {
		listenerT.EnterPointerType(s)
	}
}

func (s *PointerTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRListener); ok {
		listenerT.ExitPointerType(s)
	}
}

func (p *IRParser) PointerType() (localctx IPointerTypeContext) {
	localctx = NewPointerTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, IRParserRULE_pointerType)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(99)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case IRParserTYPE_I8, IRParserTYPE_I32, IRParserTYPE_I64:
		{
			p.SetState(97)
			p.SingleValueType()
		}

	case IRParserLBRACK:
		{
			p.SetState(98)
			p.ArrayType()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(101)
		p.Match(IRParserSTAR)
	}

	return localctx
}

// IVisibilityContext is an interface to support dynamic dispatch.
type IVisibilityContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVisibilityContext differentiates from other interfaces.
	IsVisibilityContext()
}

type VisibilityContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVisibilityContext() *VisibilityContext {
	var p = new(VisibilityContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IRParserRULE_visibility
	return p
}

func (*VisibilityContext) IsVisibilityContext() {}

func NewVisibilityContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VisibilityContext {
	var p = new(VisibilityContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IRParserRULE_visibility

	return p
}

func (s *VisibilityContext) GetParser() antlr.Parser { return s.parser }

func (s *VisibilityContext) PRIVATE() antlr.TerminalNode {
	return s.GetToken(IRParserPRIVATE, 0)
}

func (s *VisibilityContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VisibilityContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VisibilityContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRListener); ok {
		listenerT.EnterVisibility(s)
	}
}

func (s *VisibilityContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRListener); ok {
		listenerT.ExitVisibility(s)
	}
}

func (p *IRParser) Visibility() (localctx IVisibilityContext) {
	localctx = NewVisibilityContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, IRParserRULE_visibility)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(103)
		p.Match(IRParserPRIVATE)
	}

	return localctx
}

// ILocalIdentContext is an interface to support dynamic dispatch.
type ILocalIdentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLocalIdentContext differentiates from other interfaces.
	IsLocalIdentContext()
}

type LocalIdentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLocalIdentContext() *LocalIdentContext {
	var p = new(LocalIdentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IRParserRULE_localIdent
	return p
}

func (*LocalIdentContext) IsLocalIdentContext() {}

func NewLocalIdentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LocalIdentContext {
	var p = new(LocalIdentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IRParserRULE_localIdent

	return p
}

func (s *LocalIdentContext) GetParser() antlr.Parser { return s.parser }

func (s *LocalIdentContext) LocalName() ILocalNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILocalNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILocalNameContext)
}

func (s *LocalIdentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LocalIdentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LocalIdentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRListener); ok {
		listenerT.EnterLocalIdent(s)
	}
}

func (s *LocalIdentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRListener); ok {
		listenerT.ExitLocalIdent(s)
	}
}

func (p *IRParser) LocalIdent() (localctx ILocalIdentContext) {
	localctx = NewLocalIdentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, IRParserRULE_localIdent)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(105)
		p.LocalName()
	}

	return localctx
}

// ILocalNameContext is an interface to support dynamic dispatch.
type ILocalNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLocalNameContext differentiates from other interfaces.
	IsLocalNameContext()
}

type LocalNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLocalNameContext() *LocalNameContext {
	var p = new(LocalNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IRParserRULE_localName
	return p
}

func (*LocalNameContext) IsLocalNameContext() {}

func NewLocalNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LocalNameContext {
	var p = new(LocalNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IRParserRULE_localName

	return p
}

func (s *LocalNameContext) GetParser() antlr.Parser { return s.parser }

func (s *LocalNameContext) PERCENT() antlr.TerminalNode {
	return s.GetToken(IRParserPERCENT, 0)
}

func (s *LocalNameContext) NAME() antlr.TerminalNode {
	return s.GetToken(IRParserNAME, 0)
}

func (s *LocalNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LocalNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LocalNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRListener); ok {
		listenerT.EnterLocalName(s)
	}
}

func (s *LocalNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRListener); ok {
		listenerT.ExitLocalName(s)
	}
}

func (p *IRParser) LocalName() (localctx ILocalNameContext) {
	localctx = NewLocalNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, IRParserRULE_localName)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(107)
		p.Match(IRParserPERCENT)
	}
	{
		p.SetState(108)
		p.Match(IRParserNAME)
	}

	return localctx
}

// IGlobalIdentContext is an interface to support dynamic dispatch.
type IGlobalIdentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGlobalIdentContext differentiates from other interfaces.
	IsGlobalIdentContext()
}

type GlobalIdentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGlobalIdentContext() *GlobalIdentContext {
	var p = new(GlobalIdentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IRParserRULE_globalIdent
	return p
}

func (*GlobalIdentContext) IsGlobalIdentContext() {}

func NewGlobalIdentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GlobalIdentContext {
	var p = new(GlobalIdentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IRParserRULE_globalIdent

	return p
}

func (s *GlobalIdentContext) GetParser() antlr.Parser { return s.parser }

func (s *GlobalIdentContext) GlobalName() IGlobalNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGlobalNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGlobalNameContext)
}

func (s *GlobalIdentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GlobalIdentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GlobalIdentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRListener); ok {
		listenerT.EnterGlobalIdent(s)
	}
}

func (s *GlobalIdentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRListener); ok {
		listenerT.ExitGlobalIdent(s)
	}
}

func (p *IRParser) GlobalIdent() (localctx IGlobalIdentContext) {
	localctx = NewGlobalIdentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, IRParserRULE_globalIdent)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(110)
		p.GlobalName()
	}

	return localctx
}

// IGlobalNameContext is an interface to support dynamic dispatch.
type IGlobalNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGlobalNameContext differentiates from other interfaces.
	IsGlobalNameContext()
}

type GlobalNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGlobalNameContext() *GlobalNameContext {
	var p = new(GlobalNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IRParserRULE_globalName
	return p
}

func (*GlobalNameContext) IsGlobalNameContext() {}

func NewGlobalNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GlobalNameContext {
	var p = new(GlobalNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IRParserRULE_globalName

	return p
}

func (s *GlobalNameContext) GetParser() antlr.Parser { return s.parser }

func (s *GlobalNameContext) AT() antlr.TerminalNode {
	return s.GetToken(IRParserAT, 0)
}

func (s *GlobalNameContext) NAME() antlr.TerminalNode {
	return s.GetToken(IRParserNAME, 0)
}

func (s *GlobalNameContext) DOT() antlr.TerminalNode {
	return s.GetToken(IRParserDOT, 0)
}

func (s *GlobalNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GlobalNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GlobalNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRListener); ok {
		listenerT.EnterGlobalName(s)
	}
}

func (s *GlobalNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRListener); ok {
		listenerT.ExitGlobalName(s)
	}
}

func (p *IRParser) GlobalName() (localctx IGlobalNameContext) {
	localctx = NewGlobalNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, IRParserRULE_globalName)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(112)
		p.Match(IRParserAT)
	}
	p.SetState(114)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == IRParserDOT {
		{
			p.SetState(113)
			p.Match(IRParserDOT)
		}

	}
	{
		p.SetState(116)
		p.Match(IRParserNAME)
	}

	return localctx
}

// IStatGlobalDefContext is an interface to support dynamic dispatch.
type IStatGlobalDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatGlobalDefContext differentiates from other interfaces.
	IsStatGlobalDefContext()
}

type StatGlobalDefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatGlobalDefContext() *StatGlobalDefContext {
	var p = new(StatGlobalDefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IRParserRULE_statGlobalDef
	return p
}

func (*StatGlobalDefContext) IsStatGlobalDefContext() {}

func NewStatGlobalDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatGlobalDefContext {
	var p = new(StatGlobalDefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IRParserRULE_statGlobalDef

	return p
}

func (s *StatGlobalDefContext) GetParser() antlr.Parser { return s.parser }

func (s *StatGlobalDefContext) GlobalIdent() IGlobalIdentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGlobalIdentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGlobalIdentContext)
}

func (s *StatGlobalDefContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(IRParserEQUAL, 0)
}

func (s *StatGlobalDefContext) Visibility() IVisibilityContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVisibilityContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVisibilityContext)
}

func (s *StatGlobalDefContext) Type_() IType_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *StatGlobalDefContext) ExprConstant() IExprConstantContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprConstantContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprConstantContext)
}

func (s *StatGlobalDefContext) UNNAMMED_ADDR() antlr.TerminalNode {
	return s.GetToken(IRParserUNNAMMED_ADDR, 0)
}

func (s *StatGlobalDefContext) GLOBAL() antlr.TerminalNode {
	return s.GetToken(IRParserGLOBAL, 0)
}

func (s *StatGlobalDefContext) CONSTANT() antlr.TerminalNode {
	return s.GetToken(IRParserCONSTANT, 0)
}

func (s *StatGlobalDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatGlobalDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatGlobalDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRListener); ok {
		listenerT.EnterStatGlobalDef(s)
	}
}

func (s *StatGlobalDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRListener); ok {
		listenerT.ExitStatGlobalDef(s)
	}
}

func (p *IRParser) StatGlobalDef() (localctx IStatGlobalDefContext) {
	localctx = NewStatGlobalDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, IRParserRULE_statGlobalDef)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(118)
		p.GlobalIdent()
	}
	{
		p.SetState(119)
		p.Match(IRParserEQUAL)
	}
	{
		p.SetState(120)
		p.Visibility()
	}
	p.SetState(122)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == IRParserUNNAMMED_ADDR {
		{
			p.SetState(121)
			p.Match(IRParserUNNAMMED_ADDR)
		}

	}
	p.SetState(125)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == IRParserGLOBAL || _la == IRParserCONSTANT {
		{
			p.SetState(124)
			_la = p.GetTokenStream().LA(1)

			if !(_la == IRParserGLOBAL || _la == IRParserCONSTANT) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}
	{
		p.SetState(127)
		p.Type_()
	}
	{
		p.SetState(128)
		p.ExprConstant()
	}

	return localctx
}

// IStatFunctDeclContext is an interface to support dynamic dispatch.
type IStatFunctDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatFunctDeclContext differentiates from other interfaces.
	IsStatFunctDeclContext()
}

type StatFunctDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatFunctDeclContext() *StatFunctDeclContext {
	var p = new(StatFunctDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IRParserRULE_statFunctDecl
	return p
}

func (*StatFunctDeclContext) IsStatFunctDeclContext() {}

func NewStatFunctDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatFunctDeclContext {
	var p = new(StatFunctDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IRParserRULE_statFunctDecl

	return p
}

func (s *StatFunctDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *StatFunctDeclContext) DECLARE() antlr.TerminalNode {
	return s.GetToken(IRParserDECLARE, 0)
}

func (s *StatFunctDeclContext) FunctSignature() IFunctSignatureContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctSignatureContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctSignatureContext)
}

func (s *StatFunctDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatFunctDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatFunctDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRListener); ok {
		listenerT.EnterStatFunctDecl(s)
	}
}

func (s *StatFunctDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRListener); ok {
		listenerT.ExitStatFunctDecl(s)
	}
}

func (p *IRParser) StatFunctDecl() (localctx IStatFunctDeclContext) {
	localctx = NewStatFunctDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, IRParserRULE_statFunctDecl)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(130)
		p.Match(IRParserDECLARE)
	}
	{
		p.SetState(131)
		p.FunctSignature()
	}

	return localctx
}

// IStatFunctDefContext is an interface to support dynamic dispatch.
type IStatFunctDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatFunctDefContext differentiates from other interfaces.
	IsStatFunctDefContext()
}

type StatFunctDefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatFunctDefContext() *StatFunctDefContext {
	var p = new(StatFunctDefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IRParserRULE_statFunctDef
	return p
}

func (*StatFunctDefContext) IsStatFunctDefContext() {}

func NewStatFunctDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatFunctDefContext {
	var p = new(StatFunctDefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IRParserRULE_statFunctDef

	return p
}

func (s *StatFunctDefContext) GetParser() antlr.Parser { return s.parser }

func (s *StatFunctDefContext) DEFINE() antlr.TerminalNode {
	return s.GetToken(IRParserDEFINE, 0)
}

func (s *StatFunctDefContext) FunctSignature() IFunctSignatureContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctSignatureContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctSignatureContext)
}

func (s *StatFunctDefContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(IRParserLBRACE, 0)
}

func (s *StatFunctDefContext) InstructionSeq() IInstructionSeqContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstructionSeqContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstructionSeqContext)
}

func (s *StatFunctDefContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(IRParserRBRACE, 0)
}

func (s *StatFunctDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatFunctDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatFunctDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRListener); ok {
		listenerT.EnterStatFunctDef(s)
	}
}

func (s *StatFunctDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRListener); ok {
		listenerT.ExitStatFunctDef(s)
	}
}

func (p *IRParser) StatFunctDef() (localctx IStatFunctDefContext) {
	localctx = NewStatFunctDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, IRParserRULE_statFunctDef)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(133)
		p.Match(IRParserDEFINE)
	}
	{
		p.SetState(134)
		p.FunctSignature()
	}
	{
		p.SetState(135)
		p.Match(IRParserLBRACE)
	}
	{
		p.SetState(136)
		p.InstructionSeq()
	}
	{
		p.SetState(137)
		p.Match(IRParserRBRACE)
	}

	return localctx
}

// IFunctSignatureContext is an interface to support dynamic dispatch.
type IFunctSignatureContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctSignatureContext differentiates from other interfaces.
	IsFunctSignatureContext()
}

type FunctSignatureContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctSignatureContext() *FunctSignatureContext {
	var p = new(FunctSignatureContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IRParserRULE_functSignature
	return p
}

func (*FunctSignatureContext) IsFunctSignatureContext() {}

func NewFunctSignatureContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctSignatureContext {
	var p = new(FunctSignatureContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IRParserRULE_functSignature

	return p
}

func (s *FunctSignatureContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctSignatureContext) Type_() IType_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *FunctSignatureContext) GlobalName() IGlobalNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGlobalNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGlobalNameContext)
}

func (s *FunctSignatureContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(IRParserLPAREN, 0)
}

func (s *FunctSignatureContext) FunctDeclArgsList() IFunctDeclArgsListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctDeclArgsListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctDeclArgsListContext)
}

func (s *FunctSignatureContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(IRParserRPAREN, 0)
}

func (s *FunctSignatureContext) FunctAttribsList() IFunctAttribsListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctAttribsListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctAttribsListContext)
}

func (s *FunctSignatureContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctSignatureContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctSignatureContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRListener); ok {
		listenerT.EnterFunctSignature(s)
	}
}

func (s *FunctSignatureContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRListener); ok {
		listenerT.ExitFunctSignature(s)
	}
}

func (p *IRParser) FunctSignature() (localctx IFunctSignatureContext) {
	localctx = NewFunctSignatureContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, IRParserRULE_functSignature)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(139)
		p.Type_()
	}
	{
		p.SetState(140)
		p.GlobalName()
	}
	{
		p.SetState(141)
		p.Match(IRParserLPAREN)
	}
	{
		p.SetState(142)
		p.FunctDeclArgsList()
	}
	{
		p.SetState(143)
		p.Match(IRParserRPAREN)
	}
	{
		p.SetState(144)
		p.FunctAttribsList()
	}

	return localctx
}

// IFunctDeclArgsListContext is an interface to support dynamic dispatch.
type IFunctDeclArgsListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctDeclArgsListContext differentiates from other interfaces.
	IsFunctDeclArgsListContext()
}

type FunctDeclArgsListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctDeclArgsListContext() *FunctDeclArgsListContext {
	var p = new(FunctDeclArgsListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IRParserRULE_functDeclArgsList
	return p
}

func (*FunctDeclArgsListContext) IsFunctDeclArgsListContext() {}

func NewFunctDeclArgsListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctDeclArgsListContext {
	var p = new(FunctDeclArgsListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IRParserRULE_functDeclArgsList

	return p
}

func (s *FunctDeclArgsListContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctDeclArgsListContext) FunctDeclArg() IFunctDeclArgContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctDeclArgContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctDeclArgContext)
}

func (s *FunctDeclArgsListContext) FunctDeclArgsListTail() IFunctDeclArgsListTailContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctDeclArgsListTailContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctDeclArgsListTailContext)
}

func (s *FunctDeclArgsListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctDeclArgsListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctDeclArgsListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRListener); ok {
		listenerT.EnterFunctDeclArgsList(s)
	}
}

func (s *FunctDeclArgsListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRListener); ok {
		listenerT.ExitFunctDeclArgsList(s)
	}
}

func (p *IRParser) FunctDeclArgsList() (localctx IFunctDeclArgsListContext) {
	localctx = NewFunctDeclArgsListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, IRParserRULE_functDeclArgsList)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(150)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case IRParserLBRACK, IRParserTYPE_I8, IRParserTYPE_I32, IRParserTYPE_I64:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(146)
			p.FunctDeclArg()
		}
		{
			p.SetState(147)
			p.FunctDeclArgsListTail()
		}

	case IRParserRPAREN:
		p.EnterOuterAlt(localctx, 2)

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IFunctDeclArgsListTailContext is an interface to support dynamic dispatch.
type IFunctDeclArgsListTailContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctDeclArgsListTailContext differentiates from other interfaces.
	IsFunctDeclArgsListTailContext()
}

type FunctDeclArgsListTailContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctDeclArgsListTailContext() *FunctDeclArgsListTailContext {
	var p = new(FunctDeclArgsListTailContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IRParserRULE_functDeclArgsListTail
	return p
}

func (*FunctDeclArgsListTailContext) IsFunctDeclArgsListTailContext() {}

func NewFunctDeclArgsListTailContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctDeclArgsListTailContext {
	var p = new(FunctDeclArgsListTailContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IRParserRULE_functDeclArgsListTail

	return p
}

func (s *FunctDeclArgsListTailContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctDeclArgsListTailContext) COMMA() antlr.TerminalNode {
	return s.GetToken(IRParserCOMMA, 0)
}

func (s *FunctDeclArgsListTailContext) FunctDeclArgsList() IFunctDeclArgsListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctDeclArgsListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctDeclArgsListContext)
}

func (s *FunctDeclArgsListTailContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctDeclArgsListTailContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctDeclArgsListTailContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRListener); ok {
		listenerT.EnterFunctDeclArgsListTail(s)
	}
}

func (s *FunctDeclArgsListTailContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRListener); ok {
		listenerT.ExitFunctDeclArgsListTail(s)
	}
}

func (p *IRParser) FunctDeclArgsListTail() (localctx IFunctDeclArgsListTailContext) {
	localctx = NewFunctDeclArgsListTailContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, IRParserRULE_functDeclArgsListTail)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(155)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case IRParserCOMMA:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(152)
			p.Match(IRParserCOMMA)
		}
		{
			p.SetState(153)
			p.FunctDeclArgsList()
		}

	case IRParserRPAREN:
		p.EnterOuterAlt(localctx, 2)

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IFunctDeclArgContext is an interface to support dynamic dispatch.
type IFunctDeclArgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctDeclArgContext differentiates from other interfaces.
	IsFunctDeclArgContext()
}

type FunctDeclArgContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctDeclArgContext() *FunctDeclArgContext {
	var p = new(FunctDeclArgContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IRParserRULE_functDeclArg
	return p
}

func (*FunctDeclArgContext) IsFunctDeclArgContext() {}

func NewFunctDeclArgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctDeclArgContext {
	var p = new(FunctDeclArgContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IRParserRULE_functDeclArg

	return p
}

func (s *FunctDeclArgContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctDeclArgContext) Type_() IType_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *FunctDeclArgContext) ParameterAttrib() IParameterAttribContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParameterAttribContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParameterAttribContext)
}

func (s *FunctDeclArgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctDeclArgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctDeclArgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRListener); ok {
		listenerT.EnterFunctDeclArg(s)
	}
}

func (s *FunctDeclArgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRListener); ok {
		listenerT.ExitFunctDeclArg(s)
	}
}

func (p *IRParser) FunctDeclArg() (localctx IFunctDeclArgContext) {
	localctx = NewFunctDeclArgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, IRParserRULE_functDeclArg)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(157)
		p.Type_()
	}
	p.SetState(159)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == IRParserNOCAPTURE {
		{
			p.SetState(158)
			p.ParameterAttrib()
		}

	}

	return localctx
}

// IFunctAttribsListContext is an interface to support dynamic dispatch.
type IFunctAttribsListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctAttribsListContext differentiates from other interfaces.
	IsFunctAttribsListContext()
}

type FunctAttribsListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctAttribsListContext() *FunctAttribsListContext {
	var p = new(FunctAttribsListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IRParserRULE_functAttribsList
	return p
}

func (*FunctAttribsListContext) IsFunctAttribsListContext() {}

func NewFunctAttribsListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctAttribsListContext {
	var p = new(FunctAttribsListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IRParserRULE_functAttribsList

	return p
}

func (s *FunctAttribsListContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctAttribsListContext) FunctAttrib() IFunctAttribContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctAttribContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctAttribContext)
}

func (s *FunctAttribsListContext) FunctAttribsList() IFunctAttribsListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctAttribsListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctAttribsListContext)
}

func (s *FunctAttribsListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctAttribsListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctAttribsListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRListener); ok {
		listenerT.EnterFunctAttribsList(s)
	}
}

func (s *FunctAttribsListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRListener); ok {
		listenerT.ExitFunctAttribsList(s)
	}
}

func (p *IRParser) FunctAttribsList() (localctx IFunctAttribsListContext) {
	localctx = NewFunctAttribsListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, IRParserRULE_functAttribsList)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(165)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case IRParserNOUNWIND:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(161)
			p.FunctAttrib()
		}
		{
			p.SetState(162)
			p.FunctAttribsList()
		}

	case IRParserEOF, IRParserDECLARE, IRParserDEFINE, IRParserLBRACE, IRParserAT:
		p.EnterOuterAlt(localctx, 2)

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IFunctAttribContext is an interface to support dynamic dispatch.
type IFunctAttribContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctAttribContext differentiates from other interfaces.
	IsFunctAttribContext()
}

type FunctAttribContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctAttribContext() *FunctAttribContext {
	var p = new(FunctAttribContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IRParserRULE_functAttrib
	return p
}

func (*FunctAttribContext) IsFunctAttribContext() {}

func NewFunctAttribContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctAttribContext {
	var p = new(FunctAttribContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IRParserRULE_functAttrib

	return p
}

func (s *FunctAttribContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctAttribContext) NOUNWIND() antlr.TerminalNode {
	return s.GetToken(IRParserNOUNWIND, 0)
}

func (s *FunctAttribContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctAttribContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctAttribContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRListener); ok {
		listenerT.EnterFunctAttrib(s)
	}
}

func (s *FunctAttribContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRListener); ok {
		listenerT.ExitFunctAttrib(s)
	}
}

func (p *IRParser) FunctAttrib() (localctx IFunctAttribContext) {
	localctx = NewFunctAttribContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, IRParserRULE_functAttrib)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(167)
		p.Match(IRParserNOUNWIND)
	}

	return localctx
}

// IParameterAttribContext is an interface to support dynamic dispatch.
type IParameterAttribContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParameterAttribContext differentiates from other interfaces.
	IsParameterAttribContext()
}

type ParameterAttribContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParameterAttribContext() *ParameterAttribContext {
	var p = new(ParameterAttribContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IRParserRULE_parameterAttrib
	return p
}

func (*ParameterAttribContext) IsParameterAttribContext() {}

func NewParameterAttribContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParameterAttribContext {
	var p = new(ParameterAttribContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IRParserRULE_parameterAttrib

	return p
}

func (s *ParameterAttribContext) GetParser() antlr.Parser { return s.parser }

func (s *ParameterAttribContext) NOCAPTURE() antlr.TerminalNode {
	return s.GetToken(IRParserNOCAPTURE, 0)
}

func (s *ParameterAttribContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParameterAttribContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParameterAttribContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRListener); ok {
		listenerT.EnterParameterAttrib(s)
	}
}

func (s *ParameterAttribContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRListener); ok {
		listenerT.ExitParameterAttrib(s)
	}
}

func (p *IRParser) ParameterAttrib() (localctx IParameterAttribContext) {
	localctx = NewParameterAttribContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, IRParserRULE_parameterAttrib)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(169)
		p.Match(IRParserNOCAPTURE)
	}

	return localctx
}

// IFunctCallArgsListContext is an interface to support dynamic dispatch.
type IFunctCallArgsListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctCallArgsListContext differentiates from other interfaces.
	IsFunctCallArgsListContext()
}

type FunctCallArgsListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctCallArgsListContext() *FunctCallArgsListContext {
	var p = new(FunctCallArgsListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IRParserRULE_functCallArgsList
	return p
}

func (*FunctCallArgsListContext) IsFunctCallArgsListContext() {}

func NewFunctCallArgsListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctCallArgsListContext {
	var p = new(FunctCallArgsListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IRParserRULE_functCallArgsList

	return p
}

func (s *FunctCallArgsListContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctCallArgsListContext) FunctCallArg() IFunctCallArgContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctCallArgContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctCallArgContext)
}

func (s *FunctCallArgsListContext) FunctCallArgsListTail() IFunctCallArgsListTailContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctCallArgsListTailContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctCallArgsListTailContext)
}

func (s *FunctCallArgsListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctCallArgsListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctCallArgsListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRListener); ok {
		listenerT.EnterFunctCallArgsList(s)
	}
}

func (s *FunctCallArgsListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRListener); ok {
		listenerT.ExitFunctCallArgsList(s)
	}
}

func (p *IRParser) FunctCallArgsList() (localctx IFunctCallArgsListContext) {
	localctx = NewFunctCallArgsListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, IRParserRULE_functCallArgsList)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(175)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case IRParserLBRACK, IRParserTYPE_I8, IRParserTYPE_I32, IRParserTYPE_I64:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(171)
			p.FunctCallArg()
		}
		{
			p.SetState(172)
			p.FunctCallArgsListTail()
		}

	case IRParserRPAREN:
		p.EnterOuterAlt(localctx, 2)

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IFunctCallArgsListTailContext is an interface to support dynamic dispatch.
type IFunctCallArgsListTailContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctCallArgsListTailContext differentiates from other interfaces.
	IsFunctCallArgsListTailContext()
}

type FunctCallArgsListTailContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctCallArgsListTailContext() *FunctCallArgsListTailContext {
	var p = new(FunctCallArgsListTailContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IRParserRULE_functCallArgsListTail
	return p
}

func (*FunctCallArgsListTailContext) IsFunctCallArgsListTailContext() {}

func NewFunctCallArgsListTailContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctCallArgsListTailContext {
	var p = new(FunctCallArgsListTailContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IRParserRULE_functCallArgsListTail

	return p
}

func (s *FunctCallArgsListTailContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctCallArgsListTailContext) COMMA() antlr.TerminalNode {
	return s.GetToken(IRParserCOMMA, 0)
}

func (s *FunctCallArgsListTailContext) FunctCallArgsList() IFunctCallArgsListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctCallArgsListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctCallArgsListContext)
}

func (s *FunctCallArgsListTailContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctCallArgsListTailContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctCallArgsListTailContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRListener); ok {
		listenerT.EnterFunctCallArgsListTail(s)
	}
}

func (s *FunctCallArgsListTailContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRListener); ok {
		listenerT.ExitFunctCallArgsListTail(s)
	}
}

func (p *IRParser) FunctCallArgsListTail() (localctx IFunctCallArgsListTailContext) {
	localctx = NewFunctCallArgsListTailContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, IRParserRULE_functCallArgsListTail)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(180)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case IRParserCOMMA:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(177)
			p.Match(IRParserCOMMA)
		}
		{
			p.SetState(178)
			p.FunctCallArgsList()
		}

	case IRParserRPAREN:
		p.EnterOuterAlt(localctx, 2)

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IFunctCallArgContext is an interface to support dynamic dispatch.
type IFunctCallArgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctCallArgContext differentiates from other interfaces.
	IsFunctCallArgContext()
}

type FunctCallArgContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctCallArgContext() *FunctCallArgContext {
	var p = new(FunctCallArgContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IRParserRULE_functCallArg
	return p
}

func (*FunctCallArgContext) IsFunctCallArgContext() {}

func NewFunctCallArgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctCallArgContext {
	var p = new(FunctCallArgContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IRParserRULE_functCallArg

	return p
}

func (s *FunctCallArgContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctCallArgContext) Type_() IType_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *FunctCallArgContext) LocalIdent() ILocalIdentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILocalIdentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILocalIdentContext)
}

func (s *FunctCallArgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctCallArgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctCallArgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRListener); ok {
		listenerT.EnterFunctCallArg(s)
	}
}

func (s *FunctCallArgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRListener); ok {
		listenerT.ExitFunctCallArg(s)
	}
}

func (p *IRParser) FunctCallArg() (localctx IFunctCallArgContext) {
	localctx = NewFunctCallArgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, IRParserRULE_functCallArg)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(182)
		p.Type_()
	}
	{
		p.SetState(183)
		p.LocalIdent()
	}

	return localctx
}

// IInstructionSeqContext is an interface to support dynamic dispatch.
type IInstructionSeqContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInstructionSeqContext differentiates from other interfaces.
	IsInstructionSeqContext()
}

type InstructionSeqContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInstructionSeqContext() *InstructionSeqContext {
	var p = new(InstructionSeqContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IRParserRULE_instructionSeq
	return p
}

func (*InstructionSeqContext) IsInstructionSeqContext() {}

func NewInstructionSeqContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstructionSeqContext {
	var p = new(InstructionSeqContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IRParserRULE_instructionSeq

	return p
}

func (s *InstructionSeqContext) GetParser() antlr.Parser { return s.parser }

func (s *InstructionSeqContext) Instruction() IInstructionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstructionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstructionContext)
}

func (s *InstructionSeqContext) InstructionSeqTail() IInstructionSeqTailContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstructionSeqTailContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstructionSeqTailContext)
}

func (s *InstructionSeqContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstructionSeqContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InstructionSeqContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRListener); ok {
		listenerT.EnterInstructionSeq(s)
	}
}

func (s *InstructionSeqContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRListener); ok {
		listenerT.ExitInstructionSeq(s)
	}
}

func (p *IRParser) InstructionSeq() (localctx IInstructionSeqContext) {
	localctx = NewInstructionSeqContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, IRParserRULE_instructionSeq)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(185)
		p.Instruction()
	}
	{
		p.SetState(186)
		p.InstructionSeqTail()
	}

	return localctx
}

// IInstructionSeqTailContext is an interface to support dynamic dispatch.
type IInstructionSeqTailContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInstructionSeqTailContext differentiates from other interfaces.
	IsInstructionSeqTailContext()
}

type InstructionSeqTailContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInstructionSeqTailContext() *InstructionSeqTailContext {
	var p = new(InstructionSeqTailContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IRParserRULE_instructionSeqTail
	return p
}

func (*InstructionSeqTailContext) IsInstructionSeqTailContext() {}

func NewInstructionSeqTailContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstructionSeqTailContext {
	var p = new(InstructionSeqTailContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IRParserRULE_instructionSeqTail

	return p
}

func (s *InstructionSeqTailContext) GetParser() antlr.Parser { return s.parser }

func (s *InstructionSeqTailContext) InstructionSeq() IInstructionSeqContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstructionSeqContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstructionSeqContext)
}

func (s *InstructionSeqTailContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstructionSeqTailContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InstructionSeqTailContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRListener); ok {
		listenerT.EnterInstructionSeqTail(s)
	}
}

func (s *InstructionSeqTailContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRListener); ok {
		listenerT.ExitInstructionSeqTail(s)
	}
}

func (p *IRParser) InstructionSeqTail() (localctx IInstructionSeqTailContext) {
	localctx = NewInstructionSeqTailContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, IRParserRULE_instructionSeqTail)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(190)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case IRParserT__0, IRParserT__1, IRParserPERCENT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(188)
			p.InstructionSeq()
		}

	case IRParserRBRACE:
		p.EnterOuterAlt(localctx, 2)

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IInstructionContext is an interface to support dynamic dispatch.
type IInstructionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInstructionContext differentiates from other interfaces.
	IsInstructionContext()
}

type InstructionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInstructionContext() *InstructionContext {
	var p = new(InstructionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IRParserRULE_instruction
	return p
}

func (*InstructionContext) IsInstructionContext() {}

func NewInstructionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstructionContext {
	var p = new(InstructionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IRParserRULE_instruction

	return p
}

func (s *InstructionContext) GetParser() antlr.Parser { return s.parser }

func (s *InstructionContext) InstructionRet() IInstructionRetContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstructionRetContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstructionRetContext)
}

func (s *InstructionContext) InstructionCall() IInstructionCallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstructionCallContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstructionCallContext)
}

func (s *InstructionContext) InstructionGetelementptr() IInstructionGetelementptrContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstructionGetelementptrContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstructionGetelementptrContext)
}

func (s *InstructionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstructionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InstructionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRListener); ok {
		listenerT.EnterInstruction(s)
	}
}

func (s *InstructionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRListener); ok {
		listenerT.ExitInstruction(s)
	}
}

func (p *IRParser) Instruction() (localctx IInstructionContext) {
	localctx = NewInstructionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, IRParserRULE_instruction)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(195)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case IRParserT__0:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(192)
			p.InstructionRet()
		}

	case IRParserT__1:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(193)
			p.InstructionCall()
		}

	case IRParserPERCENT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(194)
			p.InstructionGetelementptr()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IInstructionRetContext is an interface to support dynamic dispatch.
type IInstructionRetContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInstructionRetContext differentiates from other interfaces.
	IsInstructionRetContext()
}

type InstructionRetContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInstructionRetContext() *InstructionRetContext {
	var p = new(InstructionRetContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IRParserRULE_instructionRet
	return p
}

func (*InstructionRetContext) IsInstructionRetContext() {}

func NewInstructionRetContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstructionRetContext {
	var p = new(InstructionRetContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IRParserRULE_instructionRet

	return p
}

func (s *InstructionRetContext) GetParser() antlr.Parser { return s.parser }

func (s *InstructionRetContext) Type_() IType_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *InstructionRetContext) Value() IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *InstructionRetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstructionRetContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InstructionRetContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRListener); ok {
		listenerT.EnterInstructionRet(s)
	}
}

func (s *InstructionRetContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRListener); ok {
		listenerT.ExitInstructionRet(s)
	}
}

func (p *IRParser) InstructionRet() (localctx IInstructionRetContext) {
	localctx = NewInstructionRetContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, IRParserRULE_instructionRet)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(197)
		p.Match(IRParserT__0)
	}
	{
		p.SetState(198)
		p.Type_()
	}
	{
		p.SetState(199)
		p.Value()
	}

	return localctx
}

// IInstructionCallContext is an interface to support dynamic dispatch.
type IInstructionCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInstructionCallContext differentiates from other interfaces.
	IsInstructionCallContext()
}

type InstructionCallContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInstructionCallContext() *InstructionCallContext {
	var p = new(InstructionCallContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IRParserRULE_instructionCall
	return p
}

func (*InstructionCallContext) IsInstructionCallContext() {}

func NewInstructionCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstructionCallContext {
	var p = new(InstructionCallContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IRParserRULE_instructionCall

	return p
}

func (s *InstructionCallContext) GetParser() antlr.Parser { return s.parser }

func (s *InstructionCallContext) Type_() IType_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *InstructionCallContext) GlobalName() IGlobalNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGlobalNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGlobalNameContext)
}

func (s *InstructionCallContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(IRParserLPAREN, 0)
}

func (s *InstructionCallContext) FunctCallArgsList() IFunctCallArgsListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctCallArgsListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctCallArgsListContext)
}

func (s *InstructionCallContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(IRParserRPAREN, 0)
}

func (s *InstructionCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstructionCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InstructionCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRListener); ok {
		listenerT.EnterInstructionCall(s)
	}
}

func (s *InstructionCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRListener); ok {
		listenerT.ExitInstructionCall(s)
	}
}

func (p *IRParser) InstructionCall() (localctx IInstructionCallContext) {
	localctx = NewInstructionCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, IRParserRULE_instructionCall)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(201)
		p.Match(IRParserT__1)
	}
	{
		p.SetState(202)
		p.Type_()
	}
	{
		p.SetState(203)
		p.GlobalName()
	}
	{
		p.SetState(204)
		p.Match(IRParserLPAREN)
	}
	{
		p.SetState(205)
		p.FunctCallArgsList()
	}
	{
		p.SetState(206)
		p.Match(IRParserRPAREN)
	}

	return localctx
}

// IInstructionGetelementptrContext is an interface to support dynamic dispatch.
type IInstructionGetelementptrContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInstructionGetelementptrContext differentiates from other interfaces.
	IsInstructionGetelementptrContext()
}

type InstructionGetelementptrContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInstructionGetelementptrContext() *InstructionGetelementptrContext {
	var p = new(InstructionGetelementptrContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IRParserRULE_instructionGetelementptr
	return p
}

func (*InstructionGetelementptrContext) IsInstructionGetelementptrContext() {}

func NewInstructionGetelementptrContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstructionGetelementptrContext {
	var p = new(InstructionGetelementptrContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IRParserRULE_instructionGetelementptr

	return p
}

func (s *InstructionGetelementptrContext) GetParser() antlr.Parser { return s.parser }

func (s *InstructionGetelementptrContext) AllLocalIdent() []ILocalIdentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ILocalIdentContext)(nil)).Elem())
	var tst = make([]ILocalIdentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ILocalIdentContext)
		}
	}

	return tst
}

func (s *InstructionGetelementptrContext) LocalIdent(i int) ILocalIdentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILocalIdentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ILocalIdentContext)
}

func (s *InstructionGetelementptrContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(IRParserEQUAL, 0)
}

func (s *InstructionGetelementptrContext) Type_() IType_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *InstructionGetelementptrContext) COMMA() antlr.TerminalNode {
	return s.GetToken(IRParserCOMMA, 0)
}

func (s *InstructionGetelementptrContext) PointerType() IPointerTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPointerTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPointerTypeContext)
}

func (s *InstructionGetelementptrContext) InstructionGetelementptrTail() IInstructionGetelementptrTailContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstructionGetelementptrTailContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstructionGetelementptrTailContext)
}

func (s *InstructionGetelementptrContext) GlobalIdent() IGlobalIdentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGlobalIdentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGlobalIdentContext)
}

func (s *InstructionGetelementptrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstructionGetelementptrContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InstructionGetelementptrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRListener); ok {
		listenerT.EnterInstructionGetelementptr(s)
	}
}

func (s *InstructionGetelementptrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRListener); ok {
		listenerT.ExitInstructionGetelementptr(s)
	}
}

func (p *IRParser) InstructionGetelementptr() (localctx IInstructionGetelementptrContext) {
	localctx = NewInstructionGetelementptrContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, IRParserRULE_instructionGetelementptr)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(208)
		p.LocalIdent()
	}
	{
		p.SetState(209)
		p.Match(IRParserEQUAL)
	}
	{
		p.SetState(210)
		p.Match(IRParserT__2)
	}
	{
		p.SetState(211)
		p.Type_()
	}
	{
		p.SetState(212)
		p.Match(IRParserCOMMA)
	}
	{
		p.SetState(213)
		p.PointerType()
	}
	p.SetState(216)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case IRParserAT:
		{
			p.SetState(214)
			p.GlobalIdent()
		}

	case IRParserPERCENT:
		{
			p.SetState(215)
			p.LocalIdent()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(218)
		p.InstructionGetelementptrTail()
	}

	return localctx
}

// IInstructionGetelementptrTailContext is an interface to support dynamic dispatch.
type IInstructionGetelementptrTailContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInstructionGetelementptrTailContext differentiates from other interfaces.
	IsInstructionGetelementptrTailContext()
}

type InstructionGetelementptrTailContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInstructionGetelementptrTailContext() *InstructionGetelementptrTailContext {
	var p = new(InstructionGetelementptrTailContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IRParserRULE_instructionGetelementptrTail
	return p
}

func (*InstructionGetelementptrTailContext) IsInstructionGetelementptrTailContext() {}

func NewInstructionGetelementptrTailContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstructionGetelementptrTailContext {
	var p = new(InstructionGetelementptrTailContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IRParserRULE_instructionGetelementptrTail

	return p
}

func (s *InstructionGetelementptrTailContext) GetParser() antlr.Parser { return s.parser }

func (s *InstructionGetelementptrTailContext) COMMA() antlr.TerminalNode {
	return s.GetToken(IRParserCOMMA, 0)
}

func (s *InstructionGetelementptrTailContext) Type_() IType_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *InstructionGetelementptrTailContext) DECIMALS() antlr.TerminalNode {
	return s.GetToken(IRParserDECIMALS, 0)
}

func (s *InstructionGetelementptrTailContext) InstructionGetelementptrTail() IInstructionGetelementptrTailContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstructionGetelementptrTailContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstructionGetelementptrTailContext)
}

func (s *InstructionGetelementptrTailContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstructionGetelementptrTailContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InstructionGetelementptrTailContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRListener); ok {
		listenerT.EnterInstructionGetelementptrTail(s)
	}
}

func (s *InstructionGetelementptrTailContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRListener); ok {
		listenerT.ExitInstructionGetelementptrTail(s)
	}
}

func (p *IRParser) InstructionGetelementptrTail() (localctx IInstructionGetelementptrTailContext) {
	localctx = NewInstructionGetelementptrTailContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, IRParserRULE_instructionGetelementptrTail)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(226)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case IRParserCOMMA:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(220)
			p.Match(IRParserCOMMA)
		}
		{
			p.SetState(221)
			p.Type_()
		}
		{
			p.SetState(222)
			p.Match(IRParserDECIMALS)
		}
		{
			p.SetState(223)
			p.InstructionGetelementptrTail()
		}

	case IRParserT__0, IRParserT__1, IRParserRBRACE, IRParserPERCENT:
		p.EnterOuterAlt(localctx, 2)

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IRParserRULE_value
	return p
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IRParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) DECIMALS() antlr.TerminalNode {
	return s.GetToken(IRParserDECIMALS, 0)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRListener); ok {
		listenerT.EnterValue(s)
	}
}

func (s *ValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRListener); ok {
		listenerT.ExitValue(s)
	}
}

func (p *IRParser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, IRParserRULE_value)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(228)
		p.Match(IRParserDECIMALS)
	}

	return localctx
}

// IExprConstantContext is an interface to support dynamic dispatch.
type IExprConstantContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExprConstantContext differentiates from other interfaces.
	IsExprConstantContext()
}

type ExprConstantContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprConstantContext() *ExprConstantContext {
	var p = new(ExprConstantContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IRParserRULE_exprConstant
	return p
}

func (*ExprConstantContext) IsExprConstantContext() {}

func NewExprConstantContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprConstantContext {
	var p = new(ExprConstantContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IRParserRULE_exprConstant

	return p
}

func (s *ExprConstantContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprConstantContext) STRING_LIT() antlr.TerminalNode {
	return s.GetToken(IRParserSTRING_LIT, 0)
}

func (s *ExprConstantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprConstantContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprConstantContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRListener); ok {
		listenerT.EnterExprConstant(s)
	}
}

func (s *ExprConstantContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRListener); ok {
		listenerT.ExitExprConstant(s)
	}
}

func (p *IRParser) ExprConstant() (localctx IExprConstantContext) {
	localctx = NewExprConstantContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, IRParserRULE_exprConstant)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(230)
		p.Match(IRParserSTRING_LIT)
	}

	return localctx
}
