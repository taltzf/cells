/**
 * rest.proto
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * OpenAPI spec version: version not set
 * 
 *
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen.git
 * Do not edit the class manually.
 *
 */


import ApiClient from '../ApiClient';
import InstallInstallConfig from './InstallInstallConfig';





/**
* The InstallGetDefaultsResponse model module.
* @module model/InstallGetDefaultsResponse
* @version version not set
*/
export default class InstallGetDefaultsResponse {
    /**
    * Constructs a new <code>InstallGetDefaultsResponse</code>.
    * @alias module:model/InstallGetDefaultsResponse
    * @class
    */

    constructor() {
        

        
        

        

        
    }

    /**
    * Constructs a <code>InstallGetDefaultsResponse</code> from a plain JavaScript object, optionally creating a new instance.
    * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
    * @param {Object} data The plain JavaScript object bearing properties of interest.
    * @param {module:model/InstallGetDefaultsResponse} obj Optional instance to populate.
    * @return {module:model/InstallGetDefaultsResponse} The populated <code>InstallGetDefaultsResponse</code> instance.
    */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new InstallGetDefaultsResponse();

            
            
            

            if (data.hasOwnProperty('config')) {
                obj['config'] = InstallInstallConfig.constructFromObject(data['config']);
            }
        }
        return obj;
    }

    /**
    * @member {module:model/InstallInstallConfig} config
    */
    config = undefined;








}

