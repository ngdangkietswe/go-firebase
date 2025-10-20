/**
 * Author : ngdangkietswe
 * Since  : 10/17/2025
 */

package repository

import (
	"context"
	"go-firebase/internal/data/ent"
	"go-firebase/internal/data/ent/user"
	"go-firebase/internal/request"

	"github.com/google/uuid"
)

type userRepo struct {
	cli *ent.Client
}

func (r *userRepo) Save(ctx context.Context, tx *ent.Tx, request *request.CreateUserRequest, firebaseUID string) (*ent.User, error) {
	builder := tx.User.Create().
		SetEmail(request.Email).
		SetFirebaseUID(firebaseUID)

	if request.FirstName != "" {
		builder = builder.SetFirstName(request.FirstName)
	}

	if request.LastName != "" {
		builder = builder.SetLastName(request.LastName)
	}

	if request.FirstName != "" && request.LastName != "" {
		builder = builder.SetDisplayName(request.FirstName + " " + request.LastName)
	}

	return builder.Save(ctx)
}

func (r *userRepo) FindByEmail(ctx context.Context, email string) (*ent.User, error) {
	return r.cli.User.Query().Where(user.Email(email)).Only(ctx)
}

func (r *userRepo) FindByID(ctx context.Context, id uuid.UUID) (*ent.User, error) {
	return r.cli.User.Query().Where(user.ID(id)).Only(ctx)
}

func (r *userRepo) FindByEmailOrID(ctx context.Context, identifier string) (*ent.User, error) {
	if _, err := uuid.Parse(identifier); err != nil {
		return r.cli.User.Query().
			Where(user.Email(identifier)).
			Only(ctx)
	}
	return r.cli.User.Query().
		Where(user.ID(uuid.MustParse(identifier))).
		Only(ctx)
}

func (r *userRepo) ExistsByID(ctx context.Context, id uuid.UUID) (bool, error) {
	return r.cli.User.Query().Where(user.ID(id)).Exist(ctx)
}

func NewUserRepository(cli *ent.Client) UserRepository {
	return &userRepo{cli: cli}
}
