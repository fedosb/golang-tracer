package entities

import (
	"golang-tracer/pkg/tracer"
	"time"
)

type Bar struct {
	tracer tracer.ITracer
}

func NewBar(tracer tracer.ITracer) *Bar {
	return &Bar{
		tracer: tracer,
	}
}

func (b *Bar) DoSomethingBar() {
	b.tracer.StartTrace()
	defer b.tracer.StopTrace()

	time.Sleep(10 * time.Millisecond)
}
