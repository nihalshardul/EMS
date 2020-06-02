import { Component, OnInit } from '@angular/core';
import { ApiService } from "../../api.service";
import { UserloginService } from "../../userlogin.service";
import { Router } from '@angular/router';
import { User } from '../../models/user';

@Component({
  selector: 'app-searchuser',
  templateUrl: './searchuser.component.html',
  styleUrls: ['./searchuser.component.css']
})
export class SearchuserComponent implements OnInit {

  message: boolean;

  constructor(private _api: ApiService, private value: UserloginService, private router: Router) { }

  ngOnInit(): void {
    this.value.sharedMessage.subscribe(message => this.message = message)
    if (this.message == false) {
      this.router.navigate(["login"])
    }
  }
  byid = false;
  users: any;
  item: any[];
  u_id : any;
  u_f_name :any;
  u_l_name : any;
  arr: Array<User> = [];

  getall() {           // for endpoint 
    this._api.getUsers().subscribe(res => {
      this.users = res["data"];
      console.log("Response", res);
      console.log(this.users)     //  List of object [{Object},{Object}]
      
      for (let index = 0; index < (this.users).length; index++) {
        let userObj = new User();     // for storing objects
        userObj.id = this.users[index][0]
        userObj.first_name = this.users[index][1]
        userObj.last_name = this.users[index][2]
        console.log(userObj)
        this.arr.push(userObj)
      }
      console.log("Array => ",this.arr)


    });
  }
  getid: any;
  id(event: any) {
    this.getid = event.target.value;
  }

  getbyid() {
    this.byid = true;
    console.log(this.getid);
    this._api.getUserById(this.getid).subscribe(res => {
      this.item = res["data"]
      this.u_id = this.item["id"]
      this.u_f_name = this.item["first_name"]
      this.u_l_name = this.item["last_name"]
      console.log("Response", res);
      console.log("Items ", this.item);
      // console.log(this.users)


    });
  }

}
