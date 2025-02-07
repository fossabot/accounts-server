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
	"context"
	"net/http"
)



// AccountsApiRouter defines the required methods for binding the api requests to a responses for the AccountsApi
// The AccountsApiRouter implementation should parse necessary information from the http request, 
// pass the data to a AccountsApiServicer to perform the required actions, then write the service results to the http response.
type AccountsApiRouter interface { 
	CreateAccount(http.ResponseWriter, *http.Request)
	DeleteAccount(http.ResponseWriter, *http.Request)
	EditAccount(http.ResponseWriter, *http.Request)
	GetAccount(http.ResponseWriter, *http.Request)
	GetAccountMe(http.ResponseWriter, *http.Request)
	GetUploadHistory(http.ResponseWriter, *http.Request)
	LoginWithForm(http.ResponseWriter, *http.Request)
}
// MutesApiRouter defines the required methods for binding the api requests to a responses for the MutesApi
// The MutesApiRouter implementation should parse necessary information from the http request, 
// pass the data to a MutesApiServicer to perform the required actions, then write the service results to the http response.
type MutesApiRouter interface { 
	AddMute(http.ResponseWriter, *http.Request)
	DeleteMute(http.ResponseWriter, *http.Request)
	GetMute(http.ResponseWriter, *http.Request)
	GetMutes(http.ResponseWriter, *http.Request)
}
// MylistApiRouter defines the required methods for binding the api requests to a responses for the MylistApi
// The MylistApiRouter implementation should parse necessary information from the http request, 
// pass the data to a MylistApiServicer to perform the required actions, then write the service results to the http response.
type MylistApiRouter interface { 
	CreateMylist(http.ResponseWriter, *http.Request)
	GetUserMylists(http.ResponseWriter, *http.Request)
}
// NotifyApiRouter defines the required methods for binding the api requests to a responses for the NotifyApi
// The NotifyApiRouter implementation should parse necessary information from the http request, 
// pass the data to a NotifyApiServicer to perform the required actions, then write the service results to the http response.
type NotifyApiRouter interface { 
	AddLineNotifyClient(http.ResponseWriter, *http.Request)
	AddWebNotifyClient(http.ResponseWriter, *http.Request)
	DeleteNotifyClient(http.ResponseWriter, *http.Request)
	DeleteNotifyCondition(http.ResponseWriter, *http.Request)
	EditNotifyClient(http.ResponseWriter, *http.Request)
	EditNotifyCondition(http.ResponseWriter, *http.Request)
	GetNotifyClient(http.ResponseWriter, *http.Request)
	GetNotifyClients(http.ResponseWriter, *http.Request)
	GetNotifyCondition(http.ResponseWriter, *http.Request)
	GetNotifyConditions(http.ResponseWriter, *http.Request)
	RegisterNotifyCondition(http.ResponseWriter, *http.Request)
}
// TimelineApiRouter defines the required methods for binding the api requests to a responses for the TimelineApi
// The TimelineApiRouter implementation should parse necessary information from the http request, 
// pass the data to a TimelineApiServicer to perform the required actions, then write the service results to the http response.
type TimelineApiRouter interface { 
	FollowArtist(http.ResponseWriter, *http.Request)
	GetFollowingArtists(http.ResponseWriter, *http.Request)
	UnfollowArtist(http.ResponseWriter, *http.Request)
}


// AccountsApiServicer defines the api actions for the AccountsApi service
// This interface intended to stay up to date with the openapi yaml used to generate it, 
// while the service implementation can ignored with the .openapi-generator-ignore file 
// and updated with the logic required for the API.
type AccountsApiServicer interface { 
	CreateAccount(context.Context, AccountStruct) (ImplResponse, error)
	DeleteAccount(context.Context, int32, string) (ImplResponse, error)
	EditAccount(context.Context, int32, AccountStruct) (ImplResponse, error)
	GetAccount(context.Context, int32) (ImplResponse, error)
	GetAccountMe(context.Context) (ImplResponse, error)
	GetUploadHistory(context.Context, int32, int32, string, string, int32) (ImplResponse, error)
	LoginWithForm(context.Context, PostLoginWithFormRequest) (ImplResponse, error)
}


// MutesApiServicer defines the api actions for the MutesApi service
// This interface intended to stay up to date with the openapi yaml used to generate it, 
// while the service implementation can ignored with the .openapi-generator-ignore file 
// and updated with the logic required for the API.
type MutesApiServicer interface { 
	AddMute(context.Context, int32, MuteStruct) (ImplResponse, error)
	DeleteMute(context.Context, int32, int32) (ImplResponse, error)
	GetMute(context.Context, int32, int32) (ImplResponse, error)
	GetMutes(context.Context, int32) (ImplResponse, error)
}


// MylistApiServicer defines the api actions for the MylistApi service
// This interface intended to stay up to date with the openapi yaml used to generate it, 
// while the service implementation can ignored with the .openapi-generator-ignore file 
// and updated with the logic required for the API.
type MylistApiServicer interface { 
	CreateMylist(context.Context, int32, MylistStruct) (ImplResponse, error)
	GetUserMylists(context.Context, int32) (ImplResponse, error)
}


// NotifyApiServicer defines the api actions for the NotifyApi service
// This interface intended to stay up to date with the openapi yaml used to generate it, 
// while the service implementation can ignored with the .openapi-generator-ignore file 
// and updated with the logic required for the API.
type NotifyApiServicer interface { 
	AddLineNotifyClient(context.Context, int32, PostRegisterLineNotifyRequest) (ImplResponse, error)
	AddWebNotifyClient(context.Context, int32, PostRegisterWebPushRequest) (ImplResponse, error)
	DeleteNotifyClient(context.Context, int32, int32) (ImplResponse, error)
	DeleteNotifyCondition(context.Context, int32, int32) (ImplResponse, error)
	EditNotifyClient(context.Context, int32, int32, NotifyClientStruct) (ImplResponse, error)
	EditNotifyCondition(context.Context, int32, int32, NotifyConditionStruct) (ImplResponse, error)
	GetNotifyClient(context.Context, int32, int32) (ImplResponse, error)
	GetNotifyClients(context.Context, int32) (ImplResponse, error)
	GetNotifyCondition(context.Context, int32, int32) (ImplResponse, error)
	GetNotifyConditions(context.Context, int32, string) (ImplResponse, error)
	RegisterNotifyCondition(context.Context, int32, NotifyConditionStruct) (ImplResponse, error)
}


// TimelineApiServicer defines the api actions for the TimelineApi service
// This interface intended to stay up to date with the openapi yaml used to generate it, 
// while the service implementation can ignored with the .openapi-generator-ignore file 
// and updated with the logic required for the API.
type TimelineApiServicer interface { 
	FollowArtist(context.Context, int32, LightArtistStruct) (ImplResponse, error)
	GetFollowingArtists(context.Context, int32, string, string, int32) (ImplResponse, error)
	UnfollowArtist(context.Context, int32, LightArtistStruct) (ImplResponse, error)
}
