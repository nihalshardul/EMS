import { Component } from '@angular/core';
import { ApiService } from './api.service';
import { UserloginService } from './userlogin.service';


@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = 'ACL';

  message:boolean;
  constructor( private value: UserloginService) { }

  ngOnInit(): void {
    this.value.sharedMessage.subscribe(message => this.message = message)
  }

}
