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

// ExportRuleGetReader is a Reader for the ExportRuleGet structure.
type ExportRuleGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExportRuleGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExportRuleGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewExportRuleGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExportRuleGetOK creates a ExportRuleGetOK with default headers values
func NewExportRuleGetOK() *ExportRuleGetOK {
	return &ExportRuleGetOK{}
}

/* ExportRuleGetOK describes a response with status code 200, with default header values.

OK
*/
type ExportRuleGetOK struct {
	Payload *models.ExportRule
}

func (o *ExportRuleGetOK) Error() string {
	return fmt.Sprintf("[GET /protocols/nfs/export-policies/{policy.id}/rules/{index}][%d] exportRuleGetOK  %+v", 200, o.Payload)
}
func (o *ExportRuleGetOK) GetPayload() *models.ExportRule {
	return o.Payload
}

func (o *ExportRuleGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ExportRule)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExportRuleGetDefault creates a ExportRuleGetDefault with default headers values
func NewExportRuleGetDefault(code int) *ExportRuleGetDefault {
	return &ExportRuleGetDefault{
		_statusCode: code,
	}
}

/* ExportRuleGetDefault describes a response with status code -1, with default header values.

Error
*/
type ExportRuleGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the export rule get default response
func (o *ExportRuleGetDefault) Code() int {
	return o._statusCode
}

func (o *ExportRuleGetDefault) Error() string {
	return fmt.Sprintf("[GET /protocols/nfs/export-policies/{policy.id}/rules/{index}][%d] export_rule_get default  %+v", o._statusCode, o.Payload)
}
func (o *ExportRuleGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ExportRuleGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
