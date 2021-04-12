/*
 * UsagiBooru Accounts API
 *
 * アカウント関連API
 *
 * API version: 2.0
 * Contact: dsgamer777@gmail.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gen

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// A AccountsApiController binds http requests to an api service and writes the service results to the http response
type AccountsApiController struct {
	service AccountsApiServicer
}

// NewAccountsApiController creates a default api controller
func NewAccountsApiController(s AccountsApiServicer) Router {
	return &AccountsApiController{service: s}
}

// Routes returns all of the api route for the AccountsApiController
func (c *AccountsApiController) Routes() Routes {
	return Routes{
		{
			"CreateAccount",
			strings.ToUpper("Post"),
			"/accounts",
			c.CreateAccount,
		},
		{
			"DeleteAccount",
			strings.ToUpper("Delete"),
			"/accounts/{accountID}",
			c.DeleteAccount,
		},
		{
			"EditAccount",
			strings.ToUpper("Patch"),
			"/accounts/{accountID}",
			c.EditAccount,
		},
		{
			"GetAccountMe",
			strings.ToUpper("Get"),
			"/accounts/me",
			c.GetAccountMe,
		},
		{
			"GetAccount",
			strings.ToUpper("Get"),
			"/accounts/{accountID}",
			c.GetAccount,
		},
		{
			"GetUploadHistory",
			strings.ToUpper("Get"),
			"/accounts/{accountID}/upload_history",
			c.GetUploadHistory,
		},
		{
			"LoginWithForm",
			strings.ToUpper("Post"),
			"/accounts/login/form",
			c.LoginWithForm,
		},
	}
}

// CreateAccount - Create account
func (c *AccountsApiController) CreateAccount(w http.ResponseWriter, r *http.Request) {
	accountStruct := &AccountStruct{}
	if err := json.NewDecoder(r.Body).Decode(&accountStruct); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	result, err := c.service.CreateAccount(r.Context(), *accountStruct)
	//If an error occured, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	//If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// DeleteAccount - Delete account info
func (c *AccountsApiController) DeleteAccount(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	accountID, err := parseInt32Parameter(params["accountID"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	password := r.Header.Get("password")
	result, err := c.service.DeleteAccount(r.Context(), accountID, password)
	//If an error occured, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	//If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// EditAccount - Edit account info
func (c *AccountsApiController) EditAccount(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	accountID, err := parseInt32Parameter(params["accountID"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	accountStruct := &AccountStruct{}
	if err := json.NewDecoder(r.Body).Decode(&accountStruct); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	result, err := c.service.EditAccount(r.Context(), accountID, *accountStruct)
	//If an error occured, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	//If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// GetAccount - Get account info
func (c *AccountsApiController) GetAccount(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	accountID, err := parseInt32Parameter(params["accountID"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	result, err := c.service.GetAccount(r.Context(), accountID)
	//If an error occured, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	//If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// GetAccountMe - Get user info (self)
func (c *AccountsApiController) GetAccountMe(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.GetAccountMe(r.Context())
	//If an error occured, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	//If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// GetUploadHistory - Get upload history
func (c *AccountsApiController) GetUploadHistory(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	query := r.URL.Query()
	accountID, err := parseInt32Parameter(params["accountID"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	page, err := parseInt32Parameter(query.Get("page"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	sort := query.Get("sort")
	order := query.Get("order")
	perPage, err := parseInt32Parameter(query.Get("per_page"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	result, err := c.service.GetUploadHistory(r.Context(), accountID, page, sort, order, perPage)
	//If an error occured, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	//If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// LoginWithForm - Login with form
func (c *AccountsApiController) LoginWithForm(w http.ResponseWriter, r *http.Request) {
	postLoginWithFormRequest := &PostLoginWithFormRequest{}
	if err := json.NewDecoder(r.Body).Decode(&postLoginWithFormRequest); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	result, err := c.service.LoginWithForm(r.Context(), *postLoginWithFormRequest)
	//If an error occured, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	//If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}
