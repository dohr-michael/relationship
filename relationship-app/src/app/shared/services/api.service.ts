import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Observable } from 'rxjs/Observable';
import { environment } from '../../../environments/environment';

@Injectable()
export class ApiService {
  constructor( private http: HttpClient ) {}

  private getHeaders(): HttpHeaders {
    const headersConfig = {
      'Content-Type': 'application/json',
      'Accept': 'application/json'
    };
    // TODO Bearer
    return new HttpHeaders( headersConfig );
  }

  private formatErrors( error: any ) {
    console.log( error );
    return Observable.throw( error.json() );
  }

  get<T = any>( path: string, params: URLSearchParams = new URLSearchParams() ): Observable<T> {
    return this.http.get<T>( `${environment.api_url}${path}?${params.toString()}`, { headers: this.getHeaders() } )
      .catch( this.formatErrors );
  }

  put<T = any>( path: string, body: Object = {} ): Observable<T> {
    return this.http.put<T>( `${environment.api_url}${path}`, JSON.stringify( body ), { headers: this.getHeaders() } )
      .catch( this.formatErrors );
  }

  post<T = any>( path: string, body: Object = {} ): Observable<T> {
    return this.http.post<T>( `${environment.api_url}${path}`, JSON.stringify( body ), { headers: this.getHeaders() } )
      .catch( this.formatErrors );
  }

  delete<T = any>( path ): Observable<T> {
    return this.http.delete<T>( `${environment.api_url}${path}`, { headers: this.getHeaders() } )
      .catch( this.formatErrors );
  }
}
