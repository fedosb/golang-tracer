package entities

import (
	"golang-tracer/pkg/tracer"
	"time"
)

type Qux struct {
	tracer tracer.ITracer
	quz    *Quz
}

func NewQux(tracer tracer.ITracer) *Qux {
	return &Qux{
		tracer: tracer,
		quz:    NewQuz(tracer),
	}
}

func (q *Qux) DoSomethingQux() {
	q.tracer.StartTrace()
	defer q.tracer.StopTrace()

	time.Sleep(10 * time.Millisecond)
	q.quz.DoSomethingQuz()
	time.Sleep(10 * time.Millisecond)
}
