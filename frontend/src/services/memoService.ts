import { apiRaw } from "./api";
import type { MemoDto } from "../types/dto";
import type { UserNotes } from "../types/types";
import { toApiError } from "./errorUtils";

const fromMemoDto = (dto: MemoDto): UserNotes => ({
  carryInMemo: dto.memo1 ?? "",
  wordList: dto.memo2 ?? "",
});

const toMemoDto = (notes: UserNotes): MemoDto => ({
  memo1: notes.carryInMemo,
  memo2: notes.wordList,
});

// メモを取得する関数
export const fetchMemo = async (): Promise<UserNotes> => {
  try {
    const response = await apiRaw.get<MemoDto>("/memo");
    return fromMemoDto(response.data);
  } catch (error) {
    throw toApiError(error, "メモの取得に失敗しました");
  }
};

// メモを保存する関数
export const saveMemo = async (notes: UserNotes): Promise<UserNotes> => {
  try {
    const response = await apiRaw.put<MemoDto>("/memo", toMemoDto(notes));
    return fromMemoDto(response.data);
  } catch (error) {
    throw toApiError(error, "メモの保存に失敗しました");
  }
};
