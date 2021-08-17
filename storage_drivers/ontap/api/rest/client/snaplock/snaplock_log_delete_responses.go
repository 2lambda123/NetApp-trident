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

// SnaplockLogDeleteReader is a Reader for the SnaplockLogDelete structure.
type SnaplockLogDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SnaplockLogDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewSnaplockLogDeleteAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSnaplockLogDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSnaplockLogDeleteAccepted creates a SnaplockLogDeleteAccepted with default headers values
func NewSnaplockLogDeleteAccepted() *SnaplockLogDeleteAccepted {
	return &SnaplockLogDeleteAccepted{}
}

/* SnaplockLogDeleteAccepted describes a response with status code 202, with default header values.

Accepted
*/
type SnaplockLogDeleteAccepted struct {
	Payload *models.JobLinkResponse
}

func (o *SnaplockLogDeleteAccepted) Error() string {
	return fmt.Sprintf("[DELETE /storage/snaplock/audit-logs/{svm.uuid}][%d] snaplockLogDeleteAccepted  %+v", 202, o.Payload)
}
func (o *SnaplockLogDeleteAccepted) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *SnaplockLogDeleteAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSnaplockLogDeleteDefault creates a SnaplockLogDeleteDefault with default headers values
func NewSnaplockLogDeleteDefault(code int) *SnaplockLogDeleteDefault {
	return &SnaplockLogDeleteDefault{
		_statusCode: code,
	}
}

/* SnaplockLogDeleteDefault describes a response with status code -1, with default header values.

 ONTAP Error Response codes
| Error code  |  Description |
|-------------|--------------|
| 14090346    | Internal Error. Wait a few minutes, then try the command again  |

*/
type SnaplockLogDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the snaplock log delete default response
func (o *SnaplockLogDeleteDefault) Code() int {
	return o._statusCode
}

func (o *SnaplockLogDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /storage/snaplock/audit-logs/{svm.uuid}][%d] snaplock_log_delete default  %+v", o._statusCode, o.Payload)
}
func (o *SnaplockLogDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SnaplockLogDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
