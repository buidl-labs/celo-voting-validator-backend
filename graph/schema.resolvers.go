package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"

	"github.com/buidl-labs/celo-voting-validator-backend/graph/generated"
	"github.com/buidl-labs/celo-voting-validator-backend/graph/model"
)

func (r *epochResolver) StartBlock(ctx context.Context, obj *model.Epoch) (int, error) {
	return int(obj.StartBlock), nil
}

func (r *epochResolver) EndBlock(ctx context.Context, obj *model.Epoch) (int, error) {
	return int(obj.EndBlock), nil
}

func (r *epochResolver) Number(ctx context.Context, obj *model.Epoch) (int, error) {
	return int(obj.Number), nil
}

func (r *mutationResolver) UpdateVGSocialInfo(ctx context.Context, vgID string, email *string, websiteURL *string, discordTag *string, twitterUsername *string, geographicLocation *string) (*model.ValidatorGroup, error) {
	vg := new(model.ValidatorGroup)
	if err := r.DB.Model(vg).Where("ID = ?", vgID).Relation("Validators").Limit(1).Select(); err != nil {
		return vg, err
	}
	vg_updated := false
	if email != nil {
		vg.Email = *email
		vg_updated = true
	}
	if websiteURL != nil {
		vg.WebsiteURL = *websiteURL
		vg_updated = true
	}
	if discordTag != nil {
		vg.DiscordTag = *discordTag
		vg_updated = true
	}
	if twitterUsername != nil {
		vg.TwitterUsername = *twitterUsername
		vg_updated = true
	}
	if geographicLocation != nil {
		vg.GeographicLocation = *geographicLocation
		vg_updated = true
	}

	if vg_updated {
		vg.TransparencyScore = calculateTransparencyScore(vg)
		_, err := r.DB.Model(vg).WherePK().Update()
		if err != nil {
			return vg, err
		}
	}

	return vg, nil
}

func (r *queryResolver) ValidatorGroups(ctx context.Context, sortByScore *bool, limit *int) ([]*model.ValidatorGroup, error) {
	var vgs_db []*model.ValidatorGroup
	if limit != nil {
		if *limit <= 0 {
			return vgs_db, errors.New("limit needs to be more than 0")
		}

		err := r.DB.Model(&vgs_db).Relation("Validators").Limit(*limit).Select()
		if err != nil {
			return vgs_db, err
		}
		return vgs_db, nil
	}

	if err := r.DB.Model(&vgs_db).Relation("Validators").Select(); err != nil {
		return vgs_db, err
	}
	return vgs_db, nil
}

func (r *queryResolver) ValidatorGroup(ctx context.Context, address string) (*model.ValidatorGroup, error) {
	vg := new(model.ValidatorGroup)
	if err := r.DB.Model(vg).Where("address = ?", address).Relation("Validators").Limit(1).Select(); err != nil {
		return vg, err
	}
	return vg, nil
}

func (r *validatorGroupResolver) EpochRegisteredAt(ctx context.Context, obj *model.ValidatorGroup) (int, error) {
	return int(obj.EpochRegisteredAt), nil
}

func (r *validatorGroupResolver) EpochsServed(ctx context.Context, obj *model.ValidatorGroup) (int, error) {
	return int(obj.EpochsServed), nil
}

func (r *validatorGroupResolver) RecievedVotes(ctx context.Context, obj *model.ValidatorGroup) (int, error) {
	return int(obj.RecievedVotes), nil
}

func (r *validatorGroupResolver) AvailableVotes(ctx context.Context, obj *model.ValidatorGroup) (int, error) {
	return int(obj.AvailableVotes), nil
}

func (r *validatorGroupResolver) LockedCelo(ctx context.Context, obj *model.ValidatorGroup) (int, error) {
	return int(obj.LockedCelo), nil
}

// Epoch returns generated.EpochResolver implementation.
func (r *Resolver) Epoch() generated.EpochResolver { return &epochResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// ValidatorGroup returns generated.ValidatorGroupResolver implementation.
func (r *Resolver) ValidatorGroup() generated.ValidatorGroupResolver {
	return &validatorGroupResolver{r}
}

type epochResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type validatorGroupResolver struct{ *Resolver }
