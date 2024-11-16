package request

import (
	"app/constant"
)

type QueryReq[T any] struct {
	Data        T                   `json:"data"`
	Datas       []T                 `json:"datas"`
	Args        []interface{}       `json:"args"`
	Condition   string              `json:"condition"`
	Preload     map[string]*string  `json:"preload"`
	Omit        map[string][]string `json:"omit"`
	Method      constant.METHOD     `json:"method"`
	Order       string              `json:"order"`
	Unscoped    bool                `json:"unscoped"`
	PreloadNull bool                `json:"preloadNull"`
}

type FindPayload struct {
	Condition string
	Preload   map[string]*string
	Omit      map[string][]string
	Order     string
	Agrs      []interface{}
}

type FirstPayload struct {
	Condition string
	Preload   map[string]*string
	Omit      map[string][]string
	Agrs      []interface{}
}
