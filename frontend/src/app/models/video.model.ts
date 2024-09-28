export interface VideoModel {
  video_id: number;
  title: string;
  description?: string;
  liked: boolean;
  disliked: boolean;

  views: number;
  likes: number;
  dislikes: number;
  comments: number;
}
