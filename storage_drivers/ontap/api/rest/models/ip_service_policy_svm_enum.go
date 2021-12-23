// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// IPServicePolicySvmEnum Built-in service policies for SVMs.
//
// swagger:model ip_service_policy_svm_enum
type IPServicePolicySvmEnum string

const (

	// IPServicePolicySvmEnumDefaultDashManagement captures enum value "default-management"
	IPServicePolicySvmEnumDefaultDashManagement IPServicePolicySvmEnum = "default-management"

	// IPServicePolicySvmEnumDefaultDashDataDashFiles captures enum value "default-data-files"
	IPServicePolicySvmEnumDefaultDashDataDashFiles IPServicePolicySvmEnum = "default-data-files"

	// IPServicePolicySvmEnumDefaultDashDataDashBlocks captures enum value "default-data-blocks"
	IPServicePolicySvmEnumDefaultDashDataDashBlocks IPServicePolicySvmEnum = "default-data-blocks"
)

// for schema
var ipServicePolicySvmEnumEnum []interface{}

func init() {
	var res []IPServicePolicySvmEnum
	if err := json.Unmarshal([]byte(`["default-management","default-data-files","default-data-blocks"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		ipServicePolicySvmEnumEnum = append(ipServicePolicySvmEnumEnum, v)
	}
}

func (m IPServicePolicySvmEnum) validateIPServicePolicySvmEnumEnum(path, location string, value IPServicePolicySvmEnum) error {
	if err := validate.EnumCase(path, location, value, ipServicePolicySvmEnumEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this ip service policy svm enum
func (m IPServicePolicySvmEnum) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateIPServicePolicySvmEnumEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this ip service policy svm enum based on context it is used
func (m IPServicePolicySvmEnum) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
