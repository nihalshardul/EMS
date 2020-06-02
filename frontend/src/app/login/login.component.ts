import { Component, OnInit, Input, Output, EventEmitter } from '@angular/core';
import { ApiService } from '../api.service';
import { Router } from '@angular/router';
import { UserloginService } from '../userlogin.service';
import { IsrootService } from "../isroot.service";


@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {

  constructor(private _isroot: IsrootService, private _api: ApiService, private router: Router, private value: UserloginService) { }
  message: boolean;
  user_root: boolean;

  ngOnInit(): void {
    this.value.sharedMessage.subscribe(islogin => this.message = this.message)
    this._isroot.sharedMessage.subscribe(rt => this.user_root = this.user_root)
  }
  root_user: Number;
  id: any;
  password: String;
  onlogin: number = 0;
  item: any;
  is_id: any;
  is_pass: any;
  stat: any;

  validation = function (data) {
    this.id = data.id,
      this.password = data.password

    // if (this.email == 'admin' && this.password == 'admin'){
    //   this.root_user = 1      
    //   this.value.nextMessage(true)
    //   this.router.navigate(["home"])
    // }else{
    //   alert("Invalid Credentials...")
    // }
    console.log(this.email);
    this._api.getUserById(this.id).subscribe(res => {
      this.item = res["message"]
      if (this.item == undefined) {
        console.log("True");
        this.item = res["data"]
        this.is_id = this.item["id"]
        this.is_pass = this.item["password"]
        this.root_user = this.item[" is_user_root "]
        console.log("Response", res);
        console.log("Items ", this.item);
        console.log("User Id", this.is_id);
        console.log("Pass", this.is_pass);
        console.log("Root", this.root_user);
        if (this.password == this.is_pass) {
          this.value.nextMessage(true)
          this._isroot.setValue(this.is_id)
          if (this.root_user == 1) {
            this._isroot.isRoot(true)
          }
          this.router.navigate(["home"])

        } else {
          alert("Invalid Credentials...");
        }
      } else {
        alert("Invalid Credentials...");
      }
    });
  }






}
