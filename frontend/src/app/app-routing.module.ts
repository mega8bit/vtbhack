import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import {StdGuard} from "./std.guard";


const routes: Routes = [

  {
    path: 'auth',
    loadChildren: () => import('./auth/auth.module')
      .then(mod => mod.AuthModule),
  },

  {
    path: 'topic',
    loadChildren: () => import('./base/base.module')
      .then(mod => mod.BaseModule),
    canActivate: [StdGuard],
  },

  {
    path: 'new',
    loadChildren: () => import('./new/new.module')
      .then(mod => mod.NewModule),
    canActivate: [StdGuard],
  },

];


@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
