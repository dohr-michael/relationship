import { flux } from 'toolbox';
import { Universe } from 'app/models';


export type LoadContext = flux.Action;
export const LoadContext = flux.Actions.ofType( 'context#load' );

export type ContextLoaded = flux.Action;
export const ContextLoaded = flux.Actions.ofType( 'context#loaded' );

export type UserInfoLoaded = flux.ActionWP<any>;
export const UserInfoLoaded = flux.Actions.withPayload<any>( 'context#user#loaded' );

export type UniversesLoaded = flux.ActionWP<Universe[]>;
export const UniversesLoaded = flux.Actions.withPayload<Universe[]>( 'context#universes#loaded' );


export type Actions = LoadContext | UserInfoLoaded | UniversesLoaded;