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

// VolumeDeleteReader is a Reader for the VolumeDelete structure.
type VolumeDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VolumeDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewVolumeDeleteAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewVolumeDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewVolumeDeleteAccepted creates a VolumeDeleteAccepted with default headers values
func NewVolumeDeleteAccepted() *VolumeDeleteAccepted {
	return &VolumeDeleteAccepted{}
}

/* VolumeDeleteAccepted describes a response with status code 202, with default header values.

Accepted
*/
type VolumeDeleteAccepted struct {
	Payload *models.JobLinkResponse
}

func (o *VolumeDeleteAccepted) Error() string {
	return fmt.Sprintf("[DELETE /storage/volumes/{uuid}][%d] volumeDeleteAccepted  %+v", 202, o.Payload)
}
func (o *VolumeDeleteAccepted) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *VolumeDeleteAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVolumeDeleteDefault creates a VolumeDeleteDefault with default headers values
func NewVolumeDeleteDefault(code int) *VolumeDeleteDefault {
	return &VolumeDeleteDefault{
		_statusCode: code,
	}
}

/* VolumeDeleteDefault describes a response with status code -1, with default header values.

Error
*/
type VolumeDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the volume delete default response
func (o *VolumeDeleteDefault) Code() int {
	return o._statusCode
}

func (o *VolumeDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /storage/volumes/{uuid}][%d] volume_delete default  %+v", o._statusCode, o.Payload)
}
func (o *VolumeDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *VolumeDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
