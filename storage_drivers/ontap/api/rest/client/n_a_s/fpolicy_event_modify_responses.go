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

// FpolicyEventModifyReader is a Reader for the FpolicyEventModify structure.
type FpolicyEventModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FpolicyEventModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFpolicyEventModifyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFpolicyEventModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFpolicyEventModifyOK creates a FpolicyEventModifyOK with default headers values
func NewFpolicyEventModifyOK() *FpolicyEventModifyOK {
	return &FpolicyEventModifyOK{}
}

/* FpolicyEventModifyOK describes a response with status code 200, with default header values.

OK
*/
type FpolicyEventModifyOK struct {
}

func (o *FpolicyEventModifyOK) Error() string {
	return fmt.Sprintf("[PATCH /protocols/fpolicy/{svm.uuid}/events/{name}][%d] fpolicyEventModifyOK ", 200)
}

func (o *FpolicyEventModifyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFpolicyEventModifyDefault creates a FpolicyEventModifyDefault with default headers values
func NewFpolicyEventModifyDefault(code int) *FpolicyEventModifyDefault {
	return &FpolicyEventModifyDefault{
		_statusCode: code,
	}
}

/* FpolicyEventModifyDefault describes a response with status code -1, with default header values.

 | Error Code | Description |
| ---------- | ----------- |
| 9764873    | The event is a cluster event |
| 9764929    | The file operation is not supported by the protocol |
| 9764955    | The filter is not supported by the protocol |
| 9764930    | The filter is not supported by any of the file operations |
| 9764946    | The protocol is specifed without file operation or a file operation and filter pair |

*/
type FpolicyEventModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the fpolicy event modify default response
func (o *FpolicyEventModifyDefault) Code() int {
	return o._statusCode
}

func (o *FpolicyEventModifyDefault) Error() string {
	return fmt.Sprintf("[PATCH /protocols/fpolicy/{svm.uuid}/events/{name}][%d] fpolicy_event_modify default  %+v", o._statusCode, o.Payload)
}
func (o *FpolicyEventModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *FpolicyEventModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
