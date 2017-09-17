import { NgModule } from '@angular/core';
import { MaterialModule as BaseMaterialModule } from '@angular/material';


@NgModule( {
  imports: [ BaseMaterialModule ],
  exports: [ BaseMaterialModule ],
} )
export class MaterialModule {}
