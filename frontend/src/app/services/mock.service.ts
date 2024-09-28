import { Injectable } from '@angular/core';
import { Observable, of } from "rxjs";
import { VideoModel } from "../models/video.model";

@Injectable({
  providedIn: 'root'
})
export class MockService {
  private readonly videos: VideoModel[] = [
    {
      video_id: 0,
      title: 'How programmers flex on each other',
      description: 'Brainfuck is a minimal esoteric programming language. It provides a 30K 8-bit array that can be modified with 8 different characters.',
      is_liked: false,
      is_disliked: false,
      likes: 1203,
      dislikes: 40,
      comments: 130,
      views: 12345
    },
    {
      video_id: 1,
      title: '10 Math Concepts for Programmers',
      description: 'Learn 10 essential math concepts for software engineering and technical interviews. Understand how programmers use mathematics in fields like AI, game dev, crypto, machine learning, and more.',
      is_liked: false,
      is_disliked: true,
      likes: 1203,
      dislikes: 40,
      comments: 130,
      views: 1234
    },
    {
      video_id: 2,
      title: 'Just the Two of Us - Cover',
      description: 'Hey guys it\'s been a while! Here is my cover of Just the Two of Us by Grover Washington Jr. Thank you for all the support, wow almost 500 subs :) some people were asking, my Instagram is @lucellis',
      is_liked: false,
      is_disliked: false,
      likes: 1203,
      dislikes: 40,
      comments: 130,
      views: 123
    },
  ]


  constructor() {
  }

  randomIntFromInterval(min: number, max: number) { // min and max included
    return Math.floor(Math.random() * (max - min + 1) + min);
  }

  get_video_by_id(id: number) {
    for (const video of this.videos) {
      if (video.video_id === id)
        return of(JSON.stringify({result: video}));
    }
    return of('{}');
  }

  get_videos(count: number = 10): Observable<any> {
    let result: VideoModel[] = [];

    for (let i = 0; i < count; i++) {
      const id = this.randomIntFromInterval(0, this.videos.length - 1);
      result.push(this.videos[id]);
    }

    return of(result);
  }

  react(video_id: number, liked: boolean, disliked: boolean) {
    if (liked) {
      this.videos[video_id].is_liked = true;
      this.videos[video_id].is_disliked = false;
    }
    if (disliked) {
      this.videos[video_id].is_liked = false;
      this.videos[video_id].is_disliked = true;
    }
    return of('done');
  }

  register() {
    return of(JSON.stringify({'uuid': '1'}));
  }
}
