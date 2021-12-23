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

// LocalHost local host
//
// swagger:model local_host
type LocalHost struct {

	// links
	Links *LocalHostLinks `json:"_links,omitempty"`

	// IPv4/IPv6 address in dotted form.
	// Example: 123.123.123.123
	Address string `json:"address,omitempty"`

	// The list of aliases.
	// Example: ["host1.sales.foo.com","host2.sakes.foo.com"]
	Aliases []string `json:"aliases,omitempty"`

	// Canonical hostname.
	// Example: host.sales.foo.com
	// Max Length: 255
	// Min Length: 1
	Hostname string `json:"hostname,omitempty"`

	// owner
	Owner *LocalHostOwner `json:"owner,omitempty"`

	// Scope of the entity. Set to "cluster" for cluster owned objects and to "svm" for SVM owned objects.
	// Enum: [cluster svm]
	Scope string `json:"scope,omitempty"`
}

// Validate validates this local host
func (m *LocalHost) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHostname(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOwner(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScope(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LocalHost) validateLinks(formats strfmt.Registry) error {
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

func (m *LocalHost) validateHostname(formats strfmt.Registry) error {
	if swag.IsZero(m.Hostname) { // not required
		return nil
	}

	if err := validate.MinLength("hostname", "body", m.Hostname, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("hostname", "body", m.Hostname, 255); err != nil {
		return err
	}

	return nil
}

func (m *LocalHost) validateOwner(formats strfmt.Registry) error {
	if swag.IsZero(m.Owner) { // not required
		return nil
	}

	if m.Owner != nil {
		if err := m.Owner.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("owner")
			}
			return err
		}
	}

	return nil
}

var localHostTypeScopePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["cluster","svm"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		localHostTypeScopePropEnum = append(localHostTypeScopePropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// local_host
	// LocalHost
	// scope
	// Scope
	// cluster
	// END DEBUGGING
	// LocalHostScopeCluster captures enum value "cluster"
	LocalHostScopeCluster string = "cluster"

	// BEGIN DEBUGGING
	// local_host
	// LocalHost
	// scope
	// Scope
	// svm
	// END DEBUGGING
	// LocalHostScopeSvm captures enum value "svm"
	LocalHostScopeSvm string = "svm"
)

// prop value enum
func (m *LocalHost) validateScopeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, localHostTypeScopePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *LocalHost) validateScope(formats strfmt.Registry) error {
	if swag.IsZero(m.Scope) { // not required
		return nil
	}

	// value enum
	if err := m.validateScopeEnum("scope", "body", m.Scope); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this local host based on the context it is used
func (m *LocalHost) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOwner(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LocalHost) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *LocalHost) contextValidateOwner(ctx context.Context, formats strfmt.Registry) error {

	if m.Owner != nil {
		if err := m.Owner.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("owner")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LocalHost) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LocalHost) UnmarshalBinary(b []byte) error {
	var res LocalHost
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// LocalHostLinks local host links
//
// swagger:model LocalHostLinks
type LocalHostLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this local host links
func (m *LocalHostLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LocalHostLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this local host links based on the context it is used
func (m *LocalHostLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LocalHostLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *LocalHostLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LocalHostLinks) UnmarshalBinary(b []byte) error {
	var res LocalHostLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// LocalHostOwner local host owner
//
// swagger:model LocalHostOwner
type LocalHostOwner struct {

	// links
	Links *LocalHostOwnerLinks `json:"_links,omitempty"`

	// The name of the SVM.
	//
	// Example: svm1
	Name string `json:"name,omitempty"`

	// The unique identifier of the SVM.
	//
	// Example: 02c9e252-41be-11e9-81d5-00a0986138f7
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this local host owner
func (m *LocalHostOwner) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LocalHostOwner) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("owner" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this local host owner based on the context it is used
func (m *LocalHostOwner) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LocalHostOwner) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("owner" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LocalHostOwner) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LocalHostOwner) UnmarshalBinary(b []byte) error {
	var res LocalHostOwner
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// LocalHostOwnerLinks local host owner links
//
// swagger:model LocalHostOwnerLinks
type LocalHostOwnerLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this local host owner links
func (m *LocalHostOwnerLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LocalHostOwnerLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("owner" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this local host owner links based on the context it is used
func (m *LocalHostOwnerLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LocalHostOwnerLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("owner" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LocalHostOwnerLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LocalHostOwnerLinks) UnmarshalBinary(b []byte) error {
	var res LocalHostOwnerLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
