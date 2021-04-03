package tests

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/UsagiBooru/accounts-server/gen"
	"github.com/UsagiBooru/accounts-server/impl"
	"github.com/UsagiBooru/accounts-server/utils"
)

func GetAccountsServer() *httptest.Server {
	AccountsApiService := impl.NewAccountsApiImplService()
	AccountsApiController := gen.NewAccountsApiController(AccountsApiService)
	router := utils.NewRouterWithInject(AccountsApiController)
	return httptest.NewServer(router)
}

func SetAdminUserHeader(req *http.Request) *http.Request {
	req.Header.Set("x-consumer-user-id", "1")
	req.Header.Set("x-consumer-user-permission", strconv.Itoa(utils.PermissionAdmin))
	return req
}

func SetModUserHeader(req *http.Request) *http.Request {
	req.Header.Set("x-consumer-user-id", "2")
	req.Header.Set("x-consumer-user-permission", strconv.Itoa(utils.PermissionModelator))
	return req
}

func SetNormalUserHeader(req *http.Request) *http.Request {
	req.Header.Set("x-consumer-user-id", "3")
	req.Header.Set("x-consumer-user-permission", strconv.Itoa(utils.PermissionUser))
	return req
}

func TestMain(m *testing.M) {
	// utils.Debug("Resetting database...")
	err := ReGenerateTestDatabase()
	if err != nil {
		utils.Error(err.Error())
	}
	// utils.Debug("Reset database success.")

	m.Run()
}

func TestCreateAccount(t *testing.T) {
	s := GetAccountsServer()
	defer s.Close()
	newAccount := gen.AccountStruct{
		Name:      "デバッグアカウント",
		DisplayID: "debug_account",
		Password:  "debug_account",
		Mail:      "mail@example.com",
		Invite: gen.AccountStructInvite{
			Code: "dev",
		},
	}
	user_json, err := json.Marshal(newAccount)
	if err != nil {
		log.Fatal("Convert struct to json failed.")
	}
	req := httptest.NewRequest(
		http.MethodPost,
		"/accounts",
		bytes.NewBuffer(user_json),
	)
	rec := httptest.NewRecorder()
	s.Config.Handler.ServeHTTP(rec, req)
	t.Log(rec.Body)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestGetAccount(t *testing.T) {
	s := GetAccountsServer()
	defer s.Close()
	req := httptest.NewRequest(http.MethodGet, "/accounts/1", nil)
	rec := httptest.NewRecorder()
	s.Config.Handler.ServeHTTP(rec, req)
	t.Log(rec.Body)
	assert.Equal(t, http.StatusOK, rec.Code)
}
