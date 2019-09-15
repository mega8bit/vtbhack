import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { TopicComponent } from './topic/topic.component';
import {MatInputModule} from "@angular/material/input";
import {MatButtonModule} from "@angular/material/button";
import {Router} from "./router";
import {MatDatepickerModule} from "@angular/material/datepicker";
import {FormsModule, ReactiveFormsModule} from "@angular/forms";
import {MatSnackBarModule} from "@angular/material/snack-bar";
import {MatNativeDateModule} from "@angular/material/core";
import { AddQuestComponent } from './add-quest/add-quest.component';
import { ItemQuestComponent } from './item-quest/item-quest.component';
import {MatStepperModule} from "@angular/material/stepper";
import {MatFormFieldModule} from "@angular/material/form-field";



@NgModule({
  declarations: [TopicComponent, AddQuestComponent, ItemQuestComponent],
  imports: [
    CommonModule,
    MatInputModule,
    MatButtonModule,
    Router,
    MatDatepickerModule,
    FormsModule,
    ReactiveFormsModule,
    MatSnackBarModule,
    MatNativeDateModule,
    MatStepperModule,
    MatFormFieldModule,
  ]
})
export class NewModule { }
