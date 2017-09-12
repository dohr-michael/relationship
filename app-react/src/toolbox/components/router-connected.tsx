import * as React from 'react';
import { connect, Dispatch } from 'react-redux';
import { Menu, MenuItemProps } from 'semantic-ui-react';
import { push } from 'react-router-redux';
import { Stores } from '../reducers';


type RouterConnectedProps = {
    to: string;
    exact?: boolean;
}

type MappedProps = {
    active: boolean,
    pushTo: () => void;
}

function mapStateToProps( state: Stores, props: RouterConnectedProps ): { active: boolean } {
    const location = state.router.location;
    return {
        active: !!location && (props.exact && location.pathname === props.to || !props.exact && location.pathname.startsWith( props.to ) ),
    };
}

function mapDispatchToProps( dispatch: Dispatch<any>, props: RouterConnectedProps ) {
    return {
        pushTo: () => { dispatch( push( props.to ) ); }
    };
}

const RouterMenuItemBase: React.SFC<RouterConnectedProps & MenuItemProps & MappedProps> = ( { to, exact, pushTo, ...props } ) => (
    <Menu.Item onClick={ pushTo } { ...props } />
);
const RouterMenuItem: React.ComponentType<RouterConnectedProps & MenuItemProps> = connect( mapStateToProps, mapDispatchToProps )( RouterMenuItemBase );


export { RouterConnectedProps, RouterMenuItem };
