import { Component, OnInit, Input } from '@angular/core';
import { ApiService } from '../../api.service';
import { UserloginService } from "../../userlogin.service";
import { Router } from '@angular/router';

@Component({
  selector: 'app-createuser',
  templateUrl: './createuser.component.html',
  styleUrls: ['./createuser.component.css']
})
export class CreateuserComponent implements OnInit {

  message:boolean;

  constructor(private _api : ApiService, private value: UserloginService, private router: Router) { }

  ngOnInit(): void {
    this.value.sharedMessage.subscribe(message => this.message = message)
    if (this.message == false){
      this.router.navigate(["login"])
    }
  }

  newuser: Object  = {} ;
  isAdded: Boolean = false;
  Confirmation: String = "New User has been Registered.";

  @Input() islogin: number;


  post = function(data){
    this.newuser = {
      "id":Number(data.id),
      "first_name":data.first_name,
      "last_name":data.last_name,
      "email":data.email,
      "password":data.password
    }

    this._api.createUser(this.newuser).subscribe(res => {
       console.log(res);
       this.isAdded=true;
       
    })
  }



}
