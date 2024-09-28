import { AfterViewInit, Component, ViewChild } from '@angular/core';
import { SideBarComponent } from "../side-bar/side-bar.component";

@Component({
  selector: 'HeaderComp',
  standalone: true,
  imports: [],
  templateUrl: './header.component.html',
  styleUrl: './header.component.css'
})
export class HeaderComponent {
  @ViewChild('moving_line') moving_line: any;

  constructor() { }

  Open() {
    this.moving_line.nativeElement.style.width = 'var(--hamburger-size)';
  }

  Close() {
    this.moving_line.nativeElement.style.width = 'var(--close-line-hamburger-size)';
  }

  Interact() {
    if (SideBarComponent.is_opened)
      this.Close();
    else
      this.Open();
    SideBarComponent.Interact();
  }
}
