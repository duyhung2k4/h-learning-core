package request

import (
	"app/model"

	"github.com/lib/pq"
)

type CreateQuizzRequest struct {
	Ask        string         `json:"ask"`
	ResultType string         `json:"resultType"`
	Result     pq.StringArray `json:"result" gorm:"type:text[]"`
	Pption     pq.StringArray `json:"option" gorm:"type:text[]"`
	Time       int            `json:"time"`

	EntityType model.ENTITY_TYPE `json:"entityType"`
	EntityId   uint              `json:"entityId"`
}
