import { addLocaleData } from 'react-intl';
import * as frLocaleData from 'react-intl/locale-data/fr';

import * as m from './messages';
import f from './formats';


addLocaleData( frLocaleData );

export const messages = {
    ...m,
};
export const formats = {
    ...f,
};