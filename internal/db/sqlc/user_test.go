package db

// func createRandomUser(t *testing.T) User {
// 	orgarg := CreateOrganiationParams{
// 		Name: util.RandString(8),
// 		Domain: pgtype.Text{
// 			String: util.GetDomain(util.RandomEmail()),
// 			Valid:  true,
// 		},
// 	}
// 	org, err := testQueries.CreateOrganiation(context.Background(), orgarg)
// 	require.NoError(t, err)
// 	require.Equal(t, org.Domain, orgarg.Domain)
// 	require.Equal(t, org.Name, orgarg.Name)
// 	require.NotZero(t, org.ID)

// 	hashPassword, err := util.NewPassword(util.RandString(6))
// 	require.NoError(t, err)

// 	arg := CreateUserParams{
// 		Email:        util.RandomEmail(),
// 		Org:          org.ID,
// 		HashPassword: hashPassword,
// 	}
// 	user, err := testQueries.CreateUser(context.Background(), arg)

// 	require.NoError(t, err)
// 	require.NotEmpty(t, user)
// 	require.Equal(t, arg.Email, user.Email)
// 	require.NotZero(t, user.ID)

// 	return user
// }

// func TestCreateUser(t *testing.T) {
// 	createRandomUser(t)
// }

// func TestGetAccount(t *testing.T) {
// 	user1 := createRandomUser(t)
// 	userGet, err := testQueries.GetUserByEntityId(context.Background(), user1.Uuid)

// 	require.NoError(t, err)
// 	require.NotEmpty(t, userGet)

// 	require.Equal(t, user1.Org, userGet.Org)
// 	require.Equal(t, user1.Email, userGet.Email)
// 	require.WithinDuration(t, user1.CreatedAt.Time, userGet.CreatedAt.Time, time.Second)
// }
