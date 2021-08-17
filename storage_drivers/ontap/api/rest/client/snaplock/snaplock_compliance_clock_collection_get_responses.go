// Code generated by go-swagger; DO NOT EDIT.

package snaplock

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/v21/storage_drivers/ontap/api/rest/models"
)

// SnaplockComplianceClockCollectionGetReader is a Reader for the SnaplockComplianceClockCollectionGet structure.
type SnaplockComplianceClockCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SnaplockComplianceClockCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSnaplockComplianceClockCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSnaplockComplianceClockCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSnaplockComplianceClockCollectionGetOK creates a SnaplockComplianceClockCollectionGetOK with default headers values
func NewSnaplockComplianceClockCollectionGetOK() *SnaplockComplianceClockCollectionGetOK {
	return &SnaplockComplianceClockCollectionGetOK{}
}

/* SnaplockComplianceClockCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type SnaplockComplianceClockCollectionGetOK struct {
	Payload *models.SnaplockComplianceClockResponse
}

func (o *SnaplockComplianceClockCollectionGetOK) Error() string {
	return fmt.Sprintf("[GET /storage/snaplock/compliance-clocks][%d] snaplockComplianceClockCollectionGetOK  %+v", 200, o.Payload)
}
func (o *SnaplockComplianceClockCollectionGetOK) GetPayload() *models.SnaplockComplianceClockResponse {
	return o.Payload
}

func (o *SnaplockComplianceClockCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SnaplockComplianceClockResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSnaplockComplianceClockCollectionGetDefault creates a SnaplockComplianceClockCollectionGetDefault with default headers values
func NewSnaplockComplianceClockCollectionGetDefault(code int) *SnaplockComplianceClockCollectionGetDefault {
	return &SnaplockComplianceClockCollectionGetDefault{
		_statusCode: code,
	}
}

/* SnaplockComplianceClockCollectionGetDefault describes a response with status code -1, with default header values.

Error
*/
type SnaplockComplianceClockCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the snaplock compliance clock collection get default response
func (o *SnaplockComplianceClockCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *SnaplockComplianceClockCollectionGetDefault) Error() string {
	return fmt.Sprintf("[GET /storage/snaplock/compliance-clocks][%d] snaplock_compliance_clock_collection_get default  %+v", o._statusCode, o.Payload)
}
func (o *SnaplockComplianceClockCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SnaplockComplianceClockCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
