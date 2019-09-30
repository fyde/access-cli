// Code generated by go-swagger; DO NOT EDIT.

package device_events

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/oNaiPs/fyde-cli/models"
)

// ListDeviceEventsReader is a Reader for the ListDeviceEvents structure.
type ListDeviceEventsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListDeviceEventsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListDeviceEventsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListDeviceEventsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListDeviceEventsOK creates a ListDeviceEventsOK with default headers values
func NewListDeviceEventsOK() *ListDeviceEventsOK {
	return &ListDeviceEventsOK{}
}

/*ListDeviceEventsOK handles this case with default header values.

successful operation
*/
type ListDeviceEventsOK struct {
	/*Total number of items (for pagination)
	 */
	Total int64

	Payload *ListDeviceEventsOKBodyTuple0
}

func (o *ListDeviceEventsOK) Error() string {
	return fmt.Sprintf("[GET /device_events][%d] listDeviceEventsOK  %+v", 200, o.Payload)
}

func (o *ListDeviceEventsOK) GetPayload() *ListDeviceEventsOKBodyTuple0 {
	return o.Payload
}

func (o *ListDeviceEventsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header total
	total, err := swag.ConvertInt64(response.GetHeader("total"))
	if err != nil {
		return errors.InvalidType("total", "header", "int64", response.GetHeader("total"))
	}
	o.Total = total

	o.Payload = new(ListDeviceEventsOKBodyTuple0)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListDeviceEventsUnauthorized creates a ListDeviceEventsUnauthorized with default headers values
func NewListDeviceEventsUnauthorized() *ListDeviceEventsUnauthorized {
	return &ListDeviceEventsUnauthorized{}
}

/*ListDeviceEventsUnauthorized handles this case with default header values.

unauthorized: invalid credentials or missing authentication headers
*/
type ListDeviceEventsUnauthorized struct {
	Payload *models.GenericUnauthorizedResponse
}

func (o *ListDeviceEventsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /device_events][%d] listDeviceEventsUnauthorized  %+v", 401, o.Payload)
}

func (o *ListDeviceEventsUnauthorized) GetPayload() *models.GenericUnauthorizedResponse {
	return o.Payload
}

func (o *ListDeviceEventsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericUnauthorizedResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ListDeviceEventsOKBodyTuple0 ListDeviceEventsOKBodyTuple0 a representation of an anonymous Tuple type
swagger:model ListDeviceEventsOKBodyTuple0
*/
type ListDeviceEventsOKBodyTuple0 struct {

	// p0
	// Required: true
	P0 *models.DeviceEventListItem `json:"-"` // custom serializer

}

// UnmarshalJSON unmarshals this tuple type from a JSON array
func (o *ListDeviceEventsOKBodyTuple0) UnmarshalJSON(raw []byte) error {
	// stage 1, get the array but just the array
	var stage1 []json.RawMessage
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&stage1); err != nil {
		return err
	}

	// stage 2: hydrates struct members with tuple elements

	if len(stage1) > 0 {
		var dataP0 models.DeviceEventListItem
		buf = bytes.NewBuffer(stage1[0])
		dec := json.NewDecoder(buf)
		dec.UseNumber()
		if err := dec.Decode(&dataP0); err != nil {
			return err
		}
		o.P0 = &dataP0

	}

	return nil
}

// MarshalJSON marshals this tuple type into a JSON array
func (o ListDeviceEventsOKBodyTuple0) MarshalJSON() ([]byte, error) {
	data := []interface{}{
		o.P0,
	}

	return json.Marshal(data)
}

// Validate validates this list device events o k body tuple0
func (o *ListDeviceEventsOKBodyTuple0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateP0(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListDeviceEventsOKBodyTuple0) validateP0(formats strfmt.Registry) error {

	if err := validate.Required("P0", "body", o.P0); err != nil {
		return err
	}

	if o.P0 != nil {
		if err := o.P0.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("P0")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ListDeviceEventsOKBodyTuple0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListDeviceEventsOKBodyTuple0) UnmarshalBinary(b []byte) error {
	var res ListDeviceEventsOKBodyTuple0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
