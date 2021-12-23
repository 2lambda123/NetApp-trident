// Code generated by go-swagger; DO NOT EDIT.

package support

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// AutoUpdateConfigurationModifyReader is a Reader for the AutoUpdateConfigurationModify structure.
type AutoUpdateConfigurationModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AutoUpdateConfigurationModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAutoUpdateConfigurationModifyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAutoUpdateConfigurationModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAutoUpdateConfigurationModifyOK creates a AutoUpdateConfigurationModifyOK with default headers values
func NewAutoUpdateConfigurationModifyOK() *AutoUpdateConfigurationModifyOK {
	return &AutoUpdateConfigurationModifyOK{}
}

/* AutoUpdateConfigurationModifyOK describes a response with status code 200, with default header values.

OK
*/
type AutoUpdateConfigurationModifyOK struct {
}

func (o *AutoUpdateConfigurationModifyOK) Error() string {
	return fmt.Sprintf("[PATCH /support/auto-update/configurations/{uuid}][%d] autoUpdateConfigurationModifyOK ", 200)
}

func (o *AutoUpdateConfigurationModifyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAutoUpdateConfigurationModifyDefault creates a AutoUpdateConfigurationModifyDefault with default headers values
func NewAutoUpdateConfigurationModifyDefault(code int) *AutoUpdateConfigurationModifyDefault {
	return &AutoUpdateConfigurationModifyDefault{
		_statusCode: code,
	}
}

/* AutoUpdateConfigurationModifyDefault describes a response with status code -1, with default header values.

 ONTAP Error Response Codes
| Error Code | Description |
| ---------- | ----------- |
| 262179 | Unexpected argument. |

*/
type AutoUpdateConfigurationModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the auto update configuration modify default response
func (o *AutoUpdateConfigurationModifyDefault) Code() int {
	return o._statusCode
}

func (o *AutoUpdateConfigurationModifyDefault) Error() string {
	return fmt.Sprintf("[PATCH /support/auto-update/configurations/{uuid}][%d] auto_update_configuration_modify default  %+v", o._statusCode, o.Payload)
}
func (o *AutoUpdateConfigurationModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AutoUpdateConfigurationModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
