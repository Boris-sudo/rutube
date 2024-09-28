from flask import Flask, request, jsonify
import pandas as pd
import json
import random

app = Flask(__name__)

df = pd.read_parquet("./data/video.parquet")
# Filter out hidden or deleted videos
df = df[(df['v_is_hidden'] == False) & (df['v_is_deleted'] == False)]

@app.route('/api/random_videos', methods=['POST'])
def get_random_videos():
    data = request.get_json()
#     print(data)

    # Select 10 random videos
    random_videos = df.sample(n=10, random_state=random.randint(0, 5))  # Set random_state for reproducibility

    # Create a list of dictionaries for the desired output
    video_list = []
    for _, row in random_videos.iterrows():
        video_info = {
            'video_id': row['video_id'],
            'description': row['description'],
            'title': row['title'],
            'views': row['v_year_views'],
            'comments': row['v_total_comments'],
            'likes': row['v_likes'],
            'dislikes': row['v_dislikes']
        }
        video_list.append(video_info)

    return jsonify(video_list)

if __name__ == '__main__':
    app.run(debug=True)
