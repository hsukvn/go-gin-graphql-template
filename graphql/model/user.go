package model

import (
	"os/user"
)

type User struct {
	UID    string   `json:"uid"`
	GID    string   `json:"gid"`
	Name   string   `json:"name"`
	Home   string   `json:"home"`
	Groups []*Group `json:"groups"`
}

func NewUserByUID(uid string) (*User, error) {
	u, err := user.LookupId(uid)
	if err != nil {
		return nil, err
	}

	return NewUser(u)
}

func NewUserByName(name string) (*User, error) {
	u, err := user.Lookup(name)
	if err != nil {
		return nil, err
	}

	return NewUser(u)
}

func NewUser(u *user.User) (*User, error) {
	gids, err := u.GroupIds()
	if err != nil {
		return nil, err
	}

	return &User{
		UID:    u.Uid,
		GID:    u.Gid,
		Name:   u.Name,
		Home:   u.HomeDir,
		Groups: getGroups(gids),
	}, nil
}

func getGroups(gids []string) []*Group {
	modelGroups := make([]*Group, len(gids))

	for i, id := range gids {
		g, err := user.LookupGroupId(id)
		if err != nil {
			continue
		}

		modelGroups[i] = &Group{
			GID:  g.Gid,
			Name: g.Name,
		}
	}

	return modelGroups
}
