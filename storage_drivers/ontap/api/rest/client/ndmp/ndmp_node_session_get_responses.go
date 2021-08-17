// Code generated by go-swagger; DO NOT EDIT.

package ndmp

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/v21/storage_drivers/ontap/api/rest/models"
)

// NdmpNodeSessionGetReader is a Reader for the NdmpNodeSessionGet structure.
type NdmpNodeSessionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NdmpNodeSessionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNdmpNodeSessionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewNdmpNodeSessionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNdmpNodeSessionGetOK creates a NdmpNodeSessionGetOK with default headers values
func NewNdmpNodeSessionGetOK() *NdmpNodeSessionGetOK {
	return &NdmpNodeSessionGetOK{}
}

/* NdmpNodeSessionGetOK describes a response with status code 200, with default header values.

OK
*/
type NdmpNodeSessionGetOK struct {
	Payload *models.NdmpSession
}

func (o *NdmpNodeSessionGetOK) Error() string {
	return fmt.Sprintf("[GET /protocols/ndmp/sessions/{owner.uuid}/{session.id}][%d] ndmpNodeSessionGetOK  %+v", 200, o.Payload)
}
func (o *NdmpNodeSessionGetOK) GetPayload() *models.NdmpSession {
	return o.Payload
}

func (o *NdmpNodeSessionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NdmpSession)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNdmpNodeSessionGetDefault creates a NdmpNodeSessionGetDefault with default headers values
func NewNdmpNodeSessionGetDefault(code int) *NdmpNodeSessionGetDefault {
	return &NdmpNodeSessionGetDefault{
		_statusCode: code,
	}
}

/* NdmpNodeSessionGetDefault describes a response with status code -1, with default header values.

 ONTAP Error Response codes
| Error code  |  Description |
|-------------|--------------|
| 68812802    | The UUID is not valid.|

*/
type NdmpNodeSessionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the ndmp node session get default response
func (o *NdmpNodeSessionGetDefault) Code() int {
	return o._statusCode
}

func (o *NdmpNodeSessionGetDefault) Error() string {
	return fmt.Sprintf("[GET /protocols/ndmp/sessions/{owner.uuid}/{session.id}][%d] ndmp_node_session_get default  %+v", o._statusCode, o.Payload)
}
func (o *NdmpNodeSessionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *NdmpNodeSessionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
