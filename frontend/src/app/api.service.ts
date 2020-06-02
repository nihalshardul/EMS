import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { User } from "./models/user";
import { Group } from "./models/group";
import { Usergroup } from "./models/usergroup";
import { Grouppermission } from "./models/grouppermission";
import { Userpermission } from "./models/userpermission";
import { Resources } from "./models/resources";

@Injectable({
  providedIn: 'root'
})
export class ApiService {

  private apiURL: string = 'webapi/v1/';

  constructor(private httpClient: HttpClient) { }

  // services for user
    public createUser(user: User) {
      return this.httpClient.post(`${this.apiURL}user`, user);

    }

  public updateUser(user: User) {
    return this.httpClient.put(`${this.apiURL}user/${user.id}`, user);
  }

  public deleteUser(id: number) {
    return this.httpClient.delete(`${this.apiURL}user/${id}`);
  }

  public getUserById(id: number) {
    return this.httpClient.get(`${this.apiURL}/user/${id}`);   //endpoint
  }

  public getUsers() {
    return this.httpClient.get(`${this.apiURL}user`);
  }

  // public loginUser(email: String , password: String){
  //   return this.httpClient.get(`${this.apiURL}/user/${user.email}/${user.password}`);   //for login we need username and password
  // }

  // services for group
  public createGroup(group: Group) {
    return this.httpClient.post(`${this.apiURL}groups`, group);

  }

  public updateGroup(group: Group) {
    return this.httpClient.put(`${this.apiURL}groups/${group.id}`, group);
  }

  public deleteGroup(id: number) {
    return this.httpClient.delete(`${this.apiURL}groups/${id}`);
  }

  public getGroupById(id: number) {
    return this.httpClient.get(`${this.apiURL}/groups/${id}`);   //endpoint
  }

  public getGroup() {
    return this.httpClient.get(`${this.apiURL}groups`);
  }

  // services for usergroup

  public addUserGroup(newusergroup: Usergroup) {
    return this.httpClient.post(`${this.apiURL}usergroup`, newusergroup);
  }

  public updateUserGroup(group: Usergroup) {
    return this.httpClient.put(`${this.apiURL}usergroup/${group.id}`, group);
  }

  public deleteUserGroup(id: number) {
    return this.httpClient.delete(`${this.apiURL}usergroup/${id}`);
  }

  public getUserGroupById(id: number) {
    return this.httpClient.get(`${this.apiURL}/usergroup/${id}`);   //endpoint
  }

  public getUserGroup() {
    return this.httpClient.get(`${this.apiURL}usergroup`);
  }

  // services for userpermission

  public createUserpermission(data: Userpermission) {
    return this.httpClient.post(`${this.apiURL}userpermission`, data);
  }

  public updateUserpermission(data: Userpermission) {
    return this.httpClient.put(`${this.apiURL}userpermission/${data.resource_id}`, data);
  }

  public deleteUserpermission(id: number) {
    return this.httpClient.delete(`${this.apiURL}userpermission/${id}`);
  }

  public getUserpermissionById(id: number) {
    return this.httpClient.get(`${this.apiURL}/userpermission/${id}`);   //endpoint
  }

  public getUserpermission() {
    return this.httpClient.get(`${this.apiURL}userpermission`);
  }



  // services for grouppermission

  public createGroupPermission(data: Grouppermission) {
    return this.httpClient.post(`${this.apiURL}grouppermission`, data);
  }

  public updateGrouppermission(data: Grouppermission) {
    return this.httpClient.put(`${this.apiURL}grouppermission/${data.id}`, data);
  }

  public deleteGrouppermission(id: number) {
    return this.httpClient.delete(`${this.apiURL}grouppermission/${id}`);
  }

  public getGrouppermissionById(id: number) {
    return this.httpClient.get(`${this.apiURL}/grouppermission/${id}`);   //endpoint
  }

  public getGrouppermission() {
    return this.httpClient.get(`${this.apiURL}grouppermission`);
  }

  // services for resources
  public createResources(data: Resources) {
    return this.httpClient.post(`${this.apiURL}resources`, data);
  }

  public updateResources(data: Resources) {
    return this.httpClient.put(`${this.apiURL}resources/${data.id}`, data);
  }

  public deleteResources(id: number) {
    return this.httpClient.delete(`${this.apiURL}resources/${id}`);
  }

  public getResourcesById(id: number) {
    return this.httpClient.get(`${this.apiURL}/resources/${id}`);   //endpoint
  }

  public getResources() {
    return this.httpClient.get(`${this.apiURL}resources`);
  }


}
