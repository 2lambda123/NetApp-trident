// Code generated by go-swagger; DO NOT EDIT.

package storage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/v21/storage_drivers/ontap/api/rest/models"
)

// VolumeGetReader is a Reader for the VolumeGet structure.
type VolumeGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VolumeGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVolumeGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewVolumeGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewVolumeGetOK creates a VolumeGetOK with default headers values
func NewVolumeGetOK() *VolumeGetOK {
	return &VolumeGetOK{}
}

/* VolumeGetOK describes a response with status code 200, with default header values.

OK
*/
type VolumeGetOK struct {
	Payload *models.Volume
}

func (o *VolumeGetOK) Error() string {
	return fmt.Sprintf("[GET /storage/volumes/{uuid}][%d] volumeGetOK  %+v", 200, o.Payload)
}
func (o *VolumeGetOK) GetPayload() *models.Volume {
	return o.Payload
}

func (o *VolumeGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Volume)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVolumeGetDefault creates a VolumeGetDefault with default headers values
func NewVolumeGetDefault(code int) *VolumeGetDefault {
	return &VolumeGetDefault{
		_statusCode: code,
	}
}

/* VolumeGetDefault describes a response with status code -1, with default header values.

Error
*/
type VolumeGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the volume get default response
func (o *VolumeGetDefault) Code() int {
	return o._statusCode
}

func (o *VolumeGetDefault) Error() string {
	return fmt.Sprintf("[GET /storage/volumes/{uuid}][%d] volume_get default  %+v", o._statusCode, o.Payload)
}
func (o *VolumeGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *VolumeGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
