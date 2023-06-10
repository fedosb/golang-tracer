package main

import (
	"golang-tracer/internal/entities"
	"golang-tracer/pkg/serializers"
	"golang-tracer/pkg/tracer"
	"os"
)

func main() {

	myTracer := tracer.NewTracer()
	foo := entities.NewFoo(myTracer)

	foo.DoSomethingFoo()
	foo.DoSomethingElseFoo()

	traceResult := myTracer.GetTraceResult()

	serializer := serializers.NewJSONTraceResultSerializer()
	if err := serializer.Serialize(traceResult, os.Stdout); err != nil {
		panic(err)
	}
}
