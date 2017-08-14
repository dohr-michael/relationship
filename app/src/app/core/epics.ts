import { Observable } from 'rxjs/Observable';
import { combineEpics, Epic } from 'redux-observable';

const allEpics: Epic<any, any>[] = [];

export const epics: Epic<any, any> = ( action$, store ) => combineEpics<any>( ...allEpics )( action$, store ).catch( ( err, source ) => {
    console.log( err );
    // TODO Throw error.
    return source;
} );
