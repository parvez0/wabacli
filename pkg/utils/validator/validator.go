package validator

import (
	"fmt"
	"github.com/parvez0/wabacli/pkg/errutil/badrequest"
	"reflect"
	"strings"
)

const (
	TagName Type = "validate"
	FieldRequired Type = "required"
)

// Type custom type for different validator names
type Type string

// Validator provides an interface for all the different
// types of validators, all those should implement the method
// validate
type Validator interface {
	Validate(string, interface{}) error
}

// RequiredFields is a type of validator which validates fields with
// tag required, and it throws and error if the value is empty
type RequiredFields struct {}


// Validate takes an interface as input and get the tags of each fields
// after processing the tags it will get a appropriate validator and
// calls the validate func
func Validate(data interface{}) []error {
	var errs []error
	rval := reflect.ValueOf(data)
	if rval.Kind() == reflect.Ptr {
		rval = reflect.Indirect(rval)
	}
	for i:=0; i < rval.NumField(); i++ {
		tag := rval.Type().Field(i).Tag.Get(string(TagName))
		if tag == "" || tag == "-" {
			continue
		}
		ops := strings.Split(tag, ",")
		for _, v := range ops {
			validator := getValidatorFromTag(Type(v))
			if validator == nil {
				continue
			}
			err := validator.Validate(rval.Type().Field(i).Name, rval.Field(i).Interface())
			if err != nil {
				errs = append(errs, err)
			}
		}
	}
	return errs
}

func (re *RequiredFields) Validate(field string, val interface{}) error {
	badReq := &badrequest.BadRequest{
		Code:        400,
		Title:       "Required filed not provided",
		Description: fmt.Sprintf("ValidationError(RequiredFiled) missing required field \"%s\";", field),
	}
	switch reflect.ValueOf(val).Kind() {
	case reflect.String:
		if s := val.(string); s == "" {
			badReq.Description = fmt.Sprintf("ValidationError(RequiredFiled) missing required field \"%s\";", field)
			return badReq
		}
	case reflect.Int:
		if s := val.(int); s == 0 {
			badReq.Description = fmt.Sprintf("ValidationError(RequiredFiled) missing required field \"%s\";", field)
			return badReq
		}
	}
	return nil
}

// getValidatorFromTag returns the validator based on tag provided in
// field, in case of multiple tags it needs to be called multiple times
func getValidatorFromTag(tag Type) Validator {
	switch tag {
	case FieldRequired:
		return &RequiredFields{}
	default:
		return nil
	}
}



