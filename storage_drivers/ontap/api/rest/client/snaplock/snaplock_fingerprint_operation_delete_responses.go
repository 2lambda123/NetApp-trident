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

// SnaplockFingerprintOperationDeleteReader is a Reader for the SnaplockFingerprintOperationDelete structure.
type SnaplockFingerprintOperationDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SnaplockFingerprintOperationDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSnaplockFingerprintOperationDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSnaplockFingerprintOperationDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSnaplockFingerprintOperationDeleteOK creates a SnaplockFingerprintOperationDeleteOK with default headers values
func NewSnaplockFingerprintOperationDeleteOK() *SnaplockFingerprintOperationDeleteOK {
	return &SnaplockFingerprintOperationDeleteOK{}
}

/* SnaplockFingerprintOperationDeleteOK describes a response with status code 200, with default header values.

OK
*/
type SnaplockFingerprintOperationDeleteOK struct {
}

func (o *SnaplockFingerprintOperationDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /storage/snaplock/file-fingerprints/{id}][%d] snaplockFingerprintOperationDeleteOK ", 200)
}

func (o *SnaplockFingerprintOperationDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSnaplockFingerprintOperationDeleteDefault creates a SnaplockFingerprintOperationDeleteDefault with default headers values
func NewSnaplockFingerprintOperationDeleteDefault(code int) *SnaplockFingerprintOperationDeleteDefault {
	return &SnaplockFingerprintOperationDeleteDefault{
		_statusCode: code,
	}
}

/* SnaplockFingerprintOperationDeleteDefault describes a response with status code -1, with default header values.

 ONTAP Error Response codes
| Error code  |  Description |
|-------------|--------------|
| 14090440    | File fingerprint operation has completed  |
| 14090446    | Invalid session ID  |

*/
type SnaplockFingerprintOperationDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the snaplock fingerprint operation delete default response
func (o *SnaplockFingerprintOperationDeleteDefault) Code() int {
	return o._statusCode
}

func (o *SnaplockFingerprintOperationDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /storage/snaplock/file-fingerprints/{id}][%d] snaplock_fingerprint_operation_delete default  %+v", o._statusCode, o.Payload)
}
func (o *SnaplockFingerprintOperationDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SnaplockFingerprintOperationDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
