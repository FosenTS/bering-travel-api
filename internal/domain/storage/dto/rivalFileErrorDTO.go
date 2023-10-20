package dto

import (
	"github.com/achillescres/pkg/object/oid"
	"saina.gitlab.yandexcloud.net/saina/backend/saina-api/internal/domain/entity"
)

type RivalFileErrorCreate struct {
	Create `json:"-" db:"-" binding:"-"`

	RivalFileID oid.ID `json:"rivalFlightsFileID" db:"rival_flights_file_id" binding:"required"`
	Error       string `json:"error" db:"error" binding:"required"`
	RowIndex    int    `json:"rowIndex" db:"row_index" binding:"required"`
}

func NewRivalFileErrorCreate(rivalFileID oid.ID, error string, rowIndex int) *RivalFileErrorCreate {
	return &RivalFileErrorCreate{RivalFileID: rivalFileID, Error: error, RowIndex: rowIndex}
}

func (rfeC *RivalFileErrorCreate) ToEntity(id oid.ID) *entity.RivalFileError {
	return entity.NewRivalFileError(
		id,
		rfeC.RivalFileID,
		rfeC.Error,
		rfeC.RowIndex,
	)
}
