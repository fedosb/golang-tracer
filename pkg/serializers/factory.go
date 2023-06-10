package serializers

import "errors"

type serializationFormatType string

const (
	FormatJSON = "json"
	FormatXML  = "xml"
)

func GetSerializer(format serializationFormatType) (ITraceResultSerializer, error) {
	switch format {
	case FormatJSON:
		return newJSONTraceResultSerializer(), nil
	case FormatXML:
		return newXMLTraceResultSerializer(), nil
	default:
		return nil, errors.New("UNKNOWN SERIALIZATION FORMAT")
	}
}
