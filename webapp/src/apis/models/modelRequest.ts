/**
 * 0G Serving User Broker API
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * OpenAPI spec version: 1.0
 * 
 *
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen.git
 * Do not edit the class manually.
 */


export interface ModelRequest { 
    fee: number;
    inputFee: number;
    nonce: number;
    previousOutputFee: number;
    providerAddress: string;
    serviceName: string;
    signature: string;
}
