// Code generated by go-swagger; DO NOT EDIT.

package s_a_n

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/v21/storage_drivers/ontap/api/rest/models"
)

// IgroupGetReader is a Reader for the IgroupGet structure.
type IgroupGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IgroupGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIgroupGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIgroupGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIgroupGetOK creates a IgroupGetOK with default headers values
func NewIgroupGetOK() *IgroupGetOK {
	return &IgroupGetOK{}
}

/* IgroupGetOK describes a response with status code 200, with default header values.

OK
*/
type IgroupGetOK struct {
	Payload *models.Igroup
}

func (o *IgroupGetOK) Error() string {
	return fmt.Sprintf("[GET /protocols/san/igroups/{uuid}][%d] igroupGetOK  %+v", 200, o.Payload)
}
func (o *IgroupGetOK) GetPayload() *models.Igroup {
	return o.Payload
}

func (o *IgroupGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Igroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIgroupGetDefault creates a IgroupGetDefault with default headers values
func NewIgroupGetDefault(code int) *IgroupGetDefault {
	return &IgroupGetDefault{
		_statusCode: code,
	}
}

/* IgroupGetDefault describes a response with status code -1, with default header values.

 ONTAP Error Response Codes
| Error Code | Description |
| ---------- | ----------- |
| 5374852 | The initiator group does not exist. |

*/
type IgroupGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the igroup get default response
func (o *IgroupGetDefault) Code() int {
	return o._statusCode
}

func (o *IgroupGetDefault) Error() string {
	return fmt.Sprintf("[GET /protocols/san/igroups/{uuid}][%d] igroup_get default  %+v", o._statusCode, o.Payload)
}
func (o *IgroupGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *IgroupGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
