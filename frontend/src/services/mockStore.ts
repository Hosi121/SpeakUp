import conversationHistory from "../mock/conversation_history.json";
import events from "../mock/events.json";
import notifications from "../mock/notifications.json";
import sessionHistory from "../mock/session_history.json";
import sessions from "../mock/sessions.json";
import topics from "../mock/topics.json";

const SEED_VERSION = "2";
const VERSION_KEY = "speakup.seed.version";
const KEY_PREFIX = "speakup.seed.";

const seedData = {
  sessions,
  conversationHistory,
  sessionHistory,
  topics,
  notifications,
  events,
};

export type MockDataKey = keyof typeof seedData;

const canUseStorage = (): boolean =>
  typeof window !== "undefined" && typeof localStorage !== "undefined";

const getStorageKey = (key: MockDataKey) => `${KEY_PREFIX}${key}`;

export const seedMockData = (): void => {
  if (!canUseStorage()) {
    return;
  }
  if (localStorage.getItem(VERSION_KEY) === SEED_VERSION) {
    return;
  }
  localStorage.setItem(VERSION_KEY, SEED_VERSION);
  (Object.keys(seedData) as MockDataKey[]).forEach((key) => {
    localStorage.setItem(getStorageKey(key), JSON.stringify(seedData[key]));
  });
};

export const loadMockData = <T,>(key: MockDataKey): T => {
  if (!canUseStorage()) {
    return seedData[key] as T;
  }
  seedMockData();
  const raw = localStorage.getItem(getStorageKey(key));
  if (!raw) {
    return seedData[key] as T;
  }
  try {
    return JSON.parse(raw) as T;
  } catch (error) {
    console.error("Failed to parse mock data", error);
    return seedData[key] as T;
  }
};

export const saveMockData = <T,>(key: MockDataKey, value: T): void => {
  if (!canUseStorage()) {
    return;
  }
  localStorage.setItem(getStorageKey(key), JSON.stringify(value));
};
