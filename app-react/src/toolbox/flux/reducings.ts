import { Action } from './action'

/**
 *
 */
export class Reducing<S, A extends Action> {
    static of<S, A extends Action>( predicate: string | string[], fn: ( s: S, a: A ) => S ): Reducing<S, A> {
        const match = ( a: Action ): boolean => {
            if(Array.isArray( predicate )) return (predicate as string[]).indexOf( a.type ) > -1;
            return predicate === a.type;
        };
        return new Reducing<S, A>( match, fn );
    }
    
    constructor( private match: ( a: Action ) => boolean, private fn: ( s: S, a: A ) => S ) {}
    
    or<A1 extends Action>( predicate: string | string[], fn: ( s: S, a: A1 ) => S ): Reducing<S, A | A1> {
        return this.orElse( Reducing.of( predicate, fn ) );
    }
    
    orElse<A1 extends Action>( c: Reducing<S, A1> ): Reducing<S, A | A1> {
        return new Reducing<S, A | A1>( a => this.match( a ) || c.match( a ), ( s: S, a: A | A1 ) => {
            if( this.match( a ) ) return this.fn( s, a as A );
            else if( c.match( a ) ) return c.fn( s, a as A1 );
            return s;
        } )
    }
    
    exec( s: S, a: A ): S {
        if( this.match( a ) ) return this.fn( s, a );
        return s;
    }
}
