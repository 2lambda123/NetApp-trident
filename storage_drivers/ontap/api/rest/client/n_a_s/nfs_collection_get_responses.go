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

// NfsCollectionGetReader is a Reader for the NfsCollectionGet structure.
type NfsCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NfsCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNfsCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewNfsCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNfsCollectionGetOK creates a NfsCollectionGetOK with default headers values
func NewNfsCollectionGetOK() *NfsCollectionGetOK {
	return &NfsCollectionGetOK{}
}

/* NfsCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type NfsCollectionGetOK struct {
	Payload *models.NfsServiceResponse
}

func (o *NfsCollectionGetOK) Error() string {
	return fmt.Sprintf("[GET /protocols/nfs/services][%d] nfsCollectionGetOK  %+v", 200, o.Payload)
}
func (o *NfsCollectionGetOK) GetPayload() *models.NfsServiceResponse {
	return o.Payload
}

func (o *NfsCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NfsServiceResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNfsCollectionGetDefault creates a NfsCollectionGetDefault with default headers values
func NewNfsCollectionGetDefault(code int) *NfsCollectionGetDefault {
	return &NfsCollectionGetDefault{
		_statusCode: code,
	}
}

/* NfsCollectionGetDefault describes a response with status code -1, with default header values.

Error
*/
type NfsCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the nfs collection get default response
func (o *NfsCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *NfsCollectionGetDefault) Error() string {
	return fmt.Sprintf("[GET /protocols/nfs/services][%d] nfs_collection_get default  %+v", o._statusCode, o.Payload)
}
func (o *NfsCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *NfsCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
