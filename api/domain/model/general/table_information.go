package general

import (
	"api/domain/model/user"
	"time"
)

type TableInformation struct {
	ID        TableInformationID `json:"tableInformationID,omitempty"`
	CreatedAt time.Time          `json:"tableInformationCreatedAt,omitempty"`
	CreatedBy user.UserID        `json:"tableInformationCreatedBy,omitempty"`
	UpdatedAt time.Time          `json:"tableInformationUpdatedAt,omitempty"`
	UpdatedBy user.UserID        `json:"tableInformationUpdatedBy,omitempty"`
}

type TableInformationID string
