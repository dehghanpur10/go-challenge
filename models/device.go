package models

//Device is model that should save in dynamoDB
type Device struct {
	Id          string `json:"id" validate:"required" example:"/devices/id1"`
	DeviceModel string `json:"deviceModel" validate:"required" example:"/devicemodels/id1"`
	Name        string `json:"name" validate:"required" example:"Sensor"`
	Note        string `json:"note" validate:"required" example:"Testing a sensor"`
	Serial      string `json:"serial" validate:"required" example:"A020000102"`
}
