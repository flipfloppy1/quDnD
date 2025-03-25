export namespace main {
	
	export enum Screen {
	    CHARACTER = "character",
	    CONCEPTS = "concepts",
	    CREATURES = "creatures",
	    ITEMS = "items",
	    LIQUIDS = "liquids",
	    LORE = "lore",
	    MECHANICS = "mechanics",
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
	    }
	}
	export class DiceRoll {
	    dice: string[];
	    offset: number;
	
	    static createFrom(source: any = {}) {
	        return new DiceRoll(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.dice = source["dice"];
	        this.offset = source["offset"];
	    }
	}
	export class DmgAffinity {
	    level: string;
	    dmgType: string;
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
	export class WpnEffect {
	    effect: string;
	    condition?: string;
	    reason?: string;
	
	    static createFrom(source: any = {}) {
	        return new WpnEffect(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.effect = source["effect"];
	        this.condition = source["condition"];
	        this.reason = source["reason"];
	    }
	}
	export class WpnUsageCondition {
	    condition: string;
	    reason?: string;
	
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
	    dmgType: string;
	    dmg: DiceRoll;
	    dmgVersatile?: DiceRoll;
	    reach: number;
	    wpnRange: WeaponRange;
	    conditions: WpnUsageCondition[];
	    effects: WpnEffect[];
	    statOffsets: StatOffset[];
	
	    static createFrom(source: any = {}) {
	        return new Weapon(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.dmgType = source["dmgType"];
	        this.dmg = this.convertValues(source["dmg"], DiceRoll);
	        this.dmgVersatile = this.convertValues(source["dmgVersatile"], DiceRoll);
	        this.reach = source["reach"];
	        this.wpnRange = this.convertValues(source["wpnRange"], WeaponRange);
	        this.conditions = this.convertValues(source["conditions"], WpnUsageCondition);
	        this.effects = this.convertValues(source["effects"], WpnEffect);
	        this.statOffsets = this.convertValues(source["statOffsets"], StatOffset);
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
	    stat: string;
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
	export class Statblock {
	    stats: Record<string, string>;
	    statOffsets: StatOffset;
	    dmgAffinities: DmgAffinity;
	    weapons: Weapon[];
	
	    static createFrom(source: any = {}) {
	        return new Statblock(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.stats = source["stats"];
	        this.statOffsets = this.convertValues(source["statOffsets"], StatOffset);
	        this.dmgAffinities = this.convertValues(source["dmgAffinities"], DmgAffinity);
	        this.weapons = this.convertValues(source["weapons"], Weapon);
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
	export class PageInfo {
	    pageType: Screen;
	    pageTitle: string;
	    description?: string;
	    statblock?: Statblock;
	
	    static createFrom(source: any = {}) {
	        return new PageInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.pageType = source["pageType"];
	        this.pageTitle = source["pageTitle"];
	        this.description = source["description"];
	        this.statblock = this.convertValues(source["statblock"], Statblock);
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

