package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/buidl-labs/celo-voting-validator-backend/graph/generated"
	"github.com/buidl-labs/celo-voting-validator-backend/graph/model"
)

func (r *queryResolver) ValidatorGroups(ctx context.Context) ([]*model.ValidatorGroup, error) {
	var vgs_db []*model.ValidatorGroup
	if err := r.DB.Model(&vgs_db).Relation("Validators").Relation("Stats").Relation("Validators.Stats").Select(); err != nil {
		log.Println(err)
		return vgs_db, err
	}

	// log.Println(len(vgs_db))
	// var vgs []*model.ValidatorGroup
	// vg := &model.ValidatorGroup{
	// 	ID:   "some id",
	// 	Name: "Some Validator",
	// }
	// return append(vgs, vg), nil
	return vgs_db, nil
}

func (r *validatorGroupStatsResolver) EstimatedDrr(ctx context.Context, obj *model.ValidatorGroupStats) (float64, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *validatorGroupStatsResolver) EstimatedMrr(ctx context.Context, obj *model.ValidatorGroupStats) (float64, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *validatorGroupStatsResolver) EstimatedArr(ctx context.Context, obj *model.ValidatorGroupStats) (float64, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *validatorGroupStatsResolver) EpochNum(ctx context.Context, obj *model.ValidatorGroupStats) (int, error) {
	epoch := new(model.Epoch)
	if err := r.DB.Model(epoch).Where("id = ?", obj.EpochId).Select(); err != nil {
		return -1, err
	}
	// PrettyPrint(obj)
	return epoch.Number, nil
}

func (r *validatorStatsResolver) EpochNum(ctx context.Context, obj *model.ValidatorStats) (int, error) {
	epoch := new(model.Epoch)
	if err := r.DB.Model(epoch).Where("id = ?", obj.EpochId).Select(); err != nil {
		return -1, err
	}
	// PrettyPrint(obj)
	return epoch.Number, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// ValidatorGroupStats returns generated.ValidatorGroupStatsResolver implementation.
func (r *Resolver) ValidatorGroupStats() generated.ValidatorGroupStatsResolver {
	return &validatorGroupStatsResolver{r}
}

// ValidatorStats returns generated.ValidatorStatsResolver implementation.
func (r *Resolver) ValidatorStats() generated.ValidatorStatsResolver {
	return &validatorStatsResolver{r}
}

type queryResolver struct{ *Resolver }
type validatorGroupStatsResolver struct{ *Resolver }
type validatorStatsResolver struct{ *Resolver }

func PrettyPrint(v interface{}) (err error) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err == nil {
		fmt.Println(string(b))
	}
	return
}
