package services

var (
	DATE_SERVICE_MOCK_GetCurrentDate string
)

type DATE_SERVICE_MOCK struct {}

func (d DATE_SERVICE_MOCK) GetCurrentDate() string {
	return DATE_SERVICE_MOCK_GetCurrentDate
}

func (d DATE_SERVICE_MOCK) GivenDateServiceReturns(s string) {
	DATE_SERVICE_MOCK_GetCurrentDate = s
}