import api from "./api";
import { toApiError } from "./errorUtils";
import type {
  FriendSummary,
  User,
  UserProfile,
} from "../types/types";
import type {
  AvatarDto,
  FriendSummaryDto,
  UserDto,
  UserProfileDto,
} from "../types/dto";

const mapUserDto = (dto: UserDto): User => ({
  id: dto.id,
  username: dto.username,
  avatarUrl: dto.avatar_url,
  email: dto.email,
  createdAt: dto.created_at,
});

const mapUserProfileDto = (dto: UserProfileDto): UserProfile => ({
  id: dto.id,
  username: dto.username,
  email: dto.email,
  avatarUrl: dto.avatar_url,
  role: dto.role,
  createdAt: dto.created_at,
  updatedAt: dto.updated_at,
});

const mapFriendSummaryDto = (dto: FriendSummaryDto): FriendSummary => ({
  id: dto.id,
  username: dto.username,
  avatarUrl: dto.avatar_url,
});

export const fetchUserProfile = async (): Promise<UserProfile> => {
  try {
    const response = await api.get<UserProfileDto>("/user/info");
    return mapUserProfileDto(response.data);
  } catch (error) {
    throw toApiError(error, "ユーザー情報の取得に失敗しました");
  }
};

export const updateUserProfile = async (update: {
  username?: string;
  email?: string;
}): Promise<void> => {
  try {
    await api.put("/user/update", update);
  } catch (error) {
    throw toApiError(error, "ユーザー情報の更新に失敗しました");
  }
};

export const uploadAvatar = async (file: File): Promise<string> => {
  try {
    const formData = new FormData();
    formData.append("avatar", file);
    const response = await api.put<AvatarDto>("/user/avatar", formData, {
      headers: {
        "Content-Type": "multipart/form-data",
      },
    });
    return response.data.avatar_url;
  } catch (error) {
    throw toApiError(error, "アバターの更新に失敗しました");
  }
};

export const searchUsers = async (query: string): Promise<User[]> => {
  try {
    const encodedQuery = encodeURIComponent(query);
    const response = await api.get<UserDto[]>(`/users/search?q=${encodedQuery}`);
    return response.data.map(mapUserDto);
  } catch (error) {
    throw toApiError(error, "ユーザーの検索に失敗しました");
  }
};

export const fetchUserSummaryById = async (
  userId: number
): Promise<FriendSummary> => {
  try {
    const [userResponse, avatarResponse] = await Promise.all([
      api.get<FriendSummaryDto>(`/users/search/id/${userId}`),
      api.get<AvatarDto>(`/users/${userId}/avatar`),
    ]);
    return mapFriendSummaryDto({
      ...userResponse.data,
      avatar_url: avatarResponse.data.avatar_url,
    });
  } catch (error) {
    throw toApiError(error, "ユーザー情報の取得に失敗しました");
  }
};
