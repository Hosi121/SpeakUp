import api from "./api";
import { isTestMode } from "./appMode";
import { loadMockData } from "./mockStore";
import {
  type ConversationHistoryItem,
  type NotificationItem,
  type SessionData,
  type SessionHistoryItem,
  type TopicGroup,
} from "../types/types";
import {
  type ConversationHistoryDto,
  type NotificationDto,
  type SessionDto,
  type SessionHistoryDto,
  type TopicDto,
} from "../types/dto";
import { toApiError } from "./errorUtils";

const mapSessionDto = (dto: SessionDto): SessionData => ({
  theme: dto.theme,
  dateTime: dto.date_time,
  sessions: dto.sessions,
});

const mapConversationHistoryDto = (
  dto: ConversationHistoryDto
): ConversationHistoryItem => ({
  date: dto.date,
  previousDate: dto.previous_date,
  sessions: dto.sessions,
  completionRate: dto.completion_rate,
  comment: dto.comment,
  examples: dto.examples.map((example) => ({
    english: example.english,
    japanese: example.japanese,
  })),
});

const mapSessionHistoryDto = (dto: SessionHistoryDto): SessionHistoryItem => ({
  avatar: dto.avatar,
  user: dto.user,
  theme: dto.theme,
  date: dto.date,
  rank: dto.rank,
  friendState: dto.friend_state,
});

const mapTopicDto = (dto: TopicDto): TopicGroup => ({
  theme: dto.theme,
  topics: dto.topics,
});

const mapNotificationDto = (dto: NotificationDto): NotificationItem => ({
  id: dto.id,
  time: dto.time,
  user: dto.user,
  type: dto.type,
  message: dto.message,
  profileIcon: dto.profile_icon,
});

export const fetchSessions = async (): Promise<SessionData[]> => {
  if (isTestMode) {
    return loadMockData<SessionData[]>("sessions");
  }
  try {
    const response = await api.get<SessionDto[]>("/sessions");
    return response.data.map(mapSessionDto);
  } catch (error) {
    throw toApiError(error, "セッションの取得に失敗しました");
  }
};

export const fetchConversationHistory = async (): Promise<
  ConversationHistoryItem[]
> => {
  if (isTestMode) {
    return loadMockData<ConversationHistoryItem[]>("conversationHistory");
  }
  try {
    const response = await api.get<ConversationHistoryDto[]>(
      "/conversation_history"
    );
    return response.data.map(mapConversationHistoryDto);
  } catch (error) {
    throw toApiError(error, "会話履歴の取得に失敗しました");
  }
};

export const fetchSessionHistory = async (): Promise<SessionHistoryItem[]> => {
  if (isTestMode) {
    return loadMockData<SessionHistoryItem[]>("sessionHistory");
  }
  try {
    const response = await api.get<SessionHistoryDto[]>("/session_history");
    return response.data.map(mapSessionHistoryDto);
  } catch (error) {
    throw toApiError(error, "セッション履歴の取得に失敗しました");
  }
};

export const fetchTopics = async (): Promise<TopicGroup[]> => {
  if (isTestMode) {
    return loadMockData<TopicGroup[]>("topics");
  }
  try {
    const response = await api.get<TopicDto[]>("/topics");
    return response.data.map(mapTopicDto);
  } catch (error) {
    throw toApiError(error, "トピックの取得に失敗しました");
  }
};

export const fetchNotifications = async (): Promise<NotificationItem[]> => {
  if (isTestMode) {
    return loadMockData<NotificationItem[]>("notifications");
  }
  try {
    const response = await api.get<NotificationDto[]>("/notifications");
    return response.data.map(mapNotificationDto);
  } catch (error) {
    throw toApiError(error, "通知の取得に失敗しました");
  }
};
