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

// NetworkIPBgpPeerGroupsGetReader is a Reader for the NetworkIPBgpPeerGroupsGet structure.
type NetworkIPBgpPeerGroupsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NetworkIPBgpPeerGroupsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNetworkIPBgpPeerGroupsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewNetworkIPBgpPeerGroupsGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNetworkIPBgpPeerGroupsGetOK creates a NetworkIPBgpPeerGroupsGetOK with default headers values
func NewNetworkIPBgpPeerGroupsGetOK() *NetworkIPBgpPeerGroupsGetOK {
	return &NetworkIPBgpPeerGroupsGetOK{}
}

/* NetworkIPBgpPeerGroupsGetOK describes a response with status code 200, with default header values.

OK
*/
type NetworkIPBgpPeerGroupsGetOK struct {
	Payload *models.BgpPeerGroupResponse
}

func (o *NetworkIPBgpPeerGroupsGetOK) Error() string {
	return fmt.Sprintf("[GET /network/ip/bgp/peer-groups][%d] networkIpBgpPeerGroupsGetOK  %+v", 200, o.Payload)
}
func (o *NetworkIPBgpPeerGroupsGetOK) GetPayload() *models.BgpPeerGroupResponse {
	return o.Payload
}

func (o *NetworkIPBgpPeerGroupsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BgpPeerGroupResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNetworkIPBgpPeerGroupsGetDefault creates a NetworkIPBgpPeerGroupsGetDefault with default headers values
func NewNetworkIPBgpPeerGroupsGetDefault(code int) *NetworkIPBgpPeerGroupsGetDefault {
	return &NetworkIPBgpPeerGroupsGetDefault{
		_statusCode: code,
	}
}

/* NetworkIPBgpPeerGroupsGetDefault describes a response with status code -1, with default header values.

Error
*/
type NetworkIPBgpPeerGroupsGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the network ip bgp peer groups get default response
func (o *NetworkIPBgpPeerGroupsGetDefault) Code() int {
	return o._statusCode
}

func (o *NetworkIPBgpPeerGroupsGetDefault) Error() string {
	return fmt.Sprintf("[GET /network/ip/bgp/peer-groups][%d] network_ip_bgp_peer_groups_get default  %+v", o._statusCode, o.Payload)
}
func (o *NetworkIPBgpPeerGroupsGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *NetworkIPBgpPeerGroupsGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
