package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Contact struct {
	bun.BaseModel `bun:"contact"`

	ID          int64
	UID         uuid.UUID  `json:"uid" bun:",nullzero,notnull,type:uuid"`
	Name        string     `json:"name"`
	Email       string     `json:"email"`
	PhoneNumber string     `json:"phoneNumber"`
	Message     string     `json:"message"`
	IsRead      bool       `json:"isRead"`
	CreatedAt   time.Time  `json:"createdAt" bun:",nullzero,notnull"`
	UpdatedAt   time.Time  `json:"updatedAt" bun:",nullzero,notnull"`
	DeletedAt   *time.Time `bun:",soft_delete"`
}

type CreateContactInput struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	Message     string `json:"message"`
	IsRead      bool   `json:"isRead"`
}

type UpdateContactInput struct {
	Name        *string `json:"name"`
	Email       *string `json:"email"`
	PhoneNumber *string `json:"phoneNumber"`
	Message     *string `json:"message"`
	IsRead      *bool   `json:"isRead"`
}

type ContactList struct {
	TotalCount int        `json:"totalCount"`
	Items      []*Contact `json:"items"`
}

type ContactFilterInput struct {
	UID         *FilterEqualTypeInput `json:"uid"`
	Name        *FilterMatchTypeInput `json:"name"`
	Email       *FilterMatchTypeInput `json:"email"`
	PhoneNumber *FilterMatchTypeInput `json:"phoneNumber"`
	Message     *FilterMatchTypeInput `json:"message"`
	IsRead      *FilterEqualTypeInput `json:"isRead"`
	CreatedAt   *FilterRangeTypeInput `json:"createdAt"`
	UpdatedAt   *FilterRangeTypeInput `json:"updatedAt"`
}

type ContactSortInput struct {
	UID         *SortEnum `json:"uid"`
	Name        *SortEnum `json:"name"`
	Email       *SortEnum `json:"email"`
	PhoneNumber *SortEnum `json:"phoneNumber"`
	IsRead      *SortEnum `json:"isRead"`
	CreatedAt   *SortEnum `json:"createdAt"`
	UpdatedAt   *SortEnum `json:"updatedAt"`
}

type ContactFilterConfig struct {
	Filter *ContactFilterInput `json:¨"filter"`
	Sort   *ContactSortInput   `json:¨"sort"`
	*FilterConfigBase
}

func (input *CreateContactInput) ToContact() *Contact {

	return &Contact{
		Name:  input.Name,
		Email: input.Email,

		PhoneNumber: input.PhoneNumber,
		Message:     input.Message,
	}
}

func (f *ContactFilterInput) ToQuery(q *bun.SelectQuery) *bun.SelectQuery {

	switch {
	case f.UID != nil:
		q = q.Where("uid = ?", f.UID.Eq)

	case f.UID != nil:
		q = q.Where("uid IN (?)", bun.In(f.UID.In))

	case f.Name != nil:
		q = q.Where("name ILIKE ?", "%"+*f.Name.Match+"%")

	case f.Email != nil:
		q = q.Where("email ILIKE ?", "%"+*f.Email.Match+"%")

	case f.PhoneNumber != nil:
		q = q.Where("phone_number ILIKE ?", "%"+*f.PhoneNumber.Match+"%")

	case f.Message != nil:
		q = q.Where("message ILIKE ?", "%"+*f.Message.Match+"%")

	case f.IsRead != nil:
		q = q.Where("is_read = ?", f.IsRead.Eq)

	case f.IsRead != nil:
		q = q.Where("is_read = ?", bun.In(f.IsRead.In))

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

func (i *UpdateContactInput) ApplyUpdate(q *bun.UpdateQuery) *bun.UpdateQuery {

	switch {
	case i.Name != nil:
		q = q.Set("name = ?", i.Name)

	case i.Email != nil:
		q = q.Set("email = ?", i.Email)

	case i.PhoneNumber != nil:
		q = q.Set("phone_number = ?", i.PhoneNumber)

	case i.Message != nil:
		q = q.Set("message = ?", i.Message)

	case i.IsRead != nil:
		q = q.Set("is_read = ?", i.IsRead)
	}

	q = q.Set("updated_at = ?", time.Now())

	return q
}
