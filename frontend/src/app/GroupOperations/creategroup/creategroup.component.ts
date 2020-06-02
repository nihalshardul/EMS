import { Component, OnInit } from '@angular/core';
import { ApiService } from '../../api.service';
import { UserloginService } from "../../userlogin.service";
import { Router } from '@angular/router';

@Component({
  selector: 'app-creategroup',
  templateUrl: './creategroup.component.html',
  styleUrls: ['./creategroup.component.css']
})
export class CreategroupComponent implements OnInit {

  message:boolean;

  constructor(private _api : ApiService, private value: UserloginService, private router: Router) { }

  ngOnInit(): void {
    this.value.sharedMessage.subscribe(message => this.message = message)
    if (this.message == false){
      this.router.navigate(["login"])
    }
  }

  newgroup: Object  = {} ;
  isAdded: Boolean = false;
  Confirmation: String = "New group has been Registered.";

  post = function(data){
    this.newgroup = {
      "id" : Number(data.id),
      "group_name":data.group_name,
    }
    this._api.createGroup(this.newgroup).subscribe(res => {
       console.log(res);
       this.isAdded=true;
       
    })
  }

}
