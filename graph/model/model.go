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
	Name             string
	Address          string
	Memory           int
	CPUCapacity      int
	DiskStorage      int
	NetworkSpeed     int
	ValidatorGroup   *ValidatorGroup
	ValidatorGroupId string
	Stats            []*ValidatorStats `pg:"rel:has-many"`
	CreatedAt        time.Time
}

func (v Validator) String() string {
	return fmt.Sprintf("V<%s %s>", v.ID, v.ValidatorGroup)
}

type ValidatorGroup struct {
	ID                 string `pg:"default:gen_random_uuid()"`
	Name               string
	Email              string
	GeographicLocation string
	WebsiteURL         string
	DiscordTag         string
	TwitterUsername    string
	Address            string
	VerifiedDNS        bool
	EpochsServed       int
	CreatedAt          time.Time
	NumValidators      int
	Validators         []*Validator           `pg:"rel:has-many"`
	Stats              []*ValidatorGroupStats `pg:"rel:has-many"`
}

func (vg ValidatorGroup) String() string {
	return fmt.Sprintf("VG<%s>", vg.ID)
}

type ValidatorGroupStats struct {
	ID                    string `pg:"default:gen_random_uuid()"`
	LockedGold            string
	GroupShare            string
	Votes                 string
	VotingCap             string
	RewardRatio           string
	AttestationPercentage float64
	Epoch                 *Epoch
	EpochId               string
	ValidatorGroup        *ValidatorGroup
	ValidatorGroupId      string
	CreatedAt             time.Time
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
