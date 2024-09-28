export interface VideoModel {
  id: number;
  name: string;
  description?: string;
  liked: boolean;
  disliked: boolean;

  views_count: number;
  likes_count: number;
  dislikes_count: number;
  comments_count: number;
}
