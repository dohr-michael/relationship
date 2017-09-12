import { Epic } from 'redux-observable';
import { Observable } from 'rxjs';
import { flux, FormErrors } from 'toolbox';
import { Stores } from 'app/core';
import * as actions from '../actions';


const validate: Epic<actions.Actions, Stores> = ( $action, state ) =>
    $action.ofType( actions.Validate.Type )
        .flatMap( ( a: actions.Validate ) => {
            console.log( 'validate' );
            const schema = a.payload.schema;
            const field = a.payload.field;
            let result: Observable<any>;
            if( schema.type === 'toto' && field === 'type' ) {
                result = Observable.throw( { type: 'type already exists' } );
            } else {
                result = Observable.empty();
            }
            return flux.epicAsync( a, result, a => a );
        } );


export const epics: Epic<any, any>[] = [ validate ];