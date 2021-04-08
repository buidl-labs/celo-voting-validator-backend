package indexer

type latestBlockResponse struct {
	LatestBlock float64 `json:"latestBlock"`
}

type validatorGroupsResponse struct {
	CeloValidatorGroups []celoValidatorGroups `json:"celoValidatorGroups"`
}

type celoValidatorGroups struct {
	Account struct {
		Address string `json:"address"`
		Claims  struct {
			Edges []struct {
				Node struct {
					Element string `json:"element"`
					Type    string `json:"type"`
				} `json:"node"`
			} `json:"edges"`
		} `json:"claims"`
	} `json:"account"`
	Affiliates struct {
		Edges []struct {
			Node struct {
				Address               string      `json:"address"`
				AttestationsFulfilled int         `json:"attestationsFulfilled"`
				AttestationsRequested int         `json:"attestationsRequested"`
				LastElected           int         `json:"lastElected"`
				LastOnline            int         `json:"lastOnline"`
				Name                  interface{} `json:"name"`
				Score                 string      `json:"score"`
				Usd                   string      `json:"usd"`
			} `json:"node"`
		} `json:"edges"`
	} `json:"affiliates"`
	Commission      string `json:"commission"`
	LockedGold      string `json:"lockedGold"`
	Name            string `json:"name"`
	NumMembers      int    `json:"numMembers"`
	ReceivableVotes string `json:"receivableVotes"`
	RewardsRatio    string `json:"rewardsRatio"`
	Votes           string `json:"votes"`
}

type electedValidatorsResponse struct {
	CeloElectedValidators []struct {
		CeloAccount struct {
			Address string `json:"address"`
		} `json:"celoAccount"`
	} `json:"celoElectedValidators"`
}
