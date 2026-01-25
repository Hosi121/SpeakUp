import api from "./api";
import type { ChatResponseDto } from "../types/dto";
import { toApiError } from "./errorUtils";

const extractMessage = (data: ChatResponseDto): string | undefined =>
  data.choices?.[0]?.message?.content;

export const askAssistant = async (content: string): Promise<string> => {
  try {
    const response = await api.post<ChatResponseDto>("/chat/ask", { content });
    const message = extractMessage(response.data);
    if (!message) {
      throw new Error("アシスタントの応答形式が不正です");
    }
    return message;
  } catch (error) {
    throw toApiError(error, "アシスタントの応答取得に失敗しました");
  }
};
