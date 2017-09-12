import * as React from 'react';
import { connect, Dispatch } from 'react-redux';
import { reduxForm, flux } from 'toolbox';

import { Stores } from 'app/core/reducers';
import * as actions from '../actions';
import { SchemaEditor as EditorBase } from '../components';
import { Schema } from '../models';


type SchemaEditorProps = {}

export function mapStateToProps( state: Stores, props: SchemaEditorProps ) {
    return {
        initialValues: {},
    };
}

export function mapDispatchToProps( dispatch: Dispatch<any>, props: SchemaEditorProps ) {
    return {
        onSubmit:      ( fields: Schema ) => flux.dispatchAsync(
            actions.Submit( fields ),
            dispatch,
        ),
        asyncValidate: ( fields: Schema, d, dd, field: string ) => flux.dispatchAsync(
            actions.Validate( { schema: fields, field } ),
            dispatch
        )
    };
}

const Form = reduxForm<Partial<Schema>, SchemaEditorProps>( {
    form:            'schemaEditor',
    asyncBlurFields: [ 'type' ],
} )( EditorBase );

const SchemaEditor: React.ComponentType<SchemaEditorProps> = connect(
    mapStateToProps,
    mapDispatchToProps
)( props => <Form { ...props } /> );

export { SchemaEditorProps, SchemaEditor };