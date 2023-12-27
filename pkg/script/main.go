package script

import (
	"context"
	"os"

	"github.com/jykuo-love-shiritori/twp/db"
	"golang.org/x/crypto/bcrypt"
)

func CheckAdminAccount(pg *db.DB) error {
	admin_name := os.Getenv("TWP_ADMIN_USER")
	password := os.Getenv("TWP_ADMIN_PASSWORD")
	db_password, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	err = pg.Queries.CreateAdmin(context.Background(), db.CreateAdminParams{Username: admin_name, Password: string(db_password)})
	return err
}
