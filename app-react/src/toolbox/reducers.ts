import * as i18n from './i18n';
import { ReducersMapObject } from 'redux';
import * as reduxForm from 'redux-form';
import * as reduxRouter from 'react-router-redux';


export type Stores = i18n.Stores & {
    form: reduxForm.FormState;
    router: reduxRouter.RouterState;
};
export const reducers: ReducersMapObject = {
    ...i18n.reducers,
    form:   reduxForm.reducer,
    router: reduxRouter.routerReducer,
};