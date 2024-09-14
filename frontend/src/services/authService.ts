import api from './api';

// サインアップリクエストの型
interface SignUpRequest {
  username: string;
  email: string;
  password: string;
}

// サインアップレスポンスの型
interface SignUpResponse {
  success: boolean;
  message: string;
}

// サインインリクエストの型
interface SignInRequest {
  email: string;
  password: string;
}

// サインインレスポンスの型
interface SignInResponse {
  token: string;
  user: {
    id: string;
    email: string;
    username: string;
  };
}

// サインアップ関数
export const signUp = async (username: string, email: string, password: string): Promise<SignUpResponse> => {
  const requestData: SignUpRequest = {
    username,
    email,
    password,
  };

  try {
    const response = await api.post<SignUpResponse>('/signup', requestData);
    return response.data;
  } catch (err) {
    throw new Error('サインアップに失敗しました。');
  }
};

// ログイン関数
export const signIn = async (email: string, password: string): Promise<SignInResponse> => {
  const requestData: SignInRequest = {
    email,
    password,
  };

  try {
    const response = await api.post<SignInResponse>('/signin', requestData);
    return response.data;
  } catch (err) {
    throw new Error('ログインに失敗しました。');
  }
};