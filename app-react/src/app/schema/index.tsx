import * as React from 'react';
import { Path } from 'toolbox';
import { Label } from 'semantic-ui-react';
import { Link } from 'react-router-dom';
import { SchemaEditor } from './containers';

export * from './models';
export * from './containers';
export * from './epics';

export const content = ( root: string = '/schemas' ): Path => ({
    path:          root,
    indexRedirect: `${root}/list`,
    paths:         [
        {
            path:      `list`,
            component: () => (
                <div>
                    <Label as={ Link } to={ `${root}/new` }>Create new schema</Label>
                </div>
            )
        },
        {
            path:      `:id`,
            component: SchemaEditor,
        }
    ]
});


