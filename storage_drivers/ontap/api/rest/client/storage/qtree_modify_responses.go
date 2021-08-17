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

// QtreeModifyReader is a Reader for the QtreeModify structure.
type QtreeModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QtreeModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewQtreeModifyAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewQtreeModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewQtreeModifyAccepted creates a QtreeModifyAccepted with default headers values
func NewQtreeModifyAccepted() *QtreeModifyAccepted {
	return &QtreeModifyAccepted{}
}

/* QtreeModifyAccepted describes a response with status code 202, with default header values.

Accepted
*/
type QtreeModifyAccepted struct {
	Payload *models.JobLinkResponse
}

func (o *QtreeModifyAccepted) Error() string {
	return fmt.Sprintf("[PATCH /storage/qtrees/{volume.uuid}/{id}][%d] qtreeModifyAccepted  %+v", 202, o.Payload)
}
func (o *QtreeModifyAccepted) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *QtreeModifyAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQtreeModifyDefault creates a QtreeModifyDefault with default headers values
func NewQtreeModifyDefault(code int) *QtreeModifyDefault {
	return &QtreeModifyDefault{
		_statusCode: code,
	}
}

/* QtreeModifyDefault describes a response with status code -1, with default header values.

 ONTAP Error Response Codes
| Error Code | Description |
| ---------- | ----------- |
| 918235 | A volume with UUID was not found. |
| 5242951 | Export policy supplied does not belong to the specified export policy ID. |
| 5242955 | The UUID of the volume is required. |
| 5242956 | Failed to obtain qtree with ID. |
| 5242958 | Failed to rename qtree in volume in SVM with ID. |
| 5242959 | Successfully renamed qtree but modify failed. |

*/
type QtreeModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the qtree modify default response
func (o *QtreeModifyDefault) Code() int {
	return o._statusCode
}

func (o *QtreeModifyDefault) Error() string {
	return fmt.Sprintf("[PATCH /storage/qtrees/{volume.uuid}/{id}][%d] qtree_modify default  %+v", o._statusCode, o.Payload)
}
func (o *QtreeModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *QtreeModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
