import { Injectable } from '@angular/core';
import { Subject } from "rxjs";

@Injectable({
  providedIn: 'root'
})
export class OpenRegistrationMenuService {
  public observer = new Subject();
  public subscribers$ = this.observer.asObservable();

  constructor() { }

  Open() {
    this.observer.next({});
  }
}
