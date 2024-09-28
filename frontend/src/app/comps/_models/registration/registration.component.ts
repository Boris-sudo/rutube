import { AfterViewInit, Component, ViewChild } from '@angular/core';
import { OpenRegistrationMenuService } from "../../../services/open-registration-menu.service";
import { ApiService } from "../../../services/api.service";
import { FormsModule } from "@angular/forms";

@Component({
  selector: 'Registration',
  standalone: true,
  imports: [
    FormsModule
  ],
  templateUrl: './registration.component.html',
  styleUrl: './registration.component.css'
})
export class RegistrationComponent implements AfterViewInit {
  @ViewChild('container') container: any;
  register: boolean = false;
  register_data = {
    login: '',
    email: '',
    password: '',
    name: '',
    surname: '',
    region: '',
    city: '',
  }
  login_data = {
    email: '',
    password: '',
  }

  constructor(
    private openRegistrationMenuService: OpenRegistrationMenuService,
    private api: ApiService,
  ) {}

  ngAfterViewInit() {
    this.openRegistrationMenuService.subscribers$.subscribe(resp => {
      this.Open();
    });
  }

  Open() {
    this.container.nativeElement.style.display = 'flex';
  }

  Close() {
    this.container.nativeElement.style.display = 'none';
    if (this.api.get_user_id() === null)
      this.api.register('').then(window.location.reload);
  }

  Login() {
    this.api.login(this.login_data.email, this.login_data.password).then(window.location.reload);
  }

  Register() {
    this.api.register(
      this.register_data.login, this.register_data.email,
      this.register_data.password, this.register_data.name,
      this.register_data.surname, this.register_data.region,
      this.register_data.city
    ).then(() => {
      window.location.reload();
    });
  }

  LoginButtonDisabled() {
    return this.login_data.password.length === 0 || this.login_data.email.length === 0;
  }

  RegisterButtonDisabled() {
    if (this.register_data.login.length < 5) return true;
    if (this.register_data.email.split('@').length < 2 || this.register_data.email.split('.').length < 2) return true;
    if (this.register_data.password.length < 8) return true;
    if (this.register_data.name.length === 0 || this.register_data.city.length === 0 || this.register_data.region.length === 0 || this.register_data.surname.length === 0) return true;
    return false;
  }
}
