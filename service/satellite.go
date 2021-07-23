package service

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"satplan/common"
	"satplan/dao/db"
	"satplan/entity"
	"strings"
)

func AddSatellite(tle *entity.TleData) (int, error) {
	satName := strings.TrimSpace(tle.Line0)

	line1Details := strings.Split(tle.Line1, " ")
	noardId := line1Details[1]

	newSat := entity.Satellite{
		Name:     satName,
		NoardId:  noardId,
		OleColor: 0,
	}
	err := db.CreateSatellite(&newSat)

	if err != nil {
		return 0, err
	}
	err = db.CreateTle(&entity.Tle{
		SatNoardId: noardId,
		Time:       common.GetUtcNowTimeStampSec(),
		Line1:      tle.Line1,
		Line2:      tle.Line2,
	})

	return newSat.Id, err
}

func GetAllSatellites() *[]entity.Satellite {
	return db.FindAllSatellites()
}

func GetSatelliteById(satId string) (*entity.Satellite, error) {
	return db.FindSatelliteByNoardId(satId)
}

func UpdateSatellite(satId string, satDTO *entity.SatDTO) error {
	satInDB, err := db.FindSatelliteByNoardId(satId)
	if err != nil || satInDB.Id == 0 {
		return errors.New("error finding satellite")
	}
	satInDB.Name = satDTO.SatName
	satInDB.OleColor = satDTO.OleColor
	return db.SaveSatellite(satInDB)
}

func DeleteSatelliteById(satId string) error {
	return db.DeleteSatelliteById(satId)
}

func UpdateTles() error {
	tleTxt, err := getNewTles()
	if err != nil {
		return err
	}
	tleDetails := strings.Split(tleTxt, "\n")
	tles := []entity.Tle{}
	for i := 0; i < len(tleDetails)-3; i += 3 {
		//get noard id
		line1Details := strings.Split(tleDetails[i+1], " ")
		noardId := line1Details[1]
		sat, _ := db.FindSatelliteByNoardId(noardId)
		if sat == nil || sat.Id == 0 {
			db.CreateSatellite(&entity.Satellite{
				Name:     strings.TrimSpace(tleDetails[i]),
				NoardId:  noardId,
				OleColor: 0,
			})
		}
		tles = append(tles, entity.Tle{
			SatNoardId: noardId,
			Time:       common.GetUtcNowTimeStampSec(),
			Line1:      tleDetails[i+1],
			Line2:      tleDetails[i+2],
		})
	}
	return db.BatCreateTle(&tles)
}

func getNewTles() (string, error) {
	url := "http://celestrak.com/NORAD/elements/resource.txt"
	/*
		SCD 1
		1 22490U 93009B   21202.41204679  .00000203  00000-0  53852-5 0  9991
		2 22490  24.9705  95.0049 0042752 280.5291 240.5150 14.44613837501644
		TECHSAT 1B (GO-32)
		1 25397U 98043D   21202.84242513 -.00000023  00000-0  95154-5 0  9991
		2 25397  98.7723 148.3036 0001524  31.1555 328.9713 14.23711884196605
	*/

	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return "", err
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return string(body), nil
}
