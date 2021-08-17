// Code generated by go-swagger; DO NOT EDIT.

package n_a_s

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/v21/storage_drivers/ontap/api/rest/models"
)

// KerberosInterfaceGetReader is a Reader for the KerberosInterfaceGet structure.
type KerberosInterfaceGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KerberosInterfaceGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewKerberosInterfaceGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewKerberosInterfaceGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewKerberosInterfaceGetOK creates a KerberosInterfaceGetOK with default headers values
func NewKerberosInterfaceGetOK() *KerberosInterfaceGetOK {
	return &KerberosInterfaceGetOK{}
}

/* KerberosInterfaceGetOK describes a response with status code 200, with default header values.

OK
*/
type KerberosInterfaceGetOK struct {
	Payload *models.KerberosInterface
}

func (o *KerberosInterfaceGetOK) Error() string {
	return fmt.Sprintf("[GET /protocols/nfs/kerberos/interfaces/{interface.uuid}][%d] kerberosInterfaceGetOK  %+v", 200, o.Payload)
}
func (o *KerberosInterfaceGetOK) GetPayload() *models.KerberosInterface {
	return o.Payload
}

func (o *KerberosInterfaceGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.KerberosInterface)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKerberosInterfaceGetDefault creates a KerberosInterfaceGetDefault with default headers values
func NewKerberosInterfaceGetDefault(code int) *KerberosInterfaceGetDefault {
	return &KerberosInterfaceGetDefault{
		_statusCode: code,
	}
}

/* KerberosInterfaceGetDefault describes a response with status code -1, with default header values.

Error
*/
type KerberosInterfaceGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the kerberos interface get default response
func (o *KerberosInterfaceGetDefault) Code() int {
	return o._statusCode
}

func (o *KerberosInterfaceGetDefault) Error() string {
	return fmt.Sprintf("[GET /protocols/nfs/kerberos/interfaces/{interface.uuid}][%d] kerberos_interface_get default  %+v", o._statusCode, o.Payload)
}
func (o *KerberosInterfaceGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *KerberosInterfaceGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
