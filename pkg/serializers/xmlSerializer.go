package serializers

import (
	"encoding/xml"
	"golang-tracer/pkg/tracer"
	"io"
)

type xmlTraceResultSerializer struct{}

func newXMLTraceResultSerializer() *xmlTraceResultSerializer {
	return &xmlTraceResultSerializer{}
}

func (x *xmlTraceResultSerializer) Serialize(traceResult *tracer.TraceResult, writer io.Writer) error {
	data, err := xml.MarshalIndent(buildTraceResultModel(traceResult), "", "  ")
	if err != nil {
		return err
	}
	_, err = writer.Write(data)
	return err
}
