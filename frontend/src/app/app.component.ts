import { Component } from '@angular/core';
import { HelloServiceService } from './services/hello-service.service';
import { GridNode } from './pb/grid_pb';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss'],
})
export class AppComponent {
  title = 'frontend';
  greeting: GridNode[] = [];

  constructor(private helloService: HelloServiceService) {}

  async handleClick(event: Event) {
    this.greeting = await this.helloService.getGreeting();
  }
}
