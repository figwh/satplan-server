package syscfg

import (
	c "satplan/controller"

	log "github.com/sirupsen/logrus"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

var swagHandler gin.HandlerFunc

func NewGinRouterWithAuth(authMiddleware *jwt.GinJWTMiddleware) *gin.Engine {

	//API_PREFIX := "/api"
	API_PREFIX := ""
	router := gin.Default()
	router.POST(API_PREFIX+"/login", authMiddleware.LoginHandler)
	router.GET(API_PREFIX+"/", c.HelloGin)
	router.GET(API_PREFIX+"/version", c.Version)
	router.GET(API_PREFIX+"/test", c.TempTest)

	router.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Debugf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	//user
	user := router.Group(API_PREFIX + "/user")
	user.Use(authMiddleware.MiddlewareFunc())
	{
		user.GET("/all", c.GetAllUsers)
		user.GET("/me", c.GetUserInfo)
	}
	//sat
	sat := router.Group(API_PREFIX + "/sat")
	sat.Use(authMiddleware.MiddlewareFunc())
	{
		sat.GET("/all", c.GetAllSatellites)
		sat.POST("/add", c.AddSatellite)
		sat.GET("/:id", c.GetSatelliteById)
		sat.PUT("/update/:id", c.UpdateSatellite)
		sat.DELETE("/delete/:id", c.DeleteSatellite)
		sat.POST("/tle/update", c.UpdateTles)
	}

	//sen
	sen := router.Group(API_PREFIX + "/sen")
	sen.Use(authMiddleware.MiddlewareFunc())
	{
		sen.GET("/all", c.GetAllSensors)
		sen.POST("/add", c.AddSensor)
		sen.GET("/:satid", c.GetSensorBySatId)
		sen.GET("/:id", c.GetSensorById)
		sen.PUT("/update/:id", c.UpdateSensor)
		sen.DELETE("/delete/:id", c.DeleteSensor)
	}

	//track
	track := router.Group(API_PREFIX + "/track")
	track.Use(authMiddleware.MiddlewareFunc())
	{
		track.GET("/", c.GetTrackBySatId)
	}

	//path
	path := router.Group(API_PREFIX + "/path")
	path.Use(authMiddleware.MiddlewareFunc())
	{
		path.GET("/:senid", c.GetPathBySenId)
		path.POST("/satplan", c.GetPathPlan)
	}
	return router
}
