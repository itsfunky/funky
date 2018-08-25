package providers

var (
	// Providers contains the configurations for each registered provider.
	Providers = make(map[string]Provider)

	// Names contains a list of registered provider names.
	Names = []string{}
)

// Provider represents configuration for a given provider.
type Provider struct {
	Shim    bool
	Compile string
}

// Register adds a provider configuration to the Providers list.
func Register(name string, config Provider) {
	// Append the provider name to the list once.
	if _, ok := Providers[name]; !ok {
		Names = append(Names, name)
	}

	Providers[name] = config
}
