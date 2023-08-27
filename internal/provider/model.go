package provider

type Collection struct {
	Id         string  `json:"_id"`
	Name       string  `json:"name"`
	Label      string  `json:"label"`
	InMenu     bool    `json:"in_menu"`
	Created    int     `json:"_created,omitempty"`
	Modified   int     `json:"_modified,omitempty"`
	ItemsCount int     `json:"itemsCount,omitempty"`
	Sort       Sort    `json:"sort"`
	Fields     []Field `json:"fields"`
}

type Sort struct {
	Column string `json:"column"`
	Dir    int    `json:"dir"`
}

type Field struct {
	Name  string `json:"name"`
	Label string `json:"label"`
	Type  string `json:"type"`
}

type CreateCollection struct {
	Name string     `json:"name"`
	Data Collection `json:"data"`
}

type UpdateCollection struct {
	Data Collection `json:"data"`
}
