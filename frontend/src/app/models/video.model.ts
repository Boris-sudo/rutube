export interface VideoModel {
  video_id: number;
  title: string;
  description?: string;
  is_liked: boolean;
  is_disliked: boolean;

  views: number;
  likes: number;
  dislikes: number;
  comments: number;
}
