package day6

func ValidateField(value string, fieldName string) error {
	ce := &CustomError{}
	if fieldName == "address" {
		if len(value) < 5 {
			ce.errcontext = "validating address"
			ce.filename = "validate_fields.go"
			ce.lineNum = 6
			ce.message = "address cannot be less than 5"
			return ce
		}
		return nil
	} else if fieldName == "name" {
		if len(value) < 3 {
			ce.errcontext = "validating name"
			ce.filename = "validate_fields.go"
			ce.lineNum = 15
			ce.message = "name cannot be less than 5"
			return ce
		}
		return nil
	}
	return nil
}
