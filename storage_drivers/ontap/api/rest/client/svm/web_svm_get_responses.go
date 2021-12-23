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

// WebSvmGetReader is a Reader for the WebSvmGet structure.
type WebSvmGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WebSvmGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWebSvmGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWebSvmGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWebSvmGetOK creates a WebSvmGetOK with default headers values
func NewWebSvmGetOK() *WebSvmGetOK {
	return &WebSvmGetOK{}
}

/* WebSvmGetOK describes a response with status code 200, with default header values.

OK
*/
type WebSvmGetOK struct {
	Payload *models.WebSvm
}

func (o *WebSvmGetOK) Error() string {
	return fmt.Sprintf("[GET /svm/svms/{svm.uuid}/web][%d] webSvmGetOK  %+v", 200, o.Payload)
}
func (o *WebSvmGetOK) GetPayload() *models.WebSvm {
	return o.Payload
}

func (o *WebSvmGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WebSvm)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWebSvmGetDefault creates a WebSvmGetDefault with default headers values
func NewWebSvmGetDefault(code int) *WebSvmGetDefault {
	return &WebSvmGetDefault{
		_statusCode: code,
	}
}

/* WebSvmGetDefault describes a response with status code -1, with default header values.

Error
*/
type WebSvmGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the web svm get default response
func (o *WebSvmGetDefault) Code() int {
	return o._statusCode
}

func (o *WebSvmGetDefault) Error() string {
	return fmt.Sprintf("[GET /svm/svms/{svm.uuid}/web][%d] web_svm_get default  %+v", o._statusCode, o.Payload)
}
func (o *WebSvmGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *WebSvmGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
