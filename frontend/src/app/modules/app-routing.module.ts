import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { HomePageComponent } from '../components/home-page/home-page.component';
import { ReportsPageComponent } from '../components/reports-page/reports-page.component';

const routes: Routes = [
  {
    path: 'home',
    component: HomePageComponent,
  },
  { path: 'home/:id', component: HomePageComponent },
  { path: 'reports', component: ReportsPageComponent },
  { path: '', redirectTo: '/home', pathMatch: 'full' },
  { path: '**', redirectTo: '/home' },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule],
})
export class AppRoutingModule {}
