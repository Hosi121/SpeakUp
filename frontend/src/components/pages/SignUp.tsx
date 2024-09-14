import { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { Box, Button, Container, TextField, Typography, Link } from '@mui/material';
import { signUp } from '../../services/authService';  // authServiceのサインアップ関数をインポート
import Logo from "../../assets/logo.tsx";

const SignUp = () => {
  const [username, setUsername] = useState('');
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');  // パスワードの状態
  const [error, setError] = useState('');
  const navigate = useNavigate();

  const handleNavigateToLogin = () => {
    navigate("/login");
  const handleSignUp = async () => {
    if (password.length < 8) {
      setError('パスワードは8文字以上である必要があります。');
      return;
    }

    try {
      const data = await signUp(username, email, password);  // authServiceの関数を使用
      console.log('サインアップ成功:', data);
      navigate('/login');  // サインアップ後、ログインページにリダイレクト
    } catch (err) {
      setError((err as Error).message);
    }
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
      <Box
        sx={{
          backgroundColor: "white",
          width: 150,
          height: 150,
          borderRadius: "50%",
          display: "flex",
          alignItems: "center",
          justifyContent: "center",
          mb: 4,
        }}
      >
        <Logo style={{ width: "80%" }} />
      </Box>

      <Typography variant="h4" component="h1" sx={{ mb: 4 }}>
        サインアップ
      </Typography>

      <TextField
        label="ユーザー名(セッション時の表示名)"
        type="text"
        fullWidth
        variant="outlined"
        margin="normal"
        value={username}
        onChange={(e) => setUsername(e.target.value)}
      />

      <TextField
        label="Email"
        type="email"
        fullWidth
        variant="outlined"
        margin="normal"
        value={email}
        onChange={(e) => setEmail(e.target.value)}
      />

      <TextField
        label="パスワード"
        type="password"
        fullWidth
        variant="outlined"
        margin="normal"
        value={password}
        onChange={(e) => setPassword(e.target.value)}
      />

      <Button
        variant="contained"
        fullWidth
        sx={{
          backgroundColor: "#FF007F",
          color: "white",
          marginY: 2,
          "&:hover": {
            backgroundColor: "#FF3399",
          },
        }}
        onClick={handleSignUp}
      >
        メールを送信して仮登録
      </Button>
      <Link
        component="button"
        variant="body2"
        sx={{ mt: 2, cursor: 'pointer', textDecoration: 'underline' }}
        onClick={() => navigate('/login')}
      >
        ログインページに戻る
      </Link>
    </Container>
  );
};

export default SignUp;
