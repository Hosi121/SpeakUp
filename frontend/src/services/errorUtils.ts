import axios from "axios";

type ErrorPayload = Partial<Record<"message" | "error" | "detail", string>>;

const isRecord = (value: unknown): value is Record<string, unknown> =>
  typeof value === "object" && value !== null;

const pickMessage = (data: unknown): string | undefined => {
  if (typeof data === "string") {
    return data;
  }
  if (!isRecord(data)) {
    return undefined;
  }
  const payload = data as ErrorPayload;
  return payload.message ?? payload.error ?? payload.detail;
};

export const getErrorMessage = (error: unknown, fallback: string): string => {
  if (axios.isAxiosError(error)) {
    const message = pickMessage(error.response?.data);
    return message ?? error.message ?? fallback;
  }
  if (error instanceof Error && error.message) {
    return error.message;
  }
  return fallback;
};

export type ApiError = Error & {
  status?: number;
  code?: string;
  data?: unknown;
  cause?: unknown;
};

export const toApiError = (error: unknown, fallback: string): ApiError => {
  const message = getErrorMessage(error, fallback);
  const normalized = new Error(message) as ApiError;
  normalized.cause = error;
  if (axios.isAxiosError(error)) {
    normalized.status = error.response?.status;
    normalized.code = error.code;
    normalized.data = error.response?.data;
  }
  return normalized;
};
