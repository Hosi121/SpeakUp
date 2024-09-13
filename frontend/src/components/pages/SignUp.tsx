import { Box, Button, Container, TextField, Typography, Link } from "@mui/material";
import { useNavigate } from "react-router-dom";
import Logo from "../../assets/logo.tsx";

const SignUp = () => {
  const navigate = useNavigate();

  const handleNavigateToLogin = () => {
    navigate("/login");
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
        backgroundColor: "#FFDD66", // Matches the yellow background color
        padding: "10%",
      }}
    >
      {/* Logo */}
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

      {/* Sign Up Title */}
      <Typography variant="h4" component="h1" sx={{ mb: 4 }}>
        サインアップ
      </Typography>

      {/* Username Field */}
      <TextField label="ユーザー名(セッション時の表示名)" type="text" fullWidth variant="outlined" margin="normal" />

      {/* Email Field */}
      <TextField label="Email" type="email" fullWidth variant="outlined" margin="normal" />

      {/* Submit Button */}
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
      >
        メールを送信して仮登録
      </Button>

      {/* Navigate to Login */}
      <Link component="button" variant="body2" sx={{ mt: 2, cursor: "pointer", textDecoration: "underline" }} onClick={handleNavigateToLogin}>
        ログインページに戻る
      </Link>
    </Container>
  );
};

export default SignUp;
