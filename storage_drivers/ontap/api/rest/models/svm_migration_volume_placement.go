// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SvmMigrationVolumePlacement Volume selection information
//
// swagger:model svm_migration_volume_placement
type SvmMigrationVolumePlacement struct {

	// Optional property used to specify the list of desired aggregates to use for volume creation in the destination.
	Aggregates []*SvmMigrationVolumePlacementAggregatesItems0 `json:"aggregates,omitempty"`
}

// Validate validates this svm migration volume placement
func (m *SvmMigrationVolumePlacement) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAggregates(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SvmMigrationVolumePlacement) validateAggregates(formats strfmt.Registry) error {
	if swag.IsZero(m.Aggregates) { // not required
		return nil
	}

	for i := 0; i < len(m.Aggregates); i++ {
		if swag.IsZero(m.Aggregates[i]) { // not required
			continue
		}

		if m.Aggregates[i] != nil {
			if err := m.Aggregates[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("aggregates" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this svm migration volume placement based on the context it is used
func (m *SvmMigrationVolumePlacement) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAggregates(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SvmMigrationVolumePlacement) contextValidateAggregates(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Aggregates); i++ {

		if m.Aggregates[i] != nil {
			if err := m.Aggregates[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("aggregates" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *SvmMigrationVolumePlacement) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SvmMigrationVolumePlacement) UnmarshalBinary(b []byte) error {
	var res SvmMigrationVolumePlacement
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SvmMigrationVolumePlacementAggregatesItems0 Aggregate
//
// swagger:model SvmMigrationVolumePlacementAggregatesItems0
type SvmMigrationVolumePlacementAggregatesItems0 struct {

	// links
	Links *SvmMigrationVolumePlacementAggregatesItems0Links `json:"_links,omitempty"`

	// name
	// Example: aggr1
	Name string `json:"name,omitempty"`

	// uuid
	// Example: 1cd8a442-86d1-11e0-ae1c-123478563412
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this svm migration volume placement aggregates items0
func (m *SvmMigrationVolumePlacementAggregatesItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SvmMigrationVolumePlacementAggregatesItems0) validateLinks(formats strfmt.Registry) error {
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

// ContextValidate validate this svm migration volume placement aggregates items0 based on the context it is used
func (m *SvmMigrationVolumePlacementAggregatesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SvmMigrationVolumePlacementAggregatesItems0) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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
func (m *SvmMigrationVolumePlacementAggregatesItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SvmMigrationVolumePlacementAggregatesItems0) UnmarshalBinary(b []byte) error {
	var res SvmMigrationVolumePlacementAggregatesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SvmMigrationVolumePlacementAggregatesItems0Links svm migration volume placement aggregates items0 links
//
// swagger:model SvmMigrationVolumePlacementAggregatesItems0Links
type SvmMigrationVolumePlacementAggregatesItems0Links struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this svm migration volume placement aggregates items0 links
func (m *SvmMigrationVolumePlacementAggregatesItems0Links) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SvmMigrationVolumePlacementAggregatesItems0Links) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this svm migration volume placement aggregates items0 links based on the context it is used
func (m *SvmMigrationVolumePlacementAggregatesItems0Links) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SvmMigrationVolumePlacementAggregatesItems0Links) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *SvmMigrationVolumePlacementAggregatesItems0Links) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SvmMigrationVolumePlacementAggregatesItems0Links) UnmarshalBinary(b []byte) error {
	var res SvmMigrationVolumePlacementAggregatesItems0Links
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
