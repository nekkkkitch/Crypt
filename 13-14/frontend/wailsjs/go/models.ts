export namespace cryptanalysis {
	
	export class Correlation {
	    amount: number;
	    coefficient: number;
	    t: number;
	
	    static createFrom(source: any = {}) {
	        return new Correlation(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.amount = source["amount"];
	        this.coefficient = source["coefficient"];
	        this.t = source["t"];
	    }
	}
	export class ACData {
	    correlations: Correlation[];
	
	    static createFrom(source: any = {}) {
	        return new ACData(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.correlations = this.convertValues(source["correlations"], Correlation);
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
	
	export class  {
	
	
	    static createFrom(source: any = {}) {
	        return new (source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	
	    }
	}
	export class KasiskiData {
	    possible_dividers: Record<number, >;
	
	    static createFrom(source: any = {}) {
	        return new KasiskiData(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.possible_dividers = this.convertValues(source["possible_dividers"], , true);
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
	export class SymbolFreq {
	    Symbol: string;
	    Freq: number;
	    Place: number;
	
	    static createFrom(source: any = {}) {
	        return new SymbolFreq(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Symbol = source["Symbol"];
	        this.Freq = source["Freq"];
	        this.Place = source["Place"];
	    }
	}
	export class Place {
	    part: number;
	    number: number;
	
	    static createFrom(source: any = {}) {
	        return new Place(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.part = source["part"];
	        this.number = source["number"];
	    }
	}
	export class FreqData {
	    frequencies: Record<string, Place>;
	    sorted: SymbolFreq[];
	    in_stupid: string;
	
	    static createFrom(source: any = {}) {
	        return new FreqData(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.frequencies = this.convertValues(source["frequencies"], Place, true);
	        this.sorted = this.convertValues(source["sorted"], SymbolFreq);
	        this.in_stupid = source["in_stupid"];
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
	export class Data {
	    freq_data: FreqData;
	    ac_data: ACData;
	    k_data: KasiskiData;
	    index_data: number;
	
	    static createFrom(source: any = {}) {
	        return new Data(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.freq_data = this.convertValues(source["freq_data"], FreqData);
	        this.ac_data = this.convertValues(source["ac_data"], ACData);
	        this.k_data = this.convertValues(source["k_data"], KasiskiData);
	        this.index_data = source["index_data"];
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

