package model

import validation "github.com/go-ozzo/ozzo-validation"

//Катомный валидатор, если условие которое мы передали true то поле мы проверяем на обязательное (Required)
func requiredIf(cond bool) validation.RuleFunc {

	return func(value interface{}) error {
		if cond {
			return validation.Validate(value, validation.Required)
		}

		return nil
	}
}
