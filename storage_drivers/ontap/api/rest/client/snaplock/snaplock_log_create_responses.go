// Code generated by go-swagger; DO NOT EDIT.

package snaplock

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/v21/storage_drivers/ontap/api/rest/models"
)

// SnaplockLogCreateReader is a Reader for the SnaplockLogCreate structure.
type SnaplockLogCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SnaplockLogCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewSnaplockLogCreateAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSnaplockLogCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSnaplockLogCreateAccepted creates a SnaplockLogCreateAccepted with default headers values
func NewSnaplockLogCreateAccepted() *SnaplockLogCreateAccepted {
	return &SnaplockLogCreateAccepted{}
}

/* SnaplockLogCreateAccepted describes a response with status code 202, with default header values.

Accepted
*/
type SnaplockLogCreateAccepted struct {
	Payload *models.JobLinkResponse
}

func (o *SnaplockLogCreateAccepted) Error() string {
	return fmt.Sprintf("[POST /storage/snaplock/audit-logs][%d] snaplockLogCreateAccepted  %+v", 202, o.Payload)
}
func (o *SnaplockLogCreateAccepted) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *SnaplockLogCreateAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSnaplockLogCreateDefault creates a SnaplockLogCreateDefault with default headers values
func NewSnaplockLogCreateDefault(code int) *SnaplockLogCreateDefault {
	return &SnaplockLogCreateDefault{
		_statusCode: code,
	}
}

/* SnaplockLogCreateDefault describes a response with status code -1, with default header values.

 ONTAP Error Response codes
| Error code  |  Description |
|-------------|--------------|
| 14090340    | {field} is a required field  |
| 14090343    | Invalid Field  |
| 14090346    | Internal Error. Wait a few minutes, then try the command again  |

*/
type SnaplockLogCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the snaplock log create default response
func (o *SnaplockLogCreateDefault) Code() int {
	return o._statusCode
}

func (o *SnaplockLogCreateDefault) Error() string {
	return fmt.Sprintf("[POST /storage/snaplock/audit-logs][%d] snaplock_log_create default  %+v", o._statusCode, o.Payload)
}
func (o *SnaplockLogCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SnaplockLogCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
