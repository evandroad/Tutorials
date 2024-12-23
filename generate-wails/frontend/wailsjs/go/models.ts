export namespace main {
	
	export class Content {
	    id: string;
	    number: number;
	    title: string;
	    content: string;
	
	    static createFrom(source: any = {}) {
	        return new Content(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.number = source["number"];
	        this.title = source["title"];
	        this.content = source["content"];
	    }
	}
	export class Tutorial {
	    id: string;
	    number: number;
	    title: string;
	    image: string;
	
	    static createFrom(source: any = {}) {
	        return new Tutorial(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.number = source["number"];
	        this.title = source["title"];
	        this.image = source["image"];
	    }
	}

}

