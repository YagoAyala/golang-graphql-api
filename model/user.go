package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"user"`

	ID        int64
	UID       uuid.UUID  `json:"uid" bun:",nullzero,notnull,type:uuid"`
	Name      string     `json:"name"`
	Username  string     `json:"username"`
	Password  string     `json:"password"`
	Picture   string     `json:"picture"`
	IsActive  bool       `json:"isActive"`
	CreatedAt time.Time  `json:"createdAt" bun:",nullzero,notnull"`
	UpdatedAt time.Time  `json:"updatedAt" bun:",nullzero,notnull"`
	DeletedAt *time.Time `bun:",soft_delete"`
}

type CreateUserInput struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Picture  string `json:"picture"`
	IsActive bool   `json:"isActive"`
}

type UpdateUserInput struct {
	Name     *string `json:"name"`
	Username *string `json:"username"`
	Password *string `json:"password"`
	Picture  *string `json:"picture"`
	IsActive *bool   `json:"isRead"`
}

type UserList struct {
	TotalCount int     `json:"totalCount"`
	Items      []*User `json:"items"`
}

type UserFilterInput struct {
	UID       *FilterEqualTypeInput `json:"uid"`
	Name      *FilterMatchTypeInput `json:"name"`
	Username  *FilterMatchTypeInput `json:"username"`
	IsActive  *FilterEqualTypeInput `json:"isActive"`
	CreatedAt *FilterRangeTypeInput `json:"createdAt"`
	UpdatedAt *FilterRangeTypeInput `json:"updatedAt"`
}

type UserSortInput struct {
	UID       *SortEnum `json:"uid"`
	Name      *SortEnum `json:"name"`
	Username  *SortEnum `json:"username"`
	Password  *SortEnum `json:"password"`
	Picture   *SortEnum `json:"picture"`
	IsActive  *SortEnum `json:"isActive"`
	CreatedAt *SortEnum `json:"createdAt"`
	UpdatedAt *SortEnum `json:"updatedAt"`
}

type UserFilterConfig struct {
	Filter *UserFilterInput `json:¨"filter"`
	Sort   *UserSortInput   `json:¨"sort"`
	*FilterConfigBase
}

func (input *CreateUserInput) ToUser() *User {

	return &User{
		Name:     input.Name,
		Username: input.Username,

		Password: input.Password,
		Picture:  input.Picture,
	}
}

func (f *UserFilterInput) ToQuery(q *bun.SelectQuery) *bun.SelectQuery {

	switch {
	case f.UID != nil:
		q = q.Where("uid = ?", f.UID.Eq)

	case f.UID != nil:
		q = q.Where("uid IN (?)", bun.In(f.UID.In))

	case f.Name != nil:
		q = q.Where("name ILIKE ?", "%"+*f.Name.Match+"%")

	case f.Username != nil:
		q = q.Where("username ILIKE ?", "%"+*f.Username.Match+"%")

	case f.IsActive != nil:
		q = q.Where("is_active = ?", f.IsActive.Eq)

	case f.IsActive != nil:
		q = q.Where("is_active = ?", f.IsActive.In)

	case f.CreatedAt != nil:
		q = q.Where("created_at <= ?", f.CreatedAt.From)

	case f.CreatedAt != nil:
		q = q.Where("created_at >= ?", f.CreatedAt.To)

	case f.UpdatedAt != nil:
		q = q.Where("updated_at <= ?", f.UpdatedAt.From)

	case f.UpdatedAt != nil:
		q = q.Where("updated_at >= ?", f.UpdatedAt.To)
	}

	return q
}

// func (i *UpdateUserInput) ApplyUpdate(q *)
