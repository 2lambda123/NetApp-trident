// Code generated by go-swagger; DO NOT EDIT.

package networking

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// NetworkIPInterfacesCreateReader is a Reader for the NetworkIPInterfacesCreate structure.
type NetworkIPInterfacesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NetworkIPInterfacesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewNetworkIPInterfacesCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewNetworkIPInterfacesCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNetworkIPInterfacesCreateCreated creates a NetworkIPInterfacesCreateCreated with default headers values
func NewNetworkIPInterfacesCreateCreated() *NetworkIPInterfacesCreateCreated {
	return &NetworkIPInterfacesCreateCreated{}
}

/* NetworkIPInterfacesCreateCreated describes a response with status code 201, with default header values.

Created
*/
type NetworkIPInterfacesCreateCreated struct {
}

func (o *NetworkIPInterfacesCreateCreated) Error() string {
	return fmt.Sprintf("[POST /network/ip/interfaces][%d] networkIpInterfacesCreateCreated ", 201)
}

func (o *NetworkIPInterfacesCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNetworkIPInterfacesCreateDefault creates a NetworkIPInterfacesCreateDefault with default headers values
func NewNetworkIPInterfacesCreateDefault(code int) *NetworkIPInterfacesCreateDefault {
	return &NetworkIPInterfacesCreateDefault{
		_statusCode: code,
	}
}

/* NetworkIPInterfacesCreateDefault describes a response with status code -1, with default header values.

 ONTAP Error Response Codes
| Error Code | Description |
| ---------- | ----------- |
| 1376656 | Cluster interfaces must be in the same subnet. Verify the address and netmask are set to the correct values. |
| 1376663 | All LIFs from a single DNS zone must be in the same SVM. |
| 1376663 | Cannot add interface to DNS zone because all interfaces from a single DNS zone must be in the same SVM. |
| 1376963 | Duplicate IP address. |
| 1966138 | The same IP address may not be used for both a mgmt interface and a gateway address. |
| 1966140 | An interface with the same name already exists. |
| 1966141 | Invalid DNS zone name. |
| 1966142 | Only data LIFs can be assigned a DNS zone. |
| 1966267 | IPv6 addresses must have a prefix length between 1 and 127. |
| 1966269 | IPv4 addresses must have a prefix length between 1 and 32. |
| 1966270 | Operation not support on SAN LIFs. |
| 1966476 | DNS Update is supported only on data LIFs. |
| 1966477 | DNS Update is supported only on LIFs configured with the NFS or CIFS protocol. |
| 1966987 | The Vserver Broadcast-Domain Home-Node and Home-Port combination is not valid. |
| 1967081 | The specified SVM must exist in the specified IPspace. |
| 1967082 | The specified ipspace.name does not match the IPspace name of ipspace.uuid. |
| 1967102 | POST operation might have left configuration in an inconsistent state. Check the configuration. |
| 1967106 | The specified location.home_port.name does not match the specified port name of location.home_port.uuid. |
| 1967107 | The location.home_port.uuid specified is not valid. |
| 1967108 | The specified location.home_node.name does not match the node name of location.home_node.uuid. |
| 1967109 | The specified location.home_port.node.name does not match the node name of location.home_node.uuid. |
| 1967110 | The specified location.home_port.node.name does not match location.home_node.name. |
| 1967111 | Home node must be specified by at least one location.home_node, location.home_port, or location.broadcast_domain field. |
| 1967112 | The specified location.home_node.name does not match the node name of location.home_port.uuid. |
| 1967120 | The specified service_policy.name does not match the specified service policy name of service_policy.uuid. |
| 1967121 | Invalid service_policy.uuid specified. |
| 1967122 | The specified location.broadcast_domain.name does not match the specified broadcast domain name of location.broadcast_domain.uuid. |
| 1967123 | The specified IPspace does not match the IPspace name of location.broadcast_domain.uuid. |
| 1967124 | The location.broadcast_domain.uuid specified is not valid. |
| 1967127 | svm.uuid or svm.name must be provided if scope is "svm". |
| 1967128 | ipspace.uuid or ipspace.name must be provided if scope is "cluster". |
| 1967129 | The specified location.home_port.uuid is not valid. |
| 1967130 | The specified location.home_port.name is not valid. |
| 1967131 | The specified location.home_port.uuid and location.home_port.name are not valid. |
| 1967135 | The specified location.broadcast_domain.uuid is not valid. |
| 1967136 | The specified location.broadcast_domain.name (and ipspace name) is not valid. |
| 1967137 | The specified location.broadcast_domain.uuid and location.broadcast_domain.name (and IPspace name) are not valid. |
| 1967145 | The specified location.failover is not valid. |
| 1967146 | The specified svm.name is not valid. |
| 1967147 | The specified svm.uuid is not valid. |
| 1967153 | No suitable port exists on location.home_node to host the interface. |
| 1967154 | Interfaces cannot be created on ports that are down. If a broadcast domain is specified, ensure that it contains at least one port that is operationally up. |
| 1967381 | Post VIP interfaces requires an effective cluster version of 9.7 or later. |
| 1967382 | VIP interfaces only reside in SVM scope. |
| 1967383 | Neither location.home_port.uuid or location.home_port.name should be set with vip=true. |
| 1967384 | Failed to create VIP interface because the home node does not have active BGP sessions to support Virtual IP (VIP) traffic. |
| 1967385 | VIP interfaces with an IPv4 address must use ip.netmask=32. VIP interfaces with an IPv6 address must use ip.netmask=128. |
| 1967387 | The specified IP address is in use by a subnet in this IPspace. |
| 1967391 | Setting the DNS zone requires an effective cluster version of 9.9.1 or later. |
| 1967392 | Setting the DDNS enable parameter requires an effective cluster version of 9.9.1 or later. |
| 1967394 | Setting the probe port parameter requires an effective cluster version of 9.10.1 or later. |
| 5373966 | An iSCSI interface cannot be created in an SVM configured for NVMe. |
| 53281036 | Setting the probe port parameter is not allowed on this platform. |
| 53281065 | The service_policy does not exist in the SVM. |
| 53281086 | LIF would exceed the maximum number of supported intercluster LIFs in IPspace. |
| 53281087 | Cannot configure SAN LIF on SVM. |

*/
type NetworkIPInterfacesCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the network ip interfaces create default response
func (o *NetworkIPInterfacesCreateDefault) Code() int {
	return o._statusCode
}

func (o *NetworkIPInterfacesCreateDefault) Error() string {
	return fmt.Sprintf("[POST /network/ip/interfaces][%d] network_ip_interfaces_create default  %+v", o._statusCode, o.Payload)
}
func (o *NetworkIPInterfacesCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *NetworkIPInterfacesCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
