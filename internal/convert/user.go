package convert

import (
	apiv1 "learn-ddd/gen/api/v1"
	"learn-ddd/internal/domain/model"
)

func ToConnectUser(u *model.User) *apiv1.User {
	return &apiv1.User{
		Id:   u.ID,
		Name: u.Name,
	}
}

func ToConnectUsers(us []*model.User) []*apiv1.User {
	ret := make([]*apiv1.User, len(us))
	for i, v := range us {
		vv := ToConnectUser(v)
		ret[i] = vv
	}
	return ret
}
