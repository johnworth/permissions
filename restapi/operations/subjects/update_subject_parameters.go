package subjects

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit"
	"github.com/go-swagger/go-swagger/httpkit/middleware"

	strfmt "github.com/go-swagger/go-swagger/strfmt"

	"permissions/models"
)

// NewUpdateSubjectParams creates a new UpdateSubjectParams object
// with the default values initialized.
func NewUpdateSubjectParams() UpdateSubjectParams {
	var ()
	return UpdateSubjectParams{}
}

// UpdateSubjectParams contains all the bound params for the update subject operation
// typically these are obtained from a http.Request
//
// swagger:parameters updateSubject
type UpdateSubjectParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request

	/*The subject ID.
	  Required: true
	  In: path
	*/
	ID string
	/*The new subject information.
	  Required: true
	  In: body
	*/
	SubjectIn *models.SubjectIn
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *UpdateSubjectParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	rID, rhkID, _ := route.Params.GetOK("id")
	if err := o.bindID(rID, rhkID, route.Formats); err != nil {
		res = append(res, err)
	}

	if httpkit.HasBody(r) {
		defer r.Body.Close()
		var body models.SubjectIn
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("subjectIn", "body"))
			} else {
				res = append(res, errors.NewParseError("subjectIn", "body", "", err))
			}

		} else {
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.SubjectIn = &body
			}
		}

	} else {
		res = append(res, errors.Required("subjectIn", "body"))
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateSubjectParams) bindID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.ID = raw

	return nil
}
