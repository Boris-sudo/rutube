import { Injectable } from '@angular/core';
import { HttpClient } from "@angular/common/http";
import { MockService } from "./mock.service";
import { VideoModel } from "../models/video.model";
import { firstValueFrom, onErrorResumeNext } from "rxjs";
import { UserModel } from "../models/user.model";

@Injectable({
  providedIn: 'root'
})
export class ApiService {
  private readonly base_url = 'http://localhost:8080/api';
  private readonly user_id_key: string = 'session';

  private request(url: string): string {
    return `${ this.base_url }/${ url }`;
  }

  constructor(
    private http: HttpClient,
    private mock: MockService,
  ) { }

  get_user_id() {
    return localStorage.getItem(this.user_id_key);
  }

  randomIntFromInterval(min: number, max: number): number {
    return Math.floor(Math.random() * (max - min + 1) + min);
  }

  save_user_id(user_id: string) {
    localStorage.setItem(this.user_id_key, user_id);
  }

  async get_videos(count: number = 10): Promise<VideoModel[]> {
    let result: VideoModel[] = [];
    const user_id = this.get_user_id();

    const resp: any = await firstValueFrom(onErrorResumeNext(
      this.http.post(this.request('recsys/videos'), {user_id: user_id}),
      this.mock.get_videos(count)
    ));

    function getRandomColor() {
      const letters = '0123456789ABCDEF';
      let color = '#';
      for (let i = 0; i < 6; i++) {
        color += letters[Math.floor(Math.random() * 16)];
      }
      return color;
    }
    function setRandomGradient() {
      const color1 = getRandomColor();
      const color2 = getRandomColor();
      const angle = Math.floor(Math.random() * 360);
      return `linear-gradient(${angle}deg, ${color1}, ${color2})`;
    }

    let index = 0;
    for (const video of resp) {
      let res_video: any = {};
      for (const key in video) res_video[key] = video[key];
      res_video.gradient = setRandomGradient();
      index++;
      result.push(res_video);
    }

    return result;
  }

  async react(video_id: number, liked: boolean, disliked: boolean) {
    const user_id = this.get_user_id();

    const resp: any = await firstValueFrom(
      this.http.post(this.request('recsys/preferences/save'), {
        user_id: String(user_id),
        video_id: String(video_id),
        is_liked: liked,
        is_disliked: disliked
      })
    );

    return resp;
  }

  async save_to_history(video_id: number) {
    const user_id = this.get_user_id();

    const resp: any = await firstValueFrom(onErrorResumeNext(
      this.http.post(this.request('recsys/preferences/save'), {
        user_id: String(user_id),
        video_id: String(video_id),
      }),
    ));

    return resp;
  }

  async login(email: string, password: string): Promise<UserModel> {
    const resp: any = await firstValueFrom(onErrorResumeNext(
      this.http.post(this.request('accounts/login'), {
        email: email,
        password: password,
      }),
    ));

    await this.save_to_history(resp.uuid);

    return resp;
  }

  async register(login: string = '', email: string | undefined = '', password: string | undefined = '', name: string | undefined = '', surname: string | undefined = '', region: string | undefined = '', city: string | undefined = ''): Promise<void> {
    let res: any;
    if (login === '') {
      // res = await firstValueFrom(this.mock.register());
      res = await firstValueFrom(this.http.post(this.request('accounts/register'), { login: '' }));
    } else {
      // res = await firstValueFrom(this.mock.register());
      res = await firstValueFrom(this.http.post(this.request('accounts/register'), { login: login, email: email, password: password, name: name, surname: surname, region: region, city: city}));
    }

    this.save_user_id(res.uuid);
  }
}
