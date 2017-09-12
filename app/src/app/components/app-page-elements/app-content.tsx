import * as React from 'react';

type AppContentProps = { className ?: string };

const AppContent: React.SFC<AppContentProps> = ( { className, children, ...props } ) => (
    <div className={ className }>
        Content ???
        { children }
    </div>
);

export { AppContentProps, AppContent };