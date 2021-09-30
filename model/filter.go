package model

type FilterEqualTypeInput struct {
	Eq *string   `json:"eq"`
	In *[]string `json:"in"`
}

type FilterMatchTypeInput struct {
	Match *string `json:"match"`
}

type FilterRangeTypeInput struct {
	From *string `json:"from"`
	To   *string `json:"to"`
}

type FilterConfigBase struct {
	Search *string `json:"search"`
	Limit  int     `json:"limit"`
	Page   int     `json:"page"`
}

func (f *FilterConfigBase) ToOffset() int {
	if f.Page == 1 {
		return 0
	}

	return f.Limit * (f.Page - 1)
}
