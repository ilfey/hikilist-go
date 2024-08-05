package sql

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

type User struct {
	ID       uint
	Username string
	Email    *string
	Password string
}

func (u *User) TableName() string {
	return "users"
}

type Pet struct {
	ID      uint
	Name    string
	OwnerID uint
	Owner   *User
}

func (p *Pet) TableName() string {
	return "users"
}

func TestSelectFromStruct(t *testing.T) {
	assert.NotPanics(t, func() {
		sel := SelectFromStruct(&User{})
		assert.Equalf(t, "users", sel.table, "table name should be 'users'")

		sel = SelectFromStruct(&User{}, "user")
		assert.Equalf(t, "user", sel.table, "table name should be 'user'")
	})

	assert.Panics(t, func() { SelectFromStruct(&struct{}{}) })
}

func TestSelectWhere(t *testing.T) {
	sel := SelectFromStruct(&User{})

	sel.Where("id = 1")

	assert.Equal(
		t,
		"SELECT users.id, users.username, users.email, users.password FROM users WHERE id = 1;",
		sel.SQL(),
	)
}

func TestSelectGroup(t *testing.T) {
	sel := SelectFromStruct(&User{})

	sel.Group("users.id")

	assert.Equal(
		t,
		"SELECT users.id, users.username, users.email, users.password FROM users GROUP BY users.id;",
		sel.SQL(),
	)
}

func TestSelectOrder(t *testing.T) {
	sel := SelectFromStruct(&User{})

	sel.Order("users.id DESC")

	assert.Equal(
		t,
		"SELECT users.id, users.username, users.email, users.password FROM users ORDER BY users.id DESC;",
		sel.SQL(),
	)
}

func TestSelectOffset(t *testing.T) {
	sel := SelectFromStruct(&User{})

	sel.Offset(10)

	assert.Equal(
		t,
		"SELECT users.id, users.username, users.email, users.password FROM users OFFSET 10;",
		sel.SQL(),
	)
}

func TestSelectLimit(t *testing.T) {
	sel := SelectFromStruct(&User{})

	sel.Limit(5)

	assert.Equal(
		t,
		"SELECT users.id, users.username, users.email, users.password FROM users LIMIT 5;",
		sel.SQL(),
	)
}

func TestSelectResolve(t *testing.T) {
	sel := SelectFromStruct(&Pet{})

	assert.Panics(t, func() {
		sel.Resolve("NonExistField", func(ctx context.Context, p *Pet) error {
			return nil
		})
	})

	sel.Resolve("Owner", func(ctx context.Context, p *Pet) error {
		return nil
	})

	assert.Len(t, sel.resolvers, 1)
	assert.Len(t, sel.fields, 3)

	assert.Panics(t, func() {
		sel.Resolve("Owner", func(ctx context.Context, p *Pet) error {
			return nil
		})
	})
}

func TestSelectAs(t *testing.T) {
	sel := SelectFromStruct(&User{})

	sel.As("u")

	assert.Equal(
		t,
		"SELECT u.id, u.username, u.email, u.password FROM users AS u;",
		sel.SQL(),
	)

	sel.Where(map[string]any{
		"ID": 1,
	})

	assert.Equal(
		t,
		"SELECT u.id, u.username, u.email, u.password FROM users AS u WHERE u.id = 1;",
		sel.SQL(),
	)
}
