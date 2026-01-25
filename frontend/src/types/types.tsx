export interface EventDetails {
  eventStart: string;
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
      finishReason: string;
    }>;
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
