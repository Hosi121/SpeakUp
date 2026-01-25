import { apiRaw } from "./api";
import { Event, EventDetails } from "../types/types";
import type { ChatThemeResponseDto, CreateEventDto, EventDto } from "../types/dto";
import { isTestMode } from "./appMode";
import { loadMockData, saveMockData } from "./mockStore";
import type { TopicGroup } from "../types/types";
import { toApiError } from "./errorUtils";

const mapEventDto = (dto: EventDto): Event => ({
  id: dto.id,
  eventStart: dto.event_start,
  eventEnd: dto.event_end,
  themeId: dto.theme_id,
  theme: {
    themeText: dto.theme.theme_text,
    topic1: dto.theme.topic1,
    topic2: dto.theme.topic2,
    topic3: dto.theme.topic3,
  },
});

const toCreateEventDto = (eventData: EventDetails): CreateEventDto => ({
  event_start: eventData.eventStart,
  theme: eventData.theme,
  topics: eventData.topics,
});

export const fetchEvents = async (): Promise<Event[]> => {
  try {
    if (isTestMode) {
      return loadMockData<Event[]>("events");
    }
    const response = await apiRaw.get<EventDto[]>("/events");
    return response.data.map(mapEventDto);
  } catch (error) {
    throw toApiError(error, "イベントの取得に失敗しました");
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
    const payload = toCreateEventDto(eventData);
    const response = await apiRaw.post<CreateEventDto>("/events", payload);
    return {
      eventStart: response.data.event_start,
      theme: response.data.theme,
      topics: response.data.topics,
    };
  } catch (error) {
    throw toApiError(error, "イベントの作成に失敗しました");
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
    const response = await apiRaw.post<ChatThemeResponseDto>("/chat/theme", {
      content: prompt,
    });
    return response.data.choices[0].message.content;
  } catch (error) {
    throw toApiError(error, "テーマの生成に失敗しました");
  }
};
