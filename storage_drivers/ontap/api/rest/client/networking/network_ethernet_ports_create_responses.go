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

// NetworkEthernetPortsCreateReader is a Reader for the NetworkEthernetPortsCreate structure.
type NetworkEthernetPortsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NetworkEthernetPortsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewNetworkEthernetPortsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewNetworkEthernetPortsCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNetworkEthernetPortsCreateCreated creates a NetworkEthernetPortsCreateCreated with default headers values
func NewNetworkEthernetPortsCreateCreated() *NetworkEthernetPortsCreateCreated {
	return &NetworkEthernetPortsCreateCreated{}
}

/* NetworkEthernetPortsCreateCreated describes a response with status code 201, with default header values.

Created
*/
type NetworkEthernetPortsCreateCreated struct {
	Payload *models.PortResponse
}

func (o *NetworkEthernetPortsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /network/ethernet/ports][%d] networkEthernetPortsCreateCreated  %+v", 201, o.Payload)
}
func (o *NetworkEthernetPortsCreateCreated) GetPayload() *models.PortResponse {
	return o.Payload
}

func (o *NetworkEthernetPortsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PortResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNetworkEthernetPortsCreateDefault creates a NetworkEthernetPortsCreateDefault with default headers values
func NewNetworkEthernetPortsCreateDefault(code int) *NetworkEthernetPortsCreateDefault {
	return &NetworkEthernetPortsCreateDefault{
		_statusCode: code,
	}
}

/* NetworkEthernetPortsCreateDefault describes a response with status code -1, with default header values.

 ONTAP Error Response Codes
| Error Code | Description |
| ---------- | ----------- |
| 1376361 | Port is already a member of a LAG. |
| 1966189 | Port is the home port or current port of an interface. |
| 1967083 | The specified type is not valid. |
| 1967084 | The specified node UUID is not valid. |
| 1967085 | The specified node name is not valid. |
| 1967086 | Node name and UUID must match if both are provided. |
| 1967087 | The specified broadcast domain UUID is not valid. |
| 1967088 | The specified broadcast domain name does not exist in the specified IPspace. |
| 1967089 | The specified broadcast domain UUID, name, and IPspace name do not match. |
| 1967090 | The specified VLAN base port UUID is not valid. |
| 1967091 | The specified VLAN base port name and node name are not valid. |
| 1967092 | The specified node does not match the node specified for the VLAN base port. |
| 1967093 | The specified VLAN base port UUID, name, and VLAN base port node name do not match. |
| 1967094 | The specified LAG member port UUID is not valid. |
| 1967095 | The specified LAG member port name and node name combination is not valid. |
| 1967096 | The specified node does not match the specified LAG member port node. |
| 1967097 | The specified LAG member ports UUID, name, and node name do not match. |
| 1967098 | VLAN POST operation has failed because admin status could not be set for the specified port. |
| 1967099 | Partial success of the VLAN POST operation. Verify the state of the created VLAN for more information. |
| 1967100 | LAG POST operation failed because admin status could not be set. |
| 1967101 | Partial success of the LAG POST operation. Verify the state of the created LAG for more information. |
| 1967102 | POST operation might have left the configuration in an inconsistent state. Check the configuration. |
| 1967148 | Failure to remove port from broadcast domain. |
| 1967149 | Failure to add port to broadcast domain. |
| 1967175 | VLANs cannot be created on ports in the Cluster IPspace. |

*/
type NetworkEthernetPortsCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the network ethernet ports create default response
func (o *NetworkEthernetPortsCreateDefault) Code() int {
	return o._statusCode
}

func (o *NetworkEthernetPortsCreateDefault) Error() string {
	return fmt.Sprintf("[POST /network/ethernet/ports][%d] network_ethernet_ports_create default  %+v", o._statusCode, o.Payload)
}
func (o *NetworkEthernetPortsCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *NetworkEthernetPortsCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
