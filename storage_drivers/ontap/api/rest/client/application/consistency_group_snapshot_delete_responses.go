// Code generated by go-swagger; DO NOT EDIT.

package application

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// ConsistencyGroupSnapshotDeleteReader is a Reader for the ConsistencyGroupSnapshotDelete structure.
type ConsistencyGroupSnapshotDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ConsistencyGroupSnapshotDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewConsistencyGroupSnapshotDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewConsistencyGroupSnapshotDeleteAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewConsistencyGroupSnapshotDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewConsistencyGroupSnapshotDeleteOK creates a ConsistencyGroupSnapshotDeleteOK with default headers values
func NewConsistencyGroupSnapshotDeleteOK() *ConsistencyGroupSnapshotDeleteOK {
	return &ConsistencyGroupSnapshotDeleteOK{}
}

/* ConsistencyGroupSnapshotDeleteOK describes a response with status code 200, with default header values.

OK
*/
type ConsistencyGroupSnapshotDeleteOK struct {
}

func (o *ConsistencyGroupSnapshotDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /application/consistency-groups/{consistency_group.uuid}/snapshots/{uuid}][%d] consistencyGroupSnapshotDeleteOK ", 200)
}

func (o *ConsistencyGroupSnapshotDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewConsistencyGroupSnapshotDeleteAccepted creates a ConsistencyGroupSnapshotDeleteAccepted with default headers values
func NewConsistencyGroupSnapshotDeleteAccepted() *ConsistencyGroupSnapshotDeleteAccepted {
	return &ConsistencyGroupSnapshotDeleteAccepted{}
}

/* ConsistencyGroupSnapshotDeleteAccepted describes a response with status code 202, with default header values.

Accepted
*/
type ConsistencyGroupSnapshotDeleteAccepted struct {
}

func (o *ConsistencyGroupSnapshotDeleteAccepted) Error() string {
	return fmt.Sprintf("[DELETE /application/consistency-groups/{consistency_group.uuid}/snapshots/{uuid}][%d] consistencyGroupSnapshotDeleteAccepted ", 202)
}

func (o *ConsistencyGroupSnapshotDeleteAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewConsistencyGroupSnapshotDeleteDefault creates a ConsistencyGroupSnapshotDeleteDefault with default headers values
func NewConsistencyGroupSnapshotDeleteDefault(code int) *ConsistencyGroupSnapshotDeleteDefault {
	return &ConsistencyGroupSnapshotDeleteDefault{
		_statusCode: code,
	}
}

/* ConsistencyGroupSnapshotDeleteDefault describes a response with status code -1, with default header values.

Error
*/
type ConsistencyGroupSnapshotDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the consistency group snapshot delete default response
func (o *ConsistencyGroupSnapshotDeleteDefault) Code() int {
	return o._statusCode
}

func (o *ConsistencyGroupSnapshotDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /application/consistency-groups/{consistency_group.uuid}/snapshots/{uuid}][%d] consistency_group_snapshot_delete default  %+v", o._statusCode, o.Payload)
}
func (o *ConsistencyGroupSnapshotDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ConsistencyGroupSnapshotDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
