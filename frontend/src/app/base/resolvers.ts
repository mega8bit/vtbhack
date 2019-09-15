import {Injectable} from "@angular/core";
import {ActivatedRouteSnapshot, Resolve, RouterStateSnapshot} from "@angular/router";
import {Observable} from "rxjs";
import {Topic} from "../models/Topic";
import {Apollo} from "apollo-angular";
import gql from "graphql-tag";
import {map} from "rxjs/operators";

@Injectable()
export class TopicsResolver implements Resolve<Topic[]> {
  constructor(private apollo: Apollo) {
  }

    resolve(
      route: ActivatedRouteSnapshot,
      state: RouterStateSnapshot
    ): Observable<Topic[]> {
        return this.apollo
        .query<{ topics: Topic[] }>({
            query: gql`{
                topics {
                    id
                    status
                    title
                    typeId
                }
            }`
        })
        .pipe(
          map(({data}) => data.topics)
        );
    }
}
