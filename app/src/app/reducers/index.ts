import { ReducersMapObject } from 'redux';
import * as toolbox from 'toolbox';
import * as context from './context';

export type Stores = toolbox.Stores & context.Store;

export const reducers: ReducersMapObject = {
    ...toolbox.reducers,
    ...context.Store,
};
