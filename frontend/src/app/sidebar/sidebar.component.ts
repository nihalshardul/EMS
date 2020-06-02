import { Component, OnInit } from '@angular/core';
import { IsrootService } from "../isroot.service";

@Component({
  selector: 'app-sidebar',
  templateUrl: './sidebar.component.html',
  styleUrls: ['./sidebar.component.css']
})
export class SidebarComponent implements OnInit {

  user_root:boolean;
  constructor( private _isroot: IsrootService) { }

  ngOnInit(): void {
    this._isroot.sharedMessage.subscribe(user_root => this.user_root = user_root)
  }
}
