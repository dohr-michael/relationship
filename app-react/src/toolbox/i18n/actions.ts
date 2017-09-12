import * as flux from '../flux';

export type Payload = {
    data: { messages: { [key: string]: Object }, formats: { [key: string]: Object } },
    defaultLocale: string
};

export type Load = flux.ActionWP<Payload>;
export const Load: flux.ActionWPFn<Payload> = flux.Actions.withPayload<Payload>( 'toolbox#i18n#load' );

export type ChangeLanguage = flux.ActionWP<string>;
export const ChangeLanguage: flux.ActionWPFn<string> = flux.Actions.withPayload<string>( 'toolbox#i18n#changeLanguage' )


export type Actions = Load | ChangeLanguage;


