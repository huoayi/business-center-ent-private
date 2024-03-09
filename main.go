package main

import (
	"context"
	"fmt"
	"github.com/huoayi/business-center-ent-private/common"
	"github.com/huoayi/business-center-ent-private/internal/db"
	"github.com/huoayi/business-center-ent-private/pkg/ent_work/user"
	"github.com/sirupsen/logrus"
)

func main() {
	ctx := context.Background()
	users, err := db.DB.User.Query().Where(user.DeletedAt(common.ZeroTime)).First(ctx)
	if err != nil {
		logrus.Errorf("%+v", err)
	}
	fmt.Printf("%+v\n", users)
}
