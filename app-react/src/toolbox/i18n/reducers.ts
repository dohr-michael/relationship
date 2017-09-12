import { createSelector } from 'reselect';
import * as flux from '../flux';
import * as actions from './actions';

export type State = {
    defaultLocale: string;
    locales: string[];
    locale: string;
    messages: { [locale: string]: Object };
    formats: { [locale: string]: Object };
};
export const defaultState = (): State => ({
    locales:       [ 'fr-FR' ],
    defaultLocale: 'fr-FR',
    locale:        'fr-FR',
    messages:      {},
    formats:       {},
});
export type Stores = {
    i18n: State;
}

export namespace selectors {
    export const getState = createSelector( ( stores: Stores ) => stores.i18n, s => s );
    export const getLocales = createSelector( getState, ( state: State ) => state.locales );
    export const getLocale = createSelector( getState, ( state: State ) => state.locale );
    export const getDefaultLocale = createSelector( getState, ( state: State ) => state.defaultLocale );
    export const getMessages = createSelector( getState, getLocale, getDefaultLocale, ( state: State, locale: string, defaultLocale: string ) => {
        if( defaultLocale === locale ) return state.messages[ locale ] || {};
        return { ...(state.messages[ defaultLocale ] || {}), ...(state.messages[ locale ] || {}) };
    } );
    export const getFormats = createSelector( getState, getLocale, getDefaultLocale, ( state: State, locale: string, defaultLocale: string ) => {
        if( defaultLocale === locale ) return state.formats[ locale ] || {};
        return { ...(state.formats[ defaultLocale ] || {}), ...(state.formats[ locale ] || {}) };
    } );
}

export const reducers = {
    i18n: ( state: State = defaultState(), action: actions.Actions ): State => {
        return flux.Reducing
            .of( actions.Load.Type, load )
            .or( actions.ChangeLanguage.Type, changeLanguage )
            .exec( state, action );
    }
};


/** Convert the provided action to a readable message format */
function load( state: State, action: actions.Load ) {
    const messagesByLocale = Object.keys( action.payload.data.messages ).reduce( ( acc, c ) => ({ ...acc, [c]: flatten( action.payload.data.messages[ c ] ) }), {} );
    return {
        ...state,
        defaultLocale: action.payload.defaultLocale,
        locale:        action.payload.defaultLocale,
        formats:       action.payload.data.formats,
        messages:      messagesByLocale
    };
}

function changeLanguage( state: State, action: actions.ChangeLanguage ) {
    return {
        ...state,
        locale: action.payload,
    };
}

function flatten( obj: Object, parent: string = '' ): { [key: string]: string } {
    return Object.keys( obj ).reduce( ( acc, key ) => {
        const ck = parent === '' ? key : `${parent}.${key}`;
        if( typeof obj[ key ] === 'object' ) {
            return { ...acc, ...flatten( obj[ key ], ck ) };
        }
        return { ...acc, [ck]: obj[ key ] };
    }, {} );
}
