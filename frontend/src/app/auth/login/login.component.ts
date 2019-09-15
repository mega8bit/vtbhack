import {Component, OnInit} from '@angular/core';
import {FormBuilder, FormGroup, Validators} from "@angular/forms";
import gql from "graphql-tag";
import {map} from "rxjs/operators";
import {Apollo} from "apollo-angular";
import {User} from "../../models/User";
import {MatSnackBar} from "@angular/material/snack-bar";
import {Router} from "@angular/router";

@Component({
  selector: 'app-login-user',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss']
})
export class LoginComponent implements OnInit {

  loginForm: FormGroup = this._fb.group({
    email: ['', [
      Validators.required,
      Validators.email,
    ]],
    password: ['', Validators.required],
  });

  constructor(
    private _fb: FormBuilder,
    private apollo: Apollo,
    private _snackBar: MatSnackBar,
    private router: Router
  ) {
  }

  get email() {
    return this.loginForm.get("email")
  }

  get password() {
    return this.loginForm.get("password")
  }

  get ready(): boolean {
    return this.loginForm.valid
  }

  getErrorMessage() {
    return this.email.hasError('required') ? 'Обязательное поле' :
      this.email.hasError('email') ? 'Некорректный емайл' :
        '';
  }

    login() {
      if (this.loginForm.invalid) {
        return
      }

      const {email, password} = this.loginForm.value;

      this.apollo
      .query<{ user: User }>({
          query: gql`
              query Q($email: String!, $pwd: String!) {
                  user(email: $email, password: $pwd) {
                      token
                  }
              }
          `, variables: {email: email, pwd: password}, fetchPolicy: "no-cache",
      })
      .pipe(
        map(({data}) => data.user.token)
      )
      .subscribe(token => {
        if (token == "") {
          localStorage.removeItem("auth_token");
          this._snackBar.open("Ошибка авторизаиции", null, {
            duration: 4500,
          });
        } else {
          localStorage.setItem("auth_token", token);
          this.router.navigate(['/topic']);
        }
      }, err => {
        this._snackBar.open(err.message, null, {
          duration: 4500,
        });
      });
    }

  ngOnInit() {
  }

}
