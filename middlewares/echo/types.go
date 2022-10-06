package echo

import (
	v "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

const (
	timeUnitHour   string = "HOUR"
	timeUnitMinute string = "MINUTE"
	timeUnitSecond string = "SECOND"
)

type emailReqData struct {
	Email string `json:"email" validate:"required,email"`
}

func (f *emailReqData) Validate() error {
	return v.ValidateStruct(f,
		v.Field(&f.Email, v.Required, is.EmailFormat),
	)
}
