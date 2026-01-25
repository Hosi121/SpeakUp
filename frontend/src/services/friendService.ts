import api from "./api";
import { toApiError } from "./errorUtils";
import type { FriendInfo, FriendSummary } from "../types/types";
import type {
  FriendInfoDto,
  FriendListDto,
  FriendSummaryDto,
} from "../types/dto";

const mapFriendSummaryDto = (dto: FriendSummaryDto): FriendSummary => ({
  id: dto.id,
  username: dto.username,
  avatarUrl: dto.avatar_url,
});

const mapFriendInfoDto = (dto: FriendInfoDto): FriendInfo => ({
  username: dto.username,
  avatarUrl: dto.avatar_url,
});

export const fetchFriendList = async (): Promise<FriendSummary[]> => {
  try {
    const response = await api.get<FriendListDto>("/friend/me");
    return response.data.friends.map(mapFriendSummaryDto);
  } catch (error) {
    throw toApiError(error, "フレンド一覧の取得に失敗しました");
  }
};

export const fetchFriendInfo = async (friendName: string): Promise<FriendInfo> => {
  try {
    const response = await api.get<FriendInfoDto>(`/friend/${friendName}`);
    return mapFriendInfoDto(response.data);
  } catch (error) {
    throw toApiError(error, "フレンド情報の取得に失敗しました");
  }
};

export const sendFriendRequest = async (
  targetUserId: number
): Promise<void> => {
  try {
    await api.post("/friend/register", { target_user_id: targetUserId });
  } catch (error) {
    throw toApiError(error, "フレンド申請に失敗しました");
  }
};
