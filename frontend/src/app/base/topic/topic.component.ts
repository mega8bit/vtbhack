import {Component, EventEmitter, Input, OnInit, Output} from '@angular/core';
import {Topic} from "../../models/Topic";

@Component({
  selector: 'app-topic',
  templateUrl: './topic.component.html',
  styleUrls: ['./topic.component.scss']
})
export class TopicComponent implements OnInit {

  @Input()
  topics: Topic[];

  @Output()
  select = new EventEmitter<Topic>();

  constructor() {
  }

  ngOnInit() {
  }

}
