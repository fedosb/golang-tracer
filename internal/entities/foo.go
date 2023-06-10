package entities

import (
	"golang-tracer/pkg/tracer"
	"time"
)

type Foo struct {
	tracer tracer.ITracer
}

func NewFoo(tracer tracer.ITracer) *Foo {
	return &Foo{
		tracer: tracer,
	}
}

func (f *Foo) FooDoSomething() {
	f.tracer.StartTrace()
	defer f.tracer.StopTrace()

	time.Sleep(time.Millisecond * 10)
}
