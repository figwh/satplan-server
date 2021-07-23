package entity

type User struct {
	Id       int
	UserName string
	Email    string
}

type PlanPara struct {
	SensorId string
	Start    int64
	Stop     int64
	Xmin     float32
	Xmax     float32
	Ymin     float32
	Ymax     float32
}

type TleData struct {
	Line0 string
	Line1 string
	Line2 string
}

type SatDTO struct {
	SatName  string
	OleColor int
}

type SensorInDTO struct {
	SatId          string
	Name           string
	Resolution     float32
	Width          float32
	RightSideAngle float32
	LeftSideAngle  float32
	ObserveAngle   float32
	InitAngle      float32
	OleColor       int
}

type CurrentUserInfo struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	AdminId int    `json:"adminId"`
	RoleId  int    `json:"roleId"`
}

type PathUnit struct {
	SatName string
	SenName string
	Start   int64
	Stop    int64
	PathGeo *Path
}
