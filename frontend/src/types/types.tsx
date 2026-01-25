export interface EventDetails {
  eventStart: string;
  theme: string;
  topics: string[];
}

export interface Event {
    id: number;
    eventStart: string;
    eventEnd: string;
    themeId: number;
    theme: {
      themeText: string;
      topic1: string;
      topic2: string;
      topic3: string;
    };
  }
  
export  interface User {
    id: number;
    username: string;
    avatarUrl: string;
    email: string;
    createdAt: string;
  }

export interface UserProfile {
  id: number;
  username: string;
  email: string;
  avatarUrl: string;
  role: string;
  createdAt: string;
  updatedAt: string;
}

export interface UserNotes {
  carryInMemo: string;
  wordList: string;
}

export interface SessionData {
  theme: string;
  dateTime: string;
  sessions: number[];
}

export interface ConversationExample {
  english: string;
  japanese: string;
}

export interface ConversationHistoryItem {
  date: string;
  previousDate: string;
  sessions: number;
  completionRate: string;
  comment: string;
  examples: ConversationExample[];
}

export type FriendState = "friend" | "pending" | "unapplied";

export interface SessionHistoryItem {
  avatar: string;
  user: string;
  theme: string;
  date: string;
  rank: number;
  friendState: FriendState;
}

export interface TopicGroup {
  theme: string;
  topics: string[];
}

export interface NotificationItem {
  id: number;
  time: string;
  user: string;
  type: string;
  message: string;
  profileIcon: string;
}

export interface FriendSummary {
  id: number;
  username: string;
  avatarUrl: string;
}

export interface FriendInfo {
  username: string;
  avatarUrl: string;
}
