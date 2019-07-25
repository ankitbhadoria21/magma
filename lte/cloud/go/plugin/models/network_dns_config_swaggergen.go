// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NetworkDNSConfig DNS configuration for a network
// swagger:model networkDnsConfig
type NetworkDNSConfig struct {

	// enable caching
	// Required: true
	EnableCaching *bool `json:"enable_caching"`

	// local ttl
	// Required: true
	LocalTTL *int32 `json:"local_ttl"`

	// records
	Records []*NetworkDNSConfigRecordsItems0 `json:"records"`
}

// Validate validates this network Dns config
func (m *NetworkDNSConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnableCaching(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocalTTL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecords(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NetworkDNSConfig) validateEnableCaching(formats strfmt.Registry) error {

	if err := validate.Required("enable_caching", "body", m.EnableCaching); err != nil {
		return err
	}

	return nil
}

func (m *NetworkDNSConfig) validateLocalTTL(formats strfmt.Registry) error {

	if err := validate.Required("local_ttl", "body", m.LocalTTL); err != nil {
		return err
	}

	return nil
}

func (m *NetworkDNSConfig) validateRecords(formats strfmt.Registry) error {

	if swag.IsZero(m.Records) { // not required
		return nil
	}

	for i := 0; i < len(m.Records); i++ {
		if swag.IsZero(m.Records[i]) { // not required
			continue
		}

		if m.Records[i] != nil {
			if err := m.Records[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *NetworkDNSConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetworkDNSConfig) UnmarshalBinary(b []byte) error {
	var res NetworkDNSConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// NetworkDNSConfigRecordsItems0 Mapping used for DNS resolving from a domain
// swagger:model NetworkDNSConfigRecordsItems0
type NetworkDNSConfigRecordsItems0 struct {

	// a record
	ARecord []string `json:"a_record"`

	// aaaa record
	AaaaRecord []string `json:"aaaa_record"`

	// cname record
	CnameRecord []string `json:"cname_record"`

	// domain
	// Required: true
	// Min Length: 1
	Domain string `json:"domain"`
}

// Validate validates this network DNS config records items0
func (m *NetworkDNSConfigRecordsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateARecord(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAaaaRecord(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCnameRecord(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDomain(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NetworkDNSConfigRecordsItems0) validateARecord(formats strfmt.Registry) error {

	if swag.IsZero(m.ARecord) { // not required
		return nil
	}

	for i := 0; i < len(m.ARecord); i++ {

		if err := validate.MinLength("a_record"+"."+strconv.Itoa(i), "body", string(m.ARecord[i]), 1); err != nil {
			return err
		}

	}

	return nil
}

func (m *NetworkDNSConfigRecordsItems0) validateAaaaRecord(formats strfmt.Registry) error {

	if swag.IsZero(m.AaaaRecord) { // not required
		return nil
	}

	for i := 0; i < len(m.AaaaRecord); i++ {

		if err := validate.MinLength("aaaa_record"+"."+strconv.Itoa(i), "body", string(m.AaaaRecord[i]), 1); err != nil {
			return err
		}

	}

	return nil
}

func (m *NetworkDNSConfigRecordsItems0) validateCnameRecord(formats strfmt.Registry) error {

	if swag.IsZero(m.CnameRecord) { // not required
		return nil
	}

	for i := 0; i < len(m.CnameRecord); i++ {

		if err := validate.MinLength("cname_record"+"."+strconv.Itoa(i), "body", string(m.CnameRecord[i]), 1); err != nil {
			return err
		}

	}

	return nil
}

func (m *NetworkDNSConfigRecordsItems0) validateDomain(formats strfmt.Registry) error {

	if err := validate.RequiredString("domain", "body", string(m.Domain)); err != nil {
		return err
	}

	if err := validate.MinLength("domain", "body", string(m.Domain), 1); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NetworkDNSConfigRecordsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetworkDNSConfigRecordsItems0) UnmarshalBinary(b []byte) error {
	var res NetworkDNSConfigRecordsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
