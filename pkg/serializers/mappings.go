package serializers

import (
	"fmt"
	"golang-tracer/pkg/tracer"
)

func buildTraceResultModel(res *tracer.TraceResult) (result *TraceResult) {
	result = &TraceResult{
		Name:    res.Name,
		Class:   res.Class,
		Time:    fmt.Sprintf("%s", res.Time),
		Methods: nil,
	}

	for _, m := range res.Methods {
		result.Methods = append(result.Methods, buildTraceResultModel(m))
	}
	
	return
}
