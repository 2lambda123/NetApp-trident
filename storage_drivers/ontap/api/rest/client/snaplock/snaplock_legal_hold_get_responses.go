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

// SnaplockLegalHoldGetReader is a Reader for the SnaplockLegalHoldGet structure.
type SnaplockLegalHoldGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SnaplockLegalHoldGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSnaplockLegalHoldGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSnaplockLegalHoldGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSnaplockLegalHoldGetOK creates a SnaplockLegalHoldGetOK with default headers values
func NewSnaplockLegalHoldGetOK() *SnaplockLegalHoldGetOK {
	return &SnaplockLegalHoldGetOK{}
}

/* SnaplockLegalHoldGetOK describes a response with status code 200, with default header values.

OK
*/
type SnaplockLegalHoldGetOK struct {
	Payload *models.SnaplockLitigation
}

func (o *SnaplockLegalHoldGetOK) Error() string {
	return fmt.Sprintf("[GET /storage/snaplock/litigations/{id}][%d] snaplockLegalHoldGetOK  %+v", 200, o.Payload)
}
func (o *SnaplockLegalHoldGetOK) GetPayload() *models.SnaplockLitigation {
	return o.Payload
}

func (o *SnaplockLegalHoldGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SnaplockLitigation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSnaplockLegalHoldGetDefault creates a SnaplockLegalHoldGetDefault with default headers values
func NewSnaplockLegalHoldGetDefault(code int) *SnaplockLegalHoldGetDefault {
	return &SnaplockLegalHoldGetDefault{
		_statusCode: code,
	}
}

/* SnaplockLegalHoldGetDefault describes a response with status code -1, with default header values.

 ONTAP Error Response codes
| Error code  |  Description |
|-------------|--------------|
| 14090346    | Internal Error. Wait a few minutes, then try the command again  |
| 14090343    | Invalid Field  |

*/
type SnaplockLegalHoldGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the snaplock legal hold get default response
func (o *SnaplockLegalHoldGetDefault) Code() int {
	return o._statusCode
}

func (o *SnaplockLegalHoldGetDefault) Error() string {
	return fmt.Sprintf("[GET /storage/snaplock/litigations/{id}][%d] snaplock_legal_hold_get default  %+v", o._statusCode, o.Payload)
}
func (o *SnaplockLegalHoldGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SnaplockLegalHoldGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
