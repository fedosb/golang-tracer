package serializers

import (
	"encoding/json"
	"golang-tracer/pkg/tracer"
	"io"
)

type JSONTraceResultSerializer struct{}

func NewJSONTraceResultSerializer() *JSONTraceResultSerializer {
	return &JSONTraceResultSerializer{}
}

func (j *JSONTraceResultSerializer) Serialize(traceResult *tracer.TraceResult, writer io.Writer) error {
	encoder := json.NewEncoder(writer)
	encoder.SetIndent("", "  ")
	return encoder.Encode(traceResult)
}
