package model

import (
	"fmt"
	"time"
)

type Epoch struct {
	tableName           struct{} `pg:"epochs"`
	ID                  string   `pg:"default:gen_random_uuid()"`
	StartBlock          int
	EndBlock            int
	Number              int
	CreatedAt           time.Time
	ValidatorGroupStats []*ValidatorGroupStats `pg:"rel:has-many"`
	ValidatorStats      []*ValidatorStats      `pg:"rel:has-many"`
}

type Validator struct {
	ID               string `pg:"default:gen_random_uuid()"`
	Address          string
	Name             string
	CreatedAt        time.Time
	CurrentlyElected bool
	ValidatorGroup   *ValidatorGroup
	Stats            []*ValidatorStats `pg:"rel:has-many"`
	ValidatorGroupId string
}

func (v Validator) String() string {
	return fmt.Sprintf("V<%s %s>", v.ID, v.ValidatorGroup)
}

type ValidatorGroup struct {
	ID                   string `pg:"default:gen_random_uuid()"`
	Address              string
	Name                 string
	Email                string
	WebsiteURL           string
	DiscordTag           string
	TwitterUsername      string
	VerifiedDNS          bool
	GeographicLocation   string
	CreatedAt            time.Time
	EpochRegisteredAt    int
	EpochsServed         int
	RecievedVotes        int
	AvailableVotes       int
	GroupScore           int
	LockedCelo           int
	LockedCeloPercentile float64
	SlashingPenaltyScore float64
	AttestationScore     float64
	EstimatedAPY         float64
	TransparencyScore    float64
	PerformanceScore     float64
	Validators           []*Validator           `pg:"rel:has-many"`
	Stats                []*ValidatorGroupStats `pg:"rel:has-many"`
}

func (vg ValidatorGroup) String() string {
	return fmt.Sprintf("VG<%s>", vg.ID)
}

type ValidatorGroupStats struct {
	ID                    string `pg:"default:gen_random_uuid()"`
	LockedCelo            string
	LockedCeloPercentile  float64
	GroupShare            string
	Votes                 string
	VotingCap             string
	AttestationPercentage float64
	SlashingScore         float64
	Epoch                 *Epoch
	EpochId               string
	ValidatorGroup        *ValidatorGroup
	ValidatorGroupId      string
	CreatedAt             time.Time
	EstimatedAPY          float64
}

type ValidatorStats struct {
	ID                     string `pg:"default:gen_random_uuid()"`
	AttestationsRequested  int
	AttenstationsFulfilled int
	LastElected            int
	Score                  string
	Epoch                  *Epoch
	EpochId                string
	Validator              *Validator
	ValidatorId            string
	CreatedAt              time.Time
}
