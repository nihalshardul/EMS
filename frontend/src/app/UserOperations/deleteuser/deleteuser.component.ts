import { Component, OnInit } from '@angular/core';
import { ApiService } from "../../api.service";
import { UserloginService } from "../../userlogin.service";
import { Router } from '@angular/router';

@Component({
  selector: 'app-deleteuser',
  templateUrl: './deleteuser.component.html',
  styleUrls: ['./deleteuser.component.css']
})
export class DeleteuserComponent implements OnInit {

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
  Confirmation: String = "User Deleted Successfully.";

  id(event:any){
    this.getid = event.target.value;
  }

  delete(){
    this._api.deleteUser(this.getid).subscribe(res => {
      console.log(res);      
      this.isdel=true;
   })
  }


}
