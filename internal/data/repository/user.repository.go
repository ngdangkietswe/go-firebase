/**
 * Author : ngdangkietswe
 * Since  : 10/17/2025
 */

package repository

import (
	"context"
	"go-firebase/internal/data/ent"
	"go-firebase/internal/data/ent/user"
	"go-firebase/pkg/request"
	"go-firebase/pkg/util"

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

func (r *userRepo) FindAll(ctx context.Context, request *request.ListUserRequest) ([]*ent.User, int, error) {
	query := r.cli.User.Query()

	if request.Search != "" {
		query = query.Where(user.Or(
			user.EmailContainsFold(request.Search),
			user.FirstNameContainsFold(request.Search),
			user.LastNameContainsFold(request.Search),
		))
	}

	totalItems, err := query.Clone().Count(ctx)
	if err != nil {
		return nil, 0, err
	}

	items, err := query.Clone().Order(util.ToSortOrder(request.Paginate)).
		Offset(request.Paginate.Page * request.Paginate.PageSize).
		Limit(request.Paginate.PageSize).
		All(ctx)

	return items, totalItems, err
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

func (r *userRepo) FindByFirebaseUID(ctx context.Context, firebaseUID string) (*ent.User, error) {
	return r.cli.User.Query().Where(user.FirebaseUID(firebaseUID)).Only(ctx)
}

func (r *userRepo) ExistsByID(ctx context.Context, id uuid.UUID) (bool, error) {
	return r.cli.User.Query().Where(user.ID(id)).Exist(ctx)
}

func NewUserRepository(cli *ent.Client) UserRepository {
	return &userRepo{cli: cli}
}
