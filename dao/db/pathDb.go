package db

import (
	"satplan/entity"
)

//=================info===============
func FindSenPathInfo(satId string, senName string) *entity.PathInfo {
	pathDb := GetSenPathDb(satId, senName)
	if pathDb == nil {
		return nil
	}

	pathInfo := entity.PathInfo{}
	pathDb.First(&pathInfo)
	return &pathInfo
}

//=================senPath===============
func FindLastSenPathPoint(satId string, senName string) *entity.SenPath {
	senPathDb := GetSenPathDb(satId, senName)
	if senPathDb == nil {
		return nil
	}

	senPath := entity.SenPath{}
	senPathDb.Order("time_offset desc").First(&senPath)
	return &senPath
}

func FindSenPath(satId string, senName string, start int64, stop int64) *[]entity.SenPath {
	senPathDb := GetSenPathDb(satId, senName)
	if senPathDb == nil {
		return nil
	}

	senPath := []entity.SenPath{}
	senPathDb.Where("time_offset >=? and time_offset <= ?", start, stop).Find(&senPath)
	return &senPath
}

func FindPathUnit(satId string, senName string, start int64, stop int64,
	xmin float32, xmax float32, ymax float32, ymin float32) *[]entity.PathUnit {

	senPathDb := GetSenPathDb(satId, senName)
	if senPathDb == nil {
		return nil
	}

	senPath := []entity.SenPath{}
	//SELECT time, lon1,lat1,lon2,lat2 FROM path where time >=$delta_t1
	//and time <=$delta_t2 and ((lon1 >$xmin and lon1 <$xmax and lat1 >$ymin and lat1 <$ymax)
	//or (lon2 >$xmin and lon2 <$xmax and lat2 >$ymin and lat2 <$ymax)) order by time"

	senPathDb.Where("time_offset >=? and time_offset <= ? and ((lon1 >? and lon1 <? "+
		" and lat1>? and lat1<?) or (lon2>? and lon2<? and lat2>? and lat2<?))",
		start, stop, xmin, xmax, ymin, ymax, xmin, xmax, ymin, ymax).Order("time_offset").Find(&senPath)
	//分隔成一个个小区域
	pathInfo := FindSenPathInfo(satId, senName)
	start_index := 0
	pathUnits := []entity.PathUnit{{
		SatName: pathInfo.SatName,
		SenName: pathInfo.SenName,
		Start:   senPath[0].TimeOffset,
	}}
	for i := 0; i < len(senPath)-1; i++ {
		curPoint := senPath[i]
		nextPath := senPath[i+1]

		if nextPath.TimeOffset-curPoint.TimeOffset > int64(100) {
			//last point
			pathUnits[len(pathUnits)-1].Stop = curPoint.TimeOffset
			pathGeo := senPath[start_index:i]
			pathUnits[len(pathUnits)-1].PathGeo = &pathGeo
			//new point
			pathUnits = append(pathUnits, entity.PathUnit{
				SatName: pathInfo.SatName,
				SenName: pathInfo.SenName,
				Start:   nextPath.TimeOffset,
				Stop:    0,
			})
			start_index = i + 1
		}
	}
	return &pathUnits
}