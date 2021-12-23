// Code generated by go-swagger; DO NOT EDIT.

package security

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// IpsecCaCertificateDeleteReader is a Reader for the IpsecCaCertificateDelete structure.
type IpsecCaCertificateDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpsecCaCertificateDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpsecCaCertificateDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIpsecCaCertificateDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpsecCaCertificateDeleteOK creates a IpsecCaCertificateDeleteOK with default headers values
func NewIpsecCaCertificateDeleteOK() *IpsecCaCertificateDeleteOK {
	return &IpsecCaCertificateDeleteOK{}
}

/* IpsecCaCertificateDeleteOK describes a response with status code 200, with default header values.

OK
*/
type IpsecCaCertificateDeleteOK struct {
}

func (o *IpsecCaCertificateDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /security/ipsec/ca-certificates/{certificate.uuid}][%d] ipsecCaCertificateDeleteOK ", 200)
}

func (o *IpsecCaCertificateDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewIpsecCaCertificateDeleteDefault creates a IpsecCaCertificateDeleteDefault with default headers values
func NewIpsecCaCertificateDeleteDefault(code int) *IpsecCaCertificateDeleteDefault {
	return &IpsecCaCertificateDeleteDefault{
		_statusCode: code,
	}
}

/* IpsecCaCertificateDeleteDefault describes a response with status code -1, with default header values.

Error
*/
type IpsecCaCertificateDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the ipsec ca certificate delete default response
func (o *IpsecCaCertificateDeleteDefault) Code() int {
	return o._statusCode
}

func (o *IpsecCaCertificateDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /security/ipsec/ca-certificates/{certificate.uuid}][%d] ipsec_ca_certificate_delete default  %+v", o._statusCode, o.Payload)
}
func (o *IpsecCaCertificateDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *IpsecCaCertificateDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
