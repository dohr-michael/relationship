import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';

import { AppComponent } from './app.component';
import { services } from './shared';
import { MaterialModule } from './material.module';
import 'rxjs/Rx';


@NgModule( {
  declarations: [
    AppComponent
  ],
  imports:      [
    BrowserModule,
    FormsModule,
    BrowserAnimationsModule,
    MaterialModule
  ],
  providers:    [ ...services ],
  bootstrap:    [ AppComponent ]
} )
export class AppModule {}
