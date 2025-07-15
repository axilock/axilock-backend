package api

// func addAuthorization(
// 	t *testing.T,
// 	request *http.Request,
// 	tokenMaker auth.Maker,
// 	authorizationType string,
// 	username string,
// 	duration time.Duration,
// ) {
// 	token, _, err := tokenMaker.CreateToken(username, duration)
// 	require.NoError(t, err)

// 	authorizationHeader := fmt.Sprintf("%s %s", authorizationType, token)
// 	request.Header.Set(authheaderkey, authorizationHeader)
// }

// func TestAuthMiddleware(t *testing.T) {
// 	username := util.RandomUUID()

// 	testCases := []struct {
// 		name          string
// 		setupAuth     func(t *testing.T, request *http.Request, tokenMaker auth.Maker)
// 		checkResponse func(t *testing.T, recorder *http.Response)
// 	}{
// 		{
// 			name: "OK",
// 			setupAuth: func(t *testing.T, request *http.Request, tokenMaker auth.Maker) {
// 				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, username, time.Minute)
// 			},
// 			checkResponse: func(t *testing.T, recorder *http.Response) {
// 				require.Equal(t, http.StatusOK, recorder.StatusCode)
// 			},
// 		},
// 		{
// 			name: "NoAuthorization",
// 			setupAuth: func(_ *testing.T, _ *http.Request, _ auth.Maker) {
// 			},
// 			checkResponse: func(t *testing.T, recorder *http.Response) {
// 				require.Equal(t, http.StatusUnauthorized, recorder.StatusCode)
// 			},
// 		},
// 		{
// 			name: "UnsupportedAuthorization",
// 			setupAuth: func(t *testing.T, request *http.Request, tokenMaker auth.Maker) {
// 				addAuthorization(t, request, tokenMaker, "unsupported", username, time.Minute)
// 			},
// 			checkResponse: func(t *testing.T, recorder *http.Response) {
// 				require.Equal(t, http.StatusUnauthorized, recorder.StatusCode)
// 			},
// 		},
// 		{
// 			name: "InvalidAuthorizationFormat",
// 			setupAuth: func(t *testing.T, request *http.Request, tokenMaker auth.Maker) {
// 				addAuthorization(t, request, tokenMaker, "", username, time.Minute)
// 			},
// 			checkResponse: func(t *testing.T, recorder *http.Response) {
// 				require.Equal(t, http.StatusUnauthorized, recorder.StatusCode)
// 			},
// 		},
// 		{
// 			name: "ExpiredToken",
// 			setupAuth: func(t *testing.T, request *http.Request, tokenMaker auth.Maker) {
// 				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, username, -time.Minute)
// 			},
// 			checkResponse: func(t *testing.T, recorder *http.Response) {
// 				require.Equal(t, http.StatusUnauthorized, recorder.StatusCode)
// 			},
// 		},
// 	}

// 	for i := range testCases {
// 		tc := testCases[i]
// 		server := newTestServer(t, nil)
// 		authPath := "/auth"
// 		server.router.Get(
// 			authPath,
// 			func(ctx fiber.Ctx) error {
// 				return ctx.JSON("")
// 			},
// 			authMiddleWare(server.tokenMaker),
// 		)
// 		t.Run(tc.name, func(t *testing.T) {
// 			request, err := http.NewRequest(http.MethodGet, authPath, nil)
// 			require.NoError(t, err)

// 			tc.setupAuth(t, request, server.tokenMaker)

// 			resp, _ := server.router.Test(request)
// 			tc.checkResponse(t, resp)
// 		})
// 	}
// }
