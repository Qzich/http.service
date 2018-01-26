package validation

type ValidatorFunc func() bool

type FieldValidator struct {
	FieldName     string
	ErrorMessage  string
	ValidatorFunc ValidatorFunc
}

type FieldErrors []FieldError

func (fieldErrors FieldErrors) IsFieldValid(fieldName string) bool {
	for _, fieldError := range fieldErrors {
		if fieldError.FieldName == fieldName {
			return false
		}
	}

	return true
}

type FieldError struct {
	FieldName    string
	ErrorMessage string
}

func (fieldError *FieldError) Error() string {
	return fieldError.ErrorMessage
}

type Validator struct {
	Validators       []FieldValidator
	ValidationErrors FieldErrors
}

func NewValidator() *Validator {
	return &Validator{
		Validators:       []FieldValidator{},
		ValidationErrors: FieldErrors{},
	}
}

func (dataValidation *Validator) GetFieldErrors(fieldName string) []string {
	fieldErrors := []string{}

	for _, fieldError := range dataValidation.ValidationErrors {
		if fieldError.FieldName == fieldName {
			fieldErrors = append(fieldErrors, fieldError.ErrorMessage)
		}
	}

	return fieldErrors
}

func (dataValidation *Validator) GetErrors() FieldErrors {
	fieldErrors := dataValidation.ValidationErrors

	return fieldErrors
}

func (dataValidation *Validator) AddValidator(validatorFunc ValidatorFunc, fieldName string, errorMessage string) {
	fieldValidator := FieldValidator{
		FieldName:     fieldName,
		ErrorMessage:  errorMessage,
		ValidatorFunc: validatorFunc,
	}

	dataValidation.Validators = append(dataValidation.Validators, fieldValidator)
}

func (dataValidation *Validator) Validate() bool {

	isValid := true

	dataValidation.ValidationErrors = FieldErrors{}

	for _, fieldValidator := range dataValidation.Validators {
		isFieldValid := fieldValidator.ValidatorFunc()

		if !isFieldValid {
			isValid = false

			fieldError := FieldError{FieldName: fieldValidator.FieldName, ErrorMessage: fieldValidator.ErrorMessage,}

			dataValidation.ValidationErrors = append(dataValidation.ValidationErrors, fieldError)
		}
	}

	return isValid
}

type ValidationQueue struct {
	currentValidator *Validator
	validatorList    []*Validator
}

func (validationSequence *ValidationQueue) Then(validator *Validator) *ValidationQueue {
	validationSequence.validatorList = append(validationSequence.validatorList, validator)

	return validationSequence
}

func (validationSequence *ValidationQueue) Validate() bool {
	if validationSequence.currentValidator.Validate() {
		for _, val := range validationSequence.validatorList {
			if !val.Validate() {
				validationSequence.currentValidator = val

				return false
			}
		}

		return true
	}
	return false
}

func (validationSequence *ValidationQueue) GetFieldErrors(fieldName string) []string {
	return validationSequence.currentValidator.GetFieldErrors(fieldName)
}

func (validationSequence *ValidationQueue) GetErrors() FieldErrors {
	return validationSequence.currentValidator.GetErrors()
}

func Validate(validator *Validator) *ValidationQueue {
	return &ValidationQueue{currentValidator: validator, validatorList: []*Validator{}}
}
