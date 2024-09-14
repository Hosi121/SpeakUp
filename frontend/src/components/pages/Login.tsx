import { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { Box, Button, Container, TextField, Typography, IconButton, InputAdornment } from '@mui/material';
import Visibility from '@mui/icons-material/Visibility';
import VisibilityOff from '@mui/icons-material/VisibilityOff';
import { signIn } from '../../services/authService';
import Logo from "../../assets/logo.tsx";        

const Login = () => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [showPassword, setShowPassword] = useState(false);
  const [error, setError] = useState('');
  const navigate = useNavigate();

  const handleClickShowPassword = () => setShowPassword(!showPassword);

  // Function to handle navigation to the sign-up page
  const handleNavigateToSignup = () => {
    navigate("/signup");
  };

  // ログイン処理
  const handleLogin = async () => {
    try {
      const data = await signIn(email, password);  // authServiceの関数を使用
      console.log('ログイン成功:', data);
      navigate('/home');  // ログイン成功後にリダイレクト
    } catch (err) {
      setError((err as Error).message);  // エラーメッセージを設定
    }
  };

  return (
    <Container
      maxWidth="md"
      sx={{
        display: "flex",
        flexDirection: "column",
        alignItems: "center",
        justifyContent: "center",
        height: "100vh",
        backgroundColor: "#FFDD66",
        padding: "16px",
        borderRadius: "16px",
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
        サインイン
      </Typography>

      {/* Email Field */}
      <TextField 
        label="Email" 
        type="email" 
        fullWidth 
        variant="outlined" 
        margin="normal"
        value={email}
        onChange={(e) => setEmail(e.target.value)}
      />

      {/* Error message */}
      {error && (
        <Typography color="error" sx={{ mb: 2 }}>
          {error}
        </Typography>
      )}

      {/* Password Field */}
      <TextField
        label="パスワード"
        type={showPassword ? "text" : "password"}
        fullWidth
        variant="outlined"
        margin="normal"
        value={password}
        onChange={(e) => setPassword(e.target.value)}
        InputProps={{
          endAdornment: (
            <InputAdornment position="end">
              <IconButton onClick={handleClickShowPassword}>
                {showPassword ? <VisibilityOff /> : <Visibility />}
              </IconButton>
            </InputAdornment>
          ),
        }}
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
        onClick={handleLogin}  // ログイン処理を実行
      >
        サインイン
      </Button>

      {/* Forgot Password and Sign Up */}
      <Box
        sx={{
          display: "flex",
          justifyContent: "space-between",
          width: "100%",
        }}
      >
        <Typography variant="body2" color="textSecondary">
          パスワードを忘れた場合はこちら
        </Typography>

        {/* Clickable Sign-Up text */}
        <Typography
          variant="body2"
          color="textSecondary"
          sx={{ cursor: "pointer", textDecoration: "underline" }}
          onClick={handleNavigateToSignup} // OnClick handler to navigate to /signup
        >
          サインアップ
        </Typography>
      </Box>
    </Container>
  );
};

export default Login;