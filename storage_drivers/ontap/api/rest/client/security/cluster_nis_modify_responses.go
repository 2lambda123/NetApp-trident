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

// ClusterNisModifyReader is a Reader for the ClusterNisModify structure.
type ClusterNisModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ClusterNisModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewClusterNisModifyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewClusterNisModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewClusterNisModifyOK creates a ClusterNisModifyOK with default headers values
func NewClusterNisModifyOK() *ClusterNisModifyOK {
	return &ClusterNisModifyOK{}
}

/* ClusterNisModifyOK describes a response with status code 200, with default header values.

OK
*/
type ClusterNisModifyOK struct {
}

func (o *ClusterNisModifyOK) Error() string {
	return fmt.Sprintf("[PATCH /security/authentication/cluster/nis][%d] clusterNisModifyOK ", 200)
}

func (o *ClusterNisModifyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewClusterNisModifyDefault creates a ClusterNisModifyDefault with default headers values
func NewClusterNisModifyDefault(code int) *ClusterNisModifyDefault {
	return &ClusterNisModifyDefault{
		_statusCode: code,
	}
}

/* ClusterNisModifyDefault describes a response with status code -1, with default header values.

 ONTAP Error Response Codes
| Error Code | Description |
| ---------- | ----------- |
| 1966253    | IPv6 is not enabled in the cluster .|
| 3276964    | The NIS domain name or NIS server domain is too long. The maximum supported for domain name is 64 characters and the maximum supported for NIS server domain is 255 characters. |
| 3276933    | A maximum of 10 NIS servers can be configured per SVM. |
| 23724109   | DNS resolution failed for one or more specified servers.  |
| 23724112   | DNS resolution failed due to an internal error. Contact technical support if this issue persists. |
| 23724132   | DNS resolution failed for all the specified servers.  |
| 23724130   | Cannot use an IPv6 name server address because there are no IPv6 interfaces |

*/
type ClusterNisModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the cluster nis modify default response
func (o *ClusterNisModifyDefault) Code() int {
	return o._statusCode
}

func (o *ClusterNisModifyDefault) Error() string {
	return fmt.Sprintf("[PATCH /security/authentication/cluster/nis][%d] cluster_nis_modify default  %+v", o._statusCode, o.Payload)
}
func (o *ClusterNisModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ClusterNisModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
