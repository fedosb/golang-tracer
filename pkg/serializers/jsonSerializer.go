package serializers

import (
	"encoding/json"
	"golang-tracer/pkg/tracer"
	"io"
)

type jsonTraceResultSerializer struct{}

func newJSONTraceResultSerializer() *jsonTraceResultSerializer {
	return &jsonTraceResultSerializer{}
}

func (j *jsonTraceResultSerializer) Serialize(traceResult *tracer.TraceResult, writer io.Writer) error {
	encoder := json.NewEncoder(writer)
	encoder.SetIndent("", "  ")
	return encoder.Encode(traceResult)
}
