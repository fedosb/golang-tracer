package entities

import (
	"golang-tracer/pkg/tracer"
	"time"
)

type Foo struct {
	bar    *Bar
	baz    *Baz
	tracer tracer.ITracer
}

func NewFoo(tracer tracer.ITracer) *Foo {
	return &Foo{
		bar:    NewBar(tracer),
		baz:    NewBaz(tracer),
		tracer: tracer,
	}
}

func (f *Foo) DoSomethingFoo() {
	f.tracer.StartTrace()
	defer f.tracer.StopTrace()

	time.Sleep(5 * time.Millisecond)
	f.bar.DoSomethingBar()
	time.Sleep(15 * time.Millisecond)
	f.baz.DoSomethingBaz()
	time.Sleep(15 * time.Millisecond)
}

func (f *Foo) DoSomethingElseFoo() {
	f.tracer.StartTrace()
	defer f.tracer.StopTrace()

	time.Sleep(5 * time.Millisecond)
	f.bar.DoSomethingBar()
	time.Sleep(15 * time.Millisecond)
}
