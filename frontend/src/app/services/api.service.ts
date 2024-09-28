import { Injectable } from '@angular/core';
import { HttpClient } from "@angular/common/http";
import { MockService } from "./mock.service";
import { VideoModel } from "../models/video.model";
import { firstValueFrom, onErrorResumeNext } from "rxjs";

@Injectable({
  providedIn: 'root'
})
export class ApiService {
  private readonly base_url = 'http://localhost:3000';

  private request(url: string): string {
    return `${ this.base_url }/${ url }/`;
  }

  constructor(
    private http: HttpClient,
    private mock: MockService,
  ) { }

  async get_video_by_id(id: number): Promise<VideoModel> {
    let result!: VideoModel;

    const resp: any = await firstValueFrom(onErrorResumeNext(
      this.http.get(this.request('')),
      this.mock.get_video_by_id(id)
    ));

    result = JSON.parse(resp).result;

    return result;
  }

  async get_videos(count: number = 10): Promise<VideoModel[]> {
    let result: VideoModel[] = [];

    const resp: any = await firstValueFrom(onErrorResumeNext(
      this.http.get(this.request('')),
      this.mock.get_videos(count)
    ));

    for (const video of resp.result) {
      result.push(video);
    }

    return result;
  }

  async react(video_id: number, liked: boolean, disliked: boolean) {
    const user_id = localStorage.getItem('session_id');

    const resp: any = await firstValueFrom(onErrorResumeNext(
      this.http.post(this.request('recsys/preferences/save'), {
        user_id: String(user_id),
        video_id: String(video_id),
        liked: String(liked),
        disliked: String(disliked)
      }),
      this.mock.react(video_id, liked, disliked)
    ));

    return resp;
  }
}
