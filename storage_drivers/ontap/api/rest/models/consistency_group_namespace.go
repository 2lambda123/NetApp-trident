// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ConsistencyGroupNamespace An NVMe namespace is a collection of addressable logical blocks presented to hosts connected to the storage virtual machine using the NVMe over Fabrics protocol.<br/>
// In ONTAP, an NVMe namespace is located within a volume. Optionally, it can be located within a qtree in a volume.<br/>
// An NVMe namespace is created to a specified size using thin or thick provisioning as determined by the volume on which it is created. NVMe namespaces support being cloned. An NVMe namespace cannot be renamed, resized, or moved to a different volume. NVMe namespaces do not support the assignment of a QoS policy for performance management, but a QoS policy can be assigned to the volume containing the namespace. See the NVMe namespace object model to learn more about each of the properties supported by the NVMe namespace REST API.<br/>
// An NVMe namespace must be mapped to an NVMe subsystem to grant access to the subsystem's hosts. Hosts can then access the NVMe namespace and perform I/O using the NVMe over Fabrics protocol.
//
//
// swagger:model consistency_group_namespace
type ConsistencyGroupNamespace struct {

	// This property marks the NVMe namespace for auto deletion when the volume containing the namespace runs out of space. This is most commonly set on namespace clones.<br/>
	// When set to _true_, the NVMe namespace becomes eligible for automatic deletion when the volume runs out of space. Auto deletion only occurs when the volume containing the namespace is also configured for auto deletion and free space in the volume decreases below a particular threshold.<br/>
	// This property is optional in POST and PATCH. The default value for a new NVMe namespace is _false_.<br/>
	// There is an added cost to retrieving this property's value. It is not populated for either a collection GET or an instance GET unless it is explicitly requested using the `fields` query parameter. See [`Requesting specific fields`](#Requesting_specific_fields) to learn more.
	//
	AutoDelete *bool `json:"auto_delete,omitempty"`

	// A configurable comment available for use by the administrator. Valid in POST and PATCH.
	//
	// Max Length: 254
	// Min Length: 0
	Comment *string `json:"comment,omitempty"`

	// The time the NVMe namespace was created.
	// Example: 2018-06-04T19:00:00Z
	// Read Only: true
	// Format: date-time
	CreateTime *strfmt.DateTime `json:"create_time,omitempty"`

	// The enabled state of the NVMe namespace. Certain error conditions cause the namespace to become disabled. If the namespace is disabled, you can check the `state` property to determine what error disabled the namespace. An NVMe namespace is enabled automatically when it is created.
	//
	// Read Only: true
	Enabled *bool `json:"enabled,omitempty"`

	// The fully qualified path name of the NVMe namespace composed of a "/vol" prefix, the volume name, the (optional) qtree name and base name of the namespace. Valid in POST.<br/>
	// NVMe namespaces do not support rename, or movement between volumes.
	//
	// Example: /vol/volume1/qtree1/namespace1
	Name string `json:"name,omitempty"`

	// The operating system type of the NVMe namespace.<br/>
	// Required in POST when creating an NVMe namespace that is not a clone of another. Disallowed in POST when creating a namespace clone.
	//
	// Enum: [aix linux vmware windows]
	OsType string `json:"os_type,omitempty"`

	// provisioning options
	ProvisioningOptions *ConsistencyGroupNamespaceProvisioningOptions `json:"provisioning_options,omitempty"`

	// The NVMe subsystem with which the NVMe namespace is associated. A namespace can be mapped to zero (0) or one (1) subsystems.<br/>
	// There is an added cost to retrieving property values for `subsystem_map`. They are not populated for either a collection GET or an instance GET unless explicitly requested using the `fields` query parameter.
	//
	SubsystemMap []*ConsistencyGroupNamespaceSubsystemMapItems0 `json:"subsystem_map,omitempty"`

	// The unique identifier of the NVMe namespace.
	//
	// Example: 1cd8a442-86d1-11e0-ae1c-123478563412
	// Read Only: true
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this consistency group namespace
func (m *ConsistencyGroupNamespace) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateComment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOsType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProvisioningOptions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubsystemMap(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConsistencyGroupNamespace) validateComment(formats strfmt.Registry) error {
	if swag.IsZero(m.Comment) { // not required
		return nil
	}

	if err := validate.MinLength("comment", "body", *m.Comment, 0); err != nil {
		return err
	}

	if err := validate.MaxLength("comment", "body", *m.Comment, 254); err != nil {
		return err
	}

	return nil
}

func (m *ConsistencyGroupNamespace) validateCreateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.CreateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("create_time", "body", "date-time", m.CreateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

var consistencyGroupNamespaceTypeOsTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["aix","linux","vmware","windows"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		consistencyGroupNamespaceTypeOsTypePropEnum = append(consistencyGroupNamespaceTypeOsTypePropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// consistency_group_namespace
	// ConsistencyGroupNamespace
	// os_type
	// OsType
	// aix
	// END DEBUGGING
	// ConsistencyGroupNamespaceOsTypeAix captures enum value "aix"
	ConsistencyGroupNamespaceOsTypeAix string = "aix"

	// BEGIN DEBUGGING
	// consistency_group_namespace
	// ConsistencyGroupNamespace
	// os_type
	// OsType
	// linux
	// END DEBUGGING
	// ConsistencyGroupNamespaceOsTypeLinux captures enum value "linux"
	ConsistencyGroupNamespaceOsTypeLinux string = "linux"

	// BEGIN DEBUGGING
	// consistency_group_namespace
	// ConsistencyGroupNamespace
	// os_type
	// OsType
	// vmware
	// END DEBUGGING
	// ConsistencyGroupNamespaceOsTypeVmware captures enum value "vmware"
	ConsistencyGroupNamespaceOsTypeVmware string = "vmware"

	// BEGIN DEBUGGING
	// consistency_group_namespace
	// ConsistencyGroupNamespace
	// os_type
	// OsType
	// windows
	// END DEBUGGING
	// ConsistencyGroupNamespaceOsTypeWindows captures enum value "windows"
	ConsistencyGroupNamespaceOsTypeWindows string = "windows"
)

// prop value enum
func (m *ConsistencyGroupNamespace) validateOsTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, consistencyGroupNamespaceTypeOsTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ConsistencyGroupNamespace) validateOsType(formats strfmt.Registry) error {
	if swag.IsZero(m.OsType) { // not required
		return nil
	}

	// value enum
	if err := m.validateOsTypeEnum("os_type", "body", m.OsType); err != nil {
		return err
	}

	return nil
}

func (m *ConsistencyGroupNamespace) validateProvisioningOptions(formats strfmt.Registry) error {
	if swag.IsZero(m.ProvisioningOptions) { // not required
		return nil
	}

	if m.ProvisioningOptions != nil {
		if err := m.ProvisioningOptions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("provisioning_options")
			}
			return err
		}
	}

	return nil
}

func (m *ConsistencyGroupNamespace) validateSubsystemMap(formats strfmt.Registry) error {
	if swag.IsZero(m.SubsystemMap) { // not required
		return nil
	}

	for i := 0; i < len(m.SubsystemMap); i++ {
		if swag.IsZero(m.SubsystemMap[i]) { // not required
			continue
		}

		if m.SubsystemMap[i] != nil {
			if err := m.SubsystemMap[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("subsystem_map" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this consistency group namespace based on the context it is used
func (m *ConsistencyGroupNamespace) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCreateTime(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEnabled(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProvisioningOptions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSubsystemMap(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUUID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConsistencyGroupNamespace) contextValidateCreateTime(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "create_time", "body", m.CreateTime); err != nil {
		return err
	}

	return nil
}

func (m *ConsistencyGroupNamespace) contextValidateEnabled(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "enabled", "body", m.Enabled); err != nil {
		return err
	}

	return nil
}

func (m *ConsistencyGroupNamespace) contextValidateProvisioningOptions(ctx context.Context, formats strfmt.Registry) error {

	if m.ProvisioningOptions != nil {
		if err := m.ProvisioningOptions.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("provisioning_options")
			}
			return err
		}
	}

	return nil
}

func (m *ConsistencyGroupNamespace) contextValidateSubsystemMap(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.SubsystemMap); i++ {

		if m.SubsystemMap[i] != nil {
			if err := m.SubsystemMap[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("subsystem_map" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ConsistencyGroupNamespace) contextValidateUUID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "uuid", "body", string(m.UUID)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConsistencyGroupNamespace) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConsistencyGroupNamespace) UnmarshalBinary(b []byte) error {
	var res ConsistencyGroupNamespace
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ConsistencyGroupNamespaceProvisioningOptions Options that are applied to the operation.
//
// swagger:model ConsistencyGroupNamespaceProvisioningOptions
type ConsistencyGroupNamespaceProvisioningOptions struct {

	// Operation to perform
	// Enum: [create]
	Action string `json:"action,omitempty"`

	// Number of elements to perform the operation on.
	Count int64 `json:"count,omitempty"`
}

// Validate validates this consistency group namespace provisioning options
func (m *ConsistencyGroupNamespaceProvisioningOptions) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAction(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var consistencyGroupNamespaceProvisioningOptionsTypeActionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["create"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		consistencyGroupNamespaceProvisioningOptionsTypeActionPropEnum = append(consistencyGroupNamespaceProvisioningOptionsTypeActionPropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// ConsistencyGroupNamespaceProvisioningOptions
	// ConsistencyGroupNamespaceProvisioningOptions
	// action
	// Action
	// create
	// END DEBUGGING
	// ConsistencyGroupNamespaceProvisioningOptionsActionCreate captures enum value "create"
	ConsistencyGroupNamespaceProvisioningOptionsActionCreate string = "create"
)

// prop value enum
func (m *ConsistencyGroupNamespaceProvisioningOptions) validateActionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, consistencyGroupNamespaceProvisioningOptionsTypeActionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ConsistencyGroupNamespaceProvisioningOptions) validateAction(formats strfmt.Registry) error {
	if swag.IsZero(m.Action) { // not required
		return nil
	}

	// value enum
	if err := m.validateActionEnum("provisioning_options"+"."+"action", "body", m.Action); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this consistency group namespace provisioning options based on context it is used
func (m *ConsistencyGroupNamespaceProvisioningOptions) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ConsistencyGroupNamespaceProvisioningOptions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConsistencyGroupNamespaceProvisioningOptions) UnmarshalBinary(b []byte) error {
	var res ConsistencyGroupNamespaceProvisioningOptions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ConsistencyGroupNamespaceSubsystemMapItems0 The NVMe subsystem with which the NVMe namespace is associated. A namespace can be mapped to zero (0) or one (1) subsystems.<br/>
// There is an added cost to retrieving property values for `subsystem_map`.
// They are not populated for either a collection GET or an instance GET unless explicitly requested using the `fields` query parameter.
//
//
// swagger:model ConsistencyGroupNamespaceSubsystemMapItems0
type ConsistencyGroupNamespaceSubsystemMapItems0 struct {

	// links
	Links *SelfLink `json:"_links,omitempty"`

	// The Asymmetric Namespace Access Group ID (ANAGRPID) of the NVMe namespace.<br/>
	// The format for an ANAGRPID is 8 hexadecimal digits (zero-filled) followed by a lower case "h".
	//
	// Example: 00103050h
	// Read Only: true
	Anagrpid string `json:"anagrpid,omitempty"`

	// The NVMe namespace identifier. This is an identifier used by an NVMe controller to provide access to the NVMe namespace.<br/>
	// The format for an NVMe namespace identifier is 8 hexadecimal digits (zero-filled) followed by a lower case "h".
	//
	// Example: 00000001h
	// Read Only: true
	Nsid string `json:"nsid,omitempty"`

	// The NVMe subsystem to which the NVMe namespace is mapped.
	//
	// Read Only: true
	Subsystem *NvmeSubsystemReference `json:"subsystem,omitempty"`
}

// Validate validates this consistency group namespace subsystem map items0
func (m *ConsistencyGroupNamespaceSubsystemMapItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubsystem(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConsistencyGroupNamespaceSubsystemMapItems0) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links")
			}
			return err
		}
	}

	return nil
}

func (m *ConsistencyGroupNamespaceSubsystemMapItems0) validateSubsystem(formats strfmt.Registry) error {
	if swag.IsZero(m.Subsystem) { // not required
		return nil
	}

	if m.Subsystem != nil {
		if err := m.Subsystem.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("subsystem")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this consistency group namespace subsystem map items0 based on the context it is used
func (m *ConsistencyGroupNamespaceSubsystemMapItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAnagrpid(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNsid(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSubsystem(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConsistencyGroupNamespaceSubsystemMapItems0) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links")
			}
			return err
		}
	}

	return nil
}

func (m *ConsistencyGroupNamespaceSubsystemMapItems0) contextValidateAnagrpid(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "anagrpid", "body", string(m.Anagrpid)); err != nil {
		return err
	}

	return nil
}

func (m *ConsistencyGroupNamespaceSubsystemMapItems0) contextValidateNsid(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "nsid", "body", string(m.Nsid)); err != nil {
		return err
	}

	return nil
}

func (m *ConsistencyGroupNamespaceSubsystemMapItems0) contextValidateSubsystem(ctx context.Context, formats strfmt.Registry) error {

	if m.Subsystem != nil {
		if err := m.Subsystem.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("subsystem")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConsistencyGroupNamespaceSubsystemMapItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConsistencyGroupNamespaceSubsystemMapItems0) UnmarshalBinary(b []byte) error {
	var res ConsistencyGroupNamespaceSubsystemMapItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
