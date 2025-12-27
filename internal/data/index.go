package data

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
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

func Load(pathway string) ([]MRDSRecord, error) {
	return generateRecords(pathway)
}

func generateRecords(csvPath string) ([]MRDSRecord, error) {
	var records []MRDSRecord

	f, err := os.Open(csvPath)
	if err != nil {
		fmt.Printf("Failed to open %s: %v\n", csvPath, err)
		return nil, err
	}
	defer f.Close()

	r := csv.NewReader(f)
	r.FieldsPerRecord = -1

	if _, err := r.Read(); err != nil {
		f.Close()
		fmt.Printf("Failed to read header of %s: %v\n", csvPath, err)
		return nil, err
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

		depID, err := strconv.ParseInt(record[0], 10, 64)
		if err != nil {
			fmt.Printf("invalid DepID %q\n", record[0])
			continue
		}

		lat, err := strconv.ParseFloat(record[5], 64)
		if err != nil {
			fmt.Printf("invalid latitude %q found at depID %q\n", record[5], record[0])
			continue
		}

		lon, err := strconv.ParseFloat(record[6], 64)
		if err != nil {
			fmt.Printf("invalid longitude %q found at depID %q\n", record[6], record[0])
			continue
		}

		score, _ := strconv.ParseFloat(record[45], 64) // score can be empty

		rec := MRDSRecord{
			DepID:      depID,
			URL:        record[1],
			MRDSID:     record[2],
			MASID:      record[3],
			SiteName:   record[4],
			Latitude:   lat,
			Longitude:  lon,
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

		records = append(records, rec)
	}

	return records, nil
}
