package permissions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-swagger/go-swagger/httpkit"

	"permissions/models"
)

/*BySubjectAndResourceOK OK

swagger:response bySubjectAndResourceOK
*/
type BySubjectAndResourceOK struct {

	// In: body
	Payload *models.PermissionList `json:"body,omitempty"`
}

// NewBySubjectAndResourceOK creates BySubjectAndResourceOK with default headers values
func NewBySubjectAndResourceOK() *BySubjectAndResourceOK {
	return &BySubjectAndResourceOK{}
}

// WithPayload adds the payload to the by subject and resource o k response
func (o *BySubjectAndResourceOK) WithPayload(payload *models.PermissionList) *BySubjectAndResourceOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the by subject and resource o k response
func (o *BySubjectAndResourceOK) SetPayload(payload *models.PermissionList) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *BySubjectAndResourceOK) WriteResponse(rw http.ResponseWriter, producer httpkit.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*BySubjectAndResourceBadRequest Bad Request

swagger:response bySubjectAndResourceBadRequest
*/
type BySubjectAndResourceBadRequest struct {

	// In: body
	Payload *models.ErrorOut `json:"body,omitempty"`
}

// NewBySubjectAndResourceBadRequest creates BySubjectAndResourceBadRequest with default headers values
func NewBySubjectAndResourceBadRequest() *BySubjectAndResourceBadRequest {
	return &BySubjectAndResourceBadRequest{}
}

// WithPayload adds the payload to the by subject and resource bad request response
func (o *BySubjectAndResourceBadRequest) WithPayload(payload *models.ErrorOut) *BySubjectAndResourceBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the by subject and resource bad request response
func (o *BySubjectAndResourceBadRequest) SetPayload(payload *models.ErrorOut) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *BySubjectAndResourceBadRequest) WriteResponse(rw http.ResponseWriter, producer httpkit.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*BySubjectAndResourceInternalServerError by subject and resource internal server error

swagger:response bySubjectAndResourceInternalServerError
*/
type BySubjectAndResourceInternalServerError struct {

	// In: body
	Payload *models.ErrorOut `json:"body,omitempty"`
}

// NewBySubjectAndResourceInternalServerError creates BySubjectAndResourceInternalServerError with default headers values
func NewBySubjectAndResourceInternalServerError() *BySubjectAndResourceInternalServerError {
	return &BySubjectAndResourceInternalServerError{}
}

// WithPayload adds the payload to the by subject and resource internal server error response
func (o *BySubjectAndResourceInternalServerError) WithPayload(payload *models.ErrorOut) *BySubjectAndResourceInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the by subject and resource internal server error response
func (o *BySubjectAndResourceInternalServerError) SetPayload(payload *models.ErrorOut) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *BySubjectAndResourceInternalServerError) WriteResponse(rw http.ResponseWriter, producer httpkit.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
