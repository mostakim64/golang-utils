package echo

const (
	timeUnitHour   string = "HOUR"
	timeUnitMinute string = "MINUTE"
	timeUnitSecond string = "SECOND"
)

type emailReqData struct {
	Email string `json:"email" validate:"required,email"`
}
