import { Injectable } from '@angular/core';
import { Http, Headers } from '@angular/http';
import { Observable } from 'rxjs/Observable';
import { environment } from '../../../environments/environment';

@Injectable()
export class ApiService {
  constructor( private http: Http ) {}
  
  private getHeaders(): Headers {
    const headersConfig = {
      'Content-Type': 'application/json',
      'Accept':       'application/json'
    };
    // TODO Bearer
    return new Headers( headersConfig );
  }
  
  private formatErrors( error: any ) {return Observable.throw( error.json() );}
  
  get<T = any>( path: string, params: URLSearchParams = new URLSearchParams() ): Observable<T> {
    return this.http.get( `${environment.api_url}${path}`, { headers: this.getHeaders(), search: params } )
      .catch( this.formatErrors )
      .switchMap( ( res: Response ) => Observable.fromPromise( res.json() ) );
  }
  
  put<T = any>( path: string, body: Object = {} ): Observable<T> {
    return this.http.put( `${environment.api_url}${path}`, JSON.stringify( body ), { headers: this.getHeaders() } )
      .catch( this.formatErrors )
      .switchMap( ( res: Response ) => Observable.fromPromise( res.json() ) );
  }
  
  post<T = any>( path: string, body: Object = {} ): Observable<T> {
    return this.http.post( `${environment.api_url}${path}`, JSON.stringify( body ), { headers: this.getHeaders() } )
      .catch( this.formatErrors )
      .switchMap( ( res: Response ) => Observable.fromPromise( res.json() ) );
  }
  
  delete<T = any>( path ): Observable<T> {
    return this.http.delete( `${environment.api_url}${path}`, { headers: this.getHeaders() } )
      .catch( this.formatErrors )
      .switchMap( ( res: Response ) => Observable.fromPromise( res.json() ) );
  }
}
