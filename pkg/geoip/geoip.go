package geoip

// Record todo
type Record struct {
	*IPv4     `bson:",inline"`
	*Location `bson:",inline"`
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

// IPv4 todo
type IPv4 struct {
	GeonameID           string  `bson:"_id" json:"geoname_id"`
	Network             string  `bson:"network" json:"network"`
	IsAnonymousProxy    bool    `bson:"is_anonymous_proxy" json:"is_anonymous_proxy"`
	IsSatelliteProvider bool    `bson:"is_satellite_provider" json:"is_satellite_provider"`
	Latitude            float32 `bson:"latitude" json:"latitude"`
	Longitude           float32 `bson:"longitude" json:"longitude"`
	AccuracyRadius      uint    `bson:"accuracy_radius" json:"accuracy_radius"`
}
