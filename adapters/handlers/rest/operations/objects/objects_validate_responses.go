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

package objects

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/weaviate/weaviate/entities/models"
)

// ObjectsValidateOKCode is the HTTP code returned for type ObjectsValidateOK
const ObjectsValidateOKCode int = 200

/*
ObjectsValidateOK Successfully validated.

swagger:response objectsValidateOK
*/
type ObjectsValidateOK struct {
}

// NewObjectsValidateOK creates ObjectsValidateOK with default headers values
func NewObjectsValidateOK() *ObjectsValidateOK {

	return &ObjectsValidateOK{}
}

// WriteResponse to the client
func (o *ObjectsValidateOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// ObjectsValidateUnauthorizedCode is the HTTP code returned for type ObjectsValidateUnauthorized
const ObjectsValidateUnauthorizedCode int = 401

/*
ObjectsValidateUnauthorized Unauthorized or invalid credentials.

swagger:response objectsValidateUnauthorized
*/
type ObjectsValidateUnauthorized struct {
}

// NewObjectsValidateUnauthorized creates ObjectsValidateUnauthorized with default headers values
func NewObjectsValidateUnauthorized() *ObjectsValidateUnauthorized {

	return &ObjectsValidateUnauthorized{}
}

// WriteResponse to the client
func (o *ObjectsValidateUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// ObjectsValidateForbiddenCode is the HTTP code returned for type ObjectsValidateForbidden
const ObjectsValidateForbiddenCode int = 403

/*
ObjectsValidateForbidden Forbidden

swagger:response objectsValidateForbidden
*/
type ObjectsValidateForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewObjectsValidateForbidden creates ObjectsValidateForbidden with default headers values
func NewObjectsValidateForbidden() *ObjectsValidateForbidden {

	return &ObjectsValidateForbidden{}
}

// WithPayload adds the payload to the objects validate forbidden response
func (o *ObjectsValidateForbidden) WithPayload(payload *models.ErrorResponse) *ObjectsValidateForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the objects validate forbidden response
func (o *ObjectsValidateForbidden) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ObjectsValidateForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ObjectsValidateUnprocessableEntityCode is the HTTP code returned for type ObjectsValidateUnprocessableEntity
const ObjectsValidateUnprocessableEntityCode int = 422

/*
ObjectsValidateUnprocessableEntity Request body is well-formed (i.e., syntactically correct), but semantically erroneous. Are you sure the class is defined in the configuration file?

swagger:response objectsValidateUnprocessableEntity
*/
type ObjectsValidateUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewObjectsValidateUnprocessableEntity creates ObjectsValidateUnprocessableEntity with default headers values
func NewObjectsValidateUnprocessableEntity() *ObjectsValidateUnprocessableEntity {

	return &ObjectsValidateUnprocessableEntity{}
}

// WithPayload adds the payload to the objects validate unprocessable entity response
func (o *ObjectsValidateUnprocessableEntity) WithPayload(payload *models.ErrorResponse) *ObjectsValidateUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the objects validate unprocessable entity response
func (o *ObjectsValidateUnprocessableEntity) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ObjectsValidateUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ObjectsValidateInternalServerErrorCode is the HTTP code returned for type ObjectsValidateInternalServerError
const ObjectsValidateInternalServerErrorCode int = 500

/*
ObjectsValidateInternalServerError An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.

swagger:response objectsValidateInternalServerError
*/
type ObjectsValidateInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewObjectsValidateInternalServerError creates ObjectsValidateInternalServerError with default headers values
func NewObjectsValidateInternalServerError() *ObjectsValidateInternalServerError {

	return &ObjectsValidateInternalServerError{}
}

// WithPayload adds the payload to the objects validate internal server error response
func (o *ObjectsValidateInternalServerError) WithPayload(payload *models.ErrorResponse) *ObjectsValidateInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the objects validate internal server error response
func (o *ObjectsValidateInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ObjectsValidateInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
