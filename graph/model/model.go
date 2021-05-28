package model

import (
	"fmt"
	"time"
)

type Epoch struct {
	//lint:ignore U1000 `tableName` field is unused, but needed for go-pg
	tableName           struct{}               `pg:"epochs"`
	ID                  string                 `pg:"default:gen_random_uuid()"`
	StartBlock          uint64                 `pg:",unique"`
	EndBlock            uint64                 `pg:",unique"`
	Number              uint64                 `pg:",unique"`
	CreatedAt           time.Time              `pg:"default:now()"`
	ValidatorGroupStats []*ValidatorGroupStats `pg:"rel:has-many"`
	ValidatorStats      []*ValidatorStats      `pg:"rel:has-many"`
}

type Validator struct {
	ID               string    `pg:"default:gen_random_uuid()"`
	Address          string    `pg:",notnull,unique"`
	Name             string    `pg:",unique"`
	CreatedAt        time.Time `pg:"default:now()"`
	CurrentlyElected bool      `pg:",use_zero"`
	ValidatorGroup   *ValidatorGroup
	Stats            []*ValidatorStats `pg:"rel:has-many"`
	ValidatorGroupId string
}

func (v Validator) String() string {
	return fmt.Sprintf("V<%s %s>", v.ID, v.ValidatorGroup)
}

type ValidatorGroup struct {
	ID                   string `pg:"default:gen_random_uuid()"`
	Address              string `pg:",notnull,unique"`
	Name                 string `pg:",unique"`
	Email                string `pg:",unique"`
	WebsiteURL           string `pg:",unique"`
	DiscordTag           string `pg:",unique"`
	TwitterUsername      string `pg:",unique"`
	VerifiedDNS          bool   `pg:",use_zero"`
	GeographicLocation   string
	CreatedAt            time.Time              `pg:"default:now()"`
	EpochRegisteredAt    uint64                 `pg:",use_zero"`
	EpochsServed         uint64                 `pg:"default:0,use_zero"`
	CurrentlyElected     bool                   `pg:"default:false,use_zero"`
	RecievedVotes        uint64                 `pg:",use_zero"`
	AvailableVotes       uint64                 `pg:",use_zero"`
	GroupScore           float64                `pg:",use_zero"`
	LockedCelo           uint64                 `pg:",use_zero"`
	LockedCeloPercentile float64                `pg:",use_zero"`
	SlashingPenaltyScore float64                `pg:",use_zero"`
	AttestationScore     float64                `pg:",use_zero"`
	EstimatedAPY         float64                `pg:",use_zero"`
	TransparencyScore    float64                `pg:",use_zero"`
	PerformanceScore     float64                `pg:",use_zero"`
	Validators           []*Validator           `pg:"rel:has-many"`
	Stats                []*ValidatorGroupStats `pg:"rel:has-many"`
}

func (vg ValidatorGroup) String() string {
	return fmt.Sprintf("VG<%s>", vg.ID)
}

type ValidatorGroupStats struct {
	ID                    string  `pg:"default:gen_random_uuid()"`
	LockedCelo            uint64  `pg:",use_zero"`
	GroupShare            float64 `pg:",use_zero"`
	Votes                 uint64  `pg:",use_zero"`
	VotingCap             uint64  `pg:",use_zero"`
	AttestationPercentage float64 `pg:",use_zero"`
	SlashingScore         float64 `pg:",use_zero"`
	GroupScore            float64 `pg:",use_zero"`
	Epoch                 *Epoch
	EpochId               string
	ValidatorGroup        *ValidatorGroup
	ValidatorGroupId      string
	CreatedAt             time.Time `pg:"default:now()"`
	EstimatedAPY          float64   `pg:"estimated_apy,use_zero"`
}

type ValidatorStats struct {
	ID                    string  `pg:"default:gen_random_uuid()"`
	AttestationsRequested int     `pg:",use_zero"`
	AttestationsFulfilled int     `pg:",use_zero"`
	LastElected           int     `pg:",use_zero"`
	Score                 float64 `pg:",use_zero"`
	Epoch                 *Epoch
	EpochId               string
	Validator             *Validator
	ValidatorId           string
	CreatedAt             time.Time `pg:"default:now()"`
}
