package secrets

// Secret represents an app secret
type Secret struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}
