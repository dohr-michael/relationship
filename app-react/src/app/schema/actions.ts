import { flux, FormErrors } from 'toolbox';
import { Schema } from './models';


export type Submit = flux.PromiseAction<Schema, FormErrors>;
export const Submit = flux.Actions.promise<Schema, FormErrors>( 'schema#submit' );

export type Submitted = flux.Action;
export const Submitted = flux.Actions.ofType( 'schema#submitted' );


export type ValidatePayload = {
    schema: Schema;
    field: string;
};
export type Validate = flux.PromiseAction<ValidatePayload, FormErrors>;
export const Validate = flux.Actions.promise<ValidatePayload, FormErrors>( 'schema#validate' );

export type Actions = Submit | Validate;
