package serializers

import (
	"golang-tracer/pkg/tracer"
	"io"
)

type ITraceResultSerializer interface {
	Serialize(traceResult *tracer.TraceResult, writer io.Writer) error
}
