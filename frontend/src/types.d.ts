// REQUEST TYPES 

declare module "tstypes" {
    export interface AccountRequest { 
        firstName: string
        lastName: string
    }


    // RETURN TYPES
    export interface AccountReturn { 
        accountUUID: string
        firstName: string
        lastName: string
        balance: number
        portfolios: [PortfolioReturn]
        updatedAt: Date
        createdAt: Date

    }
    export interface PortfolioReturn {
        portfolioUUID:     string     
        name:              string        
        associatedAccount: string     
        balance:           number       
        ledger:            [Transaction]
        updatedAt:         Date     
        createdAt:         Date     
    }

    export interface Transaction {
        transactionUUID:      string 
        name:                 string    
        amount:               number   
        associatedPortofolio: string 
        type:                 string    
        assets:               number   
        date:                 Date
        createdAt:            Date 
    }
}
