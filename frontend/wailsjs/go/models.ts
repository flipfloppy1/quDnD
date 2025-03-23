export namespace main {
	
	export enum Screen {
	    SEARCH = "Search",
	    CREATURES = "Creatures",
	    OBJECTS = "Objects",
	    LIQUIDS = "Liquids",
	    LORE = "Lore",
	    MECHANICS = "Mechanics",
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
	export class PageInfo {
	
	
	    static createFrom(source: any = {}) {
	        return new PageInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	
	    }
	}

}

