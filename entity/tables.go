package entity

type Satellite struct {
	Id        int
	Name      string
	NoardId   string
	OleColor  int
	IsChecked bool
}

type Sensor struct {
	Id             int
	SatNoardId     string
	SatName        string
	Name           string
	Resolution     float32
	Width          float32
	RightSideAngle float32
	LeftSideAngle  float32
	ObserveAngle   float32
	InitAngle      float32
	OleColor       int
	IsChecked      bool
}

type SysUser struct {
	Id       int    `gorm:"primary_key"`
	UserName string `gorm:"type:varchar(255)"`
	Password string `gorm:"type:varchar(255)"`
	Email    string `gorm:"type:varchar(255)"`
}

type Tle struct {
	Id         int
	SatNoardId string
	Time       int64
	Line1      string
	Line2      string
}

type TleSite struct {
	Id          int
	Site        string
	Url         string
	Description string
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
