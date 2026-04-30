package main

import (
	"telemedicine-api/route"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	route.RegisterDoctorRoutes(r)

	r.Run(":8080")
}
