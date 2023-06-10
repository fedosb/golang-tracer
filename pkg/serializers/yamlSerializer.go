package serializers

import (
	"github.com/go-yaml/yaml"
	"golang-tracer/pkg/tracer"
	"io"
)

type yamlTraceResultSerializer struct{}

func newYAMLTraceResultSerializer() *yamlTraceResultSerializer {
	return &yamlTraceResultSerializer{}
}

func (y *yamlTraceResultSerializer) Serialize(traceResult *tracer.TraceResult, writer io.Writer) error {
	data, err := yaml.Marshal(traceResult)
	if err != nil {
		return err
	}
	_, err = writer.Write(data)
	return err
}
