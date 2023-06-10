package serializers

type TraceResult struct {
	Name    string         `json:"name" xml:"name" yaml:"name"`
	Class   string         `json:"class,omitempty" xml:"class,omitempty" yaml:"class,omitempty"`
	Time    string         `json:"time" xml:"time" yaml:"time"`
	Methods []*TraceResult `json:"methods,omitempty" xml:"methods,omitempty" yaml:"methods"`
}
