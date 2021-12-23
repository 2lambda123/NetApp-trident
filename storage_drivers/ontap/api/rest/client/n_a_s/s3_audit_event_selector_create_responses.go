// Code generated by go-swagger; DO NOT EDIT.

package n_a_s

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// S3AuditEventSelectorCreateReader is a Reader for the S3AuditEventSelectorCreate structure.
type S3AuditEventSelectorCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *S3AuditEventSelectorCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewS3AuditEventSelectorCreateAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewS3AuditEventSelectorCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewS3AuditEventSelectorCreateAccepted creates a S3AuditEventSelectorCreateAccepted with default headers values
func NewS3AuditEventSelectorCreateAccepted() *S3AuditEventSelectorCreateAccepted {
	return &S3AuditEventSelectorCreateAccepted{}
}

/* S3AuditEventSelectorCreateAccepted describes a response with status code 202, with default header values.

Accepted
*/
type S3AuditEventSelectorCreateAccepted struct {
	Payload *models.S3AuditEventSelector
}

func (o *S3AuditEventSelectorCreateAccepted) Error() string {
	return fmt.Sprintf("[POST /protocols/event-selectors][%d] s3AuditEventSelectorCreateAccepted  %+v", 202, o.Payload)
}
func (o *S3AuditEventSelectorCreateAccepted) GetPayload() *models.S3AuditEventSelector {
	return o.Payload
}

func (o *S3AuditEventSelectorCreateAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.S3AuditEventSelector)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewS3AuditEventSelectorCreateDefault creates a S3AuditEventSelectorCreateDefault with default headers values
func NewS3AuditEventSelectorCreateDefault(code int) *S3AuditEventSelectorCreateDefault {
	return &S3AuditEventSelectorCreateDefault{
		_statusCode: code,
	}
}

/* S3AuditEventSelectorCreateDefault describes a response with status code -1, with default header values.

 ONTAP Error Response Codes
| Error Code | Description |
| ---------- | ----------- |
| 140902570 | S3 audit event selector configuration is not available. |
| 140902571 | S3 audit configuration was not created for the SVM. |
| 140902572 | S3 audit event selector operation failed. |
| 140902573 | Not all S3 audit event selectors have been removed for the SVM. |

*/
type S3AuditEventSelectorCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the s3 audit event selector create default response
func (o *S3AuditEventSelectorCreateDefault) Code() int {
	return o._statusCode
}

func (o *S3AuditEventSelectorCreateDefault) Error() string {
	return fmt.Sprintf("[POST /protocols/event-selectors][%d] s3_audit_event_selector_create default  %+v", o._statusCode, o.Payload)
}
func (o *S3AuditEventSelectorCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *S3AuditEventSelectorCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
