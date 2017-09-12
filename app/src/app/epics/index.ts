import { Observable } from 'rxjs';
import { combineEpics, Epic } from 'redux-observable';
import * as toolbox from 'toolbox';
import * as context from './context';


const allEpics: Epic<any, any>[] = [ ...toolbox.epics, ...context.epics ];

export const epics: Epic<any, any> = ( action$, store ) => combineEpics<any>( ...allEpics )( action$, store ).catch( ( err, source ) => {
    console.log( err );
    // TODO Throw error.
    return source;
} );