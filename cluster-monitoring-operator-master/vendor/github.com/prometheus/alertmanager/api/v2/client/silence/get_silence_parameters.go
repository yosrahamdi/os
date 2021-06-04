// Code generated by go-swagger; DO NOT EDIT.

// Copyright Prometheus Team
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package silence

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

// NewGetSilenceParams creates a new GetSilenceParams object
// with the default values initialized.
func NewGetSilenceParams() *GetSilenceParams {
	var ()
	return &GetSilenceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSilenceParamsWithTimeout creates a new GetSilenceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSilenceParamsWithTimeout(timeout time.Duration) *GetSilenceParams {
	var ()
	return &GetSilenceParams{

		timeout: timeout,
	}
}

// NewGetSilenceParamsWithContext creates a new GetSilenceParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSilenceParamsWithContext(ctx context.Context) *GetSilenceParams {
	var ()
	return &GetSilenceParams{

		Context: ctx,
	}
}

// NewGetSilenceParamsWithHTTPClient creates a new GetSilenceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSilenceParamsWithHTTPClient(client *http.Client) *GetSilenceParams {
	var ()
	return &GetSilenceParams{
		HTTPClient: client,
	}
}

/*GetSilenceParams contains all the parameters to send to the API endpoint
for the get silence operation typically these are written to a http.Request
*/
type GetSilenceParams struct {

	/*SilenceID
	  ID of the silence to get

	*/
	SilenceID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get silence params
func (o *GetSilenceParams) WithTimeout(timeout time.Duration) *GetSilenceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get silence params
func (o *GetSilenceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get silence params
func (o *GetSilenceParams) WithContext(ctx context.Context) *GetSilenceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get silence params
func (o *GetSilenceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get silence params
func (o *GetSilenceParams) WithHTTPClient(client *http.Client) *GetSilenceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get silence params
func (o *GetSilenceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSilenceID adds the silenceID to the get silence params
func (o *GetSilenceParams) WithSilenceID(silenceID strfmt.UUID) *GetSilenceParams {
	o.SetSilenceID(silenceID)
	return o
}

// SetSilenceID adds the silenceId to the get silence params
func (o *GetSilenceParams) SetSilenceID(silenceID strfmt.UUID) {
	o.SilenceID = silenceID
}

// WriteToRequest writes these params to a swagger request
func (o *GetSilenceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param silenceID
	if err := r.SetPathParam("silenceID", o.SilenceID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
