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

// PortMetricsData Throughput performance for the Ethernet port.
//
// swagger:model port_metrics_data
type PortMetricsData struct {

	// links
	Links *PortMetricsDataLinks `json:"_links,omitempty"`

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
	Throughput *PortMetricsDataThroughput `json:"throughput,omitempty"`

	// The timestamp of the performance data.
	// Example: 2017-01-25T11:20:13Z
	// Format: date-time
	Timestamp *strfmt.DateTime `json:"timestamp,omitempty"`
}

// Validate validates this port metrics data
func (m *PortMetricsData) Validate(formats strfmt.Registry) error {
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

func (m *PortMetricsData) validateLinks(formats strfmt.Registry) error {
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

var portMetricsDataTypeDurationPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["PT15S","PT4M","PT30M","PT2H","P1D","PT5M"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		portMetricsDataTypeDurationPropEnum = append(portMetricsDataTypeDurationPropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// port_metrics_data
	// PortMetricsData
	// duration
	// Duration
	// PT15S
	// END DEBUGGING
	// PortMetricsDataDurationPT15S captures enum value "PT15S"
	PortMetricsDataDurationPT15S string = "PT15S"

	// BEGIN DEBUGGING
	// port_metrics_data
	// PortMetricsData
	// duration
	// Duration
	// PT4M
	// END DEBUGGING
	// PortMetricsDataDurationPT4M captures enum value "PT4M"
	PortMetricsDataDurationPT4M string = "PT4M"

	// BEGIN DEBUGGING
	// port_metrics_data
	// PortMetricsData
	// duration
	// Duration
	// PT30M
	// END DEBUGGING
	// PortMetricsDataDurationPT30M captures enum value "PT30M"
	PortMetricsDataDurationPT30M string = "PT30M"

	// BEGIN DEBUGGING
	// port_metrics_data
	// PortMetricsData
	// duration
	// Duration
	// PT2H
	// END DEBUGGING
	// PortMetricsDataDurationPT2H captures enum value "PT2H"
	PortMetricsDataDurationPT2H string = "PT2H"

	// BEGIN DEBUGGING
	// port_metrics_data
	// PortMetricsData
	// duration
	// Duration
	// P1D
	// END DEBUGGING
	// PortMetricsDataDurationP1D captures enum value "P1D"
	PortMetricsDataDurationP1D string = "P1D"

	// BEGIN DEBUGGING
	// port_metrics_data
	// PortMetricsData
	// duration
	// Duration
	// PT5M
	// END DEBUGGING
	// PortMetricsDataDurationPT5M captures enum value "PT5M"
	PortMetricsDataDurationPT5M string = "PT5M"
)

// prop value enum
func (m *PortMetricsData) validateDurationEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, portMetricsDataTypeDurationPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PortMetricsData) validateDuration(formats strfmt.Registry) error {
	if swag.IsZero(m.Duration) { // not required
		return nil
	}

	// value enum
	if err := m.validateDurationEnum("duration", "body", m.Duration); err != nil {
		return err
	}

	return nil
}

var portMetricsDataTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ok","error","partial_no_data","partial_no_uuid","partial_no_response","partial_other_error","negative_delta","backfilled_data","inconsistent_delta_time","inconsistent_old_data"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		portMetricsDataTypeStatusPropEnum = append(portMetricsDataTypeStatusPropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// port_metrics_data
	// PortMetricsData
	// status
	// Status
	// ok
	// END DEBUGGING
	// PortMetricsDataStatusOk captures enum value "ok"
	PortMetricsDataStatusOk string = "ok"

	// BEGIN DEBUGGING
	// port_metrics_data
	// PortMetricsData
	// status
	// Status
	// error
	// END DEBUGGING
	// PortMetricsDataStatusError captures enum value "error"
	PortMetricsDataStatusError string = "error"

	// BEGIN DEBUGGING
	// port_metrics_data
	// PortMetricsData
	// status
	// Status
	// partial_no_data
	// END DEBUGGING
	// PortMetricsDataStatusPartialNoData captures enum value "partial_no_data"
	PortMetricsDataStatusPartialNoData string = "partial_no_data"

	// BEGIN DEBUGGING
	// port_metrics_data
	// PortMetricsData
	// status
	// Status
	// partial_no_uuid
	// END DEBUGGING
	// PortMetricsDataStatusPartialNoUUID captures enum value "partial_no_uuid"
	PortMetricsDataStatusPartialNoUUID string = "partial_no_uuid"

	// BEGIN DEBUGGING
	// port_metrics_data
	// PortMetricsData
	// status
	// Status
	// partial_no_response
	// END DEBUGGING
	// PortMetricsDataStatusPartialNoResponse captures enum value "partial_no_response"
	PortMetricsDataStatusPartialNoResponse string = "partial_no_response"

	// BEGIN DEBUGGING
	// port_metrics_data
	// PortMetricsData
	// status
	// Status
	// partial_other_error
	// END DEBUGGING
	// PortMetricsDataStatusPartialOtherError captures enum value "partial_other_error"
	PortMetricsDataStatusPartialOtherError string = "partial_other_error"

	// BEGIN DEBUGGING
	// port_metrics_data
	// PortMetricsData
	// status
	// Status
	// negative_delta
	// END DEBUGGING
	// PortMetricsDataStatusNegativeDelta captures enum value "negative_delta"
	PortMetricsDataStatusNegativeDelta string = "negative_delta"

	// BEGIN DEBUGGING
	// port_metrics_data
	// PortMetricsData
	// status
	// Status
	// backfilled_data
	// END DEBUGGING
	// PortMetricsDataStatusBackfilledData captures enum value "backfilled_data"
	PortMetricsDataStatusBackfilledData string = "backfilled_data"

	// BEGIN DEBUGGING
	// port_metrics_data
	// PortMetricsData
	// status
	// Status
	// inconsistent_delta_time
	// END DEBUGGING
	// PortMetricsDataStatusInconsistentDeltaTime captures enum value "inconsistent_delta_time"
	PortMetricsDataStatusInconsistentDeltaTime string = "inconsistent_delta_time"

	// BEGIN DEBUGGING
	// port_metrics_data
	// PortMetricsData
	// status
	// Status
	// inconsistent_old_data
	// END DEBUGGING
	// PortMetricsDataStatusInconsistentOldData captures enum value "inconsistent_old_data"
	PortMetricsDataStatusInconsistentOldData string = "inconsistent_old_data"
)

// prop value enum
func (m *PortMetricsData) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, portMetricsDataTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PortMetricsData) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *PortMetricsData) validateThroughput(formats strfmt.Registry) error {
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

func (m *PortMetricsData) validateTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.Timestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("timestamp", "body", "date-time", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this port metrics data based on the context it is used
func (m *PortMetricsData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateThroughput(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PortMetricsData) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *PortMetricsData) contextValidateThroughput(ctx context.Context, formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *PortMetricsData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PortMetricsData) UnmarshalBinary(b []byte) error {
	var res PortMetricsData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PortMetricsDataLinks port metrics data links
//
// swagger:model PortMetricsDataLinks
type PortMetricsDataLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this port metrics data links
func (m *PortMetricsDataLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PortMetricsDataLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this port metrics data links based on the context it is used
func (m *PortMetricsDataLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PortMetricsDataLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *PortMetricsDataLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PortMetricsDataLinks) UnmarshalBinary(b []byte) error {
	var res PortMetricsDataLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PortMetricsDataThroughput The rate of throughput bytes per second observed at the interface.
//
// swagger:model PortMetricsDataThroughput
type PortMetricsDataThroughput struct {

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

// Validate validates this port metrics data throughput
func (m *PortMetricsDataThroughput) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this port metrics data throughput based on context it is used
func (m *PortMetricsDataThroughput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PortMetricsDataThroughput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PortMetricsDataThroughput) UnmarshalBinary(b []byte) error {
	var res PortMetricsDataThroughput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
