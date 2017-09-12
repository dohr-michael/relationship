import * as React from 'react';
import { Field, FieldArray, InjectedFormProps, FieldArrayProps } from 'toolbox';
import { Form, Button } from 'semantic-ui-react';

import { Schema, SchemaConfiguration } from '../models';


const ConfigurationEditor: React.ComponentType<FieldArrayProps<Partial<SchemaConfiguration>>> = props => {
    const addItem = ( event: React.MouseEvent<any> ) => {
        event.preventDefault();
        props.fields.push( {} );
    };
    return (
        <Form.Group>
            { props.fields.map(
                ( name, idx, fields ) => (
                    <Form.Group key={ name }>
                        <Field.Input name={ `${name}.name` } label="Name"/>
                    </Form.Group>
                )
            )
            }
            <Button onClick={ addItem }>Add config</Button>
        </Form.Group>
    );
};


type SchemaEditorProps = {};

const SchemaEditor: React.SFC<SchemaEditorProps & InjectedFormProps<Schema>> = props => (
    <Form onSubmit={ props.handleSubmit }>
        <Field.Input name="type" label={ 'Type' } pattern="[a-z]*"/>
        <Field.TextArea name="description" label={ 'Description' }/>
        <FieldArray name="configurations" component={ ConfigurationEditor }/>
        <Button type="submit">Submit</Button>
    </Form>
);


export { SchemaEditorProps, SchemaEditor };
