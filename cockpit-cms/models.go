package cockpit_cms

type Collection struct {
	Name       string        `json:"name,omitempty"`
	Label      string        `json:"label,omitempty"`
	Id         string        `json:"_id,omitempty"`
	Fields     []Field       `json:"fields,omitempty"`
	Sortable   bool          `json:"sortable,omitempty"`
	InMenu     bool          `json:"in_menu,omitempty"`
	Created    int           `json:"_created,omitempty"`
	Modified   int           `json:"_modified,omitempty"`
	Color      string        `json:"color,omitempty"`
	Acl        []interface{} `json:"acl,omitempty"`
	Sort       Sort          `json:"sort,omitempty"`
	Rules      Rule          `json:"rules,omitempty"`
	ItemsCount int           `json:"itemsCount,omitempty"`
}

type Field struct {
	Name     string        `json:"name,omitempty"`
	Label    string        `json:"label,omitempty"`
	Type     string        `json:"type,omitempty"`
	Default  string        `json:"default,omitempty"`
	Info     string        `json:"info,omitempty"`
	Group    string        `json:"group,omitempty"`
	Localize bool          `json:"localize,omitempty"`
	Options  []interface{} `json:"options,omitempty"`
	Width    string        `json:"width,omitempty"`
	Lst      bool          `json:"lst,omitempty"`
	Acl      []interface{} `json:"acl,omitempty"`
}

type Sort struct {
	Column string `json:"column,omitempty"`
	Dir    int    `json:"dir,omitempty"`
}

type RuleSet struct {
	Enabled *bool `json:"enabled,omitempty"`
}
type Rule struct {
	Create RuleSet `json:"create,omitempty"`
	Read   RuleSet `json:"read,omitempty"`
	Update RuleSet `json:"update,omitempty"`
	Delete RuleSet `json:"delete,omitempty"`
}

type CreateCollection struct {
	Name string     `json:"name,omitempty"`
	Data Collection `json:"data,omitempty"`
}
