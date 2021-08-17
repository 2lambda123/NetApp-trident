// Code generated by go-swagger; DO NOT EDIT.

package networking

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/v21/storage_drivers/ontap/api/rest/models"
)

// NetworkEthernetBroadcastDomainGetReader is a Reader for the NetworkEthernetBroadcastDomainGet structure.
type NetworkEthernetBroadcastDomainGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NetworkEthernetBroadcastDomainGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNetworkEthernetBroadcastDomainGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewNetworkEthernetBroadcastDomainGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNetworkEthernetBroadcastDomainGetOK creates a NetworkEthernetBroadcastDomainGetOK with default headers values
func NewNetworkEthernetBroadcastDomainGetOK() *NetworkEthernetBroadcastDomainGetOK {
	return &NetworkEthernetBroadcastDomainGetOK{}
}

/* NetworkEthernetBroadcastDomainGetOK describes a response with status code 200, with default header values.

OK
*/
type NetworkEthernetBroadcastDomainGetOK struct {
	Payload *models.BroadcastDomain
}

func (o *NetworkEthernetBroadcastDomainGetOK) Error() string {
	return fmt.Sprintf("[GET /network/ethernet/broadcast-domains/{uuid}][%d] networkEthernetBroadcastDomainGetOK  %+v", 200, o.Payload)
}
func (o *NetworkEthernetBroadcastDomainGetOK) GetPayload() *models.BroadcastDomain {
	return o.Payload
}

func (o *NetworkEthernetBroadcastDomainGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BroadcastDomain)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNetworkEthernetBroadcastDomainGetDefault creates a NetworkEthernetBroadcastDomainGetDefault with default headers values
func NewNetworkEthernetBroadcastDomainGetDefault(code int) *NetworkEthernetBroadcastDomainGetDefault {
	return &NetworkEthernetBroadcastDomainGetDefault{
		_statusCode: code,
	}
}

/* NetworkEthernetBroadcastDomainGetDefault describes a response with status code -1, with default header values.

Error
*/
type NetworkEthernetBroadcastDomainGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the network ethernet broadcast domain get default response
func (o *NetworkEthernetBroadcastDomainGetDefault) Code() int {
	return o._statusCode
}

func (o *NetworkEthernetBroadcastDomainGetDefault) Error() string {
	return fmt.Sprintf("[GET /network/ethernet/broadcast-domains/{uuid}][%d] network_ethernet_broadcast_domain_get default  %+v", o._statusCode, o.Payload)
}
func (o *NetworkEthernetBroadcastDomainGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *NetworkEthernetBroadcastDomainGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
