/*
 * UsagiBooru Accounts API
 *
 * アカウント関連API
 *
 * API version: 2.0
 * Contact: dsgamer777@gmail.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package main

import (
	"log"
	"net/http"

	openapi "github.com/UsagiBooru/accounts-server/gen"
	impl "github.com/UsagiBooru/accounts-server/impl"
	"github.com/UsagiBooru/accounts-server/utils"
)

func main() {
	utils.Info("Server started")

	AccountsApiService := impl.NewAccountsApiImplService()
	AccountsApiController := openapi.NewAccountsApiController(AccountsApiService)

	MutesApiService := openapi.NewMutesApiService()
	MutesApiController := openapi.NewMutesApiController(MutesApiService)

	MylistApiService := openapi.NewMylistApiService()
	MylistApiController := openapi.NewMylistApiController(MylistApiService)

	NotifyApiService := openapi.NewNotifyApiService()
	NotifyApiController := openapi.NewNotifyApiController(NotifyApiService)

	TimelineApiService := openapi.NewTimelineApiService()
	TimelineApiController := openapi.NewTimelineApiController(TimelineApiService)

	router := openapi.NewRouter(AccountsApiController, MutesApiController, MylistApiController, NotifyApiController, TimelineApiController)

	log.Fatal(http.ListenAndServe(":8000", router))
}
