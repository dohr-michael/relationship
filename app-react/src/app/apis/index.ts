import { Observable } from 'rxjs';
import { Universe } from 'app/models';

/**
 * Load user information.
 */
export const loadUserInfo = (): Observable<any> => Observable.of( {
    name:  'Michael DOHR',
    email: 'dohr.michael@gmail.com'
} ).delay( 200 );

/**
 * Load all data universe.
 */
export const loadUniverses = (): Observable<Universe[]> => Observable.of( [
    { id: '1', name: 'L.A. By Night (Vampire)' },
    { id: '2', name: 'La règle des deux (Star Wars)' },
] ).delay( 400 );
