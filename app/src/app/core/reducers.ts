import { ReducersMapObject } from 'redux';
import { routerReducer, RouterState } from 'react-router-redux';


export type States = {
    router: RouterState;
}

export const reducers: ReducersMapObject = {
    router: routerReducer,
};
