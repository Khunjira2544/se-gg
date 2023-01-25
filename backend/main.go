package main

import (
	"github.com/Khunjira2544/se-65-g08/controller"
	"github.com/Khunjira2544/se-65-g08/entity"
	//"github.com/Khunjira2544/se-65-g08/middlewares"
	"github.com/gin-gonic/gin"
)

const PORT = "8080"

func main() {
	entity.SetupDatabase()

	 r := gin.Default()
	 r.Use(CORSMiddleware())

	//router := r.Group("/")
	{
		//router.Use(middlewares.Authorizes())
		{
			// User Routes   //ของเอา Officer officer
			r.GET("/doctors", controller.ListDoctors)
			r.GET("/doctor/:id", controller.GetDoctor)
			r.PATCH("/doctors", controller.UpdateDoctor)
			r.DELETE("/doctors/:id", controller.DeleteDoctor)

			// Resolution Routes   //ของเราคือ collegeyear  Collegeyear
			r.GET("/diseases", controller.ListDiseases)
			r.GET("/diseases/:id", controller.GetDisease)
			r.POST("/disease", controller.CreateDisease)
			r.PATCH("/disease", controller.UpdateDisease)
			r.DELETE("/disease/:id", controller.DeleteDisease)

			// ของเรา faculty Faculty
			r.GET("/statuses", controller.ListStatuses)
			r.GET("/status/:id", controller.GetStatus)
			r.POST("/statuses", controller.CreateStatus)
			r.PATCH("/statuses", controller.UpdateStatus)
			r.DELETE("/status/:id", controller.DeleteStatus)

			// Video Routes  //ของเรา teacher  Teacher
			r.GET("/patients", controller.ListPatients)
			//r.GET("/faculty/:faculty_id", controller.ListT_Facultys)
			r.GET("/patient/:id", controller.GetPatient)
			r.POST("/patients", controller.CreatePatient)
			r.PATCH("/patients", controller.UpdatePatient)
			r.DELETE("/patients/:id", controller.DeletePatient)

			// ของเรา faculty Faculty
			r.GET("/tracks", controller.ListTracks) ////
			r.GET("/tracks/:id", controller.GetTrack)
			r.POST("/tracks", controller.CreateTrack)
			r.PATCH("/tracks", controller.UpdateTrack)
			r.DELETE("/track/:id", controller.DeleteTrack)

			// WatchVideo Routes   //ของเรา Student student
			r.GET("/treatments", controller.ListTreatment)
			r.GET("/treatments/:id", controller.GetTreatment)
			r.POST("/treatments", controller.CreateTreatment)
			r.PATCH("/treatments", controller.UpdateTreatment)
			r.DELETE("/treatment/:id", controller.DeleteTreatment)

			//----ระบบการเบิกอุปกรณ์แพทย์
			r.GET("/equipments", controller.ListEquipments) ////
			r.GET("/equipments/:id", controller.GetEquipment)
			r.POST("/equipments", controller.CreateEquipment)
			r.PATCH("/equipments", controller.UpdateEquipment)
			r.DELETE("/equipment/:id", controller.DeleteEquipment)

			r.GET("/locations", controller.ListLocations) ////
			r.GET("/locations/:id", controller.GetLocation)
			r.POST("/locations", controller.CreateLocation)
			r.PATCH("/locations", controller.UpdateLocation)
			r.DELETE("/location/:id", controller.DeleteLocation)

			r.GET("/medicines", controller.ListMedicines)
			r.GET("/medicine/:id", controller.GetMedicine)
			r.PATCH("/medicines", controller.UpdateMedicine)
			r.DELETE("/medicines/:id", controller.DeleteMedicine)

			r.GET("/requests", controller.ListRequest)
			r.GET("/requests/:id", controller.GetRequest)
			r.POST("/requests", controller.CreateRequest)
			r.PATCH("/requests", controller.UpdateRequest)
			r.DELETE("/request/:id", controller.DeleteRequest)

		}
	}

	// Signup User Route
	//r.POST("/signup", controller.CreateDoctor)
	// login User Route
	//r.POST("/login", controller.Login)

	// Run the server go run main.go
	r.Run()
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
