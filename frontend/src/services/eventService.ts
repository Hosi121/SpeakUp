import api from "./api";
import { ChatResponse, Event, EventDetails } from "../types/types";
import { isTestMode } from "./appMode";
import { loadMockData, saveMockData } from "./mockStore";
import type { TopicGroup } from "./appData";

export const fetchEvents = async (): Promise<Event[]> => {
  try {
    if (isTestMode) {
      return loadMockData<Event[]>("events");
    }
    const response = await api.get("/events");
    return response.data;
  } catch (error) {
    console.error("Failed to fetch events", error);
    throw new Error("イベントの取得に失敗しました");
  }
};

export const createEvent = async (eventData: EventDetails): Promise<EventDetails> => {
  try {
    if (isTestMode) {
      const events = loadMockData<Event[]>("events");
      const nextId =
        events.length > 0 ? Math.max(...events.map((event) => event.id)) + 1 : 1;
      const startAt = new Date(eventData.eventStart);
      const endAt = new Date(startAt.getTime() + 60 * 60 * 1000);
      const [topic1 = "", topic2 = "", topic3 = ""] = eventData.topics;
      const newEvent: Event = {
        id: nextId,
        eventStart: startAt.toISOString(),
        eventEnd: endAt.toISOString(),
        themeId: nextId,
        theme: {
          themeText: eventData.theme,
          topic1,
          topic2,
          topic3,
        },
      };
      saveMockData("events", [...events, newEvent]);
      return eventData;
    }
    const response = await api.post("/events", eventData);
    return response.data;
  } catch (error) {
    console.error("Failed to create event", error);
    throw new Error("イベントの作成に失敗しました");
  }
};

export const generateTheme = async (): Promise<string> => {
  try {
    if (isTestMode) {
      const topics = loadMockData<TopicGroup[]>("topics");
      const themePool = topics.map((topic) => topic.theme).filter(Boolean);
      if (themePool.length === 0) {
        return "Sample theme";
      }
      const pickIndex = Math.floor(Math.random() * themePool.length);
      return themePool[pickIndex];
    }
    const prompt = "イベントのテーマを提案してください。";
    const response = await api.post<ChatResponse>("/chat/theme", { content: prompt });
    return response.data.choices[0].message.content;
  } catch (error) {
    console.error("Failed to generate theme", error);
    throw new Error("テーマの生成に失敗しました");
  }
};
