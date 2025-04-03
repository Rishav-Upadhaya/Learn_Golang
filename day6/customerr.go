package day6

type CustomError struct {
	message    string
	errcontext string
	lineNum    int
	filename   string
}
