import * as assert from 'power-assert';
import { JsonSchema, compileSchema, JsonSchemaVersion } from './data';


interface Data {
    id: string;
}

namespace Data {
    
    export const schema: JsonSchema = {
        $schema:    JsonSchemaVersion,
        required:   [ 'id' ],
        properties: {
            id: {
                type:    'string',
                default: 'titi'
            }
        }
    };
    
    export const validate = compileSchema( schema );
}

describe( 'toolbox/data', () => {
    it( 'do something', () => {
        const data = {};
        const isValid = Data.validate( data );
        
        console.log( isValid, Data.validate.errors, data as Data );
    } );
} );