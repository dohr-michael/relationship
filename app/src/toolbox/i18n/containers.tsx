import * as React from 'react';
import { IntlProvider } from 'react-intl';
import { connect } from 'react-redux';

import { Stores, selectors } from './reducers';


function mapStateToProps( state: Stores, props: {} ) {
    return {
        defaultLocale: selectors.getDefaultLocale( state ),
        locale:        selectors.getLocale( state ),
        messages:      selectors.getMessages( state ),
        formats:       selectors.getFormats( state ),
    };
}

const I18nProviderBase: React.ComponentType<any> = props => <IntlProvider { ...props } />;
const I18nProvider = connect( mapStateToProps )( I18nProviderBase );

export { I18nProvider };