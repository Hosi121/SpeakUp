import React, { useState } from 'react';
import { Box, Button, Container, TextField, Typography, IconButton, InputAdornment } from '@mui/material';
import Visibility from '@mui/icons-material/Visibility';
import VisibilityOff from '@mui/icons-material/VisibilityOff';

const Login = () => {
  const [showPassword, setShowPassword] = useState(false);

  const handleClickShowPassword = () => setShowPassword(!showPassword);

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
        borderRadius: '16px'
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

      {/* Sign In Title */}
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
        defaultValue="name@example.com"
      />

      {/* Password Field */}
      <TextField
        label="パスワード"
        type={showPassword ? 'text' : 'password'}
        fullWidth
        variant="outlined"
        margin="normal"
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

      {/* Sign In Button */}
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
        サインイン
      </Button>

      {/* Forgot Password and Sign Up */}
      <Box
        sx={{
          display: 'flex',
          justifyContent: 'space-between',
          width: '100%',
        }}
      >
        <Typography variant="body2" color="textSecondary">
          パスワードを忘れた場合はこちら
        </Typography>
        <Typography variant="body2" color="textSecondary" sx={{ cursor: 'pointer', textDecoration: 'underline' }}>
          サインアップ
        </Typography>
      </Box>
    </Container>
  );
};

export default Login;
