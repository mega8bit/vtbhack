import {NgModule} from '@angular/core';
import {RouterModule, Routes} from '@angular/router';
import {BaseComponent} from "./base.component";
import {TopicsResolver} from "./resolvers";

const routes: Routes = [

  {
    path: '',
    component: BaseComponent,
    resolve: {
      topics: TopicsResolver,
    }
  },

];

@NgModule({
  imports: [
    RouterModule.forChild(routes),
  ],
  providers: [
    TopicsResolver,
  ],
  exports: [RouterModule]
})
export class Router {
}
