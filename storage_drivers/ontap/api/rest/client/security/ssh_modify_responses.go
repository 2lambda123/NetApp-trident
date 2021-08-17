// Code generated by go-swagger; DO NOT EDIT.

package security

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/v21/storage_drivers/ontap/api/rest/models"
)

// SSHModifyReader is a Reader for the SSHModify structure.
type SSHModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SSHModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSSHModifyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSSHModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSSHModifyOK creates a SSHModifyOK with default headers values
func NewSSHModifyOK() *SSHModifyOK {
	return &SSHModifyOK{}
}

/* SSHModifyOK describes a response with status code 200, with default header values.

OK
*/
type SSHModifyOK struct {
}

func (o *SSHModifyOK) Error() string {
	return fmt.Sprintf("[PATCH /security/ssh][%d] sshModifyOK ", 200)
}

func (o *SSHModifyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSSHModifyDefault creates a SSHModifyDefault with default headers values
func NewSSHModifyDefault(code int) *SSHModifyDefault {
	return &SSHModifyDefault{
		_statusCode: code,
	}
}

/* SSHModifyDefault describes a response with status code -1, with default header values.

 ONTAP Error Response Codes
| Error Code | Description |
| ---------- | ----------- |
| 10682372 | There must be at least one key exchange algorithm associated with the SSH configuration. |
| 10682373 | There must be at least one cipher associated with the SSH configuration. |
| 10682375 | Failed to modify SSH key exchange algorithms. |
| 10682378 | Failed to modify SSH ciphers. |
| 10682399 | Key exchange algorithm not supported in FIPS enabled mode. |
| 10682400 | Failed to modify SSH MAC algorithms. |
| 10682401 | MAC algorithm not supported in FIPS enabled mode. |
| 10682403 | There must be at least one MAC algorithm with the SSH configuration. |
| 10682413 | Failed to modify maximum authentication retry attempts. |
| 10682413 | Failed to modify maximum authentication retry attempts. |
| 10682418 | Cipher not supported in FIPS enabled mode. |

*/
type SSHModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the ssh modify default response
func (o *SSHModifyDefault) Code() int {
	return o._statusCode
}

func (o *SSHModifyDefault) Error() string {
	return fmt.Sprintf("[PATCH /security/ssh][%d] ssh_modify default  %+v", o._statusCode, o.Payload)
}
func (o *SSHModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SSHModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
