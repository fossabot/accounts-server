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
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"github.com/gorilla/mux"
)

// A Route defines the parameters for an api endpoint
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes are a collection of defined api endpoints
type Routes []Route

// Router defines the required methods for retrieving api routes
type Router interface {
	Routes() Routes
}

// NewRouter creates a new router for any number of api routers
func NewRouter(routers ...Router) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, api := range routers {
		for _, route := range api.Routes() {
			var handler http.Handler
			handler = route.HandlerFunc
			handler = Logger(handler, route.Name)

			router.
				Methods(route.Method).
				Path(route.Pattern).
				Name(route.Name).
				Handler(handler)
		}
	}

	return router
}

// EncodeJSONResponse uses the json encoder to write an interface to the http response with an optional status code
func EncodeJSONResponse(i interface{}, status *int, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if status != nil {
		w.WriteHeader(*status)
	} else {
		w.WriteHeader(http.StatusOK)
	}

	return json.NewEncoder(w).Encode(i)
}

// ReadFormFileToTempFile reads file data from a request form and writes it to a temporary file
func ReadFormFileToTempFile(r *http.Request, key string) (*os.File, error) {
	_, fileHeader, err := r.FormFile(key)
	if err != nil {
		return nil, err
	}

	return readFileHeaderToTempFile(fileHeader)
}

// ReadFormFilesToTempFiles reads files array data from a request form and writes it to a temporary files
func ReadFormFilesToTempFiles(r *http.Request, key string) ([]*os.File, error) {
	if err := r.ParseMultipartForm(32 << 20); err != nil {
		return nil, err
	}

	files := make([]*os.File, 0, len(r.MultipartForm.File[key]))

	for _, fileHeader := range r.MultipartForm.File[key] {
		file, err := readFileHeaderToTempFile(fileHeader)
		if err != nil {
			return nil, err
		}

		files = append(files, file)
	}

	return files, nil
}

// readFileHeaderToTempFile reads multipart.FileHeader and writes it to a temporary file
func readFileHeaderToTempFile(fileHeader *multipart.FileHeader) (*os.File, error) {
	formFile, err := fileHeader.Open()
	if err != nil {
		return nil, err
	}

	defer formFile.Close()

	fileBytes, err := ioutil.ReadAll(formFile)
	if err != nil {
		return nil, err
	}

	file, err := ioutil.TempFile("", fileHeader.Filename)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	file.Write(fileBytes)

	return file, nil
}


// parseInt64Parameter parses a sting parameter to an int64
func parseInt64Parameter(param string) (int64, error) {
	return strconv.ParseInt(param, 10, 64)
}

// parseInt32Parameter parses a sting parameter to an int32
func parseInt32Parameter(param string) (int32, error) {
	val, err := strconv.ParseInt(param, 10, 32)
	if err != nil {
		return -1, err
	}
	return int32(val), nil
}
