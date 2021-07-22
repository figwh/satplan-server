package db

import "satplan/entity"

func FindSysUserByEmail(email string) *entity.SysUser {

	return nil
}

func FindSysUserByUserId(userId int) *entity.SysUser {
	return nil
}

func DeleteSensors() {
	satDb.Delete(entity.Sensor{}, "1=1")
}

func CreateSensor(sen *entity.Sensor) error {
	query := satDb.Create(sen)
	return query.Error
}

func BatCreateSensor(sens *[]entity.Sensor) error {
	query := satDb.CreateInBatches(sens, len(*sens))
	return query.Error
}

func DeleteSatellites() {
	satDb.Delete(entity.Satellite{}, "1=1")
}

func FindSatelliteByNoardId(noardId string) *entity.Satellite {
	sat := entity.Satellite{}
	satDb.Where("noard_id= ? ", noardId).First(&sat)
	return &sat
}

func CreateSatellite(sat *entity.Satellite) error {
	query := satDb.Create(sat)
	return query.Error
}

func BatCreateSatellite(sats *[]entity.Satellite) error {
	query := satDb.CreateInBatches(sats, len(*sats))
	return query.Error
}

func DeleteTles() {
	satDb.Delete(entity.Tle{}, "1=1")
}

func CreateTle(tle *entity.Tle) error {
	query := satDb.Create(tle)
	return query.Error
}

func BatCreateTle(tles *[]entity.Tle) error {
	query := satDb.CreateInBatches(tles, len(*tles))
	return query.Error
}
