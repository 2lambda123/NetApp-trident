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

// SnmpTraphostsCreateReader is a Reader for the SnmpTraphostsCreate structure.
type SnmpTraphostsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SnmpTraphostsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewSnmpTraphostsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSnmpTraphostsCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSnmpTraphostsCreateCreated creates a SnmpTraphostsCreateCreated with default headers values
func NewSnmpTraphostsCreateCreated() *SnmpTraphostsCreateCreated {
	return &SnmpTraphostsCreateCreated{}
}

/* SnmpTraphostsCreateCreated describes a response with status code 201, with default header values.

Created
*/
type SnmpTraphostsCreateCreated struct {
	Payload *models.SnmpTraphostResponse
}

func (o *SnmpTraphostsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /support/snmp/traphosts][%d] snmpTraphostsCreateCreated  %+v", 201, o.Payload)
}
func (o *SnmpTraphostsCreateCreated) GetPayload() *models.SnmpTraphostResponse {
	return o.Payload
}

func (o *SnmpTraphostsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SnmpTraphostResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSnmpTraphostsCreateDefault creates a SnmpTraphostsCreateDefault with default headers values
func NewSnmpTraphostsCreateDefault(code int) *SnmpTraphostsCreateDefault {
	return &SnmpTraphostsCreateDefault{
		_statusCode: code,
	}
}

/* SnmpTraphostsCreateDefault describes a response with status code -1, with default header values.

 ONTAP Error Response Codes
| Error Code | Description |
| ---------- | ----------- |
| 8847365 | Unknown host. |
| 9043969 | Duplicate traphost entry. |
| 9043980 | IPv6 support is disabled. |
| 9043991 | Not a USM user. |
| 9043993 | Current cluster version does not support SNMPv3 traps. |
| 9044001 | Failed to create SNMPv1 traphost. |
| 9044002 | Failed to create SNMPv3 traphost. |

*/
type SnmpTraphostsCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the snmp traphosts create default response
func (o *SnmpTraphostsCreateDefault) Code() int {
	return o._statusCode
}

func (o *SnmpTraphostsCreateDefault) Error() string {
	return fmt.Sprintf("[POST /support/snmp/traphosts][%d] snmp_traphosts_create default  %+v", o._statusCode, o.Payload)
}
func (o *SnmpTraphostsCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SnmpTraphostsCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
