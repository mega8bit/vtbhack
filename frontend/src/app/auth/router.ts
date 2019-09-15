import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import {RegUserComponent} from "./reg-user/reg-user.component";
import {AuthComponent} from "./auth.component";
import {LoginComponent} from "./login/login.component";

const routes: Routes = [

  {
    path: '',
    component: AuthComponent,
    children: [
      {
        path: 'reg_user',
        component: RegUserComponent,
      },
      {
        path: 'login',
        component: LoginComponent,
      },
    ],
  }

];

@NgModule({
  imports: [
    RouterModule.forChild(routes),
  ],
  providers: [
  ],
  exports: [RouterModule]
})
export class Router
{
}
