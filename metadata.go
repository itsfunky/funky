package funky

var (
	// FunctionName is the functions name set at build time.
	FunctionName = ""
)

type metadata struct {
	FunctionName string `json:"function_name"`
}

// Metadata provides the available function metadata.
func Metadata() metadata {
	return metadata{
		FunctionName: FunctionName,
	}
}
