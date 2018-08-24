package funky

var (
	// FunctionName is the functions name set at build time.
	FunctionName = ""
)

// FunctionMetadata represents the invoked functions metadata.
type FunctionMetadata struct {
	FunctionName string `json:"function_name"`
}

// Metadata provides the available function metadata.
func Metadata() FunctionMetadata {
	return FunctionMetadata{
		FunctionName: FunctionName,
	}
}
