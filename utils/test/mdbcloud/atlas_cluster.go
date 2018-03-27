package mdbcloud

// AtlasCluster represents an Atlas cluster
type AtlasCluster struct {
	ID               string           `json:"id"`
	Name             string           `json:"name"`
	MongoURI         string           `json:"mongoURI"`
	ProviderSettings ProviderSettings `json:"providerSettings"`
}

// ProviderSettings represents the providerSettings in an atlas cluster
type ProviderSettings struct {
	InstanceSize string `json:"instanceSizeName"`
}
