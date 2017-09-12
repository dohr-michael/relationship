import * as React from 'react';
import { Menu, MenuProps } from 'semantic-ui-react';
import { RouterMenuItem } from 'toolbox';

import { Universe } from 'app/models';


type AppSidebarProps = MenuProps & {
    universes: Universe[];
};

const AppSidebar: React.SFC<AppSidebarProps> = ( { universes, ...props } ) => (
    <Menu vertical { ...props } >
        <Menu.Item>
            <Menu.Header content={ 'Universes' }/>
            <Menu.Menu>
                { universes.map( u => <RouterMenuItem key={ u.id } content={ u.name } to={ `/universes/${u.id}` } exact/> ) }
            </Menu.Menu>
        </Menu.Item>
        <Menu.Item>
            <Menu.Header content={ 'Administration' }/>
            <Menu.Menu>
                <RouterMenuItem content="Manage data schema" to="/" exact/>
            </Menu.Menu>
        </Menu.Item>
    </Menu>
);

export { AppSidebarProps, AppSidebar };