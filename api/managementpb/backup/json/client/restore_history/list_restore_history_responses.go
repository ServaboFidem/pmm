// Code generated by go-swagger; DO NOT EDIT.

package restore_history

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ListRestoreHistoryReader is a Reader for the ListRestoreHistory structure.
type ListRestoreHistoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListRestoreHistoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListRestoreHistoryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListRestoreHistoryDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListRestoreHistoryOK creates a ListRestoreHistoryOK with default headers values
func NewListRestoreHistoryOK() *ListRestoreHistoryOK {
	return &ListRestoreHistoryOK{}
}

/*ListRestoreHistoryOK handles this case with default header values.

A successful response.
*/
type ListRestoreHistoryOK struct {
	Payload *ListRestoreHistoryOKBody
}

func (o *ListRestoreHistoryOK) Error() string {
	return fmt.Sprintf("[POST /v1/management/backup/RestoreHistory/List][%d] listRestoreHistoryOk  %+v", 200, o.Payload)
}

func (o *ListRestoreHistoryOK) GetPayload() *ListRestoreHistoryOKBody {
	return o.Payload
}

func (o *ListRestoreHistoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ListRestoreHistoryOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListRestoreHistoryDefault creates a ListRestoreHistoryDefault with default headers values
func NewListRestoreHistoryDefault(code int) *ListRestoreHistoryDefault {
	return &ListRestoreHistoryDefault{
		_statusCode: code,
	}
}

/*ListRestoreHistoryDefault handles this case with default header values.

An unexpected error response.
*/
type ListRestoreHistoryDefault struct {
	_statusCode int

	Payload *ListRestoreHistoryDefaultBody
}

// Code gets the status code for the list restore history default response
func (o *ListRestoreHistoryDefault) Code() int {
	return o._statusCode
}

func (o *ListRestoreHistoryDefault) Error() string {
	return fmt.Sprintf("[POST /v1/management/backup/RestoreHistory/List][%d] ListRestoreHistory default  %+v", o._statusCode, o.Payload)
}

func (o *ListRestoreHistoryDefault) GetPayload() *ListRestoreHistoryDefaultBody {
	return o.Payload
}

func (o *ListRestoreHistoryDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ListRestoreHistoryDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*DetailsItems0 `Any` contains an arbitrary serialized protocol buffer message along with a
// URL that describes the type of the serialized message.
//
// Protobuf library provides support to pack/unpack Any values in the form
// of utility functions or additional generated methods of the Any type.
//
// Example 1: Pack and unpack a message in C++.
//
//     Foo foo = ...;
//     Any any;
//     any.PackFrom(foo);
//     ...
//     if (any.UnpackTo(&foo)) {
//       ...
//     }
//
// Example 2: Pack and unpack a message in Java.
//
//     Foo foo = ...;
//     Any any = Any.pack(foo);
//     ...
//     if (any.is(Foo.class)) {
//       foo = any.unpack(Foo.class);
//     }
//
//  Example 3: Pack and unpack a message in Python.
//
//     foo = Foo(...)
//     any = Any()
//     any.Pack(foo)
//     ...
//     if any.Is(Foo.DESCRIPTOR):
//       any.Unpack(foo)
//       ...
//
//  Example 4: Pack and unpack a message in Go
//
//      foo := &pb.Foo{...}
//      any, err := ptypes.MarshalAny(foo)
//      ...
//      foo := &pb.Foo{}
//      if err := ptypes.UnmarshalAny(any, foo); err != nil {
//        ...
//      }
//
// The pack methods provided by protobuf library will by default use
// 'type.googleapis.com/full.type.name' as the type URL and the unpack
// methods only use the fully qualified type name after the last '/'
// in the type URL, for example "foo.bar.com/x/y.z" will yield type
// name "y.z".
//
//
// JSON
// ====
// The JSON representation of an `Any` value uses the regular
// representation of the deserialized, embedded message, with an
// additional field `@type` which contains the type URL. Example:
//
//     package google.profile;
//     message Person {
//       string first_name = 1;
//       string last_name = 2;
//     }
//
//     {
//       "@type": "type.googleapis.com/google.profile.Person",
//       "firstName": <string>,
//       "lastName": <string>
//     }
//
// If the embedded message type is well-known and has a custom JSON
// representation, that representation will be embedded adding a field
// `value` which holds the custom JSON in addition to the `@type`
// field. Example (for message [google.protobuf.Duration][]):
//
//     {
//       "@type": "type.googleapis.com/google.protobuf.Duration",
//       "value": "1.212s"
//     }
swagger:model DetailsItems0
*/
type DetailsItems0 struct {

	// A URL/resource name that uniquely identifies the type of the serialized
	// protocol buffer message. This string must contain at least
	// one "/" character. The last segment of the URL's path must represent
	// the fully qualified name of the type (as in
	// `path/google.protobuf.Duration`). The name should be in a canonical form
	// (e.g., leading "." is not accepted).
	//
	// In practice, teams usually precompile into the binary all types that they
	// expect it to use in the context of Any. However, for URLs which use the
	// scheme `http`, `https`, or no scheme, one can optionally set up a type
	// server that maps type URLs to message definitions as follows:
	//
	// * If no scheme is provided, `https` is assumed.
	// * An HTTP GET on the URL must yield a [google.protobuf.Type][]
	//   value in binary format, or produce an error.
	// * Applications are allowed to cache lookup results based on the
	//   URL, or have them precompiled into a binary to avoid any
	//   lookup. Therefore, binary compatibility needs to be preserved
	//   on changes to types. (Use versioned type names to manage
	//   breaking changes.)
	//
	// Note: this functionality is not currently available in the official
	// protobuf release, and it is not used for type URLs beginning with
	// type.googleapis.com.
	//
	// Schemes other than `http`, `https` (or the empty scheme) might be
	// used with implementation specific semantics.
	TypeURL string `json:"type_url,omitempty"`

	// Must be a valid serialized protocol buffer of the above specified type.
	// Format: byte
	Value strfmt.Base64 `json:"value,omitempty"`
}

// Validate validates this details items0
func (o *DetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DetailsItems0) UnmarshalBinary(b []byte) error {
	var res DetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ItemsItems0 RestoreHistoryItem represents single backup restore item.
swagger:model ItemsItems0
*/
type ItemsItems0 struct {

	// Machine-readable restore id.
	RestoreID string `json:"restore_id,omitempty"`

	// ID of the artifact used for restore.
	ArtifactID string `json:"artifact_id,omitempty"`

	// Artifact name used for restore.
	Name string `json:"name,omitempty"`

	// Database vendor e.g. PostgreSQL, MongoDB, MySQL.
	Vendor string `json:"vendor,omitempty"`

	// Machine-readable location ID.
	LocationID string `json:"location_id,omitempty"`

	// Location name.
	LocationName string `json:"location_name,omitempty"`

	// Machine-readable service ID.
	ServiceID string `json:"service_id,omitempty"`

	// Service name.
	ServiceName string `json:"service_name,omitempty"`

	// DataModel is a model used for performing a backup.
	// Enum: [DATA_MODEL_INVALID PHYSICAL LOGICAL]
	DataModel *string `json:"data_model,omitempty"`

	// RestoreStatus shows the current status of execution of restore.
	// Enum: [RESTORE_STATUS_INVALID RESTORE_STATUS_IN_PROGRESS RESTORE_STATUS_SUCCESS RESTORE_STATUS_ERROR]
	Status *string `json:"status,omitempty"`

	// Restore start time.
	// Format: date-time
	StartedAt strfmt.DateTime `json:"started_at,omitempty"`

	// Restore finish time.
	// Format: date-time
	FinishedAt strfmt.DateTime `json:"finished_at,omitempty"`
}

// Validate validates this items items0
func (o *ItemsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDataModel(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateStartedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateFinishedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var itemsItems0TypeDataModelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["DATA_MODEL_INVALID","PHYSICAL","LOGICAL"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		itemsItems0TypeDataModelPropEnum = append(itemsItems0TypeDataModelPropEnum, v)
	}
}

const (

	// ItemsItems0DataModelDATAMODELINVALID captures enum value "DATA_MODEL_INVALID"
	ItemsItems0DataModelDATAMODELINVALID string = "DATA_MODEL_INVALID"

	// ItemsItems0DataModelPHYSICAL captures enum value "PHYSICAL"
	ItemsItems0DataModelPHYSICAL string = "PHYSICAL"

	// ItemsItems0DataModelLOGICAL captures enum value "LOGICAL"
	ItemsItems0DataModelLOGICAL string = "LOGICAL"
)

// prop value enum
func (o *ItemsItems0) validateDataModelEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, itemsItems0TypeDataModelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *ItemsItems0) validateDataModel(formats strfmt.Registry) error {

	if swag.IsZero(o.DataModel) { // not required
		return nil
	}

	// value enum
	if err := o.validateDataModelEnum("data_model", "body", *o.DataModel); err != nil {
		return err
	}

	return nil
}

var itemsItems0TypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["RESTORE_STATUS_INVALID","RESTORE_STATUS_IN_PROGRESS","RESTORE_STATUS_SUCCESS","RESTORE_STATUS_ERROR"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		itemsItems0TypeStatusPropEnum = append(itemsItems0TypeStatusPropEnum, v)
	}
}

const (

	// ItemsItems0StatusRESTORESTATUSINVALID captures enum value "RESTORE_STATUS_INVALID"
	ItemsItems0StatusRESTORESTATUSINVALID string = "RESTORE_STATUS_INVALID"

	// ItemsItems0StatusRESTORESTATUSINPROGRESS captures enum value "RESTORE_STATUS_IN_PROGRESS"
	ItemsItems0StatusRESTORESTATUSINPROGRESS string = "RESTORE_STATUS_IN_PROGRESS"

	// ItemsItems0StatusRESTORESTATUSSUCCESS captures enum value "RESTORE_STATUS_SUCCESS"
	ItemsItems0StatusRESTORESTATUSSUCCESS string = "RESTORE_STATUS_SUCCESS"

	// ItemsItems0StatusRESTORESTATUSERROR captures enum value "RESTORE_STATUS_ERROR"
	ItemsItems0StatusRESTORESTATUSERROR string = "RESTORE_STATUS_ERROR"
)

// prop value enum
func (o *ItemsItems0) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, itemsItems0TypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *ItemsItems0) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(o.Status) { // not required
		return nil
	}

	// value enum
	if err := o.validateStatusEnum("status", "body", *o.Status); err != nil {
		return err
	}

	return nil
}

func (o *ItemsItems0) validateStartedAt(formats strfmt.Registry) error {

	if swag.IsZero(o.StartedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("started_at", "body", "date-time", o.StartedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *ItemsItems0) validateFinishedAt(formats strfmt.Registry) error {

	if swag.IsZero(o.FinishedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("finished_at", "body", "date-time", o.FinishedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ItemsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ItemsItems0) UnmarshalBinary(b []byte) error {
	var res ItemsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ListRestoreHistoryDefaultBody list restore history default body
swagger:model ListRestoreHistoryDefaultBody
*/
type ListRestoreHistoryDefaultBody struct {

	// error
	Error string `json:"error,omitempty"`

	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*DetailsItems0 `json:"details"`
}

// Validate validates this list restore history default body
func (o *ListRestoreHistoryDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListRestoreHistoryDefaultBody) validateDetails(formats strfmt.Registry) error {

	if swag.IsZero(o.Details) { // not required
		return nil
	}

	for i := 0; i < len(o.Details); i++ {
		if swag.IsZero(o.Details[i]) { // not required
			continue
		}

		if o.Details[i] != nil {
			if err := o.Details[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ListRestoreHistory default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *ListRestoreHistoryDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListRestoreHistoryDefaultBody) UnmarshalBinary(b []byte) error {
	var res ListRestoreHistoryDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ListRestoreHistoryOKBody list restore history OK body
swagger:model ListRestoreHistoryOKBody
*/
type ListRestoreHistoryOKBody struct {

	// items
	Items []*ItemsItems0 `json:"items"`
}

// Validate validates this list restore history OK body
func (o *ListRestoreHistoryOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateItems(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListRestoreHistoryOKBody) validateItems(formats strfmt.Registry) error {

	if swag.IsZero(o.Items) { // not required
		return nil
	}

	for i := 0; i < len(o.Items); i++ {
		if swag.IsZero(o.Items[i]) { // not required
			continue
		}

		if o.Items[i] != nil {
			if err := o.Items[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("listRestoreHistoryOk" + "." + "items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *ListRestoreHistoryOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListRestoreHistoryOKBody) UnmarshalBinary(b []byte) error {
	var res ListRestoreHistoryOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
