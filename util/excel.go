package util

import (
	"backend/api"
	"errors"
	"strconv"

	"github.com/xuri/excelize/v2"
)




func BindToDailyMileageRequest(str []string) (api.DailyMileageRequest, error) {

	if len(str) < 2 {
		return api.DailyMileageRequest{}, errors.New(api.DailyMileageReqStatusInvalidInput)
	}

	amount, err := strconv.Atoi(str[1])
	if err != nil {
		return api.DailyMileageRequest{}, errors.New(api.DailyMileageReqStatusInvalidInput)
	}

	request := api.DailyMileageRequest{
		Key: str[0],
		Amount: amount,
		Weekday: str[2],
	}

	if err := api.WeekdayValid(request.Weekday); err != nil {
		return api.DailyMileageRequest{}, err
	}

	return request, err
}



func ExtractDataFromExcel(filepath string) ([]api.DailyMileageRequest, error) {
	f, err := excelize.OpenFile(filepath)
	if err != nil {
		return nil, err
	}

	rows, err := f.GetRows("Sheet1")
	if err != nil {
		return nil, err
	}

	var result []api.DailyMileageRequest

	for _, row := range rows {
		item, err := BindToDailyMileageRequest(row)
		if err != nil {
			item.State = err.Error()
		}

	    result = append(result, item)
	}

	return result, nil	
}

