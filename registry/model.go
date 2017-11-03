package registry

type AccessToken struct {
	Token     string `json:"token"`
	ExpiresIn int    `json:"expires_in"`
	IssuedAt  string `json:"issued_at"`
}

type CatalogResp struct {
	Repositories []string `json:"repositories"`
}

type Config struct {
	MediaType string `json: "mediaType"`
	Size      int    `json: "size"`
	Digest    string `json: "digest"`
}
type Manifest struct {
	SchemaVersion int      `json: "schemaVersion"`
	MediaType     string   `json: "mediaType"`
	ImageConfig   Config   `json: "config"`
	Layers        []Config `json: "layers"`
}
