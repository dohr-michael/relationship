import { Dispatch } from 'react-redux';
import * as flux from '../flux';
import { Observable } from 'rxjs';


export function dispatchAsync<A = {}, Er = { [key: string]: any }>( a: flux.PromiseAction<A, Er>, dispatch: Dispatch<any> ): Promise<{}> {
    return new Promise( ( success, failed ) => {
        dispatch( { ...a, success, failed } );
    } );
}


export function fireAsyncSuccess<A = {}>( action: flux.PromiseAction<A>, obs?: Observable<any> ): Observable<any> {
    if( action.success ) action.success();
    return obs || Observable.empty();
}

export function fireAsyncError<A = {}, Er = {}>( action: flux.PromiseAction<A, Er>, err: Er, errToThrow?: any ): Observable<any> {
    if( action.failed ) action.failed( err );
    return Observable.throw( errToThrow || err );
}

export function epicAsync<A = {}, Er = { [key: string]: any }>( action: flux.PromiseAction<A, Er>, obs: Observable<any>, toEr: ( a: any ) => Er ): Observable<any> {
    return obs.map( c => {
        return fireAsyncSuccess( action, obs );
    } ).catch( ( err, _ ) => {
        return fireAsyncError( action, toEr( err ), err );
    } );
}
