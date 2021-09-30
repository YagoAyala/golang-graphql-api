package model

import (
	"fmt"
	"io"
	"strconv"
)

type SortEnum string

const (
	SortEnumAsc  SortEnum = "ASC"
	SortEnumDesc SortEnum = "DESC"
)

var AllSortEnum = []SortEnum{
	SortEnumAsc,
	SortEnumDesc,
}

func (e SortEnum) IsValid() bool {
	switch e {
	case SortEnumAsc, SortEnumDesc:
		return true
	}
	return false
}

func (e SortEnum) String() string {
	return string(e)
}

func (e *SortEnum) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = SortEnum(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid SortEnum", str)
	}
	return nil
}

func (e SortEnum) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
