package console

import (
	"bytes"

	"github.com/neelance/gopherjs/js"
)

var c = js.Global("console")

func Assert(b bool, msg interface{})     { c.Call("assert", b, msg) }
func Clear()                             { c.Call("clear") }
func Count(label string)                 { c.Call("count", label) }
func Dir(obj interface{})                { c.Call("dir", obj) }
func DirXML(obj interface{})             { c.Call("dirxml", obj) }
func Error(objs ...interface{})          { c.Call("error", objs...) }
func Group(objs ...interface{})          { c.Call("group", objs...) }
func GroupCollapsed(objs ...interface{}) { c.Call("groupCollapsed", objs...) }
func GroupEnd()                          { c.Call("groupEnd") }
func Log(objs ...interface{})            { c.Call("log", objs...) }
func Profile(label interface{})          { c.Call("profile", label) }
func ProfileEnd()                        { c.Call("profileEnd") }
func Time(label interface{})             { c.Call("time", label) }
func TimeEnd(label interface{})          { c.Call("timeEnd", label) }
func Timestamp(label interface{})        { c.Call("timeStamp", label) }
func Trace()                             { c.Call("trace") }
func Warn(objs ...interface{})           { c.Call("warn", objs...) }

type Writer struct {
	buf *bytes.Buffer
}

func (w *Writer) Write(buf []byte) (n int, err error) {
	if len(buf) == 0 {
		return 0, nil
	}

	for i := len(buf); i >= 0; i-- {
		if buf[i] == '\n' {
			w.buf.Write(buf[:i])
			Log(w.buf.String())
			w.buf.Reset()
			w.buf.Write(buf[i+1:])
			break
		}
	}

	return len(buf), nil
}

func (w *Writer) Flush() {
	w.Write([]byte{'\n'})
}

func New() *Writer {
	return &Writer{buf: new(bytes.Buffer)}
}
