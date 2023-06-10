package entities

import (
	"golang-tracer/pkg/tracer"
	"time"
)

type Baz struct {
	tracer tracer.ITracer
	qux    *Qux
}

func NewBaz(tracer tracer.ITracer) *Baz {
	return &Baz{
		tracer: tracer,
		qux:    NewQux(tracer),
	}
}

func (b *Baz) DoSomethingBaz() {
	b.tracer.StartTrace()
	defer b.tracer.StopTrace()

	time.Sleep(10 * time.Millisecond)
	b.qux.DoSomethingQux()
	time.Sleep(10 * time.Millisecond)
}
