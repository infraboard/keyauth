package geoip

import (
	"fmt"
	"strconv"
	"strings"
)

// Record todo
type Record struct {
	*IPv4     `bson:",inline"`
	*Location `bson:",inline"`
}

// ParseLocationFromCsvLine todo
func ParseLocationFromCsvLine(line string) (*Location, error) {
	row := strings.Split(line, ",")
	if len(row) != 14 {
		return nil, fmt.Errorf("row length not eqaul 14 (%d) check csv data source, row: %s", len(row), line)
	}

	return &Location{
		GeonameID:      row[0],
		LocaleCode:     row[1],
		ContinentCode:  row[2],
		ContinentName:  row[3],
		CountryISOCode: row[4],
		CountryName:    row[5],
		CityName:       row[10],
	}, nil
}

// Location todo
type Location struct {
	GeonameID      string `bson:"_id" json:"geoname_id"`
	LocaleCode     string `bson:"locale_code" json:"locale_code"`
	ContinentCode  string `bson:"continent_code" json:"continent_code"`
	ContinentName  string `bson:"continent_name" json:"continent_name"`
	CountryISOCode string `bson:"country_iso_code" json:"country_iso_code"`
	CountryName    string `bson:"country_name" json:"country_name"`
	CityName       string `bson:"city_name" json:"city_name"`
}

// NewLocationSet todo
func NewLocationSet(capacity uint) *LocationSet {
	return &LocationSet{
		items:    make([]*Location, 0, capacity),
		capacity: capacity,
	}
}

// LocationSet todo
type LocationSet struct {
	items    []*Location
	capacity uint
	length   uint
}

// Add todo
func (s *LocationSet) Add(item *Location) {
	s.items = append(s.items, item)
	s.length++
}

// IsFull todo
func (s *LocationSet) IsFull() bool {
	return s.length == s.capacity
}

// Reset todo
func (s *LocationSet) Reset() {
	s.items = make([]*Location, 0, s.capacity)
	s.length = 0
}

// Items todo
func (s *LocationSet) Items() []*Location {
	return s.items
}

// Length tood
func (s *LocationSet) Length() uint {
	return s.length
}

// IPv4 todo
type IPv4 struct {
	Network             string  `bson:"_id" json:"network"`
	GeonameID           string  `bson:"geoname_id" json:"geoname_id"`
	IsAnonymousProxy    bool    `bson:"is_anonymous_proxy" json:"is_anonymous_proxy"`
	IsSatelliteProvider bool    `bson:"is_satellite_provider" json:"is_satellite_provider"`
	Latitude            float64 `bson:"latitude" json:"latitude"`
	Longitude           float64 `bson:"longitude" json:"longitude"`
	AccuracyRadius      int64   `bson:"accuracy_radius" json:"accuracy_radius"`
}

// ParseIsAnonymousProxy todo
func (i *IPv4) ParseIsAnonymousProxy(is string) {
	if is == "0" {
		i.IsAnonymousProxy = false
	} else {
		i.IsAnonymousProxy = true
	}
}

// ParseIsSatelliteProvider tood
func (i *IPv4) ParseIsSatelliteProvider(is string) {
	if is == "0" {
		i.IsSatelliteProvider = false
	} else {
		i.IsSatelliteProvider = true
	}
}

// ParseLatitude todo
func (i *IPv4) ParseLatitude(lat string) {
	i.Latitude, _ = strconv.ParseFloat(strings.TrimSpace(lat), 32)
}

// ParseLongitude todo
func (i *IPv4) ParseLongitude(lon string) {
	i.Longitude, _ = strconv.ParseFloat(strings.TrimSpace(lon), 32)
}

// ParseAccuracyRadius todo
func (i *IPv4) ParseAccuracyRadius(radius string) {
	i.AccuracyRadius, _ = strconv.ParseInt(strings.TrimSpace(radius), 10, 32)
}

// ParseIPv4FromCsvLine todo
func ParseIPv4FromCsvLine(line string) (*IPv4, error) {
	row := strings.Split(line, ",")
	if len(row) != 10 {
		return nil, fmt.Errorf("row length not eqaul 10 (%d) check csv data source, row: %s", len(row), line)
	}

	ipv4 := &IPv4{
		Network:   row[0],
		GeonameID: row[1],
	}

	ipv4.ParseIsAnonymousProxy(row[4])
	ipv4.ParseIsSatelliteProvider(row[5])
	ipv4.ParseLatitude(row[7])
	ipv4.ParseLongitude(row[8])
	ipv4.ParseAccuracyRadius(row[9])
	return ipv4, nil
}

// NewIPv4Set todo
func NewIPv4Set(capacity uint) *IPv4Set {
	return &IPv4Set{
		items:    make([]*IPv4, 0, capacity),
		capacity: capacity,
	}
}

// IPv4Set todo
type IPv4Set struct {
	items    []*IPv4
	capacity uint
	length   uint
}

// Add todo
func (s *IPv4Set) Add(item *IPv4) {
	s.items = append(s.items, item)
	s.length++
}

// IsFull todo
func (s *IPv4Set) IsFull() bool {
	return s.length == s.capacity
}

// Reset todo
func (s *IPv4Set) Reset() {
	s.items = make([]*IPv4, 0, s.capacity)
	s.length = 0
}

// Length tood
func (s *IPv4Set) Length() uint {
	return s.length
}

// Items tood
func (s *IPv4Set) Items() []*IPv4 {
	return s.items
}
