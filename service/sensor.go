package service

import "satplan/entity"

func AddSensor(sensorIn *entity.SensorInDTO) (int, error) {

	return 0, nil
}

func GetAllSensors() *[]entity.Sensor {

	return nil
}

func GetSensorBySatId(satId string) (*entity.Sensor, error) {
	return nil, nil
}

func GetSensorById(senId string) *entity.Sensor {
	return nil
}

func UpdateSensor(satId string, senInDTO *entity.SensorInDTO) error {
	return nil
}

func DeleteSensorById(satId string) error {
	return nil
}
