package day6

import "fmt"

type CustomError struct {
	message    string
	errcontext string
	lineNum    int
	filename   string
}

func (c CustomError) Error() string {
	return fmt.Sprintln("message: ", c.message, "errContext: ", c.errcontext, "Line Number:", c.lineNum, "FilenName", c.filename)
}
