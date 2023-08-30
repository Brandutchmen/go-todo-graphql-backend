// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type NewTodo struct {
	Text string `json:"text"`
}

type Todo struct {
	ID     string `json:"id"`
	Text   string `json:"text"`
	Status Status `json:"status"`
	Done   bool   `json:"done"`
}

type Status string

const (
	StatusPending   Status = "PENDING"
	StatusComplete  Status = "COMPLETE"
	StatusCancelled Status = "CANCELLED"
)

var AllStatus = []Status{
	StatusPending,
	StatusComplete,
	StatusCancelled,
}

func (e Status) IsValid() bool {
	switch e {
	case StatusPending, StatusComplete, StatusCancelled:
		return true
	}
	return false
}

func (e Status) String() string {
	return string(e)
}

func (e *Status) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Status(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Status", str)
	}
	return nil
}

func (e Status) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
