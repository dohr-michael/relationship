import * as React from 'react';
import { FormattedMessage } from 'react-intl'
import { MenuProps, Menu, Image, Button, Icon } from 'semantic-ui-react';

import { version } from 'app-config'

type AppTopBarProps = MenuProps & {
    titleClassName?: string;
    onClick: () => void;
};

const AppTopBar: React.SFC<AppTopBarProps> = ( { titleClassName, onClick, ...props } ) => (
    <Menu fixed="top" inverted { ...props }>
        <Menu.Item>
            <Button icon="content" onClick={ onClick } size="large" inverted />
            <strong className={ titleClassName }>
                <FormattedMessage tagName="span" id="app.title"/>
                <small>
                    <FormattedMessage tagName="em" id="app.version" values={ { version } }/>
                </small>
            </strong>
        </Menu.Item>
    </Menu>
);

export { AppTopBarProps, AppTopBar };