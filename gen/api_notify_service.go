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
	"errors"
)

// NotifyApiService is a service that implents the logic for the NotifyApiServicer
// This service should implement the business logic for every endpoint for the NotifyApi API. 
// Include any external packages or services that will be required by this service.
type NotifyApiService struct {
}

// NewNotifyApiService creates a default api service
func NewNotifyApiService() NotifyApiServicer {
	return &NotifyApiService{}
}

// AddLineNotifyClient - Create line notify client
func (s *NotifyApiService) AddLineNotifyClient(ctx context.Context, accountID int32, postRegisterLineNotifyRequest PostRegisterLineNotifyRequest) (ImplResponse, error) {
	// TODO - update AddLineNotifyClient with the required logic for this service method.
	// Add api_notify_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, NotifyClientStruct{}) or use other options such as http.Ok ...
	//return Response(200, NotifyClientStruct{}), nil

	//TODO: Uncomment the next line to return response Response(400, GeneralMessageResponse{}) or use other options such as http.Ok ...
	//return Response(400, GeneralMessageResponse{}), nil

	//TODO: Uncomment the next line to return response Response(403, GeneralMessageResponse{}) or use other options such as http.Ok ...
	//return Response(403, GeneralMessageResponse{}), nil

	//TODO: Uncomment the next line to return response Response(404, GeneralMessageResponse{}) or use other options such as http.Ok ...
	//return Response(404, GeneralMessageResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("AddLineNotifyClient method not implemented")
}

// AddWebNotifyClient - Create webpush notify client
func (s *NotifyApiService) AddWebNotifyClient(ctx context.Context, accountID int32, postRegisterWebPushRequest PostRegisterWebPushRequest) (ImplResponse, error) {
	// TODO - update AddWebNotifyClient with the required logic for this service method.
	// Add api_notify_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, NotifyClientStruct{}) or use other options such as http.Ok ...
	//return Response(200, NotifyClientStruct{}), nil

	//TODO: Uncomment the next line to return response Response(403, GeneralMessageResponse{}) or use other options such as http.Ok ...
	//return Response(403, GeneralMessageResponse{}), nil

	//TODO: Uncomment the next line to return response Response(404, GeneralMessageResponse{}) or use other options such as http.Ok ...
	//return Response(404, GeneralMessageResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("AddWebNotifyClient method not implemented")
}

// DeleteNotifyClient - Delete notify client
func (s *NotifyApiService) DeleteNotifyClient(ctx context.Context, accountID int32, notifyClientID int32) (ImplResponse, error) {
	// TODO - update DeleteNotifyClient with the required logic for this service method.
	// Add api_notify_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(204, GeneralMessageResponse{}) or use other options such as http.Ok ...
	//return Response(204, GeneralMessageResponse{}), nil

	//TODO: Uncomment the next line to return response Response(403, GeneralMessageResponse{}) or use other options such as http.Ok ...
	//return Response(403, GeneralMessageResponse{}), nil

	//TODO: Uncomment the next line to return response Response(404, GeneralMessageResponse{}) or use other options such as http.Ok ...
	//return Response(404, GeneralMessageResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("DeleteNotifyClient method not implemented")
}

// DeleteNotifyCondition - Delete notify condition
func (s *NotifyApiService) DeleteNotifyCondition(ctx context.Context, conditionID int32, accountID int32) (ImplResponse, error) {
	// TODO - update DeleteNotifyCondition with the required logic for this service method.
	// Add api_notify_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(204, GeneralMessageResponse{}) or use other options such as http.Ok ...
	//return Response(204, GeneralMessageResponse{}), nil

	//TODO: Uncomment the next line to return response Response(403, GeneralMessageResponse{}) or use other options such as http.Ok ...
	//return Response(403, GeneralMessageResponse{}), nil

	//TODO: Uncomment the next line to return response Response(404, GeneralMessageResponse{}) or use other options such as http.Ok ...
	//return Response(404, GeneralMessageResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("DeleteNotifyCondition method not implemented")
}

// EditNotifyClient - Edit notify client
func (s *NotifyApiService) EditNotifyClient(ctx context.Context, accountID int32, notifyClientID int32, notifyClientStruct NotifyClientStruct) (ImplResponse, error) {
	// TODO - update EditNotifyClient with the required logic for this service method.
	// Add api_notify_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, NotifyClientStruct{}) or use other options such as http.Ok ...
	//return Response(200, NotifyClientStruct{}), nil

	//TODO: Uncomment the next line to return response Response(403, GeneralMessageResponse{}) or use other options such as http.Ok ...
	//return Response(403, GeneralMessageResponse{}), nil

	//TODO: Uncomment the next line to return response Response(404, GeneralMessageResponse{}) or use other options such as http.Ok ...
	//return Response(404, GeneralMessageResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("EditNotifyClient method not implemented")
}

// EditNotifyCondition - Edit notify condition
func (s *NotifyApiService) EditNotifyCondition(ctx context.Context, conditionID int32, accountID int32, notifyConditionStruct NotifyConditionStruct) (ImplResponse, error) {
	// TODO - update EditNotifyCondition with the required logic for this service method.
	// Add api_notify_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, NotifyConditionStruct{}) or use other options such as http.Ok ...
	//return Response(200, NotifyConditionStruct{}), nil

	//TODO: Uncomment the next line to return response Response(400, GeneralMessageResponse{}) or use other options such as http.Ok ...
	//return Response(400, GeneralMessageResponse{}), nil

	//TODO: Uncomment the next line to return response Response(403, GeneralMessageResponse{}) or use other options such as http.Ok ...
	//return Response(403, GeneralMessageResponse{}), nil

	//TODO: Uncomment the next line to return response Response(404, GeneralMessageResponse{}) or use other options such as http.Ok ...
	//return Response(404, GeneralMessageResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("EditNotifyCondition method not implemented")
}

// GetNotifyClient - Get notify client
func (s *NotifyApiService) GetNotifyClient(ctx context.Context, accountID int32, notifyClientID int32) (ImplResponse, error) {
	// TODO - update GetNotifyClient with the required logic for this service method.
	// Add api_notify_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, NotifyClientStruct{}) or use other options such as http.Ok ...
	//return Response(200, NotifyClientStruct{}), nil

	//TODO: Uncomment the next line to return response Response(403, GeneralMessageResponse{}) or use other options such as http.Ok ...
	//return Response(403, GeneralMessageResponse{}), nil

	//TODO: Uncomment the next line to return response Response(404, GeneralMessageResponse{}) or use other options such as http.Ok ...
	//return Response(404, GeneralMessageResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetNotifyClient method not implemented")
}

// GetNotifyClients - Get notify clients
func (s *NotifyApiService) GetNotifyClients(ctx context.Context, accountID int32) (ImplResponse, error) {
	// TODO - update GetNotifyClients with the required logic for this service method.
	// Add api_notify_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, GetNotifyClientsResponse{}) or use other options such as http.Ok ...
	//return Response(200, GetNotifyClientsResponse{}), nil

	//TODO: Uncomment the next line to return response Response(403, GeneralMessageResponse{}) or use other options such as http.Ok ...
	//return Response(403, GeneralMessageResponse{}), nil

	//TODO: Uncomment the next line to return response Response(404, GeneralMessageResponse{}) or use other options such as http.Ok ...
	//return Response(404, GeneralMessageResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetNotifyClients method not implemented")
}

// GetNotifyCondition - Get notify condition
func (s *NotifyApiService) GetNotifyCondition(ctx context.Context, conditionID int32, accountID int32) (ImplResponse, error) {
	// TODO - update GetNotifyCondition with the required logic for this service method.
	// Add api_notify_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, NotifyConditionStruct{}) or use other options such as http.Ok ...
	//return Response(200, NotifyConditionStruct{}), nil

	//TODO: Uncomment the next line to return response Response(403, GeneralMessageResponse{}) or use other options such as http.Ok ...
	//return Response(403, GeneralMessageResponse{}), nil

	//TODO: Uncomment the next line to return response Response(404, GeneralMessageResponse{}) or use other options such as http.Ok ...
	//return Response(404, GeneralMessageResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetNotifyCondition method not implemented")
}

// GetNotifyConditions - Get notify conditions
func (s *NotifyApiService) GetNotifyConditions(ctx context.Context, accountID int32, type_ string) (ImplResponse, error) {
	// TODO - update GetNotifyConditions with the required logic for this service method.
	// Add api_notify_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, GetNotifyConditionsResponse{}) or use other options such as http.Ok ...
	//return Response(200, GetNotifyConditionsResponse{}), nil

	//TODO: Uncomment the next line to return response Response(403, GeneralMessageResponse{}) or use other options such as http.Ok ...
	//return Response(403, GeneralMessageResponse{}), nil

	//TODO: Uncomment the next line to return response Response(404, GeneralMessageResponse{}) or use other options such as http.Ok ...
	//return Response(404, GeneralMessageResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetNotifyConditions method not implemented")
}

// RegisterNotifyCondition - Register notify condition
func (s *NotifyApiService) RegisterNotifyCondition(ctx context.Context, accountID int32, notifyConditionStruct NotifyConditionStruct) (ImplResponse, error) {
	// TODO - update RegisterNotifyCondition with the required logic for this service method.
	// Add api_notify_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, NotifyConditionStruct{}) or use other options such as http.Ok ...
	//return Response(200, NotifyConditionStruct{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("RegisterNotifyCondition method not implemented")
}

