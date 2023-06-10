package main

import (
	"encoding/json"
	"fmt"
	"golang-tracer/internal/entities"
	"golang-tracer/pkg/tracer"
)

func main() {

	myTracer := tracer.NewTracer()
	foo := entities.NewFoo(myTracer)

	foo.DoSomethingFoo()

	res := myTracer.GetTraceResult()

	encoded, _ := json.Marshal(res)
	fmt.Println(string(encoded))
}
