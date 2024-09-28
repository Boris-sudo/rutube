import { AfterViewInit, Component } from '@angular/core';
import { NavigationEnd, Router, RouterLink, RouterLinkActive } from "@angular/router";

@Component({
  selector: 'SideBar',
  standalone: true,
  imports: [
    RouterLink,
    RouterLinkActive
  ],
  styleUrl: './side-bar.component.css',
  templateUrl: './side-bar.component.html'
})
export class SideBarComponent implements AfterViewInit {
  static side_bar_cards_count = 3;
  static is_opened: boolean = true;

  constructor(
    public router: Router
  ) {
    this.router.events.subscribe((val) => {
      if (val instanceof NavigationEnd && window.innerWidth <= 600) {
        SideBarComponent.Hide();
      }
    });
  }

  ngAfterViewInit() {
    const container: HTMLElement = document.getElementById('side-menu-container')!;
    if (container.offsetWidth === window.innerWidth || container.offsetWidth === 230)
      SideBarComponent.is_opened = true;
    else SideBarComponent.is_opened = false;
  }

  static Show() {
    const container: HTMLElement = document.getElementById('side-menu-container')!;
    container.style.maxWidth = 'var(--container-width)';

    if (window.innerWidth > 600) {
      for (let i = 0; i < 1000; i++) {
        if (document.getElementById('side-bar-container-' + i) === null) break;
        const card = document.getElementById('side-bar-container-' + i)!;
        card.style.padding = '12px 12px';
      }
      for (let i = 0; i < 1000; i++) {
        if (document.getElementById('side-bar-card-' + i) === null) break;
        const card = document.getElementById('side-bar-card-' + i)!;
        card.style.fontSize = '16px';
        card.style.flexDirection = 'row';
        card.style.padding = '5px 12px';
      }
    }

    this.is_opened = true;
  }

  static Hide() {
    const container: HTMLElement = document.getElementById('side-menu-container')!;
    container.style.maxWidth = 'var(--container-mini-width)';

    if (window.innerWidth > 600) {
      for (let i = 0; i < 1000; i++) {
        if (document.getElementById('side-bar-container-'+i) === null) break;
        const card = document.getElementById('side-bar-container-'+i)!;
        card.style.padding = '12px 4px';
      }
      for (let i = 0; i < 1000; i++) {
        if (document.getElementById('side-bar-card-'+i) === null) break;
        const card = document.getElementById('side-bar-card-'+i)!;
        card.style.fontSize = '10px';
        card.style.flexDirection = 'column';
        card.style.padding = '3px 0';
      }
    }

    this.is_opened = false;
  }

  static Interact() {
    if (this.is_opened) this.Hide();
    else this.Show();
  }
}
