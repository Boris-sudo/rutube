import { AfterViewInit, Component, OnInit } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { HeaderComponent } from "./comps/header/header.component";
import { SideBarComponent } from "./comps/side-bar/side-bar.component";
import { FooterComponent } from "./comps/footer/footer.component";
import { ApiService } from "./services/api.service";
import { RegistrationComponent } from "./comps/_models/registration/registration.component";
import { OpenRegistrationMenuService } from "./services/open-registration-menu.service";

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [RouterOutlet, HeaderComponent, SideBarComponent, FooterComponent, RegistrationComponent],
  templateUrl: './app.component.html',
  styleUrl: './app.component.css'
})
export class AppComponent implements AfterViewInit {
  title = 'rutube';

  constructor(
    private api: ApiService,
    private openRegistrationMenuService: OpenRegistrationMenuService,
  ) {}

  ngAfterViewInit() {
    const user_id = this.api.get_user_id();
    if (user_id === null)
      this.openRegistrationMenuService.Open();
  }
}
