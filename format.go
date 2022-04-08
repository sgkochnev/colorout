package colorout

import (
	"bytes"
	"fmt"
	"io"
	"regexp"
	"strings"

	"github.com/mattn/go-colorable"
	"github.com/sgkochnev/colorout/color"
	"github.com/sgkochnev/colorout/style"
)

const escape = "\x1b"

const space byte = ' '

// Reset all stiles
func Reset() string {
	return fmt.Sprintf("%s[%dm", escape, style.Reset)
}

func unset(w io.Writer) {
	if _, err := fmt.Fprint(w, Reset()); err != nil {
		panic("cannot write control sequence in output")
	}
}

func sequence(a ...uint8) string {
	seq := make([]string, 0, len(a))
	for _, v := range a {
		seq = append(seq, fmt.Sprintf("%v", v))
	}
	return strings.Join(seq, ";")
}

func ControlSequence(a ...uint8) string {
	return fmt.Sprintf("%s[%sm", escape, sequence(a...))
}

func countArgs(str string) int {
	reR := regexp.MustCompile(`%[ ]{0,2}%`)
	reF := regexp.MustCompile(`%[ +\-#]?[ +\-#\d]?[.]?[\d]*[bcdefgopqstvwxEFGOTUX]{1}`)
	return len(reF.FindAllStringSubmatch(
		reR.ReplaceAllString(str, ""), -1),
	)
}

func correctCountArgs(expectedCountArgs int, countArgs int) int {
	if expectedCountArgs >= countArgs {
		return countArgs
	}
	return expectedCountArgs
}

type format struct {
	params            bytes.Buffer
	buf               strings.Builder
	expectedCountArgs int
	out               io.Writer
}

func New() *format {
	f := &format{
		params:            bytes.Buffer{},
		buf:               strings.Builder{},
		out:               colorable.NewColorableStdout(),
		expectedCountArgs: 0,
	}
	return f
}

func (f *format) paramsIsEmpty() bool {
	return f.params.Len() == 0
}

func (f *format) bufIsEmpty() bool {
	return f.buf.Len() == 0
}

func (f *format) set(w io.Writer) (n int, err error) {
	if !f.paramsIsEmpty() {
		return fmt.Fprint(w, ControlSequence(f.params.Bytes()...))
	}
	return
}

func (f *format) write(p []byte) (n int, err error) {
	defer f.params.Reset()
	if n, err = f.set(&f.buf); err != nil {
		return
	}
	return f.buf.Write(p)
}

func (f *format) writeString(s string) (n int, err error) {
	return f.write([]byte(s))
}

// Change output
func (f *format) Out(w io.Writer) *format {
	f.out = w
	return f
}

// Styles and buffer reset
func (f *format) Reset() {
	f.buf.Reset()
	f.params.Reset()
	f.expectedCountArgs = 0
}

type fprint = func(w io.Writer, a ...interface{}) (n int, err error)

// wrapper for fmt.Fprint and fmt.Fprintln
func (f *format) print(fp fprint, a ...interface{}) (n int, err error) {
	if n, err = f.set(f.out); err != nil {
		return
	}
	return fp(f.out, a...)
}

type fprintf = func(w io.Writer, format string, a ...interface{}) (n int, err error)

// wrapper for fmt.Fprintf
func (f *format) printf(fpf fprintf, format string, a ...interface{}) (n int, err error) {
	if n, err = f.set(f.out); err != nil {
		return
	}
	return fpf(f.out, format, a...)
}

func (f *format) doPrint(fp fprint, a ...interface{}) (n int, err error) {
	countArgs := len(a)
	count := correctCountArgs(f.expectedCountArgs, countArgs)
	if n, err = f.printf(fmt.Fprintf, f.buf.String(), a[:count]...); err != nil {
		return
	}
	if count != len(a) {
		if _, err = f.out.Write([]byte{space}); err != nil {
			return
		}
	}
	n2, err := f.print(fp, a[count:]...)
	return n + n2 + 1, err
}

func (f *format) Print(a ...interface{}) (n int, err error) {
	if f.bufIsEmpty() {
		return f.print(fmt.Fprint, a...)
	}
	defer unset(f.out)
	return f.doPrint(fmt.Fprint, a...)
}

func (f *format) Println(a ...interface{}) (n int, err error) {
	if f.bufIsEmpty() {
		return f.print(fmt.Fprintln, a...)
	}
	defer unset(f.out)
	return f.doPrint(fmt.Fprintln, a...)
}

func (f *format) Printf(format string, a ...interface{}) (n int, err error) {
	thisFormatExpectedCountArgs := countArgs(format)
	countArgs := len(a)
	k := correctCountArgs(thisFormatExpectedCountArgs, countArgs)
	if n, err = f.doPrint(fmt.Fprint, a[k:]...); err != nil {
		return
	}
	if countArgs-k > 0 {
		defer unset(f.out)
		if _, err = f.out.Write([]byte{space}); err != nil {
			return
		}
	}
	n2, err := f.printf(fmt.Fprintf, format, a[:k]...)
	return n + n2, err

}

func (f *format) Format(format string) (n int, err error) {
	if n, err = f.writeString(format); err != nil {
		return
	}
	f.expectedCountArgs += countArgs(format)
	return
}

// Adds options to create a custom sequence that defines styles.
func (f *format) Style(a ...style.Style) *format {
	f.params.Write(a)
	return f
}

// Adds color-defining options to create a custom style-defining sequence.
func (f *format) CustomFg(c *color.Color) *format {
	f.params.Write(c.Fore())
	return f
}

// Adds background color definition options to create a custom style sequence.
func (f *format) CustomBg(c *color.Color) *format {
	f.params.Write(c.Back())
	return f
}

// String returns the accumulated string.
func (f *format) String() string {
	return f.buf.String() + Reset()
}

func (f *format) ExpectedCountArgs() int {
	return f.expectedCountArgs
}
