package repository

import (
	"context"

	"multiplier/website/app"
	"multiplier/website/model"

	"github.com/google/uuid"
)

func FetchContact(ctx context.Context, uid uuid.UUID) (*model.Contact, error) {
	contact := &model.Contact{}
	err := app.DB.NewSelect().
		Model(contact).
		Where("uid = ?", uid).
		Scan(ctx)

	if err != nil {
		return nil, err
	}

	return contact, nil
}

func FetchAllContacts(ctx context.Context, f *model.ContactFilterConfig) ([]*model.Contact, int, error) {
	contacts := []*model.Contact{}

	total, err := app.DB.NewSelect().
		Model(&contacts).
		Limit(f.Limit).
		Offset(f.ToOffset()).
		Apply(f.Filter.ToQuery).
		ScanAndCount(ctx)

	if err != nil {
		return nil, 0, err
	}

	return contacts, total, nil
}

func StoreContact(ctx context.Context, input model.CreateContactInput) (*model.Contact, error) {
	// TODO: remove thiss
	item := input.ToContact()
	_, err := app.DB.NewInsert().Model(item).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return item, nil
}

func UpdateContact(ctx context.Context, uid uuid.UUID, input model.UpdateContactInput) (*model.Contact, error) {
	contact := &model.Contact{}

	_, err := app.DB.NewUpdate().
		Model(contact).
		Where("uid = ?", uid).
		Apply(input.ApplyUpdate).
		Returning("*").
		Exec(ctx)

	if err != nil {
		return nil, err
	}

	return contact, nil
}

func DeleteContact(ctx context.Context, uid uuid.UUID) (*string, error) {
	contact := &model.Contact{}

	errorDelete := "Error to delete"
	sucessDelete := "Deleted"

	err := app.DB.NewSelect().
		Model(contact).
		Where("uid = ?", uid).
		Scan(ctx)

	if err != nil {
		return &errorDelete, err
	}

	_, err = app.DB.NewDelete().
		Model(contact).
		Where("uid = ?", uid).
		Exec(ctx)

	if err != nil {
		return &errorDelete, err
	}

	return &sucessDelete, nil
}
