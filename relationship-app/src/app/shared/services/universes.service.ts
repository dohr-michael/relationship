import {Injectable} from '@angular/core';
import {ApiService} from './api.service';
import {Observable} from 'rxjs/Observable';
import {BehaviorSubject} from 'rxjs/BehaviorSubject';
import {Universe, neo} from '../models';

@Injectable()
export class UniversesService {
  // Internal subjects
  private _universes: BehaviorSubject<neo.Node<Universe>[]> = new BehaviorSubject([]);
  private _isLoaded: BehaviorSubject<boolean> = new BehaviorSubject(false);

  // Public observable
  readonly universes = this._universes.asObservable();
  readonly isLoaded = this._isLoaded.asObservable();


  constructor(private api: ApiService) {}

  load() {
    this._isLoaded.next(false);

  }

  create(obj: Universe.Creation) {

  }

  update(obj: Universe.Update) {

  }

}
