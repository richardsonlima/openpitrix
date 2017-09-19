// Code generated by go-swagger; DO NOT EDIT.

package repos

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"AppHub/src/api/swagger/models"
)

// GetV1ReposRepoIDOKCode is the HTTP code returned for type GetV1ReposRepoIDOK
const GetV1ReposRepoIDOKCode int = 200

/*GetV1ReposRepoIDOK A repo

swagger:response getV1ReposRepoIdOK
*/
type GetV1ReposRepoIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.Repo `json:"body,omitempty"`
}

// NewGetV1ReposRepoIDOK creates GetV1ReposRepoIDOK with default headers values
func NewGetV1ReposRepoIDOK() *GetV1ReposRepoIDOK {
	return &GetV1ReposRepoIDOK{}
}

// WithPayload adds the payload to the get v1 repos repo Id o k response
func (o *GetV1ReposRepoIDOK) WithPayload(payload *models.Repo) *GetV1ReposRepoIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get v1 repos repo Id o k response
func (o *GetV1ReposRepoIDOK) SetPayload(payload *models.Repo) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetV1ReposRepoIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetV1ReposRepoIDNotFoundCode is the HTTP code returned for type GetV1ReposRepoIDNotFound
const GetV1ReposRepoIDNotFoundCode int = 404

/*GetV1ReposRepoIDNotFound The repo does not exist.

swagger:response getV1ReposRepoIdNotFound
*/
type GetV1ReposRepoIDNotFound struct {
}

// NewGetV1ReposRepoIDNotFound creates GetV1ReposRepoIDNotFound with default headers values
func NewGetV1ReposRepoIDNotFound() *GetV1ReposRepoIDNotFound {
	return &GetV1ReposRepoIDNotFound{}
}

// WriteResponse to the client
func (o *GetV1ReposRepoIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
}

// GetV1ReposRepoIDInternalServerErrorCode is the HTTP code returned for type GetV1ReposRepoIDInternalServerError
const GetV1ReposRepoIDInternalServerErrorCode int = 500

/*GetV1ReposRepoIDInternalServerError An unexpected error occured.

swagger:response getV1ReposRepoIdInternalServerError
*/
type GetV1ReposRepoIDInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetV1ReposRepoIDInternalServerError creates GetV1ReposRepoIDInternalServerError with default headers values
func NewGetV1ReposRepoIDInternalServerError() *GetV1ReposRepoIDInternalServerError {
	return &GetV1ReposRepoIDInternalServerError{}
}

// WithPayload adds the payload to the get v1 repos repo Id internal server error response
func (o *GetV1ReposRepoIDInternalServerError) WithPayload(payload *models.Error) *GetV1ReposRepoIDInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get v1 repos repo Id internal server error response
func (o *GetV1ReposRepoIDInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetV1ReposRepoIDInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}