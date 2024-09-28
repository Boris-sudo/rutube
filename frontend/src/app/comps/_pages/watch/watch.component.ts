import { Component, OnInit, ViewChild } from '@angular/core';
import { VideoModel } from "../../../models/video.model";
import { ActivatedRoute, Router } from "@angular/router";
import { ApiService } from "../../../services/api.service";

@Component({
  selector: 'app-watch',
  standalone: true,
  imports: [],
  templateUrl: './watch.component.html',
  styleUrl: './watch.component.css'
})
export class WatchComponent implements OnInit {
  @ViewChild('like') like_icon: any;
  @ViewChild('like_container') like_container: any;
  public video!: VideoModel;

  constructor(
    private router: Router,
    private route: ActivatedRoute,
    private api: ApiService,
  ) {
  }

  ngOnInit() {
    this.route.paramMap.subscribe(params => {
      let video_id = Number(params.get('video_id')!);
      this.api.get_video_by_id(video_id).then(resp => this.video = resp);
    });
  }

  parseViewsCount(): string {
    let res = '';
    let number = String(this.video.views);
    for(let i=number.length-1; i>=0; --i) {
      if ((number.length - i + 1) % 3 === 0)
        res = res + ' ';
      res = res + number[i];
    }
    return res;
  }

  likeVideo() {
    if (!this.video.liked) {
      this.like_container.nativeElement.style.paddingLeft = '25px'
      this.like_icon.nativeElement.style.position = 'absolute';
      this.like_icon.nativeElement.classList.add('animated_like')

      setTimeout(() => {
        this.like_icon.nativeElement.classList.remove('animated_like')
      }, 600);

      if (this.video.disliked) this.video.dislikes--;
      this.video.disliked = false;
      this.video.liked = true;
      this.video.likes++;
      this.api.react(this.video.video_id, this.video.liked, this.video.disliked).then().catch();
    } else {
      if (this.video.liked)
        this.video.likes--;
      this.video.liked = false;
      this.api.react(this.video.video_id, this.video.liked, this.video.disliked).then().catch();
    }
  }

  dislikeVideo() {
    if (!this.video.disliked) {
      if (this.video.liked) this.video.likes--;
      this.video.disliked = true;
      this.video.liked = false;
      this.video.dislikes++;
      this.api.react(this.video.video_id, this.video.liked, this.video.disliked).then().catch();
    } else {
      if (this.video.disliked) this.video.dislikes--;
      this.video.disliked = false;
      this.api.react(this.video.video_id, this.video.liked, this.video.disliked).then().catch();
    }
  }
}
