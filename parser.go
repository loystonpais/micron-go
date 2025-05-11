package main

import (
	"fmt"
	"strconv"
)

type Value interface {
	Display
	DisplayCode
	isValue()
}

type StrValue struct {
	Val string
}

func (StrValue) isValue() {}
func (v StrValue) Repr() string {
	return fmt.Sprintf("Str(%s)", v.Val)
}
func (v StrValue) Code() string {
	// TODO: Implement a valid string code
	return fmt.Sprintf("%#v", v.Val)
}

type IntValue struct {
	Val int
}

func (IntValue) isValue() {}
func (v IntValue) Repr() string {
	return fmt.Sprintf("Int(%d)", v.Val)
}
func (v IntValue) Code() string {
	// TODO: Make sure that this returns a valid
	// micron integer datatype
	return strconv.Itoa(v.Val)
}

type NoneValue struct{}

func (NoneValue) isValue() {}
func (v NoneValue) Repr() string {
	return "None"
}

type Expr interface {
	Display
	DisplayCode
	isExpr()
}

type ValueExpr struct {
	Expr Value
}

func (ValueExpr) isExpr() {}
func (e ValueExpr) Repr() string {
	return fmt.Sprintf("Value(%s)", e.Expr.Repr())
}
func (e ValueExpr) Code() string {
	return e.Expr.Code()
}

type FunCallExpr struct {
	Expr Fun
}

func (e FunCallExpr) Repr() string {
	return fmt.Sprintf("FunCall(%s)", (e.Expr).Repr())
}
func (e FunCallExpr) Code() string {
	return e.Expr.Code()
}

func (FunCallExpr) isExpr() {}

type Fun interface {
	Display
	DisplayCode
	isFun()
}

type SetFun struct {
	Expr1 Expr
	Expr2 Expr
}

func (SetFun) isFun() {}
func (f SetFun) Repr() string {
	return fmt.Sprintf("Set(%s, %s)", f.Expr1.Repr(), f.Expr2.Repr())
}
func (f SetFun) Code() string {
	return fmt.Sprintf("s: %s %s", f.Expr1.Code(), f.Expr2.Code())
}

type GetFun struct {
	Expr Expr
}

func (f GetFun) Repr() string {
	return fmt.Sprintf("Get(%s)", f.Expr.Repr())
}
func (f GetFun) Code() string {
	return fmt.Sprintf("g: %s", f.Expr.Code())
}

type WriteFun struct {
	Expr Expr
}

func (f WriteFun) Repr() string {
	return fmt.Sprintf("Write(%s)", f.Expr.Repr())
}
func (f WriteFun) Code() string {
	return fmt.Sprintf("w: %s", f.Expr.Code())
}

type PrintFun struct {
	Expr Expr
}

func (f PrintFun) Repr() string {
	return fmt.Sprintf("Print(%s)", f.Expr.Repr())
}
func (f PrintFun) Code() string {
	return fmt.Sprintf("p: %s", f.Expr.Code())
}

type AddFun struct {
	Expr1 Expr
	Expr2 Expr
}

func (AddFun) isFun() {}
func (f AddFun) Repr() string {
	return fmt.Sprintf("Add(%s, %s)", f.Expr1.Repr(), f.Expr2.Repr())
}
func (f AddFun) Code() string {
	return fmt.Sprintf("a: %s %s", f.Expr1.Code(), f.Expr2.Code())
}

type JumpFun struct {
	Expr Expr
}

func (f JumpFun) Repr() string {
	return fmt.Sprintf("Jump(%s)", f.Expr.Repr())
}
func (f JumpFun) Code() string {
	return fmt.Sprintf("j: %s", f.Expr.Code())
}

type ConvertFun struct {
	Expr Expr
}

func (f ConvertFun) Repr() string {
	return fmt.Sprintf("Convert(%s)", f.Expr.Repr())
}
func (f ConvertFun) Code() string {
	return fmt.Sprintf("c: %s", f.Expr.Code())
}

type IfFun struct {
	Expr1 Expr
	Expr2 Expr
}

func (f IfFun) Repr() string {
	return fmt.Sprintf("If(%s, %s)", f.Expr1.Repr(), f.Expr2.Repr())
}

type InputFun struct{}

func (InputFun) Repr() string {
	return "Input"
}
func (InputFun) Code() string {
	return "i"
}

type KeyCharFun struct{}

func (KeyCharFun) Repr() string {
	return "KeyChar"
}
func (KeyCharFun) Code() string {
	return "k"
}

type TextFun struct {
	Expr Expr
}

func (f TextFun) Repr() string {
	return fmt.Sprintf("Text(%s)", f.Expr.Repr())
}
func (f TextFun) Code() string {
	return fmt.Sprintf("t: %s", f.Expr.Repr())
}

type NumberFun struct {
	Expr Expr
}

func (f NumberFun) Repr() string {
	return fmt.Sprintf("Number(%s)", f.Expr.Repr())
}

func (f NumberFun) Code() string {
	return fmt.Sprintf("n: %s", f.Expr.Repr())
}

type CatchErrorFun struct {
	Expr1 Expr
	Expr2 Expr
}

func (f CatchErrorFun) Repr() string {
	return fmt.Sprintf("CatchError(%s, %s)", f.Expr1.Repr(), f.Expr2.Repr())
}
func (f CatchErrorFun) Code() string {
	return fmt.Sprintf("#: %s %s", f.Expr1.Repr(), f.Expr2.Repr())
}

type ThrowErrorFun struct {
	Expr Expr
}

func (f ThrowErrorFun) Repr() string {
	return fmt.Sprintf("ThrowError(%s)", f.Expr.Repr())
}
func (f ThrowErrorFun) Code() string {
	return fmt.Sprintf("!: %s", f.Expr.Repr())
}

type ReturnFun struct {
	Expr Expr
}

func (f ReturnFun) Repr() string {
	return fmt.Sprintf("Return(%s)", f.Expr.Repr())
}
func (f ReturnFun) Code() string {
	return fmt.Sprintf("r: %s", f.Expr.Repr())
}

type FunJumpFun struct {
	Expr Expr
}

func (f FunJumpFun) Repr() string {
	return fmt.Sprintf("FunJump(%s)", f.Expr.Repr())
}
func (f FunJumpFun) Code() string {
	return fmt.Sprintf("f: %s", f.Expr.Repr())
}

type EmptySlotFun struct{}

func (EmptySlotFun) Repr() string {
	return "EmptySlot"
}
func (EmptySlotFun) Code() string {
	return "~"
}

type ExitFun struct{}

func (ExitFun) Repr() string {
	return "Exit"
}

func (ExitFun) Code() string {
	return "$"
}

type Instr interface {
	isInstr()
}

type SetLabelInstr struct {
	Label string
}

func (SetLabelInstr) isInstr() {}

type LabelPlaceHolderInstr struct {
	Label string
}

func (LabelPlaceHolderInstr) isInstr() {}

type FunCallInstr struct {
	Fun Fun
}

func (FunCallInstr) isInstr() {}

type InstrInfo struct {
	Instr Instr
	Start uint
	End   uint
}

type ParseError interface {
	error
	isParseError()
	ErrorCode() uint
}

type LabelAlreadySetParseError struct {
	Label string
	Line  []TokenInfo
}

func (LabelAlreadySetParseError) isParseError() {}
func (e LabelAlreadySetParseError) Error() string {
	return fmt.Sprintf("Label '%s' is already set", e.Label)
}
func (e LabelAlreadySetParseError) ErrorCode() uint {
	return 301
}

type UnexpectedTokenParseError struct {
	TokenInfo TokenInfo
}

func (UnexpectedTokenParseError) isParseError() {}
func (e UnexpectedTokenParseError) Error() string {
	return fmt.Sprintf("Unexpected token '%s'", e.TokenInfo.Token.Repr())
}
func (e UnexpectedTokenParseError) ErrorCode() uint {
	return 302
}

type InvalidSyntaxParseError struct{}

func (InvalidSyntaxParseError) isParseError() {}
func (e InvalidSyntaxParseError) Error() string {
	return fmt.Sprintf("Invalid Syntax")
}
func (e InvalidSyntaxParseError) ErrorCode() uint {
	return 303
}

type NotEnoughArgumentParseError struct {
	TokenInfo TokenInfo
	Got       uint
	Expected  uint
}

type UnknownFunctionNameParseError struct {
	TokenInfo TokenInfo
}
