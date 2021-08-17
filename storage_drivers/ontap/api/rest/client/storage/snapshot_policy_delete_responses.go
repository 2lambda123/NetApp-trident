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

// SnapshotPolicyDeleteReader is a Reader for the SnapshotPolicyDelete structure.
type SnapshotPolicyDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SnapshotPolicyDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSnapshotPolicyDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSnapshotPolicyDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSnapshotPolicyDeleteOK creates a SnapshotPolicyDeleteOK with default headers values
func NewSnapshotPolicyDeleteOK() *SnapshotPolicyDeleteOK {
	return &SnapshotPolicyDeleteOK{}
}

/* SnapshotPolicyDeleteOK describes a response with status code 200, with default header values.

OK
*/
type SnapshotPolicyDeleteOK struct {
}

func (o *SnapshotPolicyDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /storage/snapshot-policies/{uuid}][%d] snapshotPolicyDeleteOK ", 200)
}

func (o *SnapshotPolicyDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSnapshotPolicyDeleteDefault creates a SnapshotPolicyDeleteDefault with default headers values
func NewSnapshotPolicyDeleteDefault(code int) *SnapshotPolicyDeleteDefault {
	return &SnapshotPolicyDeleteDefault{
		_statusCode: code,
	}
}

/* SnapshotPolicyDeleteDefault describes a response with status code -1, with default header values.

 ONTAP Error Response Code
| Error Code | Description |
| ---------- | ----------- |
| 1638415    | Cannot delete policy. Reason: Policy is in use by at least one volume. |
| 1638416    | Cannot delete policy. Reason: Cannot verify whether policy is in use. |
| 1638430    | Cannot delete policy. Reason: Policy is in use by at least one Vserver. |
| 1638430    | Cannot delete built-in policy. |

*/
type SnapshotPolicyDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the snapshot policy delete default response
func (o *SnapshotPolicyDeleteDefault) Code() int {
	return o._statusCode
}

func (o *SnapshotPolicyDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /storage/snapshot-policies/{uuid}][%d] snapshot_policy_delete default  %+v", o._statusCode, o.Payload)
}
func (o *SnapshotPolicyDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SnapshotPolicyDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
