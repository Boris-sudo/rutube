import pandas as pd
import torch
import torch.nn as nn
import torch.optim as optim
from torch.utils.data import Dataset, DataLoader
import numpy as np
import random

parquet_file_path = './data/video_1.csv'
log_file_path = './data/log_5_1.csv'

interactions_log = pd.read_csv(log_file_path, engine='pyarrow')
video_df = pd.read_csv(parquet_file_path, engine='pyarrow')

interactions = pd.merge(interactions_log, video_df[['video_id', 'v_duration']], on='video_id', how='left')

# Удаляем строки, где video_id отсутствует (None или NaN) или v_duration равно 0
interactions = interactions.dropna(subset=['video_id'])  # Удаляем строки с отсутствующим video_id
interactions = interactions[interactions['v_duration'] > 0]  # Удаляем строки с нулевой продолжительностью

# Вычисляем interaction_score
interactions['interaction_score'] = (interactions['watchtime'] / interactions['v_duration']) * 2

# Ограничиваем значения в пределах [0, 1] и берем нижнюю целую часть
interactions['interaction_score'] = np.clip(interactions['interaction_score'], 0, 1)

final_interactions = interactions[['user_id', 'video_id', 'watchtime', 'v_duration', 'interaction_score']]
df = pd.DataFrame(final_interactions)

# Костыль перевода строк в числа
user_mapping = {user: idx for idx, user in enumerate(df['user_id'].unique())}
video_mapping = {video: idx for idx, video in enumerate(df['video_id'].unique())}
df['user_idx'] = df['user_id'].map(user_mapping)
df['video_idx'] = df['video_id'].map(video_mapping)

# Преобразование данных
user_ids = df['user_idx'].values
video_ids = df['video_idx'].values
interaction_scores = df['interaction_score'].values

num_users = len(df['user_id']) + 1
num_videos = len(df['video_id']) + 1


# Модель фильтрации
class ImplicitMatrixFactorization(nn.Module):
    def __init__(self, num_users, num_videos, embedding_dim=10):
        super(ImplicitMatrixFactorization, self).__init__()
        self.user_embedding = nn.Embedding(num_users, embedding_dim, sparse=True)
        self.video_embedding = nn.Embedding(num_videos, embedding_dim, sparse=True)

    def forward(self, user_id, video_id):
        user_embed = self.user_embedding(user_id)
        video_embed = self.video_embedding(video_id)
        interaction = (user_embed * video_embed).sum(dim=-1)
        return interaction


model = ImplicitMatrixFactorization(num_users=num_users, num_videos=num_videos, embedding_dim=10)

# Move the model to GPU
device = torch.device('cuda' if torch.cuda.is_available() else 'cpu')
torch.load('./models/model.pt')
model = model.to(device)


def get_videos_ai(user, num_recommendations):
    print(user)
    interaction_list = []
    for preference in user['video_preferences']:
        video_id = preference['VideoId']
        if preference['IsLiked']:
            score = 1
        elif preference['IsDisliked']:
            score = 0
        else:
            score = 0.5
        interaction_list.append({"video_id": video_id, "interaction_score": score})

    interaction_list = pd.DataFrame(interaction_list)

    with torch.no_grad():
        predictions = model.predict(interaction_list).numpy()

    interaction_list['predicted_interaction'] = predictions

    top_videos = interaction_list.sort_values(by='predicted_interaction', ascending=False)['video_id'].values

    print(f"Топ рекомендованных видео для нового пользователя: {top_videos}")
    return top_videos
