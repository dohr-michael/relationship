import * as React from 'react';
import * as ReactDom from 'react-dom';
import { Route, Switch, Redirect } from 'react-router';
import { ConnectedRouter, routerMiddleware } from 'react-router-redux';
import { Provider } from 'react-redux';
import createHistory from 'history/createHashHistory';
import { createStore, combineReducers, applyMiddleware } from 'redux';
import { composeWithDevTools } from 'redux-devtools-extension';
import { createEpicMiddleware } from 'redux-observable';
// Polyfills
import 'rxjs';
import 'babel-polyfill';
import 'whatwg-fetch';
// Styles
import './index.scss';
// Config
import { epics, reducers, States } from './core';

import { Flag } from 'semantic-ui-react';


// Create a history of your choosing (we're using a browser history in this case)
const history = createHistory();

// Build the middleware for intercepting and dispatching navigation actions
const middleware = routerMiddleware( history );


// Add the reducer to your store on the `router` key
// Also apply our middleware for navigating
const store = createStore(
    combineReducers<States>( {
        ...reducers
    } ),
    composeWithDevTools(
        applyMiddleware(
            createEpicMiddleware( epics ),
            middleware
        )
    )
);

// Init Authentication.
// store.dispatch( auth.actions.InitContext( 'ea00ba7e-1a25-4d08-9feb-b79ee34f1972', 'bcaa61da-9015-4cbe-9d92-ed22358357c7' ) );
/*store.dispatch( i18n.Load( {
    defaultLocale: 'en', data: {
        messages: { en: en.messages }, formats: { en: en.formats }
    }
} ) );
*/


const App = () => (
    <Provider store={ store }>
        <div>
            <div>Salut !</div>
            <Flag name="fr"/>
        </div>
    </Provider>
);

export default function () {
    const app = document.getElementById( 'app' );
    if( app != null ) {
        ReactDom.render( <App/>, app );
    }
}