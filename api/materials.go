package main

type Location struct {
	Lat      float64 `json:"lat"`
	Lon      float64 `json:"lon"`
	SiteName string  `json:"siteName"`
}

type Material struct {
	Name      string     `json:"name"`
	Locations []Location `json:"locations"`
}
