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
	
	
	    static createFrom(source: any = {}) {
	        return new CategoryMembers(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	
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

