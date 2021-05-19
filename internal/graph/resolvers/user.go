package resolvers

import (
	"context"

	"github.com/xinliangnote/go-gin-api/internal/graph/model"

	"github.com/pkg/errors"
)

func (r *queryResolver) BySex(ctx context.Context, sex string) ([]*model.User, error) {
	if sex == "" {
		return nil, errors.New("sex required")
	}

	// Simulation data
	var users []*model.User
	users = append(users, &model.User{ID: "1", Name: "Tom", Sex: sex, Mobile: "13266666666"})
	users = append(users, &model.User{ID: "1", Name: "Jack", Sex: sex, Mobile: "13288888888"})

	return users, nil
}

func (r *mutationResolver) UpdateUserMobile(ctx context.Context, data model.UpdateUserMobileInput) (*model.User, error) {
	if data.ID == "" {
		return nil, errors.New("id required")
	}

	if data.Mobile == "" {
		return nil, errors.New("mobile required")
	}

	// Simulation data
	user := new(model.User)
	user.ID = data.ID
	user.Mobile = data.Mobile
	user.Sex = "male"
	user.Name = "Jack"

	// Operational database
	//userData, err := r.userService.GetUserByUserName(r.getCoreContextByCtx(ctx), "test_user")
	//if err != nil {
	//	return nil, err
	//}

	return user, nil
}
