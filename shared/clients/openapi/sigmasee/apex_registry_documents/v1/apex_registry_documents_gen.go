// Package apexregistrydocumentsv1 provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.16.2 DO NOT EDIT.
package apexregistrydocumentsv1

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"strings"
	"time"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/oapi-codegen/runtime"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

// Document defines model for Document.
type Document struct {
	DocumentDate *time.Time          `json:"documentDate,omitempty"`
	DocumentId   *openapi_types.UUID `json:"documentId,omitempty"`
	FileName     *string             `json:"fileName"`
	InvestmentId *openapi_types.UUID `json:"investmentId"`
	InvestorId   *openapi_types.UUID `json:"investorId"`
	Location     *string             `json:"location"`
	Metadata     *interface{}        `json:"metadata"`
}

// DocumentsResponse defines model for DocumentsResponse.
type DocumentsResponse struct {
	Items *[]Document `json:"items"`
}

// SystemIOStream defines model for System.IO.Stream.
type SystemIOStream struct {
	CanRead      *bool  `json:"canRead,omitempty"`
	CanSeek      *bool  `json:"canSeek,omitempty"`
	CanTimeout   *bool  `json:"canTimeout,omitempty"`
	CanWrite     *bool  `json:"canWrite,omitempty"`
	Length       *int64 `json:"length,omitempty"`
	Position     *int64 `json:"position,omitempty"`
	ReadTimeout  *int32 `json:"readTimeout,omitempty"`
	WriteTimeout *int32 `json:"writeTimeout,omitempty"`
}

// GetInvestmentsDocumentsParams defines parameters for GetInvestmentsDocuments.
type GetInvestmentsDocumentsParams struct {
	FromDate *time.Time `form:"fromDate,omitempty" json:"fromDate,omitempty"`
	ToDate   *time.Time `form:"toDate,omitempty" json:"toDate,omitempty"`
}

// GetInvestmentDocumentsParams defines parameters for GetInvestmentDocuments.
type GetInvestmentDocumentsParams struct {
	FromDate *time.Time `form:"fromDate,omitempty" json:"fromDate,omitempty"`
	ToDate   *time.Time `form:"toDate,omitempty" json:"toDate,omitempty"`
}

// GetInvestorsDocumentsParams defines parameters for GetInvestorsDocuments.
type GetInvestorsDocumentsParams struct {
	IncludeInvestmentDocuments *bool      `form:"includeInvestmentDocuments,omitempty" json:"includeInvestmentDocuments,omitempty"`
	FromDate                   *time.Time `form:"fromDate,omitempty" json:"fromDate,omitempty"`
	ToDate                     *time.Time `form:"toDate,omitempty" json:"toDate,omitempty"`
}

// GetInvestorDocumentsParams defines parameters for GetInvestorDocuments.
type GetInvestorDocumentsParams struct {
	IncludeInvestmentDocuments *bool      `form:"includeInvestmentDocuments,omitempty" json:"includeInvestmentDocuments,omitempty"`
	FromDate                   *time.Time `form:"fromDate,omitempty" json:"fromDate,omitempty"`
	ToDate                     *time.Time `form:"toDate,omitempty" json:"toDate,omitempty"`
}

// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A list of callbacks for modifying requests which are generated before sending over
	// the network.
	RequestEditors []RequestEditorFn
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(server string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Server: server,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the server URL always has a trailing slash
	if !strings.HasSuffix(client.Server, "/") {
		client.Server += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = &http.Client{}
	}
	return &client, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithRequestEditorFn allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestEditorFn(fn RequestEditorFn) ClientOption {
	return func(c *Client) error {
		c.RequestEditors = append(c.RequestEditors, fn)
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {
	// GetInvestmentsDocuments request
	GetInvestmentsDocuments(ctx context.Context, params *GetInvestmentsDocumentsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetInvestmentDocuments request
	GetInvestmentDocuments(ctx context.Context, id openapi_types.UUID, params *GetInvestmentDocumentsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetInvestmentDocument request
	GetInvestmentDocument(ctx context.Context, id openapi_types.UUID, documentId openapi_types.UUID, fileName string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetInvestorsDocuments request
	GetInvestorsDocuments(ctx context.Context, params *GetInvestorsDocumentsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetInvestorDocuments request
	GetInvestorDocuments(ctx context.Context, id openapi_types.UUID, params *GetInvestorDocumentsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetInvestorDocument request
	GetInvestorDocument(ctx context.Context, id openapi_types.UUID, documentId openapi_types.UUID, fileName string, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) GetInvestmentsDocuments(ctx context.Context, params *GetInvestmentsDocumentsParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetInvestmentsDocumentsRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetInvestmentDocuments(ctx context.Context, id openapi_types.UUID, params *GetInvestmentDocumentsParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetInvestmentDocumentsRequest(c.Server, id, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetInvestmentDocument(ctx context.Context, id openapi_types.UUID, documentId openapi_types.UUID, fileName string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetInvestmentDocumentRequest(c.Server, id, documentId, fileName)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetInvestorsDocuments(ctx context.Context, params *GetInvestorsDocumentsParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetInvestorsDocumentsRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetInvestorDocuments(ctx context.Context, id openapi_types.UUID, params *GetInvestorDocumentsParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetInvestorDocumentsRequest(c.Server, id, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetInvestorDocument(ctx context.Context, id openapi_types.UUID, documentId openapi_types.UUID, fileName string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetInvestorDocumentRequest(c.Server, id, documentId, fileName)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewGetInvestmentsDocumentsRequest generates requests for GetInvestmentsDocuments
func NewGetInvestmentsDocumentsRequest(server string, params *GetInvestmentsDocumentsParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/investments/documents")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		queryValues := queryURL.Query()

		if params.FromDate != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "fromDate", runtime.ParamLocationQuery, *params.FromDate); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}

		}

		if params.ToDate != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "toDate", runtime.ParamLocationQuery, *params.ToDate); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}

		}

		queryURL.RawQuery = queryValues.Encode()
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetInvestmentDocumentsRequest generates requests for GetInvestmentDocuments
func NewGetInvestmentDocumentsRequest(server string, id openapi_types.UUID, params *GetInvestmentDocumentsParams) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/investments/%s/documents", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		queryValues := queryURL.Query()

		if params.FromDate != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "fromDate", runtime.ParamLocationQuery, *params.FromDate); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}

		}

		if params.ToDate != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "toDate", runtime.ParamLocationQuery, *params.ToDate); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}

		}

		queryURL.RawQuery = queryValues.Encode()
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetInvestmentDocumentRequest generates requests for GetInvestmentDocument
func NewGetInvestmentDocumentRequest(server string, id openapi_types.UUID, documentId openapi_types.UUID, fileName string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParamWithLocation("simple", false, "documentId", runtime.ParamLocationPath, documentId)
	if err != nil {
		return nil, err
	}

	var pathParam2 string

	pathParam2, err = runtime.StyleParamWithLocation("simple", false, "fileName", runtime.ParamLocationPath, fileName)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/investments/%s/documents/%s/%s", pathParam0, pathParam1, pathParam2)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetInvestorsDocumentsRequest generates requests for GetInvestorsDocuments
func NewGetInvestorsDocumentsRequest(server string, params *GetInvestorsDocumentsParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/investors/documents")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		queryValues := queryURL.Query()

		if params.IncludeInvestmentDocuments != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "includeInvestmentDocuments", runtime.ParamLocationQuery, *params.IncludeInvestmentDocuments); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}

		}

		if params.FromDate != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "fromDate", runtime.ParamLocationQuery, *params.FromDate); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}

		}

		if params.ToDate != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "toDate", runtime.ParamLocationQuery, *params.ToDate); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}

		}

		queryURL.RawQuery = queryValues.Encode()
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetInvestorDocumentsRequest generates requests for GetInvestorDocuments
func NewGetInvestorDocumentsRequest(server string, id openapi_types.UUID, params *GetInvestorDocumentsParams) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/investors/%s/documents", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		queryValues := queryURL.Query()

		if params.IncludeInvestmentDocuments != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "includeInvestmentDocuments", runtime.ParamLocationQuery, *params.IncludeInvestmentDocuments); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}

		}

		if params.FromDate != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "fromDate", runtime.ParamLocationQuery, *params.FromDate); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}

		}

		if params.ToDate != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "toDate", runtime.ParamLocationQuery, *params.ToDate); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}

		}

		queryURL.RawQuery = queryValues.Encode()
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetInvestorDocumentRequest generates requests for GetInvestorDocument
func NewGetInvestorDocumentRequest(server string, id openapi_types.UUID, documentId openapi_types.UUID, fileName string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParamWithLocation("simple", false, "documentId", runtime.ParamLocationPath, documentId)
	if err != nil {
		return nil, err
	}

	var pathParam2 string

	pathParam2, err = runtime.StyleParamWithLocation("simple", false, "fileName", runtime.ParamLocationPath, fileName)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/investors/%s/documents/%s/%s", pathParam0, pathParam1, pathParam2)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (c *Client) applyEditors(ctx context.Context, req *http.Request, additionalEditors []RequestEditorFn) error {
	for _, r := range c.RequestEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	for _, r := range additionalEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	return nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(server string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(server, opts...)
	if err != nil {
		return nil, err
	}
	return &ClientWithResponses{client}, nil
}

// WithBaseURL overrides the baseURL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.Server = newBaseURL.String()
		return nil
	}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {
	// GetInvestmentsDocumentsWithResponse request
	GetInvestmentsDocumentsWithResponse(ctx context.Context, params *GetInvestmentsDocumentsParams, reqEditors ...RequestEditorFn) (*GetInvestmentsDocumentsResponse, error)

	// GetInvestmentDocumentsWithResponse request
	GetInvestmentDocumentsWithResponse(ctx context.Context, id openapi_types.UUID, params *GetInvestmentDocumentsParams, reqEditors ...RequestEditorFn) (*GetInvestmentDocumentsResponse, error)

	// GetInvestmentDocumentWithResponse request
	GetInvestmentDocumentWithResponse(ctx context.Context, id openapi_types.UUID, documentId openapi_types.UUID, fileName string, reqEditors ...RequestEditorFn) (*GetInvestmentDocumentResponse, error)

	// GetInvestorsDocumentsWithResponse request
	GetInvestorsDocumentsWithResponse(ctx context.Context, params *GetInvestorsDocumentsParams, reqEditors ...RequestEditorFn) (*GetInvestorsDocumentsResponse, error)

	// GetInvestorDocumentsWithResponse request
	GetInvestorDocumentsWithResponse(ctx context.Context, id openapi_types.UUID, params *GetInvestorDocumentsParams, reqEditors ...RequestEditorFn) (*GetInvestorDocumentsResponse, error)

	// GetInvestorDocumentWithResponse request
	GetInvestorDocumentWithResponse(ctx context.Context, id openapi_types.UUID, documentId openapi_types.UUID, fileName string, reqEditors ...RequestEditorFn) (*GetInvestorDocumentResponse, error)
}

type GetInvestmentsDocumentsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *DocumentsResponse
}

// Status returns HTTPResponse.Status
func (r GetInvestmentsDocumentsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetInvestmentsDocumentsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetInvestmentDocumentsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *DocumentsResponse
}

// Status returns HTTPResponse.Status
func (r GetInvestmentDocumentsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetInvestmentDocumentsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetInvestmentDocumentResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *SystemIOStream
}

// Status returns HTTPResponse.Status
func (r GetInvestmentDocumentResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetInvestmentDocumentResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetInvestorsDocumentsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *DocumentsResponse
}

// Status returns HTTPResponse.Status
func (r GetInvestorsDocumentsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetInvestorsDocumentsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetInvestorDocumentsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *DocumentsResponse
}

// Status returns HTTPResponse.Status
func (r GetInvestorDocumentsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetInvestorDocumentsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetInvestorDocumentResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *SystemIOStream
}

// Status returns HTTPResponse.Status
func (r GetInvestorDocumentResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetInvestorDocumentResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// GetInvestmentsDocumentsWithResponse request returning *GetInvestmentsDocumentsResponse
func (c *ClientWithResponses) GetInvestmentsDocumentsWithResponse(ctx context.Context, params *GetInvestmentsDocumentsParams, reqEditors ...RequestEditorFn) (*GetInvestmentsDocumentsResponse, error) {
	rsp, err := c.GetInvestmentsDocuments(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetInvestmentsDocumentsResponse(rsp)
}

// GetInvestmentDocumentsWithResponse request returning *GetInvestmentDocumentsResponse
func (c *ClientWithResponses) GetInvestmentDocumentsWithResponse(ctx context.Context, id openapi_types.UUID, params *GetInvestmentDocumentsParams, reqEditors ...RequestEditorFn) (*GetInvestmentDocumentsResponse, error) {
	rsp, err := c.GetInvestmentDocuments(ctx, id, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetInvestmentDocumentsResponse(rsp)
}

// GetInvestmentDocumentWithResponse request returning *GetInvestmentDocumentResponse
func (c *ClientWithResponses) GetInvestmentDocumentWithResponse(ctx context.Context, id openapi_types.UUID, documentId openapi_types.UUID, fileName string, reqEditors ...RequestEditorFn) (*GetInvestmentDocumentResponse, error) {
	rsp, err := c.GetInvestmentDocument(ctx, id, documentId, fileName, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetInvestmentDocumentResponse(rsp)
}

// GetInvestorsDocumentsWithResponse request returning *GetInvestorsDocumentsResponse
func (c *ClientWithResponses) GetInvestorsDocumentsWithResponse(ctx context.Context, params *GetInvestorsDocumentsParams, reqEditors ...RequestEditorFn) (*GetInvestorsDocumentsResponse, error) {
	rsp, err := c.GetInvestorsDocuments(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetInvestorsDocumentsResponse(rsp)
}

// GetInvestorDocumentsWithResponse request returning *GetInvestorDocumentsResponse
func (c *ClientWithResponses) GetInvestorDocumentsWithResponse(ctx context.Context, id openapi_types.UUID, params *GetInvestorDocumentsParams, reqEditors ...RequestEditorFn) (*GetInvestorDocumentsResponse, error) {
	rsp, err := c.GetInvestorDocuments(ctx, id, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetInvestorDocumentsResponse(rsp)
}

// GetInvestorDocumentWithResponse request returning *GetInvestorDocumentResponse
func (c *ClientWithResponses) GetInvestorDocumentWithResponse(ctx context.Context, id openapi_types.UUID, documentId openapi_types.UUID, fileName string, reqEditors ...RequestEditorFn) (*GetInvestorDocumentResponse, error) {
	rsp, err := c.GetInvestorDocument(ctx, id, documentId, fileName, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetInvestorDocumentResponse(rsp)
}

// ParseGetInvestmentsDocumentsResponse parses an HTTP response from a GetInvestmentsDocumentsWithResponse call
func ParseGetInvestmentsDocumentsResponse(rsp *http.Response) (*GetInvestmentsDocumentsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetInvestmentsDocumentsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest DocumentsResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case rsp.StatusCode == 200:
		// Content-type (text/plain) unsupported

	}

	return response, nil
}

// ParseGetInvestmentDocumentsResponse parses an HTTP response from a GetInvestmentDocumentsWithResponse call
func ParseGetInvestmentDocumentsResponse(rsp *http.Response) (*GetInvestmentDocumentsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetInvestmentDocumentsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest DocumentsResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case rsp.StatusCode == 200:
		// Content-type (text/plain) unsupported

	}

	return response, nil
}

// ParseGetInvestmentDocumentResponse parses an HTTP response from a GetInvestmentDocumentWithResponse call
func ParseGetInvestmentDocumentResponse(rsp *http.Response) (*GetInvestmentDocumentResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetInvestmentDocumentResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest SystemIOStream
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case rsp.StatusCode == 200:
		// Content-type (text/plain) unsupported

	}

	return response, nil
}

// ParseGetInvestorsDocumentsResponse parses an HTTP response from a GetInvestorsDocumentsWithResponse call
func ParseGetInvestorsDocumentsResponse(rsp *http.Response) (*GetInvestorsDocumentsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetInvestorsDocumentsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest DocumentsResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case rsp.StatusCode == 200:
		// Content-type (text/plain) unsupported

	}

	return response, nil
}

// ParseGetInvestorDocumentsResponse parses an HTTP response from a GetInvestorDocumentsWithResponse call
func ParseGetInvestorDocumentsResponse(rsp *http.Response) (*GetInvestorDocumentsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetInvestorDocumentsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest DocumentsResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case rsp.StatusCode == 200:
		// Content-type (text/plain) unsupported

	}

	return response, nil
}

// ParseGetInvestorDocumentResponse parses an HTTP response from a GetInvestorDocumentWithResponse call
func ParseGetInvestorDocumentResponse(rsp *http.Response) (*GetInvestorDocumentResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetInvestorDocumentResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest SystemIOStream
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case rsp.StatusCode == 200:
		// Content-type (text/plain) unsupported

	}

	return response, nil
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/+yWTW/TTBDHv8qjeTiaOG0RB58roV4oapA4oB6m9iTd4n3p7LgQWf7uaNeNnTQOjUNV",
	"JMhtk8z8Zzfzm5cacqudNWTEQ1aDz29JYzye27zSZCScsSiUKGuw/MTWEYsiD9kcS08JuLWvaige3c5R",
	"KHyeW9YokEGBQm9FaYIEZOkIMvDCyiygSTqvi2LDp6pUMWQ+VyV9RB0DmKos8aYkyIQrGjBW5oG87Fbf",
	"U8Dyge6lzTH8eXtdVpNggYLbxk3TmdubO8olmK+S5K/IO2s8jcyWEtKbhzdMc8jg/7QHI32kIu2QaHY+",
	"G5lxCYN3nS29kJ5cXE5mwoR65FVzNFeEMQVMWFyacvkk9o21JaEJsXI0M6Jvext/VppsJXvbf2HV4v28",
	"dUlmIbcb5Cgj799BsstZGaEFcXB21qsVO1vu2+ZBcO0p6x5np4Me38NDxrhsZ7aJJTK3sfrJ56xce2W4",
	"ooXywsv/OkyDoJJAzfCPD8S+9T2ZTMP9rCODTkEGZ5Pp5AQScCi3EYi0L2yfFp1IVsOC4lsCPbHyQuHC",
	"B5KL3mE9qENGTULsIftagwrh7yviZaju2GVgzlbHjpY8tsh9W1uTDAuKPUzuOiS5LfX41NPpNBaHNbLq",
	"1c6Vqu046Z1v0emD7FPefTOJ2aYf8nJCrkT1m0phXGxgNqvynLxvg+AiZBH6VPeZvg4GG9TUqmjGovMc",
	"OYHPPs9xRDDdV4qpWNX5QM6HJ90ueo44/gM4pnW/FTVpvdp5mnGYvialm6JrO93Li3cb4K+kX5PWre3m",
	"QFh364xkdUDoRVC1PGraWh47a5XJy6qg4Y67ldxu0zr2yr+qV1rejd/Ysb2u9eeG9pHrI9djuD58/vcx",
	"jtP/OP0PmP5PIG2anwEAAP//5TuutigUAAA=",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}