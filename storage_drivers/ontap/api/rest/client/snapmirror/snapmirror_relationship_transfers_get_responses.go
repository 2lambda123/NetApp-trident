// Code generated by go-swagger; DO NOT EDIT.

package snapmirror

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/v21/storage_drivers/ontap/api/rest/models"
)

// SnapmirrorRelationshipTransfersGetReader is a Reader for the SnapmirrorRelationshipTransfersGet structure.
type SnapmirrorRelationshipTransfersGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SnapmirrorRelationshipTransfersGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSnapmirrorRelationshipTransfersGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSnapmirrorRelationshipTransfersGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSnapmirrorRelationshipTransfersGetOK creates a SnapmirrorRelationshipTransfersGetOK with default headers values
func NewSnapmirrorRelationshipTransfersGetOK() *SnapmirrorRelationshipTransfersGetOK {
	return &SnapmirrorRelationshipTransfersGetOK{}
}

/* SnapmirrorRelationshipTransfersGetOK describes a response with status code 200, with default header values.

OK
*/
type SnapmirrorRelationshipTransfersGetOK struct {
	Payload *models.SnapmirrorTransferResponse
}

func (o *SnapmirrorRelationshipTransfersGetOK) Error() string {
	return fmt.Sprintf("[GET /snapmirror/relationships/{relationship.uuid}/transfers][%d] snapmirrorRelationshipTransfersGetOK  %+v", 200, o.Payload)
}
func (o *SnapmirrorRelationshipTransfersGetOK) GetPayload() *models.SnapmirrorTransferResponse {
	return o.Payload
}

func (o *SnapmirrorRelationshipTransfersGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SnapmirrorTransferResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSnapmirrorRelationshipTransfersGetDefault creates a SnapmirrorRelationshipTransfersGetDefault with default headers values
func NewSnapmirrorRelationshipTransfersGetDefault(code int) *SnapmirrorRelationshipTransfersGetDefault {
	return &SnapmirrorRelationshipTransfersGetDefault{
		_statusCode: code,
	}
}

/* SnapmirrorRelationshipTransfersGetDefault describes a response with status code -1, with default header values.

Error
*/
type SnapmirrorRelationshipTransfersGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the snapmirror relationship transfers get default response
func (o *SnapmirrorRelationshipTransfersGetDefault) Code() int {
	return o._statusCode
}

func (o *SnapmirrorRelationshipTransfersGetDefault) Error() string {
	return fmt.Sprintf("[GET /snapmirror/relationships/{relationship.uuid}/transfers][%d] snapmirror_relationship_transfers_get default  %+v", o._statusCode, o.Payload)
}
func (o *SnapmirrorRelationshipTransfersGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SnapmirrorRelationshipTransfersGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
