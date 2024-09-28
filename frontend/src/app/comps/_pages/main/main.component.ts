import { Component, OnInit } from '@angular/core';
import { VideoModel } from "../../../models/video.model";
import { ApiService } from "../../../services/api.service";
import { MainVideoCoverComponent } from "../../_models/main-video-cover/main-video-cover.component";
import { Router, RouterLink } from "@angular/router";

@Component({
  selector: 'app-main',
  standalone: true,
  imports: [
    MainVideoCoverComponent,
    RouterLink
  ],
  templateUrl: './main.component.html',
  styleUrl: './main.component.css'
})
export class MainComponent implements OnInit {
  public videos: VideoModel[] = [];

  constructor(
    private api: ApiService,
    private router: Router
  ) {}

  addVideos(arr: VideoModel[]) {
    this.videos = arr;
  }

  go(video: VideoModel) {
    localStorage.setItem('video', JSON.stringify(video));
    this.router.navigate(['watch']).then();
  }

  ngOnInit() {
    this.api.get_videos(10).then(resp => this.addVideos(resp));
  }
}
