import api from "./api";

// メモを取得する関数
export const fetchMemo = async () => {
  try {
    const response = await api.get("/memo");
    return response.data;
  } catch (error) {
    console.error("Failed to fetch memos", error);
    throw error; // エラーを投げることで、呼び出し元で処理できる
  }
};

// メモを保存する関数
export const saveMemo = async (memo1: string, memo2: string) => {
  try {
    const response = await api.put("/memo", { memo1, memo2 });
    return response.data;
  } catch (error) {
    console.error("Failed to save memos", error);
    throw error; // エラーを投げることで、呼び出し元で処理できる
  }
};
