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

// InterfaceMetrics Throughput performance for the interfaces.
//
// swagger:model interface_metrics
type InterfaceMetrics struct {

	// links
	Links *InterfaceMetricsLinks `json:"_links,omitempty"`

	// The duration over which this sample is calculated. The time durations are represented in the ISO-8601 standard format. Samples can be calculated over the following durations:
	//
	// Example: PT15S
	// Enum: [PT15S PT4M PT30M PT2H P1D PT5M]
	Duration string `json:"duration,omitempty"`

	// Errors associated with the sample. For example, if the aggregation of data over multiple nodes fails, then any partial errors might return "ok" on success or "error" on an internal uncategorized failure. Whenever a sample collection is missed but done at a later time, it is back filled to the previous 15 second timestamp and tagged with "backfilled_data". "inconsistent_delta_time" is encountered when the time between two collections is not the same for all nodes. Therefore, the aggregated value might be over or under inflated. "Negative_delta" is returned when an expected monotonically increasing value has decreased in value. "inconsistent_old_data" is returned when one or more nodes do not have the latest data.
	// Example: ok
	// Enum: [ok error partial_no_data partial_no_uuid partial_no_response partial_other_error negative_delta backfilled_data inconsistent_delta_time inconsistent_old_data]
	Status string `json:"status,omitempty"`

	// throughput
	Throughput *InterfaceMetricsThroughput `json:"throughput,omitempty"`

	// The timestamp of the performance data.
	// Example: 2017-01-25T11:20:13Z
	// Format: date-time
	Timestamp *strfmt.DateTime `json:"timestamp,omitempty"`

	// The UUID that uniquely identifies the interface.
	// Example: 1cd8a442-86d1-11e0-ae1c-123478563412
	// Read Only: true
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this interface metrics
func (m *InterfaceMetrics) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDuration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateThroughput(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InterfaceMetrics) validateLinks(formats strfmt.Registry) error {
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

var interfaceMetricsTypeDurationPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["PT15S","PT4M","PT30M","PT2H","P1D","PT5M"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		interfaceMetricsTypeDurationPropEnum = append(interfaceMetricsTypeDurationPropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// interface_metrics
	// InterfaceMetrics
	// duration
	// Duration
	// PT15S
	// END DEBUGGING
	// InterfaceMetricsDurationPT15S captures enum value "PT15S"
	InterfaceMetricsDurationPT15S string = "PT15S"

	// BEGIN DEBUGGING
	// interface_metrics
	// InterfaceMetrics
	// duration
	// Duration
	// PT4M
	// END DEBUGGING
	// InterfaceMetricsDurationPT4M captures enum value "PT4M"
	InterfaceMetricsDurationPT4M string = "PT4M"

	// BEGIN DEBUGGING
	// interface_metrics
	// InterfaceMetrics
	// duration
	// Duration
	// PT30M
	// END DEBUGGING
	// InterfaceMetricsDurationPT30M captures enum value "PT30M"
	InterfaceMetricsDurationPT30M string = "PT30M"

	// BEGIN DEBUGGING
	// interface_metrics
	// InterfaceMetrics
	// duration
	// Duration
	// PT2H
	// END DEBUGGING
	// InterfaceMetricsDurationPT2H captures enum value "PT2H"
	InterfaceMetricsDurationPT2H string = "PT2H"

	// BEGIN DEBUGGING
	// interface_metrics
	// InterfaceMetrics
	// duration
	// Duration
	// P1D
	// END DEBUGGING
	// InterfaceMetricsDurationP1D captures enum value "P1D"
	InterfaceMetricsDurationP1D string = "P1D"

	// BEGIN DEBUGGING
	// interface_metrics
	// InterfaceMetrics
	// duration
	// Duration
	// PT5M
	// END DEBUGGING
	// InterfaceMetricsDurationPT5M captures enum value "PT5M"
	InterfaceMetricsDurationPT5M string = "PT5M"
)

// prop value enum
func (m *InterfaceMetrics) validateDurationEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, interfaceMetricsTypeDurationPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *InterfaceMetrics) validateDuration(formats strfmt.Registry) error {
	if swag.IsZero(m.Duration) { // not required
		return nil
	}

	// value enum
	if err := m.validateDurationEnum("duration", "body", m.Duration); err != nil {
		return err
	}

	return nil
}

var interfaceMetricsTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ok","error","partial_no_data","partial_no_uuid","partial_no_response","partial_other_error","negative_delta","backfilled_data","inconsistent_delta_time","inconsistent_old_data"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		interfaceMetricsTypeStatusPropEnum = append(interfaceMetricsTypeStatusPropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// interface_metrics
	// InterfaceMetrics
	// status
	// Status
	// ok
	// END DEBUGGING
	// InterfaceMetricsStatusOk captures enum value "ok"
	InterfaceMetricsStatusOk string = "ok"

	// BEGIN DEBUGGING
	// interface_metrics
	// InterfaceMetrics
	// status
	// Status
	// error
	// END DEBUGGING
	// InterfaceMetricsStatusError captures enum value "error"
	InterfaceMetricsStatusError string = "error"

	// BEGIN DEBUGGING
	// interface_metrics
	// InterfaceMetrics
	// status
	// Status
	// partial_no_data
	// END DEBUGGING
	// InterfaceMetricsStatusPartialNoData captures enum value "partial_no_data"
	InterfaceMetricsStatusPartialNoData string = "partial_no_data"

	// BEGIN DEBUGGING
	// interface_metrics
	// InterfaceMetrics
	// status
	// Status
	// partial_no_uuid
	// END DEBUGGING
	// InterfaceMetricsStatusPartialNoUUID captures enum value "partial_no_uuid"
	InterfaceMetricsStatusPartialNoUUID string = "partial_no_uuid"

	// BEGIN DEBUGGING
	// interface_metrics
	// InterfaceMetrics
	// status
	// Status
	// partial_no_response
	// END DEBUGGING
	// InterfaceMetricsStatusPartialNoResponse captures enum value "partial_no_response"
	InterfaceMetricsStatusPartialNoResponse string = "partial_no_response"

	// BEGIN DEBUGGING
	// interface_metrics
	// InterfaceMetrics
	// status
	// Status
	// partial_other_error
	// END DEBUGGING
	// InterfaceMetricsStatusPartialOtherError captures enum value "partial_other_error"
	InterfaceMetricsStatusPartialOtherError string = "partial_other_error"

	// BEGIN DEBUGGING
	// interface_metrics
	// InterfaceMetrics
	// status
	// Status
	// negative_delta
	// END DEBUGGING
	// InterfaceMetricsStatusNegativeDelta captures enum value "negative_delta"
	InterfaceMetricsStatusNegativeDelta string = "negative_delta"

	// BEGIN DEBUGGING
	// interface_metrics
	// InterfaceMetrics
	// status
	// Status
	// backfilled_data
	// END DEBUGGING
	// InterfaceMetricsStatusBackfilledData captures enum value "backfilled_data"
	InterfaceMetricsStatusBackfilledData string = "backfilled_data"

	// BEGIN DEBUGGING
	// interface_metrics
	// InterfaceMetrics
	// status
	// Status
	// inconsistent_delta_time
	// END DEBUGGING
	// InterfaceMetricsStatusInconsistentDeltaTime captures enum value "inconsistent_delta_time"
	InterfaceMetricsStatusInconsistentDeltaTime string = "inconsistent_delta_time"

	// BEGIN DEBUGGING
	// interface_metrics
	// InterfaceMetrics
	// status
	// Status
	// inconsistent_old_data
	// END DEBUGGING
	// InterfaceMetricsStatusInconsistentOldData captures enum value "inconsistent_old_data"
	InterfaceMetricsStatusInconsistentOldData string = "inconsistent_old_data"
)

// prop value enum
func (m *InterfaceMetrics) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, interfaceMetricsTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *InterfaceMetrics) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *InterfaceMetrics) validateThroughput(formats strfmt.Registry) error {
	if swag.IsZero(m.Throughput) { // not required
		return nil
	}

	if m.Throughput != nil {
		if err := m.Throughput.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("throughput")
			}
			return err
		}
	}

	return nil
}

func (m *InterfaceMetrics) validateTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.Timestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("timestamp", "body", "date-time", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this interface metrics based on the context it is used
func (m *InterfaceMetrics) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateThroughput(ctx, formats); err != nil {
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

func (m *InterfaceMetrics) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *InterfaceMetrics) contextValidateThroughput(ctx context.Context, formats strfmt.Registry) error {

	if m.Throughput != nil {
		if err := m.Throughput.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("throughput")
			}
			return err
		}
	}

	return nil
}

func (m *InterfaceMetrics) contextValidateUUID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "uuid", "body", string(m.UUID)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InterfaceMetrics) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InterfaceMetrics) UnmarshalBinary(b []byte) error {
	var res InterfaceMetrics
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// InterfaceMetricsLinks interface metrics links
//
// swagger:model InterfaceMetricsLinks
type InterfaceMetricsLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this interface metrics links
func (m *InterfaceMetricsLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InterfaceMetricsLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this interface metrics links based on the context it is used
func (m *InterfaceMetricsLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InterfaceMetricsLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *InterfaceMetricsLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InterfaceMetricsLinks) UnmarshalBinary(b []byte) error {
	var res InterfaceMetricsLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// InterfaceMetricsThroughput The rate of throughput bytes per second observed at the interface.
//
// swagger:model InterfaceMetricsThroughput
type InterfaceMetricsThroughput struct {

	// Performance metric for read I/O operations.
	// Example: 200
	Read int64 `json:"read,omitempty"`

	// Performance metric aggregated over all types of I/O operations.
	// Example: 1000
	Total int64 `json:"total,omitempty"`

	// Peformance metric for write I/O operations.
	// Example: 100
	Write int64 `json:"write,omitempty"`
}

// Validate validates this interface metrics throughput
func (m *InterfaceMetricsThroughput) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this interface metrics throughput based on context it is used
func (m *InterfaceMetricsThroughput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *InterfaceMetricsThroughput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InterfaceMetricsThroughput) UnmarshalBinary(b []byte) error {
	var res InterfaceMetricsThroughput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
