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
		MetamorphosisMutation.Id:        MetamorphosisMutation,
		MultipleArmsMutation.Id:         MultipleArmsMutation,
		MultipleLegsMutation.Id:         MultipleLegsMutation,
		NightVisionMutation.Id:          NightVisionMutation,
		PhasingMutation.Id:              PhasingMutation,
		PhotosyntheticSkinMutation.Id:   PhotosyntheticSkinMutation,
	}
	ChimeraMutation = Mutation{"chimera",
		"Chimera",
		Morphotypes,
		"https://wiki.cavesofqud.com/images/a/a9/Chimera_mutation.png",
		"You only manifest physical mutations",
		[]statblock.Ability{
			{
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
			},
		},
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
		[]statblock.Ability{
			{
				Id:         "esper",
				Indefinite: true,
				Summary:    "You can only manifest mental mutations",
				Description: `When you expend mutation points to acquire a new mutation,
			all of your choices will be mental mutations. You cannot
			manifest permanent physical mutations through any means.`,
				Conditions: []string{},
				Attacks:    []statblock.Attack{},
				Effects:    []statblock.Effect{},
			},
		},
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
		[]statblock.Ability{
			{
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
			},
		},
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
		[]statblock.Ability{
			{
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
				WillPoints: 3,
			},
		},
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
		[]statblock.Ability{
			{
				Id:         "beak",
				Indefinite: true,
				Summary:    "Your face bears a sightly beak",
				Description: `This physical mutation has the following
			forms: beak, bill, rostrum, frill and proboscis.
			In addition to the charisma bonus, when you gain
			this mutation the party's reputation with birds
			increases by 8.`,
				Conditions: []string{},
				Attacks:    []statblock.Attack{},
				Effects:    []statblock.Effect{},
			},
		},
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
		[]statblock.Ability{
			{
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
			},
		},
		[]string{},
		[]string{},
		[]statblock.FeatBuff{
			{
				Stat:       statblock.TOHIT,
				Value:      "3 * MUT",
				Conditions: []string{"targetting a wall"},
			},
		},
		3,
	}
	CarapaceMutation = Mutation{
		Id:          "carapace",
		Name:        "Carapace",
		Category:    PhysicalMutations,
		ImageUrl:    "https://wiki.cavesofqud.com/images/d/d5/Carapace_mutation.png",
		Description: "You are protected by a durable carapace",
		Abilities: []statblock.Ability{
			{
				Id:         "carapace",
				Indefinite: true,
				Summary:    "Your carapace provides protection from harm and the elements",
				Description: `You gain an AC bonus equal to this mutation's rank,
				but you cannot wear body armor. When you acquire this mutation, the party's
				reputation with tortoises increases by 16. Your carapace thermally insulates you,
				such that at MUT 10 and greater you are resistant to heat and cold
				damage.`,
				Conditions: []string{},
				Attacks:    []statblock.Attack{},
				Effects:    []statblock.Effect{},
			},
			{
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
		Conditions:        []string{},
		Incompatibilities: []string{"quills"},
		Buffs:             []statblock.FeatBuff{},
		Cost:              3,
	}
	CorrosiveGasGeneration = Mutation{"corrosive gas generation",
		"Corrosive Gas Generation",
		PhysicalMutations,
		"https://wiki.cavesofqud.com/images/b/b4/Corrosive_gas_generation_mutation.png",
		"You release bursts of corrosive gas around yourself",
		[]statblock.Ability{
			{
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
				Duration:   &statblock.DiceRoll{Dice: []string{}, StatBonus: statblock.MUT},
				WillPoints: 2,
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
			{
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
				Effects: []statblock.Effect{
					{
						Effect:     statblock.DAZED,
						Conditions: []string{"greater than 20 - MUT on a d20"},
						Rounds:     statblock.DiceRoll{Dice: []string{"1d4"}},
						Reasons:    []string{},
					},
				},
				Indefinite: true,
			},
		},
		[]string{},
		[]string{},
		[]statblock.FeatBuff{{Stat: statblock.STR, Value: "2 + Mutation rank / 2"}},
		3,
	}
	ElectricalGenerationMutation = Mutation{"electrical generation",
		"Electrical Generation",
		PhysicalMutations,
		"https://wiki.cavesofqud.com/images/0/0e/Electrical_generation_mutation.png",
		"You accrue electrical charge that can be used in combat or otherwise",
		[]statblock.Ability{
			{
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
			{
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
			{
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
				Duration:   &statblock.DiceRoll{Dice: []string{}, StatBonus: statblock.MUT},
				WillPoints: 3,
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
			{
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
				WillPoints: 1,
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
			{
				Id:      "freezing ray",
				UseTime: statblock.Action,
				Summary: "You can emit a ray of frost to freeze targets",
				Description: `When you gain this mutation, you may choose whether
				it is emitted from your hands, feet or face. You emit a 9-square
				ray of frost in the direction of your choice, dealing MUTd4 cold
				damage and dropping the temperature of the target. The freezing ray
				may also freeze liquids. With this mutation, your melee attacks
				drop the temperature of your target. If you have double the body parts
				usually used by this mutation, you may double the damage of the
				freezing ray. This mutation makes you more resistant to being frozen.`,
				Conditions: []string{"freezing ray was not used last round"},
				Attacks:    []statblock.Attack{},
				Effects:    []statblock.Effect{},
				WillPoints: 1,
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
			{
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
		[]statblock.FeatBuff{
			{
				Stat:       statblock.WIS,
				Value:      "advantage",
				Conditions: []string{"on passive or active perception towards the presence of creatures"},
			},
		},
		2,
	}
	HeightenedQuicknessMutation = Mutation{"heightened quickness",
		"Heightened Quickness",
		PhysicalMutations,
		"https://wiki.cavesofqud.com/images/7/7e/Heightened_quickness_mutation.png",
		"You are gifted with tremendous speed",
		[]statblock.Ability{
			{
				Id:      "heightened hearing",
				Summary: "Your speed allows you to take more actions than usual sometimes",
				Description: `When you take the first action of your turn, roll a
				percentile die. If you roll within 13 + MUT * 2 of 100, you may
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
			{
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
				constitution saving throw or begin bleeding. When you gain
				this mutation, the party's reputation with antelopes and
				goatfolk increases by 4.`,
				Conditions: []string{},
				Attacks:    []statblock.Attack{},
				Effects: []statblock.Effect{
					{
						Effect:     statblock.BLEEDING,
						Conditions: []string{"on enemy failing CON save"},
						Reasons:    []string{"enemy is hit by horns"},
						Rounds:     statblock.DiceRoll{Dice: []string{}, StatBonus: statblock.MUT},
					},
				},
			},
		},
		[]string{},
		[]string{},
		[]statblock.FeatBuff{
			{
				Stat:  statblock.AC,
				Value: "1 + MUT / 3",
			},
		},
		4,
	}
	MetamorphosisMutation = Mutation{"metamorphosis",
		"Metamorphosis",
		PhysicalMutations,
		"https://wiki.cavesofqud.com/images/6/6d/Metamorphosis_mutation.png",
		"You assume the form of any creature you touch",
		[]statblock.Ability{
			{
				Id:      "metamorphosis",
				Summary: "You may inherit the properties of creatures adjacent to you",
				Description: `You may touch a creature, assuming their form.
				You become a near-perfect copy of the creature, inheriting their level,
				attributes, anatomy, mutations, skills, etc. The only stats you retain from
				your original body are your intelligence, wisdom and charisma. You lose
				access to your own skills and mutations, but retain your inventory.
				You may keep the ability active indefinitely, but once you end it,
				the ability will be on cooldown until the end of your next long rest.
				When you end Metamorphosis, you return to your original body.`,
				Conditions: []string{"inaccessible from character creation and random mutations"},
				Attacks:    []statblock.Attack{},
				Effects:    []statblock.Effect{},
			},
		},
		[]string{},
		[]string{},
		[]statblock.FeatBuff{},
		4,
	}
	MultipleArmsMutation = Mutation{"multiple arms",
		"Multiple Arms",
		PhysicalMutations,
		"https://wiki.cavesofqud.com/images/a/a0/Multiple_arms_mutation.png",
		"You have an extra set of arms",
		[]statblock.Ability{
			{
				Id:      "multiple arms",
				Summary: "Your extra arms can be used in melee",
				Description: `You gain two new hands, two new arms, and one new 'worn on hands' slot when
				you acquire this mutation. Whenever you make a melee attack, roll a d20 for each
				arm you get from this mutation. If you roll greater than 20 - MUT, you make an
				attack with that arm. This chance stacks with multiweapon fighting (roll both odds per
				arm, if at least one succeeds make an extra attack with that arm).`,
				Conditions: []string{},
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
	MultipleLegsMutation = Mutation{"multiple legs",
		"Multiple Legs",
		PhysicalMutations,
		"https://wiki.cavesofqud.com/images/1/17/Multiple_legs_mutation.png",
		"You have an extra set of legs",
		[]statblock.Ability{
			{
				Id:      "multiple legs",
				Summary: "Your extra legs augment your speed and carry capacity",
				Description: `You gain 5ft of movement speed and 10lbs carry weight for
				each rank you have in Multiple Legs.`,
				Conditions: []string{},
				Attacks:    []statblock.Attack{},
				Effects:    []statblock.Effect{},
				Indefinite: true,
			},
		},
		[]string{},
		[]string{},
		[]statblock.FeatBuff{{Stat: statblock.Speed, Value: "MUT * 5ft"}},
		5,
	}
	NightVisionMutation = Mutation{"night vision",
		"Night Vision",
		PhysicalMutations,
		"https://wiki.cavesofqud.com/images/4/49/Night_vision_mutation.png",
		"You can see in the dark",
		[]statblock.Ability{
			{
				Id:      "night vision",
				Summary: "Your night vision allows you to see in the dark",
				Description: `You can see up to 15ft in darkness as if it were bright light,
				and from 15 to 25ft as if it were dim light.`,
				Conditions: []string{},
				Attacks:    []statblock.Attack{},
				Effects:    []statblock.Effect{},
				Indefinite: true,
			},
		},
		[]string{},
		[]string{},
		[]statblock.FeatBuff{},
		1,
	}
	PhasingMutation = Mutation{"phasing",
		"Phasing",
		PhysicalMutations,
		"https://wiki.cavesofqud.com/images/4/47/Phasing_mutation.png",
		"You can phase through solid objects for brief periods of time",
		[]statblock.Ability{
			{
				Id:      "phasing",
				Summary: "You can phase through solid objects for brief periods of time",
				Description: `When you activate this ability, you shift out of phase, becoming
				physically intangible to any creature or object that is not out-of-phase itself.
				Most creatures cannot harm you in this state, and you cannot harm them. This
				effect lasts 6 + MUT rounds, and can be used once per long rest from mutation
				ranks 1-4, once per short rest from mutation ranks 5-9, and for 3 will points
				at mutation rank 10+. If you are forced to phase back in by reality-stabilizing
				effects, you take 2d6 astral damage.`,
				Conditions: []string{},
				Attacks:    []statblock.Attack{},
				Effects:    []statblock.Effect{},
				Duration:   &statblock.DiceRoll{StatBonus: statblock.MUT, Offset: 6, Dice: []string{}},
			},
		},
		[]string{},
		[]string{},
		[]statblock.FeatBuff{},
		4,
	}
	PhotosyntheticSkinMutation = Mutation{"photosynthetic skin",
		"Photosynthetic Skin",
		PhysicalMutations,
		"https://wiki.cavesofqud.com/images/b/bb/Photosynthetic_skin_mutation.png",
		"You replenish yourself by absorbing sunlight through your hearty green skin",
		[]statblock.Ability{
			{
				Id:      "photosynthetic sustenance",
				Summary: "You can bask in sunlight instead of eating meals",
				Description: `When above ground during the day, you can bask in sunlight,
				increasing your movement speed by MUT * 5ft and giving you an AC bonus
				of 1 + MUT / 4. This effect lasts for 1 + MUT / 4 days.`,
				Conditions: []string{},
				Attacks:    []statblock.Attack{},
				Effects:    []statblock.Effect{},
			},
			{
				Id:      "photosynthetic skin",
				Summary: "Your photosynthetic skin gives you a number of bonuses",
				Description: `While in sunlight, you accrue a maximum of 1 + MUT / 4 starch and
				lignin that you can use as ingredients in meals that you cook. You gain
				a 1 + MUT / 4 AC bonus while in dense foliage, and upon gaining this
				mutation your reputation with roots, trees, vines and the Consortium of Phyta
				increases by 8.`,
				Conditions: []string{},
				Attacks:    []statblock.Attack{},
				Effects:    []statblock.Effect{},
			},
		},
		[]string{},
		[]string{},
		[]statblock.FeatBuff{},
		2,
	}
	QuillsMutation = Mutation{
		Id:          "quills",
		Name:        "Quills",
		Category:    PhysicalMutations,
		ImageUrl:    "https://wiki.cavesofqud.com/images/a/a5/Quills_mutation.png",
		Description: "Hundreds of needle-pointed quills cover your body",
		Abilities: []statblock.Ability{
			{
				Id:      "quill burst",
				UseTime: statblock.Action,
				Summary: "You can expel quills in a burst around yourself.",
				Description: `When you use this ability, you expend 10% of your maximum quills, attacking all creatures and objects within 5ft of you.
				The to-hit bonus for this attack starts at 0 and increases by one for every odd level you gain in this mutation, to a maximum of 6.
			 	Each successful attack deals 1d4+MUT piercing damage to the affected item or creature.
				Once you have used your quills, you must wait for them to regenerate. You regenerate half your maximum quills every hour.`,
				Attacks: []statblock.Attack{
					{
						DamageType: statblock.Piercing,
						Damage:     statblock.DiceRoll{StatBonus: statblock.MUT, Dice: []string{"1d4"}},
					},
				},
				WillPoints: 1,
			},
			{
				Id:         "quill defense",
				Indefinite: true,
				Summary:    "Your quills shield you from attacks",
				Description: `You gain a +2 bonus to your AV that increases by one for every odd level you gain in this mutation.
				This bonus is halved (rounded up) whenever you have 50% or less of your maximum quills.`,
			},
			{
				Id:         "impaling quills",
				Indefinite: true,
				Summary:    "Your quills impale attackers",
				Description: `When you suffer damage in melee combat and have more than 20% of your maximum quills left, roll a d4.
				On a 4, you deal 1d8 piercing damage to the attacker up to a maximum of the damage you received.
				This damage increases to 2d8 at mutation level 10.`,
				Conditions: []string{"attacked in melee", "more than 20% of quills left"},
			},
		},
		Conditions:        []string{"cannot wear body armour"},
		Incompatibilities: []string{"carapace"},
		Cost:              3,
	}
	RegenerationMutation = Mutation{
		Id:          "regeneration",
		Name:        "Regeneration",
		Category:    PhysicalMutations,
		ImageUrl:    "https://wiki.cavesofqud.com/images/c/c3/Regeneration_mutation.png",
		Description: "Your wounds heal very quickly",
		Abilities: []statblock.Ability{
			{
				Id:         "quick recovery",
				Indefinite: true,
				Summary:    "Your maximum hit dice is doubled.",
			},
			{
				Id:         "limb regeneration",
				Indefinite: true,
				Summary:    "Your supernatural regeneration allows you to regrow missing limbs",
				Description: `Each round you have a missing limb, roll 1d10.
				If you roll equal to or less than MUT, you regrow the missing limb.
				If this mutation's rank exceeds 10, you regrow limbs in one round and are immune to decapitation.`,
			},
			{
				Id:         "debuff resistance",
				Indefinite: true,
				Summary:    "Your regeneration grants you resistance to physical debuffs",
				Description: `Each round you have a minor physical ailment, roll a 1d100.
				On a 1 + MUT / 3 (rounded down) or higher, the ailment is removed.
				You may roll for multiple rounds by multiplying 1 + MUT / 3 (rounded down) by the number of rounds you have waited for.
				If this value exceeds 100, you automatically recover from the ailment.
				At mutation level 5, this ability is no longer limited to minor ailments.
				
				Minor ailments include: bleeding, blind, choking on ash, cleaved, covered in spores, cracked, ill, itchy skin, paralyzed,
				phase spider venom, poisoned by gas, poisoned, rusted, sluggish, stony, stunned by gas
				
				Major ailments include: broken, glotrot, sore throat, ironshank, stiff legs, blurry vision, monochrome, mutating`,
			},
		},
		Cost: 4,
	}
	SleepGasGenerationMutation = Mutation{
		Id:          "sleep gas generation",
		Name:        "Sleep Gas Generation",
		Category:    PhysicalMutations,
		ImageUrl:    "https://wiki.cavesofqud.com/images/f/f2/Sleep_gas_generation_mutation.png",
		Description: "You release bursts of sleep gas around yourself",
		Abilities: []statblock.Ability{
			{
				Id:      "sleep gas generation",
				Summary: "You release sleep gas for a number of turns",
				Description: `For 2 + MUT rounds after activating this ability, you release sleep gas.
				While you are releasing sleep gas, any creature that ends its turn within 5ft of you must succeed on a CON save or fall asleep for 1d4 + MUT rounds. The DC of this save is 10 + MUT.
				In enclosed areas or areas with low ventilation, the save DC increases by 2 for each round that you release sleep gas, and creatures in a 5ft radius of any place you have released sleep gas can suffer its effects.`,
				WillPoints: 2,
				Duration:   &statblock.DiceRoll{Offset: 2, StatBonus: statblock.MUT, Dice: []string{}},
				UseTime:    statblock.Action,
				Effects: []statblock.Effect{
					{
						Effect:     statblock.ASLEEP,
						Conditions: []string{"is not immune to sleep", "does not have the 'Wakeful' effect"},
						Reasons:    []string{"affected by sleeping gas"},
						Rounds:     statblock.DiceRoll{Dice: []string{"1d4"}, StatBonus: statblock.MUT},
					},
				},
			},
		},
		Cost: 3,
	}
	SlimeGlandsMutation = Mutation{
		Id:          "slime glands",
		Name:        "Slime Glands",
		Category:    PhysicalMutations,
		ImageUrl:    "https://wiki.cavesofqud.com/images/c/ca/Slime_glands_mutation.png",
		Description: "You produce a viscous slime that you can spit at things",
		Abilities: []statblock.Ability{
			{
				Id:      "spit slime",
				Summary: "You spit slime into an area",
				Description: `You spit slime into a 15 by 15 foot square, turning it into difficult terrain for any creature that does not have slimewalking.
				For each square of slime a creature moves through, they must make a DC 12 DEX save or slip on the slime, moving into a random adjacent tile.
				If such a tile is also slime, they must make another check with disadvantage.
				On a 5 or lower on any movement check, the creature falls prone.`,
				WillPoints: 2,
			},
			{
				Id:          "slimewalking",
				Summary:     "You are immune to getting stuck in or slipping on liquids",
				Description: "You automatically succeed any saves to avoid slipping or getting stuck in liquids.",
				Indefinite:  true,
			},
		},
		Cost: 1,
	}
	SpinneretsMutation = Mutation{
		Id:          "spinnerets",
		Name:        "Spinnerets",
		Category:    PhysicalMutations,
		ImageUrl:    "https://wiki.cavesofqud.com/images/b/bf/Spinnerets_mutation.png",
		Description: "You can spin sticky silk webs",
		Abilities: []statblock.Ability{
			{
				Id:      "web spinning",
				UseTime: statblock.FreeAction,
				Summary: "You can spin webs, leaving a trail of them behind you",
				Description: `You begin spinning a web, leaving a 5ft wide, 2ft tall trail of webs behind you.
				You can stop spinning webs at any time and the webs will remain in place.
				The length of the trail is 25ft, plus 5ft for every level you have in this mutation.
				The webs are treated as difficult terrain for any creature without slimewalking or web immunity.
				Creatures that move through your webs must succeed on a STR saving throw to push through the webs or become stuck.
				The DC for this save is 15 + MUT. Any creature stuck in the web can repeat the save as an action on their next turn,
				freeing themselves on a success or remaining stuck on failure.`,
				WillPoints: 3,
			},
			{
				Id:          "web immunity",
				Summary:     "You are immune to getting stuck in webs",
				Description: "You are immune to getting stuck in another creature's webs or in any other substance.",
				Indefinite:  true,
			},
			{
				Id:          "arachnid kin",
				Summary:     "You are favoured by arachnids",
				Description: "When you gain this mutation, the party gains +6 reputation with arachnids.",
				Indefinite:  true,
			},
		},
		Cost: 3,
	}
	StingerConfusingMutation = Mutation{
		Id:          "confusing stinger",
		Name:        "Stinger (Confusing Venom)",
		Category:    PhysicalMutations,
		ImageUrl:    "https://wiki.cavesofqud.com/images/b/b7/Stinger_confusing_venom_mutation.png",
		Description: "You bear a tail with a stinger that delivers confusing venom to your enemies",
		Abilities: []statblock.Ability{
			{
				Id:      "confusing venom sting",
				Summary: "You can sting opponents with your stinger, inflicting confusing venom",
				Description: `Whenever you make a melee attack, roll a d10. On a 9 or 10, you make an additional attack with your stinger.
				The stinger is a natural long blade that can only be attacked with once per turn. Its to-hit bonus is 6 + (MUT + 2) / 3 (rounded down).
				The stinger's damage starts at 1d6 and increases at 4th, 7th, 10th, 13th, and 20th mutation levels to 1d8, 1d10, 1d12, 2d6, 2d6+1 and 2d8, respectively.
				`,
			},
		},
		Incompatibilities: []string{"paralyzing stinger", "poisoning stinger"},
		Cost:              3,
	}
)
