import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import {RegUserComponent} from "./reg-user/reg-user.component";
import {AuthComponent} from "./auth.component";
import {LoginComponent} from "./login/login.component";
import {TopicComponent} from "./topic/topic.component";

const routes: Routes = [

  {
    path: 'topic',
    component: TopicComponent,
  },

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
