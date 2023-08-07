package validator

import (
	"fmt"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
)

func New() *Validator {
	validate := validator.New()
	english := en.New()
	uni := ut.New(english, english)
	trans, _ := uni.GetTranslator("en")
	_ = enTranslations.RegisterDefaultTranslations(validate, trans)

	return &Validator{
		validate,
		trans,
	}
}

type Validator struct {
	validate *validator.Validate
	trans    ut.Translator
}

func (v *Validator) Validate(data any) []error {
	var errs []error

	err := v.validate.Struct(data)

	if err == nil {
		return nil
	}

	validatorErrs := err.(validator.ValidationErrors)

	for _, e := range validatorErrs {
		translatedErr := fmt.Errorf(e.Translate(v.trans))
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
