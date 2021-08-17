// Code generated by go-swagger; DO NOT EDIT.

package security

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/netapp/trident/v21/storage_drivers/ontap/api/rest/models"
)

// SecurityKeyManagerMigrateReader is a Reader for the SecurityKeyManagerMigrate structure.
type SecurityKeyManagerMigrateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SecurityKeyManagerMigrateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewSecurityKeyManagerMigrateAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSecurityKeyManagerMigrateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSecurityKeyManagerMigrateAccepted creates a SecurityKeyManagerMigrateAccepted with default headers values
func NewSecurityKeyManagerMigrateAccepted() *SecurityKeyManagerMigrateAccepted {
	return &SecurityKeyManagerMigrateAccepted{}
}

/* SecurityKeyManagerMigrateAccepted describes a response with status code 202, with default header values.

Accepted
*/
type SecurityKeyManagerMigrateAccepted struct {
	Payload *models.JobLinkResponse
}

func (o *SecurityKeyManagerMigrateAccepted) Error() string {
	return fmt.Sprintf("[POST /security/key-managers/{source.uuid}/migrate][%d] securityKeyManagerMigrateAccepted  %+v", 202, o.Payload)
}
func (o *SecurityKeyManagerMigrateAccepted) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *SecurityKeyManagerMigrateAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSecurityKeyManagerMigrateDefault creates a SecurityKeyManagerMigrateDefault with default headers values
func NewSecurityKeyManagerMigrateDefault(code int) *SecurityKeyManagerMigrateDefault {
	return &SecurityKeyManagerMigrateDefault{
		_statusCode: code,
	}
}

/* SecurityKeyManagerMigrateDefault describes a response with status code -1, with default header values.

Error
*/
type SecurityKeyManagerMigrateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the security key manager migrate default response
func (o *SecurityKeyManagerMigrateDefault) Code() int {
	return o._statusCode
}

func (o *SecurityKeyManagerMigrateDefault) Error() string {
	return fmt.Sprintf("[POST /security/key-managers/{source.uuid}/migrate][%d] security_key_manager_migrate default  %+v", o._statusCode, o.Payload)
}
func (o *SecurityKeyManagerMigrateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SecurityKeyManagerMigrateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*SecurityKeyManagerMigrateBody Migration destination key manager UUID
swagger:model SecurityKeyManagerMigrateBody
*/
type SecurityKeyManagerMigrateBody struct {

	// links
	Links *SecurityKeyManagerMigrateParamsBodyLinks `json:"_links,omitempty"`

	// Key manager UUID
	// Example: 1cd8a442-86d1-11e0-ae1c-123478563434
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this security key manager migrate body
func (o *SecurityKeyManagerMigrateBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SecurityKeyManagerMigrateBody) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(o.Links) { // not required
		return nil
	}

	if o.Links != nil {
		if err := o.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this security key manager migrate body based on the context it is used
func (o *SecurityKeyManagerMigrateBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SecurityKeyManagerMigrateBody) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if o.Links != nil {
		if err := o.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *SecurityKeyManagerMigrateBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SecurityKeyManagerMigrateBody) UnmarshalBinary(b []byte) error {
	var res SecurityKeyManagerMigrateBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*SecurityKeyManagerMigrateParamsBodyLinks security key manager migrate params body links
swagger:model SecurityKeyManagerMigrateParamsBodyLinks
*/
type SecurityKeyManagerMigrateParamsBodyLinks struct {

	// self
	Self *models.Href `json:"self,omitempty"`
}

// Validate validates this security key manager migrate params body links
func (o *SecurityKeyManagerMigrateParamsBodyLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SecurityKeyManagerMigrateParamsBodyLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(o.Self) { // not required
		return nil
	}

	if o.Self != nil {
		if err := o.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this security key manager migrate params body links based on the context it is used
func (o *SecurityKeyManagerMigrateParamsBodyLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SecurityKeyManagerMigrateParamsBodyLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if o.Self != nil {
		if err := o.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *SecurityKeyManagerMigrateParamsBodyLinks) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SecurityKeyManagerMigrateParamsBodyLinks) UnmarshalBinary(b []byte) error {
	var res SecurityKeyManagerMigrateParamsBodyLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
