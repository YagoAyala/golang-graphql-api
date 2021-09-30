package repository

import (
	"context"

	"multiplier/website/app"
	"multiplier/website/model"

	"github.com/google/uuid"
)

func FetchNewsletterByEmail(ctx context.Context, email string) (*model.Newsletter, error) {
	newsletter := &model.Newsletter{}

	err := app.DB.NewSelect().
		Model(newsletter).
		Where("email = ?", email).
		Scan(ctx)

	if err != nil {
		return nil, err
	}

	return newsletter, nil
}

func FetchNewsletter(ctx context.Context, uid uuid.UUID) (*model.Newsletter, error) {
	newsletter := &model.Newsletter{}

	err := app.DB.NewSelect().
		Model(newsletter).
		Where("uid = ?", uid).
		Scan(ctx)

	if err != nil {
		return nil, err
	}

	return newsletter, nil
}

func FetchAllNewsletter(ctx context.Context, f *model.NewsletterFilterConfig) ([]*model.Newsletter, int, error) {
	newsletter := []*model.Newsletter{}

	total, err := app.DB.NewSelect().
		Model(&newsletter).
		Limit(f.Limit).
		Offset(f.ToOffset()).
		Apply(f.Filter.ToQuery).
		ScanAndCount(ctx)

	if err != nil {
		return nil, 0, err
	}

	return newsletter, total, nil
}

func StoreNewsletter(ctx context.Context, input model.CreateNewsletterInput) (*model.Newsletter, error) {

	item := input.ToNewsletter()

	_, err := app.DB.NewInsert().Model(item).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return item, nil
}

func UpdateNewsletter(ctx context.Context, uid uuid.UUID, input model.UpdateNewsletterInput) (*model.Newsletter, error) {
	newsletter := &model.Newsletter{}

	_, err := app.DB.NewUpdate().
		Model(newsletter).
		Where("uid = ?", uid).
		Apply(input.ApplyUpdate).
		Returning("*").
		Exec(ctx)

	if err != nil {
		return nil, err
	}

	return newsletter, nil
}

func DeleteNewsletter(ctx context.Context, uid uuid.UUID) (*string, error) {
	newsletter := &model.Newsletter{}

	errorDelete := "Error to delete"
	sucessDelete := "Deleted"

	err := app.DB.NewSelect().
		Model(newsletter).
		Where("uid = ?", uid).
		Scan(ctx)
		//TODO: melhorar o retorno, pois não faz sentido retornar duas strings
	if err != nil {
		return &errorDelete, err
	}

	_, err = app.DB.NewDelete().
		Model(newsletter).
		Where("uid = ?", uid).
		Exec(ctx)

	if err != nil {
		return &errorDelete, err
	}

	return &sucessDelete, nil
}
