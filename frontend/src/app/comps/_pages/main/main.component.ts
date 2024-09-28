import { Component, OnInit } from '@angular/core';
import { VideoModel } from "../../../models/video.model";
import { ApiService } from "../../../services/api.service";
import { MainVideoCoverComponent } from "../../_models/main-video-cover/main-video-cover.component";
import { RouterLink } from "@angular/router";

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
  ) {}

  addVideos(arr: VideoModel[]) {
    this.videos = arr;
  }

  ngOnInit() {
    this.api.get_videos(10).then(resp => this.addVideos(resp));
  }
}
