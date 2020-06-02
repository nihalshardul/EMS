import { Injectable } from '@angular/core';
import { BehaviorSubject } from 'rxjs';


@Injectable({
  providedIn: 'root'
})
export class UserloginService {
    private message = new BehaviorSubject(false);
    sharedMessage = this.message.asObservable();
  
    constructor() { }
  
    nextMessage(message: boolean) {
      this.message.next(message)
    }
    
}
