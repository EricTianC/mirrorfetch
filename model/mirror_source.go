package model

type MirrorSource struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	HomeUrl     string `json:"homeUrl"`
}

func (source *MirrorSource) GetName() string {
	return source.Name
}

func (source *MirrorSource) GetHomeUrl() string {
	return source.HomeUrl
}
