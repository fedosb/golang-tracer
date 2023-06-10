package entities

import (
	"golang-tracer/pkg/tracer"
	"time"
)

type Quz struct {
	tracer tracer.ITracer
}

func NewQuz(tracer tracer.ITracer) *Quz {
	return &Quz{
		tracer: tracer,
	}
}

func (q *Quz) DoSomethingQuz() {
	q.tracer.StartTrace()
	defer q.tracer.StopTrace()

	time.Sleep(10 * time.Millisecond)
}
