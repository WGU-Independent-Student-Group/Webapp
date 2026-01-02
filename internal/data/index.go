package data

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

type MRDSRecord struct {
	DepID     int64   `csv:"dep_id"`
	URL       string  `csv:"url"`
	MRDSID    string  `csv:"mrds_id"`
	MASID     string  `csv:"mas_id"`
	SiteName  string  `csv:"site_name"`
	Latitude  float64 `csv:"latitude"`
	Longitude float64 `csv:"longitude"`
	Region    string  `csv:"region"`
	Country   string  `csv:"country"`
	State     string  `csv:"state"`
	County    string  `csv:"county"`
	ComType   string  `csv:"com_type"`

	Commod1 string `csv:"commod1"`
	Commod2 string `csv:"commod2"`
	Commod3 string `csv:"commod3"`

	OperType string `csv:"oper_type"`
	DepType  string `csv:"dep_type"`
	ProdSize string `csv:"prod_size"`
	DevStat  string `csv:"dev_stat"`

	Ore       string `csv:"ore"`
	Gangue    string `csv:"gangue"`
	OtherMatl string `csv:"other_matl"`
	OrebodyFm string `csv:"orebody_fm"`

	WorkType   string `csv:"work_type"`
	Model      string `csv:"model"`
	Alteration string `csv:"alteration"`
	ConcProc   string `csv:"conc_proc"`

	Names    string `csv:"names"`
	OreCtrl  string `csv:"ore_ctrl"`
	Reporter string `csv:"reporter"`

	HRockUnit string `csv:"hrock_unit"`
	HRockType string `csv:"hrock_type"`
	ARockUnit string `csv:"arock_unit"`
	ARockType string `csv:"arock_type"`

	Structure string `csv:"structure"`
	Tectonic  string `csv:"tectonic"`
	Ref       string `csv:"ref"`

	YFPBA    string `csv:"yfp_ba"`
	YrFstPrd string `csv:"yr_fst_prd"`
	YLPBA    string `csv:"ylp_ba"`
	YrLstPrd string `csv:"yr_lst_prd"`
	DYBA     string `csv:"dy_ba"`
	DiscYr   string `csv:"disc_yr"`
	ProdYrs  string `csv:"prod_yrs"`

	Discr string  `csv:"discr"`
	Score float64 `csv:"score"`
}

type Metadata struct {
	MetadataID int64
	Url        string
	Score      float64
}

type SiteIdentification struct {
	DepID      int64
	MRDSID     string
	MASID      string
	SiteName   string
	MetadataID int64
}

type SiteCharacteristics struct {
	SiteCharID int64
	DepID      int64
	OperType   string
	DepType    string
	ProdSize   string
	DevStat    string
	Ore        string
	OreBody    string
	Gangue     string
	OtherMatl  string
	WorkType   string
	Model      string
	Alteration string
}

type GeographicLocation struct {
	GeoID     int64
	DepID     int64
	Latitude  float64
	Longitude float64
	Country   string
	State     string
	County    string
}

var (
	metaIDCounter     int64 = 1
	siteCharIDCounter int64 = 1
	geoIDCounter      int64 = 1
)

func Load(pathway string) ([]Metadata, []SiteIdentification, []SiteCharacteristics, []GeographicLocation, error) {
	return generateRecords(pathway)
}

func generateRecords(csvPath string) ([]Metadata, []SiteIdentification, []SiteCharacteristics, []GeographicLocation, error) {
	//var records []MRDSRecord
	var metadataEntries []Metadata
	var siteIdentEntries []SiteIdentification
	var siteCharEntries []SiteCharacteristics
	var geoEntries []GeographicLocation

	f, err := os.Open(csvPath)
	if err != nil {
		fmt.Printf("Failed to open %s: %v\n", csvPath, err)
		return nil, nil, nil, nil, err
	}
	defer f.Close()

	r := csv.NewReader(f)
	r.FieldsPerRecord = -1

	if _, err := r.Read(); err != nil {
		f.Close()
		fmt.Printf("Failed to read header of %s: %v\n", csvPath, err)
		return nil, nil, nil, nil, err
	}

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("error reading %s: %v\n", csvPath, err)
			continue
		}

		if len(record) < 46 {
			continue
		}

		depID, err := parseIntField(record[0], "DepID")
		if err != nil {
			fmt.Println(err)
			continue
		}

		validLon, validLat, err := validateCoordinates(record[5], record[6])
		if err != nil {
			fmt.Println(err)
			continue
		}

		usa := validateCountry(record[8])
		if !usa {
			continue
		}

		validState := validateState(record[9])
		if !validState {
			continue
		}

		score, _ := parseFloatField(record[45], "score")

		rec := MRDSRecord{
			DepID:      depID,
			URL:        record[1],
			MRDSID:     record[2],
			MASID:      record[3],
			SiteName:   record[4],
			Latitude:   validLat,
			Longitude:  validLon,
			Region:     record[7],
			Country:    record[8],
			State:      record[9],
			County:     record[10],
			ComType:    record[11],
			Commod1:    record[12],
			Commod2:    record[13],
			Commod3:    record[14],
			OperType:   record[15],
			DepType:    record[16],
			ProdSize:   record[17],
			DevStat:    record[18],
			Ore:        record[19],
			Gangue:     record[20],
			OtherMatl:  record[21],
			OrebodyFm:  record[22],
			WorkType:   record[23],
			Model:      record[24],
			Alteration: record[25],
			ConcProc:   record[26],
			Names:      record[27],
			OreCtrl:    record[28],
			Reporter:   record[29],
			HRockUnit:  record[30],
			HRockType:  record[31],
			ARockUnit:  record[32],
			ARockType:  record[33],
			Structure:  record[34],
			Tectonic:   record[35],
			Ref:        record[36],
			YFPBA:      record[37],
			YrFstPrd:   record[38],
			YLPBA:      record[39],
			YrLstPrd:   record[40],
			DYBA:       record[41],
			DiscYr:     record[42],
			ProdYrs:    record[43],
			Discr:      record[44],
			Score:      score,
		}

		//records = append(records, rec)

		meta := createMetadata(rec)
		metadataEntries = append(metadataEntries, meta)

		siteID := createSiteIdentification(rec)
		siteIdentEntries = append(siteIdentEntries, siteID)

		siteChar := createSiteCharacteristics(rec)
		siteCharEntries = append(siteCharEntries, siteChar)

		geo := createGeographicLocation(rec)
		geoEntries = append(geoEntries, geo)

		metaIDCounter++
		siteCharIDCounter++
		geoIDCounter++
	}

	return metadataEntries, siteIdentEntries, siteCharEntries, geoEntries, nil
}

func createMetadata(record MRDSRecord) Metadata {
	return Metadata{
		MetadataID: metaIDCounter,
		Url:        record.URL,
		Score:      record.Score,
	}
}

func createSiteIdentification(record MRDSRecord) SiteIdentification {
	return SiteIdentification{
		DepID:      record.DepID,
		MRDSID:     record.MRDSID,
		MASID:      record.MASID,
		SiteName:   record.SiteName,
		MetadataID: metaIDCounter,
	}
}

func createSiteCharacteristics(record MRDSRecord) SiteCharacteristics {
	return SiteCharacteristics{
		SiteCharID: siteCharIDCounter,
		DepID:      record.DepID,
		OperType:   record.OperType,
		DepType:    record.DepType,
		ProdSize:   record.ProdSize,
		DevStat:    record.DevStat,
		Ore:        record.Ore,
		OreBody:    record.OrebodyFm,
		Gangue:     record.Gangue,
		OtherMatl:  record.OtherMatl,
		WorkType:   record.WorkType,
		Model:      record.Model,
		Alteration: record.Alteration,
	}
}

func createGeographicLocation(record MRDSRecord) GeographicLocation {
	return GeographicLocation{
		GeoID:     geoIDCounter,
		DepID:     record.DepID,
		Latitude:  record.Latitude,
		Longitude: record.Longitude,
		Country:   record.Country,
		State:     record.State,
		County:    record.County,
	}
}
