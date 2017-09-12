//import { Reducer } from 'redux';
//
//
//export type Values<T = any> = { [field: string]: FieldValue<T> }
//
//export type FieldValue<T = any> = {
//    value: T;
//    isArray: boolean;
//    items?: State<any>[];
//};
//
//export type State<T = {}> = {
//    initialValues: Partial<T>;
//    values: Values;
//}
//
//
//
//
//function parseValue<T = any>( v: T ): FieldValue<T> {
//    const isArray = Array.isArray( v );
//    //const isPrimitive =
//    return {
//        value: v,
//    }
//}
//
//function init<T = any>( state: State<T>, a: { type: string, initialValues: T } ): State<T> {
//    const values = Object.keys( a.initialValues ).reduce( ( acc, c ) => {
//        return {
//            iv: Object.assign( acc.iv, { [c]: a.initialValues[ c ] } ),
//            v:  { ...acc.v, [c]: parseValue( a.initialValues[ c ] ) }
//        };
//    }, { iv: {} as T, v: {} as Values<T> } );
//    return {
//        initialValues: values.iv,
//        values:        values.v,
//    };
//}

