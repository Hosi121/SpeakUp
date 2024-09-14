import api from './api';

// サインアップ関数
export const signUp = async (username, email, password) => {
  try {
    const response = await api.post('/signup', {
      username,
      email,
      password,
    });
    return response.data;  // 成功時のレスポンスデータを返す
  } catch (err) {
    throw new Error('サインアップに失敗しました。');
  }
};

// ログイン関数
export const signIn = async (email, password) => {
  try {
    const response = await api.post('/signin', {
      email,
      password,
    });
    return response.data;  // 成功時のレスポンスデータを返す
  } catch (err) {
    throw new Error('ログインに失敗しました。');
  }
};
