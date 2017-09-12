import { Action } from 'redux';

/** Redux Action */
export { Action };

/** Redux Action with payload */
export interface ActionWP<A> extends Action {payload: A}

/** Action builder */
export type ActionFn = (() => Action) & { Type: string; };
/** Action with payload build */
export type ActionWPFn<A> = (( payload: A ) => ActionWP<A> ) & { Type: string; };

export interface PromiseAction<A = {}, Er = { [key: string]: any }> extends ActionWP<A> {
    success?: () => void;
    failed?: ( errors: Er ) => void;
}

export type PromiseActionFn<A, Er = { [key: string]: any }> = (( data: A, success?: () => void, failed?: ( errors: Er ) => void ) => PromiseAction<A, Er> ) & { Type: string };


export namespace Actions {
    function addStaticPart( a: any, type: string ): ActionFn | ActionWPFn<any> {
        a.Type = type;
        return a;
    }
    
    export function ofType( type: string ): ActionFn {
        return addStaticPart( () => ({ type }), type ) as ActionFn;
    }
    
    export function withPayload<A>( type: string ): ActionWPFn<A> {
        return addStaticPart( ( payload: A ) => ({ type, payload }), type ) as ActionWPFn<A>;
    }
    
    export function promise<A, Er = { [key: string]: any }>( type: string ): PromiseActionFn<A, Er> {
        return addStaticPart( ( payload: A, success?: () => void, failed?: ( errors: Er ) => void ) => ({ type, payload, success, failed }), type ) as PromiseActionFn<A, Er>;
    }
}
