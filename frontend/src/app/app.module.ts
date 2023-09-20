import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './modules/app-routing.module';
import { AppComponent } from './app.component';
import { NgxGraphModule } from '@swimlane/ngx-graph';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { MaterialModule } from './modules/material.module';
import { HomePageComponent } from './components/home-page/home-page.component';
import { GridMapComponent } from './components/grid-map/grid-map.component';
import { GridNodeDetailsComponent } from './components/grid-node-details/grid-node-details.component';
import { GridService } from './services/grid/grid.service';
import { FormsModule } from '@angular/forms';
import { ReportsPageComponent } from './components/reports-page/reports-page.component';

@NgModule({
  declarations: [
    AppComponent,
    HomePageComponent,
    GridMapComponent,
    GridNodeDetailsComponent,
    ReportsPageComponent,
  ],
  imports: [
    BrowserModule,
    FormsModule,
    AppRoutingModule,
    NgxGraphModule,
    BrowserAnimationsModule,
    MaterialModule,
  ],
  providers: [GridService],
  bootstrap: [AppComponent],
})
export class AppModule {}
