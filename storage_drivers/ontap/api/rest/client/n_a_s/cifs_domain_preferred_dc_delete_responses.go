// Code generated by go-swagger; DO NOT EDIT.

package n_a_s

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// CifsDomainPreferredDcDeleteReader is a Reader for the CifsDomainPreferredDcDelete structure.
type CifsDomainPreferredDcDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CifsDomainPreferredDcDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCifsDomainPreferredDcDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCifsDomainPreferredDcDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCifsDomainPreferredDcDeleteOK creates a CifsDomainPreferredDcDeleteOK with default headers values
func NewCifsDomainPreferredDcDeleteOK() *CifsDomainPreferredDcDeleteOK {
	return &CifsDomainPreferredDcDeleteOK{}
}

/* CifsDomainPreferredDcDeleteOK describes a response with status code 200, with default header values.

OK
*/
type CifsDomainPreferredDcDeleteOK struct {
}

func (o *CifsDomainPreferredDcDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /protocols/cifs/domains/{svm.uuid}/preferred-domain-controllers/{fqdn}/{server_ip}][%d] cifsDomainPreferredDcDeleteOK ", 200)
}

func (o *CifsDomainPreferredDcDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCifsDomainPreferredDcDeleteDefault creates a CifsDomainPreferredDcDeleteDefault with default headers values
func NewCifsDomainPreferredDcDeleteDefault(code int) *CifsDomainPreferredDcDeleteDefault {
	return &CifsDomainPreferredDcDeleteDefault{
		_statusCode: code,
	}
}

/* CifsDomainPreferredDcDeleteDefault describes a response with status code -1, with default header values.

 ONTAP Error Response Codes
| Error Code | Description |
| ---------- | ----------- |
| 655507     | Failed to remove preferred-dc. |

*/
type CifsDomainPreferredDcDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the cifs domain preferred dc delete default response
func (o *CifsDomainPreferredDcDeleteDefault) Code() int {
	return o._statusCode
}

func (o *CifsDomainPreferredDcDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /protocols/cifs/domains/{svm.uuid}/preferred-domain-controllers/{fqdn}/{server_ip}][%d] cifs_domain_preferred_dc_delete default  %+v", o._statusCode, o.Payload)
}
func (o *CifsDomainPreferredDcDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CifsDomainPreferredDcDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
