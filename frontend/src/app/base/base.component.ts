import {Component, OnInit} from "@angular/core";
import {ActivatedRoute} from "@angular/router";
import {Topic} from "../models/Topic";
import {Question} from "../models/Question";
import gql from "graphql-tag";
import {map} from "rxjs/operators";
import {Apollo} from "apollo-angular";

@Component({
  selector: 'app-base',
  templateUrl: './base.html',
  styleUrls: ['./base.scss'],
})
export class BaseComponent implements OnInit {
  selectedQuestions: Question[];
  selectedTopic: Topic = null;
  topics: Topic[];

  constructor(
    private route: ActivatedRoute,
    private apollo: Apollo,
  ) {
  }

  get getQuestions(): Question[] {
    if (!this.selectedTopic) {
      return []
    }
    return this.selectedQuestions
  }

  ngOnInit(): void {
    this.route.data
      .subscribe(({topics}: { topics: Topic[] }) => {
        this.topics = topics
      })
  }

    selectTopic(topic: Topic) {
    this.selectedTopic = topic;
      this.apollo
      .query<{ questions: Question[] }>({
          query: gql`
              query Q($topic: Int!) {
                  questions(topic_id: $topic) {
                      id
                      title
                      status
                      topicId
                  }
              }
          `, variables: {topic: topic.id}
      })
      .pipe(
        map(({data}) => data.questions)
      )
      .subscribe(questions => {
        this.selectedQuestions = questions
      });
    }
}
