package dto

import (
	"github.com/achillescres/pkg/object/oid"
	"saina.gitlab.yandexcloud.net/saina/backend/saina-api/internal/domain/entity"
)

type AirportCreate struct {
	Create `json:"-" db:"-" csv:"-" binding:"-"`

	IATA        string `json:"iata" db:"iata" csv:"code" binding:"required"`
	TimeZone    string `json:"timeZone" db:"time_zone" csv:"time_zone" binding:"required"`
	Name        string `json:"name" db:"name" csv:"name" binding:"required"`
	City        string `json:"city" db:"city" csv:"city" binding:"required"`
	CityCode    string `json:"cityCode" db:"city_code" csv:"city_code" binding:"required"`
	CountryCode string `json:"countryCode" db:"country_code" csv:"country_code" binding:"required"`
	Location    string `json:"location" db:"location" csv:"location" binding:"required"`
	Elevation   string `json:"elevation" db:"elevation" csv:"elevation" binding:"required"`
	ICAO        string `json:"icao" db:"icao" csv:"icao" binding:"required"`
	County      string `json:"county" db:"county" csv:"county" binding:"required"`
	State       string `json:"state" db:"state" csv:"state" binding:"required"`
}

func NewAirportCreate(IATA string, timeZone string, name string, city string, cityCode string, countryCode string, location string, elevation string, ICAO string, county string, state string) *AirportCreate {
	return &AirportCreate{IATA: IATA, TimeZone: timeZone, Name: name, City: city, CityCode: cityCode, CountryCode: countryCode, Location: location, Elevation: elevation, ICAO: ICAO, County: county, State: state}
}

func (a *AirportCreate) ToEntity(id oid.ID) *entity.Airport {
	return entity.NewAirport(
		id,
		a.IATA,
		a.TimeZone,
		a.Name,
		a.City,
		a.CityCode,
		a.CountryCode,
		a.Location,
		a.Elevation,
		a.ICAO,
		a.County,
		a.State,
	)
}
