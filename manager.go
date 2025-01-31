package exit

import (
	"io"
	"log"
	"os"
	"runtime"
)

type nullInt struct {
	Value int
	Valid bool
}

type Manager struct {
	code   nullInt
	logger *log.Logger
}

func (m *Manager) Handle() {
	if m.code.Valid {
		os.Exit(m.code.Value)
	}
}

func (m *Manager) Exit(code int) {
	m.code.Value = code
	m.code.Valid = true
	runtime.Goexit()
}

func (m *Manager) Fatal(v ...any) {
	m.logger.Print(v...)
	m.Exit(1)
}
func (m *Manager) Fatalf(format string, v ...any) {
	m.logger.Printf(format, v...)
	m.Exit(1)
}
func (m *Manager) Fatalln(v ...any) {
	m.logger.Println(v...)
	m.Exit(1)
}
func (m *Manager) Flags() int {
	return m.logger.Flags()
}
func (m *Manager) Output(calldepth int, s string) error {
	return m.logger.Output(calldepth, s)
}
func (m *Manager) Panic(v ...any) {
	m.logger.Panic(v...)
}
func (m *Manager) Panicf(format string, v ...any) {
	m.logger.Panicf(format, v...)
}
func (m *Manager) Panicln(v ...any) {
	m.logger.Panicln(v...)
}
func (m *Manager) Prefix() string {
	return m.logger.Prefix()
}
func (m *Manager) Print(v ...any) {
	m.logger.Println(v...)
}
func (m *Manager) Printf(format string, v ...any) {
	m.logger.Printf(format, v...)
}
func (m *Manager) Println(v ...any) {
	m.logger.Println(v...)
}
func (m *Manager) SetFlags(flag int) {
	m.logger.SetFlags(flag)
}
func (m *Manager) SetOutput(w io.Writer) {
	m.logger.SetOutput(w)
}
func (m *Manager) SetPrefix(prefix string) {
	m.logger.SetPrefix(prefix)
}
func (m *Manager) Writer() io.Writer {
	return m.logger.Writer()
}
