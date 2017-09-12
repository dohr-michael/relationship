import { Reducer, ReducersMapObject } from 'redux';
import { flux } from 'toolbox';

import { Universe } from 'app/models';
import * as actions from 'app/actions';


export type State = {
    loaded: boolean;
    totalStep: number;
    currentStep: number;
    userInfo?: any;
    universes?: Universe[];
};
export const defaultState = (): State => ({ loaded: false, totalStep: 2, currentStep: 0 });

export type Store = {
    context: State,
}

function userInfoLoaded( state: State, action: actions.context.UserInfoLoaded ) {
    return { ...state, currentStep: state.currentStep + 1, userInfo: action.payload };
}

function universesLoaded( state: State, action: actions.context.UniversesLoaded ) {
    return { ...state, currentStep: state.currentStep + 1, universes: action.payload };
}

export const reducer: Reducer<State> = ( state = defaultState(), action: actions.context.Actions ) =>
    flux.Reducing
        .of( actions.context.LoadContext.Type, ( s: State, a: actions.context.LoadContext ) => defaultState() )
        .or( actions.context.ContextLoaded.Type, ( s: State, a: actions.context.ContextLoaded ) => ({ ...s, loaded: true }) )
        .or( actions.context.UserInfoLoaded.Type, userInfoLoaded )
        .or( actions.context.UniversesLoaded.Type, universesLoaded )
        .exec( state, action );


export const Store: ReducersMapObject = {
    context: reducer,
};