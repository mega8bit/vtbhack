import {Component, Input, OnInit} from '@angular/core';
import {Question} from "../../models/Question";
import {Topic} from "../../models/Topic";

@Component({
  selector: 'app-questions',
  templateUrl: './questions.component.html',
  styleUrls: ['./questions.component.scss']
})
export class QuestionsComponent implements OnInit {

  @Input()
  questions: Question[];

  @Input()
  topic: Topic;

  constructor() {
  }

  ngOnInit() {
  }

}
