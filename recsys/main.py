from flask import Flask, request, jsonify
import pandas as pd
import torch
import torch.nn as nn
import numpy as np
import service
import random

# Инициализация Flask приложения
app = Flask(__name__)

# Загрузка данных и модели

video_file_path = './data/video.parquet'
video_df = pd.read_parquet(video_file_path, engine='pyarrow')
get_video = service.get_videos_ai


@app.route('/api/predicted_videos', methods=['POST'])
def get_predicted_videos():
    data = request.get_json()
    print(data)
    top_video_ids = get_video(data, num_recommendations=40)
    top_video_ids = random.sample(list(top_video_ids), 10)

    video_list = []
    for video_id in top_video_ids:
        video_info = video_df.iloc[video_id]
        video_data = {
            'video_id': video_info['video_id'],
            'description': video_info['description'],
            'title': video_info['title'],
            'category': video_info['category_id'],
            'views': video_info['v_year_views'],
            'comments': video_info['v_total_comments'],
            'likes': video_info['v_likes'],
            'dislikes': video_info['v_dislikes']
        }
        video_list.append(video_data)

    return jsonify(video_list)


if __name__ == '__main__':
    app.run(debug=True)
