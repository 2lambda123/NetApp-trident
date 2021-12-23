// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ZappS3BucketApplicationComponents The list of application components to be created.
//
// swagger:model zapp_s3_bucket_application_components
type ZappS3BucketApplicationComponents struct {

	// The list of S3 objectstore policies to be created.
	// Max Items: 10
	// Min Items: 0
	AccessPolicies []*ZappS3BucketApplicationComponentsAccessPolicies `json:"access_policies,omitempty"`

	// Prefer lower latency storage under similar media costs.
	// Enum: [false true]
	CapacityTier bool `json:"capacity_tier,omitempty"`

	// Object Store Server Bucket Description Usage: &lt;(size 1..256)&gt;
	// Max Length: 256
	// Min Length: 1
	Comment string `json:"comment,omitempty"`

	// exclude aggregates
	ExcludeAggregates []*ZappS3BucketApplicationComponentsExcludeAggregatesItems0 `json:"exclude_aggregates,omitempty"`

	// The name of the application component.
	// Required: true
	// Max Length: 63
	// Min Length: 3
	Name *string `json:"name"`

	// qos
	Qos *ZappS3BucketApplicationComponentsQos `json:"qos,omitempty"`

	// The total size of the S3 Bucket, split across the member components. Usage: {&lt;integer&gt;[KB|MB|GB|TB|PB]}
	// Required: true
	Size *int64 `json:"size"`

	// storage service
	StorageService *ZappS3BucketApplicationComponentsStorageService `json:"storage_service,omitempty"`

	// Object Store Server Bucket UUID Usage: &lt;UUID&gt;
	// Read Only: true
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this zapp s3 bucket application components
func (m *ZappS3BucketApplicationComponents) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccessPolicies(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCapacityTier(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateComment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExcludeAggregates(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQos(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSize(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStorageService(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ZappS3BucketApplicationComponents) validateAccessPolicies(formats strfmt.Registry) error {
	if swag.IsZero(m.AccessPolicies) { // not required
		return nil
	}

	iAccessPoliciesSize := int64(len(m.AccessPolicies))

	if err := validate.MinItems("access_policies", "body", iAccessPoliciesSize, 0); err != nil {
		return err
	}

	if err := validate.MaxItems("access_policies", "body", iAccessPoliciesSize, 10); err != nil {
		return err
	}

	for i := 0; i < len(m.AccessPolicies); i++ {
		if swag.IsZero(m.AccessPolicies[i]) { // not required
			continue
		}

		if m.AccessPolicies[i] != nil {
			if err := m.AccessPolicies[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("access_policies" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var zappS3BucketApplicationComponentsTypeCapacityTierPropEnum []interface{}

func init() {
	var res []bool
	if err := json.Unmarshal([]byte(`[false,true]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		zappS3BucketApplicationComponentsTypeCapacityTierPropEnum = append(zappS3BucketApplicationComponentsTypeCapacityTierPropEnum, v)
	}
}

// prop value enum
func (m *ZappS3BucketApplicationComponents) validateCapacityTierEnum(path, location string, value bool) error {
	if err := validate.EnumCase(path, location, value, zappS3BucketApplicationComponentsTypeCapacityTierPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ZappS3BucketApplicationComponents) validateCapacityTier(formats strfmt.Registry) error {
	if swag.IsZero(m.CapacityTier) { // not required
		return nil
	}

	// value enum
	if err := m.validateCapacityTierEnum("capacity_tier", "body", m.CapacityTier); err != nil {
		return err
	}

	return nil
}

func (m *ZappS3BucketApplicationComponents) validateComment(formats strfmt.Registry) error {
	if swag.IsZero(m.Comment) { // not required
		return nil
	}

	if err := validate.MinLength("comment", "body", m.Comment, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("comment", "body", m.Comment, 256); err != nil {
		return err
	}

	return nil
}

func (m *ZappS3BucketApplicationComponents) validateExcludeAggregates(formats strfmt.Registry) error {
	if swag.IsZero(m.ExcludeAggregates) { // not required
		return nil
	}

	for i := 0; i < len(m.ExcludeAggregates); i++ {
		if swag.IsZero(m.ExcludeAggregates[i]) { // not required
			continue
		}

		if m.ExcludeAggregates[i] != nil {
			if err := m.ExcludeAggregates[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("exclude_aggregates" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ZappS3BucketApplicationComponents) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", *m.Name, 3); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", *m.Name, 63); err != nil {
		return err
	}

	return nil
}

func (m *ZappS3BucketApplicationComponents) validateQos(formats strfmt.Registry) error {
	if swag.IsZero(m.Qos) { // not required
		return nil
	}

	if m.Qos != nil {
		if err := m.Qos.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("qos")
			}
			return err
		}
	}

	return nil
}

func (m *ZappS3BucketApplicationComponents) validateSize(formats strfmt.Registry) error {

	if err := validate.Required("size", "body", m.Size); err != nil {
		return err
	}

	return nil
}

func (m *ZappS3BucketApplicationComponents) validateStorageService(formats strfmt.Registry) error {
	if swag.IsZero(m.StorageService) { // not required
		return nil
	}

	if m.StorageService != nil {
		if err := m.StorageService.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("storage_service")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this zapp s3 bucket application components based on the context it is used
func (m *ZappS3BucketApplicationComponents) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAccessPolicies(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateExcludeAggregates(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateQos(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStorageService(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUUID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ZappS3BucketApplicationComponents) contextValidateAccessPolicies(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AccessPolicies); i++ {

		if m.AccessPolicies[i] != nil {
			if err := m.AccessPolicies[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("access_policies" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ZappS3BucketApplicationComponents) contextValidateExcludeAggregates(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ExcludeAggregates); i++ {

		if m.ExcludeAggregates[i] != nil {
			if err := m.ExcludeAggregates[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("exclude_aggregates" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ZappS3BucketApplicationComponents) contextValidateQos(ctx context.Context, formats strfmt.Registry) error {

	if m.Qos != nil {
		if err := m.Qos.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("qos")
			}
			return err
		}
	}

	return nil
}

func (m *ZappS3BucketApplicationComponents) contextValidateStorageService(ctx context.Context, formats strfmt.Registry) error {

	if m.StorageService != nil {
		if err := m.StorageService.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("storage_service")
			}
			return err
		}
	}

	return nil
}

func (m *ZappS3BucketApplicationComponents) contextValidateUUID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "uuid", "body", string(m.UUID)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ZappS3BucketApplicationComponents) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ZappS3BucketApplicationComponents) UnmarshalBinary(b []byte) error {
	var res ZappS3BucketApplicationComponents
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ZappS3BucketApplicationComponentsExcludeAggregatesItems0 zapp s3 bucket application components exclude aggregates items0
//
// swagger:model ZappS3BucketApplicationComponentsExcludeAggregatesItems0
type ZappS3BucketApplicationComponentsExcludeAggregatesItems0 struct {

	// The name of the aggregate to exclude. Usage: &lt;aggr0_svctridentci_vsim3&gt;
	Name string `json:"name,omitempty"`

	// The ID of the aggregate to exclude. Usage: &lt;UUID&gt;
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this zapp s3 bucket application components exclude aggregates items0
func (m *ZappS3BucketApplicationComponentsExcludeAggregatesItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this zapp s3 bucket application components exclude aggregates items0 based on context it is used
func (m *ZappS3BucketApplicationComponentsExcludeAggregatesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ZappS3BucketApplicationComponentsExcludeAggregatesItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ZappS3BucketApplicationComponentsExcludeAggregatesItems0) UnmarshalBinary(b []byte) error {
	var res ZappS3BucketApplicationComponentsExcludeAggregatesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ZappS3BucketApplicationComponentsQos zapp s3 bucket application components qos
//
// swagger:model ZappS3BucketApplicationComponentsQos
type ZappS3BucketApplicationComponentsQos struct {

	// policy
	Policy *ZappS3BucketApplicationComponentsQosPolicy `json:"policy,omitempty"`
}

// Validate validates this zapp s3 bucket application components qos
func (m *ZappS3BucketApplicationComponentsQos) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePolicy(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ZappS3BucketApplicationComponentsQos) validatePolicy(formats strfmt.Registry) error {
	if swag.IsZero(m.Policy) { // not required
		return nil
	}

	if m.Policy != nil {
		if err := m.Policy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("qos" + "." + "policy")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this zapp s3 bucket application components qos based on the context it is used
func (m *ZappS3BucketApplicationComponentsQos) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePolicy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ZappS3BucketApplicationComponentsQos) contextValidatePolicy(ctx context.Context, formats strfmt.Registry) error {

	if m.Policy != nil {
		if err := m.Policy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("qos" + "." + "policy")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ZappS3BucketApplicationComponentsQos) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ZappS3BucketApplicationComponentsQos) UnmarshalBinary(b []byte) error {
	var res ZappS3BucketApplicationComponentsQos
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ZappS3BucketApplicationComponentsQosPolicy zapp s3 bucket application components qos policy
//
// swagger:model ZappS3BucketApplicationComponentsQosPolicy
type ZappS3BucketApplicationComponentsQosPolicy struct {

	// The name of an existing QoS policy.
	Name string `json:"name,omitempty"`

	// The UUID of an existing QoS policy. Usage: &lt;UUID&gt;
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this zapp s3 bucket application components qos policy
func (m *ZappS3BucketApplicationComponentsQosPolicy) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this zapp s3 bucket application components qos policy based on context it is used
func (m *ZappS3BucketApplicationComponentsQosPolicy) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ZappS3BucketApplicationComponentsQosPolicy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ZappS3BucketApplicationComponentsQosPolicy) UnmarshalBinary(b []byte) error {
	var res ZappS3BucketApplicationComponentsQosPolicy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ZappS3BucketApplicationComponentsStorageService zapp s3 bucket application components storage service
//
// swagger:model ZappS3BucketApplicationComponentsStorageService
type ZappS3BucketApplicationComponentsStorageService struct {

	// The storage service of the application component.
	// Enum: [extreme performance value]
	Name *string `json:"name,omitempty"`
}

// Validate validates this zapp s3 bucket application components storage service
func (m *ZappS3BucketApplicationComponentsStorageService) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var zappS3BucketApplicationComponentsStorageServiceTypeNamePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["extreme","performance","value"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		zappS3BucketApplicationComponentsStorageServiceTypeNamePropEnum = append(zappS3BucketApplicationComponentsStorageServiceTypeNamePropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// ZappS3BucketApplicationComponentsStorageService
	// ZappS3BucketApplicationComponentsStorageService
	// name
	// Name
	// extreme
	// END DEBUGGING
	// ZappS3BucketApplicationComponentsStorageServiceNameExtreme captures enum value "extreme"
	ZappS3BucketApplicationComponentsStorageServiceNameExtreme string = "extreme"

	// BEGIN DEBUGGING
	// ZappS3BucketApplicationComponentsStorageService
	// ZappS3BucketApplicationComponentsStorageService
	// name
	// Name
	// performance
	// END DEBUGGING
	// ZappS3BucketApplicationComponentsStorageServiceNamePerformance captures enum value "performance"
	ZappS3BucketApplicationComponentsStorageServiceNamePerformance string = "performance"

	// BEGIN DEBUGGING
	// ZappS3BucketApplicationComponentsStorageService
	// ZappS3BucketApplicationComponentsStorageService
	// name
	// Name
	// value
	// END DEBUGGING
	// ZappS3BucketApplicationComponentsStorageServiceNameValue captures enum value "value"
	ZappS3BucketApplicationComponentsStorageServiceNameValue string = "value"
)

// prop value enum
func (m *ZappS3BucketApplicationComponentsStorageService) validateNameEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, zappS3BucketApplicationComponentsStorageServiceTypeNamePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ZappS3BucketApplicationComponentsStorageService) validateName(formats strfmt.Registry) error {
	if swag.IsZero(m.Name) { // not required
		return nil
	}

	// value enum
	if err := m.validateNameEnum("storage_service"+"."+"name", "body", *m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this zapp s3 bucket application components storage service based on context it is used
func (m *ZappS3BucketApplicationComponentsStorageService) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ZappS3BucketApplicationComponentsStorageService) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ZappS3BucketApplicationComponentsStorageService) UnmarshalBinary(b []byte) error {
	var res ZappS3BucketApplicationComponentsStorageService
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
