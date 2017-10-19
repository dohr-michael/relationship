import { Component, OnInit } from '@angular/core';
import { UniversesService } from '../../shared/services/universes.service';
import { Universe, neo } from '../../shared/models';

@Component( {
  selector: 'app-sidebar',
  templateUrl: './sidebar.component.html',
  styleUrls: [ './sidebar.component.scss' ]
} )
export class SidebarComponent implements OnInit {
  universes: neo.Node<Universe>[] = [];

  constructor( private universesService: UniversesService ) { }

  ngOnInit() {
    this.universesService.universes.subscribe( universes => {
      this.universes = universes;
    } );
  }

}
