import {Component, OnInit} from '@angular/core';
import {Apollo} from "apollo-angular";
import gql from "graphql-tag";
import {map} from "rxjs/operators";
import {FormBuilder, FormGroup, Validators} from "@angular/forms";
import {MatSnackBar} from "@angular/material/snack-bar";
import {Router} from "@angular/router";

@Component({
  selector: 'app-topic',
  templateUrl: './topic.component.html',
  styleUrls: ['./topic.component.scss']
})
export class TopicComponent implements OnInit {

  isLinear = false;
  firstFormGroup: FormGroup;
  secondFormGroup: FormGroup;

  item = [];

  constructor(
    private _fb: FormBuilder,
    private apollo: Apollo,
    private _snackBar: MatSnackBar,
    private router: Router,
    private _formBuilder: FormBuilder
    ,
  ) {
  }

  topicId: number = null;

    create() {
      if (this.firstFormGroup.invalid) {
        return
      }

      const {title,
        type,
        start,
        end,
        main_user,} = this.firstFormGroup.value;

      this.apollo
      .query<{ createTopic: number }>({
          query: gql`
              mutation CreateTopic(
                  $title:String!,
                  $typeId:Int!,
                  $startDateTime:String!,
                  $endDateTime:String!,
                  $chairmanId:Int!
              ) {
                  createTopic(
                      title:$title,
                      typeId:$typeId,
                      startDateTime:$startDateTime,
                      endDateTime:$endDateTime,
                      chairmanId:$chairmanId
                  )
              }
          `,
        variables: {
          title: title,
          typeId: type,
          startDateTime: start,
          endDateTime: end,
          chairmanId: main_user,
        },
        fetchPolicy: "no-cache",
      })
      .pipe(
        map(({data}) => data.createTopic)
      )
      .subscribe(topicId => {
        this.topicId = topicId
      }, err => {
        this._snackBar.open(err.message, null, {
          duration: 4500,
        });
      });
    }

    add(v: string) {
      if (this.topicId === null) {
        return
      }
      this.item.push(v);
      this.secondFormGroup.reset();

      this.apollo.mutate({
          mutation: gql`
              mutation CreateQuestion(
                  $title:String!,
                  $topicId:Int!
              ) {
                  createQuestion(
                      title:$title,
                      topicId:$topicId
                  )
              }
          `
          ,
        variables: {
          title: v,
          topicId: this.topicId,
        }
      })
      .subscribe({
        error : (err) => {
          this._snackBar.open(err.message, null, {
          duration: 4500,
        });
        }
      });

    }

  ngOnInit() {
    this.firstFormGroup = this._formBuilder.group({
      title: ['', Validators.required],
      type: ['', Validators.required],
      start: ['', Validators.required],
      end: ['', Validators.required],
      main_user: ['', Validators.required],
    });
    this.secondFormGroup = this._formBuilder.group({
      title: ['', Validators.required]
    });
  }

}
