package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Newsletter struct {
	bun.BaseModel `bun:"newsletter"`

	ID        int64
	UID       uuid.UUID  `json:"uid" bun:",nullzero,notnull,type:uuid"`
	Email     string     `json:"email"`
	CreatedAt time.Time  `json:"createdAt" bun:",nullzero,notnull"`
	UpdatedAt time.Time  `json:"updatedAt" bun:",nullzero,notnull"`
	DeletedAt *time.Time `bun:",soft_delete"`
}

type CreateNewsletterInput struct {
	Email string `json:"email"`
}

type UpdateNewsletterInput struct {
	Email *string `json:"email"`
}

type NewsletterList struct {
	TotalCount int           `json:"totalCount"`
	Items      []*Newsletter `json:"items"`
}

type NewsletterFilterInput struct {
	UID       *FilterEqualTypeInput `json:"uid"`
	Email     *FilterMatchTypeInput `json:"email"`
	CreatedAt *FilterRangeTypeInput `json:"createdAt"`
	UpdatedAt *FilterRangeTypeInput `json:"updatedAt"`
}

type NewsletterSortInput struct {
	UID       *SortEnum `json:"uid"`
	Email     *SortEnum `json:"email"`
	CreatedAt *SortEnum `json:"createdAt"`
	UpdatedAt *SortEnum `json:"updatedAt"`
}

type NewsletterFilterConfig struct {
	Filter *NewsletterFilterInput `json:¨"filter"`
	Sort   *NewsletterSortInput   `json:¨"sort"`
	*FilterConfigBase
}

func (input *CreateNewsletterInput) ToNewsletter() *Newsletter {
	return &Newsletter{
		Email: input.Email,
	}
}

func (f *NewsletterFilterInput) ToQuery(q *bun.SelectQuery) *bun.SelectQuery {

	switch {
	case f.UID != nil:
		q = q.Where("uid = ?", f.UID.Eq)

	case f.UID != nil:
		q = q.Where("uid IN (?)", bun.In(f.UID.In))

	case f.Email != nil:
		q = q.Where("email ILIKE ?", "%"+*f.Email.Match+"%")

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

func (i *UpdateNewsletterInput) ApplyUpdate(q *bun.UpdateQuery) *bun.UpdateQuery {

	if i.Email != nil {
		q = q.Set("email = ?", i.Email)
	}

	q = q.Set("updated_at = ?", time.Now())

	return q
}
