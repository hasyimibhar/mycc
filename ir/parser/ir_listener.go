// Code generated from IR.g4 by ANTLR 4.9. DO NOT EDIT.

package parser // IR

import "github.com/antlr/antlr4/runtime/Go/antlr"

// IRListener is a complete listener for a parse tree produced by IRParser.
type IRListener interface {
	antlr.ParseTreeListener

	// EnterIr is called when entering the ir production.
	EnterIr(c *IrContext)

	// EnterStatSeq is called when entering the statSeq production.
	EnterStatSeq(c *StatSeqContext)

	// EnterStatSeqTail is called when entering the statSeqTail production.
	EnterStatSeqTail(c *StatSeqTailContext)

	// EnterStat is called when entering the stat production.
	EnterStat(c *StatContext)

	// EnterType_ is called when entering the type_ production.
	EnterType_(c *Type_Context)

	// EnterSingleValueType is called when entering the singleValueType production.
	EnterSingleValueType(c *SingleValueTypeContext)

	// EnterArrayType is called when entering the arrayType production.
	EnterArrayType(c *ArrayTypeContext)

	// EnterPointerType is called when entering the pointerType production.
	EnterPointerType(c *PointerTypeContext)

	// EnterVisibility is called when entering the visibility production.
	EnterVisibility(c *VisibilityContext)

	// EnterLocalIdent is called when entering the localIdent production.
	EnterLocalIdent(c *LocalIdentContext)

	// EnterLocalName is called when entering the localName production.
	EnterLocalName(c *LocalNameContext)

	// EnterGlobalIdent is called when entering the globalIdent production.
	EnterGlobalIdent(c *GlobalIdentContext)

	// EnterGlobalName is called when entering the globalName production.
	EnterGlobalName(c *GlobalNameContext)

	// EnterStatGlobalDef is called when entering the statGlobalDef production.
	EnterStatGlobalDef(c *StatGlobalDefContext)

	// EnterStatFunctDecl is called when entering the statFunctDecl production.
	EnterStatFunctDecl(c *StatFunctDeclContext)

	// EnterStatFunctDef is called when entering the statFunctDef production.
	EnterStatFunctDef(c *StatFunctDefContext)

	// EnterFunctSignature is called when entering the functSignature production.
	EnterFunctSignature(c *FunctSignatureContext)

	// EnterFunctDeclArgsList is called when entering the functDeclArgsList production.
	EnterFunctDeclArgsList(c *FunctDeclArgsListContext)

	// EnterFunctDeclArgsListTail is called when entering the functDeclArgsListTail production.
	EnterFunctDeclArgsListTail(c *FunctDeclArgsListTailContext)

	// EnterFunctDeclArg is called when entering the functDeclArg production.
	EnterFunctDeclArg(c *FunctDeclArgContext)

	// EnterFunctAttribsList is called when entering the functAttribsList production.
	EnterFunctAttribsList(c *FunctAttribsListContext)

	// EnterFunctAttrib is called when entering the functAttrib production.
	EnterFunctAttrib(c *FunctAttribContext)

	// EnterParameterAttrib is called when entering the parameterAttrib production.
	EnterParameterAttrib(c *ParameterAttribContext)

	// EnterFunctCallArgsList is called when entering the functCallArgsList production.
	EnterFunctCallArgsList(c *FunctCallArgsListContext)

	// EnterFunctCallArgsListTail is called when entering the functCallArgsListTail production.
	EnterFunctCallArgsListTail(c *FunctCallArgsListTailContext)

	// EnterFunctCallArg is called when entering the functCallArg production.
	EnterFunctCallArg(c *FunctCallArgContext)

	// EnterInstructionSeq is called when entering the instructionSeq production.
	EnterInstructionSeq(c *InstructionSeqContext)

	// EnterInstructionSeqTail is called when entering the instructionSeqTail production.
	EnterInstructionSeqTail(c *InstructionSeqTailContext)

	// EnterInstruction is called when entering the instruction production.
	EnterInstruction(c *InstructionContext)

	// EnterInstructionRet is called when entering the instructionRet production.
	EnterInstructionRet(c *InstructionRetContext)

	// EnterInstructionCall is called when entering the instructionCall production.
	EnterInstructionCall(c *InstructionCallContext)

	// EnterInstructionGetelementptr is called when entering the instructionGetelementptr production.
	EnterInstructionGetelementptr(c *InstructionGetelementptrContext)

	// EnterInstructionGetelementptrTail is called when entering the instructionGetelementptrTail production.
	EnterInstructionGetelementptrTail(c *InstructionGetelementptrTailContext)

	// EnterValue is called when entering the value production.
	EnterValue(c *ValueContext)

	// EnterExprConstant is called when entering the exprConstant production.
	EnterExprConstant(c *ExprConstantContext)

	// ExitIr is called when exiting the ir production.
	ExitIr(c *IrContext)

	// ExitStatSeq is called when exiting the statSeq production.
	ExitStatSeq(c *StatSeqContext)

	// ExitStatSeqTail is called when exiting the statSeqTail production.
	ExitStatSeqTail(c *StatSeqTailContext)

	// ExitStat is called when exiting the stat production.
	ExitStat(c *StatContext)

	// ExitType_ is called when exiting the type_ production.
	ExitType_(c *Type_Context)

	// ExitSingleValueType is called when exiting the singleValueType production.
	ExitSingleValueType(c *SingleValueTypeContext)

	// ExitArrayType is called when exiting the arrayType production.
	ExitArrayType(c *ArrayTypeContext)

	// ExitPointerType is called when exiting the pointerType production.
	ExitPointerType(c *PointerTypeContext)

	// ExitVisibility is called when exiting the visibility production.
	ExitVisibility(c *VisibilityContext)

	// ExitLocalIdent is called when exiting the localIdent production.
	ExitLocalIdent(c *LocalIdentContext)

	// ExitLocalName is called when exiting the localName production.
	ExitLocalName(c *LocalNameContext)

	// ExitGlobalIdent is called when exiting the globalIdent production.
	ExitGlobalIdent(c *GlobalIdentContext)

	// ExitGlobalName is called when exiting the globalName production.
	ExitGlobalName(c *GlobalNameContext)

	// ExitStatGlobalDef is called when exiting the statGlobalDef production.
	ExitStatGlobalDef(c *StatGlobalDefContext)

	// ExitStatFunctDecl is called when exiting the statFunctDecl production.
	ExitStatFunctDecl(c *StatFunctDeclContext)

	// ExitStatFunctDef is called when exiting the statFunctDef production.
	ExitStatFunctDef(c *StatFunctDefContext)

	// ExitFunctSignature is called when exiting the functSignature production.
	ExitFunctSignature(c *FunctSignatureContext)

	// ExitFunctDeclArgsList is called when exiting the functDeclArgsList production.
	ExitFunctDeclArgsList(c *FunctDeclArgsListContext)

	// ExitFunctDeclArgsListTail is called when exiting the functDeclArgsListTail production.
	ExitFunctDeclArgsListTail(c *FunctDeclArgsListTailContext)

	// ExitFunctDeclArg is called when exiting the functDeclArg production.
	ExitFunctDeclArg(c *FunctDeclArgContext)

	// ExitFunctAttribsList is called when exiting the functAttribsList production.
	ExitFunctAttribsList(c *FunctAttribsListContext)

	// ExitFunctAttrib is called when exiting the functAttrib production.
	ExitFunctAttrib(c *FunctAttribContext)

	// ExitParameterAttrib is called when exiting the parameterAttrib production.
	ExitParameterAttrib(c *ParameterAttribContext)

	// ExitFunctCallArgsList is called when exiting the functCallArgsList production.
	ExitFunctCallArgsList(c *FunctCallArgsListContext)

	// ExitFunctCallArgsListTail is called when exiting the functCallArgsListTail production.
	ExitFunctCallArgsListTail(c *FunctCallArgsListTailContext)

	// ExitFunctCallArg is called when exiting the functCallArg production.
	ExitFunctCallArg(c *FunctCallArgContext)

	// ExitInstructionSeq is called when exiting the instructionSeq production.
	ExitInstructionSeq(c *InstructionSeqContext)

	// ExitInstructionSeqTail is called when exiting the instructionSeqTail production.
	ExitInstructionSeqTail(c *InstructionSeqTailContext)

	// ExitInstruction is called when exiting the instruction production.
	ExitInstruction(c *InstructionContext)

	// ExitInstructionRet is called when exiting the instructionRet production.
	ExitInstructionRet(c *InstructionRetContext)

	// ExitInstructionCall is called when exiting the instructionCall production.
	ExitInstructionCall(c *InstructionCallContext)

	// ExitInstructionGetelementptr is called when exiting the instructionGetelementptr production.
	ExitInstructionGetelementptr(c *InstructionGetelementptrContext)

	// ExitInstructionGetelementptrTail is called when exiting the instructionGetelementptrTail production.
	ExitInstructionGetelementptrTail(c *InstructionGetelementptrTailContext)

	// ExitValue is called when exiting the value production.
	ExitValue(c *ValueContext)

	// ExitExprConstant is called when exiting the exprConstant production.
	ExitExprConstant(c *ExprConstantContext)
}
