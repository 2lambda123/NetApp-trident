// Code generated by go-swagger; DO NOT EDIT.

package n_v_me

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/v21/storage_drivers/ontap/api/rest/models"
)

// NvmeNamespaceCollectionGetReader is a Reader for the NvmeNamespaceCollectionGet structure.
type NvmeNamespaceCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NvmeNamespaceCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNvmeNamespaceCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewNvmeNamespaceCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNvmeNamespaceCollectionGetOK creates a NvmeNamespaceCollectionGetOK with default headers values
func NewNvmeNamespaceCollectionGetOK() *NvmeNamespaceCollectionGetOK {
	return &NvmeNamespaceCollectionGetOK{}
}

/* NvmeNamespaceCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type NvmeNamespaceCollectionGetOK struct {
	Payload *models.NvmeNamespaceResponse
}

func (o *NvmeNamespaceCollectionGetOK) Error() string {
	return fmt.Sprintf("[GET /storage/namespaces][%d] nvmeNamespaceCollectionGetOK  %+v", 200, o.Payload)
}
func (o *NvmeNamespaceCollectionGetOK) GetPayload() *models.NvmeNamespaceResponse {
	return o.Payload
}

func (o *NvmeNamespaceCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NvmeNamespaceResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNvmeNamespaceCollectionGetDefault creates a NvmeNamespaceCollectionGetDefault with default headers values
func NewNvmeNamespaceCollectionGetDefault(code int) *NvmeNamespaceCollectionGetDefault {
	return &NvmeNamespaceCollectionGetDefault{
		_statusCode: code,
	}
}

/* NvmeNamespaceCollectionGetDefault describes a response with status code -1, with default header values.

Error
*/
type NvmeNamespaceCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the nvme namespace collection get default response
func (o *NvmeNamespaceCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *NvmeNamespaceCollectionGetDefault) Error() string {
	return fmt.Sprintf("[GET /storage/namespaces][%d] nvme_namespace_collection_get default  %+v", o._statusCode, o.Payload)
}
func (o *NvmeNamespaceCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *NvmeNamespaceCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
