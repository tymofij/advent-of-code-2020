package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"strconv"
	"strings"
)

func calcExpr(expr interface{}) int {
	exprType := fmt.Sprintf("%T", expr)
	switch exprType {
	case "*ast.BasicLit":
		lit, _ := expr.(*ast.BasicLit)
		n, _ := strconv.Atoi(lit.Value)
		return n
	case "*ast.BinaryExpr":
		binExpr, _ := expr.(*ast.BinaryExpr)
		switch binExpr.Op {
		case token.QUO: // division, but we know that for us it means addition
			return calcExpr(binExpr.X) + calcExpr(binExpr.Y)
		case token.SUB: // same trick. We just coded multiplication as subtraction
			return calcExpr(binExpr.X) * calcExpr(binExpr.Y)
		}
	case "*ast.ParenExpr":
		parenExpr := expr.(*ast.ParenExpr)
		return calcExpr(parenExpr.X)
	}
	return -1
}

func calcAst(s string) int {
	s = strings.ReplaceAll(s, "+", "/") // hacking the string to give "+" priority over "*"
	s = strings.ReplaceAll(s, "*", "-") // lowering priority of "*" for the same reason
	expr, _ := parser.ParseExpr(s)
	return calcExpr(expr)
}

func main() {
	data, _ := ioutil.ReadFile("../input.txt")
	text := strings.TrimSpace(string(data))
	lines := strings.Split(text, "\n")

	s := 0
	for _, line := range lines {
		s += calcAst(line)
	}
	fmt.Println(s)
}
