package resolver

import (
	"context"
	"multiplier/website/model"
	"multiplier/website/repository"

	"github.com/Kichiyaki/goutil/safeptr"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func (r *queryResolver) Newsletter(ctx context.Context, uid uuid.UUID) (*model.Newsletter, error) {
	return repository.FetchNewsletter(ctx, uid)
}

func (r *queryResolver) Newsletters(ctx context.Context, search *string, filter *model.NewsletterFilterInput,
	limit *int, page *int, sort *model.NewsletterSortInput) (*model.NewsletterList, error) {

	var err error
	var list = &model.NewsletterList{}

	list.Items, list.TotalCount, err = repository.FetchAllNewsletter(ctx, &model.NewsletterFilterConfig{
		Filter: filter,
		Sort:   sort,
		FilterConfigBase: &model.FilterConfigBase{
			Search: search,
			Limit:  safeptr.SafeIntPointer(limit, 30),
			Page:   safeptr.SafeIntPointer(page, 1),
		},
	})

	if err != nil {
		return nil, err
	}

	return list, nil
}

func (r *mutationResolver) CreateNewsletter(ctx context.Context, input model.CreateNewsletterInput) (*model.Newsletter, error) {
	data, err := repository.FetchNewsletterByEmail(ctx, input.Email)

	if err == nil {
		return nil, errors.New("E-mail is alredy on database")
	}

	if data != nil {
		return nil, err
	}

	return repository.StoreNewsletter(ctx, input)
}

func (r *mutationResolver) UpdateNewsletter(ctx context.Context, uid uuid.UUID, input model.UpdateNewsletterInput) (*model.Newsletter, error) {
	return repository.UpdateNewsletter(ctx, uid, input)
}

func (r *mutationResolver) DeleteNewsletter(ctx context.Context, uid uuid.UUID) (*string, error) {
	return repository.DeleteNewsletter(ctx, uid)
}
