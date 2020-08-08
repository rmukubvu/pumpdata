package repository

import (
	"github.com/rmukubvu/pumpdata/model"
	"github.com/rmukubvu/pumpdata/store"
	"reflect"
)

func GenerateAlarmSensorReport(p model.AlarmsDataReportRequest) (model.ReportResponse, error) {
	if err := validateSerialNumber(p.SerialNumber); err != nil {
		return model.ReportResponse{
			Error: err.Error(),
		}, err
	}
	result, err := store.GenerateAlarmSensorReport(p.SerialNumber, p.StartDate, p.EndDate)
	if ok, val := takeSliceArg(result); ok {
		return paginate(p.PageLimit, p.Page, val), err
	}
	return model.ReportResponse{}, nil
}

func GenerateSensorDataReport(p model.SensorDataReportRequest) (model.ReportResponse, error) {
	result, err := store.GenerateSensorDataReport(p.CompanyId, p.StartDate, p.EndDate)
	if ok, val := takeSliceArg(result); ok {
		return paginate(p.PageLimit, p.Page, val), err
	}
	return model.ReportResponse{}, nil
}

func GenerateServiceReport(p model.ServiceReportRequest) (model.ReportResponse, error) {
	result, err := store.GenerateServiceReport(p.SerialNumber)
	if ok, val := takeSliceArg(result); ok {
		return paginate(p.PageLimit, p.Page, val), err
	}
	return model.ReportResponse{}, nil
}

func paginate(pageLimit, currentPage int, result []interface{}) model.ReportResponse {
	m := model.ReportResponse{}
	m.Data = result
	count := len(result)
	m.TotalRecords = count
	if pageLimit == 0 {
		m.CurrentPage = 1
		m.TotalPages = 1
	} else {
		//check the limit that its not above count
		/*if pageLimit >= count {
			m.Error = "the page limit is more than the number of items"
			return m
		}*/
		mod := count % pageLimit
		pages := count / pageLimit
		if mod > 0 {
			pages = pages + 1
		}
		//pages
		currentPage--
		if currentPage >= pages {
			//slice the data
			slice := result[currentPage:]
			m.Data = slice
			m.CurrentPage = 1
		} else {
			//A slice is formed by specifying two indices, a low and high bound, separated by a colon:
			//a[low : high]
			//scenario - result has 50 items , we want page 2 and limit to 5
			//page 2 is page + 1 ... to limit
			low := currentPage * pageLimit
			high := low + pageLimit
			//check if high is not above cap
			if high > cap(result) {
				high = cap(result)
			}
			slice := result[low:high]
			m.Data = slice
			m.CurrentPage = currentPage
		}
		m.TotalPages = pages
		m.TotalRecords = count
	}
	return m
}

//https://ahmet.im/blog/golang-take-slices-of-any-type-as-input-parameter/
func takeSliceArg(arg interface{}) (ok bool, out []interface{}) {
	slice, success := takeArg(arg, reflect.Slice)
	if !success {
		ok = false
		return
	}
	c := slice.Len()
	out = make([]interface{}, c)
	for i := 0; i < c; i++ {
		out[i] = slice.Index(i).Interface()
	}
	return true, out
}

func takeArg(arg interface{}, kind reflect.Kind) (val reflect.Value, ok bool) {
	val = reflect.ValueOf(arg)
	if val.Kind() == kind {
		ok = true
	}
	return
}
