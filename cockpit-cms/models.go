package cockpit_cms

type Collection struct {
	Name       string        `json:"name"`
	Label      string        `json:"label"`
	Id         string        `json:"_id"`
	Fields     []Field       `json:"fields"`
	Sortable   bool          `json:"sortable"`
	InMenu     bool          `json:"in_menu"`
	Created    int           `json:"_created"`
	Modified   int           `json:"_modified"`
	Color      string        `json:"color"`
	Acl        []interface{} `json:"acl"`
	Sort       Sort          `json:"sort"`
	Rules      Rule          `json:"rules"`
	ItemsCount int           `json:"itemsCount,omitempty"`
}

type Field struct {
	Name     string        `json:"name"`
	Label    string        `json:"label"`
	Type     string        `json:"type"`
	Default  string        `json:"default"`
	Info     string        `json:"info"`
	Group    string        `json:"group"`
	Localize bool          `json:"localize"`
	Options  []interface{} `json:"options"`
	Width    string        `json:"width"`
	Lst      bool          `json:"lst"`
	Acl      []interface{} `json:"acl"`
}

type Sort struct {
	Column string `json:"column"`
	Dir    int    `json:"dir"`
}

type RuleSet struct {
	Enabled bool `json:"enabled"`
}
type Rule struct {
	Create RuleSet `json:"create"`
	Read   RuleSet `json:"read"`
	Update RuleSet `json:"update"`
	Delete RuleSet `json:"delete"`
}

type CreateCollection struct {
	Name string     `json:"name"`
	Data Collection `json:"data"`
}
