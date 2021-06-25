package security

import (
	"context"
	"golang-project/models"
)

type AuthRepository interface {
	AddSession(ctx context.Context, session *models.Session) error
}