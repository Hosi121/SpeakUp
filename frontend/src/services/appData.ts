import api from "./api";
import { isTestMode } from "./appMode";
import { loadMockData } from "./mockStore";

export type SessionData = {
  theme: string;
  dateTime: string;
  sessions: number[];
};

export type ConversationExample = {
  english: string;
  japanese: string;
};

export type ConversationHistoryItem = {
  date: string;
  previousDate: string;
  sessions: number;
  completionRate: string;
  comment: string;
  examples: ConversationExample[];
};

export type SessionHistoryItem = {
  avatar: string;
  user: string;
  theme: string;
  date: string;
  rank: number;
  friedstate: "friend" | "pending" | "unapplied";
};

export type TopicGroup = {
  theme: string;
  topics: string[];
};

export type NotificationItem = {
  id: number;
  time: string;
  user: string;
  type: string;
  message: string;
  profileIcon: string;
};

export const fetchSessions = async (): Promise<SessionData[]> => {
  if (isTestMode) {
    return loadMockData<SessionData[]>("sessions");
  }
  const response = await api.get<SessionData[]>("/sessions");
  return response.data;
};

export const fetchConversationHistory = async (): Promise<
  ConversationHistoryItem[]
> => {
  if (isTestMode) {
    return loadMockData<ConversationHistoryItem[]>("conversationHistory");
  }
  const response = await api.get<ConversationHistoryItem[]>(
    "/conversation_history"
  );
  return response.data;
};

export const fetchSessionHistory = async (): Promise<SessionHistoryItem[]> => {
  if (isTestMode) {
    return loadMockData<SessionHistoryItem[]>("sessionHistory");
  }
  const response = await api.get<SessionHistoryItem[]>("/session_history");
  return response.data;
};

export const fetchTopics = async (): Promise<TopicGroup[]> => {
  if (isTestMode) {
    return loadMockData<TopicGroup[]>("topics");
  }
  const response = await api.get<TopicGroup[]>("/topics");
  return response.data;
};

export const fetchNotifications = async (): Promise<NotificationItem[]> => {
  if (isTestMode) {
    return loadMockData<NotificationItem[]>("notifications");
  }
  const response = await api.get<NotificationItem[]>("/notifications");
  return response.data;
};
