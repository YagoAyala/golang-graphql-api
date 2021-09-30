package resolver

import (
	"context"
	"multiplier/website/model"
	"multiplier/website/repository"

	"github.com/Kichiyaki/goutil/safeptr"
	"github.com/google/uuid"
)

func (r *queryResolver) Contact(ctx context.Context, uid uuid.UUID) (*model.Contact, error) {
	return repository.FetchContact(ctx, uid)
}

func (r *queryResolver) Contacts(ctx context.Context, search *string, filter *model.ContactFilterInput,
	limit *int, page *int, sort *model.ContactSortInput) (*model.ContactList, error) {

	var err error
	var list = &model.ContactList{}

	list.Items, list.TotalCount, err = repository.FetchAllContacts(ctx, &model.ContactFilterConfig{
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

func (r *mutationResolver) CreateContact(ctx context.Context, input model.CreateContactInput) (*model.Contact, error) {
	return repository.StoreContact(ctx, input)
}

func (r *mutationResolver) UpdateContact(ctx context.Context, uid uuid.UUID, input model.UpdateContactInput) (*model.Contact, error) {
	return repository.UpdateContact(ctx, uid, input)
}

func (r *mutationResolver) DeleteContact(ctx context.Context, uid uuid.UUID) (*string, error) {
	return repository.DeleteContact(ctx, uid)
}
