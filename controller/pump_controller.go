package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rmukubvu/pumpdata/model"
	"github.com/rmukubvu/pumpdata/rabbit"
	"github.com/rmukubvu/pumpdata/repository"
	"net/http"
	"strconv"
	"time"
)

type internalError struct {
	Message string `json:"message"`
}

var rb *rabbit.QueueService

func InitRouter(q *rabbit.QueueService) *gin.Engine {
	r := gin.Default()
	rb = q
	gin.SetMode(gin.ReleaseMode)
	v1 := r.Group("/api/v1/")
	{
		v1.POST("pump", createPump)
		v1.POST("pump/types", createPumpType)
		v1.POST("service", createPumpService)
		v1.POST("service/history", createPumpServiceHistory)
		v1.POST("company", createCompany)
		v1.POST("sensor", createSensor)
		v1.POST("sensor-alarm", createSensorAlarm)
		v1.POST("sensor-contact", createSensorAlarmContact)
		v1.POST("slack", createSlackMessage)
		v1.POST("sms", createSmsAlerts)
		v1.POST("annunciator", createAnnunciator)
		v1.POST("pump-test-by-serial", createPumpTest)
		v1.POST("remote-calls", remoteUnitCall)

		v1.GET("annunciator", getAnnunciators)
		v1.GET("annunciator-by-serial/:serial", getAnnunciatorBySerial)
		v1.GET("sensor/:tid/:pid", getSensorByTypeAndId)
		v1.GET("sensor-data/:serial", getSensorDataBySerial)
		v1.GET("sensor-alarm", getAlarms)
		v1.GET("sensor-contact/:id", getContactsByCompany)
		v1.GET("company/:id", getCompanyById)
		v1.GET("pump/types", getPumpTypes)
		v1.GET("pump", getPumps)
		v1.GET("pump/by-serial/:serial", getPumpBySerial)
		v1.GET("pump/sensor-types", getPumpSensorTypes)
		v1.GET("pump/by-company/:id", getPumpsUnderCompanyById)
		v1.GET("pump/dashboard", getDashboardInformation)
		v1.GET("water-tank/:serial", getWaterTankLevel)
		v1.GET("service", getPumpServices)
		v1.GET("service/by-serial/:serial", getServiceBySerial)
		v1.GET("service/history/:serial", getServiceHistory)
		v1.GET("sensor-by-type-serial/:id/:serial", getSensorByTypeAndSerial)
		v1.GET("pump-test-by-serial/:serial", getPumpTestBySerial)
		v1.GET("orange-notification/:date", getOrangeNotifications)
		v1.GET("red-notification/:date", getRedNotifications)
		//reports
		v1.POST("sensor/report/by-company", getSensorReportByCompany)
		v1.POST("alarm/sensor/report/by-serial", getAlarmSensorReportBySerial)
		v1.POST("service/report/by-serial", getServiceReportBySerial)
	}
	return r
}

func createPump(c *gin.Context) {
	p := model.Pump{}
	err := c.Bind(&p)
	if err != nil {
		c.JSON(http.StatusInternalServerError, generateErrorMessage(err.Error()))
		return
	}
	err = repository.AddPump(p)
	if err != nil {
		c.JSON(http.StatusInternalServerError, generateErrorMessage(err.Error()))
		return
	}
	c.JSON(http.StatusCreated, p)
}

func createPumpService(c *gin.Context) {
	p := model.PumpService{}
	err := c.Bind(&p)
	if err != nil {
		c.JSON(http.StatusInternalServerError, generateErrorMessage(err.Error()))
		return
	}
	err = repository.AddPumpService(p)
	if err != nil {
		c.JSON(http.StatusInternalServerError, generateErrorMessage(err.Error()))
		return
	}
	c.JSON(http.StatusCreated, p)
}

func createPumpTest(c *gin.Context) {
	p := model.PumpTest{}
	err := c.Bind(&p)
	if err != nil {
		c.JSON(http.StatusInternalServerError, generateErrorMessage(err.Error()))
		return
	}
	err = repository.AddPumpTests(p)
	if err != nil {
		c.JSON(http.StatusInternalServerError, generateErrorMessage(err.Error()))
		return
	}
	c.JSON(http.StatusCreated, p)
}

func remoteUnitCall(c *gin.Context) {
	p := model.Remote{}
	err := c.Bind(&p)
	if err != nil {
		c.JSON(http.StatusInternalServerError, generateErrorMessage(err.Error()))
		return
	}
	res, err := rb.RemoteCalls(p)
	if err != nil {
		c.JSON(http.StatusInternalServerError, generateErrorMessage(err.Error()))
		return
	}
	c.JSON(http.StatusCreated, res)
}

func createPumpServiceHistory(c *gin.Context) {
	p := model.PumpServiceHistory{}
	err := c.Bind(&p)
	if err != nil {
		c.JSON(http.StatusInternalServerError, generateErrorMessage(err.Error()))
		return
	}
	err = repository.AddPumpServiceHistory(p)
	if err != nil {
		c.JSON(http.StatusInternalServerError, generateErrorMessage(err.Error()))
		return
	}
	c.JSON(http.StatusCreated, p)
}

func createSlackMessage(c *gin.Context) {
	p := model.DigitalMessage{}
	err := c.Bind(&p)
	if err != nil {
		c.JSON(http.StatusInternalServerError, generateErrorMessage(err.Error()))
		return
	}
	err = repository.SendSlack(p.Message)
	if err != nil {
		c.JSON(http.StatusInternalServerError, generateErrorMessage(err.Error()))
		return
	}
	c.JSON(http.StatusCreated, p)
}

func createCompany(c *gin.Context) {
	p := model.Company{}
	err := c.Bind(&p)
	if err != nil {
		c.JSON(http.StatusInternalServerError, generateErrorMessage(err.Error()))
		return
	}
	err = repository.AddCompany(p)
	if err != nil {
		c.JSON(http.StatusInternalServerError, generateErrorMessage(err.Error()))
		return
	}
	c.JSON(http.StatusCreated, p)
}

func createSensor(c *gin.Context) {
	p := model.Sensor{}
	err := c.Bind(&p)
	if err != nil {
		c.JSON(http.StatusInternalServerError, generateErrorMessage(err.Error()))
		return
	}
	err = repository.AddSensor(p)
	if err != nil {
		c.JSON(http.StatusInternalServerError, generateErrorMessage(err.Error()))
		return
	}
	c.JSON(http.StatusCreated, p)
}

func createSensorAlarm(c *gin.Context) {
	p := model.SensorAlarm{}
	err := c.Bind(&p)
	if err != nil {
		c.JSON(http.StatusInternalServerError, generateErrorMessage(err.Error()))
		return
	}
	err = repository.AddSensorAlarm(p)
	if err != nil {
		c.JSON(http.StatusInternalServerError, generateErrorMessage(err.Error()))
		return
	}
	c.JSON(http.StatusCreated, p)
}

func createAnnunciator(c *gin.Context) {
	p := model.Annunciator{}
	err := c.Bind(&p)
	if err != nil {
		c.JSON(http.StatusInternalServerError, generateErrorMessage(err.Error()))
		return
	}
	err = repository.AddAnnunciator(p)
	if err != nil {
		c.JSON(http.StatusInternalServerError, generateErrorMessage(err.Error()))
		return
	}
	c.JSON(http.StatusCreated, p)
}

func createSmsAlerts(c *gin.Context) {
	var p []model.DigitalMessage
	err := c.Bind(&p)
	if err != nil {
		c.JSON(http.StatusInternalServerError, generateErrorMessage(err.Error()))
		return
	}
	err = repository.SendSms(p)
	if err != nil {
		c.JSON(http.StatusInternalServerError, generateErrorMessage(err.Error()))
		return
	}
	c.JSON(http.StatusCreated, p)
}

func createSensorAlarmContact(c *gin.Context) {
	p := model.SensorAlarmContact{}
	err := c.Bind(&p)
	if err != nil {
		c.JSON(http.StatusInternalServerError, generateErrorMessage(err.Error()))
		return
	}
	err = repository.AddSensorAlarmContact(p)
	if err != nil {
		c.JSON(http.StatusInternalServerError, generateErrorMessage(err.Error()))
		return
	}
	c.JSON(http.StatusCreated, p)
}

func createPumpType(c *gin.Context) {
	p := model.PumpTypes{}
	err := c.Bind(&p)
	if err != nil {
		c.JSON(http.StatusInternalServerError, generateErrorMessage(err.Error()))
		return
	}
	err = repository.AddPumpType(p)
	if err != nil {
		c.JSON(http.StatusInternalServerError, generateErrorMessage(err.Error()))
		return
	}
	c.JSON(http.StatusCreated, p)
}

func getCompanyById(c *gin.Context) {
	id := c.Param("id")
	if len(id) == 0 {
		c.JSON(http.StatusInternalServerError, generateErrorMessage("company id is required"))
		return
	}
	i, _ := strconv.Atoi(id)
	p, err := repository.GetCompanyById(i)
	if err != nil {
		c.JSON(http.StatusInternalServerError, generateErrorMessage(err.Error()))
		return
	}
	c.JSON(http.StatusOK, p)
}

func getPumpsUnderCompanyById(c *gin.Context) {
	id := c.Param("id")
	if len(id) == 0 {
		c.JSON(http.StatusInternalServerError, generateErrorMessage("company id is required"))
		return
	}
	i, _ := strconv.Atoi(id)
	p, err := repository.GetPumpsUnderCompany(i)
	if err != nil {
		c.JSON(http.StatusInternalServerError, generateErrorMessage(err.Error()))
		return
	}
	c.JSON(http.StatusOK, p)
}

func getSensorDataBySerial(c *gin.Context) {
	serial := c.Param("serial")
	if len(serial) == 0 {
		c.JSON(http.StatusInternalServerError, generateErrorMessage("serial number is required"))
		return
	}
	p, err := repository.GetSensorDataBySerial(serial)
	if err != nil {
		c.JSON(http.StatusInternalServerError, generateErrorMessage(err.Error()))
		return
	}
	c.JSON(http.StatusOK, p)
}

func getContactsByCompany(c *gin.Context) {
	id := c.Param("id")
	if len(id) == 0 {
		c.JSON(http.StatusInternalServerError, generateErrorMessage("company id is required"))
		return
	}
	i, _ := strconv.Atoi(id)
	p := repository.GetAlarmContacts(i)
	c.JSON(http.StatusOK, p)
}

func getSensorByTypeAndId(c *gin.Context) {
	typeId := c.Param("tid")
	pumpId := c.Param("pid")
	if len(typeId) == 0 && len(pumpId) == 0 {
		c.JSON(http.StatusInternalServerError, generateErrorMessage("type id & pump id is required"))
		return
	}
	tid, _ := strconv.Atoi(typeId)
	pid, _ := strconv.Atoi(pumpId)
	p, err := repository.SensorByTypeAndId(tid, pid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, generateErrorMessage(err.Error()))
		return
	}
	c.JSON(http.StatusOK, p)
}

func getSensorByTypeAndSerial(c *gin.Context) {
	typeId := c.Param("id")
	serial := c.Param("serial")
	if len(typeId) == 0 && len(serial) == 0 {
		c.JSON(http.StatusInternalServerError, generateErrorMessage("type id & serial is required"))
		return
	}
	tid, _ := strconv.Atoi(typeId)
	p, err := repository.SensorViewModelBySerialAndType(tid, serial)
	if err != nil {
		c.JSON(http.StatusInternalServerError, generateErrorMessage(err.Error()))
		return
	}
	c.JSON(http.StatusOK, p)
}

func getPumpBySerial(c *gin.Context) {
	serialNumber := c.Param("serial")
	if len(serialNumber) == 0 {
		c.JSON(http.StatusInternalServerError, generateErrorMessage("serial number is required"))
		return
	}
	p, err := repository.GetPumpBySerialNumber(serialNumber, false)
	if err != nil {
		c.JSON(http.StatusInternalServerError, generateErrorMessage(err.Error()))
		return
	}
	c.JSON(http.StatusOK, p)
}

func getWaterTankLevel(c *gin.Context) {
	serialNumber := c.Param("serial")
	if len(serialNumber) == 0 {
		c.JSON(http.StatusInternalServerError, generateErrorMessage("serial number is required"))
		return
	}
	p, err := repository.GetWaterTankLevelForSerial(serialNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, generateErrorMessage(err.Error()))
		return
	}
	c.JSON(http.StatusOK, p)
}

func getServiceHistory(c *gin.Context) {
	serialNumber := c.Param("serial")
	if len(serialNumber) == 0 {
		c.JSON(http.StatusInternalServerError, generateErrorMessage("serial number is required"))
		return
	}
	p, err := repository.GetServiceHistoryForPump(serialNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, generateErrorMessage(err.Error()))
		return
	}
	c.JSON(http.StatusOK, p)
}

func getServiceBySerial(c *gin.Context) {
	serialNumber := c.Param("serial")
	if len(serialNumber) == 0 {
		c.JSON(http.StatusInternalServerError, generateErrorMessage("serial number is required"))
		return
	}
	p, err := repository.PumpService(serialNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, generateErrorMessage(err.Error()))
		return
	}
	c.JSON(http.StatusOK, p)
}

func getPumpTestBySerial(c *gin.Context) {
	serialNumber := c.Param("serial")
	if len(serialNumber) == 0 {
		c.JSON(http.StatusInternalServerError, generateErrorMessage("serial number is required"))
		return
	}
	p, err := repository.PumpTestBySerial(serialNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, generateErrorMessage(err.Error()))
		return
	}
	c.JSON(http.StatusOK, p)
}

func getOrangeNotifications(c *gin.Context) {
	date := c.Param("date")
	if len(date) == 0 {
		c.JSON(http.StatusInternalServerError, generateErrorMessage("serial number is required"))
		return
	}

	d, err := time.Parse("2006-01-02", date)
	if err != nil {
		c.JSON(http.StatusInternalServerError, generateErrorMessage(err.Error()))
		return
	}
	p, err := repository.GetOrangeNotification(d)
	if err != nil {
		c.JSON(http.StatusInternalServerError, generateErrorMessage(err.Error()))
		return
	}
	c.JSON(http.StatusOK, p)
}

func getRedNotifications(c *gin.Context) {
	date := c.Param("date")
	if len(date) == 0 {
		c.JSON(http.StatusInternalServerError, generateErrorMessage("serial number is required"))
		return
	}

	d, err := time.Parse("2006-01-02", date)
	if err != nil {
		c.JSON(http.StatusInternalServerError, generateErrorMessage(err.Error()))
		return
	}
	p, err := repository.GetRedNotification(d)
	if err != nil {
		c.JSON(http.StatusInternalServerError, generateErrorMessage(err.Error()))
		return
	}
	c.JSON(http.StatusOK, p)
}

func getPumps(c *gin.Context) {
	res, err := repository.GetAllPumps()
	if err != nil {
		c.JSON(http.StatusInternalServerError, generateErrorMessage(err.Error()))
		return
	}
	c.JSON(http.StatusOK, res)
}

func getPumpTypes(c *gin.Context) {
	res, err := repository.FetchPumpTypes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, generateErrorMessage(err.Error()))
		return
	}
	c.JSON(http.StatusOK, res)
}

func getDashboardInformation(c *gin.Context) {
	res, err := repository.DashboardAlarms()
	if err != nil {
		c.JSON(http.StatusInternalServerError, generateErrorMessage(err.Error()))
		return
	}
	c.JSON(http.StatusOK, res)
}

func getAlarms(c *gin.Context) {
	res, err := repository.GetAllSensorAlarms()
	if err != nil {
		c.JSON(http.StatusInternalServerError, generateErrorMessage(err.Error()))
		return
	}
	c.JSON(http.StatusOK, res)
}

func getPumpSensorTypes(c *gin.Context) {
	res, err := repository.FetchSensorTypes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, generateErrorMessage(err.Error()))
		return
	}
	c.JSON(http.StatusOK, res)
}

func getAnnunciators(c *gin.Context) {
	res, err := repository.GetAnnunciators()
	if err != nil {
		c.JSON(http.StatusInternalServerError, generateErrorMessage(err.Error()))
		return
	}
	c.JSON(http.StatusOK, res)
}

func getAnnunciatorBySerial(c *gin.Context) {
	serialNumber := c.Param("serial")
	if len(serialNumber) == 0 {
		c.JSON(http.StatusInternalServerError, generateErrorMessage("serial number is required"))
		return
	}
	p, err := repository.GetAnnunciatorBySerial(serialNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, generateErrorMessage(err.Error()))
		return
	}
	c.JSON(http.StatusOK, p)
}

func getPumpServices(c *gin.Context) {
	res, err := repository.GetPumpsServices()
	if err != nil {
		c.JSON(http.StatusInternalServerError, generateErrorMessage(err.Error()))
		return
	}
	c.JSON(http.StatusOK, res)
}

/******************************* reports ***********************************************/

func getSensorReportByCompany(c *gin.Context) {
	p := model.SensorDataReportRequest{}
	err := c.Bind(&p)
	if err != nil {
		c.JSON(http.StatusInternalServerError, generateErrorMessage("invalid json string"))
		return
	}
	res, err := repository.GenerateSensorDataReport(p)
	if err != nil {
		c.JSON(http.StatusInternalServerError, generateErrorMessage(err.Error()))
		return
	}
	c.JSON(http.StatusOK, res)
}

func getServiceReportBySerial(c *gin.Context) {
	p := model.ServiceReportRequest{}
	err := c.Bind(&p)
	if err != nil {
		c.JSON(http.StatusInternalServerError, generateErrorMessage("invalid json string"))
		return
	}
	res, err := repository.GenerateServiceReport(p)
	if err != nil {
		c.JSON(http.StatusInternalServerError, generateErrorMessage(err.Error()))
		return
	}
	c.JSON(http.StatusOK, res)
}

func getAlarmSensorReportBySerial(c *gin.Context) {
	p := model.AlarmsDataReportRequest{}
	err := c.Bind(&p)
	if err != nil {
		c.JSON(http.StatusInternalServerError, generateErrorMessage("invalid json string"))
		return
	}
	res, err := repository.GenerateAlarmSensorReport(p)
	if err != nil {
		c.JSON(http.StatusInternalServerError, generateErrorMessage(err.Error()))
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
