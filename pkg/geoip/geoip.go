package geoip

// The Record struct corresponds to the data in the GeoIP2/GeoLite2 Result
// databases.
type Record struct {
	Continent          Continent          `maxminddb:"continent"`
	City               City               `maxminddb:"city"`
	Country            Country            `maxminddb:"country"`
	Location           Location           `maxminddb:"location"`
	RegisteredCountry  RegisteredCountry  `maxminddb:"registered_country"`
	RepresentedCountry RepresentedCountry `maxminddb:"represented_country"`
	Subdivisions       []Subdivisions     `maxminddb:"subdivisions"`
	Traits             Traits             `maxminddb:"traits"`
}

// Continent 标识符
type Continent struct {
	Code      string            `maxminddb:"code"`
	GeoNameID uint              `maxminddb:"geoname_id"`
	Names     map[string]string `maxminddb:"names"`
}

// City 城市
type City struct {
	GeoNameID uint              `maxminddb:"geoname_id"`
	Names     map[string]string `maxminddb:"names"`
}

// Country 国家
type Country struct {
	GeoNameID         uint              `maxminddb:"geoname_id"`
	IsInEuropeanUnion bool              `maxminddb:"is_in_european_union"`
	IsoCode           string            `maxminddb:"iso_code"`
	Names             map[string]string `maxminddb:"names"`
}

// Location 位置信息
type Location struct {
	AccuracyRadius uint16  `maxminddb:"accuracy_radius"`
	Latitude       float64 `maxminddb:"latitude"`
	Longitude      float64 `maxminddb:"longitude"`
	MetroCode      uint    `maxminddb:"metro_code"`
	TimeZone       string  `maxminddb:"time_zone"`
}

// RegisteredCountry todo
type RegisteredCountry struct {
	GeoNameID         uint              `maxminddb:"geoname_id"`
	IsInEuropeanUnion bool              `maxminddb:"is_in_european_union"`
	IsoCode           string            `maxminddb:"iso_code"`
	Names             map[string]string `maxminddb:"names"`
}

// RepresentedCountry todo
type RepresentedCountry struct {
	GeoNameID         uint              `maxminddb:"geoname_id"`
	IsInEuropeanUnion bool              `maxminddb:"is_in_european_union"`
	IsoCode           string            `maxminddb:"iso_code"`
	Names             map[string]string `maxminddb:"names"`
	Type              string            `maxminddb:"type"`
}

// Subdivisions todo
type Subdivisions struct {
	GeoNameID uint              `maxminddb:"geoname_id"`
	IsoCode   string            `maxminddb:"iso_code"`
	Names     map[string]string `maxminddb:"names"`
}

// Traits todo
type Traits struct {
	IsAnonymousProxy    bool `maxminddb:"is_anonymous_proxy"`
	IsSatelliteProvider bool `maxminddb:"is_satellite_provider"`
}
