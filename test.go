package test

func Test(function string) (result string) {

	if function == "fmt.Println()" {
		result = `func Println(a ...interface{}) (n int, err error)
		Println formats using the default formats for its operands and writes to
		standard output. Spaces are always added between operands and a newline is
		appended. It returns the number of bytes written and any write error
		encountered.
	`

	} else if function == "fmt.Printf" {
		result = `func Printf(format string, a ...interface{}) (n int, err error)
		Printf formats according to a format specifier and writes to standard
		output. It returns the number of bytes written and any write error
		encountered.`

	} else if function == "fmt.Scanln" {
		result = `func Scanln(a ...interface{}) (n int, err error)
		Scanln is similar to Scan, but stops scanning at a newline and after the
		final item there must be a newline or EOF.`

	} else if function == "fmt.Scanf" {
		result = `func Scanf(format string, a ...interface{}) (n int, err error)
		Scanf scans text read from standard input, storing successive
		space-separated values into successive arguments as determined by the
		format. It returns the number of items successfully scanned. If that is less
		than the number of arguments, err will report why. Newlines in the input
		must match newlines in the format. The one exception: the verb %c always
		scans the next rune in the input, even if it is a space (or tab etc.) or
		newline.
	`

	} else if function == "variadic function" {
		result = `A variadic function is a function that accepts a variable number of arguments. 
		In Golang, it is possible to pass a varying number of arguments of the same type 
		as referenced in the function signature.`

	} else if function == "niladic function" {
		result = "niladic function is a function that not returns any value"

	} else {
		result = "There is no any docs for this function, but Inshaallah, we will add it"
	}

	return result
}