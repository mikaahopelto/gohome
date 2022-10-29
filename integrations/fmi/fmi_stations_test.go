package fmi

import (
	"encoding/xml"
	"io/ioutil"
	"testing"
)

func TestStations(t *testing.T) {
	data, err := ioutil.ReadFile("testdata/exampleStations.xml")
	if err != nil {
		t.Fatal("Error opening exampleStations.xml file")
	}
	fmi := &FMI_StationsModel{}
	s := &fmi.StationsCol
	err = xml.Unmarshal(data, s)
	if err != nil {
		t.Fatalf("Could not parse exampleStations.xml: %v", err)
	}
	err = fmi.Validate()
	if err != nil {
		t.Errorf("Stations model validation failed: %v", err)
	}
	if l := 452; len(s.Stations) != l {
		t.Errorf("Stations length, got %d, want, %d", len(s.Stations), l)
	}
	if id := "100539"; s.Stations[0].Id != id {
		t.Errorf("First Station Id, got %s, want %s", s.Stations[0].Id, id)
	}
	if point := "65.673370 24.515260"; s.Stations[0].Point != point {
		t.Errorf("Stations[0].Point, got %s, want %s", s.Stations[0].Point, point)
	}
	if key := "http://xml.fmi.fi/namespace/locationcode/name"; s.Stations[0].Names[0].Key != key {
		t.Errorf("Stations[0].Names[0].Key, got %s, want %s", s.Stations[0].Names[0].Key, key)
	}
	if name := "Kemi Ajos"; s.Stations[0].Names[0].Value != name {
		t.Errorf("Stations[0].Names[0].Value, got %s, want %s", s.Stations[0].Names[0].Value, name)
	}
	if id := "874863"; s.Stations[len(s.Stations)-1].Id != id {
		t.Errorf("Last Station Id, got %s, want %s", s.Stations[0].Id, id)
	}

	ws, err := ConvertToWeatherStations(*fmi)
	if err != nil {
		t.Errorf("Error converting to weather stations: %v", err)
	}
	err = ws.Validate()
	if err != nil {
		t.Errorf("Error validating weatherstation model: %v", err)
	}
}

func TestValidator(t *testing.T) {
	fmi := &FMI_StationsModel{}
	sc := &fmi.StationsCol
	LoadStationsXml(t, "testdata/exampleInvalid.xml", sc)
	err := fmi.Validate()
	if err == nil {
		t.Errorf("Expected error in validation")
	}
}

func LoadStationsXml(t *testing.T, fn string, sc *StationCollection) {
	data, err := ioutil.ReadFile(fn)
	if err != nil {
		t.Fatalf("Could not retrieve %s file: %v", fn, err)
	}
	err = xml.Unmarshal(data, sc)
	if err != nil {
		t.Fatalf("Could not retrieve %s file: %v", fn, err)
	}
}
