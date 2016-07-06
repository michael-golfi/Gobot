package types

type ActionParameterValue struct {
	Entity string  `json:"entity"`
	Score  float64 `json:"score"`
	Type   string  `json:"type"`
}

type ActionParameter struct {
	Name     string                 `json:"name"`
	Required bool                   `json:"required"`
	Value    []ActionParameterValue `json:"value"`
}

type Action struct {
	Name       string            `json:"name"`
	Triggered  bool              `json:"triggered"`
	Parameters []ActionParameter `json:"parameters"`
}

type IntentRecommendation struct {
	Intent  string   `json:"intent"`
	Score   float64  `json:"score"`
	Actions []Action `json:"actions"`
}

type EntityRecommendation struct {
	Role       string            `json:"role"`
	Entity     string            `json:"entity"`
	Type       string            `json:"type"`
	StartIndex string            `json:"startIndex"`
	EndIndex   string            `json:"endIndex"`
	Score      string            `json:"score"`
	Resolution map[string]string `json:"resolution"`
}

type LuisResult struct {
	Query    string                 `json:"query"`
	Intents  []IntentRecommendation `json:"intents"`
	Entities []EntityRecommendation `json:"entities"`
}