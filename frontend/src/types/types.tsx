export interface EventDetails {
    dateTime: string;
    theme: string;
    topics: string[];
  }
  
 export interface ChatResponse {
    choices: Array<{
      message: {
        role: string;
        content: string;
      };
      index: number;
      finish_reason: string;
    }>;
  }
  
export interface Event {
    id: number;
    event_start: string;
    event_end: string;
    theme_id: number;
    theme: {
      theme_text: string;
      topic1: string;
      topic2: string;
      topic3: string;
    };
  }
  
export  interface User {
    id: number;
    username: string;
    avatarURL: string;
    email: string;
    createdAt: string;
  }