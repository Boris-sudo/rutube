import { Component, Input } from '@angular/core';
import { VideoModel } from "../../../models/video.model";
import { SafePipe } from "../../../pipes/safe.pipe";
import { NgIf } from "@angular/common";
import { ApiService } from "../../../services/api.service";

@Component({
  selector: 'VideoComp',
  standalone: true,
  imports: [
    SafePipe,
    NgIf
  ],
  templateUrl: './video.component.html',
  styleUrl: './video.component.css'
})
export class VideoComponent {
  @Input() video!: VideoModel;

  constructor(
    private api: ApiService,
  ) {}

  like() {
    if (!this.video.liked) {
      this.video.disliked = false;
      this.video.liked = true;
      this.api.react(this.video.id, this.video.liked, this.video.disliked).then().catch();
    } else {
      this.video.liked = false;
      this.api.react(this.video.id, this.video.liked, this.video.disliked).then().catch();
    }
  }

  dislike() {
    if (!this.video.disliked) {
      this.video.disliked = true;
      this.video.liked = false;
      this.api.react(this.video.id, this.video.liked, this.video.disliked).then().catch();
    } else {
      this.video.disliked = false;
      this.api.react(this.video.id, this.video.liked, this.video.disliked).then().catch();
    }
  }
}
