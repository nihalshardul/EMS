import { Component, OnInit } from '@angular/core';
import { ApiService } from '../../api.service';
import { UserloginService } from "../../userlogin.service";
import { Router } from '@angular/router';

@Component({
  selector: 'app-grant-permission',
  templateUrl: './grant-permission.component.html',
  styleUrls: ['./grant-permission.component.css']
})
export class GrantPermissionComponent implements OnInit {
  message: boolean;

  constructor(private _api: ApiService, private value: UserloginService, private router: Router) { }

  ngOnInit(): void {
    this.value.sharedMessage.subscribe(message => this.message = message)
    if (this.message == false) {
      this.router.navigate(["login"])
    }
  }
  res_id: number;
  _id : number;
  per: any;
  groupres: Object = {};
  isAdded: Boolean = false;
  Confirmation: String = "Permission Granted for resource...";

  post = function (data) {

    if(data.value == "read"){
      this.per = 100
    }else if(data.value == "write"){
      this.per = 110
    }else{
      this.per = 111
    }
    this.res_id = Number(data.resource_id)
    this._id = Number(data.id)

    this.groupres = {
      "resource_id": this.res_id,
      "id": this._id,
      "permission_id": this.per,
    }
    console.log(this.per);
    console.log(this.groupres);
    this._api.createGroupPermission(this.groupres).subscribe(res => {
      console.log(res);
      this.isAdded = true;
      alert("User Added");

    })
  }


}
