package indexer

import (
	"context"
	"log"

	"github.com/machinebox/graphql"
)

func getLatestBlock(client *graphql.Client) int {

	req := graphql.NewRequest(`{
		latestBlock
	}`)

	ctx := context.Background()
	var resp latestBlockResponse
	if err := client.Run(ctx, req, &resp); err != nil {
		log.Fatal(err)
	}
	return int(resp.LatestBlock)
}

func getValidatorGroups(client *graphql.Client) []celoValidatorGroups {

	req := graphql.NewRequest(`{
		celoValidatorGroups{
			account{
				address
				claims(first:10){
					edges{
						node{
							element
							type
						}
					}
				}
			}
			name
			numMembers
      activeGold
      lockedGold
      commission
      rewardsRatio
      receivableVotes
      votes
      
			affiliates(first: 5){ 
				edges{
					node{
						name
						address
						attestationsRequested
						attestationsFulfilled
						lastElected
						lastOnline
						score
						usd
					}
				}
			}
		}	
	}`)

	ctx := context.Background()
	var resp validatorGroupsResponse
	if err := client.Run(ctx, req, &resp); err != nil {
		log.Fatal(err)
	}

	return resp.CeloValidatorGroups
}

// func getElectedValidators(client *graphql.Client, block_number int) map[string]bool {

// 	req := graphql.NewRequest(`
// 	query($block_num: Int!){
// 		celoElectedValidators(blockNumber: $block_num) {
// 			celoAccount{
// 				address
// 			}
// 		}
// 	}`)

// 	req.Var("block_num", block_number)
// 	ctx := context.Background()
// 	var resp electedValidatorsResponse
// 	if err := client.Run(ctx, req, &resp); err != nil {
// 		log.Fatal(err)
// 	}
// 	electedValidators := make(map[string]bool)
// 	for _, v := range resp.CeloElectedValidators {
// 		electedValidators[v.CeloAccount.Address] = true
// 	}

// 	return electedValidators

// }
