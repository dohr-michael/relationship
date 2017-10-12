import { NgModule } from '@angular/core';
import * as mat from '@angular/material';


const modules = [
  mat.MatAutocompleteModule,
  mat.MatButtonModule,
  mat.MatButtonToggleModule,
  mat.MatCardModule,
  mat.MatCheckboxModule,
  mat.MatChipsModule,
  mat.MatTableModule,
  mat.MatDatepickerModule,
  mat.MatDialogModule,
  mat.MatGridListModule,
  mat.MatIconModule,
  mat.MatInputModule,
  mat.MatListModule,
  mat.MatMenuModule,
  mat.MatPaginatorModule,
  mat.MatProgressBarModule,
  mat.MatProgressSpinnerModule,
  mat.MatRadioModule,
  mat.MatSelectModule,
  mat.MatSidenavModule,
  mat.MatSlideToggleModule,
  mat.MatSliderModule,
  mat.MatTabsModule,
  mat.MatToolbarModule,
  mat.MatTooltipModule
];


@NgModule( {
  imports: modules,
  exports: modules,
} )
export class MaterialModule {}
