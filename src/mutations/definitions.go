package mutations

import "github.com/flipfloppy1/quDnD/src/statblock"

type MutationCategory string

var AllMutationCategories = []struct {
	Value  MutationCategory
	TSName string
}{
	{Morphotypes, "MORPHOTYPE"},
	{PhysicalMutations, "PHYSICAL_MUTATIONS"},
	{PhysicalDefects, "PHYSICAL_DEFECTS"},
	{MentalMutations, "MENTAL_MUTATIONS"},
	{MentalDefects, "MENTAL_DEFECTS"},
}

const (
	Morphotypes       MutationCategory = "morphotype"
	PhysicalMutations MutationCategory = "physical mutations"
	PhysicalDefects   MutationCategory = "physical defects"
	MentalMutations   MutationCategory = "mental mutations"
	MentalDefects     MutationCategory = "mental defects"
)

type Mutation struct {
	Id                string               `json:"id"`
	Name              string               `json:"name"`
	Category          MutationCategory     `json:"mutationCategory"`
	ImageUrl          string               `json:"imgUrl"`
	Description       string               `json:"description"`
	Abilities         []statblock.Ability  `json:"abilities"`
	Conditions        []string             `json:"conditions"`
	Incompatibilities []string             `json:"incompatibilities"`
	Buffs             []statblock.FeatBuff `json:"buffs"`
	Cost              int                  `json:"cost"`
}

var (
	Mutations map[string]Mutation = map[string]Mutation{
		ChimeraMutation.Id:        ChimeraMutation,
		EsperMutation.Id:          EsperMutation,
		UnstableGenomeMutation.Id: UnstableGenomeMutation,
		AdrenalControlMutation.Id: AdrenalControlMutation,
		BeakMutation.Id:           BeakMutation,
		BurrowingClawsMutation.Id: BurrowingClawsMutation,
		CarapaceMutation.Id:       CarapaceMutation,
	}
	ChimeraMutation = Mutation{"chimera",
		"Chimera",
		Morphotypes,
		"https://wiki.cavesofqud.com/images/a/a9/Chimera_mutation.png",
		"You only manifest physical mutations",
		[]statblock.Ability{statblock.Ability{
			Id:         "chimera",
			Indefinite: true,
			Summary:    "You can only manifest physical mutations",
			Description: `When you expend mutation points to acquire a new mutation,
			all of your choices will be physical mutations. Additionally,
			one choice will allow you to grow a new limb in a random place
			on your body in addition to gaining the mutation. You cannot
			manifest permanent mental mutations through any means.`,
			Conditions: []string{},
			Attacks:    []statblock.Attack{},
			Effects:    []statblock.Effect{},
		}},
		[]string{"can only be chosen at character creation"},
		[]string{"esper", "mental mutations", "mental defects"},
		[]statblock.FeatBuff{},
		1,
	}
	EsperMutation = Mutation{"esper",
		"Esper",
		Morphotypes,
		"https://wiki.cavesofqud.com/images/0/01/Esper_mutation.png",
		"You only manifest mental mutations",
		[]statblock.Ability{statblock.Ability{
			Id:         "esper",
			Indefinite: true,
			Summary:    "You can only manifest mental mutations",
			Description: `When you expend mutation points to acquire a new mutation,
			all of your choices will be mental mutations. You cannot
			manifest permanent physical mutations through any means.`,
			Conditions: []string{},
			Attacks:    []statblock.Attack{},
			Effects:    []statblock.Effect{},
		}},
		[]string{"can only be chosen at character creation"},
		[]string{"chimera", "physical mutations", "physical defects"},
		[]statblock.FeatBuff{},
		1,
	}
	UnstableGenomeMutation = Mutation{"unstable genome",
		"Unstable Genome",
		Morphotypes,
		"https://wiki.cavesofqud.com/images/7/7a/Unstable_genome_mutation.png",
		"You have a chance to manifest a mutation on level up",
		[]statblock.Ability{statblock.Ability{
			Id:         "unstable genome",
			Indefinite: true,
			Summary:    "You have a chance to manifest a mutation on level up",
			Description: `If you have unstable genome when you level up,
			roll a d6. On a 5 or higher, you can choose one of
			3 random mutations to manifest. One stack of unstable genome
			is consumed upon triggering this ability. You may acquire multiple
			stacks of unstable genome on character creation.`,
			Conditions: []string{"unstable genome is active"},
			Attacks:    []statblock.Attack{},
			Effects:    []statblock.Effect{},
		}},
		[]string{"can only be chosen at character creation"},
		[]string{},
		[]statblock.FeatBuff{},
		3,
	}
	AdrenalControlMutation = Mutation{"adrenal control",
		"Adrenal Control",
		PhysicalMutations,
		"https://wiki.cavesofqud.com/images/b/b8/Adrenal_control_mutation.png",
		"You regulate your body's release of adrenaline",
		[]statblock.Ability{statblock.Ability{
			Id:       "adrenal control",
			UseTime:  statblock.FreeAction,
			Duration: &statblock.DiceRoll{Dice: []string{}, Offset: 20, StatBonus: statblock.StatNone},
			Summary:  "You can increase your body's adrenaline flow for 2 minutes",
			Description: `Activate this ability to increase your adrenaline flow
			for 20 rounds. While it's flowing, you gain 9 + MUT quickness and
			the rank of all other physical mutations increase by
			1 + MUT divided by 3 (rounded down). While this ability is
			active, you gain one quickness point.`,
			Conditions: []string{"once per encounter"},
			Attacks:    []statblock.Attack{},
			Effects:    []statblock.Effect{},
		}},
		[]string{"can only be chosen at character creation"},
		[]string{},
		[]statblock.FeatBuff{
			{Stat: statblock.MUT, Value: "1 + Mutation rank", Conditions: []string{"ability is active", "physical mutation", "not adrenaline control"}},
		},
		4,
	}
	BeakMutation = Mutation{"beak",
		"Beak",
		PhysicalMutations,
		"https://wiki.cavesofqud.com/images/5/5f/Beak_mutation.png",
		"Your face bears a sightly beak",
		[]statblock.Ability{statblock.Ability{
			Id:         "beak",
			Indefinite: true,
			Summary:    "Your face bears a sightly beak",
			Description: `This physical mutation has the following
			forms: beak, bill, rostrum, frill and proboscis.
			In addition to the charisma bonus, you find it easier
			to talk to and reason with birds.`,
			Conditions: []string{},
			Attacks:    []statblock.Attack{},
			Effects:    []statblock.Effect{},
		}},
		[]string{},
		[]string{},
		[]statblock.FeatBuff{
			{Stat: statblock.CHA, Value: "1", Conditions: []string{}},
		},
		1,
	}
	BurrowingClawsMutation = Mutation{"burrowing claws",
		"Burrowing Claws",
		PhysicalMutations,
		"https://wiki.cavesofqud.com/images/e/ea/Burrowing_claws_mutation.png",
		"You bear spade-like claws that can burrow through the earth",
		[]statblock.Ability{statblock.Ability{
			Id:         "burrowing claws",
			Indefinite: true,
			Summary:    "Your claws can tunnel through walls and burrow out staircases",
			Description: `You gain a to-hit bonus against walls equal to
			three times this mutation's rank, and 4 successful attacks with
			your claws will destroy a wall. Your claws count as
			natural short blades that deal 1d4 damage from mutation
			rank 1-4, 1d6 damage from MUT 5-9, 1d10 damage
			from MUT 10-12, and 1d12 damage from MUT 13+. While outside
			of combat, you may spend 60 minutes digging a staircase up
			or down a floor.`,
			Conditions: []string{},
			Attacks:    []statblock.Attack{},
			Effects:    []statblock.Effect{},
		}},
		[]string{},
		[]string{},
		[]statblock.FeatBuff{},
		3,
	}
	CarapaceMutation = Mutation{"carapace",
		"Carapace",
		PhysicalMutations,
		"https://wiki.cavesofqud.com/images/d/d5/Carapace_mutation.png",
		"You are protected by a durable carapace",
		[]statblock.Ability{
			statblock.Ability{
				Id:         "carapace",
				Indefinite: true,
				Summary:    "Your carapace provides protection from harm and the elements",
				Description: `You gain an AC bonus equal to this mutation's rank,
				but you cannot wear body armor. You find it much easier to talk to
				and reason with tortoises. Your carapace thermally insulates you,
				such that at MUT 10 and greater you are resistant to heat and cold
				damage.`,
				Conditions: []string{},
				Attacks:    []statblock.Attack{},
				Effects:    []statblock.Effect{},
			},
			statblock.Ability{
				Id:         "carapace tightening",
				UseTime:    statblock.Action,
				Indefinite: true,
				Summary:    "You can tighten your carapace for an extra AC bonus",
				Description: `You may spend an action tightening your carapace,
				which doubles its AC bonus. Moving or otherwise doing anything
				other than dodging or using mental mutations will loosen the
				carapace again.`,
				Conditions: []string{"not moving or making actions"},
				Attacks:    []statblock.Attack{},
				Effects:    []statblock.Effect{},
			},
		},
		[]string{},
		[]string{"quills"},
		[]statblock.FeatBuff{},
		3,
	}
)
