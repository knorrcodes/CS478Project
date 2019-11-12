package graphql

import (
	"context"
	"errors"

	"github.com/99designs/gqlgen/graphql"

	"koala.pos/src/auth"
)

func NewConfig() Config {
	c := Config{Resolvers: &Resolver{}}
	c.Directives.MinRole = func(ctx context.Context, obj interface{}, next graphql.Resolver, role Role) (interface{}, error) {
		server := auth.GetServerFromContext(ctx)
		if server == nil {
			return nil, errors.New("Failed to check permissions")
		}

		switch role {
		case RoleManager:
			if !server.Manager {
				return nil, errors.New("insufficient privilages")
			}
		case RoleServer:
			if server.CustCode != "" {
				return nil, errors.New("insufficient privilages")
			}
		}

		return next(ctx)
	}

	return c
}
