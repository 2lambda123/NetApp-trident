// Code generated by go-swagger; DO NOT EDIT.

package name_services

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// HostRecordGetReader is a Reader for the HostRecordGet structure.
type HostRecordGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HostRecordGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewHostRecordGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewHostRecordGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewHostRecordGetOK creates a HostRecordGetOK with default headers values
func NewHostRecordGetOK() *HostRecordGetOK {
	return &HostRecordGetOK{}
}

/* HostRecordGetOK describes a response with status code 200, with default header values.

OK
*/
type HostRecordGetOK struct {
	Payload *models.HostRecord
}

func (o *HostRecordGetOK) Error() string {
	return fmt.Sprintf("[GET /name-services/host-record/{svm.uuid}/{host}][%d] hostRecordGetOK  %+v", 200, o.Payload)
}
func (o *HostRecordGetOK) GetPayload() *models.HostRecord {
	return o.Payload
}

func (o *HostRecordGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HostRecord)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHostRecordGetDefault creates a HostRecordGetDefault with default headers values
func NewHostRecordGetDefault(code int) *HostRecordGetDefault {
	return &HostRecordGetDefault{
		_statusCode: code,
	}
}

/* HostRecordGetDefault describes a response with status code -1, with default header values.

 ONTAP Error Response Codes
| Error Code | Description |
| ---------- | ----------- |
|  8912900   | Invalid IP address. |

*/
type HostRecordGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the host record get default response
func (o *HostRecordGetDefault) Code() int {
	return o._statusCode
}

func (o *HostRecordGetDefault) Error() string {
	return fmt.Sprintf("[GET /name-services/host-record/{svm.uuid}/{host}][%d] host_record_get default  %+v", o._statusCode, o.Payload)
}
func (o *HostRecordGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *HostRecordGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
