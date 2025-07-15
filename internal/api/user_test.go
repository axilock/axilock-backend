package api

// import (
// 	"bytes"
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// 	"testing"

// 	"github.com/axilock/axilock-backend/internal/axierr"
// 	db "github.com/axilock/axilock-backend/internal/db/sqlc"
// 	"github.com/axilock/axilock-backend/internal/mocks"
// 	"github.com/axilock/axilock-backend/pkg/util"
// 	"github.com/gofiber/fiber/v3"
// 	"github.com/jackc/pgx/v5"
// 	"github.com/stretchr/testify/mock"
// 	"github.com/stretchr/testify/require"
// )

// func TestCreateUserAPI(t *testing.T) {
// 	user, _, org := randomUser(t)

// 	testCases := []struct {
// 		name          string
// 		body          fiber.Map
// 		buildStubs    func(store *mocks.Store, svc *mocks.MockServiceStore)
// 		checkResponse func(recoder *http.Response)
// 	}{
// 		{
// 			name: "OK",
// 			body: fiber.Map{
// 				"email":            user.Email,
// 				"password":         "admin@12345",
// 				"confirm_password": "admin@12345",
// 			},
// 			buildStubs: func(store *mocks.Store, svc *mocks.MockServiceStore) {
// 				store.EXPECT().
// 					GetOrgByEntity(mock.Anything, mock.Anything).
// 					Times(1).
// 					Return(org, nil)
// 				store.EXPECT().
// 					CreateUser(mock.Anything, mock.Anything).
// 					Times(1).
// 					Return(user, nil)
// 			},
// 			checkResponse: func(recorder *http.Response) {
// 				require.Equal(t, http.StatusOK, recorder.StatusCode)
// 			},
// 		},
// 		{
// 			name: "InternalError",
// 			body: fiber.Map{
// 				"email":            user.Email,
// 				"password":         "admin@12345",
// 				"confirm_password": "admin@12345",
// 			},
// 			buildStubs: func(store *mocks.Store, svc *mocks.MockServiceStore) {
// 				store.EXPECT().
// 					GetOrgByEntity(mock.Anything, mock.Anything).
// 					Times(1).
// 					Return(org, nil)
// 				store.EXPECT().
// 					CreateUser(mock.Anything, mock.Anything).
// 					Times(1).
// 					Return(db.User{}, pgx.ErrTxClosed)
// 			},
// 			checkResponse: func(recorder *http.Response) {
// 				require.Equal(t, http.StatusInternalServerError, recorder.StatusCode)
// 			},
// 		},
// 		{
// 			name: "DuplicateEmail",
// 			body: fiber.Map{
// 				"email":            user.Email,
// 				"password":         "admin@12345",
// 				"confirm_password": "admin@12345",
// 			},
// 			buildStubs: func(store *mocks.Store, svc *mocks.MockServiceStore) {
// 				store.EXPECT().
// 					GetOrgByEntity(mock.Anything, mock.Anything).
// 					Times(1).
// 					Return(org, nil)
// 				store.EXPECT().
// 					CreateUser(mock.Anything, mock.Anything).
// 					Times(1).
// 					Return(user, axierr.ErrUniqueViolation)
// 			},
// 			checkResponse: func(recorder *http.Response) {
// 				require.Equal(t, http.StatusForbidden, recorder.StatusCode)
// 			},
// 		},
// 		{
// 			name: "InvalidEmail",
// 			body: fiber.Map{
// 				"email":            "abcdfasfs",
// 				"password":         "admin@12345",
// 				"confirm_password": "admin@12345",
// 			},
// 			buildStubs: func(store *mocks.Store, svc *mocks.MockServiceStore) {
// 			},
// 			checkResponse: func(recorder *http.Response) {
// 				require.Equal(t, http.StatusBadRequest, recorder.StatusCode)
// 			},
// 		},
// 		{
// 			name: "TooShortPassword",
// 			body: fiber.Map{
// 				"email":            user.Email,
// 				"password":         "admin",
// 				"confirm_password": "admin",
// 			},
// 			buildStubs: func(store *mocks.Store, svc *mocks.MockServiceStore) {
// 			},
// 			checkResponse: func(recorder *http.Response) {
// 				require.Equal(t, http.StatusBadRequest, recorder.StatusCode)
// 			},
// 		},
// 	}

// 	for i := range testCases {
// 		tc := testCases[i]

// 		t.Run(tc.name, func(t *testing.T) {
// 			store := mocks.NewStore(t)
// 			services := mocks.NewMockServiceStore(t)
// 			tc.buildStubs(store, services)

// 			server := newTestServer(t, store)

// 			// Marshal body data to JSON
// 			data, err := json.Marshal(tc.body)
// 			require.NoError(t, err)

// 			url := "/v1//user/create"
// 			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
// 			require.NoError(t, err)

// 			resp, _ := server.router.Test(request)
// 			tc.checkResponse(resp)
// 		})
// 	}
// }

// func TestLoginUserAPI(t *testing.T) {
// 	user, password, _ := randomUser(t)

// 	testCases := []struct {
// 		name          string
// 		body          fiber.Map
// 		buildStubs    func(store *mocks.Store)
// 		checkResponse func(recoder *http.Response)
// 	}{
// 		{
// 			name: "OK",
// 			body: fiber.Map{
// 				"email":    user.Email,
// 				"password": password,
// 			},
// 			buildStubs: func(store *mocks.Store) {
// 				store.EXPECT().
// 					GetUserByEmail(mock.Anything, mock.Anything).
// 					Times(1).
// 					Return(user, nil)
// 			},
// 			checkResponse: func(recorder *http.Response) {
// 				require.Equal(t, http.StatusOK, recorder.StatusCode)
// 			},
// 		},
// 		{
// 			name: "UserNotFound",
// 			body: fiber.Map{
// 				"email":    "abcd@abcd.com",
// 				"password": password,
// 			},
// 			buildStubs: func(store *mocks.Store) {
// 				store.EXPECT().
// 					GetUserByEmail(mock.Anything, mock.Anything).
// 					Times(1).
// 					Return(db.User{}, axierr.ErrRecordNotFound)
// 			},
// 			checkResponse: func(recorder *http.Response) {
// 				require.Equal(t, http.StatusNotFound, recorder.StatusCode)
// 			},
// 		},
// 		{
// 			name: "IncorrectPassword",
// 			body: fiber.Map{
// 				"email":    user.Email,
// 				"password": "incorrect",
// 			},
// 			buildStubs: func(store *mocks.Store) {
// 				store.EXPECT().
// 					GetUserByEmail(mock.Anything, mock.Anything).
// 					Times(1).
// 					Return(user, nil)
// 			},
// 			checkResponse: func(recorder *http.Response) {
// 				require.Equal(t, http.StatusUnauthorized, recorder.StatusCode)
// 			},
// 		},
// 		{
// 			name: "InternalError",
// 			body: fiber.Map{
// 				"email":    user.Email,
// 				"password": password,
// 			},
// 			buildStubs: func(store *mocks.Store) {
// 				store.EXPECT().
// 					GetUserByEmail(mock.Anything, mock.Anything).
// 					Times(1).
// 					Return(user, fmt.Errorf("Internal error"))
// 			},
// 			checkResponse: func(recorder *http.Response) {
// 				require.Equal(t, http.StatusInternalServerError, recorder.StatusCode)
// 			},
// 		},
// 		{
// 			name: "InvalidUsername",
// 			body: fiber.Map{
// 				"email":    "invalid-user",
// 				"password": password,
// 			},
// 			buildStubs: func(store *mocks.Store) {
// 			},
// 			checkResponse: func(recorder *http.Response) {
// 				require.Equal(t, http.StatusBadRequest, recorder.StatusCode)
// 			},
// 		},
// 	}

// 	for i := range testCases {
// 		tc := testCases[i]

// 		t.Run(tc.name, func(t *testing.T) {
// 			store := mocks.NewStore(t)
// 			tc.buildStubs(store)

// 			server := newTestServer(t, store)

// 			// Marshal body data to JSON
// 			data, err := json.Marshal(tc.body)
// 			require.NoError(t, err)

// 			url := "/v1/user/login"
// 			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
// 			require.NoError(t, err)

// 			resp, _ := server.router.Test(request)
// 			tc.checkResponse(resp)
// 		})
// 	}
// }

// func randomUser(t *testing.T) (user db.User, password string, org db.Organisation) {
// 	password = util.RandString(6)
// 	hashedPassword, err := util.NewPassword(password)
// 	require.NoError(t, err)

// 	email := util.RandomEmail()
// 	domain, err := util.GetDomain(email)
// 	require.NoError(t, err)
// 	orgId := util.RandomInt(4, 8)

// 	user = db.User{
// 		HashPassword: hashedPassword,
// 		Email:        email,
// 		EntityID:     util.RandomUUID(),
// 		ID:           util.RandomInt(4, 8),
// 		Org:          orgId,
// 	}
// 	org = db.Organisation{
// 		ID:       orgId,
// 		Name:     "sekrit",
// 		EntityID: util.RandomUUID(),
// 		Domain:   domain,
// 	}
// 	return
// }
