import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { HelloServiceService } from './services/hello-service.service';
import { GridGraphComponent } from './components/grid-graph/grid-graph.component';

@NgModule({
  declarations: [AppComponent, GridGraphComponent],
  imports: [BrowserModule, AppRoutingModule],
  providers: [HelloServiceService],
  bootstrap: [AppComponent],
})
export class AppModule {}
