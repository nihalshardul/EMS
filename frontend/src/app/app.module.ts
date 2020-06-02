import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { HttpClientModule, HttpClient } from "@angular/common/http";
import { CommonModule } from '@angular/common';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { LoginComponent } from './login/login.component';
import { HomeComponent } from './home/home.component';
import { FormsModule } from '@angular/forms';
import { CreateuserComponent } from './UserOperations/createuser/createuser.component';
import { UpdateuserComponent } from './UserOperations/updateuser/updateuser.component';
import { DeleteuserComponent } from './UserOperations/deleteuser/deleteuser.component';
import { SearchuserComponent } from './UserOperations/searchuser/searchuser.component';
import { SearchgroupComponent } from './GroupOperations/searchgroup/searchgroup.component';
import { DeletegroupComponent } from './GroupOperations/deletegroup/deletegroup.component';
import { UpdategroupComponent } from './GroupOperations/updategroup/updategroup.component';
import { CreategroupComponent } from './GroupOperations/creategroup/creategroup.component';
import { SidebarComponent } from './sidebar/sidebar.component';
import { MembershipComponent } from './GroupOperations/membership/membership.component';
import { GrantPermissionComponent } from './GroupOperations/grant-permission/grant-permission.component';
import { ListResourcesComponent } from './ResourceOperations/list-resources/list-resources.component';
import { CreateResourcesComponent } from './ResourceOperations/create-resources/create-resources.component';
import { DeleteResourceComponent } from './ResourceOperations/delete-resource/delete-resource.component';


@NgModule({
  declarations: [
    AppComponent,
    LoginComponent,
    HomeComponent,
    CreateuserComponent,
    UpdateuserComponent,
    DeleteuserComponent,
    SearchuserComponent,
    SearchgroupComponent,
    DeletegroupComponent,
    UpdategroupComponent,
    CreategroupComponent,
    SidebarComponent,
    MembershipComponent,
    GrantPermissionComponent,
    ListResourcesComponent,
    CreateResourcesComponent,
    DeleteResourceComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    HttpClientModule,
    FormsModule,
    CommonModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
