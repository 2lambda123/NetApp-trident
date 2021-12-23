// Code generated by go-swagger; DO NOT EDIT.

package support

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// CoredumpDeleteReader is a Reader for the CoredumpDelete structure.
type CoredumpDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CoredumpDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCoredumpDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCoredumpDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCoredumpDeleteOK creates a CoredumpDeleteOK with default headers values
func NewCoredumpDeleteOK() *CoredumpDeleteOK {
	return &CoredumpDeleteOK{}
}

/* CoredumpDeleteOK describes a response with status code 200, with default header values.

OK
*/
type CoredumpDeleteOK struct {
}

func (o *CoredumpDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /support/coredump/coredumps/{node.uuid}/{name}][%d] coredumpDeleteOK ", 200)
}

func (o *CoredumpDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCoredumpDeleteDefault creates a CoredumpDeleteDefault with default headers values
func NewCoredumpDeleteDefault(code int) *CoredumpDeleteDefault {
	return &CoredumpDeleteDefault{
		_statusCode: code,
	}
}

/* CoredumpDeleteDefault describes a response with status code -1, with default header values.

Error
*/
type CoredumpDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the coredump delete default response
func (o *CoredumpDeleteDefault) Code() int {
	return o._statusCode
}

func (o *CoredumpDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /support/coredump/coredumps/{node.uuid}/{name}][%d] coredump_delete default  %+v", o._statusCode, o.Payload)
}
func (o *CoredumpDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CoredumpDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
