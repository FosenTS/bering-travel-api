package dto

import (
	"github.com/achillescres/pkg/object/oid"
	"saina.gitlab.yandexcloud.net/saina/backend/saina-api/internal/domain/entity"
	"time"
)

type BookingDataParsingErrorCreate struct {
	Create `json:"-" db:"-" binding:"-"`

	BookingDataFileID oid.ID              `json:"bookingDataFileID" db:"booking_data_file_id" binding:"required"`
	Error             string              `json:"error" db:"error" binding:"required"`
	ErrorType         entity.ParsingError `json:"errorType" db:"error_type" binding:"required"`
	RowIndex          uint                `json:"rowIndex" db:"row_index" binding:"required"`
	Time              time.Time           `json:"time" db:"snatch_time" binding:"required"`
}

func NewBookingDataParsingErrorCreate(bookingDataFileID oid.ID, error string, errorType entity.ParsingError, rowIndex uint, time time.Time) *BookingDataParsingErrorCreate {
	return &BookingDataParsingErrorCreate{BookingDataFileID: bookingDataFileID, Error: error, ErrorType: errorType, RowIndex: rowIndex, Time: time}
}

func (tpeC *BookingDataParsingErrorCreate) ToEntity(id oid.ID) *entity.BookingDataParsingError {
	return entity.NewBookingDataParsingError(id, tpeC.BookingDataFileID, tpeC.Error, tpeC.ErrorType, tpeC.RowIndex)
}
