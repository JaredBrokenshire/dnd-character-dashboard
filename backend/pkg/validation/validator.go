package validation

import (
	"errors"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	v "github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
)

const minPasswordLength = 12

type CustomValidator struct {
	validator *v.Validate
	uni       *ut.UniversalTranslator
	trans     ut.Translator
}

func NewCustomValidator(validator *v.Validate) *CustomValidator {

	cv := &CustomValidator{validator: validator}

	translator := en.New()
	cv.uni = ut.New(translator, translator)
	trans, _ := cv.uni.GetTranslator("translator")
	cv.trans = trans

	_ = enTranslations.RegisterDefaultTranslations(validator, trans)

	return cv
}

func (cv *CustomValidator) Validate(i interface{}) error {
	err := cv.validator.Struct(i)
	if err == nil {
		return nil
	}
	var errs v.ValidationErrors
	errors.As(err, &errs)
	msg := ""
	for _, valErrorTranslation := range errs.Translate(cv.trans) {
		if msg != "" {
			msg += ", "
		}
		msg += valErrorTranslation
	}
	return errors.New(msg)
}
