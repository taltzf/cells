/**
 * Pydio Cell Rest API
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * OpenAPI spec version: 1.0
 * 
 *
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen.git
 * Do not edit the class manually.
 *
 */


import ApiClient from '../ApiClient';
import InstallCheckResult from './InstallCheckResult';





/**
* The InstallInstallConfig model module.
* @module model/InstallInstallConfig
* @version 1.0
*/
export default class InstallInstallConfig {
    /**
    * Constructs a new <code>InstallInstallConfig</code>.
    * @alias module:model/InstallInstallConfig
    * @class
    */

    constructor() {
        

        
        

        

        
    }

    /**
    * Constructs a <code>InstallInstallConfig</code> from a plain JavaScript object, optionally creating a new instance.
    * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
    * @param {Object} data The plain JavaScript object bearing properties of interest.
    * @param {module:model/InstallInstallConfig} obj Optional instance to populate.
    * @return {module:model/InstallInstallConfig} The populated <code>InstallInstallConfig</code> instance.
    */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new InstallInstallConfig();

            
            
            

            if (data.hasOwnProperty('dbConnectionType')) {
                obj['dbConnectionType'] = ApiClient.convertToType(data['dbConnectionType'], 'String');
            }
            if (data.hasOwnProperty('dbTCPHostname')) {
                obj['dbTCPHostname'] = ApiClient.convertToType(data['dbTCPHostname'], 'String');
            }
            if (data.hasOwnProperty('dbTCPPort')) {
                obj['dbTCPPort'] = ApiClient.convertToType(data['dbTCPPort'], 'String');
            }
            if (data.hasOwnProperty('dbTCPName')) {
                obj['dbTCPName'] = ApiClient.convertToType(data['dbTCPName'], 'String');
            }
            if (data.hasOwnProperty('dbTCPUser')) {
                obj['dbTCPUser'] = ApiClient.convertToType(data['dbTCPUser'], 'String');
            }
            if (data.hasOwnProperty('dbTCPPassword')) {
                obj['dbTCPPassword'] = ApiClient.convertToType(data['dbTCPPassword'], 'String');
            }
            if (data.hasOwnProperty('dbSocketFile')) {
                obj['dbSocketFile'] = ApiClient.convertToType(data['dbSocketFile'], 'String');
            }
            if (data.hasOwnProperty('dbSocketName')) {
                obj['dbSocketName'] = ApiClient.convertToType(data['dbSocketName'], 'String');
            }
            if (data.hasOwnProperty('dbSocketUser')) {
                obj['dbSocketUser'] = ApiClient.convertToType(data['dbSocketUser'], 'String');
            }
            if (data.hasOwnProperty('dbSocketPassword')) {
                obj['dbSocketPassword'] = ApiClient.convertToType(data['dbSocketPassword'], 'String');
            }
            if (data.hasOwnProperty('dbManualDSN')) {
                obj['dbManualDSN'] = ApiClient.convertToType(data['dbManualDSN'], 'String');
            }
            if (data.hasOwnProperty('dsName')) {
                obj['dsName'] = ApiClient.convertToType(data['dsName'], 'String');
            }
            if (data.hasOwnProperty('dsPort')) {
                obj['dsPort'] = ApiClient.convertToType(data['dsPort'], 'String');
            }
            if (data.hasOwnProperty('dsFolder')) {
                obj['dsFolder'] = ApiClient.convertToType(data['dsFolder'], 'String');
            }
            if (data.hasOwnProperty('externalMicro')) {
                obj['externalMicro'] = ApiClient.convertToType(data['externalMicro'], 'String');
            }
            if (data.hasOwnProperty('externalGateway')) {
                obj['externalGateway'] = ApiClient.convertToType(data['externalGateway'], 'String');
            }
            if (data.hasOwnProperty('externalWebsocket')) {
                obj['externalWebsocket'] = ApiClient.convertToType(data['externalWebsocket'], 'String');
            }
            if (data.hasOwnProperty('externalFrontPlugins')) {
                obj['externalFrontPlugins'] = ApiClient.convertToType(data['externalFrontPlugins'], 'String');
            }
            if (data.hasOwnProperty('externalDex')) {
                obj['externalDex'] = ApiClient.convertToType(data['externalDex'], 'String');
            }
            if (data.hasOwnProperty('externalDexID')) {
                obj['externalDexID'] = ApiClient.convertToType(data['externalDexID'], 'String');
            }
            if (data.hasOwnProperty('externalDexSecret')) {
                obj['externalDexSecret'] = ApiClient.convertToType(data['externalDexSecret'], 'String');
            }
            if (data.hasOwnProperty('frontendHosts')) {
                obj['frontendHosts'] = ApiClient.convertToType(data['frontendHosts'], 'String');
            }
            if (data.hasOwnProperty('frontendLogin')) {
                obj['frontendLogin'] = ApiClient.convertToType(data['frontendLogin'], 'String');
            }
            if (data.hasOwnProperty('frontendPassword')) {
                obj['frontendPassword'] = ApiClient.convertToType(data['frontendPassword'], 'String');
            }
            if (data.hasOwnProperty('frontendRepeatPassword')) {
                obj['frontendRepeatPassword'] = ApiClient.convertToType(data['frontendRepeatPassword'], 'String');
            }
            if (data.hasOwnProperty('fpmAddress')) {
                obj['fpmAddress'] = ApiClient.convertToType(data['fpmAddress'], 'String');
            }
            if (data.hasOwnProperty('licenseRequired')) {
                obj['licenseRequired'] = ApiClient.convertToType(data['licenseRequired'], 'Boolean');
            }
            if (data.hasOwnProperty('licenseString')) {
                obj['licenseString'] = ApiClient.convertToType(data['licenseString'], 'String');
            }
            if (data.hasOwnProperty('CheckResults')) {
                obj['CheckResults'] = ApiClient.convertToType(data['CheckResults'], [InstallCheckResult]);
            }
        }
        return obj;
    }

    /**
    * @member {String} dbConnectionType
    */
    dbConnectionType = undefined;
    /**
    * @member {String} dbTCPHostname
    */
    dbTCPHostname = undefined;
    /**
    * @member {String} dbTCPPort
    */
    dbTCPPort = undefined;
    /**
    * @member {String} dbTCPName
    */
    dbTCPName = undefined;
    /**
    * @member {String} dbTCPUser
    */
    dbTCPUser = undefined;
    /**
    * @member {String} dbTCPPassword
    */
    dbTCPPassword = undefined;
    /**
    * @member {String} dbSocketFile
    */
    dbSocketFile = undefined;
    /**
    * @member {String} dbSocketName
    */
    dbSocketName = undefined;
    /**
    * @member {String} dbSocketUser
    */
    dbSocketUser = undefined;
    /**
    * @member {String} dbSocketPassword
    */
    dbSocketPassword = undefined;
    /**
    * @member {String} dbManualDSN
    */
    dbManualDSN = undefined;
    /**
    * @member {String} dsName
    */
    dsName = undefined;
    /**
    * @member {String} dsPort
    */
    dsPort = undefined;
    /**
    * @member {String} dsFolder
    */
    dsFolder = undefined;
    /**
    * @member {String} externalMicro
    */
    externalMicro = undefined;
    /**
    * @member {String} externalGateway
    */
    externalGateway = undefined;
    /**
    * @member {String} externalWebsocket
    */
    externalWebsocket = undefined;
    /**
    * @member {String} externalFrontPlugins
    */
    externalFrontPlugins = undefined;
    /**
    * @member {String} externalDex
    */
    externalDex = undefined;
    /**
    * @member {String} externalDexID
    */
    externalDexID = undefined;
    /**
    * @member {String} externalDexSecret
    */
    externalDexSecret = undefined;
    /**
    * @member {String} frontendHosts
    */
    frontendHosts = undefined;
    /**
    * @member {String} frontendLogin
    */
    frontendLogin = undefined;
    /**
    * @member {String} frontendPassword
    */
    frontendPassword = undefined;
    /**
    * @member {String} frontendRepeatPassword
    */
    frontendRepeatPassword = undefined;
    /**
    * @member {String} fpmAddress
    */
    fpmAddress = undefined;
    /**
    * @member {Boolean} licenseRequired
    */
    licenseRequired = undefined;
    /**
    * @member {String} licenseString
    */
    licenseString = undefined;
    /**
    * @member {Array.<module:model/InstallCheckResult>} CheckResults
    */
    CheckResults = undefined;








}


