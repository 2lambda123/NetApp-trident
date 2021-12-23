// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// LunMapReportingNodeReference A cluster node from which network paths to the LUN are advertised by ONTAP via the SAN protocols.
//
//
// swagger:model lun_map_reporting_node_reference
type LunMapReportingNodeReference struct {

	// links
	Links *LunMapReportingNodeReferenceLinks `json:"_links,omitempty"`

	// The name of the node.<br/>
	// Either `uuid` or `name` are required in POST.
	//
	// Example: node1
	Name string `json:"name,omitempty"`

	// The unique identifier of the node.<br/>
	// Either `uuid` or `name` are required in POST.
	//
	// Example: 5ac8eb9c-4e32-dbaa-57ca-fb905976f54e
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this lun map reporting node reference
func (m *LunMapReportingNodeReference) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LunMapReportingNodeReference) validateLinks(formats strfmt.Registry) error {
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

// ContextValidate validate this lun map reporting node reference based on the context it is used
func (m *LunMapReportingNodeReference) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LunMapReportingNodeReference) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *LunMapReportingNodeReference) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LunMapReportingNodeReference) UnmarshalBinary(b []byte) error {
	var res LunMapReportingNodeReference
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// LunMapReportingNodeReferenceLinks lun map reporting node reference links
//
// swagger:model LunMapReportingNodeReferenceLinks
type LunMapReportingNodeReferenceLinks struct {

	// node
	Node *Href `json:"node,omitempty"`

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this lun map reporting node reference links
func (m *LunMapReportingNodeReferenceLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LunMapReportingNodeReferenceLinks) validateNode(formats strfmt.Registry) error {
	if swag.IsZero(m.Node) { // not required
		return nil
	}

	if m.Node != nil {
		if err := m.Node.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "node")
			}
			return err
		}
	}

	return nil
}

func (m *LunMapReportingNodeReferenceLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this lun map reporting node reference links based on the context it is used
func (m *LunMapReportingNodeReferenceLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LunMapReportingNodeReferenceLinks) contextValidateNode(ctx context.Context, formats strfmt.Registry) error {

	if m.Node != nil {
		if err := m.Node.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "node")
			}
			return err
		}
	}

	return nil
}

func (m *LunMapReportingNodeReferenceLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *LunMapReportingNodeReferenceLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LunMapReportingNodeReferenceLinks) UnmarshalBinary(b []byte) error {
	var res LunMapReportingNodeReferenceLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
