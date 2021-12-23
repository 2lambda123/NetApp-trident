// Code generated by go-swagger; DO NOT EDIT.

package security

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// RoleDeleteReader is a Reader for the RoleDelete structure.
type RoleDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RoleDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRoleDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewRoleDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRoleDeleteOK creates a RoleDeleteOK with default headers values
func NewRoleDeleteOK() *RoleDeleteOK {
	return &RoleDeleteOK{}
}

/* RoleDeleteOK describes a response with status code 200, with default header values.

OK
*/
type RoleDeleteOK struct {
}

func (o *RoleDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /security/roles/{owner.uuid}/{name}][%d] roleDeleteOK ", 200)
}

func (o *RoleDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRoleDeleteDefault creates a RoleDeleteDefault with default headers values
func NewRoleDeleteDefault(code int) *RoleDeleteDefault {
	return &RoleDeleteDefault{
		_statusCode: code,
	}
}

/* RoleDeleteDefault describes a response with status code -1, with default header values.

 ONTAP Error Response Codes
| Error Code | Description |
| ---------- | ----------- |
| 5636169 | Specified URI path is invalid or not supported. Resource-qualified endpoints are not supported. |
| 5636170 | URI does not exist. |
| 5636172 | User accounts detected with this role assigned. Update or delete those accounts before deleting this role. |
| 5636173 | Features require an effective cluster version of 9.6 or later. |
| 5636184 | Expanded REST roles for granular resource control feature is currently disabled. |
| 5636185 | The specified UUID was not found. |
| 5636186 | Expanded REST roles for granular resource control requires an effective cluster version of 9.10.1 or later. |
| 13434890 | Vserver-ID failed for Vserver roles. |
| 13434893 | The SVM does not exist. |

*/
type RoleDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the role delete default response
func (o *RoleDeleteDefault) Code() int {
	return o._statusCode
}

func (o *RoleDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /security/roles/{owner.uuid}/{name}][%d] role_delete default  %+v", o._statusCode, o.Payload)
}
func (o *RoleDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *RoleDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
