import { JsonSchema } from 'toolbox/data';


export interface SchemaConfiguration {
    type: string;
    value: any;
}

export interface Schema {
    /**
     * Hash of the schema (version)
     */
    id: string;
    /**
     * Category of the schema (e.g. Person)
     */
    type: string;
    /**
     * Current description of schema.
     */
    description: string;
    /**
     * Map of configurations.
     * Can be representation config (e.g. how to display entities in list view ...)
     */
    configurations: SchemaConfiguration[] ;
    /**
     * The definition of the model.
     */
    definition: JsonSchema
}

export namespace Schema {

}