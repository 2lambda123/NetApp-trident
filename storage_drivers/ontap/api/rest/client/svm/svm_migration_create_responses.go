// Code generated by go-swagger; DO NOT EDIT.

package svm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// SvmMigrationCreateReader is a Reader for the SvmMigrationCreate structure.
type SvmMigrationCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SvmMigrationCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewSvmMigrationCreateAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSvmMigrationCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSvmMigrationCreateAccepted creates a SvmMigrationCreateAccepted with default headers values
func NewSvmMigrationCreateAccepted() *SvmMigrationCreateAccepted {
	return &SvmMigrationCreateAccepted{}
}

/* SvmMigrationCreateAccepted describes a response with status code 202, with default header values.

Accepted
*/
type SvmMigrationCreateAccepted struct {
	Payload *models.JobLinkResponse
}

func (o *SvmMigrationCreateAccepted) Error() string {
	return fmt.Sprintf("[POST /svm/migrations][%d] svmMigrationCreateAccepted  %+v", 202, o.Payload)
}
func (o *SvmMigrationCreateAccepted) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *SvmMigrationCreateAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSvmMigrationCreateDefault creates a SvmMigrationCreateDefault with default headers values
func NewSvmMigrationCreateDefault(code int) *SvmMigrationCreateDefault {
	return &SvmMigrationCreateDefault{
		_statusCode: code,
	}
}

/* SvmMigrationCreateDefault describes a response with status code -1, with default header values.

 ONTAP Error Response Codes
| Error Code | Description |
| ---------- | ----------- |
| 262245 | The value provided is invalid. |
| 13172746 | SVM migration cannot be started. This is a generic code, see the response message for details. |

*/
type SvmMigrationCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the svm migration create default response
func (o *SvmMigrationCreateDefault) Code() int {
	return o._statusCode
}

func (o *SvmMigrationCreateDefault) Error() string {
	return fmt.Sprintf("[POST /svm/migrations][%d] svm_migration_create default  %+v", o._statusCode, o.Payload)
}
func (o *SvmMigrationCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SvmMigrationCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
