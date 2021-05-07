// Code generated from IR.g4 by ANTLR 4.9. DO NOT EDIT.

package parser // IR

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseIRListener is a complete listener for a parse tree produced by IRParser.
type BaseIRListener struct{}

var _ IRListener = &BaseIRListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseIRListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseIRListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseIRListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseIRListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterIr is called when production ir is entered.
func (s *BaseIRListener) EnterIr(ctx *IrContext) {}

// ExitIr is called when production ir is exited.
func (s *BaseIRListener) ExitIr(ctx *IrContext) {}

// EnterStatSeq is called when production statSeq is entered.
func (s *BaseIRListener) EnterStatSeq(ctx *StatSeqContext) {}

// ExitStatSeq is called when production statSeq is exited.
func (s *BaseIRListener) ExitStatSeq(ctx *StatSeqContext) {}

// EnterStatSeqTail is called when production statSeqTail is entered.
func (s *BaseIRListener) EnterStatSeqTail(ctx *StatSeqTailContext) {}

// ExitStatSeqTail is called when production statSeqTail is exited.
func (s *BaseIRListener) ExitStatSeqTail(ctx *StatSeqTailContext) {}

// EnterStat is called when production stat is entered.
func (s *BaseIRListener) EnterStat(ctx *StatContext) {}

// ExitStat is called when production stat is exited.
func (s *BaseIRListener) ExitStat(ctx *StatContext) {}

// EnterType_ is called when production type_ is entered.
func (s *BaseIRListener) EnterType_(ctx *Type_Context) {}

// ExitType_ is called when production type_ is exited.
func (s *BaseIRListener) ExitType_(ctx *Type_Context) {}

// EnterSingleValueType is called when production singleValueType is entered.
func (s *BaseIRListener) EnterSingleValueType(ctx *SingleValueTypeContext) {}

// ExitSingleValueType is called when production singleValueType is exited.
func (s *BaseIRListener) ExitSingleValueType(ctx *SingleValueTypeContext) {}

// EnterArrayType is called when production arrayType is entered.
func (s *BaseIRListener) EnterArrayType(ctx *ArrayTypeContext) {}

// ExitArrayType is called when production arrayType is exited.
func (s *BaseIRListener) ExitArrayType(ctx *ArrayTypeContext) {}

// EnterPointerType is called when production pointerType is entered.
func (s *BaseIRListener) EnterPointerType(ctx *PointerTypeContext) {}

// ExitPointerType is called when production pointerType is exited.
func (s *BaseIRListener) ExitPointerType(ctx *PointerTypeContext) {}

// EnterVisibility is called when production visibility is entered.
func (s *BaseIRListener) EnterVisibility(ctx *VisibilityContext) {}

// ExitVisibility is called when production visibility is exited.
func (s *BaseIRListener) ExitVisibility(ctx *VisibilityContext) {}

// EnterLocalIdent is called when production localIdent is entered.
func (s *BaseIRListener) EnterLocalIdent(ctx *LocalIdentContext) {}

// ExitLocalIdent is called when production localIdent is exited.
func (s *BaseIRListener) ExitLocalIdent(ctx *LocalIdentContext) {}

// EnterLocalName is called when production localName is entered.
func (s *BaseIRListener) EnterLocalName(ctx *LocalNameContext) {}

// ExitLocalName is called when production localName is exited.
func (s *BaseIRListener) ExitLocalName(ctx *LocalNameContext) {}

// EnterGlobalIdent is called when production globalIdent is entered.
func (s *BaseIRListener) EnterGlobalIdent(ctx *GlobalIdentContext) {}

// ExitGlobalIdent is called when production globalIdent is exited.
func (s *BaseIRListener) ExitGlobalIdent(ctx *GlobalIdentContext) {}

// EnterGlobalName is called when production globalName is entered.
func (s *BaseIRListener) EnterGlobalName(ctx *GlobalNameContext) {}

// ExitGlobalName is called when production globalName is exited.
func (s *BaseIRListener) ExitGlobalName(ctx *GlobalNameContext) {}

// EnterStatGlobalDef is called when production statGlobalDef is entered.
func (s *BaseIRListener) EnterStatGlobalDef(ctx *StatGlobalDefContext) {}

// ExitStatGlobalDef is called when production statGlobalDef is exited.
func (s *BaseIRListener) ExitStatGlobalDef(ctx *StatGlobalDefContext) {}

// EnterStatFunctDecl is called when production statFunctDecl is entered.
func (s *BaseIRListener) EnterStatFunctDecl(ctx *StatFunctDeclContext) {}

// ExitStatFunctDecl is called when production statFunctDecl is exited.
func (s *BaseIRListener) ExitStatFunctDecl(ctx *StatFunctDeclContext) {}

// EnterStatFunctDef is called when production statFunctDef is entered.
func (s *BaseIRListener) EnterStatFunctDef(ctx *StatFunctDefContext) {}

// ExitStatFunctDef is called when production statFunctDef is exited.
func (s *BaseIRListener) ExitStatFunctDef(ctx *StatFunctDefContext) {}

// EnterFunctSignature is called when production functSignature is entered.
func (s *BaseIRListener) EnterFunctSignature(ctx *FunctSignatureContext) {}

// ExitFunctSignature is called when production functSignature is exited.
func (s *BaseIRListener) ExitFunctSignature(ctx *FunctSignatureContext) {}

// EnterFunctDeclArgsList is called when production functDeclArgsList is entered.
func (s *BaseIRListener) EnterFunctDeclArgsList(ctx *FunctDeclArgsListContext) {}

// ExitFunctDeclArgsList is called when production functDeclArgsList is exited.
func (s *BaseIRListener) ExitFunctDeclArgsList(ctx *FunctDeclArgsListContext) {}

// EnterFunctDeclArgsListTail is called when production functDeclArgsListTail is entered.
func (s *BaseIRListener) EnterFunctDeclArgsListTail(ctx *FunctDeclArgsListTailContext) {}

// ExitFunctDeclArgsListTail is called when production functDeclArgsListTail is exited.
func (s *BaseIRListener) ExitFunctDeclArgsListTail(ctx *FunctDeclArgsListTailContext) {}

// EnterFunctDeclArg is called when production functDeclArg is entered.
func (s *BaseIRListener) EnterFunctDeclArg(ctx *FunctDeclArgContext) {}

// ExitFunctDeclArg is called when production functDeclArg is exited.
func (s *BaseIRListener) ExitFunctDeclArg(ctx *FunctDeclArgContext) {}

// EnterFunctAttribsList is called when production functAttribsList is entered.
func (s *BaseIRListener) EnterFunctAttribsList(ctx *FunctAttribsListContext) {}

// ExitFunctAttribsList is called when production functAttribsList is exited.
func (s *BaseIRListener) ExitFunctAttribsList(ctx *FunctAttribsListContext) {}

// EnterFunctAttrib is called when production functAttrib is entered.
func (s *BaseIRListener) EnterFunctAttrib(ctx *FunctAttribContext) {}

// ExitFunctAttrib is called when production functAttrib is exited.
func (s *BaseIRListener) ExitFunctAttrib(ctx *FunctAttribContext) {}

// EnterParameterAttrib is called when production parameterAttrib is entered.
func (s *BaseIRListener) EnterParameterAttrib(ctx *ParameterAttribContext) {}

// ExitParameterAttrib is called when production parameterAttrib is exited.
func (s *BaseIRListener) ExitParameterAttrib(ctx *ParameterAttribContext) {}

// EnterFunctCallArgsList is called when production functCallArgsList is entered.
func (s *BaseIRListener) EnterFunctCallArgsList(ctx *FunctCallArgsListContext) {}

// ExitFunctCallArgsList is called when production functCallArgsList is exited.
func (s *BaseIRListener) ExitFunctCallArgsList(ctx *FunctCallArgsListContext) {}

// EnterFunctCallArgsListTail is called when production functCallArgsListTail is entered.
func (s *BaseIRListener) EnterFunctCallArgsListTail(ctx *FunctCallArgsListTailContext) {}

// ExitFunctCallArgsListTail is called when production functCallArgsListTail is exited.
func (s *BaseIRListener) ExitFunctCallArgsListTail(ctx *FunctCallArgsListTailContext) {}

// EnterFunctCallArg is called when production functCallArg is entered.
func (s *BaseIRListener) EnterFunctCallArg(ctx *FunctCallArgContext) {}

// ExitFunctCallArg is called when production functCallArg is exited.
func (s *BaseIRListener) ExitFunctCallArg(ctx *FunctCallArgContext) {}

// EnterInstructionSeq is called when production instructionSeq is entered.
func (s *BaseIRListener) EnterInstructionSeq(ctx *InstructionSeqContext) {}

// ExitInstructionSeq is called when production instructionSeq is exited.
func (s *BaseIRListener) ExitInstructionSeq(ctx *InstructionSeqContext) {}

// EnterInstructionSeqTail is called when production instructionSeqTail is entered.
func (s *BaseIRListener) EnterInstructionSeqTail(ctx *InstructionSeqTailContext) {}

// ExitInstructionSeqTail is called when production instructionSeqTail is exited.
func (s *BaseIRListener) ExitInstructionSeqTail(ctx *InstructionSeqTailContext) {}

// EnterInstruction is called when production instruction is entered.
func (s *BaseIRListener) EnterInstruction(ctx *InstructionContext) {}

// ExitInstruction is called when production instruction is exited.
func (s *BaseIRListener) ExitInstruction(ctx *InstructionContext) {}

// EnterInstructionRet is called when production instructionRet is entered.
func (s *BaseIRListener) EnterInstructionRet(ctx *InstructionRetContext) {}

// ExitInstructionRet is called when production instructionRet is exited.
func (s *BaseIRListener) ExitInstructionRet(ctx *InstructionRetContext) {}

// EnterInstructionCall is called when production instructionCall is entered.
func (s *BaseIRListener) EnterInstructionCall(ctx *InstructionCallContext) {}

// ExitInstructionCall is called when production instructionCall is exited.
func (s *BaseIRListener) ExitInstructionCall(ctx *InstructionCallContext) {}

// EnterInstructionGetelementptr is called when production instructionGetelementptr is entered.
func (s *BaseIRListener) EnterInstructionGetelementptr(ctx *InstructionGetelementptrContext) {}

// ExitInstructionGetelementptr is called when production instructionGetelementptr is exited.
func (s *BaseIRListener) ExitInstructionGetelementptr(ctx *InstructionGetelementptrContext) {}

// EnterInstructionGetelementptrTail is called when production instructionGetelementptrTail is entered.
func (s *BaseIRListener) EnterInstructionGetelementptrTail(ctx *InstructionGetelementptrTailContext) {
}

// ExitInstructionGetelementptrTail is called when production instructionGetelementptrTail is exited.
func (s *BaseIRListener) ExitInstructionGetelementptrTail(ctx *InstructionGetelementptrTailContext) {}

// EnterValue is called when production value is entered.
func (s *BaseIRListener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production value is exited.
func (s *BaseIRListener) ExitValue(ctx *ValueContext) {}

// EnterExprConstant is called when production exprConstant is entered.
func (s *BaseIRListener) EnterExprConstant(ctx *ExprConstantContext) {}

// ExitExprConstant is called when production exprConstant is exited.
func (s *BaseIRListener) ExitExprConstant(ctx *ExprConstantContext) {}
