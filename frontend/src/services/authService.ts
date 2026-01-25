import { apiRaw } from './api';
import { isTestMode } from "./appMode";
import { toApiError } from "./errorUtils";

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
    const response = await apiRaw.post<SignUpResponse>('/signup', requestData);
    return response.data;
  } catch (error) {
    throw toApiError(error, "サインアップに失敗しました。");
  }
};

// ログイン関数
export const signIn = async (email: string, password: string): Promise<void> => {
  const requestData: SignInRequest = {
    email,
    password,
  };

  try {
    if (isTestMode) {
      localStorage.setItem("token", "test-token");
      return;
    }
    const response = await apiRaw.post<SignInResponse>('/signin', requestData);
    // Use the correct 'token' from SignInResponse
    localStorage.setItem("token", response.data.token);
  } catch (error) {
    throw toApiError(error, "ログインに失敗しました。");
  }
};
