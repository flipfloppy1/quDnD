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
		ChimeraMutation.Id:              ChimeraMutation,
		EsperMutation.Id:                EsperMutation,
		UnstableGenomeMutation.Id:       UnstableGenomeMutation,
		AdrenalControlMutation.Id:       AdrenalControlMutation,
		BeakMutation.Id:                 BeakMutation,
		BurrowingClawsMutation.Id:       BurrowingClawsMutation,
		CarapaceMutation.Id:             CarapaceMutation,
		CorrosiveGasGeneration.Id:       CorrosiveGasGeneration,
		DoubleMuscledMutation.Id:        DoubleMuscledMutation,
		ElectricalGenerationMutation.Id: ElectricalGenerationMutation,
		ElectromagneticPulseMutation.Id: ElectromagneticPulseMutation,
		FlamingRayMutation.Id:           FlamingRayMutation,
		FreezingRayMutation.Id:          FreezingRayMutation,
		HeightenedHearingMutation.Id:    HeightenedHearingMutation,
		HeightenedQuicknessMutation.Id:  HeightenedQuicknessMutation,
		HornsMutation.Id:                HornsMutation,
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
		[]statblock.FeatBuff{statblock.FeatBuff{
			Stat:       statblock.TOHIT,
			Value:      "3 * MUT",
			Conditions: []string{"targetting a wall"},
		}},
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
	CorrosiveGasGeneration = Mutation{"corrosive gas generation",
		"Corrosive Gas Generation",
		PhysicalMutations,
		"https://wiki.cavesofqud.com/images/b/b4/Corrosive_gas_generation_mutation.png",
		"You release bursts of corrosive gas around yourself",
		[]statblock.Ability{
			statblock.Ability{
				Id:      "generate corrosive gas",
				UseTime: statblock.Action,
				Summary: "You release a cloud of corrosive gas",
				Description: `You are immune to all sources of corrosive gas.
				Every turn this ability is active, a cloud of gas seeps from
				you and fills the adjacent squares. When another creature steps
				through any length of gas on their turn, roll a number of d4s
				equal to the rank of this mutation. The creature must succeed
				a constitution saving throw with DC 10 + MUT or take the d4s
				in acid damage. Once the duration ends, the corrosive gas
				dissipates. You may use this ability once per encounter.`,
				Conditions: []string{},
				Attacks:    []statblock.Attack{},
				Effects:    []statblock.Effect{},
				Duration:   &statblock.DiceRoll{StatBonus: statblock.MUT},
			},
		},
		[]string{},
		[]string{},
		[]statblock.FeatBuff{},
		3,
	}
	DoubleMuscledMutation = Mutation{"double muscled",
		"Double-muscled",
		PhysicalMutations,
		"https://wiki.cavesofqud.com/images/8/84/Doublemuscled_mutation.png",
		"You are possessed of hulking strength",
		[]statblock.Ability{
			statblock.Ability{
				Id:      "double muscled",
				Summary: "Your hulking strength gives you various bonuses",
				Description: `Your strength attribute is increased by 2, plus
				this mutation's rank divided by 2 (rounded down). Whenever you
				make a melee attack, roll a d20. If your result is greater
				than 20 - MUT, you daze your opponent for 1d4 rounds. Similarly
				to bludgeoning, if you daze an already dazed opponent, they are
				stunned. Dazing an opponent that is dazed by bludgeoning in the
				same attack will instantly stun them.`,
				Conditions: []string{},
				Attacks:    []statblock.Attack{},
				Effects: []statblock.Effect{statblock.Effect{
					Effect:     statblock.DAZED,
					Conditions: []string{"greater than 20 - MUT on a d20"},
					Rounds:     statblock.DiceRoll{Dice: []string{"1d4"}},
					Reasons:    []string{},
				}},
				Indefinite: true,
			},
		},
		[]string{},
		[]string{},
		[]statblock.FeatBuff{statblock.FeatBuff{Stat: statblock.STR, Value: "2 + Mutation rank / 2"}},
		3,
	}
	ElectricalGenerationMutation = Mutation{"electrical generation",
		"Electrical Generation",
		PhysicalMutations,
		"https://wiki.cavesofqud.com/images/0/0e/Electrical_generation_mutation.png",
		"You accrue electrical charge that can be used in combat or otherwise",
		[]statblock.Ability{
			statblock.Ability{
				Id:      "electrical generation",
				Summary: "You can accrue electrical energy",
				Description: `You accrue 1 charge of electricity every minute,
				up to a maximum of twice this mutation's rank, plus 2. Whenever
				you take electrical damage, you gain charges equal to the
				amount of damage you took divided by 10, minimum 1. You can
				drink charges from energy cells and capacitors, and provide
				charge to equipped devices with integrated power systems.`,
				Conditions: []string{},
				Attacks:    []statblock.Attack{},
				Effects:    []statblock.Effect{},
				Indefinite: true,
			},
			statblock.Ability{
				Id:      "electrical discharge",
				Summary: "You can discharge electricity you've accrued to damage enemies",
				Description: `You expend all your accumulated charges, dealing 1d8
				of thunder damage for each charge to a target within 5ft. You may
				extend the electrical discharge to target an additional creature
				or object within 5ft of the previous arc, reducing the damage for
				both targets by 1 charge. You can continue to do this until each
				target is dealt 1d8 damage.`,
				Conditions: []string{"at least one charge accrued"},
				Attacks:    []statblock.Attack{},
				Effects:    []statblock.Effect{},
				Indefinite: true,
			},
		},
		[]string{},
		[]string{},
		[]statblock.FeatBuff{},
		4,
	}
	ElectromagneticPulseMutation = Mutation{"electromagnetic pulse",
		"Electromagnetic Pulse",
		PhysicalMutations,
		"https://wiki.cavesofqud.com/images/2/2d/Electromagnetic_pulse_mutation.png",
		"You generate an electromagnetic pulse that disables nearby artifacts and machines",
		[]statblock.Ability{
			statblock.Ability{
				Id:      "electromagnetic pulse",
				UseTime: statblock.BonusAction,
				Summary: "You can emit an electromagnetic pulse",
				Description: `Once per encounter, you may emit an
				electromagnetic pulse, disabling artifacts, machines and robots
				in a radius around you for MUT rounds. If this mutation's rank
				is less than 5, the radius is 2. If this mutation's rank is
				between 5 and 9, the radius is 5. For MUT 10+, the radius is 9.`,
				Conditions: []string{"once per encounter"},
				Attacks:    []statblock.Attack{},
				Effects:    []statblock.Effect{},
				Duration:   &statblock.DiceRoll{StatBonus: statblock.MUT},
			},
		},
		[]string{},
		[]string{},
		[]statblock.FeatBuff{},
		2,
	}
	FlamingRayMutation = Mutation{"flaming ray",
		"Flaming Ray",
		PhysicalMutations,
		"https://wiki.cavesofqud.com/images/1/10/Flaming_ray_mutation.png",
		"You emit a ray of flame from your hands, feet, or face",
		[]statblock.Ability{
			statblock.Ability{
				Id:      "flaming ray",
				UseTime: statblock.Action,
				Summary: "You can emit a ray of flames to incinerate targets",
				Description: `When you gain this mutation, you may choose whether
				it is emitted from your hands, feet or face. You emit a 9-square
				ray of flame in the direction of your choice, dealing MUTd4 fire
				damage and raising the temperature of the target. The flaming ray
				may also evaporate liquids. With this mutation, your melee attacks
				raise the temperature of your target.
				`,
				Conditions: []string{},
				Attacks:    []statblock.Attack{},
				Effects:    []statblock.Effect{},
			},
		},
		[]string{},
		[]string{"freezing ray"},
		[]statblock.FeatBuff{},
		4,
	}
	FreezingRayMutation = Mutation{"freezing ray",
		"Freezing Ray",
		PhysicalMutations,
		"https://wiki.cavesofqud.com/images/a/af/Freezing_ray_mutation.png",
		"You emit a ray of frost from your hands, feet, or face",
		[]statblock.Ability{
			statblock.Ability{
				Id:      "freezing ray",
				UseTime: statblock.Action,
				Summary: "You can emit a ray of frost to freeze targets",
				Description: `When you gain this mutation, you may choose whether
				it is emitted from your hands, feet or face. You emit a 9-square
				ray of frost in the direction of your choice, dealing MUTd4 cold
				damage and dropping the temperature of the target. The freezing ray
				may also freeze liquids. With this mutation, your melee attacks
				drop the temperature of your target. You can use this mutation
				once every two rounds. If you have double the body parts
				usually used by this mutation, you may double the damage of the
				freezing ray. This mutation makes you more resistant to being frozen.`,
				Conditions: []string{"freezing ray was not used last round"},
				Attacks:    []statblock.Attack{},
				Effects:    []statblock.Effect{},
			},
		},
		[]string{},
		[]string{"flaming ray"},
		[]statblock.FeatBuff{},
		5,
	}
	HeightenedHearingMutation = Mutation{"heightened hearing",
		"Heightened Hearing",
		PhysicalMutations,
		"https://wiki.cavesofqud.com/images/d/d0/Heightened_hearing_mutation.png",
		"You are possessed of unnaturally acute hearing",
		[]statblock.Ability{
			statblock.Ability{
				Id:      "heightened hearing",
				UseTime: statblock.Action,
				Summary: "Your acute hearing allows you to detect creatures",
				Description: `You gain advantage to perceiving the presence of
				creatures, both passively and actively. Whenever you successfully
				perceive a creature, you gain knowledge of its precise location.
				You may spend an action listening to a perceived creature,
				rolling a DC 16 perception check and adding this mutation's
				rank. If you succeed, you identify the type of creature.`,
				Conditions: []string{},
				Attacks:    []statblock.Attack{},
				Effects:    []statblock.Effect{},
			},
		},
		[]string{},
		[]string{},
		[]statblock.FeatBuff{statblock.FeatBuff{
			Stat:       statblock.WIS,
			Value:      "advantage",
			Conditions: []string{"on passive or active perception towards the presence of creatures"},
		}},
		2,
	}
	HeightenedQuicknessMutation = Mutation{"heightened quickness",
		"Heightened Quickness",
		PhysicalMutations,
		"https://wiki.cavesofqud.com/images/7/7e/Heightened_quickness_mutation.png",
		"You are gifted with tremendous speed",
		[]statblock.Ability{
			statblock.Ability{
				Id:      "heightened hearing",
				Summary: "Your speed allows you to take more actions than usual sometimes",
				Description: `When you take the first action of your turn, roll a
				percentile die. If you roll within 13 + MUT * 2 from 100, you may
				take another action.`,
				Conditions: []string{},
				Attacks:    []statblock.Attack{},
				Effects:    []statblock.Effect{},
			},
		},
		[]string{},
		[]string{},
		[]statblock.FeatBuff{},
		3,
	}
	HornsMutation = Mutation{"horns",
		"Horns",
		PhysicalMutations,
		"https://wiki.cavesofqud.com/images/0/0d/Horns_mutation.png",
		"Horns jut out of your head",
		[]statblock.Ability{
			statblock.Ability{
				Id:      "horns",
				Summary: "Your horns provide versatility in combat",
				Description: `When you acquire this mutation, you may choose
				from the following variants: horns, a horn, a casque, antlers,
				or a spiral horn. If you have multiple heads, the horns will
				only be on one head. Whenever you make a melee attack, roll a
				d4. On a 4, you make a short blade attack using your horns, with
				a to-hit of 1 + MUT / 2, causing 2d4 damage on hit. This damage
				increases to 2d6 at mutation rank 8, 2d8 at mutation rank
				16, and 2d12 at mutation rank 24. Whenever you make an
				attack with your horns, your target must succeed a DC 14 + MUT
				constitution saving throw or begin bleeding.`,
				Conditions: []string{},
				Attacks:    []statblock.Attack{},
				Effects: []statblock.Effect{statblock.Effect{
					Effect:     statblock.BLEEDING,
					Conditions: []string{"on enemy failing CON save"},
					Reasons:    []string{"enemy is hit by horns"},
					Rounds:     statblock.DiceRoll{StatBonus: statblock.MUT},
				}},
			},
		},
		[]string{},
		[]string{},
		[]statblock.FeatBuff{statblock.FeatBuff{
			Stat:  statblock.AC,
			Value: "1 + MUT / 3",
		}},
		4,
	}
)
