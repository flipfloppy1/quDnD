export namespace main {
	
	export class PageInfo {
	    pageType: pageUtils.Screen;
	    pageTitle: string;
	    imgSrc?: string;
	    description?: string;
	    statblock?: statblock.Statblock;
	    pageid: number;
	
	    static createFrom(source: any = {}) {
	        return new PageInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.pageType = source["pageType"];
	        this.pageTitle = source["pageTitle"];
	        this.imgSrc = source["imgSrc"];
	        this.description = source["description"];
	        this.statblock = this.convertValues(source["statblock"], statblock.Statblock);
	        this.pageid = source["pageid"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

export namespace pageUtils {
	
	export enum Screen {
	    SEARCH = "search",
	    CHARACTER = "character",
	    CONCEPTS = "concepts",
	    CREATURES = "creatures",
	    ITEMS = "items",
	    LIQUIDS = "liquids",
	    LORE = "lore",
	    MECHANICS = "mechanics",
	    MUTATIONS = "mutations",
	    OTHER = "other",
	    CUSTOM = "custom",
	}
	export class CategoryMembers {
	    liquids: number[];
	    creatures: number[];
	    items: number[];
	    characters: number[];
	    concepts: number[];
	    world: number[];
	    mechanics: number[];
	    mutations: number[];
	
	    static createFrom(source: any = {}) {
	        return new CategoryMembers(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.liquids = source["liquids"];
	        this.creatures = source["creatures"];
	        this.items = source["items"];
	        this.characters = source["characters"];
	        this.concepts = source["concepts"];
	        this.world = source["world"];
	        this.mechanics = source["mechanics"];
	        this.mutations = source["mutations"];
	    }
	}
	export class RestPageSearchResults {
	    id: number;
	    key: string;
	    title: string;
	    excerpt: string;
	    matched_title: any;
	    description: any;
	    // Go type: struct { Mimetype string "json:\"mimetype\""; Size int "json:\"size\""; Width int "json:\"width\""; Height int "json:\"height\""; Duration interface {} "json:\"duration\""; URL string "json:\"url\"" }
	    thumbnail: any;
	
	    static createFrom(source: any = {}) {
	        return new RestPageSearchResults(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.key = source["key"];
	        this.title = source["title"];
	        this.excerpt = source["excerpt"];
	        this.matched_title = source["matched_title"];
	        this.description = source["description"];
	        this.thumbnail = this.convertValues(source["thumbnail"], Object);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class RestPageSearch {
	    pages: RestPageSearchResults[];
	
	    static createFrom(source: any = {}) {
	        return new RestPageSearch(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.pages = this.convertValues(source["pages"], RestPageSearchResults);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

export namespace statblock {
	
	export enum Stat {
	    AC = "ac",
	    SPEED = "speed",
	    LEVEL = "level",
	    PROFICIENCY = "proficiency",
	    HP = "hp",
	    STR = "str",
	    DEX = "dex",
	    CON = "con",
	    INT = "int",
	    WIS = "wis",
	    CHA = "cha",
	    INITIATIVE = "initiative",
	    STRSAVE = "strsave",
	    DEXSAVE = "dexsave",
	    CONSAVE = "consave",
	    INTSAVE = "intsave",
	    WISSAVE = "wissave",
	    CHASAVE = "chasave",
	}
	export enum Duration {
	    ACTION = "action",
	    REACTION = "reaction",
	    ITEM_INTERACTION = "item_interaction",
	    BONUS_ACTION = "bonus_action",
	    FREE_ACTION = "free_action",
	}
	export enum DamageType {
	    DMGACID = "dmgacid",
	    DMGBLUDGEONING = "dmgbludgeoning",
	    DMGCOLD = "dmgcold",
	    DMGFIRE = "dmgfire",
	    DMGFORCE = "dmgforce",
	    DMGLIGHTNING = "dmglightning",
	    DMGNECROTIC = "dmgnecrotic",
	    DMGPIERCING = "dmgpiercing",
	    DMGPOISON = "dmgpoison",
	    DMGPSYCHIC = "dmgpsychic",
	    DMGRADIANT = "dmgradiant",
	    DMGSLASHING = "dmgslashing",
	    DMGTHUNDER = "dmgthunder",
	}
	export enum DmgAffinityLevel {
	    RESISTANT = "resistant",
	    WEAK = "weak",
	    IMMUNE = "immune",
	}
	export class Effect {
	    effect: string;
	    conditions: string[];
	    reasons: string[];
	    Rounds: DiceRoll;
	
	    static createFrom(source: any = {}) {
	        return new Effect(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.effect = source["effect"];
	        this.conditions = source["conditions"];
	        this.reasons = source["reasons"];
	        this.Rounds = this.convertValues(source["Rounds"], DiceRoll);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class DiceRoll {
	    dice: string[];
	    offset: number;
	    statBonus: Stat;
	
	    static createFrom(source: any = {}) {
	        return new DiceRoll(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.dice = source["dice"];
	        this.offset = source["offset"];
	        this.statBonus = source["statBonus"];
	    }
	}
	export class Attack {
	    dmgType: DamageType;
	    damage: DiceRoll;
	    conditions: string[];
	
	    static createFrom(source: any = {}) {
	        return new Attack(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.dmgType = source["dmgType"];
	        this.damage = this.convertValues(source["damage"], DiceRoll);
	        this.conditions = source["conditions"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class Ability {
	    duration: Duration;
	    summary: string;
	    description: string;
	    conditions: string[];
	    attacks: Attack[];
	    effects: Effect[];
	
	    static createFrom(source: any = {}) {
	        return new Ability(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.duration = source["duration"];
	        this.summary = source["summary"];
	        this.description = source["description"];
	        this.conditions = source["conditions"];
	        this.attacks = this.convertValues(source["attacks"], Attack);
	        this.effects = this.convertValues(source["effects"], Effect);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	
	
	export class DmgAffinity {
	    level: DmgAffinityLevel;
	    dmgType: DamageType;
	    reason?: string;
	    condition?: string;
	
	    static createFrom(source: any = {}) {
	        return new DmgAffinity(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.level = source["level"];
	        this.dmgType = source["dmgType"];
	        this.reason = source["reason"];
	        this.condition = source["condition"];
	    }
	}
	
	export class FeatBuff {
	    stat: Stat;
	    value: string;
	    conditions: string[];
	
	    static createFrom(source: any = {}) {
	        return new FeatBuff(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.stat = source["stat"];
	        this.value = source["value"];
	        this.conditions = source["conditions"];
	    }
	}
	export class Feat {
	    id: string;
	    name: string;
	    buffs: FeatBuff[];
	    abilities: Ability[];
	    description: string;
	    prereqs: string[];
	
	    static createFrom(source: any = {}) {
	        return new Feat(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.buffs = this.convertValues(source["buffs"], FeatBuff);
	        this.abilities = this.convertValues(source["abilities"], Ability);
	        this.description = source["description"];
	        this.prereqs = source["prereqs"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	
	export class StatOffset {
	    stat: Stat;
	    value: string;
	    reason?: string;
	    condition?: string;
	
	    static createFrom(source: any = {}) {
	        return new StatOffset(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.stat = source["stat"];
	        this.value = source["value"];
	        this.reason = source["reason"];
	        this.condition = source["condition"];
	    }
	}
	export class WpnUsageCondition {
	    condition: string;
	    reason: string[];
	
	    static createFrom(source: any = {}) {
	        return new WpnUsageCondition(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.condition = source["condition"];
	        this.reason = source["reason"];
	    }
	}
	export class WeaponRange {
	    normal: number;
	    long: number;
	
	    static createFrom(source: any = {}) {
	        return new WeaponRange(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.normal = source["normal"];
	        this.long = source["long"];
	    }
	}
	export class Weapon {
	    name: string;
	    imageUrl: string;
	    dmgType: DamageType;
	    dmg: DiceRoll;
	    dmgVersatile?: DiceRoll;
	    penetration: number;
	    wpnRange: WeaponRange;
	    conditions: WpnUsageCondition[];
	    effects: Effect[];
	    statOffsets: StatOffset[];
	    pageid: number;
	
	    static createFrom(source: any = {}) {
	        return new Weapon(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.imageUrl = source["imageUrl"];
	        this.dmgType = source["dmgType"];
	        this.dmg = this.convertValues(source["dmg"], DiceRoll);
	        this.dmgVersatile = this.convertValues(source["dmgVersatile"], DiceRoll);
	        this.penetration = source["penetration"];
	        this.wpnRange = this.convertValues(source["wpnRange"], WeaponRange);
	        this.conditions = this.convertValues(source["conditions"], WpnUsageCondition);
	        this.effects = this.convertValues(source["effects"], Effect);
	        this.statOffsets = this.convertValues(source["statOffsets"], StatOffset);
	        this.pageid = source["pageid"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class Statblock {
	    stats: Record<string, string>;
	    statOffsets: StatOffset[];
	    dmgAffinities: DmgAffinity[];
	    items: Weapon[];
	    feats: Feat[];
	
	    static createFrom(source: any = {}) {
	        return new Statblock(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.stats = source["stats"];
	        this.statOffsets = this.convertValues(source["statOffsets"], StatOffset);
	        this.dmgAffinities = this.convertValues(source["dmgAffinities"], DmgAffinity);
	        this.items = this.convertValues(source["items"], Weapon);
	        this.feats = this.convertValues(source["feats"], Feat);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	
	

}

