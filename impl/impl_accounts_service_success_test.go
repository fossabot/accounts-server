package impl_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/UsagiBooru/accounts-server/gen"
	"github.com/UsagiBooru/accounts-server/impl"
	"github.com/UsagiBooru/accounts-server/utils/server"
	"github.com/UsagiBooru/accounts-server/utils/tests"
)

func GetAccountsServer() (*httptest.Server, func(), bool) {
	db, shutdown, isParallel := tests.GetDatabaseConnection()
	AccountsApiService := impl.NewAccountsApiImplService(db, tests.JWT_SECRET)
	AccountsApiController := gen.NewAccountsApiController(AccountsApiService)
	router := server.NewRouterWithInject(AccountsApiController)
	return httptest.NewServer(router), shutdown, isParallel
}

func TestGetAccountSuccessOnValid(t *testing.T) {
	s, shutdown, isParallel := GetAccountsServer()
	if isParallel {
		t.Parallel()
	}
	defer s.Close()
	defer shutdown()
	req := httptest.NewRequest(http.MethodGet, "/accounts/1", nil)
	rec := httptest.NewRecorder()
	s.Config.Handler.ServeHTTP(rec, req)
	t.Log(rec.Body)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestGetAccountSuccessOnDeletedIdFromMod(t *testing.T) {
	s, shutdown, isParallel := GetAccountsServer()
	if isParallel {
		t.Parallel()
	}
	defer s.Close()
	defer shutdown()
	req := httptest.NewRequest(http.MethodGet, "/accounts/4", nil)
	req = tests.SetModUserHeader(req)
	rec := httptest.NewRecorder()
	s.Config.Handler.ServeHTTP(rec, req)
	t.Log(rec.Body)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestCreateAccountSuccessOnValid(t *testing.T) {
	s, shutdown, isParallel := GetAccountsServer()
	if isParallel {
		t.Parallel()
	}
	defer s.Close()
	defer shutdown()
	newAccount := gen.AccountStruct{
		Name:      "デバッグアカウント",
		DisplayID: "debugaccount",
		Password:  "debugaccount",
		Mail:      "mail@example.com",
		Invite: gen.AccountStructInvite{
			Code: "devcode1",
		},
	}
	user_json, _ := json.Marshal(newAccount)
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

func TestEditAccountSuccessOnChangeName(t *testing.T) {
	s, shutdown, isParallel := GetAccountsServer()
	if isParallel {
		t.Parallel()
	}
	defer s.Close()
	defer shutdown()
	editAccount := gen.AccountStruct{
		Name: "デバッグアカウント2",
	}
	req_json, _ := json.Marshal(editAccount)
	req := httptest.NewRequest(
		http.MethodPatch,
		"/accounts/1",
		bytes.NewBuffer(req_json),
	)
	req = tests.SetAdminUserHeader(req)
	rec := httptest.NewRecorder()
	s.Config.Handler.ServeHTTP(rec, req)
	t.Log(rec.Body)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestGetAccountMeSuccessFromAdmin(t *testing.T) {
	s, shutdown, isParallel := GetAccountsServer()
	if isParallel {
		t.Parallel()
	}
	defer s.Close()
	defer shutdown()
	req := httptest.NewRequest(
		http.MethodGet,
		"/accounts/me",
		nil,
	)
	req = tests.SetAdminUserHeader(req)
	rec := httptest.NewRecorder()
	s.Config.Handler.ServeHTTP(rec, req)
	t.Log(rec.Body)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestGetAccountMeSuccessFromMod(t *testing.T) {
	s, shutdown, isParallel := GetAccountsServer()
	if isParallel {
		t.Parallel()
	}
	defer s.Close()
	defer shutdown()
	req := httptest.NewRequest(
		http.MethodGet,
		"/accounts/me",
		nil,
	)
	req = tests.SetModUserHeader(req)
	rec := httptest.NewRecorder()
	s.Config.Handler.ServeHTTP(rec, req)
	t.Log(rec.Body)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestGetAccountMeSuccessFromNormal(t *testing.T) {
	s, shutdown, isParallel := GetAccountsServer()
	if isParallel {
		t.Parallel()
	}
	defer s.Close()
	defer shutdown()
	req := httptest.NewRequest(
		http.MethodGet,
		"/accounts/me",
		nil,
	)
	req = tests.SetNormalUserHeader(req)
	rec := httptest.NewRecorder()
	s.Config.Handler.ServeHTTP(rec, req)
	t.Log(rec.Body)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestLoginWithFormSuccessOnValid(t *testing.T) {
	s, shutdown, isParallel := GetAccountsServer()
	if isParallel {
		t.Parallel()
	}
	defer s.Close()
	defer shutdown()
	editAccount := gen.PostLoginWithFormRequest{
		Id:       "domao",
		Password: tests.PASSWORD,
	}
	req_json, _ := json.Marshal(editAccount)
	req := httptest.NewRequest(
		http.MethodPost,
		"/accounts/login/form",
		bytes.NewBuffer(req_json),
	)
	rec := httptest.NewRecorder()
	s.Config.Handler.ServeHTTP(rec, req)
	t.Log(rec.Body)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestDeleteAccountSuccessFromAdmin(t *testing.T) {
	s, shutdown, isParallel := GetAccountsServer()
	if isParallel {
		t.Parallel()
	}
	defer s.Close()
	defer shutdown()
	req := httptest.NewRequest(
		http.MethodDelete,
		"/accounts/1",
		nil,
	)
	req = tests.SetAdminUserHeader(req)
	rec := httptest.NewRecorder()
	s.Config.Handler.ServeHTTP(rec, req)
	t.Log(rec.Body)
	assert.Equal(t, http.StatusNoContent, rec.Code)
}

func TestDeleteAccountSuccessFromMod(t *testing.T) {
	s, shutdown, isParallel := GetAccountsServer()
	if isParallel {
		t.Parallel()
	}
	defer s.Close()
	defer shutdown()
	req := httptest.NewRequest(
		http.MethodDelete,
		"/accounts/1",
		nil,
	)
	req = tests.SetModUserHeader(req)
	rec := httptest.NewRecorder()
	s.Config.Handler.ServeHTTP(rec, req)
	t.Log(rec.Body)
	assert.Equal(t, http.StatusNoContent, rec.Code)
}

func TestDeleteAccountSuccessFromSelf(t *testing.T) {
	s, shutdown, isParallel := GetAccountsServer()
	if isParallel {
		t.Parallel()
	}
	defer s.Close()
	defer shutdown()
	req := httptest.NewRequest(
		http.MethodDelete,
		"/accounts/3",
		nil,
	)
	req = tests.SetNormalUserHeader(req)
	req.Header.Set("password", tests.PASSWORD)
	rec := httptest.NewRecorder()
	s.Config.Handler.ServeHTTP(rec, req)
	t.Log(rec.Body)
	assert.Equal(t, http.StatusNoContent, rec.Code)
}
