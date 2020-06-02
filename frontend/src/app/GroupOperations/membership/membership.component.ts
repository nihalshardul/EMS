import { Component, OnInit } from '@angular/core';
import { ApiService } from '../../api.service';
import { UserloginService } from "../../userlogin.service";
import { Router } from '@angular/router';

@Component({
  selector: 'app-membership',
  templateUrl: './membership.component.html',
  styleUrls: ['./membership.component.css']
})
export class MembershipComponent implements OnInit {


  message:boolean;

  constructor(private _api : ApiService, private value: UserloginService, private router: Router) { }

  ngOnInit(): void {
    this.value.sharedMessage.subscribe(message => this.message = message)
    if (this.message == false){
      this.router.navigate(["login"])
    }
  }

  newusergroup: Object  = {} ;
  isAdded: Boolean = false;
  Confirmation: String = "User Added to Group";

  post = function(data){
    this.newusergroup = {
      "id":Number(data.id),
      "user_id":Number(data.user_id)
    }

    this._api.addUserGroup(this.newusergroup).subscribe(res => {
       console.log(res);
       this.isAdded=true;
       alert("User Added");
       
    })
  }


}
