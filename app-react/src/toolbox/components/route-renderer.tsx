import * as React from 'react';
import { Switch, Route, Redirect } from 'react-router-dom';


type Path = {
    path: string;
    indexRedirect?: string;
    component?: React.ComponentType<any>;
    paths?: Path[];
}


type PathRendererProps = {
    paths: Path[];
    parentPath: string;
    indexRedirect?: string;
    defaultCmp?: React.ComponentType<any>;
    defaultChild?: React.ReactNode;
}

function getPath( path: Path, parent: string ): string { return path.path.startsWith( '/' ) ? path.path : `${parent}/${path.path}`; }

const RouteRenderer: React.ComponentType<PathRendererProps> = props => (
    <Switch>
        {
            props.paths.reduce( ( acc: React.ReactNode[], path ) => {
                const route = getPath( path, props.parentPath );
                const cmp = path.component;
                const key = `path_renderer.${props.parentPath}.${acc.length}`;
                let r: React.ReactNode = null;
                if( path.paths && path.paths.length > 0 ) {
                    r = (
                        <Route key={ key } path={ route }>
                            <RouteRenderer paths={ path.paths }
                                           parentPath={ route }
                                           indexRedirect={ path.indexRedirect }
                                           defaultCmp={ cmp }
                                           defaultChild={ props.defaultChild }/>
                        </Route>
                    );
                } else if( cmp ) {
                    r = <Route strict key={ key } path={ route } component={ cmp }/>;
                }
                return r ? [ ...acc, r ] : acc;
            }, [] )
        }
        { props.defaultChild ? <Route children={ props.defaultChild }/> : null }
        { props.defaultCmp ? <Route component={ props.defaultCmp }/> : null }
        { props.indexRedirect ? <Redirect to={ props.indexRedirect }/> : null }
    </Switch>
);

export { Path, PathRendererProps, RouteRenderer };