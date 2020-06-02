import { NgModule, Component } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { LoginComponent } from './login/login.component';
import { HomeComponent } from "./home/home.component";
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


const routes: Routes = [
  {path: '', component: LoginComponent},
  {path: 'login', component: LoginComponent},
  {path: 'home' ,component: HomeComponent},
  {path: 'sidebar' , component: SidebarComponent},
  {path: 'createuser' , component: CreateuserComponent},
  {path: 'updateuser' , component: UpdateuserComponent},
  {path: 'deleteuser' , component: DeleteuserComponent},
  {path: 'searchuser' , component: SearchuserComponent},
  {path: 'creategroup' , component: CreategroupComponent},
  {path: 'updategroup' , component: UpdategroupComponent},
  {path: 'deletegroup' , component: DeletegroupComponent},
  {path: 'membership', component: MembershipComponent},
  {path: 'grouppermission', component: GrantPermissionComponent},
  {path: 'listresources', component: ListResourcesComponent},
  {path: 'createresources', component: CreateResourcesComponent},
  {path: 'searchgroup' , component: SearchgroupComponent},
  {path: 'deleteresource' , component: DeleteResourceComponent}
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
