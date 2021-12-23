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

// AccountModifyReader is a Reader for the AccountModify structure.
type AccountModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AccountModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAccountModifyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAccountModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAccountModifyOK creates a AccountModifyOK with default headers values
func NewAccountModifyOK() *AccountModifyOK {
	return &AccountModifyOK{}
}

/* AccountModifyOK describes a response with status code 200, with default header values.

OK
*/
type AccountModifyOK struct {
}

func (o *AccountModifyOK) Error() string {
	return fmt.Sprintf("[PATCH /security/accounts/{owner.uuid}/{name}][%d] accountModifyOK ", 200)
}

func (o *AccountModifyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAccountModifyDefault creates a AccountModifyDefault with default headers values
func NewAccountModifyDefault(code int) *AccountModifyDefault {
	return &AccountModifyDefault{
		_statusCode: code,
	}
}

/* AccountModifyDefault describes a response with status code -1, with default header values.

 ONTAP Error Response Codes
| Error Code | Description |
| ---------- | ----------- |
| 1261215 | The role was not found. |
| 1261218 | The user was not found. |
| 1263343 | Cannot lock user with password not set or non-password authentication method. |
| 5636096 | Cannot perform the operation for this user account since the password is not set. |
| 5636097 | The operation for user account failed since user password is not set. |
| 5636100 | Modification of a service-processor user's role to a non-admin role is not supported. |
| 5636125 | The operation not supported on AutoSupport user account which is reserved. |
| 5636129 | The role does not exist. |
| 5636154 | The second-authentication-method parameter is supported for ssh application. |
| 5636155 | The second-authentication-method parameter can be specified only if the authentication-method password or public key nsswitch. |
| 5636156 | Same value cannot be specified for the second-authentication-method and the authentication-method. |
| 5636157 | If the authentication-method is domain, the second-authentication-method cannot be specified. |
| 5636159 | For a given user and application, if the second-authentication-method is specified, only one such login entry is supported. |
| 5636164 | If the value for either the authentication-method second-authentication-method is nsswitch or password, the other parameter must differ. |
| 7077896 | Cannot lock the account of the last console admin user. |
| 7077906 | A role with that name has not been defined for the Vserver. |
| 7077911 | The user is not configured to use the password authentication method. |
| 7077918 | The password cannot contain the username. |
| 7077919 | The minimum length for new password does not meet the policy. |
| 7077920 | The new password must have both letters and numbers. |
| 7077921 | The minimum number of special characters required do not meet the policy. |
| 7077924 | The new password must be different than last N passwords. |
| 7077925 | The new password must be different to the old password. |
| 7077929 | Cannot lock user with password not set or non-password authentication method. |
| 7077940 | The password exceeds maximum supported length. |
| 7077941 | Defined password composition exceeds the maximum password length of 128 characters. |
| 7078900 | An aAdmin password is not set. Set the password by including it in the request. |

*/
type AccountModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the account modify default response
func (o *AccountModifyDefault) Code() int {
	return o._statusCode
}

func (o *AccountModifyDefault) Error() string {
	return fmt.Sprintf("[PATCH /security/accounts/{owner.uuid}/{name}][%d] account_modify default  %+v", o._statusCode, o.Payload)
}
func (o *AccountModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AccountModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
