// Code generated by go-swagger; DO NOT EDIT.

package device_events

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetDeviceEventParams creates a new GetDeviceEventParams object
// with the default values initialized.
func NewGetDeviceEventParams() *GetDeviceEventParams {
	var ()
	return &GetDeviceEventParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDeviceEventParamsWithTimeout creates a new GetDeviceEventParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDeviceEventParamsWithTimeout(timeout time.Duration) *GetDeviceEventParams {
	var ()
	return &GetDeviceEventParams{

		timeout: timeout,
	}
}

// NewGetDeviceEventParamsWithContext creates a new GetDeviceEventParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDeviceEventParamsWithContext(ctx context.Context) *GetDeviceEventParams {
	var ()
	return &GetDeviceEventParams{

		Context: ctx,
	}
}

// NewGetDeviceEventParamsWithHTTPClient creates a new GetDeviceEventParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDeviceEventParamsWithHTTPClient(client *http.Client) *GetDeviceEventParams {
	var ()
	return &GetDeviceEventParams{
		HTTPClient: client,
	}
}

/*GetDeviceEventParams contains all the parameters to send to the API endpoint
for the get device event operation typically these are written to a http.Request
*/
type GetDeviceEventParams struct {

	/*ID
	  The ID of the event to retrieve

	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get device event params
func (o *GetDeviceEventParams) WithTimeout(timeout time.Duration) *GetDeviceEventParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get device event params
func (o *GetDeviceEventParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get device event params
func (o *GetDeviceEventParams) WithContext(ctx context.Context) *GetDeviceEventParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get device event params
func (o *GetDeviceEventParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get device event params
func (o *GetDeviceEventParams) WithHTTPClient(client *http.Client) *GetDeviceEventParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get device event params
func (o *GetDeviceEventParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get device event params
func (o *GetDeviceEventParams) WithID(id strfmt.UUID) *GetDeviceEventParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get device event params
func (o *GetDeviceEventParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetDeviceEventParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
