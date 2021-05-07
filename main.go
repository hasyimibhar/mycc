package main

import (
	"fmt"
	"os"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/hasyimibhar/mycc/ir/parser"
)

type TreeShapeListener struct {
	*parser.BaseIRListener
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

func (this *TreeShapeListener) EnterStatFunctDecl(ctx *parser.StatFunctDeclContext) {
	fmt.Println("found function declaration")
}

func (this *TreeShapeListener) EnterFunctSignature(ctx *parser.FunctSignatureContext) {
	fmt.Println("name:", ctx.GlobalName().GetText())
	fmt.Println("type:", ctx.Type_().GetText())
	fmt.Println("arguments:")
}

func (this *TreeShapeListener) EnterFunctDeclArg(ctx *parser.FunctDeclArgContext) {
	fmt.Println("-", ctx.Type_().GetText())
}

func main() {
	input, _ := antlr.NewFileStream(os.Args[1])
	lexer := parser.NewIRLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewIRParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	tree := p.Ir()
	antlr.ParseTreeWalkerDefault.Walk(NewTreeShapeListener(), tree)
}
