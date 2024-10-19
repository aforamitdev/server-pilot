export namespace apiv1 {
	
	export class GetStatusResponse {
	    system?: string;
	
	    static createFrom(source: any = {}) {
	        return new GetStatusResponse(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.system = source["system"];
	    }
	}

}

