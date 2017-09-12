import { Epic } from 'redux-observable';

import * as actions from 'app/actions';
import { Stores } from 'app/reducers';

import * as apis from 'app/apis';
import { Observable } from 'rxjs/Observable';


const listenLoadContext: Epic<actions.context.Actions, Stores> = ( action$, state ) =>
    action$.ofType( actions.context.LoadContext.Type )
        .flatMap( _ =>
            Observable.forkJoin(
                apis.loadUserInfo().do( c => state.dispatch( actions.context.UserInfoLoaded( c ) ) ),
                apis.loadUniverses().do( u => state.dispatch( actions.context.UniversesLoaded( u ) ) ),
                ( user, universes ) => actions.context.ContextLoaded()
            )
        );

export const epics: Epic<any, any>[] = [ listenLoadContext ];
