package graphql

import (
	"github.com/jinzhu/gorm"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

func NewUserHandler(db *gorm.DB) (*handler.Handler, error) {
	schema, err := graphql.NewSchema(
		graphql.SchemaConfig{
			Query:    queries(db),
			Mutation: mutation(db),
		},
	)
	if err != nil {
		return nil, err
	}

	return handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	}), nil
}
