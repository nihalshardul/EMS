import { Component, OnInit } from '@angular/core';
import { ApiService } from "../../api.service";
import { UserloginService } from "../../userlogin.service";
import { Router } from '@angular/router';

@Component({
  selector: 'app-updategroup',
  templateUrl: './updategroup.component.html',
  styleUrls: ['./updategroup.component.css']
})
export class UpdategroupComponent implements OnInit {
  message:boolean;

  constructor(private _api : ApiService, private value: UserloginService, private router: Router) { }

  ngOnInit(): void {
    this.value.sharedMessage.subscribe(message => this.message = message)
    if (this.message == false){
      this.router.navigate(["login"])
    }
  }
  updategroup: Object  = {} ;
  isUpdated: Boolean = false;
  Confirmation: String = "Group details has been Updated.";

  put = function(data){
    this.updategroup = {
      "id":Number(data.id),
      "group_name":data.group_name
    }

    this._api.updateGroup(this.updategroup).subscribe(res => {
       console.log(res);
       this.isUpdated=true;
       
    })
  }

}
