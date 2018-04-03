package mdbcloud

// AtlasIPWhitelistEntry represents an Atlas Group IP Whitelist entry
type AtlasIPWhitelistEntry struct {
	CIDRBlock string `json:"cidrBlock"`
	Comment   string `json:"comment"`
}
