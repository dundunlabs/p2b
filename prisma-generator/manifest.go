package prisma

type Manifest struct {
	PrettyName    string `json:"prettyName"`
	DefaultOutput string `json:"defaultOutput"`
}

type ManifestResult struct {
	Manifest Manifest `json:"manifest"`
}
