package graph

import "github.com/buidl-labs/celo-voting-validator-backend/graph/model"

func calculateTransparencyScore(vg *model.ValidatorGroup) float64 {
	transparencyScore := float64(0)
	if vg.WebsiteURL != "" {
		transparencyScore += 0.15
		if vg.VerifiedDNS {
			transparencyScore += 0.25
		}
	}
	if vg.Name != "" {
		transparencyScore += 0.15
	}
	if vg.Email != "" {
		transparencyScore += 0.15
	}
	if vg.GeographicLocation != "" {
		transparencyScore += 0.1
	}
	if vg.TwitterUsername != "" {
		transparencyScore += 0.1
	}
	if vg.DiscordTag != "" {
		transparencyScore += 0.1
	}

	return transparencyScore
}
