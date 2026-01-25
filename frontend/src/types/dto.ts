export interface EventThemeDto {
  theme_text: string;
  topic1: string;
  topic2: string;
  topic3: string;
}

export interface EventDto {
  id: number;
  event_start: string;
  event_end: string;
  theme_id: number;
  theme: EventThemeDto;
}

export interface CreateEventDto {
  event_start: string;
  theme: string;
  topics: string[];
}

export interface ChatMessageDto {
  role: string;
  content: string;
}

export interface ChatChoiceDto {
  message: ChatMessageDto;
  index: number;
  finish_reason: string;
}

export interface ChatResponseDto {
  choices: ChatChoiceDto[];
}

export type ChatThemeResponseDto = ChatResponseDto;

export interface MemoDto {
  memo1?: string;
  memo2?: string;
}

export interface SessionDto {
  theme: string;
  date_time: string;
  sessions: number[];
}

export interface ConversationExampleDto {
  english: string;
  japanese: string;
}

export interface ConversationHistoryDto {
  date: string;
  previous_date: string;
  sessions: number;
  completion_rate: string;
  comment: string;
  examples: ConversationExampleDto[];
}

export interface SessionHistoryDto {
  avatar: string;
  user: string;
  theme: string;
  date: string;
  rank: number;
  friend_state: "friend" | "pending" | "unapplied";
}

export interface TopicDto {
  theme: string;
  topics: string[];
}

export interface NotificationDto {
  id: number;
  time: string;
  user: string;
  type: string;
  message: string;
  profile_icon: string;
}

export interface UserProfileDto {
  id: number;
  username: string;
  email: string;
  avatar_url: string;
  role: string;
  created_at: string;
  updated_at: string;
}

export interface UserDto {
  id: number;
  username: string;
  email: string;
  avatar_url: string;
  created_at: string;
}

export interface FriendSummaryDto {
  id: number;
  username: string;
  avatar_url: string;
}

export interface FriendListDto {
  friends: FriendSummaryDto[];
}

export interface FriendInfoDto {
  username: string;
  avatar_url: string;
}

export interface AvatarDto {
  avatar_url: string;
}
