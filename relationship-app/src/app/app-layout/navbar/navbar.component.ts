import { Component, ComponentFactoryResolver, Input, OnInit, ViewChild, ViewContainerRef } from '@angular/core';
import { ActivatedRoute } from '@angular/router';

@Component( {
  selector: 'app-navbar',
  templateUrl: './navbar.component.html',
  styleUrls: [ './navbar.component.scss' ]
} )
export class NavbarComponent implements OnInit {
  @ViewChild( 'toolbar', { read: ViewContainerRef } )
  toolbar: ViewContainerRef;
  @Input()
  onNavClick: () => void;
  @Input()
  title: string;

  constructor( private route: ActivatedRoute,
               private _componentFactoryResolver: ComponentFactoryResolver ) {}

  ngOnInit() {
    this.route.data.subscribe( data => {
      const getPart = ( name: string ) => !!data && !!data[ name ] ? this._componentFactoryResolver.resolveComponentFactory( data[ name ] ) : null;
      const toolbar = getPart( 'toolbar' );
      if( toolbar ) {
        this.toolbar.createComponent( toolbar );
      } else {
        this.toolbar.clear();
      }
    } );
  }

}
