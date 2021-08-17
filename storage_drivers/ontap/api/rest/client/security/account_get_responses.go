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

// AccountGetReader is a Reader for the AccountGet structure.
type AccountGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AccountGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAccountGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAccountGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAccountGetOK creates a AccountGetOK with default headers values
func NewAccountGetOK() *AccountGetOK {
	return &AccountGetOK{}
}

/* AccountGetOK describes a response with status code 200, with default header values.

OK
*/
type AccountGetOK struct {
	Payload *models.AccountResponse
}

func (o *AccountGetOK) Error() string {
	return fmt.Sprintf("[GET /security/accounts/{owner.uuid}/{name}][%d] accountGetOK  %+v", 200, o.Payload)
}
func (o *AccountGetOK) GetPayload() *models.AccountResponse {
	return o.Payload
}

func (o *AccountGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AccountResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAccountGetDefault creates a AccountGetDefault with default headers values
func NewAccountGetDefault(code int) *AccountGetDefault {
	return &AccountGetDefault{
		_statusCode: code,
	}
}

/* AccountGetDefault describes a response with status code -1, with default header values.

Error
*/
type AccountGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the account get default response
func (o *AccountGetDefault) Code() int {
	return o._statusCode
}

func (o *AccountGetDefault) Error() string {
	return fmt.Sprintf("[GET /security/accounts/{owner.uuid}/{name}][%d] account_get default  %+v", o._statusCode, o.Payload)
}
func (o *AccountGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AccountGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
