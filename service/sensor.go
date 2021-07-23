package service

import (
	"errors"
	"satplan/dao/db"
	"satplan/entity"
)

func AddSensor(sensorIn *entity.SensorInDTO) (int, error) {
	sat, err := db.FindSatelliteByNoardId(sensorIn.SatId)
	if err != nil {
		return 0, err
	}
	if sat == nil || sat.Id == 0 {
		return 0, errors.New("cannot find sat: " + sensorIn.SatId)
	}
	sensor := entity.Sensor{
		SatNoardId:     sensorIn.SatId,
		SatName:        sat.Name,
		Name:           sensorIn.Name,
		Resolution:     sensorIn.Resolution,
		Width:          sensorIn.Width,
		RightSideAngle: sensorIn.RightSideAngle,
		LeftSideAngle:  sensorIn.LeftSideAngle,
		ObserveAngle:   sensorIn.ObserveAngle,
		InitAngle:      sensorIn.InitAngle,
		OleColor:       sensorIn.OleColor,
	}
	err = db.CreateSensor(&sensor)
	return sensor.Id, err
}

func GetAllSensors() *[]entity.Sensor {
	return db.FindAllSensors()
}

func GetSensorBySatId(satId string) (*[]entity.Sensor, error) {
	return db.FindSensorBySatId(satId)
}

func GetSensorBySatIdAndName(satId string, senName string) (*entity.Sensor, error) {
	return db.FindSensorBySatIdAndName(satId, senName)
}

func GetSensorById(id int) *entity.Sensor {
	sen, _ := db.FindSensorById(id)
	return sen
}

func UpdateSensor(senId int, senDTO *entity.SensorInDTO) error {
	senInDB, err := db.FindSensorById(senId)
	if err != nil || senInDB.Id == 0 {
		return errors.New("error finding sensor")
	}
	satInDB, err := db.FindSatelliteByNoardId(senDTO.SatId)
	if err != nil || satInDB.Id == 0 {
		return errors.New("error finding satellite")
	}
	senInDB.SatNoardId = senDTO.SatId
	senInDB.SatName = satInDB.Name
	senInDB.Name = senDTO.Name
	senInDB.Resolution = senDTO.Resolution
	senInDB.Width = senDTO.Width
	senInDB.RightSideAngle = senDTO.RightSideAngle
	senInDB.LeftSideAngle = senDTO.LeftSideAngle
	senInDB.ObserveAngle = senDTO.ObserveAngle
	senInDB.InitAngle = senDTO.InitAngle
	senInDB.OleColor = senDTO.OleColor

	return db.SaveSensor(senInDB)
}

func DeleteSensorBySatIdAndName(satId string, senName string) error {
	return db.DeleteSensorBySatIdAndName(satId, senName)
}

func DeleteSensorById(id int) error {
	return db.DeleteSensorById(id)
}
