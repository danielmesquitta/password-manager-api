package validator

import (
	"fmt"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
)

var validate *validator.Validate
var trans ut.Translator

func NewValidator() *Validator {
	return &Validator{}
}

func Init() {
	validate = validator.New()
	english := en.New()
	uni := ut.New(english, english)
	trans, _ = uni.GetTranslator("en")
	_ = enTranslations.RegisterDefaultTranslations(validate, trans)
}

type Validator struct{}

func (v *Validator) Validate(data any) []error {
	var errs []error

	err := validate.Struct(data)

	if err == nil {
		return nil
	}

	validatorErrs := err.(validator.ValidationErrors)

	for _, e := range validatorErrs {
		translatedErr := fmt.Errorf(e.Translate(trans))
		errs = append(errs, translatedErr)
	}

	return errs
}

func (v *Validator) FormatErrs(errs []error) string {
	var s string

	for index, err := range errs {
		s += err.Error()
		if index < len(errs)-1 {
			s += ", "
		} else {
			s += "."
		}
	}

	return s
}
