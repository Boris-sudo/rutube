import { bootstrapApplication } from '@angular/platform-browser';
import { appConfig } from './app/app.config';
import { AppComponent } from './app/app.component';

console.log(`%c

==================================
= . . . Project is done by . . . =
= . . . . . Boris Kiva . . . . . =
==================================

`, 'font-family: monospace; color: red')

bootstrapApplication(AppComponent, appConfig)
  .catch((err) => console.error(err));
