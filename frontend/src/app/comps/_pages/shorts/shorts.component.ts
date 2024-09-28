import {
  AfterViewInit,
  Component,
  ComponentFactoryResolver, ComponentRef,
  OnInit,
  Renderer2,
  ViewChild,
  ViewContainerRef
} from '@angular/core';
import { VideoModel } from "../../../models/video.model";
import { ApiService } from "../../../services/api.service";
import { VideoComponent } from "../../_models/video/video.component";

@Component({
  selector: 'app-shorts',
  standalone: true,
  imports: [
    VideoComponent
  ],
  templateUrl: './shorts.component.html',
  styleUrl: './shorts.component.css'
})
export class ShortsComponent implements OnInit, AfterViewInit {
  @ViewChild('container') container: any;
  public videos: VideoModel[] = [];
  public current_index: number = 0;

  constructor(
    private api: ApiService,
    private renderer: Renderer2,
    private viewContainerRef: ViewContainerRef,
    private resolver: ComponentFactoryResolver,
  ) {}

  ngOnInit() {
    this.addVideos();
  }

  addVideos() {
    console.log('adding videos')
    this.api.get_videos().then((resp) => {
      console.log('added videos')
      for (const video of resp)
        this.videos.push(video);
      if (this.videos.length > 100) {
        this.current_index -= this.videos.length - 100;
        this.videos = this.videos.splice(this.videos.length - 100, 100);
      }
    })
  }

  ngAfterViewInit() {
    const keys = {37: 1, 38: 1, 39: 1, 40: 1};

    function preventDefault(e: any) {
      e.preventDefault();
    }

    function preventDefaultForScrollKeys(e: any) {
      // @ts-ignore
      if (keys[e.keyCode]) {
        preventDefault(e);
        return false;
      }
      return true;
    }

    var supportsPassive = false;
    try {
      // @ts-ignore
      window.addEventListener("test", null, Object.defineProperty({}, 'passive', {
        get: function () { supportsPassive = true; }
      }));
    } catch(e) {}

    var wheelOpt = supportsPassive ? { passive: false } : false;
    var wheelEvent = 'onwheel' in document.createElement('div') ? 'wheel' : 'mousewheel';

    const disableScroll = () => {
      window.addEventListener('DOMMouseScroll', preventDefault, false); // older FF
      // window.addEventListener(wheelEvent, preventDefault, wheelOpt); // modern desktop
      // window.addEventListener('touchmove', preventDefault, wheelOpt); // mobile
      window.addEventListener('keydown', preventDefaultForScrollKeys, false);
    }

    disableScroll();
    window.addEventListener('keydown', (e) => {
      if (e.key == 'ArrowDown')
        this.scrollDown();
      if (e.key === 'ArrowUp')
        this.scrollUp();
    })
  }

  reload() {
    this.videos = [];
    this.addVideos();
  }

  scrollUp() {
    const prev = document.getElementsByClassName('prev').item(0);
    const current = document.getElementsByClassName('current').item(0)!;
    const next = document.getElementsByClassName('next').item(0)!;

    if (this.current_index === 0 || prev === null) {
      // @ts-ignore
      current.style.transform = 'ScaleY(1.1) translateX(-50%)';
      // @ts-ignore
      next.style.top = 'calc(90% * 1.1 + 30px)';
      setTimeout(() => {
        // @ts-ignore
        current.style.transform = 'ScaleY(1) translateX(-50%)';
        // @ts-ignore
        next.style.top = '';
      }, 300);
      setTimeout(this.reload, 600)
      return;
    }

    prev.classList.add('current');
    prev.classList.remove('prev');

    current.classList.add('next');
    current.classList.remove('current');

    next.remove();

    const div = this.renderer.createElement('div');
    this.renderer.addClass(div, 'prev');
    this.renderer.addClass(div, 'video');
    this.renderer.appendChild(this.container.nativeElement, div);

    const componentFactory = this.resolver.resolveComponentFactory(VideoComponent);
    const componentRef: ComponentRef<VideoComponent> = this.viewContainerRef.createComponent(componentFactory);
    this.renderer.appendChild(div, componentRef.location.nativeElement);

    componentRef.instance.video = this.videos[this.current_index - 1];

    this.current_index--;
  }

  scrollDown() {
    const prev = document.getElementsByClassName('prev').item(0);
    const current = document.getElementsByClassName('current').item(0)!;
    const next = document.getElementsByClassName('next').item(0)!;

    if (prev !== null) {
      prev.remove();
    }

    current.classList.add('prev');
    current.classList.remove('current');

    next.classList.remove('next');
    next.classList.add('current');

    const div = this.renderer.createElement('div');
    this.renderer.addClass(div, 'video');
    this.renderer.appendChild(this.container.nativeElement, div);

    const componentFactory = this.resolver.resolveComponentFactory(VideoComponent);
    const componentRef: ComponentRef<VideoComponent> = this.viewContainerRef.createComponent(componentFactory);
    this.renderer.appendChild(div, componentRef.location.nativeElement);

    componentRef.instance.video = this.videos[this.current_index + 2];
    setTimeout(() => {
      this.renderer.addClass(div, 'next');
    });

    this.current_index++;

    if (this.current_index % 10 === 8) {
      this.addVideos();
    }
  }
}
