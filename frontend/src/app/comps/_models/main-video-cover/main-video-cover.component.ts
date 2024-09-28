import { Component, Input } from '@angular/core';
import { VideoModel } from "../../../models/video.model";

@Component({
  selector: 'VideoCover',
  standalone: true,
  imports: [],
  templateUrl: './main-video-cover.component.html',
  styleUrl: './main-video-cover.component.css'
})
export class MainVideoCoverComponent {
  @Input() video!: VideoModel;

  constructor() {
  }
}
