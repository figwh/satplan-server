package service

import "satplan/entity"

func AddSatellite(tle *entity.Tle) (int, error) {

	return 0, nil
}

func GetAllSatellites() *[]entity.Satellite {

	return nil
}

func GetSatelliteById(satId string) (*entity.Satellite, error) {
	return nil, nil
}

func UpdateSatellite(satId string, satDTO *entity.SatDTO) error {
	return nil
}

func DeleteSatelliteById(satId string) error {
	return nil
}
