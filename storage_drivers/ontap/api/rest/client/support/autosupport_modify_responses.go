// Code generated by go-swagger; DO NOT EDIT.

package support

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/v21/storage_drivers/ontap/api/rest/models"
)

// AutosupportModifyReader is a Reader for the AutosupportModify structure.
type AutosupportModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AutosupportModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAutosupportModifyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAutosupportModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAutosupportModifyOK creates a AutosupportModifyOK with default headers values
func NewAutosupportModifyOK() *AutosupportModifyOK {
	return &AutosupportModifyOK{}
}

/* AutosupportModifyOK describes a response with status code 200, with default header values.

OK
*/
type AutosupportModifyOK struct {
}

func (o *AutosupportModifyOK) Error() string {
	return fmt.Sprintf("[PATCH /support/autosupport][%d] autosupportModifyOK ", 200)
}

func (o *AutosupportModifyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAutosupportModifyDefault creates a AutosupportModifyDefault with default headers values
func NewAutosupportModifyDefault(code int) *AutosupportModifyDefault {
	return &AutosupportModifyDefault{
		_statusCode: code,
	}
}

/* AutosupportModifyDefault describes a response with status code -1, with default header values.

 ONTAP Error Response Codes
| Error Code | Description |
| ---------- | ----------- |
| 8650862    | The SMTP mail host provided cannot be empty |
| 8650863    | A maximum of 5 SMTP mail hosts can be provided |
| 8650864    | A maximum of 5 email addresses can be provided |
| 8650865    | A maximum of 5 partner email addresses can be provided |
| 53149727   | The proxy URI provided is invalid |
| 53149728   | The mailhost URI provided is invalid |

*/
type AutosupportModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the autosupport modify default response
func (o *AutosupportModifyDefault) Code() int {
	return o._statusCode
}

func (o *AutosupportModifyDefault) Error() string {
	return fmt.Sprintf("[PATCH /support/autosupport][%d] autosupport_modify default  %+v", o._statusCode, o.Payload)
}
func (o *AutosupportModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AutosupportModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
