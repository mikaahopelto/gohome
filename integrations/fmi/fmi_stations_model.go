package fmi

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
)

type FMI_StationsModel struct {
	StationsCol StationCollection `xml:"FeatureCollection" validate:"required"`
}

type StationCollection struct {
	Stations []Station `xml:"member>EnvironmentalMonitoringFacility" validate:"required,dive"` // Weather stations
}
type Station struct {
	Id    StationId `xml:"identifier" validate:"required"`
	Names []Name    `xml:"name" validate:"gt=1,dive"`
	Point string    `xml:"representativePoint>Point>pos" validate:"required"`
}
type StationId string
type Name struct {
	Key   string `xml:"codeSpace,attr"`
	Value string `xml:",chardata"`
}

var (
	errValidate = fmt.Errorf("Validation error")
)

func (f FMI_StationsModel) Validate() error {
	validate := validator.New()
	err := validate.Struct(f)
	if err != nil {
		return errors.Wrap(err, "Validation error")
	}
	return nil
}

func ConvertToWeatherStations(s FMI_StationsModel) (WeatherStationModel, error) {
	wsm := WeatherStationModel{}
	for _, station := range s.StationsCol.Stations {
		weatherStation := WeatherStation{
			Id: string(station.Id),
		}
		for _, name := range station.Names {
			switch name.Key {
			case "http://xml.fmi.fi/namespace/locationcode/name":
				weatherStation.Name = name.Value
			case "http://xml.fmi.fi/namespace/location/region":
				weatherStation.Region = name.Value
			}
		}
		wsm.WeatherStations = append(wsm.WeatherStations, weatherStation)
	}
	return wsm, wsm.Validate()
}
