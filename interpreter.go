package main

// TODO: Implement error strings

type InterpreterError interface {
	error
	isInterpreterError()
}

type InterpreterTypeError struct {
	Expected Value
	Got      Value
}

func (InterpreterTypeError) isInterpreterError() {}
func (InterpreterTypeError) Error() string {
	return "Unimplemented"
}

type InterpreterLabelError struct {
	// TODO: Figure out what to name this variable
	Label string
}

func (InterpreterLabelError) isInterpreterError() {}
func (InterpreterLabelError) Error() string {
	return "Unimplemented"
}

type InterpreterValueError struct {
	Value Value
}

func (InterpreterValueError) isInterpreterError() {}
func (InterpreterValueError) Error() string {
	return "Unimplemented"
}

type InterpreterNoSlotError struct{}

func (InterpreterNoSlotError) isInterpreterError() {}
func (InterpreterNoSlotError) Error() string {
	return "Unimplemented"
}

type InterpreterGenericError struct {
	ErrorString string
}

func (InterpreterGenericError) isInterpreterError() {}
func (e InterpreterGenericError) Error() string {
	// return "Unimplemented"
	return e.ErrorString
}

type InterpreterErrorInfo struct {
	Error InterpreterError
	Fun   Fun
	Note  string
}

type InterpreterSignal interface {
	isSignal()
}

type ErrorSignal struct {
	ErrorInfo InterpreterErrorInfo
}

func (ErrorSignal) isSignal() {}

type InterpreterErrorSignal struct {
	InterpreterError InterpreterError
}

func (InterpreterErrorSignal) isSignal() {}

type JumpSignal struct {
	n uint
}

func (JumpSignal) isSignal() {}

type ReturnSignal struct {
	Value Value
}

func (ReturnSignal) isSignal() {}
