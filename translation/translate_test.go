package translation

import (
	"errors"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"testing"
)

func TestTranslateError(t *testing.T) {
	type args struct {
		err  error
		lang string
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "Test case : Nil",
			args: args{
				err:  nil,
				lang: "en",
			},
			wantErr: nil,
		},
		{
			name: "Test case : Validation Ok",
			args: args{
				err: validation.Errors{
					"name": validation.ErrRequired,
				},
				lang: "ja",
			},
			wantErr: validation.Errors{
				"name": errors.New(translations["validation_required"]["ja"]),
			},
		},
		{
			name: "Test case : Validation Error Not Found",
			args: args{
				err: validation.Errors{
					"name": errors.New("fake error"),
				},
				lang: "ja",
			},
			wantErr: validation.Errors{
				"name": errors.New("fake error"),
			},
		},
		{
			name: "Test case : Invalid Lang",
			args: args{
				err: validation.Errors{
					"name": validation.ErrRequired,
				},
				lang: "bd",
			},
			wantErr: validation.Errors{
				"name": validation.ErrRequired,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := TranslateError(tt.args.err, tt.args.lang)
			var validationErrors validation.Errors
			if errors.As(err, &validationErrors) {
				if len(validationErrors) != len(tt.wantErr.(validation.Errors)) {
					t.Errorf("TranslateError() error = %v, wantErr %v", err, tt.wantErr)
				}

				for field, validationErr := range validationErrors {
					var ve validation.Error
					if errors.As(validationErr, &ve) {
						if ve.Error() != tt.wantErr.(validation.Errors)[field].Error() {
							t.Errorf("TranslateError() error = %v, wantErr %v", err, tt.wantErr)
						}
					}
				}
			} else if err != nil && validationErrors != nil {
				t.Errorf("TranslateError() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
