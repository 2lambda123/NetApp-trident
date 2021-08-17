// Code generated by go-swagger; DO NOT EDIT.

package security

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/v21/storage_drivers/ontap/api/rest/models"
)

// SecurityAuditModifyReader is a Reader for the SecurityAuditModify structure.
type SecurityAuditModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SecurityAuditModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewSecurityAuditModifyAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSecurityAuditModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSecurityAuditModifyAccepted creates a SecurityAuditModifyAccepted with default headers values
func NewSecurityAuditModifyAccepted() *SecurityAuditModifyAccepted {
	return &SecurityAuditModifyAccepted{}
}

/* SecurityAuditModifyAccepted describes a response with status code 202, with default header values.

Accepted
*/
type SecurityAuditModifyAccepted struct {
	Payload *models.SecurityAudit
}

func (o *SecurityAuditModifyAccepted) Error() string {
	return fmt.Sprintf("[PATCH /security/audit][%d] securityAuditModifyAccepted  %+v", 202, o.Payload)
}
func (o *SecurityAuditModifyAccepted) GetPayload() *models.SecurityAudit {
	return o.Payload
}

func (o *SecurityAuditModifyAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SecurityAudit)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSecurityAuditModifyDefault creates a SecurityAuditModifyDefault with default headers values
func NewSecurityAuditModifyDefault(code int) *SecurityAuditModifyDefault {
	return &SecurityAuditModifyDefault{
		_statusCode: code,
	}
}

/* SecurityAuditModifyDefault describes a response with status code -1, with default header values.

Error
*/
type SecurityAuditModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the security audit modify default response
func (o *SecurityAuditModifyDefault) Code() int {
	return o._statusCode
}

func (o *SecurityAuditModifyDefault) Error() string {
	return fmt.Sprintf("[PATCH /security/audit][%d] security_audit_modify default  %+v", o._statusCode, o.Payload)
}
func (o *SecurityAuditModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SecurityAuditModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
