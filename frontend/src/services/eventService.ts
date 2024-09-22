import api from "./api";
import { ChatResponse, Event, EventDetails } from "../types/types";

export const fetchEvents = async (): Promise<Event[]> => {
  try {
    const response = await api.get("/events");
    return response.data;
  } catch (error) {
    console.error("Failed to fetch events", error);
    throw new Error("イベントの取得に失敗しました");
  }
};

export const createEvent = async (eventData: EventDetails): Promise<EventDetails> => {
  try {
    const response = await api.post("/events", eventData);
    return response.data;
  } catch (error) {
    console.error("Failed to create event", error);
    throw new Error("イベントの作成に失敗しました");
  }
};

export const generateTheme = async (): Promise<string> => {
  try {
    const prompt = "イベントのテーマを提案してください。";
    const response = await api.post<ChatResponse>("/chat/theme", { content: prompt });
    return response.data.choices[0].message.content;
  } catch (error) {
    console.error("Failed to generate theme", error);
    throw new Error("テーマの生成に失敗しました");
  }
};