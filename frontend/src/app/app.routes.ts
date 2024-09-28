import { Routes } from '@angular/router';
import { MainComponent } from "./comps/_pages/main/main.component";
import { SavedComponent } from "./comps/_pages/saved/saved.component";
import { ShortsComponent } from "./comps/_pages/shorts/shorts.component";
import { WatchComponent } from "./comps/_pages/watch/watch.component";

export const routes: Routes = [
  { path: '', pathMatch: 'full', redirectTo: 'main' },
  { path: 'main', component: MainComponent },
  { path: 'saved', component: SavedComponent },
  { path: 'shorts', component: ShortsComponent },
  { path: 'watch/:video_id', component: WatchComponent },
  { path: '**', pathMatch: 'full', redirectTo: 'main' },
];
