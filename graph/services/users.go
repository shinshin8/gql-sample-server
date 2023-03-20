package services

import (
	"context"
	"fmt"
	"github.com/shinshin8/gql-sample-server/graph/db"
	"github.com/shinshin8/gql-sample-server/graph/model"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type userService struct {
	exec boil.ContextExecutor
}

type UserService interface {
	GetUserByName(ctx context.Context, name string) (*model.User, error)
}

func convertUser(user *db.User) *model.User {
	return &model.User{
		ID:   user.ID,
		Name: user.Name,
	}
}

func (u *userService) GetUserByName(ctx context.Context, name string) (*model.User, error) {
	user, err := db.Users(
		qm.Select(db.UserTableColumns.ID, db.UserTableColumns.Name),
		db.UserWhere.Name.EQ(name),
	).One(ctx, u.exec)

	if err != nil {
		return nil, fmt.Errorf("failed to get user name: %w", err)
	}
	return convertUser(user), nil
}
