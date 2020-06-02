import { Injectable } from '@angular/core';
import { BehaviorSubject } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class IsrootService {

  private id_of_user;
  private message = new BehaviorSubject(false);
  sharedMessage = this.message.asObservable();

  constructor() { }

  isRoot(message: boolean) {
    this.message.next(message)
  }

  public setValue(id_of_user){
    this.id_of_user = id_of_user;
  }

  public getValue(){
    return this.id_of_user;
  }

}
