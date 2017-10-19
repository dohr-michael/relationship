import { AfterViewInit, Component, Input, OnInit, ViewChild } from '@angular/core';
import { MatSidenav } from '@angular/material';

@Component( {
  selector: 'app-layout',
  templateUrl: './layout.component.html',
  styleUrls: [ './layout.component.scss' ]
} )
export class LayoutComponent implements OnInit, AfterViewInit {
  @ViewChild( 'side' )
  side: MatSidenav;
  @Input()
  title: string;

  toggleSidebar = () => {
    if( this.side.opened ) {
      this.side.close();
    } else {
      this.side.open();
    }
  }

  constructor() { }

  ngOnInit() {}

  ngAfterViewInit() {}
}
