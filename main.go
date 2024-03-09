package main

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"graduation-project-ent/common"
	"graduation-project-ent/internal/db"
	"graduation-project-ent/pkg/ent_work/user"
)

func main() {
	ctx := context.Background()
	users, err := db.DB.User.Query().Where(user.DeletedAt(common.ZeroTime)).First(ctx)
	if err != nil {
		logrus.Errorf("%+v", err)
	}
	fmt.Printf("%+v\n", users)
}
