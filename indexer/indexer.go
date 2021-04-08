package indexer

import (
	"fmt"
	"log"
	"time"

	"github.com/buidl-labs/celo-voting-validator-backend/graph/model"
	"github.com/go-pg/pg/v10"
	"github.com/machinebox/graphql"
)

func Indexer(DB *pg.DB) {

	const NO_RESULT_ERROR = "pg: no rows in result set"

	client := graphql.NewClient("https://explorer.celo.org/graphiql")

	latestBlock := getLatestBlock(client)
	// electedValidators := getElectedValidators(client, latestBlock)

	currentEpoch := calculateCurrentEpoch(latestBlock)

	epoch := new(model.Epoch)
	if err := DB.Model(epoch).Where("number = ?", currentEpoch).Select(); err != nil {
		fmt.Println(err)
		if err.Error() == NO_RESULT_ERROR {
			// This epoch is not indexed. Add it to DB.

			startBlock := currentEpoch * blocksPerEpoch
			endBlock := startBlock + blocksPerEpoch - 1

			epoch = &model.Epoch{
				StartBlock: startBlock,
				EndBlock:   endBlock,
				Number:     currentEpoch,
				CreatedAt:  time.Now(),
			}
			res, err := DB.Model(epoch).Insert()
			fmt.Println("Added epoch to DB")
			fmt.Println(res.RowsAffected())
			if err != nil {
				panic(err)
			}

		}
	}

	// get validator groups from GraphQL API
	validatorGroups := getValidatorGroups(client)
	fmt.Println(validatorGroups[0].Account.Address)
	vg_db := new(model.ValidatorGroup)
	for _, vg := range validatorGroups {
		// select VG from DB based on address
		if err := DB.Model(vg_db).Where("address = ?", vg.Account.Address).Select(); err != nil {

			if err.Error() == NO_RESULT_ERROR {
				// Insert the VG into the DB
				vg_db = &model.ValidatorGroup{
					Name:          vg.Name,
					Address:       vg.Account.Address,
					VerifiedDNS:   false,
					CreatedAt:     time.Now(),
					NumValidators: vg.NumMembers,
				}

				for _, claim := range vg.Account.Claims.Edges {
					if claim.Node.Type == "domain" {
						vg_db.WebsiteURL = claim.Node.Element
					}
				}

				if _, err := DB.Model(vg_db).Insert(); err != nil {
					fmt.Println(err)
				}
			}
			if err := DB.Model(vg_db).Where("address = ?", vg_db.Address).Select(); err != nil {
				fmt.Println(err)
			}
		}

		if vg.NumMembers != vg_db.NumValidators {
			_, err := DB.Model(vg_db).Set("num_validators = ?", vg.NumMembers).Where("id = ?", vg_db.ID).Update()
			if err != nil {
				log.Println(err)
			}
		}

		v_db := new(model.Validator)
		var validator_stats []*model.ValidatorStats

		for _, validator := range vg.Affiliates.Edges {
			if err := DB.Model(v_db).Where("validator.address = ?", validator.Node.Address).Select(); err != nil {
				if err.Error() == NO_RESULT_ERROR {

					v_db = &model.Validator{
						Name:             fmt.Sprintf("%v", validator.Node.Name),
						ValidatorGroupId: vg_db.ID,
						Address:          validator.Node.Address,
						CreatedAt:        time.Now(),
					}
					if _, err := DB.Model(v_db).Insert(); err != nil {
						fmt.Println(err)
					}

					if err := DB.Model(v_db).Where("validator.address = ?", validator.Node.Address).Select(); err != nil {
						fmt.Println(err)
					}

				}

			}

			v_stats := &model.ValidatorStats{
				AttestationsRequested:  validator.Node.AttestationsRequested,
				AttenstationsFulfilled: validator.Node.AttestationsFulfilled,
				Score:                  validator.Node.Score,
				EpochId:                epoch.ID,
				ValidatorId:            v_db.ID,
				CreatedAt:              time.Now(),
				LastElected:            validator.Node.LastElected,
			}

			validator_stats = append(validator_stats, v_stats)

		}
		if len(validator_stats) > 0 {

			if _, err := DB.Model(&validator_stats).Insert(); err != nil {
				log.Println(err)
			}
		}

		vg_stats := &model.ValidatorGroupStats{
			LockedGold:       vg.LockedGold,
			GroupShare:       vg.Commission,
			RewardRatio:      vg.RewardsRatio,
			Votes:            vg.Votes,
			VotingCap:        vg.ReceivableVotes,
			EpochId:          epoch.ID,
			ValidatorGroupId: vg_db.ID,
			CreatedAt:        time.Now(),
		}

		if _, err := DB.Model(vg_stats).Insert(); err != nil {
			log.Println(err)
		}
		// log.Println("Inserted VG Stats...")

	}
	fmt.Println("Indexing complete...")
}
