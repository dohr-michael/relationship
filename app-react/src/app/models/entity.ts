export interface Entity {
    /**
     * Id of the entity.
     */
    id: string;
    /**
     * Universe.
     */
    universe: string;
    /**
     * The name of the entity.
     */
    name: string;
    /**
     * Schema id.
     */
    schema: string;
    /**
     * Data sets.
     */
    props: any;
}

export namespace Entity {}