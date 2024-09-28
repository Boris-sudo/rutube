import { Injectable } from '@angular/core';
import { Observable, of } from "rxjs";
import { VideoModel } from "../models/video.model";

@Injectable({
  providedIn: 'root'
})
export class MockService {
  private readonly videos: VideoModel[] = [
    {
      id: 0,
      name: 'How programmers flex on each other',
      description: 'Brainfuck is a minimal esoteric programming language. It provides a 30K 8-bit array that can be modified with 8 different characters.',
      liked: false,
      disliked: false,
      likes_count: 1203,
      dislikes_count: 40,
      comments_count: 130,
      views_count: 6540402
    },
    {
      id: 1,
      name: '10 Math Concepts for Programmers',
      description: 'Learn 10 essential math concepts for software engineering and technical interviews. Understand how programmers use mathematics in fields like AI, game dev, crypto, machine learning, and more.',
      liked: false,
      disliked: true,
      likes_count: 1203,
      dislikes_count: 40,
      comments_count: 130,
      views_count: 6540402
    },
    {
      id: 2,
      name: 'Just the Two of Us - Cover',
      description: 'Hey guys it\'s been a while! Here is my cover of Just the Two of Us by Grover Washington Jr. Thank you for all the support, wow almost 500 subs :) some people were asking, my Instagram is @lucellis',
      liked: false,
      disliked: false,
      likes_count: 1203,
      dislikes_count: 40,
      comments_count: 130,
      views_count: 6540402
    },
  ]


  constructor() {
  }

  randomIntFromInterval(min: number, max: number) { // min and max included
    return Math.floor(Math.random() * (max - min + 1) + min);
  }

  get_video_by_id(id: number) {
    for (const video of this.videos) {
      if (video.id === id)
        return of(JSON.stringify({result: video}));
    }
    return of('{}');
  }

  get_videos(count: number = 10): Observable<any> {
    let result: { result: VideoModel[] } = { result: [] };

    for (let i = 0; i < count; i++) {
      const id = this.randomIntFromInterval(0, this.videos.length - 1);
      result.result.push(this.videos[id]);
    }

    return of(result);
  }

  react(video_id: number, liked: boolean, disliked: boolean) {
    if (liked) {
      this.videos[video_id].liked = true;
      this.videos[video_id].disliked = false;
    }
    if (disliked) {
      this.videos[video_id].liked = false;
      this.videos[video_id].disliked = true;
    }
    return of('done');
  }

  register() {
    return of(JSON.stringify({'uuid': '1'}));
  }
}
