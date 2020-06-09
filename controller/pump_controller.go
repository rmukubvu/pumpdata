package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rmukubvu/pumpdata/model"
	"github.com/rmukubvu/pumpdata/repository"
	"net/http"
	"strconv"
)

type internalError struct {
	Message string `json:"message"`
}

func InitRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/api/v1/")
	{
		v1.POST("pump", createPump)
		v1.POST("pump/types", createPumpType)
		v1.POST("company", createCompany)
		v1.POST("sensor", createSensor)
		v1.GET("sensor/:tid/:pid", getSensorByTypeAndId)
		v1.GET("company/:id", getCompanyById)
		v1.GET("pump/types", getPumpTypes)
		v1.GET("pump/by-serial/:serial", getPumpBySerial)
		v1.GET("pump/sensor-types", getPumpSensorTypes)
	}
	return r
}

func createPump(c *gin.Context) {
	p := model.Pump{}
	err := c.Bind(&p)
	if err != nil {
		c.JSON(http.StatusBadGateway, generateErrorMessage("invalid json string"))
		return
	}
	err = repository.AddPump(p)
	if err != nil {
		c.JSON(http.StatusBadGateway, generateErrorMessage(err.Error()))
		return
	}
	c.JSON(http.StatusCreated, p)
}

func createCompany(c *gin.Context) {
	p := model.Company{}
	err := c.Bind(&p)
	if err != nil {
		c.JSON(http.StatusBadGateway, generateErrorMessage("invalid json string"))
		return
	}
	err = repository.AddCompany(p)
	if err != nil {
		c.JSON(http.StatusBadGateway, generateErrorMessage(err.Error()))
		return
	}
	c.JSON(http.StatusCreated, p)
}

func createSensor(c *gin.Context) {
	p := model.Sensor{}
	err := c.Bind(&p)
	if err != nil {
		c.JSON(http.StatusBadGateway, generateErrorMessage("invalid json string"))
		return
	}
	err = repository.AddSensor(p)
	if err != nil {
		c.JSON(http.StatusBadGateway, generateErrorMessage(err.Error()))
		return
	}
	c.JSON(http.StatusCreated, p)
}

func createPumpType(c *gin.Context) {
	p := model.PumpTypes{}
	err := c.Bind(&p)
	if err != nil {
		c.JSON(http.StatusBadGateway, generateErrorMessage("invalid json string"))
		return
	}
	err = repository.AddPumpType(p)
	if err != nil {
		c.JSON(http.StatusBadGateway, generateErrorMessage(err.Error()))
		return
	}
	c.JSON(http.StatusCreated, p)
}

func getCompanyById(c *gin.Context) {
	id := c.Param("id")
	if len(id) == 0 {
		c.JSON(http.StatusBadGateway, generateErrorMessage("company id is required"))
		return
	}
	i, _ := strconv.Atoi(id)
	p, err := repository.GetCompanyById(i)
	if err != nil {
		c.JSON(http.StatusBadGateway, generateErrorMessage(err.Error()))
		return
	}
	c.JSON(http.StatusOK, p)
}

func getSensorByTypeAndId(c *gin.Context) {
	typeId := c.Param("tid")
	pumpId := c.Param("pid")
	if len(typeId) == 0 && len(pumpId) == 0 {
		c.JSON(http.StatusBadGateway, generateErrorMessage("type id & pump id is required"))
		return
	}
	tid, _ := strconv.Atoi(typeId)
	pid, _ := strconv.Atoi(pumpId)
	p, err := repository.SensorByTypeAndId(tid, pid)
	if err != nil {
		c.JSON(http.StatusBadGateway, generateErrorMessage(err.Error()))
		return
	}
	c.JSON(http.StatusOK, p)
}

func getPumpBySerial(c *gin.Context) {
	serialNumber := c.Param("serial")
	if len(serialNumber) == 0 {
		c.JSON(http.StatusBadGateway, generateErrorMessage("serial number is required"))
		return
	}
	p, err := repository.GetPumpBySerialNumber(serialNumber)
	if err != nil {
		c.JSON(http.StatusBadGateway, generateErrorMessage(err.Error()))
		return
	}
	c.JSON(http.StatusOK, p)
}

func getPumpTypes(c *gin.Context) {
	res, err := repository.FetchPumpTypes()
	if err != nil {
		c.JSON(http.StatusBadGateway, generateErrorMessage(err.Error()))
		return
	}
	c.JSON(http.StatusOK, res)
}

func getPumpSensorTypes(c *gin.Context) {
	res, err := repository.FetchSensorTypes()
	if err != nil {
		c.JSON(http.StatusBadGateway, generateErrorMessage(err.Error()))
		return
	}
	c.JSON(http.StatusOK, res)
}

func generateErrorMessage(e string) string {
	ie := internalError{Message: e}
	buf, err := json.Marshal(ie)
	if err != nil {
		return fmt.Sprintf(`{"message": "%s"}`, e)
	}
	return string(buf)
}
