import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';
import {RegUserComponent} from './reg-user/reg-user.component';
import {Router} from "./router";
import {AuthComponent} from "./auth.component";
import {MatSelectModule} from "@angular/material/select";
import {MatInputModule} from "@angular/material/input";
import {MatButtonModule} from "@angular/material/button";
import {LoginComponent} from "./login/login.component";
import {FormsModule, ReactiveFormsModule} from "@angular/forms";
import {MatSnackBarModule} from "@angular/material/snack-bar";
import {RouterModule} from "@angular/router";


@NgModule({
  declarations: [
    RegUserComponent,
    AuthComponent,
    LoginComponent,
  ],
  imports: [
    CommonModule,
    Router,
    RouterModule,
    MatSelectModule,
    MatInputModule,
    MatButtonModule,
    FormsModule,
    ReactiveFormsModule,
    MatSnackBarModule,
  ]
})
export class AuthModule {
}
