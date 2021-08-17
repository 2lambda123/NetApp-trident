// Code generated by go-swagger; DO NOT EDIT.

package s_a_n

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/v21/storage_drivers/ontap/api/rest/models"
)

// LunMapCollectionGetReader is a Reader for the LunMapCollectionGet structure.
type LunMapCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LunMapCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLunMapCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewLunMapCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewLunMapCollectionGetOK creates a LunMapCollectionGetOK with default headers values
func NewLunMapCollectionGetOK() *LunMapCollectionGetOK {
	return &LunMapCollectionGetOK{}
}

/* LunMapCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type LunMapCollectionGetOK struct {
	Payload *models.LunMapResponse
}

func (o *LunMapCollectionGetOK) Error() string {
	return fmt.Sprintf("[GET /protocols/san/lun-maps][%d] lunMapCollectionGetOK  %+v", 200, o.Payload)
}
func (o *LunMapCollectionGetOK) GetPayload() *models.LunMapResponse {
	return o.Payload
}

func (o *LunMapCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LunMapResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLunMapCollectionGetDefault creates a LunMapCollectionGetDefault with default headers values
func NewLunMapCollectionGetDefault(code int) *LunMapCollectionGetDefault {
	return &LunMapCollectionGetDefault{
		_statusCode: code,
	}
}

/* LunMapCollectionGetDefault describes a response with status code -1, with default header values.

Error
*/
type LunMapCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the lun map collection get default response
func (o *LunMapCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *LunMapCollectionGetDefault) Error() string {
	return fmt.Sprintf("[GET /protocols/san/lun-maps][%d] lun_map_collection_get default  %+v", o._statusCode, o.Payload)
}
func (o *LunMapCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *LunMapCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
