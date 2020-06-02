import { Component, OnInit } from '@angular/core';
import {  ApiService } from "../../api.service";
import { UserloginService } from "../../userlogin.service";
import { Router } from '@angular/router';
import { Group } from '../../models/group';

@Component({
  selector: 'app-searchgroup',
  templateUrl: './searchgroup.component.html',
  styleUrls: ['./searchgroup.component.css']
})
export class SearchgroupComponent implements OnInit {

  message:boolean;

  constructor(private _api : ApiService, private value: UserloginService, private router: Router) { }

  ngOnInit(): void {
    this.value.sharedMessage.subscribe(message => this.message = message)
    if (this.message == false){
      this.router.navigate(["login"])
    }
  }
  byid=false;
  groups : any;
  item : any [];
  arr: Array<Group> = [];

  getall(){           // for endpoint 
    this._api.getGroup().subscribe(res => {
    this.groups = res["data"]; 
    console.log("Response", res);
    console.log(this.groups)

    for (let index = 0; index < (this.groups).length; index++) {
      let userObj = new Group();     // for storing objects
      userObj.id = this.groups[index][0]
      userObj.groups_name = this.groups[index][1]
      console.log(userObj)
      this.arr.push(userObj)
    }
    console.log("Array => ",this.arr)

    });
  }


  getid: any;
  id(event:any){
    this.getid = event.target.value;
  }
  g_id : any;
  g_name : any;

  getbyid(){
    this.byid = true;
    console.log(this.getid);
    this._api.getGroupById(this.getid).subscribe( res => {
      this.item = res["data"] 
      this.g_id = this.item["id"]
      this.g_name = this.item["group_name"]
      console.log("Response", res);
      console.log("Items ",this.item);
      
    });
  }

}
