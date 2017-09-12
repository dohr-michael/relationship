import * as React from 'react';
import { connect, Dispatch } from 'react-redux';
import { Optional, i18n } from 'toolbox';
import { AppPage as BaseAppPage, AppLoader } from 'app/components';
import { Stores } from 'app/reducers';
import * as actions from 'app/actions';
import { Universe } from '../models/universe';


type AppPageProps = {};


type StP = {
    loaded: boolean;
    totalStep: number;
    currentStep: number;
    universes?: Universe[];
    userInfo?: any;
    currentLanguage: string;
    languages: string[];
}

type DtP = {
    load: () => void;
}

function mapStateToProps( state: Stores, props: AppPageProps ): StP {
    return {
        ...state.context,
        currentLanguage: i18n.selectors.getLocale( state ),
        languages:       i18n.selectors.getLocales( state ),
    };
}

function mapDispatchToProps( dispatch: Dispatch<any>, props: AppPageProps ): DtP {
    return {
        load: () => { dispatch( actions.context.LoadContext() );},
    };
}

const Base: React.ComponentType<StP & DtP> = ( { loaded, totalStep, currentStep, load, universes, userInfo, ...props } ) => (
    <AppLoader loaded={ loaded } totalStep={ totalStep } step={ currentStep } load={ load }>
        { userInfo && universes && loaded ? <BaseAppPage { ...props } universes={ universes } userInfo={ userInfo }/> : null }
    </AppLoader>
);

const AppPage = connect( mapStateToProps, mapDispatchToProps )( Base );

export { AppPage, AppPageProps };
