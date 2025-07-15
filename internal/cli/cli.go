package cli

import (
	"encoding/json"
	"fmt"

	db "github.com/axilock/axilock-backend/internal/db/sqlc"
	"github.com/axilock/axilock-backend/pkg/util"
)

var Cli struct {
	//TODO: version ?
	Serve              ServeCmd              `cmd:"" help:"Start the server" default:"true"`
	UpdateUserPassword UpdateUserPasswordCmd `cmd:"" help:"Update user password"`
}

// Serve is handled by the main package, this is just a placeholder
type ServeCmd struct{}

func (c *ServeCmd) Run() error { return nil }

type UpdateUserPasswordCmd struct {
	Email    string `arg:"" name:"email" help:"Email of the user to reset password for"`
	Password string `arg:"" name:"password" help:"New password for the user"`
}

func (c *UpdateUserPasswordCmd) Run(store db.Store) error {
	ctx, cancel := timeoutCtx()
	defer cancel()

	userObj, err := store.GetUserByEmail(ctx, c.Email)
	if err != nil {
		return fmt.Errorf("error fetching user: %w", err)
	}

	user, _ := json.MarshalIndent(userObj, "", "  ")
	fmt.Println("Updating password for user:", string(user))

	hash, err := util.NewPassword(c.Password)
	if err != nil {
		return err
	}

	err = store.UpdateUserPassword(ctx, db.UpdateUserPasswordParams{
		HashPassword: hash,
		Email:        c.Email,
	})
	if err != nil {
		return err
	}

	fmt.Println("Password updated successfully for user:", c.Email)
	return nil
}
