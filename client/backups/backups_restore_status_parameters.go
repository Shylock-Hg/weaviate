//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2025 Weaviate B.V. All rights reserved.
//
//  CONTACT: hello@weaviate.io
//

// Code generated by go-swagger; DO NOT EDIT.

package backups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewBackupsRestoreStatusParams creates a new BackupsRestoreStatusParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewBackupsRestoreStatusParams() *BackupsRestoreStatusParams {
	return &BackupsRestoreStatusParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewBackupsRestoreStatusParamsWithTimeout creates a new BackupsRestoreStatusParams object
// with the ability to set a timeout on a request.
func NewBackupsRestoreStatusParamsWithTimeout(timeout time.Duration) *BackupsRestoreStatusParams {
	return &BackupsRestoreStatusParams{
		timeout: timeout,
	}
}

// NewBackupsRestoreStatusParamsWithContext creates a new BackupsRestoreStatusParams object
// with the ability to set a context for a request.
func NewBackupsRestoreStatusParamsWithContext(ctx context.Context) *BackupsRestoreStatusParams {
	return &BackupsRestoreStatusParams{
		Context: ctx,
	}
}

// NewBackupsRestoreStatusParamsWithHTTPClient creates a new BackupsRestoreStatusParams object
// with the ability to set a custom HTTPClient for a request.
func NewBackupsRestoreStatusParamsWithHTTPClient(client *http.Client) *BackupsRestoreStatusParams {
	return &BackupsRestoreStatusParams{
		HTTPClient: client,
	}
}

/*
BackupsRestoreStatusParams contains all the parameters to send to the API endpoint

	for the backups restore status operation.

	Typically these are written to a http.Request.
*/
type BackupsRestoreStatusParams struct {

	/* Backend.

	   Backup backend name e.g. `filesystem`, `gcs`, `s3`, `azure`.
	*/
	Backend string

	/* Bucket.

	   Name of the bucket, container, volume, etc
	*/
	Bucket *string

	/* ID.

	   The ID of a backup. Must be URL-safe and work as a filesystem path, only lowercase, numbers, underscore, minus characters allowed.
	*/
	ID string

	/* Path.

	   The path within the bucket
	*/
	Path *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the backups restore status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BackupsRestoreStatusParams) WithDefaults() *BackupsRestoreStatusParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the backups restore status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BackupsRestoreStatusParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the backups restore status params
func (o *BackupsRestoreStatusParams) WithTimeout(timeout time.Duration) *BackupsRestoreStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the backups restore status params
func (o *BackupsRestoreStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the backups restore status params
func (o *BackupsRestoreStatusParams) WithContext(ctx context.Context) *BackupsRestoreStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the backups restore status params
func (o *BackupsRestoreStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the backups restore status params
func (o *BackupsRestoreStatusParams) WithHTTPClient(client *http.Client) *BackupsRestoreStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the backups restore status params
func (o *BackupsRestoreStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBackend adds the backend to the backups restore status params
func (o *BackupsRestoreStatusParams) WithBackend(backend string) *BackupsRestoreStatusParams {
	o.SetBackend(backend)
	return o
}

// SetBackend adds the backend to the backups restore status params
func (o *BackupsRestoreStatusParams) SetBackend(backend string) {
	o.Backend = backend
}

// WithBucket adds the bucket to the backups restore status params
func (o *BackupsRestoreStatusParams) WithBucket(bucket *string) *BackupsRestoreStatusParams {
	o.SetBucket(bucket)
	return o
}

// SetBucket adds the bucket to the backups restore status params
func (o *BackupsRestoreStatusParams) SetBucket(bucket *string) {
	o.Bucket = bucket
}

// WithID adds the id to the backups restore status params
func (o *BackupsRestoreStatusParams) WithID(id string) *BackupsRestoreStatusParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the backups restore status params
func (o *BackupsRestoreStatusParams) SetID(id string) {
	o.ID = id
}

// WithPath adds the path to the backups restore status params
func (o *BackupsRestoreStatusParams) WithPath(path *string) *BackupsRestoreStatusParams {
	o.SetPath(path)
	return o
}

// SetPath adds the path to the backups restore status params
func (o *BackupsRestoreStatusParams) SetPath(path *string) {
	o.Path = path
}

// WriteToRequest writes these params to a swagger request
func (o *BackupsRestoreStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param backend
	if err := r.SetPathParam("backend", o.Backend); err != nil {
		return err
	}

	if o.Bucket != nil {

		// query param bucket
		var qrBucket string

		if o.Bucket != nil {
			qrBucket = *o.Bucket
		}
		qBucket := qrBucket
		if qBucket != "" {

			if err := r.SetQueryParam("bucket", qBucket); err != nil {
				return err
			}
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if o.Path != nil {

		// query param path
		var qrPath string

		if o.Path != nil {
			qrPath = *o.Path
		}
		qPath := qrPath
		if qPath != "" {

			if err := r.SetQueryParam("path", qPath); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
