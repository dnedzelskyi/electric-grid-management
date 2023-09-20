import { Component, Inject } from '@angular/core';

const CONSTANTS = {
  Title: 'Electric Grid Management',
  SideNavMenuItems: [
    { path: '/home', icon: 'home', title: 'Home' },
    { path: '/reports', icon: 'description', title: 'Reports' },
  ],
};

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss'],
  providers: [{ provide: 'CONSTANTS', useValue: CONSTANTS }],
})
export class AppComponent {
  constructor(@Inject('CONSTANTS') public Constants: any) {}
}
