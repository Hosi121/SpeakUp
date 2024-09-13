import { Box, Button, Container, TextField, Typography, Link } from '@mui/material';
import { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import api from '../../services/api';  // api インスタンスをインポート

const SignUp = () => {
  const navigate = useNavigate();

  // フォームの状態を管理
  const [username, setUsername] = useState('');
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');  // パスワードの状態
  const [error, setError] = useState('');

  // サインアップのPOSTリクエスト
  const handleSignUp = async () => {
    // パスワードが8文字以上か確認
    if (password.length < 8) {
      setError('パスワードは8文字以上である必要があります。');
      return;
    }

    try {
      const response = await api.post('/signup', {
        username,  // ユーザー名
        email,     // メールアドレス
        password,  // パスワード
      });

      // 成功した場合、リダイレクトまたはメッセージ表示
      console.log('サインアップ成功:', response.data);
      navigate('/login'); // サインアップ後、ログインページにリダイレクト
    } catch (err) {
      console.error('サインアップエラー:', err);
      setError('サインアップに失敗しました。もう一度お試しください。');
    }
  };

  const handleNavigateToLogin = () => {
    navigate('/login');
  };

  return (
    <Container
      maxWidth="md"
      sx={{
        display: 'flex',
        flexDirection: 'column',
        alignItems: 'center',
        justifyContent: 'center',
        height: '100vh',
        backgroundColor: '#FFDD66',
        padding: '16px',
        borderRadius: '16px',
      }}
    >
      {/* Logo */}
      <Box
        sx={{
          backgroundColor: 'white',
          width: 100,
          height: 100,
          borderRadius: '50%',
          display: 'flex',
          alignItems: 'center',
          justifyContent: 'center',
          mb: 4,
        }}
      >
        <Typography variant="h5" sx={{ fontWeight: 'bold', color: '#FF007F' }}>
          SPEAK UP!
        </Typography>
      </Box>

      {/* Sign Up Title */}
      <Typography variant="h4" component="h1" sx={{ mb: 4 }}>
        サインアップ
      </Typography>

      {/* Error Message */}
      {error && (
        <Typography color="error" sx={{ mb: 2 }}>
          {error}
        </Typography>
      )}

      {/* Username Field */}
      <TextField
        label="ユーザー名(セッション時の表示名)"
        type="text"
        fullWidth
        variant="outlined"
        margin="normal"
        value={username}
        onChange={(e) => setUsername(e.target.value)} // ユーザー名の状態を更新
      />

      {/* Email Field */}
      <TextField
        label="Email"
        type="email"
        fullWidth
        variant="outlined"
        margin="normal"
        value={email}
        onChange={(e) => setEmail(e.target.value)} // メールアドレスの状態を更新
      />

      {/* Password Field */}
      <TextField
        label="パスワード"
        type="password"
        fullWidth
        variant="outlined"
        margin="normal"
        value={password}
        onChange={(e) => setPassword(e.target.value)} // パスワードの状態を更新
      />

      {/* Submit Button */}
      <Button
        variant="contained"
        fullWidth
        sx={{
          backgroundColor: '#FF007F',
          color: 'white',
          marginY: 2,
          '&:hover': {
            backgroundColor: '#FF3399',
          },
        }}
        onClick={handleSignUp}  // ボタンクリックでサインアップリクエスト
      >
        メールを送信して仮登録
      </Button>

      {/* Navigate to Login */}
      <Link
        component="button"
        variant="body2"
        sx={{ mt: 2, cursor: 'pointer', textDecoration: 'underline' }}
        onClick={handleNavigateToLogin}
      >
        ログインページに戻る
      </Link>
    </Container>
  );
};

export default SignUp;