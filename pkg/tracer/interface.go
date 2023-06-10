package tracer

type ITracer interface {
	StartTrace()
	StopTrace()
	GetTraceResult() *TraceResult
}
