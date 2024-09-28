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
  ) {
  }

  getSrc() {
    console.log(this.video.shorts_src);
    return this.video.shorts_src;
  }

  like() {
    if (!this.video.is_liked) {
      if (this.video.is_disliked) this.video.dislikes--;
      this.video.is_disliked = false;
      this.video.is_liked = true;
      this.video.likes++;
      this.api.react(this.video.video_id, this.video.is_liked, this.video.is_disliked).then().catch();
    } else {
      if (this.video.is_liked) this.video.likes--;
      this.video.is_liked = false;
      this.api.react(this.video.video_id, this.video.is_liked, this.video.is_disliked).then().catch();
    }
  }

  dislike() {
    if (!this.video.is_disliked) {
      if (this.video.is_liked) this.video.likes--;
      this.video.is_disliked = true;
      this.video.is_liked = false;
      this.video.dislikes++;
      this.api.react(this.video.video_id, this.video.is_liked, this.video.is_disliked).then().catch();
    } else {
      if (this.video.is_disliked) this.video.dislikes--;
      this.video.is_disliked = false;
      this.api.react(this.video.video_id, this.video.is_liked, this.video.is_disliked).then().catch();
    }
  }
}
