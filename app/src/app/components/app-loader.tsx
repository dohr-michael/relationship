import * as React from 'react';
import { Loader, Progress, Dimmer, Segment, Transition } from 'semantic-ui-react';

type AppLoaderProps = {
    totalStep: number;
    step: number;
    load: () => void;
    loaded: boolean;
};

class AppLoader extends React.Component<AppLoaderProps> {
    
    componentDidMount(): void { if( !this.props.loaded ) { this.props.load();} }
    
    render(): JSX.Element | any {
        return (
            <div>
                <Transition visible={ !this.props.loaded } animation='fade' duration={ 500 } unmountOnHide>
                    <Dimmer active>
                        <Loader size="big">
                            Loading ... { `${this.props.step} / ${this.props.totalStep}` }
                        </Loader>
                    </Dimmer>
                </Transition>
                { this.props.loaded ? this.props.children : null }
            </div>
        );
    }
}

export { AppLoaderProps, AppLoader };