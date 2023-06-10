package tracer

import (
	"runtime"
	"strings"
	"time"
)

type TraceResult struct {
	Name    string
	Class   string
	Time    time.Duration
	Methods []*TraceResult
}

type Tracer struct {
	stack      []*TraceResult
	rootResult *TraceResult
	startTime  time.Time
}

func NewTracer() ITracer {
	rootResult := &TraceResult{
		Name: "root",
	}
	return &Tracer{
		stack:      []*TraceResult{rootResult},
		rootResult: rootResult,
	}
}

func (t *Tracer) StartTrace() {
	t.startTime = time.Now()
	methodResult := &TraceResult{
		Name:  getCallingMethodName(),
		Class: getCallingClassName(),
	}
	t.stack[len(t.stack)-1].Methods = append(t.stack[len(t.stack)-1].Methods, methodResult)
	t.stack = append(t.stack, methodResult)
}

func (t *Tracer) StopTrace() {
	t.stack[len(t.stack)-1].Time += time.Since(t.startTime)
	t.stack = t.stack[:len(t.stack)-1]

	if t.stack[len(t.stack)-1].Time == 0 {
		for _, m := range t.stack[len(t.stack)-1].Methods {
			t.stack[len(t.stack)-1].Time += m.Time
		}
	}

}

func (t *Tracer) GetTraceResult() *TraceResult {
	return t.rootResult
}

func getCallingMethodName() string {
	pc, _, _, _ := runtime.Caller(2)
	method := runtime.FuncForPC(pc)
	methodParts := strings.Split(method.Name(), ".")
	return methodParts[len(methodParts)-1]
}

func getCallingClassName() string {
	pc, _, _, _ := runtime.Caller(2)
	method := runtime.FuncForPC(pc)
	methodParts := strings.Split(method.Name(), ".")
	return methodParts[len(methodParts)-2]
}
