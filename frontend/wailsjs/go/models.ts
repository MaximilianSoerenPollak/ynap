export namespace main {
	
	export class Transaction {
	    transaction_uuid?: number[];
	    name?: string;
	    amount?: number;
	    associated_portofolio?: number[];
	    type?: string;
	    assets?: number;
	    // Go type: time
	    date?: any;
	    // Go type: time
	    created_at?: any;
	
	    static createFrom(source: any = {}) {
	        return new Transaction(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.transaction_uuid = source["transaction_uuid"];
	        this.name = source["name"];
	        this.amount = source["amount"];
	        this.associated_portofolio = source["associated_portofolio"];
	        this.type = source["type"];
	        this.assets = source["assets"];
	        this.date = this.convertValues(source["date"], null);
	        this.created_at = this.convertValues(source["created_at"], null);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
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
	export class Portfolio {
	    portfolio_uuid?: number[];
	    name?: string;
	    associated_account?: number[];
	    balance?: number;
	    ledger?: Transaction[];
	    // Go type: time
	    updated_at?: any;
	    // Go type: time
	    created_at?: any;
	
	    static createFrom(source: any = {}) {
	        return new Portfolio(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.portfolio_uuid = source["portfolio_uuid"];
	        this.name = source["name"];
	        this.associated_account = source["associated_account"];
	        this.balance = source["balance"];
	        this.ledger = this.convertValues(source["ledger"], Transaction);
	        this.updated_at = this.convertValues(source["updated_at"], null);
	        this.created_at = this.convertValues(source["created_at"], null);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
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
	export class Account {
	    account_uuid?: number[];
	    first_name?: string;
	    last_name?: string;
	    balance?: number;
	    portfolios?: Portfolio[];
	    // Go type: time
	    updated_at?: any;
	    // Go type: time
	    created_at?: any;
	
	    static createFrom(source: any = {}) {
	        return new Account(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.account_uuid = source["account_uuid"];
	        this.first_name = source["first_name"];
	        this.last_name = source["last_name"];
	        this.balance = source["balance"];
	        this.portfolios = this.convertValues(source["portfolios"], Portfolio);
	        this.updated_at = this.convertValues(source["updated_at"], null);
	        this.created_at = this.convertValues(source["created_at"], null);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
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

