import { Component, OnInit } from '@angular/core';
import { ApiService } from "../../api.service";
import { UserloginService } from "../../userlogin.service";
import { Router } from '@angular/router';

@Component({
  selector: 'app-deletegroup',
  templateUrl: './deletegroup.component.html',
  styleUrls: ['./deletegroup.component.css']
})
export class DeletegroupComponent implements OnInit {

  message:boolean;

  constructor(private _api : ApiService, private value: UserloginService, private router: Router) { }

  ngOnInit(): void {
    this.value.sharedMessage.subscribe(message => this.message = message)
    if (this.message == false){
      this.router.navigate(["login"])
    }
  }
  
  getid: any;
  isdel: Boolean = false;
  Confirmation: String = "Group Deleted Successfully.";

  id(event:any){
    this.getid = event.target.value;
  }

  delete(){
    this._api.deleteGroup(this.getid).subscribe(res => {
      console.log(res);      
      this.isdel=true;
   })
  }

}
