package model

import (
	"errors"
	"io"
	"strconv"

	"github.com/99designs/gqlgen/graphql"
	"github.com/google/uuid"
)

type UUID uuid.UUID

func MarshalUUID(u uuid.UUID) graphql.Marshaler {
	_, err := uuid.Parse(u.String())
	if err != nil {
		return nil
	}

	return graphql.WriterFunc(func(w io.Writer) {
		io.WriteString(w, strconv.Quote(u.String()))
	})
}

func UnmarshalUUID(v interface{}) (uuid.UUID, error) {
	if tmpStr, ok := v.(string); ok {
		return uuid.Parse(tmpStr)
	}
	return uuid.UUID{}, errors.New("uuid should be string")
}
