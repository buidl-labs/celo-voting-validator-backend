package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"log"

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

func (r *mutationResolver) UpdateVGSocialInfo(ctx context.Context, vgID string, email *string, discordTag *string, twitterUsername *string, geographicLocation *string) (*model.ValidatorGroup, error) {
	vg := new(model.ValidatorGroup)
	if err := r.DB.Model(vg).Where("ID = ?", vgID).Relation("Validators").Limit(1).Select(); err != nil {
		return vg, err
	}
	vg_updated := false
	if email != nil {
		// Validate Email before updating.
		vg.Email = *email
		vg_updated = true
	}

	if discordTag != nil {
		// Validate discordTag before updating
		vg.DiscordTag = *discordTag
		vg_updated = true
	}

	if twitterUsername != nil {
		// Validate twitterUsername before updating
		vg.TwitterUsername = *twitterUsername
		vg_updated = true
	}
	if geographicLocation != nil {
		// Validate geographicLocation before updating
		vg.GeographicLocation = *geographicLocation
		vg_updated = true
	}

	if vg_updated {
		_, err := r.DB.Model(vg).WherePK().Update()
		if err != nil {
			return vg, err
		}
	}

	return vg, nil
}

func (r *queryResolver) ValidatorGroups(ctx context.Context, sortByScore *bool, limit *int) ([]*model.ValidatorGroup, error) {

	const OrderExpression = "(validator_group.performance_score * 0.9 + validator_group.transparency_score * 0.1) desc"

	var vgs_db []*model.ValidatorGroup

	doSort := false
	if sortByScore != nil {
		doSort = *sortByScore
	}

	if limit != nil {
		if *limit <= 0 {
			return vgs_db, errors.New("limit needs to be more than 0")
		}

		// If limit is provided and sortByScore is true
		if doSort {
			err := r.DB.Model(&vgs_db).Relation("Validators").Limit(*limit).OrderExpr(OrderExpression).Select()
			if err != nil {
				log.Println(err)
				return vgs_db, err
			}
			return vgs_db, nil
		}

		// If limit is provided but sortByScore is false
		err := r.DB.Model(&vgs_db).Relation("Validators").Limit(*limit).Select()
		if err != nil {
			return vgs_db, err
		}
		return vgs_db, nil
	}

	// If limit is not provided and sortByScore is true
	if doSort {
		err := r.DB.Model(&vgs_db).Relation("Validators").OrderExpr(OrderExpression).Select()
		if err != nil {
			return vgs_db, err
		}
		return vgs_db, nil
	}

	// If limit isn't provided and sortByScore is false.
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
