// Code generated by go-swagger; DO NOT EDIT.

package support

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/v21/storage_drivers/ontap/api/rest/models"
)

// SnmpTraphostsGetReader is a Reader for the SnmpTraphostsGet structure.
type SnmpTraphostsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SnmpTraphostsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSnmpTraphostsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSnmpTraphostsGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSnmpTraphostsGetOK creates a SnmpTraphostsGetOK with default headers values
func NewSnmpTraphostsGetOK() *SnmpTraphostsGetOK {
	return &SnmpTraphostsGetOK{}
}

/* SnmpTraphostsGetOK describes a response with status code 200, with default header values.

OK
*/
type SnmpTraphostsGetOK struct {
	Payload *models.SnmpTraphost
}

func (o *SnmpTraphostsGetOK) Error() string {
	return fmt.Sprintf("[GET /support/snmp/traphosts/{host}][%d] snmpTraphostsGetOK  %+v", 200, o.Payload)
}
func (o *SnmpTraphostsGetOK) GetPayload() *models.SnmpTraphost {
	return o.Payload
}

func (o *SnmpTraphostsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SnmpTraphost)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSnmpTraphostsGetDefault creates a SnmpTraphostsGetDefault with default headers values
func NewSnmpTraphostsGetDefault(code int) *SnmpTraphostsGetDefault {
	return &SnmpTraphostsGetDefault{
		_statusCode: code,
	}
}

/* SnmpTraphostsGetDefault describes a response with status code -1, with default header values.

Error
*/
type SnmpTraphostsGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the snmp traphosts get default response
func (o *SnmpTraphostsGetDefault) Code() int {
	return o._statusCode
}

func (o *SnmpTraphostsGetDefault) Error() string {
	return fmt.Sprintf("[GET /support/snmp/traphosts/{host}][%d] snmp_traphosts_get default  %+v", o._statusCode, o.Payload)
}
func (o *SnmpTraphostsGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SnmpTraphostsGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
