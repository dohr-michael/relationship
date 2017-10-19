import { Component, OnInit } from '@angular/core';
import { UniversesService } from './shared/services/universes.service';

@Component( {
  selector: 'app-root',
  template: '<app-layout [title]="title"></app-layout>',
  styleUrls: [ './app.component.scss' ]
} )
export class AppComponent implements OnInit {
  title = 'Relationship App';


  constructor( private universesService: UniversesService ) {}

  ngOnInit(): void {
    this.universesService.load();
  }
}
