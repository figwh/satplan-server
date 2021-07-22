package entity

type Satellite struct {
	Id        int
	Name      string
	OleColor  int
	IsChecked int
}

type Sensor struct {
	SatId          int
	SatName        string
	Id             int
	Name           string
	Resolution     float32
	Width          float32
	RightSideAngle float32
	LeftSideAngle  float32
	ObserveAngle   float32
	InitAngle      float32
	OleColor       int
	IsChecked      int
}

type SysUser struct {
	Id       int    `gorm:"primary_key"`
	Email    string `gorm:"type:varchar(255)"`
	Name     string `gorm:"type:varchar(255)"`
	AdminId  int
	RoleId   int
	Password string `gorm:"type:varchar(255)"`
}

type Track struct {
	Time int64
	X    float32
	Y    float32
	Z    float32
	Vx   float32
	Vy   float32
	Vz   float32
	Lon  float32
	Lat  float32
	Alt  float32
}

type Path struct {
	Time int64
	Lon1 float32
	Lat1 float32
	Lon2 float32
	Lat2 float32
}
