import * as React from 'react';
import { Field as BaseField, WrappedFieldInputProps, WrappedFieldMetaProps } from 'redux-form';
import { Form, FormSelectProps, FormInputProps, FormCheckboxProps, FormDropdownProps, FormTextAreaProps } from 'semantic-ui-react';


type SemanticFieldProps = {
    [key: string]: any;
    as?: React.ComponentType<any>;
    label?: React.ReactNode;
    canAsync?: boolean;
}

type AdditionalProps = {
    onChangeAdditional?: ( e: React.SyntheticEvent<any>, d: any ) => void;
}

type SemanticFormFieldComponentProps = SemanticFieldProps & {
    input: WrappedFieldInputProps;
    meta: WrappedFieldMetaProps;
};

const SemanticFormFieldComponent: React.SFC<SemanticFormFieldComponentProps & AdditionalProps> = ( { input, label, meta, as: As = Form.Input, onChangeAdditional, canAsync, ...props } ) => {
    const { touched, error, warning, asyncValidating } = meta;
    
    function handleChange( e, data) {
        if( onChangeAdditional ) onChangeAdditional( e, data );
        return input.onChange( data.value );
    }
    
    const additionalProps: any = {};
    if( canAsync && asyncValidating ) additionalProps.loading = true;
    let onError = touched && !!error;
    const field = (
        <Form.Field control={ As }
                    { ...input }
                    value={ input.value }
                    label={ label }
                    { ...props }
                    { ...additionalProps }
                    onChange={ handleChange }
                    error={ onError }/>
    );
    return field;
};

const defaultProps = { component: SemanticFormFieldComponent, canAsync: true, };

class IntField<T = {}> extends BaseField<FormInputProps & T & AdditionalProps> {static defaultProps = { ...defaultProps, parse: ( value: any, field: string ): any => parseInt( value ), };}
class FloatField<T = {}> extends BaseField<FormInputProps & T & AdditionalProps> {static defaultProps = { ...defaultProps, parse: ( value: any, field: string ): any => parseFloat( value ), };}
class SelectField<T = {}> extends BaseField<FormSelectProps & T & AdditionalProps> {static defaultProps = { ...defaultProps, as: Form.Select, };}
class CheckboxField<T = {}> extends BaseField<FormCheckboxProps & T & AdditionalProps> {static defaultProps = { ...defaultProps, canAsync: false, as: Form.Checkbox, };}
class DropdownField<T = {}> extends BaseField<FormDropdownProps & T & AdditionalProps> {static defaultProps = { ...defaultProps, as: Form.Dropdown, };}
class TextAreaField<T = {}> extends BaseField<FormTextAreaProps & T & AdditionalProps> {static defaultProps = { ...defaultProps, canAsync: false, as: Form.TextArea, };}

class Field<T = {}> extends BaseField<FormInputProps & T & AdditionalProps> {
    static defaultProps = { ...defaultProps };
    static Int = IntField;
    static Float = FloatField;
    static Checkbox = CheckboxField;
    static Dropdown = DropdownField;
    static TextArea = TextAreaField;
    static Select = SelectField;
    static Input = Field;
    
}

export { SemanticFieldProps, IntField, FloatField, SelectField, CheckboxField, DropdownField, TextAreaField, Field, SemanticFormFieldComponentProps, SemanticFormFieldComponent };



