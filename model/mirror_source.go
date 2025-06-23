package model

type MirrorSource struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	HomeURL     string `json:"homeUrl"`
}

func (source *MirrorSource) GetName() string {
	return source.Name
}

func (source *MirrorSource) GetURL() string {
	return source.HomeURL
}
