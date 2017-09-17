import { Injectable } from '@angular/core';
import { ApiService } from './api.service';
import { Observable } from 'rxjs/Observable';
import { Universe, Paginate } from '../models';

@Injectable()
export class UniversesService {
  
  constructor( private api: ApiService ) {}
  
  getAll(): Observable<Paginate<Universe>> {
    return this.api.get<Paginate<Universe>>( '/universes' );
  }
  
  getOne( id: string ): Observable<Universe> {
    return this.api.get<Universe>( `/universes/${id}` );
  }
}
