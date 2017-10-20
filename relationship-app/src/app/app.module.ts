import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { RouterModule, Routes } from '@angular/router';
import { HTTP_INTERCEPTORS, HttpClientModule } from '@angular/common/http';

import { AppComponent } from './app.component';
import * as layout from './app-layout';
import { services, modules } from './shared';
import { ApiInterceptor } from './interceptors';

import * as home from './home';
import * as universe from './universe';

import 'rxjs/Rx';


const routes: Routes = [
  { path: 'home', component: home.Page, data: { toolbar: null } },
  { path: 'universes/:id', component: universe.Page, data: { toolbar: universe.Toolbar } },
  { path: '', redirectTo: '/home', pathMatch: 'full' },
  { path: '**', component: home.Page, data: { toolbar: null } }, // TODO Not Found
];

@NgModule( {
  declarations: [
    AppComponent,
    ...layout.Components,
    ...home.Components,
    ...universe.Components,
  ],
  imports: [
    HttpClientModule,
    BrowserModule,
    FormsModule,
    BrowserAnimationsModule,
    RouterModule.forRoot( routes, { enableTracing: true } ),
    ...modules
  ],
  providers: [
    ...services,
    { provide: HTTP_INTERCEPTORS, useClass: ApiInterceptor, multi: true }
  ],
  bootstrap: [ AppComponent ]
} )
export class AppModule {}
