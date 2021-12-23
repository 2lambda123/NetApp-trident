// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// QosWorkload qos workload
//
// swagger:model qos_workload
type QosWorkload struct {

	// links
	Links *QosWorkloadLinks `json:"_links,omitempty"`

	// Name of the file.
	// Read Only: true
	File string `json:"file,omitempty"`

	// Name of the LUN. The name of the LUN will be displayed as "(unknown)" if the name cannot be retrieved.
	// Read Only: true
	Lun string `json:"lun,omitempty"`

	// Name of the QoS workload.
	// Example: volume1-wid123
	Name string `json:"name,omitempty"`

	// policy
	Policy *QosWorkloadPolicy `json:"policy,omitempty"`

	// Name of the Qtree.
	// Read Only: true
	Qtree string `json:"qtree,omitempty"`

	// svm
	Svm *QosWorkloadSvm `json:"svm,omitempty"`

	// uuid
	// Example: 1cd8a442-86d1-11e0-ae1c-123478563412
	// Read Only: true
	UUID string `json:"uuid,omitempty"`

	// Name of the volume. The name of the volume will be displayed as "(unknown)" if the name cannot be retrieved.
	// Example: volume1
	// Read Only: true
	Volume string `json:"volume,omitempty"`

	// Workload ID of the QoS workload.
	// Example: 123
	// Read Only: true
	Wid int64 `json:"wid,omitempty"`

	// Class of the QoS workload.
	// Example: autovolume
	// Read Only: true
	// Enum: [undefined preset user_defined system_defined autovolume load_control]
	WorkloadClass string `json:"workload_class,omitempty"`
}

// Validate validates this qos workload
func (m *QosWorkload) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSvm(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWorkloadClass(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *QosWorkload) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links")
			}
			return err
		}
	}

	return nil
}

func (m *QosWorkload) validatePolicy(formats strfmt.Registry) error {
	if swag.IsZero(m.Policy) { // not required
		return nil
	}

	if m.Policy != nil {
		if err := m.Policy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("policy")
			}
			return err
		}
	}

	return nil
}

func (m *QosWorkload) validateSvm(formats strfmt.Registry) error {
	if swag.IsZero(m.Svm) { // not required
		return nil
	}

	if m.Svm != nil {
		if err := m.Svm.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm")
			}
			return err
		}
	}

	return nil
}

var qosWorkloadTypeWorkloadClassPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["undefined","preset","user_defined","system_defined","autovolume","load_control"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		qosWorkloadTypeWorkloadClassPropEnum = append(qosWorkloadTypeWorkloadClassPropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// qos_workload
	// QosWorkload
	// workload_class
	// WorkloadClass
	// undefined
	// END DEBUGGING
	// QosWorkloadWorkloadClassUndefined captures enum value "undefined"
	QosWorkloadWorkloadClassUndefined string = "undefined"

	// BEGIN DEBUGGING
	// qos_workload
	// QosWorkload
	// workload_class
	// WorkloadClass
	// preset
	// END DEBUGGING
	// QosWorkloadWorkloadClassPreset captures enum value "preset"
	QosWorkloadWorkloadClassPreset string = "preset"

	// BEGIN DEBUGGING
	// qos_workload
	// QosWorkload
	// workload_class
	// WorkloadClass
	// user_defined
	// END DEBUGGING
	// QosWorkloadWorkloadClassUserDefined captures enum value "user_defined"
	QosWorkloadWorkloadClassUserDefined string = "user_defined"

	// BEGIN DEBUGGING
	// qos_workload
	// QosWorkload
	// workload_class
	// WorkloadClass
	// system_defined
	// END DEBUGGING
	// QosWorkloadWorkloadClassSystemDefined captures enum value "system_defined"
	QosWorkloadWorkloadClassSystemDefined string = "system_defined"

	// BEGIN DEBUGGING
	// qos_workload
	// QosWorkload
	// workload_class
	// WorkloadClass
	// autovolume
	// END DEBUGGING
	// QosWorkloadWorkloadClassAutovolume captures enum value "autovolume"
	QosWorkloadWorkloadClassAutovolume string = "autovolume"

	// BEGIN DEBUGGING
	// qos_workload
	// QosWorkload
	// workload_class
	// WorkloadClass
	// load_control
	// END DEBUGGING
	// QosWorkloadWorkloadClassLoadControl captures enum value "load_control"
	QosWorkloadWorkloadClassLoadControl string = "load_control"
)

// prop value enum
func (m *QosWorkload) validateWorkloadClassEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, qosWorkloadTypeWorkloadClassPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *QosWorkload) validateWorkloadClass(formats strfmt.Registry) error {
	if swag.IsZero(m.WorkloadClass) { // not required
		return nil
	}

	// value enum
	if err := m.validateWorkloadClassEnum("workload_class", "body", m.WorkloadClass); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this qos workload based on the context it is used
func (m *QosWorkload) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFile(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLun(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePolicy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateQtree(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSvm(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUUID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVolume(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWid(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWorkloadClass(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *QosWorkload) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links")
			}
			return err
		}
	}

	return nil
}

func (m *QosWorkload) contextValidateFile(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "file", "body", string(m.File)); err != nil {
		return err
	}

	return nil
}

func (m *QosWorkload) contextValidateLun(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "lun", "body", string(m.Lun)); err != nil {
		return err
	}

	return nil
}

func (m *QosWorkload) contextValidatePolicy(ctx context.Context, formats strfmt.Registry) error {

	if m.Policy != nil {
		if err := m.Policy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("policy")
			}
			return err
		}
	}

	return nil
}

func (m *QosWorkload) contextValidateQtree(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "qtree", "body", string(m.Qtree)); err != nil {
		return err
	}

	return nil
}

func (m *QosWorkload) contextValidateSvm(ctx context.Context, formats strfmt.Registry) error {

	if m.Svm != nil {
		if err := m.Svm.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm")
			}
			return err
		}
	}

	return nil
}

func (m *QosWorkload) contextValidateUUID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "uuid", "body", string(m.UUID)); err != nil {
		return err
	}

	return nil
}

func (m *QosWorkload) contextValidateVolume(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "volume", "body", string(m.Volume)); err != nil {
		return err
	}

	return nil
}

func (m *QosWorkload) contextValidateWid(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "wid", "body", int64(m.Wid)); err != nil {
		return err
	}

	return nil
}

func (m *QosWorkload) contextValidateWorkloadClass(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "workload_class", "body", string(m.WorkloadClass)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *QosWorkload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *QosWorkload) UnmarshalBinary(b []byte) error {
	var res QosWorkload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// QosWorkloadLinks qos workload links
//
// swagger:model QosWorkloadLinks
type QosWorkloadLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this qos workload links
func (m *QosWorkloadLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *QosWorkloadLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this qos workload links based on the context it is used
func (m *QosWorkloadLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *QosWorkloadLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *QosWorkloadLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *QosWorkloadLinks) UnmarshalBinary(b []byte) error {
	var res QosWorkloadLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// QosWorkloadPolicy QoS policy group reference.
//
// swagger:model QosWorkloadPolicy
type QosWorkloadPolicy struct {

	// links
	Links *QosWorkloadPolicyLinks `json:"_links,omitempty"`

	// The QoS policy group name. This is mutually exclusive with UUID and other QoS attributes during POST and PATCH.
	// Example: performance
	Name string `json:"name,omitempty"`

	// The QoS policy group UUID. This is mutually exclusive with name and other QoS attributes during POST and PATCH.
	// Example: 1cd8a442-86d1-11e0-ae1c-123478563412
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this qos workload policy
func (m *QosWorkloadPolicy) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *QosWorkloadPolicy) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("policy" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this qos workload policy based on the context it is used
func (m *QosWorkloadPolicy) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *QosWorkloadPolicy) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("policy" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *QosWorkloadPolicy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *QosWorkloadPolicy) UnmarshalBinary(b []byte) error {
	var res QosWorkloadPolicy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// QosWorkloadPolicyLinks qos workload policy links
//
// swagger:model QosWorkloadPolicyLinks
type QosWorkloadPolicyLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this qos workload policy links
func (m *QosWorkloadPolicyLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *QosWorkloadPolicyLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("policy" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this qos workload policy links based on the context it is used
func (m *QosWorkloadPolicyLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *QosWorkloadPolicyLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("policy" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *QosWorkloadPolicyLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *QosWorkloadPolicyLinks) UnmarshalBinary(b []byte) error {
	var res QosWorkloadPolicyLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// QosWorkloadSvm qos workload svm
//
// swagger:model QosWorkloadSvm
type QosWorkloadSvm struct {

	// links
	Links *QosWorkloadSvmLinks `json:"_links,omitempty"`

	// The name of the SVM.
	//
	// Example: svm1
	Name string `json:"name,omitempty"`

	// The unique identifier of the SVM.
	//
	// Example: 02c9e252-41be-11e9-81d5-00a0986138f7
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this qos workload svm
func (m *QosWorkloadSvm) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *QosWorkloadSvm) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this qos workload svm based on the context it is used
func (m *QosWorkloadSvm) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *QosWorkloadSvm) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *QosWorkloadSvm) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *QosWorkloadSvm) UnmarshalBinary(b []byte) error {
	var res QosWorkloadSvm
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// QosWorkloadSvmLinks qos workload svm links
//
// swagger:model QosWorkloadSvmLinks
type QosWorkloadSvmLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this qos workload svm links
func (m *QosWorkloadSvmLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *QosWorkloadSvmLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this qos workload svm links based on the context it is used
func (m *QosWorkloadSvmLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *QosWorkloadSvmLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *QosWorkloadSvmLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *QosWorkloadSvmLinks) UnmarshalBinary(b []byte) error {
	var res QosWorkloadSvmLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
