import {Component, OnInit} from '@angular/core';
import {Apollo} from "apollo-angular";
import {HttpLink} from "apollo-angular-link-http";
import {InMemoryCache} from "apollo-cache-inmemory";
import gql from "graphql-tag";

@Component({
  selector: 'app-chat',
  templateUrl: './chat.component.html',
  styleUrls: ['./chat.component.scss']
})
export class ChatComponent implements OnInit {

  // ws = new WebSocket(this.url('/char-x'));

  constructor(
  ) {

    setInterval(function (t) {
      t.message.push('Lorem ipsum dolor sit amet, consectetur adipisicing elit. Aliquam beatae ducimus mollitia, quos repellendus sit ullam. Ducimus expedita fugiat ipsa ipsum quibusdam quis sapiente similique soluta veritatis! Aut, ducimus, rem?')
    }, Math.floor(Math.random() * (10000 - 5000 + 1)) + 4000, this);
  }

  url(s) :string{
    var l = window.location;
    return ((l.protocol === "https:") ? "wss://" : "ws://") + l.hostname +  l.port + s;
  }

  message = [];

  send(m:HTMLInputElement) {
    // this.ws.send(m.value);
    this.message.push(m.value)
    m.value = "";
  }

  ngOnInit() {
    // this.ws.onmessage = (event) => {
    //   this.message.push(event)
    // };
  }

}
