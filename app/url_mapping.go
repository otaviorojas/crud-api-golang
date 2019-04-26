package app

import "github.com/crud_golang/controllers"

func mapUrlsToControllers() {

	//*** PING AND INDEX----------------------------------------------------------- ***//

	router.GET("/", controllers.GetIndex)

	router.GET("/ping", controllers.GetPing)

	//*** STUDENT GET * POST * DELETE * PUT --------------------------------------- ***//

	router.GET("/students", controllers.GetAllStudents)

	router.GET("/student/:id", controllers.GetStudent)

	router.POST("/student", controllers.Save)

	router.DELETE("/student/:id", controllers.Delete)

	router.PUT("/student/:id", controllers.Update)

}
