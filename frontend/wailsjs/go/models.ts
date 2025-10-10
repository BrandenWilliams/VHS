export namespace linuxcliargs {
	
	export class LinuxCLICfg {
	    vcodec: string;
	    acodec: string;
	    container: string;
	    mode: string;
	    crf: number;
	    cq: number;
	    vbitrate: string;
	    maxrate: string;
	    buffsize: string;
	    preset: string;
	    tune: string;
	    profile: string;
	    level: string;
	    pixfmt: string;
	    scale: string;
	    fps: string;
	    audiobitrate: string;
	    audoch: number;
	    audiorate: number;
	    extrainputargs: string[];
	    extrafilterargs: string[];
	    extraoutputargs: string[];
	
	    static createFrom(source: any = {}) {
	        return new LinuxCLICfg(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.vcodec = source["vcodec"];
	        this.acodec = source["acodec"];
	        this.container = source["container"];
	        this.mode = source["mode"];
	        this.crf = source["crf"];
	        this.cq = source["cq"];
	        this.vbitrate = source["vbitrate"];
	        this.maxrate = source["maxrate"];
	        this.buffsize = source["buffsize"];
	        this.preset = source["preset"];
	        this.tune = source["tune"];
	        this.profile = source["profile"];
	        this.level = source["level"];
	        this.pixfmt = source["pixfmt"];
	        this.scale = source["scale"];
	        this.fps = source["fps"];
	        this.audiobitrate = source["audiobitrate"];
	        this.audoch = source["audoch"];
	        this.audiorate = source["audiorate"];
	        this.extrainputargs = source["extrainputargs"];
	        this.extrafilterargs = source["extrafilterargs"];
	        this.extraoutputargs = source["extraoutputargs"];
	    }
	}
	export class PreSet {
	    id: number;
	    name: string;
	    lcfg: LinuxCLICfg;
	
	    static createFrom(source: any = {}) {
	        return new PreSet(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.lcfg = this.convertValues(source["lcfg"], LinuxCLICfg);
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

