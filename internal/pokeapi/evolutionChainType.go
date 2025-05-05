package pokeapi

import (
	"bytes"
	"fmt"
	"strings"
)

type Species struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type evoChainRes struct {
	Chain *EvolutionChainLink `json:"chain"`
}

type EvolutionChainLink struct {
	EvolutionDetails []EvolutionDetails    `json:"evolution_details"`
	EvolvesTo        []*EvolutionChainLink `json:"evolves_to"`
	IsBaby           bool                  `json:"is_baby"`
	Species          Species               `json:"species"`
}

type EvolutionDetails struct {
	Trigger               NamedApiResource  `json:"trigger"`
	Gender                *int              `json:"gender"`
	HeldItem              *NamedApiResource `json:"held_item"`
	Item                  *NamedApiResource `json:"item"`
	KnownMove             *NamedApiResource `json:"known_move"`
	KnownMoveType         *NamedApiResource `json:"known_move_type"`
	Location              *NamedApiResource `json:"location"`
	MinAffection          *int              `json:"min_affection"`
	MinBeauty             *int              `json:"min_beauty"`
	MinHappiness          *int              `json:"min_happiness"`
	MinLevel              *int              `json:"min_level"`
	NeedsOverworldRain    bool              `json:"needs_overworld_rain"`
	PartySpecies          *NamedApiResource `json:"party_species"`
	PartyType             *NamedApiResource `json:"party_type"`
	RelativePhysicalStats *int              `json:"relative_physical_stats"`
	TimeOfDay             string            `json:"time_of_day"`
	TradeSpecies          *NamedApiResource `json:"trade_species"`
	TurnUpsideDown        bool              `json:"turn_upside_down"`
}

func (e *EvolutionDetails) String() string {
	var out bytes.Buffer
	if e.Trigger.Name != "" {
		fmt.Fprintf(&out, "%s ", cleanName(e.Trigger.Name))
	}
	if e.MinLevel != nil {
		fmt.Fprintf(&out, "At level: %d + ", *e.MinLevel)
	}
	if e.MinHappiness != nil {
		fmt.Fprintf(&out, "Min happiness: %d + ", *e.MinHappiness)
	}
	if e.MinBeauty != nil {
		fmt.Fprintf(&out, "Min beauty: %d + ", *e.MinBeauty)
	}
	if e.MinAffection != nil {
		fmt.Fprintf(&out, "Min affection: %d + ", *e.MinAffection)
	}
	if e.Gender != nil {
		fmt.Fprintf(&out, "Gender: %s + ", getGender(*e.Gender))
	}
	if e.HeldItem != nil {
		fmt.Fprintf(&out, "Held item: %s + ", cleanName(e.HeldItem.Name))
	}
	if e.Item != nil {
		fmt.Fprintf(&out, "Use item: %s + ", cleanName(e.Item.Name))
	}
	if e.KnownMove != nil {
		fmt.Fprintf(&out, "Knowing move: %s + ", cleanName(e.KnownMove.Name))
	}
	if e.KnownMoveType != nil {
		fmt.Fprintf(&out, "Knowing a move of type: %s + ", cleanName(e.KnownMoveType.Name))
	}
	if e.Location != nil {
		fmt.Fprintf(&out, "At location: %s + ", cleanName(e.Location.Name))
	}
	if e.NeedsOverworldRain != false {
		fmt.Fprintf(&out, "While raining: %t + ", e.NeedsOverworldRain)
	}
	if e.PartySpecies != nil {
		fmt.Fprintf(&out, "While pokemon in party: %s + ", cleanName(e.PartySpecies.Name))
	}
	if e.PartyType != nil {
		fmt.Fprintf(&out, "While pokemon in party of type: %s + ", cleanName(e.PartyType.Name))
	}
	if e.RelativePhysicalStats != nil {
		fmt.Fprintf(&out, "When stats: %s + ", getRelativePhysStats(*e.RelativePhysicalStats))
	}
	if e.TimeOfDay != "" {
		fmt.Fprintf(&out, "When time of day: %s + ", e.TimeOfDay)
	}
	if e.TradeSpecies != nil {
		fmt.Fprintf(&out, "When trading for: %s + ", cleanName(e.TradeSpecies.Name))
	}
	if e.TradeSpecies != nil {
		fmt.Fprintf(&out, "When trading for: %s + ", cleanName(e.TradeSpecies.Name))
	}
	if e.TurnUpsideDown != false {
		fmt.Fprintf(&out, "When turning upside down + ")
	}

	res := out.String()


	if len(res) > 2 {
		plus := res[len(res)-2:]
		if plus == "+ " {
			res = res[:len(res)-2] // remove last '+'
		}  
	}

	return res
}

func getGender(gendID int) string {
	if gendID == 1 {
		return "female"
	}
	return "male"
}

func getRelativePhysStats(relID int) string {
	if relID == 1 {
		return "ATK > DEF"
	}
	if relID == 0 {
		return "ATK = DEF"
	}
	return "ATK < DEF"
}

func cleanName(name string) string {
	return strings.ReplaceAll(name, "-", " ")
}
