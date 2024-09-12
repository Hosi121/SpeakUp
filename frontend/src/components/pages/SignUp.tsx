import React from 'react';
import { Box, Button, Container, TextField, Typography, Link } from '@mui/material';
import { useNavigate } from 'react-router-dom';

const SignUp = () => {
  const navigate = useNavigate();

  const handleNavigateToLogin = () => {
    navigate('/login');
  };

  return (
    <Container
      maxWidth="xs"
      sx={{
        display: 'flex',
        flexDirection: 'column',
        alignItems: 'center',
        justifyContent: 'center',
        height: '100vh',
        backgroundColor: '#FFDD66', // Matches the yellow background color
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

      {/* Username Field */}
      <TextField
        label="ユーザー名(セッション時の表示名)"
        type="text"
        fullWidth
        variant="outlined"
        margin="normal"
        defaultValue="name@example.com"
      />

      {/* Email Field */}
      <TextField
        label="Email"
        type="email"
        fullWidth
        variant="outlined"
        margin="normal"
        defaultValue="name@example.com"
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