import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';
import {TopicComponent} from './topic/topic.component';
import {MatToolbarModule} from "@angular/material/toolbar";
import {Router} from "./router";
import {QuestionsComponent} from './questions/questions.component';
import {ChatComponent} from './chat/chat.component';
import {BaseComponent} from "./base.component";
import {MatIconModule} from "@angular/material/icon";
import {MatRippleModule} from "@angular/material/core";
import {MatInputModule} from "@angular/material/input";
import {MatButtonModule} from "@angular/material/button";


@NgModule({
  declarations: [
    TopicComponent, QuestionsComponent, ChatComponent,
    BaseComponent,
  ],
  imports: [
    Router,
    CommonModule,
    MatToolbarModule,
    MatIconModule,
    MatRippleModule,
    MatInputModule,
    MatButtonModule
  ]
})
export class BaseModule {
}
