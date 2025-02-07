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

	"github.com/UsagiBooru/accounts-server/gen"
	"github.com/UsagiBooru/accounts-server/impl"
	"github.com/UsagiBooru/accounts-server/utils/server"
)

func main() {
	server.Info("Server started")
	conf := server.GetConfig()
	md := server.NewMongoDBClient(conf.MongoHost, conf.MongoUser, conf.MongoPass)

	AccountsApiService := impl.NewAccountsApiImplService(md, conf.JwtSecret)
	AccountsApiController := gen.NewAccountsApiController(AccountsApiService)

	MutesApiService := gen.NewMutesApiService()
	MutesApiController := gen.NewMutesApiController(MutesApiService)

	MylistApiService := gen.NewMylistApiService()
	MylistApiController := gen.NewMylistApiController(MylistApiService)

	NotifyApiService := gen.NewNotifyApiService()
	NotifyApiController := gen.NewNotifyApiController(NotifyApiService)

	TimelineApiService := gen.NewTimelineApiService()
	TimelineApiController := gen.NewTimelineApiController(TimelineApiService)

	router := server.NewRouterWithInject(AccountsApiController, MutesApiController, MylistApiController, NotifyApiController, TimelineApiController)

	log.Fatal(http.ListenAndServe(":8000", router))
}
