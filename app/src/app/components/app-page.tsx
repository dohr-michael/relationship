import * as React from 'react';
import { Sidebar } from 'semantic-ui-react';

import { Universe } from 'app/models';
import { AppSidebar, AppTopBar, AppContent } from './app-page-elements';

import './app-page.scss';


type AppPageProps = {
    universes: Universe[];
    userInfo: any;
    languages: string[];
    currentLanguage: string;
};

class AppPage extends React.Component<AppPageProps, { open: boolean }> {
    state = { open: true };
    updateOpen = () => { this.setState( { open: !this.state.open } );};
    
    
    render(): JSX.Element | any {
        const { universes, userInfo, languages, currentLanguage, ...props } = this.props;
        const { open } = this.state;
        return (
            <div className="app-page">
                <AppTopBar className="app-page__topbar"
                           titleClassName="app-page__topbar__title"
                           onClick={ this.updateOpen }/>
                <div className="app-page__body">
                    <Sidebar.Pushable>
                        <Sidebar className="app-page__body__sidebar"
                                 as={ AppSidebar }
                                 animation="push"
                                 visible={ open }
                                 universes={ universes }/>
                        <Sidebar.Pusher>
                            <AppContent className="app-page__body__content">
                                { props.children }
                            </AppContent>
                        </Sidebar.Pusher>
                    </Sidebar.Pushable>
                </div>
            </div>
        );
    }
}

export { AppPageProps, AppPage };
