# golang-tracer
### A method execution time measurer

## Tracing Methods
The `tracer` package implements the following interface for tracing methods:

```go
type ITracer interface {
    StartTrace()                    // Called at the beginning of the method being measured
    StopTrace()                     // Called at the end of the method being measured
    GetTraceResult() *TraceResult   // Retrieve the measurement results
}
```

The Tracer collects the following information about the measured method:
- Method name
- Class name with the measured method
- Execution time of the method

The `TraceResult` struct contains the following information about the measured method:
```go
type TraceResult struct {
    Name    string             // Method name
    Class   string             // Class name with the measured method
    Time    time.Duration      // Execution time of the method
    Methods []*TraceResult     // Nested method traces
}
```

## Result Serialization
The measurement results can be serialized in three formats: JSON, XML, and YAML. Serialization is performed using classes that implement the following interface:
```go
type ITraceResultSerializer interface {
    Serialize(traceResult *tracer.TraceResult, writer io.Writer) error
}
```

## Usage Example
```go
func main() {

	myTracer := tracer.NewTracer()
	foo := entities.NewFoo(myTracer)

	foo.DoSomethingFoo()
	foo.DoSomethingElseFoo()

	traceResult := myTracer.GetTraceResult()

	serializer, err := serializers.GetSerializer(serializers.FormatJSON)
	if err != nil {
		panic(err)
	}

	err = serializer.Serialize(traceResult, os.Stdout)
	if err != nil {
		panic(err)
	}
}
```

## JSON output:
```json
{
  "name": "root",
  "time": "60.205459ms",
  "methods": [
    {
      "name": "DoSomethingFoo",
      "class": "(*Foo)",
      "time": "60.205459ms",
      "methods": [
        {
          "name": "DoSomethingBar",
          "class": "(*Bar)",
          "time": "11.092417ms"
        },
        {
          "name": "DoSomethingBaz",
          "class": "(*Baz)",
          "time": "66.252668ms",
          "methods": [
            {
              "name": "DoSomethingQux",
              "class": "(*Qux)",
              "time": "33.153959ms",
              "methods": [
                {
                  "name": "DoSomethingQuz",
                  "class": "(*Quz)",
                  "time": "11.069209ms"
                }
              ]
            }
          ]
        }
      ]
    },
    {
      "name": "DoSomethingElseFoo",
      "class": "(*Foo)",
      "time": "38.096459ms",
      "methods": [
        {
          "name": "DoSomethingBar",
          "class": "(*Bar)",
          "time": "11.028334ms"
        }
      ]
    }
  ]
}
```