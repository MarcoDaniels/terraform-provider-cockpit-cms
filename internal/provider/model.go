package provider

type Collection struct {
	Id         string        `json:"_id"`
	Name       string        `json:"name"`
	Label      string        `json:"label"`
	InMenu     bool          `json:"in_menu"`
	Created    int           `json:"_created,omitempty"`
	Modified   int           `json:"_modified,omitempty"`
	ItemsCount int           `json:"itemsCount,omitempty"`
	Fields     []Field       `json:"fields"`
}

type Field struct {
	Name     string        `json:"name"`
	Label    string        `json:"label"`
	Type     string        `json:"type"`
}
