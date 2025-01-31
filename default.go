package exit

import (
	"io"
	"log"
)

var Default = Manager{logger: log.Default()}

func Handle() {
	Default.Handle()
}

func Exit(code int) {
	Default.Exit(code)
}

func Fatal(v ...any) {
	Default.Fatal(v...)
}
func Fatalf(format string, v ...any) {
	Default.Fatalf(format, v...)
}
func Fatalln(v ...any) {
	Default.Fatalln(v...)
}
func Flags() int {
	return Default.Flags()
}
func Output(calldepth int, s string) error {
	return Default.Output(calldepth, s)
}
func Panic(v ...any) {
	Default.Panic(v...)
}
func Panicf(format string, v ...any) {
	Default.Panicf(format, v...)
}
func Panicln(v ...any) {
	Default.Panicln(v...)
}
func Prefix() string {
	return Default.Prefix()
}
func Print(v ...any) {
	Default.Println(v...)
}
func Printf(format string, v ...any) {
	Default.Printf(format, v...)
}
func Println(v ...any) {
	Default.Println(v...)
}
func SetFlags(flag int) {
	Default.SetFlags(flag)
}
func SetOutput(w io.Writer) {
	Default.SetOutput(w)
}
func SetPrefix(prefix string) {
	Default.SetPrefix(prefix)
}
func Writer() io.Writer {
	return Default.Writer()
}
