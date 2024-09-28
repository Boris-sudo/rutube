import { Component } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { HeaderComponent } from "./comps/header/header.component";
import { SideBarComponent } from "./comps/side-bar/side-bar.component";
import { FooterComponent } from "./comps/footer/footer.component";
import { ApiService } from "./services/api.service";

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [RouterOutlet, HeaderComponent, SideBarComponent, FooterComponent],
  templateUrl: './app.component.html',
  styleUrl: './app.component.css'
})
export class AppComponent {
  title = 'rutube';

  constructor(
    private api: ApiService,
  ) {
  }
}
